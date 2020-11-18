package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("start server\n")

	file, err := os.Open("iplist.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	//  bufio.NewScanner() передаем прочитанный файл
	// bufio.ScanLines читает файл до спецсимволов "\n , \r"
	// scanner.Split разделяет строки по спецсимволам "\n , \r"
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var iplines []string

	// Scan() перемещаемся  по файлу по строкам
	//	добавляем  строки по строчно в слайс строк
	for scanner.Scan() {
		iplines = append(iplines, scanner.Text())
	}

	file.Close()

	for _, val := range iplines {
		fmt.Println(val)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
