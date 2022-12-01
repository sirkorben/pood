import axios from "axios"
import { useState } from "react"
import { Link } from "react-router-dom"

const SignInPage = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()

    const res = axios
      .post(
        "http://localhost/8080/signin",
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
        console.log(error)
      })
    console.log(res)
  }

  return (
    <div>
      <div>
        <form onSubmit={handleSubmit}>
          <label>Email</label>
          <input
            required
            type="email"
            onChange={(e) => setEmail(e.target.value)}
          />
          <br />
          <label>Password</label>
          <input
            required
            type="text"
            onChange={(e) => setPassword(e.target.value)}
          />
          <br />
          <button>Sign in</button>
          <br />
          or <Link to={"/signup"}>Sign Up</Link>
        </form>
      </div>
    </div>
  )
}

export default SignInPage
