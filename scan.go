package main

import (
	"fmt"
	"os"
	"net/http"
)

func main(){
	saudacoes()
	exibeMenu()
	opcao := obterOpcao()

	switch opcao{
	case 1:
		fmt.Println("Escaneando sites...")
		escanearSites()
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

func escanearSites(){
	sites := []string{"https://youtube.com", "https://instagram.com","https://httpbin.org/status/404"}

	for _, site := range sites{
		resp, err := http.Get(site)

		if err != nil{
			fmt.Println("Ocorreu um erro:", err)
		}

		if(resp.StatusCode == 200){
			fmt.Println("O site", site, "está online!")
		} else{
			fmt.Println("O site", site, "está fora do ar!")
		}
	}
}