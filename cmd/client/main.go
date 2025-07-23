// entry point do cliente CLI/GUI
package main

import (
    "fmt"
    "log"
    "time"
    "triviaMultiplayer/internal/client"
)

func main() {
    // Conecta ao servidor
    c, err := client.NewClient("localhost:8080")
    if err != nil {
        log.Fatal("Erro ao conectar ao servidor:", err)
    }
    defer c.Close()
    
    fmt.Println("Conectado ao servidor!")
    
    // Envia hello inicial
    if err := c.SendMessage("hello"); err != nil {
        log.Fatal("Erro ao enviar hello:", err)
    }
    
    // Lê resposta do hello
    response, err := c.ReadMessage()
    if err != nil {
        log.Fatal("Erro ao ler resposta:", err)
    }
    fmt.Printf("Servidor respondeu: %s\n", response)
    
    // Inicia loop de ping/pong a cada 5 segundos
    fmt.Println("Iniciando ping/pong...")
    go c.StartPingPongLoop(5 * time.Second)
    
    // Mantém o cliente rodando por 30 segundos como exemplo
    time.Sleep(30 * time.Second)
    
    // Envia quit antes de sair
    c.SendMessage("quit")
    fmt.Println("Cliente encerrando...")
}