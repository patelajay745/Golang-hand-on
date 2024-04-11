package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {

	fmt.Println("Server is started")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":3002", r))

}
