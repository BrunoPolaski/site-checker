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

func main() {

	welcome()

	for {
		command := choseCommand()

		switch command {
		case 1:
			initMonitor()
			break
		case 2:
			fmt.Println("Exibindo logs...")
			printLogs()
			break
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
			break
		default:
			fmt.Println("Não conheço este comando")
			break
		}
	}
}

func welcome() {
	version := 1.0
	fmt.Println("Qual o seu nome?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Olá,", name)
	fmt.Println("Versão do programa:", version)
}

func choseCommand() int {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	return command
}

func initMonitor() {
	fmt.Println("Monitorando...")

	sites := readSitesFromFile()

	for i := 0; i < 2; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(5 * time.Second)
	}
}

func testSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("An error has occurred: ", err)
	}
	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registerLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		registerLogs(site, false)
	}
	fmt.Println("------------------------------------------------")
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error has occurred: ", err)
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

	return sites
}

func registerLogs(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " status " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
