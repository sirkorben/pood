import React from "react"
import { useState } from "react"
import styles from "./Search.module.scss"

import { useContext } from "react"
import { SearchContext } from "../../components/utils/SearchContext"
import { useNavigate } from "react-router-dom"

const SearchBar = () => {
  return (
    <div className={styles.searchContainer}>
      <Input />
    </div>
  )
}

const Input = () => {
  const [article, setArticle] = useState()
  const nav = useNavigate()
  const { SearchRequest } = useContext(SearchContext)

  const handleClick = async (e) => {
    e.preventDefault()
    SearchRequest(article)
  }

  return (
    <div>
      <div className={styles.search}>
        <input type="text" onChange={(e) => setArticle(e.target.value)} />
        <button onClick={handleClick}>Search</button>
      </div>
    </div>
  )
}

export default SearchBar
