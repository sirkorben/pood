import React from "react"
import { useState } from "react"
import styles from "./Admin.module.scss"
import { useEffect } from "react"
import { Link, useParams } from "react-router-dom"
import { useContext } from "react"
import { UserContext } from "../../components/utils/UserContext"
import "../../styles/orders-list.scss"
import myAxios from "../../components/utils/api/axios"
import styles2 from "../../styles/cart-orders-styles.module.scss"

const AdminPage = () => {
  const { me } = useContext(UserContext)
  if (me.is_admin !== 1) return <div className="no_task">Not allowed</div>

  return (
    <div className={styles.admin_wrapper}>
      <h2 className={styles.title}>Admin panel</h2>
      <h2 className={styles.section_title}>Active orders</h2>
      <AdminManageOrders />
      <h2 className={styles.section_title}>Manage Percent</h2>
      <AdminManagePercent />
      <h2 className={styles.section_title}>Approve users</h2>
      <AdminApprove />
    </div>
  )
}

const AdminApprove = () => {
  const { me } = useContext(UserContext)

  const [users, setUsers] = useState([])
  const [approve, setApprove] = useState(false)
  useEffect(() => {
    myAxios
      .get("/api/admin/approve")

      .then((res) => {
        //console.log(res.data)
        setUsers(res.data)
        setApprove(false)
      })
  }, [approve])

  const handleApprove = async (user) => {
    setApprove(true)
    myAxios.patch(
      "/api/admin/approve",
      JSON.stringify({
        id: user.id,
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
        activated: user.activated,
        date_created: user.date_created,
      })
    )
  }

  if (users === null) return <div className="no_task">No users to approve</div>

  return (
    <div className={styles.approve_wrapper}>
      <div className={styles.users_container}>
        {users?.map((user) => (
          <div className={styles.user_card} key={user.id}>
            <ul>
              <li>
                {user.firstname} {user.lastname}
              </li>
              <li>{user.email}</li>
            </ul>
            <button onClick={() => handleApprove(user)}>Approve</button>
          </div>
        ))}
      </div>
    </div>
  )
}

const AdminManagePercent = () => {
  const { me } = useContext(UserContext)

  const [users, setUsers] = useState([])
  const [manage, setManage] = useState(false)
  const [percent, setPercent] = useState(1.15)

  useEffect(() => {
    myAxios
      .get("/api/admin/managepercent")

      .then((res) => {
        // console.log(res.data)
        setUsers(res.data)
        setManage(false)
      })
  }, [manage])

  const handleManage = async (user) => {
    setManage(true)
    myAxios
      .patch(
        "/api/admin/managepercent",
        JSON.stringify({ id: user.id, user_percent: Number(percent) })
      )

      .catch((err) => {
        console.log(err)
      })
    Array.from(document.querySelectorAll("input")).forEach(
      (input) => (input.value = "")
    )
    setPercent(1.15)
  }
  // console.log("USERS ", users)

  if (users === null) return <div className="no_task">No users to manage</div>

  return (
    <div className={styles.manage_percent_wrapper}>
      {/*   <div className={styles.user_search_bar}>
        <input
          type="text"
          placeholder="Find user"
          onChange={(e) => setSearch(e.target.value)}
        />
        <button onClick={() => handleSearch(search)}>search</button>
      </div> */}
      <div className={styles.percents_container}>
        {users?.map((user) => (
          <div className={styles.percent_card} key={user.id}>
            <div className={styles.user_info}>
              <ul>
                <li>{`${user.firstname} ${user.lastname}`}</li>
                <li>{user.email}</li>
                <li>Percent: {user.user_percent}%</li>
              </ul>
            </div>
            <div className={styles.confirm_manage}>
              <input
                type="number"
                placeholder="set percent"
                onChange={(e) => setPercent(e.target.value)}
              />
              <button onClick={() => handleManage(user)}>confirm</button>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export const AdminManageOrders = () => {
  const { me } = useContext(UserContext)
  const [orders, setOrders] = useState([])

  useEffect(() => {
    myAxios
      .get("/api/admin/orders")

      .then((res) => {
        /* console.log(res.data.orders) */
        setOrders(res.data.orders)
      })
  }, [])
  if (orders === null) return <div className="no_task">No active orders</div>

  return (
    <div className="orders_wrapper">
      <div className="orders_wrapper">
        <div className="orders_card">
          {orders?.map((order) => (
            <div className="orders_list" key={order.order_id}>
              <span className="orders_list__order_id">Order ID</span>:{" "}
              <Link
                className="orders_list__link"
                to={`/admin/orders/${order.order_id}`}
              >
                {order.order_id}
              </Link>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export const AdminSingleOrder = () => {
  const { id } = useParams()
  const [order, setOrder] = useState({})
  const { me } = useContext(UserContext)
  useEffect(() => {
    myAxios
      .get(`/api/admin/orders/order?id=${id}`)

      .then((res) => {
        // console.log(res.data)
        setOrder(res.data)
      })
      .catch((err) => {
        console.log(err)
      })
  }, [id])

  if (me.is_admin !== 1) return <div className="no_task">Not allowed</div>

  return (
    <div className={styles2.cart_wrapper}>
      <div className={styles2.cart_wrapper__total_cart_info}>
        <div className={styles2.total_cart_info__info}>
          <ul className={styles2.admin}>
            <li>
              Order ID: <span className="mark">{order.order_id}</span>
            </li>
            <li>Total items: {order.positions?.length}</li>
            <li>
              Total price:{" "}
              <span className="mark">
                {order.total_price?.toFixed(2)}&euro;
              </span>
            </li>
            <li>
              <span className="mark">{order.user?.email}</span>
            </li>
            <li>
              <span className="mark">
                {order.user?.firstname} {order.user?.lastname}
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

  /*   return (
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
            <div className={styles.not_allowed}>Not allowed</div>
          )}
        </div>
      )}
    </div>
  ) */
}

export default AdminPage
