import { useState, useEffect } from "react"
import { createContext } from "react"
import axios from "axios"
import { backend_url } from "../../App"
import myAxios from "./api/axios"
export const UserContext = createContext()

export const UserContextProvider = (props) => {
  const [logged, setLogged] = useState(false)
  const [isSignUp, setIsSignUp] = useState(false)
  const [me, setMe] = useState({})

  // console.log(logged)
  useEffect(() => {
    myAxios
      .get("/api/me")
      .then((res) => {
        setMe(res.data)
        setLogged(true)
      })
      .catch((err) => {
        setLogged(false)
        console.log("ERROR!", err)
      })
  }, [logged])

  const value = { logged, setLogged, isSignUp, setIsSignUp, me }

  return (
    <UserContext.Provider value={value}>{props.children}</UserContext.Provider>
  )
}
