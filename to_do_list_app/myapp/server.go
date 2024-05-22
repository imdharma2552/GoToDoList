package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

var db *sql.DB

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Status      string `json:"status"`
}

type Login struct {
	Username string `json: "UserName"`
	Email    string `json: "email"`
	Password string `json : "password"`
}

func connectDB() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "todolist",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// Now you can execute SQL queries using the 'db' object
	return db
}

func main() {

	e := echo.New()
	db := connectDB()

	e.Use(echo.WrapMiddleware(cors.Default().Handler))
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/send-data", func(c echo.Context) error {
		var data map[string]string
		fmt.Println("Hello")

		if err := c.Bind(&data); err != nil {
			return err
		}
		rows, err := db.Query("SELECT username, password FROM user WHERE username = ? AND password = ?", data["username"], data["password"])
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		// Process the query results
		rowCount :=0

		for rows.Next() {
			var username string
			var password string
			if err := rows.Scan(&username, &password); err != nil {
				panic(err.Error())
			}
			rowCount++
			fmt.Printf("%s  %s\n", username, password)
			print("Login Successful :)")

		}

		if(rowCount == 0){
			return c.String(http.StatusNotFound,"user not found")
		}

		return c.String(http.StatusOK, "Data Sent Successfully")
	})

	e.GET("/tasks", func(c echo.Context) error {
		db := connectDB()
		rows, err := db.Query("SELECT id, title, description, start_date, end_date, start_time, end_time, status FROM tasks WHERE status='incomplete' AND end_date >= CURRENT_DATE()")

		fmt.Println(rows)
		if err != nil {
			return err
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var task Task
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.StartDate, &task.EndDate, &task.StartTime, &task.EndTime, &task.Status)
			if err != nil {
				return err
			}
			fmt.Println(task)
			tasks = append(tasks, task)
		}
		if err := rows.Err(); err != nil {
			return err
		}
		fmt.Println(tasks)
		return c.JSONPretty(http.StatusOK, tasks, "")
	})

	e.POST("/tasks/complete", func(c echo.Context) error {
		var data int

		if err := c.Bind(&data); err != nil {
			return err
		}
		fmt.Println(data)
		_, _ = db.Exec("UPDATE tasks SET status = 'completed' WHERE id = ?", data)
		return c.NoContent(http.StatusOK)
	})

	e.POST("/tasks/delete", func(c echo.Context) error {
		var data int

		if err := c.Bind(&data); err != nil {
			return err
		}
		fmt.Println(data)
		_, _ = db.Exec("UPDATE tasks SET status = 'deleted' WHERE id = ?", data)
		return c.NoContent(http.StatusOK)
	})

	e.GET("/tasks/history", func(c echo.Context) error {
		db := connectDB()

		rows, err := db.Query("SELECT title, description, status FROM tasks WHERE status IN ('deleted', 'incomplete','completed')")
		if err != nil {
			return err
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.Title, &task.Description, &task.Status); err != nil {
				return err
			}
			tasks = append(tasks, task)
		}

		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/Create", func(c echo.Context) error {
		db := connectDB()
		var task Task
		if err := c.Bind(&task); err != nil {
			return err
		}
		fmt.Println(task)

		_, _ = db.Exec("INSERT INTO tasks (title, description, start_date, end_date, start_time, end_time, status) VALUES (?, ?, ?, ?, ?, ?, ?)",
			task.Title, task.Description, task.StartDate, task.EndDate, task.StartTime, task.EndTime, task.Status)

		return c.String(http.StatusOK, "Data inserted successfully")
	})

	e.POST("/Createaccount", func(c echo.Context) error {
		db := connectDB()
		var data Login
		if err := c.Bind(&data); err != nil {
			return err
		}
		fmt.Println(data)

		_, _ = db.Exec("INSERT INTO user (username,email,password) VALUES (?, ?, ?)",
			data.Username, data.Email, data.Password)

		return c.String(http.StatusOK, "Data inserted successfully")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
