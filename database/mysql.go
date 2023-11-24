// database/mysql.go
package database

import (
	"fmt"
	"go-fiber/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectMySQL initializes the MySQL database connection
func ConnectMySQL() {
	// Initialize the database connection
	var err error
	DB, err = gorm.Open(mysql.Open(config.GetMySQLConnection()), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error Connect To MySQL Database: %s\n", err)
		panic(err)
	}

}
