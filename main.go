// PERFORMING CRUD with (GO-LANG/GORM/GIN-FRAMEWORK/POSTSGRES-DATABASE)

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Creating a variable `db` of type GORM.DB that store our database connection and will be used every time to interact
// with database to perform CRUD operations.
var db *gorm.DB
var err error

// User is a model.
type User struct {
	//Parsing JSON into the struct
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	// Establishing the connection with database (Postgres) [ You can provide your own db connection credentials here ]
	db, err = gorm.Open("postgres", "user=postgres password=postgres dbname=test_db")
	if err != nil {
		fmt.Println(err)
	}
	// Close the resource of db with the keyword `defer` .
	defer db.Close()

	// create table `user` in database
	db.AutoMigrate(&User{})

	// Gin framework routes
	route := gin.Default()

	// route path eg:
	route.POST("/user", CreateUser)
	route.GET("/users/", GetAllUsers)
	route.GET("/user/:id", GetUserByID)   // path variable `id`
	route.PUT("/user/:id", UpdateUser)    // path variable `id`
	route.DELETE("/user/:id", DeleteUser) // path variable `id`

	// listening at :8080
	route.Run(":8080")
}

//CreateUser func (create a new record in the database).
func CreateUser(c *gin.Context) {
	// (c *gin.Context) mentioned in above function is the pointer of Gin Framework.
	var user User
	c.BindJSON(&user)
	if err := db.Create(&user); err.Error != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, user)
	}
}

//GetAllUsers func (get all the records from the database).
func GetAllUsers(c *gin.Context) {
	var user []User
	if err := db.Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

//GetUserByID func (get one specific record from the database against the id provided).
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

//UpdateUser func (updates the record in the database against the id provided).
func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
	}
	//binding the data in JSON and saving in DB.
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}

//DeleteUser func (delete the record in the database against the id provided).
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		db.Delete(&user)
		c.JSON(200, "User is deleted")
	}
}
