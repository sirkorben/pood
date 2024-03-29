import { createContext, useContext, useEffect, useState } from "react"
import myAxios from "./api/axios"
import { UserContext } from "./UserContext"

export const CartContext = createContext()

export const CartContextProvider = (props) => {
  const [cart, setCart] = useState({})
  const [added, setAdded] = useState(false)
  const [removed, setRemoved] = useState(false)
  const [confirmed, setConfirmed] = useState(false)

  const [itemsInCart, setItemsInCart] = useState(0)
  const { logged } = useContext(UserContext)
  //console.log(added)
  useEffect(() => {
    myAxios.get("/api/cart").then((res) => {
      setCart(res.data)
      setItemsInCart(res.data.products?.length)
    })
  }, [added, logged, removed, confirmed])
  /* console.log("q ", itemsInCart) */
  const value = { cart, added, setAdded, itemsInCart, setRemoved, setConfirmed }
  return (
    <CartContext.Provider value={value}>{props.children}</CartContext.Provider>
  )
}
