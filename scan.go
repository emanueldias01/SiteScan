package main

import (
	"fmt"
	"os"
)

func main(){
	saudacoes()
	exibeMenu()
	opcao := obterOpcao()

	switch opcao{
	case 1:
		fmt.Println("Escaneando sites...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}

}

func saudacoes(){
	versao := 0.1
	asciiArt := `
.+"+.+"+.+"+.+"+. 
(    SITESCAN    )
 )              ( 
(                )
 "+.+"+.+"+.+"+.+"                               
    `

    fmt.Println(asciiArt)
	fmt.Println("versão:", versao)
}

func exibeMenu(){
	fmt.Println("1- escanear sites")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func obterOpcao() int{
	var opcao int
	fmt.Scan(&opcao)
	return opcao
}