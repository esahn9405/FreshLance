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

func hirePortfolioHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/hire.html")
}

func isaacHirePortfolioHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/isaac_hire_portfolio.html")
}

func workPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/work/portfolio/0", http.StatusFound)
}

func isaacWorkPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/work_portfolio.html")
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/resume.pdf")
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
	r.HandleFunc("/hire", hirePortfolioHandler)
	r.HandleFunc("/hire/portfolio/0", isaacHirePortfolioHandler)
	r.HandleFunc("/work", workPortfolioHandler)
	r.HandleFunc("/work/portfolio/0", isaacWorkPortfolioHandler)
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
