import React from "react"
import { useState, useContext } from "react"
import { useNavigate, Link } from "react-router-dom"
import { UserContext } from "../../components/utils/UserContext"
import axios from "axios"
import { ToastContainer, toast } from "react-toastify"
import "./Login.scss"
import { backend_url } from "../../App"
import myAxios from "../../components/utils/api/axios"
import { options } from "../../components/utils/toast/options"

const Login = () => {
  return (
    <div className="formContainer">
      <Forms />
      <ToastContainer />
    </div>
  )
}

const Forms = () => {
  const { isSignUp } = useContext(UserContext)
  return <div>{isSignUp ? <SignUpForm /> : <SignInForm />}</div>
}

const SignInForm = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const { logged, setLogged } = useContext(UserContext)
  const { isSignUp, setIsSignUp } = useContext(UserContext)
  /*  console.log(email)
  console.log(password) */
  // console.log(logged)
  const navigate = useNavigate()
  const handleSubmit = async (e) => {
    e.preventDefault()
    const res = await myAxios
      .post("/api/signin", JSON.stringify({ email, password }))
      .then(() => {
        navigate("/me")
        setLogged(true)
      })
      .catch((error) => {
        console.log("error ", error)
        setLogged(false)
        toast.error(error.response["data"]["error_description"], options)
      })
    console.log(res)
  }

  return (
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
        or <span onClick={() => setIsSignUp(!isSignUp)}>Sign Up</span>
      </p>
    </div>
  )
}

const SignUpForm = () => {
  const [firstname, setFirstname] = useState("")
  const [lastname, setLastname] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [confirm, setConfirm] = useState("")
  const [registered, setRegistered] = useState(false)
  //const [err, setErr] = useState(null)
  const { isSignUp, setIsSignUp } = useContext(UserContext)

  const handleSubmit = async (e) => {
    e.preventDefault()

    const res = await myAxios
      .post(
        "/api/signup",
        JSON.stringify({
          firstname,
          lastname,
          email,
          password,
          confirmed_password: confirm,
        })
      )

      .then(() => {
        setRegistered(true)
        return <ApproveMessage />
      })
      .catch((error) => {
        toast.error(error.response?.data["error_description"], options)
      })
    //setRegistered(true)
    console.log("res ", res)
  }
  if (registered) return <ApproveMessage className="approve" />

  return (
    <div className="formWrapper">
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
        <input
          type="password"
          placeholder="Confirm password"
          onChange={(e) => setConfirm(e.target.value)}
        />
        <button>Sign Up</button>
      </form>
      <p>
        Already have an account?{" "}
        <span className="redirect" onClick={() => setIsSignUp(!isSignUp)}>
          Sign In
        </span>
      </p>
    </div>
  )
}

const ApproveMessage = () => {
  return (
    <div className="approve">
      <p>Your application will be reviewed by the site administrator.</p>
    </div>
  )
}

export default Login
