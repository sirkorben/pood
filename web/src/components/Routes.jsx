import React from "react"
import { Routes, Route } from "react-router-dom"
import Home from "../pages/home/Home"
import AdminPage from "../pages/admin/Admin"
import { AdminApprove, AdminManagePercent } from "../pages/admin/Admin"
import Me from "../pages/me/Me"
import SearchPage from "../pages/search/Search"
import Cart from "../pages/cart/Cart"
import MyOrders from "../pages/myorders/MyOrders"
import { MyOrder } from "../pages/myorders/MyOrders"
const AllRoutes = () => {
  return (
    <div>
      <div>
        <Routes>
          <Route path="/admin" element={<AdminPage />} />
          <Route path="/admin/approve" element={<AdminApprove />} />
          <Route path="/admin/managepercent" element={<AdminManagePercent />} />
          <Route path="/me" element={<Me />} />
          <Route path="/search" element={<SearchPage />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/myorders" element={<MyOrders />} />
          <Route path="/myorders/:id" element={<MyOrder />} />
          <Route path="/home" element={<Home />} />
        </Routes>
      </div>
    </div>
  )
}

export default AllRoutes
