import axios from "axios"
import { createContext, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { backend_url } from "../../App"
import myAxios from "./api/axios"
export const SearchContext = createContext()

/* const PARTS_API_KEY =
  "f22d9a0a6a65a45e5fe9cd652deb0e98e9051b286d36709ae12d1900da516c10"
const APP_LINK = "https://originalparts.pro/api/search?" */
const SearchContextProvider = (props) => {
  const [results, setResults] = useState([])

  const SearchRequest = (article) => {
    const response = myAxios
      .post("/api/search", JSON.stringify({ article }))
      .then((res) => setResults(res.data))

    console.log(response)
  }

  //console.log("results ", results)

  const value = { results, setResults, SearchRequest }

  return (
    <SearchContext.Provider value={value}>
      {props.children}
    </SearchContext.Provider>
  )
}

export default SearchContextProvider
