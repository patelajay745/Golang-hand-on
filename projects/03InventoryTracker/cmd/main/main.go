package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/patelajay745/projects/03InventoryTracker/pkg/middleware"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3002"
	}

	r := routes.SetUpRoutes()

	r.Use(middleware.LoggerMiddleware)

	http.Handle("/", r)
	fmt.Println("Starting server at port 3002 ")
	log.Fatal(http.ListenAndServe(":3002", r))

}
