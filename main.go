package main

import (
    "fmt"
    "log"
    "net"
    "flag"
)


func flow_collector(packet []byte, maxHeader int) {
    // Verifica se o tamanho do pacote é suficiente para conter o cabeçalho mínimo
    if len(packet) < 208 {
        fmt.Println("Pacote muito pequeno para ser um pacote sFlow válido")
        return
    }
    
    // Extrair os campos do cabeçalho de acordo com a estrutura fornecida
    genericHeader(&packet)

    // Verificar o tipo de flow
    sflow_type := checkFlowType(&packet)

	switch sflow_type {
        // se o flow type for igual a 3, então é um flow sample
        case 3:
            flowSample(&packet, maxHeader)
	    // se o flow type for igual a 4, então é um flow interval
        case 4:
			flowInterval(&packet)
		default:
			fmt.Println("Tipo de flow desconhecido")
	}
}

func main() {

    ip := flag.String("ip", "0.0.0.0", "Endereço IP para escutar")
    port := flag.Int("port", 6343, "Porta para escutar")
    maxHeader := flag.Int("maxHeader", 512, "Tamanho máximo do cabeçalho coletado por Amostragem")

    // Parse as flags
    flag.Parse()
    
    addr := net.UDPAddr{
        Port: *port, // Usa a porta passada como argumento
        IP:   net.ParseIP(*ip), // Usa o IP passado como argumento
    }


    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        log.Fatalf("Erro ao iniciar o servidor UDP: %v", err)
    }
    defer conn.Close()

    fmt.Printf("Servidor sFlow ouvindo no endereço %s:%d\n", *ip, *port)
    fmt.Printf("Configurado o max-header de: %d\n", *maxHeader)

    buf := make([]byte, 2048) // Buffer para armazenar pacotes recebidos
    for {
        n, _, err := conn.ReadFromUDP(buf)
        if err != nil {
            log.Printf("Erro ao ler o pacote UDP: %v", err)
            continue
        }

        // Função para processar o pacote sFlow e imprimir os contadores
        flow_collector(buf[:n], *maxHeader)
    }
}
