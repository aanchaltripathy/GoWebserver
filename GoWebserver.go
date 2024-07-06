package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

// homeHandler handles requests to the root URL
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
}

// helloHandler handles requests to the /hello URL
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

// incrementCounter handles requests to increment the counter
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	// Serve static files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Register the other handlers
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/increment", incrementCounter)

	// Start the server
	log.Fatal(http.ListenAndServe(":8081", nil))
}
