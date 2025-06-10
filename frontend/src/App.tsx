import { useState } from "react";
import { Graphviz } from "graphviz-react";
import "./App.css";
import { ChooseInputFile, GetDotGraphs } from "../wailsjs/go/main/App";

function App() {
  const [file, setFile] = useState<File | null>(null);
  const [fileError, setFileError] = useState<string | null>(null);
  const [data, setData] = useState<string>("");
  const [uploadSuccess, setUploadSuccess] = useState<boolean>(false);
  const [currentGraphIndex, setCurrentGraphIndex] = useState<number>(0);
  const [graphs, setGraphs] = useState<string[]>([]);

  const ALLOWED_EXTENSIONS = [".json", ".graphml"];

  const handleUpload = async (event: React.FormEvent) => {
    event.preventDefault();

    try {
      await ChooseInputFile();
      const arrayOfGraphs = await GetDotGraphs();
      setGraphs(arrayOfGraphs);

      if (graphs.length > 0) {
        setCurrentGraphIndex(0);
        setData(arrayOfGraphs[currentGraphIndex]);
      }

      setUploadSuccess(true);
    } catch (error) {
      setFileError("Upload failed. Please try again.");
      console.error("Upload error:", error);
    }
  };

  function Item({ data }: { data: string }) {
    if (data) {
      return (
        <Graphviz
          className="graph-visualization"
          options={{ height: 800, width: 800, zoom: true }}
          dot={data}
        />
      );
    }
    return "";
  }

  const handlePrevious = () => {
    if (currentGraphIndex < 1 || graphs.length === 0) return;

    const newIndex = currentGraphIndex - 1;
    setCurrentGraphIndex(newIndex);
    setData(graphs[newIndex]);
  };

  const handleNext = () => {
    if (graphs.length === 0) return;

    const newIndex = (currentGraphIndex + 1) % graphs.length;
    setCurrentGraphIndex(newIndex);
    setData(graphs[newIndex]);
  };

  return (
    <div className="home-page">
      <div className="upload-container">
        <h1>File Upload</h1>
        <p className="description">
          Select a file to upload to the application
        </p>

        <form onSubmit={handleUpload}>
          <div className="file-input-container">
            <div className="button-row">
              <button type="submit" className="button">
                Choose Input File
              </button>
              <button
                type="button"
                className="button"
                onClick={() => alert("gSpan clicked!")}
                // onClick={PerformgSpan}
              >
                gSpan
              </button>
              <button
                type="button"
                className="button"
                onClick={() => alert("FFSM clicked!")}
              >
                FFSM
              </button>
              {file && (
                <div className="selected-file">
                  <span className="file-name" style={{ color: "black" }}>
                    {file.name}
                  </span>
                  <span className="file-size" style={{ color: "black" }}>
                    ({(file.size / 1024).toFixed(1)} KB)
                  </span>
                </div>
              )}
            </div>
          </div>

          <div className="navigation-container">
            <div className="button-row">
              <button type="button" className="button" onClick={handlePrevious}>
                Previous
              </button>
              <button type="button" className="button" onClick={handleNext}>
                Next
              </button>
            </div>
          </div>

          {fileError && <div className="error-message">{fileError}</div>}
          {uploadSuccess && (
            <div className="success-message">File uploaded successfully!</div>
          )}
        </form>

        <div className="input-file-info">
          <p style={{ color: "black" }}>
            <strong>Allowed file types:</strong> {ALLOWED_EXTENSIONS.join(", ")}
          </p>
        </div>
      </div>
      <div className="wrapper">
        <Item data={data} />
      </div>
    </div>
  );
}

export default App;
