package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/patelajay745/projects/03InventoryTracker/pkg/routes"
)

func main() {

	r := routes.SetUpRoutes()

	http.Handle("/", r)
	fmt.Println("Starting server at port 3002 ")
	log.Fatal(http.ListenAndServe(":3002", r))

}
