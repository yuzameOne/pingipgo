package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	fmt.Println("start server")

	ip, err := ioutil.ReadFile("iplist.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nFile contents: \n%s", ip)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
