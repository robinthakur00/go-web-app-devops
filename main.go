package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)

    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/about.html"))
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Handle form submission here
        name := r.FormValue("name")
        email := r.FormValue("email")
        message := r.FormValue("message")
        fmt.Printf("Received contact form submission: Name=%s, Email=%s, Message=%s\n", name, email, message)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    tmpl := template.Must(template.ParseFiles("templates/contact.html"))
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
