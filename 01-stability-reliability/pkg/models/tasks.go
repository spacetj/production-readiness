package models

import (
	"fmt"

	"github.com/go-pg/pg"
)

// Task blob schema
type Task struct {
	ID   int    `json:"id"`
	Name string `sql:",notnull" json:"name"`
}

// GetTasks from the DB
func GetTasks(db *pg.DB) []Task {
	tasks := []Task{}
	err := db.Model(&tasks).Select()
	if err != nil {
		fmt.Println(err)
	}
	return tasks
}

// PutTask into DB
func PutTask(db *pg.DB, name string) (int, error) {
	Name := &Task{
		Name: name,
	}
	err := db.Insert(Name)
	if err != nil {
		fmt.Println(err)
	}
	return Name.ID, nil
}

// DeleteTask from DB
func DeleteTask(db *pg.DB, id int) (int, error) {

	idToDelete := &Task{
		ID: id,
	}

	err := db.Delete(idToDelete)
	if err != nil {
		fmt.Println(err)
	}
	return id, nil
}
