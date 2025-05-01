import { useState, useRef } from "react";
import { Graphviz } from "graphviz-react";
import "./App.css";
import {
  ChooseInputFile,
  GetDotGraph,
  ValidateInputFile,
} from "../wailsjs/go/main/App";

function App() {
  const [file, setFile] = useState<File | null>(null);
  const [fileError, setFileError] = useState<string | null>(null);
  const [data, setData] = useState<string>("");
  const [uploadSuccess, setUploadSuccess] = useState<boolean>(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const ALLOWED_EXTENSIONS = [".json", ".graphml"];

  const handleUpload = async (event: React.FormEvent) => {
    event.preventDefault();

    if (!file) {
      setFileError("Please select a file first");
      return;
    }

    try {
      // const filePath = await ChooseInputFile();
      await ChooseInputFile();
      await ValidateInputFile();
      setUploadSuccess(true);
      setFile(null);
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }
    } catch (error) {
      setFileError("Upload failed. Please try again.");
      console.error("Upload error:", error);
    }
  };

  const triggerFileInput = async () => {
    ChooseInputFile();
    const res = await GetDotGraph();
    setData(res);
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

  // TODO: fix the flow on subimt and think about success and error after choosing a file
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
              <button
                type="button"
                className="button"
                onClick={triggerFileInput}
              >
                Choose Input File
              </button>
              <button
                type="button"
                className="button"
                onClick={() => alert("gSpan clicked!")}
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
