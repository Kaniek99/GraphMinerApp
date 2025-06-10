package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	gr "GraphMinerApp/pkg/graph"
	"GraphMinerApp/pkg/parser"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

/*
 **********************************************************
 *                                                        *
 *                                                        *
 *  Reading this code is painful, refactor required ASAP  *
 *                                                        *
 *                                                        *
 **********************************************************
 */

type InputFile struct {
	Extension string
	Path      string
	Valid     bool

	Parser parser.Parser
}

func newInputFile(filePath string) (*InputFile, error) {
	index := strings.LastIndex(filePath, ".")
	if index == -1 {
		return nil, fmt.Errorf("invalid file path: %s", filePath)
	}
	extension := strings.ToLower(filePath[index+1:])

	// Validation in this case should not be more complicated so it is a part of creation
	if extension != "json" && extension != "graphml" {
		return nil, fmt.Errorf("unsupported file extension: %s", extension)
	}

	return &InputFile{Extension: extension, Path: filePath}, nil
}

func (inf *InputFile) setParser(parser parser.Parser) {
	inf.Parser = parser
}

func (inf *InputFile) Validate() error {
	if inf.Parser == nil {
		return fmt.Errorf("parser not set")
	}

	isValid, err := inf.Parser.Validate(inf.Path)
	if err != nil {
		return err
	}

	inf.Valid = isValid
	return nil
}

// App struct
type App struct {
	bufferWithGraph        *bytes.Buffer
	InputGraphs            []graph.Graph[string, string]
	InputGraphsDotLanguage []string

	ctx context.Context
	inf *InputFile
	jp  *parser.JSONParser
	gs  *gr.GraphFileSchema
	// gp  *parser.GraphMLParser
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dirName, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get current working directory:", err)
	}

	jp, err := parser.NewJSONParser("file:///" + dirName + "/schemas/graph_schema.json")
	if err != nil {
		fmt.Println(err)
	}
	a.jp = jp
}

func (a *App) ChooseInputFile() error {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "JSON Files (*.json)",
				Pattern:     "*.json",
			},
			{
				DisplayName: "GraphML Files (*.graphml)",
				Pattern:     "*.graphml",
			},
		},
	})
	if err != nil {
		fmt.Println("cannot open file dialog: ", err)
	}
	fmt.Println("File Path: ", filePath)

	file, err := newInputFile(filePath)
	if err != nil {
		fmt.Println("Error while creating InputFile struct: ", err)
	}

	file.setParser(a.jp)
	a.inf = file

	fmt.Println("Input File: ", a.inf)

	// isValid, err := a.ValidateInputFile()
	err = a.inf.Validate()
	if err != nil {
		fmt.Println("Error with validation: ", err)
	}

	fmt.Println("valid: ", a.inf.Valid)

	graphSchema, err := a.inf.Parser.Parse(a.inf.Path)
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	// clear old data
	a.InputGraphs = nil
	a.InputGraphsDotLanguage = nil

	a.gs = graphSchema

	if graphSchema.Graph != nil {
		graphSchema.Graph.PrintGraph()

		inputGraph := graph.New(graph.StringHash)
		for _, node := range graphSchema.Graph.Nodes {
			inputGraph.AddVertex(node.Label)
		}
		for _, edge := range graphSchema.Graph.Edges {
			inputGraph.AddEdge(edge.Source, edge.Target)
		}
		a.InputGraphs = append(a.InputGraphs, inputGraph)

		inputDotGraph, err := a.CreateDotGraph(inputGraph)
		if err != nil {
			fmt.Println("Error while creating DOT graph:", err)
			return err
		}

		a.InputGraphsDotLanguage = append(a.InputGraphsDotLanguage, inputDotGraph)
	} else if len(graphSchema.Graphs) > 0 {
		for _, graphTest := range graphSchema.Graphs {
			graphTest.PrintGraph()

			inputGraph := graph.New(graph.StringHash)
			for _, node := range graphTest.Nodes {
				inputGraph.AddVertex(node.Label)
			}
			for _, edge := range graphTest.Edges {
				inputGraph.AddEdge(edge.Source, edge.Target)
			}
			a.InputGraphs = append(a.InputGraphs, inputGraph)

			inputDotGraph, err := a.CreateDotGraph(inputGraph)
			if err != nil {
				fmt.Println("Error while creating DOT graph:", err)
				return err
			}

			a.InputGraphsDotLanguage = append(a.InputGraphsDotLanguage, inputDotGraph)
		}
	} else {
		fmt.Println("No graph found in the file")
	}

	return nil
}

func (a *App) CreateDotGraph(g graph.Graph[string, string]) (string, error) {
	// make this parameter to be set by the user
	layout := "circo"
	buf := bytes.NewBuffer(nil)

	// not sure if this should be saved as a temporary file or maybe it can be stored in the buffer
	// tmp, _ := os.Create("./simple.gv")
	err := draw.DOT(g, buf)
	if err != nil {
		fmt.Println("Error while creating DOT graph: ", err)
		return "", err
	}

	// Get the generated DOT string
	dotString := buf.String()

	// to change layout try to use Attributes in description
	// https://github.com/dominikbraun/graph/blob/main/draw/draw.go
	if strings.Contains(dotString, "digraph {") {
		dotString = strings.Replace(dotString, "digraph {", fmt.Sprintf("digraph {\n\tgraph [layout=%s];", layout), 1)
	} else if strings.Contains(dotString, "graph {") {
		dotString = strings.Replace(dotString, "graph {", fmt.Sprintf("graph {\n\tgraph [layout=%s];", layout), 1)
	}

	a.bufferWithGraph = bytes.NewBufferString(dotString)

	// fmt.Println("DOT graph: ", buf.String())
	return buf.String(), nil
}

func (a *App) GetDotGraphs() ([]string, error) {
	// if a.bufferWithGraph == nil {
	// return []string{""}, fmt.Errorf("no graph created")
	// }
	for _, dotGraph := range a.InputGraphsDotLanguage {
		fmt.Println("Get DOT graph: ", dotGraph)
	}
	return a.InputGraphsDotLanguage, nil
}

func (a *App) ValidateInputFile() (bool, error) {
	err := a.inf.Validate()
	if err != nil {
		return false, err
	}

	fmt.Println("Valid: ", a.inf.Valid)
	return true, nil
}
