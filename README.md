# How to start project locally

Go server is set to run on port :8080

simply `go run .` from backend repository

React app is set to use `local_backend_ip` and will be available on http://localhost:3000

simply `yarn install and npm start` from web repository

## Running locally inside docker containers

To allow Go server accepting requests from our React(that will run from container on different port) make changes in `/helpers/cors.go`

To make React app(running from one container) call right backend IP (running from another container) change uri link to
`` const res = axios.post(`${process.env.REACT_APP_BACKEND_URL}/signup` // where it is needed ``

and in the root of project
simply `docker compose up`
go to http://localhost
