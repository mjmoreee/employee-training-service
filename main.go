package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type EmployeeData struct {
	EmployeeID int
	FirstName  string
	LastName   string
	Role       string
	Department string
	HiredDate  string
	Email      string
	Telephone  string
	ImagePath  string
}

type RankingData struct {
	Ranking    int
	FirstName  string
	LastName   string
	Role       string
	Department string
	Points     int
}

type EmployeeDetail struct {
	EmployeeID int
	FirstName  string
	LastName   string
	// Informasi lainnya
}

// Execute query from database
func executeQuery(db *sql.DB, query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func main() {
	// Router Setup
	router := gin.Default()

	// Serve static files from the '/assets' directory
	router.Static("/assets", "./assets")

	// Database Connection
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/employee_training_service")
	if err != nil {
		// Handle database connection error
		panic(err.Error())
	}
	defer db.Close()

	// Example of using the 'executeQuery' function
	query1 := "SELECT employee_id, first_name, last_name, role, department, hired_date, email, telephone, image_path FROM employee"
	query2 := "SELECT ranking, first_name, last_name, role, department, points FROM employee ORDER BY ranking ASC, points DESC;"

	// Execute the first query and handle errors
	employeeProfileQuery, err := executeQuery(db, query1)
	if err != nil {
		log.Fatal(err)
	}
	defer employeeProfileQuery.Close()

	// Process the results of the first query here

	// Execute the second query and handle errors
	rankingQuery, err := executeQuery(db, query2)
	if err != nil {
		log.Fatal(err)
	}
	defer rankingQuery.Close()

	// Initialize a slice to store employee data
	var employeeData []EmployeeData

	// Iterate through the rows of the first query and process data
	for employeeProfileQuery.Next() {
		var employee EmployeeData
		err := employeeProfileQuery.Scan(&employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.Role, &employee.Department, &employee.HiredDate, &employee.Email, &employee.Telephone, &employee.ImagePath)
		if err != nil {
			// Handle scanning error
			panic(err.Error())
		}
		employeeData = append(employeeData, employee)
	}

	// Initialize a slice to store employee data
	var rankingData []RankingData

	// Iterate through the rows of the first query and process data
	for rankingQuery.Next() {
		var ranking RankingData
		err := rankingQuery.Scan(&ranking.Ranking, &ranking.FirstName, &ranking.LastName, &ranking.Role, &ranking.Department, &ranking.Points)
		if err != nil {
			// Handle scanning error
			panic(err.Error())
		}
		rankingData = append(rankingData, ranking)
	}

	///////////////////////////////////////////////
	// ROUTER SETUP
	//
	// In this section, we configure the web router
	// and define routes for different pages.
	///////////////////////////////////////////////

	// Load HTML templates from the 'templates' directory
	router.LoadHTMLGlob("templates/*")

	// Define the root route for the dashboard page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Dashboard",
		})
	})

	// Define the 'leaderboard' route for the leaderboard page
	router.GET("/leaderboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "leaderboard.html", gin.H{
			"title":   "Leaderboard",
			"ranking": rankingData,
		})
	})

	// Define the 'leaderboard' route for the leaderboard page
	router.GET("/employees", func(c *gin.Context) {
		c.HTML(http.StatusOK, "employees.html", gin.H{
			"title":    "Employees",
			"employee": employeeData,
		})
	})

	// Define the 'leaderboard' route for the leaderboard page
	router.GET("/employees/profile/:id", func(c *gin.Context) {
		// Get ID
		id := c.Param("id")

		var employeeDetailData EmployeeDetail
		err = db.QueryRow("SELECT employee_id, first_name, last_name FROM employee WHERE employee_id = ?", id).Scan(&employeeDetailData.EmployeeID, &employeeDetailData.FirstName, &employeeDetailData.LastName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "employees-details.html", gin.H{
			"title":          "Employees' Detail",
			"employeeDetail": employeeDetailData,
		})
	})

	// Start the router and listen on port 8000
	router.Run(":8000")
}
