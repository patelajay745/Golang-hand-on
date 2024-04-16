package model

// Recipe represents a recipe entity
type Recipe struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
	PrepTime     int           `json:"prepTime"` // in minutes
	CookTime     int           `json:"cookTime"` // in minutes
	Servings     int           `json:"servings"`
}

// Ingredient represents an ingredient in a recipe
type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

// Instruction represents a step in a recipe
type Instruction struct {
	StepNumber int    `json:"stepNumber"`
	StepDesc   string `json:"stepDesc"`
}
