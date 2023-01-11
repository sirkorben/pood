import { useState, useEffect } from "react"
import { createContext } from "react"
import axios from "axios"
import { local_backend_ip } from "../../App"
export const UserContext = createContext()

export const UserContextProvider = (props) => {
  const [logged, setLogged] = useState(false)
  const [isSignUp, setIsSignUp] = useState(false)
  const [me, setMe] = useState({})

  // console.log(logged)
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/me`, { withCredentials: true })
      .then((res) => {
        setMe(res.data)
        setLogged(true)
      })
      .catch((err) => {
        setLogged(false)
        console.log("ERROR! ", err)
      })
  }, [logged])

  const value = { logged, setLogged, isSignUp, setIsSignUp, me }

  return (
    <UserContext.Provider value={value}>{props.children}</UserContext.Provider>
  )
}
