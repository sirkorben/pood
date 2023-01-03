# How to start project locally

Go server is set to run on port :8080

simply `go run .`

React app is set to export `local_backend_ip` and will be available on http://localhost:3000

simply `yarn install and npm start`

## Running locally inside docker containers

To allow Go server accepting requests from our origin make changes in `handlers.go`

To make React app(running from one container) call right backend IP (running from another container) change uri link to
`` const res = axios.post(`${process.env.REACT_APP_BACKEND_URL}/signup` // where it is needed ``

and in the root of project
simply `docker compose up`
