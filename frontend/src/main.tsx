import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { TodoProvider } from "./contexts/TodoContext";
import { Router } from "./router";
import "./index.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <TodoProvider>
      <Router />
    </TodoProvider>
  </StrictMode>
);
