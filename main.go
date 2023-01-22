package main

import (
	"github.com/kkyr/go-recipe/pkg/recipe"
)

func main() {

	// n.AppTitleIntro()
	// n.PickAnIngredient()
	// n.ChooseYourRecipe("CHICKEN")
	// n.NAVChooseYourRecipe()
	// n.Recipe("Chicken Fettucini Alfredo")
	// n.NAVRecipe()

	url := "https://minimalistbaker.com/quick-pickled-jalapenos/"

	recipe, err := recipe.ScrapeURL(url)
	if err != nil {
		// handle err
	}

	ingredients, ok := recipe.Ingredients()
	if !ok {
		println("Not ok")
	}
	instructions, ok := recipe.Instructions()
	if !ok {
		println("Not ok")
	}

	for i := range ingredients {
		println(i)
	}

	for i := range instructions {
		println(i)
	}
	println(recipe.Author())
	println(recipe.CookTime())
	println(recipe.Cuisine())
	println(recipe.Description())
	println(recipe.Yields())
}
