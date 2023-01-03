import axios from "axios"
import { useState } from "react"
import { Link } from "react-router-dom"
import { ToastContainer, toast } from "react-toastify"
import "./SignIn.scss"
import { local_backend_ip } from "../../App"

const SignIn = () => {
  return (
    <div>
      <Form />
      <ToastContainer />
    </div>
  )
}

const Form = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()
    const res = axios
      .post(
        `${local_backend_ip}/signin`,  // for dev use ${local_backend_ip} in containers/prod use ${process.env.REACT_APP_BACKEND_URL}
        JSON.stringify({
          email,
          password,
        }),
        {
          headers: { "Content-Type": "application/json" },
          withCredentials: true,
        }
      )
      .catch((error) => {
        toast.error(error.response["data"]["error_description"], {
          position: "top-right",
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "dark",
        })
      })
    console.log(res)
  }

  return (
    <div className="formContainer">
      <div className="formWrapper">
        <span className="title">Sign In</span>
        <form onSubmit={handleSubmit}>
          <input
            required
            placeholder="email"
            type="email"
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            placeholder="password"
            required
            type="password"
            onChange={(e) => setPassword(e.target.value)}
          />
          <button>Sign in</button>
        </form>
        <p>
          or <Link to={"/signup"}>Sign Up</Link>
        </p>
      </div>
    </div>
  )
}

export default SignIn
