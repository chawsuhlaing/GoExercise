package main
import "fmt"
func main () {
	var username string
	var password string
	fmt.Print("Please Enter Your Username: ")
	fmt.Scan(&username)
	fmt.Print("Please Enter Your Password: ")
	fmt.Scan(&password)
	fmt.Println("Username",username, "Password",password)
	if username == "admin" && password == "password123," {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed.")
	}
}