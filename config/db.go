package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * Created by Goland & VS Code.
 * User : 1. Roby Purba Sakty 			: obykao26@gmail.com
		  2. Muhammad Irfan 			: mhd.irfann00@gmail.com
   		  3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
 * Date: Saturday, 12 Juni 2023
 * Time: 08.30 AM
 * Description: BRI-CMP-Service-Backend
 **/

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	dbPort := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbname)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
