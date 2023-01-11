import React from "react"
import ReactDOM from "react-dom/client"
import "./styles/styles.scss"
import App from "./App"

import { BrowserRouter } from "react-router-dom"
import { UserContextProvider } from "./components/utils/UserContext"
import SearchContextProvider from "./components/utils/SearchContext"
import { CartContextProvider } from "./components/utils/CartContext"
const root = ReactDOM.createRoot(document.getElementById("root"))

root.render(
  <React.StrictMode>
    <UserContextProvider>
      <SearchContextProvider>
        <CartContextProvider>
          <BrowserRouter>
            <App />
          </BrowserRouter>
        </CartContextProvider>
      </SearchContextProvider>
    </UserContextProvider>
  </React.StrictMode>
)
