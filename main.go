package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}
       
        fmt.Fprintf(w, "<h1>Hello from:  "+title+" </h1>")
        fmt.Fprintf(w, "<img src='assets/jxlogo.png' alt=‘jxlogo’ style='width:500;height:180;'>")        
}

func main() {
	http.HandleFunc("/", handler)
        http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
