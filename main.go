package main

import(
	"fmt"
	"net/http"
	"text/template"
	"log"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Inside Root Handler")
	t := template.New("index.html")
	t.ParseFiles("public/index.html")
	if err := t.Execute(w, t); err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		fmt.Println("Not Connecting to Host:", err)
	}
}
