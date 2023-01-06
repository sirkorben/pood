import React from "react"
import styles from "./Header.module.scss"
import SearchBar from "../../pages/search/Search"
import { useContext } from "react"
import { UserContext } from "../utils/UserContext"
import { useNavigate } from "react-router-dom"
import axios from "axios"
import { local_backend_ip } from "../../App"
const Header = () => {
  const { logged } = useContext(UserContext)
  return (
    <div className={styles.header}>
      {logged === false ? null : (
        <div className={styles.header__content}>
          <div>
            <span className={styles.logo}>4H0615108F</span>
          </div>
          <div>
            {" "}
            <nav className={styles.nav}>
              <div className={styles.nav__item}>
                <SearchBar />
              </div>
              {/*   <div className={styles.nav__button__container}>
                <SignOutButton />
              </div> */}
            </nav>
          </div>
          <div>
            <div className={styles.header__button__container}>
              <SignOutButton />
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

const SignOutButton = () => {
  const nav = useNavigate()
  const { setLogged } = useContext(UserContext)
  const handleClick = () => {
    axios
      .get(`${local_backend_ip}/signout`, { withCredentials: true })
      .then(() => {
        setLogged(false)
        nav("/")
      })
  }
  return (
    <div>
      <button onClick={handleClick}>Sign Out</button>
    </div>
  )
}

export default Header
