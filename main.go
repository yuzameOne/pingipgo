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

var lifeIp []string
var deadIp []string

func ping(ip ...string) {

	// c - sending count ECHO_REQUEST packets
	// i - Wait interval seconds between sending each packet
	// Only super-user may set interval to values less than 0.2 seconds
	// w - deadline Specify a timeout, in seconds, before ping exits regardless of how
	// many packets have been sent or received.

	for _, val := range ip {

		out, _ := exec.Command("ping", val, "-c 2", "-i 1", "-w 2").Output()

		if strings.Contains(string(out), "2 received") {
			fmt.Printf("Yeap, I'am ALIVEEE : %s \n", val)
			lifeIp = append(lifeIp, val)
		} else {
			fmt.Printf("Dead Mother Fucker : %s \n", val)
			deadIp = append(deadIp, val)
		}
	}
}

func saveIpFile() {
	file, err := os.OpenFile("lifeIp.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("file not create : %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, val := range lifeIp {
		datawriter.WriteString(val + "\n")
	}

	datawriter.Flush()
	file.Close()

	os.Exit(3)
}

func main() {

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

	ping(iplines...)

	saveIpFile()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
