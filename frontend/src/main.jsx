import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";  // Optional: You can add a global CSS file for styling

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

