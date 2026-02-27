package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:6767/")
	if err != nil {
		log.Fatal("Error connecting to server:",err)
	}
	//close the response body when the function exits
	defer response.Body.Close()
	//read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response:",err)
	}
	fmt.Printf("Server Response: %s",body)
}
