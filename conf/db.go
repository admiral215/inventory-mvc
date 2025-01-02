package conf

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"inventory-bee/models"
	"log"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB == nil {
		ConnectDB()
	}
	return DB
}

func ConnectDB() {
	dbHost, _ := web.AppConfig.String("db_host")
	dbPort, _ := web.AppConfig.String("db_port")
	dbUser, _ := web.AppConfig.String("db_user")
	dbPass, _ := web.AppConfig.String("db_password")
	dbName, _ := web.AppConfig.String("db_name")
	dbSsl, _ := web.AppConfig.String("db_ssl")
	dbTimeZone, _ := web.AppConfig.String("db_timezone")

	// buat DSN (Data Source Name) dan ambil config dari app.conf
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSsl, dbTimeZone,
	)

	// Connect ke database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// migrate entity ke db
	err = database.AutoMigrate(
		// tambah entity disini
		&models.Category{},
		&models.Inventory{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	DB = database
}
