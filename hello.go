package main

import (
	"fmt"
	"os"
) 

func main(){
	exibirIntroducao()
	exibeMenu()

	//formas de você printar na tela 
	//Com o Scanf preciso passar as duas informações modificador(%d) e o local(&comando)
	//var comando int
	//fmt.Scanf("%d", &comando)

	//Mais tem um meio mais simples usando uma outra função do Go o fmt.Scan()
	//Para saber qual é o tipo da variável  precisa importar um packge "reflect" e daí usa-lo como mostra abiaxo

	comando := leComando()
	
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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

func exibirIntroducao() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá Sr(a).", nome)
	fmt.Println("Este programa estar na verão", versao)
}

func leComando() int {
	var comandoLido int
	//reflect.TypeOf(comando)
	fmt.Scan(&comandoLido)
	fmt.Println("O camando escolhido foi", comandoLido)
	//fmt.Println("O tipo da variável é", reflect.TypeOf(comando))

	return comandoLido
}

func exibeMenu() {
	//Menu de opçoes para sselecionar na aplicação
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}