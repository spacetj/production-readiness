package handlers

import (
	"net/http"
	"strconv"

	"01-stability-reliability/pkg/models"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask endpoint
func PutTask(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// Map incoming JSON body to the new Task
		if err := c.Bind(&task); err != nil {
			return err
		}

		// Add a task using our new model
		id, err := models.PutTask(db, task.Name)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		}
		// Handle any errors
		return err
	}
}

// DeleteTask endpoint
func DeleteTask(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteTask(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		}
		// Handle errors
		return err
	}
}
