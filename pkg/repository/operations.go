package repository

import (
	"coinkeeper/db"
	"coinkeeper/models"
)

func AddTask(o models.Task) error {
	return nil
}

func GetAllTasks() (operations []models.Task, err error) {
	return operations, nil

}

func UpdateTask(o models.Task) error {

	err := db.GetDBConn().Save(&o).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteTAsk(id int) error {

	db.GetDBConn().Table("tasks").Where("id = ?", id).Update("is_deleted", true)

	return nil
}
