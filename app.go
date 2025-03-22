package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// UploadFile handles file uploads from the frontend
func (a *App) UploadFile(fileData map[string]any) (string, error) {
	fmt.Println(fileData)
	// fh.ValidateJSONFile(fileData["path"].(string))
	return "TBI", nil
}

func (a *App) OpenFileDialog() (string, error) {
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
	return filePath, err
}

func (a *App) GetDotGraph() (string, error) {
	// return "digraph { a -> b; b -> c; c -> d; d -> e; e -> f; f -> g; g -> h; h -> i; i -> j; j -> k; k -> l; l -> m; m -> n; n -> o; o -> p; p -> q; q -> r; r -> s; s -> t; t -> u; u -> v; v -> w; w -> x; x -> y; y -> z; a -> e; a -> i; a -> m; b -> f; b -> j; b -> n; c -> g; c -> k; c -> o; d -> h; d -> l; d -> p; z -> a; y -> b; x -> c; w -> d; v -> e; u -> f; m -> a; n -> c; o -> e; p -> g; q -> i; r -> k; s -> m; t -> o; a -> z -> y -> x -> a; b -> v -> u -> t -> b; c -> s -> r -> q -> c; aa; bb; cc;}", nil
	return "digraph K6 { graph [  layout = circo,];	a -> b; a -> c; a -> d; a -> e; a -> f; b -> a; b -> c; b -> d; b -> e; b -> f; c -> a; c -> b; c -> d; c -> e; c -> f; d -> a; d -> b; d -> c; d -> e; d -> f; e -> a; e -> b; e -> c; e -> d; e -> f; f -> a; f -> b; f -> c; f -> d; f -> e; }", nil
}
