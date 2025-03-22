import { useState, useRef } from "react";
import { Graphviz } from "graphviz-react";
import "./App.css";
import {
  UploadFile,
  GetDotGraph,
  OpenFileDialog,
} from "../wailsjs/go/main/App";

function App() {
  const [file, setFile] = useState<File | null>(null);
  const [fileError, setFileError] = useState<string | null>(null);
  const [data, setData] = useState<string>("");
  const [isUploading, setIsUploading] = useState<boolean>(false);
  const [uploadSuccess, setUploadSuccess] = useState<boolean>(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const ALLOWED_EXTENSIONS = [".json", ".graphml"];

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setFileError(null);
    setUploadSuccess(false);

    if (event.target.files && event.target.files.length > 0) {
      const selectedFile = event.target.files[0];

      const fileExtension =
        "." + selectedFile.name.split(".").pop()?.toLowerCase();
      if (!ALLOWED_EXTENSIONS.includes(fileExtension)) {
        setFileError(
          `File type not allowed. Allowed types: ${ALLOWED_EXTENSIONS.join(", ")}`,
        );
        return;
      }

      setFile(selectedFile);
    }
  };

  const handleUpload = async (event: React.FormEvent) => {
    event.preventDefault();

    if (!file) {
      setFileError("Please select a file first");
      return;
    }

    setIsUploading(true);

    try {
      const filePath = await OpenFileDialog();
      UploadFile({ name: file.name, path: filePath });

      setUploadSuccess(true);
      setFile(null);
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }
    } catch (error) {
      setFileError("Upload failed. Please try again.");
      console.error("Upload error:", error);
    } finally {
      setIsUploading(false);
    }
  };

  const triggerFileInput = async () => {
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

  return (
    <div className="home-page">
      <div className="upload-container">
        <h1>File Upload</h1>
        <p className="description">
          Select a file to upload to the application
        </p>

        <form onSubmit={handleUpload}>
          <div className="file-input-container">
            <input
              type="file"
              onChange={handleFileChange}
              className="file-input"
              ref={fileInputRef}
            />
            <button
              type="button"
              className="browse-button"
              onClick={triggerFileInput}
            >
              Visualize graph
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

          {fileError && <div className="error-message">{fileError}</div>}
          {uploadSuccess && (
            <div className="success-message">File uploaded successfully!</div>
          )}

          <button
            type="submit"
            className="upload-button"
            disabled={!file || isUploading}
          >
            {isUploading ? "Uploading..." : "Upload File"}
          </button>
        </form>

        <div className="upload-info">
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
