package connection

import (
	"fmt"
	"restaurant-management/pkg/config"
	"restaurant-management/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	dbConfig := config.LocalConfig
	//d, err := gorm.Open("mysql", "root:Salman12#@/restaurantmanagemensystem?charset=utf8&parseTime=True&loc=Local")
	dsn := fmt.
		Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DBUser, dbConfig.DBPass, dbConfig.DbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	db = d
}

func migrate() {
	// db.Migrator().DropTable(&models.Food{})
	// db.Migrator().DropTable(&models.Menu{})
	// db.Migrator().DropTable(&models.OrderItem{})
	db.Migrator().AutoMigrate(
		&models.Food{},
		&models.User{},
		&models.Menu{},
		&models.OrderItem{},
	)
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
