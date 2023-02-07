import React from "react"
import styles from "./Header.module.scss"
import { useContext } from "react"
import { UserContext } from "../utils/UserContext"
import { useNavigate, Link } from "react-router-dom"
import { SearchContext } from "../utils/SearchContext"
import { CartContext } from "../utils/CartContext"
import myAxios from "../utils/api/axios"

const Header = () => {
  const { logged } = useContext(UserContext)
  const { itemsInCart } = useContext(CartContext)
  if (!logged) return null

  return (
    <div className={styles.headerWrapper}>
      <div className={styles.headerWrapper__title}>
        Parts<span className={styles.hub}>hub</span>
      </div>
      <nav>
        <ul className={styles.nav_list}>
          <li>
            <Link className={styles.nav_list__item} to="/me">
              Me
            </Link>
          </li>
          <li>
            <Link className={styles.nav_list__item} to="/search">
              Search
            </Link>
          </li>
          <li>
            <Link className={styles.nav_list__item} to="/myorders">
              My orders
            </Link>
          </li>
          <li>
            <Link className={styles.nav_list__item} to="/cart">
              Cart
            </Link>
          </li>
          {itemsInCart ? (
            <li className={styles.quantity}>{itemsInCart}</li>
          ) : null}
        </ul>
      </nav>
      <SignOutButton />
    </div>
  )
}

/* const Header = () => {
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
} */

const SignOutButton = () => {
  const nav = useNavigate()
  const { setLogged } = useContext(UserContext)
  const { setResults } = useContext(SearchContext)

  const handleClick = () => {
    myAxios.get("/api/signout").then(() => {
      setLogged(false)
      setResults([])
      nav("/")
    })
  }
  return (
    <div className={styles.button}>
      <button onClick={handleClick}>Logout</button>
    </div>
  )
}

export default Header
