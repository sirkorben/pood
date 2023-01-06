import React from "react"
import { useContext } from "react"
import { SearchContext } from "../utils/SearchContext"
import styles from "./Products.module.scss"

/* "price": 16.46,
"article": "352316171195",
"supplier": "Magneti Marelli",
"supplier_price_num": 180,
"brand": "magneti marelli",
"currency": "euro",
"currency_rate": "90",
"delivery": "21",
"weight": 0.7,
"name": "" */

const Products = () => {
  const { results } = useContext(SearchContext)
  console.log(results.prices)
  return (
    <div className={styles.products}>
      {results.prices?.map((result, index) => (
        <div key={index}>
          <h3>{result.article}</h3>
          <ul>
            <li>Article: {result.article}</li>
            <li>Price: {result.price}</li>
            <li>Supplier: {result.supplier}</li>
            <li>Supplier_price_num: {result.supplier_price_num}</li>
            <li>Brand: {result.brand}</li>
            <li>Currency: {result.currency}</li>
            <li>Currency rate: {result.currency_rate}</li>
            <li>Delivery: {result.delivery}</li>
            <li>Weight: {result.weight}</li>
            <li>Name: {result.name === "" ? result.name : "No name"}</li>
          </ul>
        </div>
      ))}
    </div>
  )
}

export default Products
