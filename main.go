package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("start server")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
