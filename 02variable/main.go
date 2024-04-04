package main

import "fmt"

const LoginToken string = "agasgsgasg"

func main() {
	var username string = "ajay"
	fmt.Printf("variable is type: %T \n", username)

	var isLoginIn bool = false
	fmt.Printf("variable is type: %T \n", isLoginIn)

	var smallInt int16 = 255
	fmt.Printf("variable is type: %T \n", smallInt)

	var smallFloat float32 = 45.05515154
	fmt.Printf("variable is type: %T \n", smallFloat)
	fmt.Println(smallFloat)

	fmt.Println(LoginToken)
	fmt.Printf("variable is type: %T \n", LoginToken)
}
