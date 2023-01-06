import axios from "axios"
import React from "react"
import Header from "../components/header/Header"
import Products from "../components/products/Products"

const Main = () => {
  /*  const [results, setResults] = useState */
  const SearchRequest = async () => {
    await axios.get()
  }

  return (
    <div>
      <Header />
      <div>
        <Products />
      </div>
    </div>
  )
}

export default Main
