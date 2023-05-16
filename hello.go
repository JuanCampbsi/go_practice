package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	viewIntrod()
	for {
		viewMenu()
		comand := readComand()

		switch comand {
		case 1:
			initialMonitoring()
		case 2:
			viewLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}
func initialMonitoring() {
	conditionClearConsole()
	fmt.Println("Monitorando....")
	site := readSitesFiles()
	isVerifiyMonitoring := false

	for i := 0; i < monitoring; i++ {
		if isVerifiyMonitoring {
			fmt.Println("Monitorando....")
		}
		for _, url := range site {
			fmt.Println("\nTestando site...")
			testSite(url)
		}
		isVerifiyMonitoring = true
		time.Sleep(delay * time.Second)
		conditionClearConsole()
	}

	fmt.Println("\nTamanho do array", len(site))
	fmt.Println("Capacidade do array", cap(site))

	site = append(site, "https://www.google.com/")
}

func viewIntrod() {
	name := "Juan"
	version := 1.1
	age := 28

	fmt.Printf("\nHello Mr. %s. This age %d \n", name, age)
	fmt.Println("This is system of version", version)
}

func viewMenu() {
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func readComand() int {
	var comand int
	fmt.Scan(&comand)
	fmt.Println("\nO comando escolhido foi", comand)

	return comand
}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso! Status code:", resp.StatusCode)
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func conditionClearConsole() {
	clearCmd := ""
	if runtime.GOOS == "windows" {
		clearCmd = "cls"
	} else {
		clearCmd = "clear"
	}
	cmd := exec.Command(clearCmd)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readSitesFiles() []string {
	var sites []string
	files, err := os.Open("helloApp/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	read := bufio.NewReader(files)
	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	files.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("helloApp/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func viewLog() {
	fmt.Println("Exibindo Logs...")
	file, err := ioutil.ReadFile("helloApp/log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
