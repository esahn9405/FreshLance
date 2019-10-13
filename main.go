package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/choose", http.StatusFound)
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/portfolio/0", http.StatusFound)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/resume.pdf")
}

func isaacPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/portfolio.html")
}

func choiceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/choice.html")
}

func main() {
	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/resume", resumeHandler)
	r.HandleFunc("/portfolio", portfolioHandler)
	r.HandleFunc("/portfolio/0", isaacPortfolioHandler)
	r.HandleFunc("/choose", choiceHandler)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets/"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("templates/images/"))))

	// Serve at port 8080
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
