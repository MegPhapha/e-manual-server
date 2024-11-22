package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received: " + r.URL.Path)
	t, err := template.ParseFiles("index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 3001...")

	http.ListenAndServe(":3001", nil)
}