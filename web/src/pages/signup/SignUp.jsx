import { useEffect } from "react"
import { useState } from "react"
import axios from "axios"
import { Link } from "react-router-dom"
import { ToastContainer, toast } from "react-toastify"
import "react-toastify/dist/ReactToastify.css"

const SignUp = () => {
  return (
    <div>
      <Form />
      <ToastContainer />
    </div>
  )
}

const Form = () => {
  const [firstname, setFirstname] = useState("")
  const [lastname, setLastname] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [confirm, setConfirm] = useState("")
  const [registered, setRegistered] = useState(false)
  const [err, setErr] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault()

    const res = axios
      .post(
        "http://localhost:8080/signup",
        JSON.stringify({
          firstname,
          lastname,
          email,
          password,
        }),
        {
          headers: { "Content-Type": "application/json" },
          withcredentials: true,
        }
      )
      .then(() => {
        setRegistered(true)
        return <ApproveMessage />
      })
      .catch((error) => {
        toast.error(error.response["data"]["error_description"], {
          position: "top-right",
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "light",
        })
      })
    //setRegistered(true)
    console.log("res ", res)
  }
  console.log("err ", err)
  return (
    <div className="formContainer">
      <div className="formWrapper">
        {registered ? (
          <ApproveMessage className="approve" />
        ) : (
          <div>
            {" "}
            <span className="title">Sign Up</span>
            <form onSubmit={handleSubmit}>
              <input
                type="text"
                placeholder="Firstname"
                onChange={(e) => setFirstname(e.target.value)}
              />
              <input
                type="text"
                placeholder="Lastname"
                onChange={(e) => setLastname(e.target.value)}
              />

              <input
                type="email"
                placeholder="Email"
                onChange={(e) => setEmail(e.target.value)}
              />

              <input
                id="password"
                type="password"
                placeholder="Password"
                onChange={(e) => setPassword(e.target.value)}
              />
              <input type="password" placeholder="Confirm password" />
              <button>Sign Up</button>
            </form>
            <p>
              Already have an account? <Link to={"/signin"}>Login</Link>
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

const ApproveMessage = () => {
  return (
    <div className="approve">
      <p>Ваша заявка будет рассмотрена администратором сайта</p>
    </div>
  )
}

export default SignUp
