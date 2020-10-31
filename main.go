package main

import (
	database "demo-postgres/config"
	router "demo-postgres/router"
)

// CreateServer creates a new Fiber instance
func main() {
	database.ConnectToDB()
	router.SetupRoutes()
}
