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

const qtdMonitoramentos = 3
const delayMonitoramentos = 2 // em segundos
const arquivoSites = "sites.txt"
const arquivoLog = "log.txt"
const dateTimeFormat = "02/01/2006 15:04:05"

func main() {

	showGreetings()

	for {
		showMenu()

		comando := getCommand()

		switch comando {
		case 1:
			initMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Unknown command (", comando, ")")
			os.Exit(-1)
		}
	}
}

func showGreetings() {

	nome := "Jonnes" // declara uma variavel e pede para o Go inferir o tipo
	versao := 1.2

	fmt.Println("\nHello Mr.", nome)
	fmt.Println("Program version is", versao)
}

func getCommand() int {

	var comando int

	fmt.Print("Command > ")
	fmt.Scanf("%d", &comando)
	fmt.Println()

	return comando
}

func showMenu() {

	fmt.Println("\nCommand Menu")
	fmt.Println("==========================")
	fmt.Println("( 1 ) Start sites monitoring")
	fmt.Println("( 2 ) Show logs")
	fmt.Println("( 0 ) Quit")
	fmt.Println("==========================")
}

func initMonitoring() {

	fmt.Println("Monitoring...")

	sites := loadSitesFile()

	if sites != nil {
		for i := 0; i < qtdMonitoramentos; i++ {
			for _, site := range sites {

				resp, err := http.Get(site) // o _ diz ao Go que nao estou interessado no segundo retorno (que seria o erro)

				if err != nil {
					fmt.Println("[ERROR]", err)
				}

				if resp.StatusCode == 200 {
					fmt.Println("Site:", site, "sucessfully loaded!")
					log(site, true)
				} else {
					fmt.Println("Error loading", site, ". Error Code:", resp.StatusCode)
					log(site, false)
				}
			}

			time.Sleep(delayMonitoramentos * time.Second)
			fmt.Println("")
		}
	} else {
		fmt.Println("Nothing to do.")
	}
}

func loadSitesFile() []string {

	var sites []string

	arq, err := os.Open(arquivoSites)

	if err != nil {
		fmt.Println("[ERROR]", err)
	} else {

		leitor := bufio.NewReader(arq)

		for {
			linha, err := leitor.ReadString('\n')

			linha = strings.TrimSpace(linha)
			sites = append(sites, linha)

			if err == io.EOF {
				break
			}
		}
	}

	arq.Close()

	return sites
}

func log(site string, online bool) {

	arq, err := os.OpenFile(arquivoLog, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arq.WriteString(time.Now().Format(dateTimeFormat) + " - " + site + " - online: " + strconv.FormatBool(online) + "\n")

	arq.Close()
}

func showLogs() {

	fmt.Println("Showing logs...")

	arq, err := ioutil.ReadFile(arquivoLog)

	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	fmt.Println(string(arq))
}
