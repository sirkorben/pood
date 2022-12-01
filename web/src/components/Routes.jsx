import { Routes, Route } from "react-router-dom"
import SignUpPage from "../pages/signup/SignUp"
import SignInPage from "../pages/signin/SignIn"
import HomePage from "../pages/home/Home"

const AllRoutes = () => {
  return (
    <div>
      <div>
        <Routes>
          <Route path="signin" element={<SignInPage />} />
          <Route path="signup" element={<SignUpPage />} />
          <Route path="/" element={<HomePage />} />
        </Routes>
      </div>
    </div>
  )
}

export default AllRoutes
