package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 1

func main() {

	introduction()
	for {
		menu()

		switch readCommand() {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs...")
			printLogs()
		case 0:
			fmt.Println("exit...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func introduction() {
	name := "Thalys"
	years := 30
	version := 1.1

	fmt.Println("Hello world", name, "you are", years)
	fmt.Println("This program is on the version", version)
}

func menu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")
	fmt.Println()
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen option was", command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	// sites := []string{}
	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfully loaded")
		storeLog(site, true)
	} else {
		fmt.Println("Site:", site, "is with some problem. Status code:", resp.StatusCode)
		storeLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()
	fmt.Println(sites)
	return sites
}

func storeLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("erro:", err)
	}

	fmt.Println(string(file))
}
