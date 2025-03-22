# GraphMiner
GraphMiner is an application which can be used for extracting structural patterns and valuable insights from
graph-structured data. It takes a JSON or GraphML file as input, validates it, and visualizes it. The user
can then choose which operations they want to perform on the graph.

## Planed features
- [ ] JSON schema validator
- [ ] GraphML schema validator
- [ ] JSON to graph structure
- [ ] GraphML to graph structure
- [ ] Visualize graphs
- [ ] gSPAN implementation
- [ ] FSM implementation

## Prerequisites
- Wails 2.10+
- Go 1.23+
- Node.js 22+

## Getting started
You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

To build a redistributable, production mode package, use ` wails build`.

## Sequence diagram of app usage
```mermaid
sequenceDiagram
User->>GraphMinerApp GUI: insert a file
GraphMinerApp GUI->>GraphMinerApp GUI: check the file extension
alt extension is not .json
    GraphMinerApp GUI->>User: display an error to the user
else extension is valid
    GraphMinerApp GUI->>GraphMinerApp backend: pass the file to the backend
    GraphMinerApp backend->>GraphMinerApp backend: validate the file content
    alt content is invalid
        GraphMinerApp backend->>GraphMinerApp GUI: return an error
        GraphMinerApp GUI->>User: display an error to the user
    else content is valid
        GraphMinerApp backend->>GraphMinerApp GUI: return the input graph described in DOT
        GraphMinerApp GUI->>GraphMinerApp GUI: create visualization of the input file
        GraphMinerApp GUI->>User: display the visualization to the user
        User->>GraphMinerApp GUI: select an operation to perform on the graph
        GraphMinerApp GUI->>GraphMinerApp backend: trigger the operation
        GraphMinerApp backend->>GraphMinerApp backend: execute the operation
        GraphMinerApp backend->>GraphMinerApp GUI: return the result
        GraphMinerApp GUI->>GraphMinerApp GUI: create visualization of the result
        GraphMinerApp GUI->>User: display the result visualization
    end
end
```

## Used technologies
Application is created with [wails.io](https://wails.io/).

<div>
  <img src="https://github.com/wailsapp/wails/blob/master/assets/images/logo_cropped.png" title="wails" alt="wails" width="40" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/go/go-original-wordmark.svg" title="go" alt="go" width="40" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/typescript/typescript-original.svg" title="typescript" alt="typescript" width="40" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/vitejs/vitejs-original.svg" title="vitejs" alt="vitejs" width="40" height="40"/>&nbsp;
</div>
