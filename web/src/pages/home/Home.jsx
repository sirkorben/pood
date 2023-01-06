import React from "react"
import { useEffect } from "react"
import { useContext } from "react"
import { Link } from "react-router-dom"
import { SearchContext } from "../../components/utils/SearchContext"
import { SignInForm } from "../signin/SignIn"
import styles from "./Home.module.scss"
const Home = () => {
  const { setResults } = useContext(SearchContext)
  /*  setResults([]) */
  /*  useEffect(() => {
    setResults([])
  }, []) */
  return <div className={styles.home}>Hey</div>
}

export default Home
