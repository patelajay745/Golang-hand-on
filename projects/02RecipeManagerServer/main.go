package main

import (
	"log"
	"net/http"

	"github.com/patelajay745/Golang-hand-on/projects/RecipeManagerServer/router"
)

func main() {

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))

}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
