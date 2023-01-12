import React from "react"
import { useState } from "react"
import { useContext } from "react"
import { SearchContext } from "../utils/SearchContext"
import styles from "./Products.module.scss"
import axios from "axios"
import { local_backend_ip } from "../../App"
import { CartContext } from "../utils/CartContext"
import { toast } from "react-toastify"
/* "price": 16.46,
"article": "352316171195",
"supplier": "Magneti Marelli",
"supplier_price_num": 180,
"brand": "magneti marelli",
"currency": "euro",
"currency_rate": "90",
"delivery": "21",
"weight": 0.7,
"name": "" */

/* {
  "price": 822.906,
  "article": "4H0615108F",
  "supplier": "Bonn",
  "supplier_price_num": 192,
  "brand": "vw",
  "currency": "euro",
  "currency_rate": "85",
  "delivery": "10",
  "weight": 7.31,
  "quantity": 7
} */
const Products = () => {
  const { results } = useContext(SearchContext)
  const [quantity, setQuantity] = useState(0)
  const { setAdded } = useContext(CartContext)

  console.log(results)
  const handleAdd = async (result) => {
    console.log(quantity)
    /* setAdded(true) */

    if (quantity !== 0) {
      axios
        .post(
          `${local_backend_ip}/cart/add`,
          JSON.stringify({
            price: result.price,
            article: result.article,
            supplier: result.supplier,
            brand: result.brand,
            currency: result.currency,
            currency_rate: result.currency_rate,
            delivery: result.delivery,
            weight: result.weight,
            quantity: Number(quantity),
          }),
          {
            withCredentials: true,
            headers: { "Content-Type": "application/json" },
          }
        )
        .then(() => {
          setAdded(true)
          Array.from(document.querySelectorAll("input")).forEach(
            (input) => (input.value = "")
          )

          toast.success("Item added to cart", {
            position: "top-right",
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "dark",
          })
        })
        .catch((err) => {
          setAdded(false)
          console.log(err)
        })
      setAdded(false)
      setQuantity(0)
    } else {
      toast.warn("Enter amount!", {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
      })
    }
  }

  /*   const handleChange = (e, id) => {
    
    setQuantity(e.target.value)
    console.log(quantity)
  } */

  return (
    <div className={styles.products_wrapper}>
      {results.prices?.map((result, index) => (
        <div key={index} className={styles.product_card}>
          <h3 className={styles.title}>{result.article}</h3>
          <div className={styles.product_card__items}>
            <ul>
              <li>Article: {result.article}</li>
              <li>Price: {result.price.toFixed(2)}&euro;</li>
              <li>Supplier: {result.supplier}</li>
              <li>Supplier_price_num: {result.supplier_price_num}</li>
              <li>Brand: {result.brand}</li>
              <li>Currency: {result.currency}</li>
              <li>Currency rate: {result.currency_rate}</li>
              <li>Delivery: {result.delivery}</li>
              {result.weight ? (
                <li>Weight: {result.weight?.toFixed(2)} kg</li>
              ) : null}
              {result.name ? <li>Name: {result.name}</li> : null}
            </ul>
          </div>
          <div className={styles.product_card__quantity}>
            <div className={styles.product_card__quantity__input}>
              {/* <button onClick={() => handleMinus(index)}>-</button> */}
              <input
                id="quantity-input"
                type="text"
                placeholder="0"
                //value={quantity.item_id === index ? quantity.quantity : null}
                /* value={result.quantity} */
                onChange={(e) => {
                  setQuantity(e.target.value)
                  // console.log(quantity)
                }}
              />

              {/* <button onClick={() => handlePlus(result, index)}>+</button> */}
            </div>
            <button
              onClick={() => handleAdd(result)}
              className={styles.add_to_cart_button}
            >
              Add to cart
            </button>
          </div>
        </div>
      ))}
    </div>
  )
}

export default Products
