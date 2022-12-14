package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const VERSION = "v1.0"
const CICLES = 2
const DELAY = 5

func main() {
	clearScreen()
	fmt.Println("Este programa está na versão:", VERSION)
	fmt.Println("Salve salve Aurigod")

	startingMonitoring()
	os.Exit(0)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func executeRequest(site string) string {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Vishi deu ruim erro:", err)
	}

	if resp.StatusCode == 200 {
		return "Site " + site + " carregado com sucesso"
	}

	return "Site " + site + " está com problemas. status code:" + strconv.Itoa(resp.StatusCode)
}

func readSiteInFile() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	return sites
}

func startingMonitoring() {
	clearScreen()
	fmt.Println("\nIniciando monitoramento...")

	sites := readSiteInFile()

	for i := 0; i < CICLES; i++ {
		for _, site := range sites {
			fmt.Println("\nAcessando site:", site)

			statusMessage := executeRequest(site)
			fmt.Println(statusMessage)
		}

		time.Sleep(DELAY * time.Second)
	}
}
