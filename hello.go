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

const monitoramentos = 2
const delay = 2

func main(){
	exibirIntroducao()
for {
	exibeMenu()
	comando := leComando()
	
	switch comando {
	case 1:
		inciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
		imprimeLogs()
	case 0: 
		fmt.Println("Saindo do Prorama...")
		os.Exit(0)
	default: 
		fmt.Println("Não conheço esse comando...")
		os.Exit(1)
	}
}
}

func exibirIntroducao() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá Sr(a).", nome)
	fmt.Println("Este programa estar na verão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O camando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func inciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Minute)
	}
	fmt.Println("")
}

func testaSite(site string){
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi careagado com sucesso !!!")
		registaLog(site, true)
	}else {
		fmt.Println("Site:", site, "esta com problema. Status Code:", resp.StatusCode)
		registaLog(site, false)
	}
}

func leSitesDoArquivo() []string {

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

		fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites

}

func registaLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE| os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string (arquivo))
}