import axios from "axios"
import React from "react"
import { useState } from "react"
import styles from "./Admin.module.scss"
import { local_backend_ip } from "../../App"
import { useEffect } from "react"
import { Link, useParams } from "react-router-dom"
import { useContext } from "react"
import { UserContext } from "../../components/utils/UserContext"
import "../../styles/orders-list.scss"

const AdminPage = () => {
  const { me } = useContext(UserContext)
  return (
    <div>
      {me.is_admin === 1 ? (
        <div className={styles.admin}>
          <Link to={"/admin/approve"}>ADMIN APPROVE</Link>
          <Link to={"/admin/managepercent"}>ADMIN MANAGE PERCENT</Link>
          <Link to={"/admin/orders"}>ADMIN MANAGE ORDERS</Link>
        </div>
      ) : (
        <div>Not allowed</div>
      )}
    </div>
  )
}

export const AdminApprove = () => {
  const { me } = useContext(UserContext)

  const [users, setUsers] = useState([])
  const [approve, setApprove] = useState(false)
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/admin/approve`, { withCredentials: true })
      .then((res) => {
        console.log(res.data)
        setUsers(res.data)
        setApprove(false)
      })
  }, [approve])

  const handleApprove = async (user) => {
    setApprove(true)
    axios.patch(
      `${local_backend_ip}/admin/approve`,
      JSON.stringify({
        id: user.id,
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
        activated: user.activated,
        date_created: user.date_created,
      }),
      { withCredentials: true, headers: { "Content-Type": "application/json" } }
    )
  }
  /*  "id": 3,
  "firstname": "Artem",
  "lastname": "Non-active",
  "email": "tema@bravo.com",
  "activated": 0,
  "date_created": 1672740186 */

  return (
    <div className={styles.admin_approve_wrapper}>
      {me.is_admin === 1 ? (
        <div>
          <span>ADMIN APPROVE</span>
          <div className={styles.users}>
            {users ? (
              users?.map((user) => (
                <div key={user.id} className={styles.users__card}>
                  <div>
                    <label>Name: {user.firstname}</label>
                  </div>
                  <div>
                    <label>ID: {user.id}</label>
                  </div>
                  <button
                    className={styles.button}
                    onClick={() => handleApprove(user)}
                  >
                    approve
                  </button>
                </div>
              ))
            ) : (
              <div>No users to approve</div>
            )}
          </div>
        </div>
      ) : (
        <div>Not allowed</div>
      )}
    </div>
  )
}

export const AdminManagePercent = () => {
  const { me } = useContext(UserContext)

  const [users, setUsers] = useState([])
  const [manage, setManage] = useState(false)
  const [percent, setPercent] = useState(1.15)
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/admin/managepercent`, { withCredentials: true })
      .then((res) => {
        console.log(res.data)
        setUsers(res.data)
        setManage(false)
      })
  }, [manage])

  const handleManage = async (user) => {
    setManage(true)
    axios
      .patch(
        `${local_backend_ip}/admin/managepercent`,
        JSON.stringify({
          id: user.id,
          user_percent: Number(percent),
        }),
        {
          withCredentials: true,
          headers: { "Content-Type": "application/json" },
        }
      )
      .catch((err) => {
        console.log(err)
      })
    Array.from(document.querySelectorAll("input")).forEach(
      (input) => (input.value = "")
    )
    setPercent(1.15)
  }

  return (
    <div className={styles.admin_approve_wrapper}>
      {me.is_admin === 1 ? (
        <div>
          <span>ADMIN MANAGE PERCENT</span>
          <div className={styles.users}>
            {users ? (
              users?.map((user) => (
                <div key={user.id} className={styles.users__card}>
                  <div>
                    <label>{`${user.firstname} ${user.lastname}`}</label>
                  </div>
                  <div>
                    <label>{user.email}</label>
                  </div>
                  {/* <div>
                <label>ID: {user.id}</label>
              </div> */}
                  <div>
                    <label>Percent: {user.user_percent}</label>
                  </div>
                  <div>
                    <input
                      required
                      type="text"
                      placeholder="set percent"
                      onChange={(e) => setPercent(e.target.value)}
                    />
                  </div>
                  <button
                    className={styles.button}
                    onClick={() => handleManage(user)}
                  >
                    manage
                  </button>
                </div>
              ))
            ) : (
              <div>No users to approve</div>
            )}
          </div>
        </div>
      ) : (
        <div>Not allowed</div>
      )}
    </div>
  )
}

export const AdminManageOrders = () => {
  const { me } = useContext(UserContext)
  const [orders, setOrders] = useState([])

  useEffect(() => {
    axios
      .get(`${local_backend_ip}/admin/orders`, { withCredentials: true })
      .then((res) => {
        /* console.log(res.data.orders) */
        setOrders(res.data.orders)
      })
  }, [])

  return (
    <div>
      {me.is_admin === 1 ? (
        <div className="orders_wrapper">
          {orders ? (
            <div className="orders_card">
              <h2>Active orders</h2>
              {orders?.map((order) => (
                <div className="orders_list" key={order.order_id}>
                  Order ID:{" "}
                  <Link to={`${order.order_id}`}>{order.order_id}</Link>
                </div>
              ))}
            </div>
          ) : (
            <div>You have no orders</div>
          )}
        </div>
      ) : (
        <div>Not allowed</div>
      )}
    </div>
  )
}

export const AdminSingleOrder = () => {
  const { id } = useParams()
  const [order, setOrder] = useState({})
  const { me } = useContext(UserContext)
  useEffect(() => {
    axios
      .get(`${local_backend_ip}/admin/orders/order?id=${id}`, {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res.data)
        setOrder(res.data)
      })
      .catch((err) => {
        console.log(err)
      })
  }, [id])
  return (
    <div className="order_wrapper">
      {order && (
        <div className="order_card">
          {me.is_admin === 1 ? (
            <div>
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

              <div className="user_information">
                <ul>
                  <li>{order.user?.email}</li>
                  <li>
                    {order.user?.firstname} {order.user?.lastname}
                  </li>
                </ul>
              </div>
              <div className="total_order_price">
                Total order price: {order.total_price?.toFixed(2)}&euro;
              </div>
            </div>
          ) : (
            <div>Not allowed</div>
          )}
        </div>
      )}
    </div>
  )
}

export default AdminPage
