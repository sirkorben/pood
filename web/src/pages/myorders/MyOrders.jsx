import React from "react"
import { useState } from "react"
import { useEffect } from "react"
import { useParams, Link } from "react-router-dom"
import "../../styles/orders-list.scss"
import myAxios from "../../components/utils/api/axios"
import styles2 from "../../styles/cart-orders-styles.module.scss"

const MyOrders = () => {
  const [orders, setOrders] = useState([])
  useEffect(() => {
    myAxios.get("/api/myorders").then((res) => {
      //console.log(res.data)
      setOrders(res.data.orders)
    })
  }, [])

  if (orders === null) return <div className="no_task">You have no orders</div>

  return (
    <div className="orders_wrapper">
      <h2> My orders</h2>
      <div className="orders_card">
        {orders?.map((order) => (
          <div className="orders_list" key={order.order_id}>
            <span className="orders_list__order_id">Order ID</span>:{" "}
            <Link className="orders_list__link" to={`${order.order_id}`}>
              {order.order_id}
            </Link>
          </div>
        ))}
      </div>
    </div>
  )
}

export const MyOrder = () => {
  const { id } = useParams()
  const [order, setOrder] = useState({})
  useEffect(() => {
    myAxios
      .get(`/api/myorders/order?id=${id}`)

      .then((res) => {
        console.log(res.data)
        setOrder(res.data)
      })
      .catch((err) => {
        console.log(err)
      })
  }, [id])
  console.log("ORDER ", order)

  return (
    <div className={styles2.cart_wrapper}>
      <div className={styles2.cart_wrapper__total_cart_info}>
        <div className={styles2.total_cart_info__info}>
          <ul>
            <li>
              Order ID: <span className="mark">{order.order_id}</span>
            </li>
            <li>
              Total items:{" "}
              <span className="mark">{order.positions?.length}</span>
            </li>
            <li>
              Total price:{" "}
              <span className="mark">
                {order.total_price?.toFixed(2)}&euro;
              </span>
            </li>
          </ul>
        </div>
      </div>
      <div className={styles2.cart_wrapper__cart_items}>
        {order.positions?.map((product, index) => (
          <div
            className={styles2.cart_wrapper__cart_items__item_card}
            key={index}
          >
            <ul>
              <li>
                Article: <span className="mark">{product.article}</span>
              </li>
              <li>Brand: {product.brand}</li>
              <li>Currency: {product.currency}</li>
              <li>Currency rate: {product.currency_rate}</li>
              <li>Delivery: {product.delivery}</li>
              <li>Position_id: {product.position_id}</li>
              <li>
                Price:{" "}
                <span className="mark">{product.price?.toFixed(2)}&euro;</span>
              </li>
              <li>
                Product quantity price:{" "}
                <span className="mark">
                  {product.product_quantity_price?.toFixed(2)}&euro;
                </span>
              </li>
              <li>
                Quantity: <span className="mark">{product.quantity}</span>
              </li>
              <li>
                Supplier: <span className="mark">{product.supplier}</span>
              </li>
              {product.weight ? (
                <li>Weight: {product.weight?.toFixed(2)}kg</li>
              ) : null}
            </ul>
          </div>
        ))}
      </div>
    </div>
  )

  /*  return (
    <div className="order_wrapper">
      {order && (
        <div className="order_card">
          <h2>Order ID: {order.order_id}</h2>
          {order.positions?.map((pos) => (
            <div className="position_card" key={pos.position_id}>
              <ul>
                <li>Article: {pos.article}</li>
                <li>Brand: {pos.brand}</li>
                <li>Delivery: {pos.delivery}</li>
                <li>Price: {pos.price.toFixed(2)}&euro;</li>
                <li>
                  Product quantity price:{" "}
                  {pos.product_quantity_price.toFixed(2)}&euro;
                </li>
                <li>Quantity: {pos.quantity}</li>
                <li>Supplier: {pos.supplier}</li>
                {pos.weight ? (
                  <li>Weight: {pos.weight.toFixed(2)} kg</li>
                ) : null}
              </ul>
            </div>
          ))}

          <div className="total_order_price">
            Total order price: {order.total_price?.toFixed(2)}&euro;
          </div>
        </div>
      )}
    </div>
  ) */
}

export default MyOrders
