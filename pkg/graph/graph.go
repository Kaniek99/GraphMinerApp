package graph

type GraphFileSchema struct {
	Graph  *Graph   `json:"graph,omitempty"`
	Graphs []*Graph `json:"graphs,omitempty"`
}

type Graph struct {
	ID       string          `json:"id,omitempty"`
	Label    string          `json:"label,omitempty"`
	Directed bool            `json:"directed,omitempty"`
	Type     string          `json:"type,omitempty"`
	Metadata map[string]any  `json:"metadata,omitempty"`
	Nodes    map[string]Node `json:"nodes,omitempty"`
	Edges    []Edge          `json:"edges,omitempty"`
}

type Node struct {
	Data map[string]any `json:"metadata,omitempty"`
	ID   string         `json:"label"`
}

type Edge struct {
	Data     map[string]any `json:"metadata,omitempty"`
	Directed bool           `json:"directed,omitempty"`
	ID       string         `json:"id"`
	Label    string         `json:"label"`
	Source   string         `json:"source"`
	Target   string         `json:"target"`
}

func (graph *Graph) PrintGraph() {
	println("Graph ID:", graph.ID)
	println("Graph Label:", graph.Label)
	println("Directed:", graph.Directed)
	println("Type:", graph.Type)

	if len(graph.Nodes) > 0 {
		println("Nodes:")
		for _, node := range graph.Nodes {
			println("  Node ID:", node.ID)
		}
	} else {
		println("No nodes found.")
	}

	if len(graph.Edges) > 0 {
		println("Edges:")
		for _, edge := range graph.Edges {
			println("  Edge ID:", edge.ID)
			println("    Source:", edge.Source)
			println("    Target:", edge.Target)
			for key, value := range edge.Data {
				println("    ", key+":", value)
			}
		}
	} else {
		println("No edges found.")
	}
}
