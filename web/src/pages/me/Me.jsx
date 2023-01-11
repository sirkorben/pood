import React from "react"
import { useContext, useEffect } from "react"
import { SearchContext } from "../../components/utils/SearchContext"
import { UserContext } from "../../components/utils/UserContext"
import styles from "./Me.module.scss"
const Me = () => {
  const { me } = useContext(UserContext)
  const { setResults } = useContext(SearchContext)
  useEffect(() => {
    setResults([])
  }, [])
  console.log(me)
  return (
    <div className={styles.me_wrapper}>
      {me && (
        <div className={styles.me_card}>
          <h2>
            {me.firstname} {me.lastname}
          </h2>
          <div>{me.email}</div>
        </div>
      )}
    </div>
  )
}

export default Me
