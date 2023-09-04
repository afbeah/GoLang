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

const( 
    monitoramentos = 3 
    delay = 3
)

func main(){

    exibeIntro()
    
    for {
        exibeMenu()
        comando:= leComando()

        switch comando {
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo logs...")
            imprimeLogs()
        case 0:
            fmt.Println("Saindo do Programa.")
            os.Exit(0)
        default:
            fmt.Println("Comando Inexistente!")
            os.Exit(-1)     
        }
    }
}

func exibeIntro(){
    nome := "Vegetti"
    versao := 1.1
    fmt.Println("Olá, sr:", nome)
    fmt.Println("Este programa está na versão", versao)
} 

func exibeMenu(){
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir logs")
    fmt.Println("0- Sair do Programa")
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O Comando escolhido foi:", comandoLido)
    fmt.Println("")

    return comandoLido
}

func iniciarMonitoramento(){
    fmt.Println("Monitorando...")
    
    //fmt.Println("O meu slice tem", len(sites),"itens. E tem capacidade", cap(sites))

    sites := leSitesDoArquivo()

    for i:=0; i < monitoramentos ; i++{
        for i, site := range sites{
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Second)
          
    }
    fmt.Println("")
}

func testaSite(site string){
    resp, err := http.Get(site)
     
    if err != nil{
        fmt.Println("Ocorreu um erro!", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site:",site, "foi carregado com sucesso!")
        registraLog(site, true)
    }else{
        fmt.Println("Site:", site, "ocorreu um problema.")
        registraLog(site, false) 
    }

    fmt.Println("")   
}

func leSitesDoArquivo() []string{
    
    var sites [] string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro.", err)
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
    
    arquivo.Close()
    return sites
}

func registraLog(site string, status bool){
    arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

    if err != nil{
        fmt.Println(err)
    }
    
    arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}

func imprimeLogs(){
    
    arquivo, err := ioutil.ReadFile("log.txt")

    if err != nil{
        fmt.Println(err)
    }

    fmt.Println(string(arquivo))
}