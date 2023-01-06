import React from "react"
import ReactDOM from "react-dom/client"
import "./styles/styles.scss"
import App from "./App"

import { BrowserRouter } from "react-router-dom"
import { UserContextProvider } from "./components/utils/UserContext"
import SearchContextProvider from "./components/utils/SearchContext"
const root = ReactDOM.createRoot(document.getElementById("root"))

root.render(
  <React.StrictMode>
    <UserContextProvider>
      <SearchContextProvider>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </SearchContextProvider>
    </UserContextProvider>
  </React.StrictMode>
)
