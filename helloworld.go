package main

import (
	"net/http"
	"log"
)

func main(){
	log.Println("start listening on port 8080...")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		log.Println("hello there")
		w.Write([]byte("Hello there"))
	})

	http.Handle("/files", http.StripPrefix("/files", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
