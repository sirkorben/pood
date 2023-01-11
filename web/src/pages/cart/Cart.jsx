import axios from "axios"
import React from "react"
import { useContext } from "react"
import { local_backend_ip } from "../../App"
import { CartContext } from "../../components/utils/CartContext"
import styles from "./Cart.module.scss"
import { toast } from "react-toastify"

const Cart = () => {
  const { cart, setRemoved, setConfirmed } = useContext(CartContext)
  console.log(cart)

  const handleRemove = async (order_id) => {
    setRemoved(true)
    await axios
      .delete(`${local_backend_ip}/cart/remove`, {
        headers: { "Content-Type": "application/json" },
        withCredentials: true,
        data: JSON.stringify({
          order_id,
        }),
      })
      .then(() => {
        toast.info("Cart was removed", {
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
        toast.error(err.response["data"]["error_description"], {
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
    setRemoved(false)
  }

  const handleConfirm = async () => {
    setConfirmed(true)

    await axios
      .post(`${local_backend_ip}/cart/confirm`, null, { withCredentials: true })
      .then((res) => {
        console.log(res)

        toast.info("Your order was placed", {
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
        setConfirmed(false)
        toast.error(err.response["data"]["error_description"], {
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
    setConfirmed(false)
  }
  const handleItemRemove = async (position_id) => {
    setRemoved(true)
    await axios
      .delete(`${local_backend_ip}/cart/removeitem`, {
        headers: { "Content-Type": "application/json" },
        withCredentials: true,
        data: JSON.stringify({
          position_id,
        }),
      })
      .then(() => {
        toast.info("Item was removed", {
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

    setRemoved(false)
  }
  return (
    <div className={styles.cart_wrapper}>
      {cart.products !== null ? (
        <div>
          {cart && (
            <div className={styles.cart}>
              <h2>Order ID: {cart.order_id}</h2>
              {cart.products?.map((product, index) => (
                <div key={index} className={styles.cart__items_cards}>
                  <ul>
                    <li>Article: {product.article}</li>
                    <li>Brand: {product.brand}</li>
                    <li>Currency: {product.currency}</li>
                    <li>Currency rate: {product.currency_rate}</li>
                    <li>Delivery: {product.delivery}</li>
                    <li>Position_id: {product.position_id}</li>
                    <li>Price: {product.price?.toFixed(2)}&euro;</li>
                    <li>
                      Product quantity price:{" "}
                      {product.product_quantity_price?.toFixed(2)}&euro;
                    </li>
                    <li>Quantity: {product.quantity}</li>
                    <li>Supplier: {product.supplier}</li>
                    {product.weight ? (
                      <li>Weight: {product.weight?.toFixed(2)}kg</li>
                    ) : null}
                  </ul>
                  <button
                    onClick={() => handleItemRemove(product.position_id)}
                    className={styles.remove_button}
                  >
                    Remove item
                  </button>
                </div>
              ))}
              <label className={styles.total_price}>
                Total price: {cart.total_price?.toFixed(2)}&euro;
              </label>
            </div>
          )}

          <button
            onClick={() => handleRemove(cart.order_id)}
            className={styles.remove_button}
          >
            Remove cart
          </button>
          <button onClick={handleConfirm} className={styles.confirm_button}>
            Confirm order
          </button>
        </div>
      ) : (
        <div>Your cart is empty now</div>
      )}
    </div>
  )
}

export default Cart
