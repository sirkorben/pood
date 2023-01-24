import React from "react"
import styles from "./Header.module.scss"
import SearchBar from "../../pages/search/Search"
import { useContext } from "react"
import { UserContext } from "../utils/UserContext"
import { useNavigate, Link } from "react-router-dom"
import axios from "axios"
import { SearchContext } from "../utils/SearchContext"
import { CartContext } from "../utils/CartContext"
/* const Header = () => {
  const { logged } = useContext(UserContext)
  return (
    <div className={styles.header}>
      {logged === false ? null : (
        <div className={styles.header__content}>
          <div>
            <span className={styles.logo}>4H0615108F</span>
            <br></br>
            <span className={styles.logo}>4M2820160</span>

          </div>
          <div>
            {" "}
            <nav className={styles.nav}>
              <div className={styles.nav__item}>
                <SearchBar />
              </div>
           
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
} */

const Header = () => {
  const { logged } = useContext(UserContext)
  const { itemsInCart } = useContext(CartContext)
  return (
    <div className={styles.header_wrapper}>
      {logged === false ? null : (
        <div>
          <h1 className={styles.title}>PartsHub</h1>
          <nav className={styles.nav}>
            <ul className={styles.nav__items}>
              <li>
                <Link className={styles.nav__link} to="/cart">
                  Cart
                </Link>
                {itemsInCart ? (
                  <div className={styles.quantity}>{itemsInCart}</div>
                ) : null}
              </li>
              <li>
                <Link className={styles.nav__link} to="/search">
                  Search
                </Link>
              </li>
              <li>
                <Link className={styles.nav__link} to={"/myorders"}>
                  My orders
                </Link>
              </li>
              <li>
                <Link className={styles.nav__link} to="/me">
                  Me
                </Link>
              </li>
            </ul>
          </nav>
          <SignOutButton />
        </div>
      )}
    </div>
  )
}

const SignOutButton = () => {
  const nav = useNavigate()
  const { setLogged } = useContext(UserContext)
  const { setResults } = useContext(SearchContext)

  const handleClick = () => {
    axios
      .get(`/api/signout`, { withCredentials: true })
      .then(() => {
        setLogged(false)
        setResults([])
        nav("/")
      })
  }
  return (
    <div className={styles.button}>
      <button onClick={handleClick}>Sign Out</button>
    </div>
  )
}

export default Header
