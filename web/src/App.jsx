import "./App.css"
import AllRoutes from "./components/Routes"
import { useContext } from "react"
import { UserContext } from "./components/utils/UserContext"
import Login from "./pages/login/Login"
import { ToastContainer } from "react-toastify"
import Main from "./layout/Main"

export const local_backend_ip = "http://localhost:8080" //http://146.190.118.167:8080/

function App() {
  const { logged } = useContext(UserContext)

  return (
    <div>
      {logged ? (
        <div>
          <Main />
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
