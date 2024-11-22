package main

import ( 
	"fmt"
	"net/http"
	"html/template"
	"time"

)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved: " + r.URL.Path)
	//fmt.Fprint(w, "<h1>Welcome To Night Owl</h1>")
	t, err := template.ParseFiles("templates/index.html")

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	pages, _ := scandir ("./manuals")
	fmt.Println(pages)
	t.Execute(w, pages)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved: " + r.URL.Path)
	//fmt.Fprint(w, "<h1>Welcome To Night Owl</h1>")
	t, err := template.ParseFiles("templates/news.html")

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	//get today's date
date := time.Now().String()
	t.Execute(w, date)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/news", newsHandler)
	http.Handle("/images/", http.StripPrefix ("/images/" ,(http.FileServer (http.Dir("images")))))
	http.Handle("/css/", http.StripPrefix ("/css/" ,(http.FileServer (http.Dir("css")))))
	http.Handle("/manuals/", http.StripPrefix ("/manuals/" ,(http.FileServer (http.Dir("manuals")))))
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)

}