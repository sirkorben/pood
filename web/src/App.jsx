import logo from "./logo.svg"
import "./App.css"
import AllRoutes from "./components/Routes"
import { Link } from "react-router-dom"

export const local_backend_ip = "http://localhost:8080"

function App() {
  return (
    <div>
      <AllRoutes />
    </div>
  )
}

export default App
