import axios from "axios"
import { createContext, useEffect, useState } from "react"
import { local_backend_ip } from "../../App"

export const CartContext = createContext()

export const CartContextProvider = (props) => {
  const [cart, setCart] = useState({})
  const [added, setAdded] = useState(false)
  const [itemsInCart, setItemsInCart] = useState(0)

  console.log(added)
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/cart`, { withCredentials: true })
      .then((res) => {
        console.log(res)
        setCart(res.data)
        setItemsInCart(res.data.products?.length)
      })
  }, [added])
  /* console.log("q ", itemsInCart) */
  const value = { cart, added, setAdded, itemsInCart }
  return (
    <CartContext.Provider value={value}>{props.children}</CartContext.Provider>
  )
}
