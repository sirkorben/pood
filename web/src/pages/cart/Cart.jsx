import React from "react"
import { useContext } from "react"
import { CartContext } from "../../components/utils/CartContext"
import styles from "../../styles/cart-orders-styles.module.scss"
import { toast } from "react-toastify"
import myAxios from "../../components/utils/api/axios"
import { options } from "../../components/utils/toast/options"

const Cart = () => {
  const { cart, setRemoved, setConfirmed, itemsInCart } =
    useContext(CartContext)
  // console.log(cart)
  //const [orderConfirmed, setOrderConfirmed] = useState(false)
  const handleRemove = async (order_id) => {
    setRemoved(true)
    await myAxios
      .delete("/api/cart/remove", { data: JSON.stringify({ order_id }) })
      .then(() => {
        toast.info("Cart was removed", options)
      })
      .catch((err) => {
        toast.error(err.response["data"]["error_description"], options)
      })
    setRemoved(false)
    //console.log("RES ", res)
  }

  const handleConfirm = async () => {
    setConfirmed(true)

    await myAxios
      .post("/api/cart/confirm", null)
      .then((res) => {
        // console.log(res)
        //setOrderConfirmed(true)
        toast.info("Your order was placed", options)
      })
      .catch((err) => {
        setConfirmed(false)
        toast.error(err.response["data"]["error_description"], options)
      })
    setConfirmed(false)
  }
  const handleItemRemove = async (position_id) => {
    setRemoved(true)
    await myAxios
      .delete("/api/cart/removeitem", { data: JSON.stringify({ position_id }) })

      .then(() => {
        toast.info("Item was removed", options)
      })

    setRemoved(false)
  }

  if (cart.products === null)
    return <div className="no_task">Your cart is empty now</div>

  return (
    <div className={styles.cart_wrapper}>
      <div className={styles.cart_wrapper__total_cart_info}>
        <div className={styles.total_cart_info__info}>
          <ul>
            {/* <li>Order ID: {cart.order_id}</li> */}
            <li>
              Total items: <span className="mark">{itemsInCart}</span>
            </li>
            <li>
              Total price:{" "}
              <span className={styles.price}>
                {cart.total_price?.toFixed(2)}&euro;
              </span>
            </li>
          </ul>
        </div>
        <div className={styles.buttons}>
          <button className={styles.buttons__button} onClick={handleConfirm}>
            Confirm order
          </button>
          <button
            className={styles.buttons__button}
            onClick={() => handleRemove(cart.order_id)}
          >
            Remove order
          </button>
        </div>
      </div>
      <div className={styles.cart_wrapper__cart_items}>
        {cart.products?.map((product, index) => (
          <div
            className={styles.cart_wrapper__cart_items__item_card}
            key={index}
          >
            <ul>
              <li>
                Article: <span className="mark">{product.article}</span>
              </li>
              <li>
                Brand: <span className="mark">{product.brand}</span>
              </li>
              <li>Currency: {product.currency}</li>
              <li>Currency rate: {product.currency_rate}</li>
              <li>Delivery: {product.delivery}</li>
              <li>Position_id: {product.position_id}</li>
              <li>
                Price:{" "}
                <span className={styles.mark}>
                  {product.price?.toFixed(2)}&euro;
                </span>
              </li>
              <li>
                Product quantity price:{" "}
                <span className={styles.mark}>
                  {product.product_quantity_price?.toFixed(2)}&euro;
                </span>
              </li>
              <li>
                Quantity:{" "}
                <span className={styles.mark}>{product.quantity}</span>
              </li>
              <li>
                Supplier: <span className="mark">{product.supplier}</span>
              </li>
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
      </div>
    </div>
  )
}
//TODO - understand how to implement it
/* const PlacedOrderMessage = ({ order_id }) => {
  return (
    <div>
      <div>Your order was successfuly placed.</div>
      <div>
        Check my order: <Link to={`/${order_id}`}>{order_id}</Link>
      </div>
    </div>
  )
} */

export default Cart
