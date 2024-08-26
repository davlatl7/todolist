package db

import "coinkeeper/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.Task{})
	if err != nil {
		return err
	}
	return nil
}
