import React from "react"
import { useState } from "react"
import styles from "./Search.module.scss"

import { useContext } from "react"
import { SearchContext } from "../../components/utils/SearchContext"
import { useNavigate } from "react-router-dom"
import Products from "../../components/products/Products"
import { useEffect } from "react"

/* const SearchBar = () => {
  return (
    <div className={styles.searchContainer}>
      <Input />
    </div>
  )
} */

const SearchPage = () => {
  const { setResults } = useContext(SearchContext)
  useEffect(() => {
    setResults([])
  }, [])
  return (
    <div className={styles.search_wrapper}>
      <Input />
      <Products />
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
    <div className={styles.search}>
      <input
        placeholder="Enter article number"
        type="text"
        onChange={(e) => setArticle(e.target.value)}
      />
      <button onClick={handleClick}>Search</button>
    </div>
  )
}

export default SearchPage
