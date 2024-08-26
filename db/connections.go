package db

import (
	"coinkeeper/configs"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {

	connStr := fmt.Sprintf(`host=%s 
									port=%s 
									user=%s 
									dbname=%s 
									password=%s`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")

	dbConn = db
	return nil
}

func CloseDBConn() error {

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
