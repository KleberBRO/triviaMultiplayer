// entry point do servidor TCP
package main

import (
    "fmt"
    "net"
    "log"
    "bufio"
    "strings"
)

func main() {
    // Cria um listener TCP na porta 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal("Erro ao criar listener:", err)
    }
    defer listener.Close()
    
    fmt.Println("Servidor escutando na porta 8080...")
    
    for {
        // Aceita uma nova conexão
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Erro ao aceitar conexão:", err)
            continue
        }
        
        // Trata a conexão em uma goroutine separada
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    clientAddr := conn.RemoteAddr().String()
    fmt.Printf("Nova conexão de: %s\n", clientAddr)
    
    scanner := bufio.NewScanner(conn)
    
    for scanner.Scan() {
        message := strings.TrimSpace(scanner.Text())
        fmt.Printf("[%s] Recebido: %s\n", clientAddr, message)
        
        // Responde aos pings com pongs
        switch message {
        case "ping":
            response := "pong\n"
            conn.Write([]byte(response))
            fmt.Printf("[%s] Enviado: pong\n", clientAddr)
        case "hello":
            response := "hello\n"
            conn.Write([]byte(response))
            fmt.Printf("[%s] Enviado: hello\n", clientAddr)
        case "quit":
            fmt.Printf("[%s] Cliente desconectando\n", clientAddr)
            return
        default:
            response := fmt.Sprintf("Mensagem recebida: %s\n", message)
            conn.Write([]byte(response))
        }
    }
    
    if err := scanner.Err(); err != nil {
        log.Printf("[%s] Erro ao ler dados: %v\n", clientAddr, err)
    }
    
    fmt.Printf("[%s] Conexão encerrada\n", clientAddr)
}