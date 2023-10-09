package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
) 

const monitoramentos = 5
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
	sites := []string{"https://www.alura.com.br", "https://ramdom-status-code.herokuapp.com/", "https://www.caelum.com.br"}

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
	resp, _ := http.Get(site)
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi careagado com sucesso !!!")
	}else {
		fmt.Println("Site:", site, "esta com problema. Status Code:", resp.StatusCode)
	}
}