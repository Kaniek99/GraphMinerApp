package parser

import (
	"GraphMinerApp/pkg/graph"
	"encoding/json"
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

type Parser interface {
	Parse(filePath string) (*graph.GraphFileSchema, error)
	Validate(filePath string) (bool, error)
}

type JSONParser struct {
	schemaLoader *gojsonschema.JSONLoader
}

func NewJSONParser(schemaPath string) (*JSONParser, error) {
	schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)

	return &JSONParser{schemaLoader: &schemaLoader}, nil
}

func (jp *JSONParser) Parse(filePath string) (*graph.GraphFileSchema, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed reading from %s: %v", filePath, err)
	}

	var graphSchema graph.GraphFileSchema
	err = json.Unmarshal(data, &graphSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &graphSchema, nil
}

func (jp *JSONParser) Validate(filePath string) (bool, error) {
	documentLoader := gojsonschema.NewReferenceLoader("file:///" + filePath)
	_, err := gojsonschema.Validate(*jp.schemaLoader, documentLoader)
	if err != nil {
		return false, err
	}
	return true, nil
}

type GraphMLParser struct{}
