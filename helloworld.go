package main

import (
	"net/http"
	"log"
)

func main(){
	log.Println("start listening...")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		log.Println("hello there")
		w.Write([]byte("Hello there"))
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}