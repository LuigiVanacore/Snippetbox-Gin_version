package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		log.Print("error home {}", err)
	}
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a specific snippet..."))
	if err != nil {
		log.Print("error snippetView {}", err)
	}
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		log.Print("error snippetCreate {}", err)
	}
}

func main() {
	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
