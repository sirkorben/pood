import axios from "axios"
import React from "react"
import { useContext } from "react"
import { local_backend_ip } from "../../App"
import { CartContext } from "../../components/utils/CartContext"
import styles from "./Cart.module.scss"

const Cart = () => {
  const { cart } = useContext(CartContext)
  console.log(cart)

  const handleRemove = async (order_id) => {
    await axios.delete(`${local_backend_ip}/cart/remove`, {
      headers: { "Content-Type": "application/json" },
      withCredentials: true,
      data: JSON.stringify({
        order_id,
      }),
    })
  }
  return (
    <div className={styles.cart_wrapper}>
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
            </div>
          ))}
          <label>Total price: {cart.total_price?.toFixed(2)}&euro;</label>
        </div>
      )}
      <button onClick={() => handleRemove(cart.order_id)}>Remove cart</button>
    </div>
  )
}

export default Cart
