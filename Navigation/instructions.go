package Navigation

import "fmt"

func AppTitleIntro() {
	println()
	println()
	println("*****  *****  *****  *****  *****  *****    *****  *****  *    *")
	println("*   *  *      *        *    *   *  *        *   *  *   *  *    *")
	println("*   *  *      *        *    *   *  *        *  *   *   *   *  * ")
	println("*****  ****   *        *    *****  ****     ***    *   *    **  ")
	println("**     *      *        *    *      *        *  *   *   *    **  ")
	println("* *    *      *        *    *      *        *   *  *   *   *  * ")
	println("*  *   *      *        *    *      *        *   *  *   *   *  * ")
	println("*   *  *****  *****  *****  *      *****    *****  *****  *    *")
	println()
	println()
	println("                 WELCOME TO YOUR RECIPE BOX!")
	println()
	println("                    * Pick an ingredient")
	println("                    * Choose a recipe")
	println("                    * Make your meal!")
}

func PickAnIngredient() {
	println()
	println("***** CHOOSE YOUR MAIN INGREDIENT *****")
	println()
	println("1. Chicken")
	println("2. Beef")
	println("3. Pork")
	println("4. Seafood")
	println("5. Vegetable")
	println()
	println("6. EXIT RECIPE BOX")
	println("(Type number and press enter to make your selection.)")
}

func ChooseYourRecipe(ingredient string) {
	println()
	fmt.Printf("***** %s RECIPES *****", ingredient)
	println()
}

func NAVChooseYourRecipe() {
	println()
	println("1. PICK DIFFERENT MAIN INGREDIENT")
	println("2. EXIT RECIPE BOX")
	println("(Type number and press enter to make your selection.)")
}

func Recipe(recipeName string) {
	println()
	fmt.Printf("***** RECIPE: %s *****", recipeName)
	println()
}

func NAVRecipe() {
	println()
	println("1. PRINT RECIPE")
	println("2. EMAIL RECIPE")
	println("3. DOWNLOAD RECIPE")
	println("4. CONVERT RECIPE TO METRIC")
	println("5. CHOOSE A DIFFERENT RECIPE")
	println("6. PICK DIFFERENT MAIN INGREDIENT")
	println("7. EXIT RECIPE BOX")
	println("(Type number and press enter to make your selection.)")
}
