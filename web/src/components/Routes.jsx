import React from "react"
import { Routes, Route } from "react-router-dom"

import Home from "../pages/home/Home"

import Products from "./products/Products"

const AllRoutes = () => {
  return (
    <div>
      <div>
        <Routes>
          {/*  <Route path="signin" element={<SignIn />} />
          <Route path="signup" element={<SignUp />} /> */}
          {/* <Route path="products" element={<Products />} /> */}
          <Route path="/home" element={<Home />} />
          {/* <Route path="/search" element={<SearchBar />} /> */}
        </Routes>
      </div>
    </div>
  )
}

export default AllRoutes
