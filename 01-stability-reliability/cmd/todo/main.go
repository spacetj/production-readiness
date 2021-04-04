package main

import (
	"fmt"
	"net/http"

	"01-stability-reliability/pkg/handlers"

	"github.com/go-pg/pg"

	"os"

	"github.com/labstack/echo"
)

// For displaying the generated SQL in the command-line.
type dbLogger struct{}

func main() {
	db := initDB()
	defer db.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	port := ":8000"
	fmt.Print("App running at http://localhost", port)
	e.Logger.Fatal(e.Start(port))
}

func initDB() *pg.DB {
	username, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		username = "admin"
	}

	pass, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		pass = ""
	}

	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		dbname = "postgresdb"
	}

	options := &pg.Options{
		Addr:     "postgres:5432",
		User:     username,
		Password: pass,
		Database: dbname,
	}
	db := pg.Connect(options)
	db.AddQueryHook(dbLogger{})
	err := migrate(db)
	if err != nil {
		panic(err)
	}
	return db
}

// Create the table
func migrate(db *pg.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Tasks(id BIGSERIAL PRIMARY KEY, name text)`)

	if err != nil {
		return err
	}
	return nil
}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

// Show the generated SQL in the command-line.
func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}
