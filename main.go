package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w is the response writer
	// w will be used to write the response back to the client
	// r is the request
	// r will be used to get the request information

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// parseForm is needed to parse the form
	// this will parse the form and update r.PostForm and r.Form

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// this will print out the form information
	fmt.Fprintf(w,"POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	// this simple web server can serve static files
	// there are 3 routes: /, /form, /hello
	// / is the root route, it serves static files
	// /form is the route to the form page
	// /hello is the route to the hello page
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	// helloHandler will just print out the hello message
	http.HandleFunc("/hello", helloHandler)

	// this is the server address
	fmt.Println("Server is listening on port 8080")
	// create the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
