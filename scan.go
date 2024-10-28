package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main(){
	saudacoes()
	exibeMenu()
	opcao := obterOpcao()

	switch opcao{
	case 1:
		fmt.Println("Escaneando sites...")
		escanearSites(abreArquivo())
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
func abreArquivo() []string{
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil{
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(-1)
	}

	leitor := bufio.NewReader(arquivo)

	for{
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
	}
	arquivo.Close()
	return sites
}

func escanearSites(sites []string){
	for i:= 0; i < 5; i++{
		fmt.Println("teste", i+1)
		for _, site := range sites{
				testaSite(site)
				fmt.Println()
				time.Sleep(5 * time.Second)
			}		
	}
}

func testaSite(site string){
	resp, err := http.Get(site)

		if err != nil{
			fmt.Println("Ocorreu um erro:", err)
			os.Exit(-1)
		}

		if(resp.StatusCode == 200){
			fmt.Println("O site", site, "está online!")
			fmt.Println("status:", resp.StatusCode)
		} else{
			fmt.Println("O site", site, "está fora do ar!")
			fmt.Println("status:", resp.StatusCode)
		}
}