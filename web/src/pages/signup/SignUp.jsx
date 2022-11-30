import { useEffect } from "react"
import { useState } from "react"
import axios from "axios"

const SignUpPage = () => {
  const [firstname, setFirstname] = useState("")
  const [lastname, setLastname] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [confirm, setConfirm] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      axios.post(
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
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <div>
      <div>
        <form onSubmit={handleSubmit}>
          <label>First name</label>
          <input
            required
            type="text"
            onChange={(e) => setFirstname(e.target.value)}
          />
          <br />
          <label>Last name</label>
          <input
            required
            type="text"
            onChange={(e) => setLastname(e.target.value)}
          />
          <br />
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
          <label>Confirm password</label>
          <input
            required
            type="text"
            onChange={(e) => setConfirm(e.target.value)}
          />
          <br />
          <button>Sign Up</button>
        </form>
      </div>
    </div>
  )
}

export default SignUpPage
