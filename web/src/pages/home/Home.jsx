import { Link } from "react-router-dom"

const HomePage = () => {
  return (
    <div>
      <div>
        <Link to={"/signin"}>Login</Link>
      </div>
      <div>
        <Link to={"signup"}>Sign Up</Link>
      </div>
    </div>
  )
}

export default HomePage
