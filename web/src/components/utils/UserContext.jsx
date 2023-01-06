import { useState, useEffect } from "react"
import { createContext } from "react"
import axios from "axios"

export const UserContext = createContext()

export const UserContextProvider = (props) => {
  const [logged, setLogged] = useState(false)
  const [isSignUp, setIsSignUp] = useState(false)

  // console.log(logged)
  useEffect(() => {
    axios
      .get("http://localhost:8080/me", { withCredentials: true })
      .then((res) => {
        setLogged(true)
      })
      .catch((err) => {
        setLogged(false)
        console.log("ERROR! ", err)
      })
  }, [logged])

  const value = { logged, setLogged, isSignUp, setIsSignUp }

  return (
    <UserContext.Provider value={value}>{props.children}</UserContext.Provider>
  )
}
