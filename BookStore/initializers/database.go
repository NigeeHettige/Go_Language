package initializers

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("❌ Failed to open sql.DB:", err)
		return
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("❌ Failed to connect to GORM:", err)
		return
	}

	DB = gormDB
	fmt.Println("✅ Successfully connected to MySQL database!")
}
