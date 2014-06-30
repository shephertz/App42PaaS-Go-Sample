package main

import(
	"fmt"
	"net/http"
	"text/template"
	"log"
	"io"
  "io/ioutil"
  "os"
)

var (
    MyFile  *log.Logger
)

func Init(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

    file, _ := os.OpenFile("log/production.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    
    MyFile = log.New(file,
        "PREFIX: ",
        log.Ldate|log.Ltime|log.Lshortfile)    
}

func rootHandler(w http.ResponseWriter, r *http.Request){
  MyFile.Println("Inside Root Handler")
	fmt.Println("Inside Root Handler")
	t := template.New("index.html")
	t.ParseFiles("public/index.html")
	if err := t.Execute(w, t); err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func main() {
  Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	MyFile.Println("Inside Main Function")
	http.HandleFunc("/", rootHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		fmt.Println("Not Connecting to Host:", err)
	}
}
