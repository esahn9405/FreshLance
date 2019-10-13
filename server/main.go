package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"html/template"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles()
	fmt.Fprintf(w, "%s", "hello")
}

func main() {
	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	// Serve at port 8080
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
