package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	dbPort := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database", err)
	}
	fmt.Println("connect database success")
}
func helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello word"})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", helloworld)

	return r
}

func main() {
	initDB()
	r := setupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
