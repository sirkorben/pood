import axios from "axios"
//import { backend_url } from "../../../App"

const backend_url = "http://localhost:8080"

const myAxios = axios.create({
  /* baseURL: backend_url, */
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
})

export default myAxios
