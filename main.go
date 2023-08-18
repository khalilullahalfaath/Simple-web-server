package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	// this simple web server can serve static files
	// there are 3 routes: /, /form, /hello
	// / is the root route, it serves static files
	// /form is the route to the form page
	// /hello is the route to the hello page
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
}
