package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	r = routes.SetUpRoutes()

	http.Handle("/", r)
	fmt.Println("Starting server at port 3001 ")
	log.Fatal(http.ListenAndServe(":3001", r))

}
