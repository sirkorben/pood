import { useEffect } from "react"
import { useState } from "react"
import axios from "axios"
import { Link } from "react-router-dom"
import { ToastContainer, toast } from "react-toastify"
import "react-toastify/dist/ReactToastify.css"

const SignUpPage = () => {
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
      .then(() => setRegistered(true))
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
    <div>
      <div>
        {registered ? (
          <h2>Вашая заявка будет рассмотрена администратором сайта</h2>
        ) : (
          <div>
            <form onSubmit={handleSubmit}>
              <label>First name</label>
              <input
                type="text"
                onChange={(e) => setFirstname(e.target.value)}
              />
              <br />
              <label>Last name</label>
              <input
                type="text"
                onChange={(e) => setLastname(e.target.value)}
              />
              <br />
              <label>Email</label>
              <input type="email" onChange={(e) => setEmail(e.target.value)} />
              <br />
              <label>Password</label>
              <input
                type="text"
                onChange={(e) => setPassword(e.target.value)}
              />
              <br />
              <label>Confirm password</label>
              <input type="text" onChange={(e) => setConfirm(e.target.value)} />
              <br />
              <button>Sign Up</button>
              <br />
              or <Link to={"/signin"}>Sign In</Link>
            </form>
          </div>
        )}
      </div>
    </div>
  )
}

export default SignUpPage
