package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func ping(ip ...string) {

	// c - sending count ECHO_REQUEST packets
	// i - Wait interval seconds between sending each packet
	// Only super-user may set interval to values less than 0.2 seconds
	// w - deadline Specify a timeout, in seconds, before ping exits regardless of how
	// many packets have been sent or received.
	fmt.Printf("ip array : %s", ip[3]) //dedug out
	
	out, _ := exec.Command("ping", ip[3], "-c 2", "-i 1", "-w 2").Output()

	fmt.Println(string(string(out))) //dedug out

	if strings.Contains(string(out), "2 received") {
		fmt.Println("Yeap, I'am ALIVEEE")
	} else {
		fmt.Println("Dead Mother Fucker")
	}

	for _, val := range ip {
		fmt.Println(val)
	}

}

func main() {

	fmt.Println("start server\n")

	file, err := os.Open("iplist.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// bufio.NewScanner() передаем прочитанный файл
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

	// for _, val := range iplines {
	// 	fmt.Println(val)
	// }

	ping(iplines...)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
