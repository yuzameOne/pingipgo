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

		fmt.Printf("ip array : %s", val) //dedug out

		out, _ := exec.Command("ping", val, "-c 2", "-i 1", "-w 2").Output()

		fmt.Println(string(string(out))) //dedug out

		if strings.Contains(string(out), "2 received") {
			fmt.Println("Yeap, I'am ALIVEEE")
			lifeIp = append(lifeIp, val)
		} else {
			fmt.Println("Dead Mother Fucker")
			deadIp = append(deadIp, val)
		}
	}

	fmt.Println(lifeIp) //dedug out
	fmt.Printf("\n")    //dedug out
	fmt.Println(deadIp) //dedug out

}

func saveIpFile() {
	file, err := os.OpenFile("lifeIp.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("file not create : %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, val := range lifeIp {
		datawriter.WriteString(val)
		datawriter.WriteString("\n")
	}

	datawriter.Flush()
	file.Close()
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

	ping(iplines...)

	saveIpFile()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
