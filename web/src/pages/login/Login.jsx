import React from "react"
import { useState, useContext } from "react"
import { useNavigate, Link } from "react-router-dom"
import { UserContext } from "../../components/utils/UserContext"
import axios from "axios"
import { local_backend_ip } from "../../App"
import { ToastContainer, toast } from "react-toastify"
import "./Login.scss"

const Login = () => {
  return (
    <div>
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
    const res = await axios
      .post(
        `${local_backend_ip}/api/signin`,
        JSON.stringify({
          email,
          password,
        }),
        {
          headers: { "Content-Type": "application/json" },
          withCredentials: true,
        }
      )
      .then(() => {
        navigate("/me")
        setLogged(true)
      })
      .catch((error) => {
        console.log("error ", error)
        setLogged(false)
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
          or <span onClick={() => setIsSignUp(!isSignUp)}>Sign Up</span>
        </p>
      </div>
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
  const [err, setErr] = useState(null)
  const { isSignUp, setIsSignUp } = useContext(UserContext)

  const handleSubmit = async (e) => {
    e.preventDefault()

    const res = await axios
      .post(
        `${local_backend_ip}/api/signup`, // for dev use ${local_backend_ip} in containers/prod use ${process.env.REACT_APP_BACKEND_URL}
        JSON.stringify({
          firstname,
          lastname,
          email,
          password,
          confirmed_password: confirm,
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
        toast.error(error.response?.data["error_description"], {
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
  // console.log("err ", err)
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
        )}
      </div>
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
