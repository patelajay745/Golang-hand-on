package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/ajay", ajayHandler)

	http.HandleFunc("/contact", contactHandler)

	// http.HandleFunc()

	fmt.Println("Starting server at port 8000 \n")

	http.ListenAndServe(":8000", nil)

}

func ajayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ajay")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	message := r.FormValue("message")

	fmt.Printf("First Name: %s\n", fname)
	fmt.Printf("Last Name: %s\n", lname)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Message: %s\n", message)

}
