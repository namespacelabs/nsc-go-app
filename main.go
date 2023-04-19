package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello KubeCon '23 from Namespace!\n")
}

func main() {
	fmt.Println("Hello KubeCon '23 from Namespace!")

	log.Println("Hello web server started")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	log.Printf("Web servers exists, bye bye! %v\n", err)
}
