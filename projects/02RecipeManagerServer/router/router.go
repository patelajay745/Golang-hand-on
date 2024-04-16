package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/patelajay745/Golang-hand-on/projects/RecipeManagerServer/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	controller.AddRecipes()

	r.HandleFunc("/recipes", controller.CreteRecipeHandler).Methods("POST")
	r.HandleFunc("/recipes", controller.GetRecipesHandler).Methods("GET")
	r.HandleFunc("/recipes/{id}", controller.DeleteRecipesHandler).Methods("Delete")
	r.HandleFunc("/recipes/{id}", controller.UpdateRecipesHandler).Methods("PUT")
	r.HandleFunc("/recipe/{id}", controller.GetRecipeHandler).Methods("GET")

	fmt.Println("Starting server ar port 8000")

	return r
}
