# PERFORMING CRUD with (GO-LANG/GORM/GIN-FRAMEWORK/POSTSGRES-DATABASE)
## Steps to run this project are shown below : 
### NOTE: Remember to setup Postgres database first in your system before running this code
Run these below commands in your terminal
- step:1) `go get -u github.com/gin-gonic/gin`
- step:2) `go get -u github.com/jinzhu/gorm`
- step:3) `go get github.com/jinzhu/gorm/dialects/postgres`
- step:4) `go get github.com/lib/pq`
- step:5) `go run main.go`

## Rest Api's that are being served here:
- *POST* -> `localhost:8080/user`
- *GET* -> `localhost:8080/user/{id}`
- *PUT* -> `localhost:8080/user/{id}`
- *GET* -> `localhost:8080/users`
- *DELETE* -> `localhost:8080/user/{id}`

- listening at port `:8080`

