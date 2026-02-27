package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! (~ From the Server ~)")
}

func main() {
	//register the handler at the root path
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on 6767")
	//start the server on 6767 and log errors that occur
	err := http.ListenAndServe(":6767", nil)
	if err != nil {
		log.Fatal("Server failed to start:",err)
	}
}