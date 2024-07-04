package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JFMcFaull/gamingCompanion/controllers"
	"github.com/JFMcFaull/gamingCompanion/initalizers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"} // Adjust this to your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.POST("/games", controllers.GamesCreate)
	r.PUT("/games/:id", controllers.GamesUpdate)
	r.GET("/games", controllers.GamesIndex)
	r.GET("/games/:id", controllers.GamesShow)
	r.DELETE("/games/:id", controllers.GamesDelete)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (or any desired port)

	// Read environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	sslMode := os.Getenv("SSL_MODE")

	// Construct connection string
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbHost, dbPort, dbName, dbUser, dbPass, sslMode)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	// Test the database connection
	if err = db.Ping(); err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		return
	}
	fmt.Println("Successfully connected to the PostgreSQL database")

	// Example query
	rows, err := db.Query("SELECT NOW()")
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close()

	// Process query results
	var currentTime string
	for rows.Next() {
		if err := rows.Scan(&currentTime); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		fmt.Println("Current time from database:", currentTime)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating over results: %v\n", err)
		return
	}

}
