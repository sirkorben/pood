import React from "react"
import { useContext, useEffect } from "react"
import { SearchContext } from "../../components/utils/SearchContext"
import { UserContext } from "../../components/utils/UserContext"
import styles from "./Me.module.scss"
import { Link } from "react-router-dom"

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
          {me.is_admin === 1 ? (
            <div>
              <Link to={"/admin"}>admin panel</Link>
            </div>
          ) : null}
        </div>
      )}
    </div>
  )
}

export default Me
