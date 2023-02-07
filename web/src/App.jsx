import "./App.css"
import AllRoutes from "./components/Routes"
import { useContext } from "react"
import { UserContext } from "./components/utils/UserContext"
import Login from "./pages/login/Login"
import { ToastContainer } from "react-toastify"
import "react-toastify/dist/ReactToastify.css"
import Header from "./components/header/Header"

export const backend_url = "http://localhost:8080"

function App() {
  const { logged } = useContext(UserContext)

  return (
    <div>
      {logged ? (
        <div className="main-container">
          <Header />
          <AllRoutes />
        </div>
      ) : (
        <Login />
      )}
      <ToastContainer />
    </div>
  )
}

export default App
