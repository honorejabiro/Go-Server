package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "wrong input in the form", http.StatusBadRequest)
		return
	}

	firstName := r.FormValue("first-name")
	lastName := r.FormValue("last-name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	street := r.FormValue("street")
	city := r.FormValue("city")
	state := r.FormValue("state")
	zip := r.FormValue("zip")
	service := r.FormValue("service")
	experience := r.FormValue("experience")
	contactMethod := r.FormValue("contact-method")
	bestTime := r.FormValue("time")
	goals := r.FormValue("goals")
	topics := r.Form["topics"]

	if firstName == "" || lastName == "" || email == "" || phone == "" || service == "" {
		http.Error(w, "please fill out all required fields", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Thanks, %s %s!\n\n", firstName, lastName)
	fmt.Fprintf(w, "We received your request for: %s\n", service)
	fmt.Fprintf(w, "Email: %s\n", email)
	fmt.Fprintf(w, "Phone: %s\n", phone)
	fmt.Fprintf(w, "Address: %s, %s, %s %s\n", street, city, state, zip)
	fmt.Fprintf(w, "Experience: %s\n", experience)
	fmt.Fprintf(w, "Topics: %s\n", strings.Join(topics, ", "))
	fmt.Fprintf(w, "Preferred contact: %s during the %s\n", contactMethod, bestTime)
	fmt.Fprintf(w, "Goals: %s\n", goals)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at 8000\n")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
