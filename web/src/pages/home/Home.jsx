import React from "react"
import { Link } from "react-router-dom"

const Home = () => {
  return (
    <div>
      <div>
        <Link to={"/signin"}>Sign In</Link>
      </div>
      <div>
        <Link to={"signup"}>Sign Up</Link>
      </div>
    </div>
  )
}

export default Home
