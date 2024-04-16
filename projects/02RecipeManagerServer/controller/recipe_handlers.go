package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/patelajay745/Golang-hand-on/projects/RecipeManagerServer/model"
)

var globalRecipes []*model.Recipe

func CreteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/Json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}
	var recipe model.Recipe

	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipe.ID = rand.Intn(10000000)

	err = InsertOneRecipe(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(globalRecipes)

}

func GetRecipesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Recipes")
	w.Header().Set("Content-type", "application/Json")
	json.NewEncoder(w).Encode(globalRecipes)
}
func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get  a Recipe")
	w.Header().Set("Content-type", "application/Json")

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		// ... handle error
		panic(err)
	}

	for _, item := range globalRecipes {
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
		}
	}

}

func DeleteRecipesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Recipes")
	w.Header().Set("Content-type", "application/Json")

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		// ... handle error
		panic(err)
	}

	for index, item := range globalRecipes {
		if item.ID == i {
			globalRecipes = append(globalRecipes[:index], globalRecipes[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Deleted")
}

func UpdateRecipesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Recipes")
	w.Header().Set("Content-type", "application/Json")

	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		// ... handle error
		panic(err)
	}

	for index, item := range globalRecipes {
		if item.ID == i {
			globalRecipes = append(globalRecipes[:index], globalRecipes[index+1:]...)

			var recipe model.Recipe

			_ = json.NewDecoder(r.Body).Decode(&recipe)
			recipe.ID = i
			globalRecipes = append(globalRecipes, &recipe)
			json.NewEncoder(w).Encode(globalRecipes)

			break
		}
	}

}

func AddRecipes() {
	recipe1 := &model.Recipe{
		ID:          1,
		Title:       "Spaghetti Carbonara",
		Description: "Classic Italian pasta dish with bacon, eggs, and cheese.",
		Ingredients: []model.Ingredient{
			{Name: "Spaghetti", Amount: "200g"},
			{Name: "Bacon", Amount: "100g"},
			{Name: "Eggs", Amount: "2"},
			{Name: "Parmesan Cheese", Amount: "50g"},
			{Name: "Black Pepper", Amount: "to taste"},
		},
		Instructions: []model.Instruction{
			{StepNumber: 1, StepDesc: "Cook spaghetti according to package instructions."},
			{StepNumber: 2, StepDesc: "In a pan, cook bacon until crispy. Remove excess fat."},
			{StepNumber: 3, StepDesc: "In a bowl, mix eggs and grated Parmesan cheese."},
			{StepNumber: 4, StepDesc: "Add cooked spaghetti to the pan with bacon. Pour egg mixture over spaghetti. Stir until eggs are cooked."},
			{StepNumber: 5, StepDesc: "Season with black pepper and serve immediately."},
		},
		PrepTime: 10,
		CookTime: 15,
		Servings: 4,
	}
	globalRecipes = append(globalRecipes, recipe1)

	recipe2 := &model.Recipe{
		ID:          2,
		Title:       "Chicken Tikka Masala",
		Description: "Popular Indian dish with grilled chicken in a creamy tomato sauce.",
		Ingredients: []model.Ingredient{
			{Name: "Chicken Breast", Amount: "500g"},
			{Name: "Yogurt", Amount: "150g"},
			{Name: "Tomatoes", Amount: "400g"},
			{Name: "Onion", Amount: "1"},
			{Name: "Garlic", Amount: "3 cloves"},
			{Name: "Ginger", Amount: "1 inch piece"},
			{Name: "Garam Masala", Amount: "2 tsp"},
			{Name: "Cumin Powder", Amount: "1 tsp"},
			{Name: "Coriander Powder", Amount: "1 tsp"},
			{Name: "Cream", Amount: "100ml"},
		},
		Instructions: []model.Instruction{
			{StepNumber: 1, StepDesc: "Marinate chicken with yogurt, garlic, ginger, and spices for 1 hour."},
			{StepNumber: 2, StepDesc: "Grill marinated chicken until cooked through."},
			{StepNumber: 3, StepDesc: "In a pan, sauté chopped onions until soft. Add chopped tomatoes and cook until they break down."},
			{StepNumber: 4, StepDesc: "Blend onion-tomato mixture until smooth. Return to pan."},
			{StepNumber: 5, StepDesc: "Add grilled chicken pieces and cream to the pan. Simmer for 10 minutes."},
		},
		PrepTime: 60,
		CookTime: 30,
		Servings: 6,
	}
	globalRecipes = append(globalRecipes, recipe2)

	recipe3 := &model.Recipe{
		ID:          3,
		Title:       "Caprese Salad",
		Description: "Simple Italian salad with fresh tomatoes, mozzarella cheese, and basil.",
		Ingredients: []model.Ingredient{
			{Name: "Tomatoes", Amount: "4"},
			{Name: "Fresh Mozzarella", Amount: "200g"},
			{Name: "Fresh Basil Leaves", Amount: "handful"},
			{Name: "Extra Virgin Olive Oil", Amount: "2 tbsp"},
			{Name: "Balsamic Vinegar", Amount: "1 tbsp"},
			{Name: "Salt", Amount: "to taste"},
			{Name: "Black Pepper", Amount: "to taste"},
		},
		Instructions: []model.Instruction{
			{StepNumber: 1, StepDesc: "Slice tomatoes and mozzarella into equal-sized slices."},
			{StepNumber: 2, StepDesc: "Arrange tomato and mozzarella slices on a plate, alternating them."},
			{StepNumber: 3, StepDesc: "Tear basil leaves and scatter them over the tomatoes and mozzarella."},
			{StepNumber: 4, StepDesc: "Drizzle olive oil and balsamic vinegar over the salad."},
			{StepNumber: 5, StepDesc: "Season with salt and black pepper to taste."},
		},
		PrepTime: 10,
		CookTime: 0,
		Servings: 2,
	}
	globalRecipes = append(globalRecipes, recipe3)

	recipe4 := &model.Recipe{
		ID:          4,
		Title:       "Vegetable Stir-Fry",
		Description: "Quick and healthy stir-fried vegetables with soy sauce and garlic.",
		Ingredients: []model.Ingredient{
			{Name: "Broccoli", Amount: "200g"},
			{Name: "Carrots", Amount: "2"},
			{Name: "Bell Peppers", Amount: "2"},
			{Name: "Snow Peas", Amount: "100g"},
			{Name: "Garlic", Amount: "3 cloves"},
			{Name: "Soy Sauce", Amount: "2 tbsp"},
			{Name: "Vegetable Oil", Amount: "2 tbsp"},
			{Name: "Salt", Amount: "to taste"},
			{Name: "Black Pepper", Amount: "to taste"},
		},
		Instructions: []model.Instruction{
			{StepNumber: 1, StepDesc: "Cut broccoli into florets, carrots into matchsticks, and bell peppers into strips."},
			{StepNumber: 2, StepDesc: "Heat oil in a pan. Add minced garlic and cook until fragrant."},
			{StepNumber: 3, StepDesc: "Add broccoli, carrots, bell peppers, and snow peas to the pan. Stir-fry until vegetables are tender-crisp."},
			{StepNumber: 4, StepDesc: "Drizzle soy sauce over the vegetables. Season with salt and black pepper to taste."},
			{StepNumber: 5, StepDesc: "Serve hot as a side dish or over rice."},
		},
		PrepTime: 15,
		CookTime: 10,
		Servings: 4,
	}
	globalRecipes = append(globalRecipes, recipe4)

	recipe5 := &model.Recipe{
		ID:          5,
		Title:       "Banana Bread",
		Description: "Moist and delicious banana bread made with ripe bananas and walnuts.",
		Ingredients: []model.Ingredient{
			{Name: "Ripe Bananas", Amount: "3"},
			{Name: "All-Purpose Flour", Amount: "200g"},
			{Name: "Granulated Sugar", Amount: "100g"},
			{Name: "Butter", Amount: "100g"},
			{Name: "Eggs", Amount: "2"},
			{Name: "Baking Powder", Amount: "1 tsp"},
			{Name: "Walnuts", Amount: "50g"},
			{Name: "Vanilla Extract", Amount: "1 tsp"},
		},
		Instructions: []model.Instruction{
			{StepNumber: 1, StepDesc: "Preheat oven to 350°F (175°C). Grease a loaf pan."},
			{StepNumber: 2, StepDesc: "In a bowl, mash ripe bananas with a fork."},
			{StepNumber: 3, StepDesc: "Cream butter and sugar until light and fluffy. Beat in eggs and vanilla extract."},
			{StepNumber: 4, StepDesc: "Add mashed bananas to the butter mixture. Stir in flour and baking powder until combined."},
			{StepNumber: 5, StepDesc: "Fold in chopped walnuts. Pour batter into prepared loaf pan."},
			{StepNumber: 6, StepDesc: "Bake for 50-60 minutes or until a toothpick inserted into the center comes out clean."},
			{StepNumber: 7, StepDesc: "Let cool in the pan for 10 minutes before transferring to a wire rack to cool completely."},
		},
		PrepTime: 15,
		CookTime: 60,
		Servings: 8,
	}
	globalRecipes = append(globalRecipes, recipe5)
}

func InsertOneRecipe(recipe *model.Recipe) error {
	globalRecipes = append(globalRecipes, recipe)
	return nil

}
