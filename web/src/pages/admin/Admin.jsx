import axios from "axios"
import React from "react"
import { useState } from "react"
import styles from "./Admin.module.scss"
import { local_backend_ip } from "../../App"
import { useEffect } from "react"
import { Link } from "react-router-dom"

const AdminPage = () => {
  return (
    <div className={styles.admin}>
      <Link to={"/admin/approve"}>ADMIN APPROVE</Link>
      <Link to={"/admin/managepercent"}>ADMIN MANAGE PERCENT</Link>
    </div>
  )
}

export const AdminApprove = () => {
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
  )
}

export const AdminManagePercent = () => {
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
      });
      Array.from(document.querySelectorAll("input")).forEach(
        input => (input.value = "")
      );
      setPercent(1.15)
  }

  return (
    <div className={styles.admin_approve_wrapper}>
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
  )
}

export default AdminPage
