package main

import (
	n "c361CourseProject_UI/Navigation"
)

func main() {

	n.AppTitleIntro()
	n.PickAnIngredient()
	n.ChooseYourRecipe("CHICKEN")
	n.NAVChooseYourRecipe()
	n.Recipe("Chicken Fettucini Alfredo")
	n.NAVRecipe()
}
