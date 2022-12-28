import React from "react"
import { Routes, Route } from "react-router-dom"
import SignIn from "../pages/signin/SignIn"
import SignUp from "../pages/signup/SignUp"
import Home from "../pages/home/Home"

const AllRoutes = () => {
  return (
    <div>
      <div>
        <Routes>
          <Route path="signin" element={<SignIn />} />
          <Route path="signup" element={<SignUp />} />
          <Route path="/" element={<Home />} />
        </Routes>
      </div>
    </div>
  )
}

export default AllRoutes
