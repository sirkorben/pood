import axios from "axios"
import React from "react"
import { useState } from "react"
import { useEffect } from "react"
import { useParams, Link } from "react-router-dom"
import { local_backend_ip } from "../../App"
import styles from "./MyOrders.module.scss"

const MyOrders = () => {
  const [myOrders, setMyOrders] = useState([])
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/myorders`, { withCredentials: true })
      .then((res) => {
        console.log(res.data)
        setMyOrders(res.data.orders)
      })
  }, [])

  //console.log(myOrders.orders?.map((order) => console.log(order)))

  return (
    <div className={styles.orders_wrapper}>
      {myOrders ? (
        <div className={styles.orders_card}>
          <h2>Your orders</h2>
          {myOrders?.map((order) => (
            <div className={styles.orders_list} key={order.order_id}>
              Order ID: <Link to={`${order.order_id}`}>{order.order_id}</Link>
            </div>
          ))}
        </div>
      ) : (
        <div>You have no orders</div>
      )}
    </div>
  )
}

export const MyOrder = () => {
  const { id } = useParams()
  const [order, setOrder] = useState({})
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/order?id=${id}`, { withCredentials: true })
      .then((res) => {
        console.log(res.data)
        setOrder(res.data)
      })
  }, [id])
  return (
    <div className={styles.order_wrapper}>
      {order && (
        <div className={styles.order_card}>
          <h2>Order ID: {order.order_id}</h2>
          {order.positions?.map((pos) => (
            <div className={styles.position_card} key={pos.position_id}>
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
                {pos.weight ? <li>Weight: {pos.weight} kg</li> : null}
              </ul>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default MyOrders
