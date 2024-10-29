package main

import (
    "encoding/binary"
    "fmt"
    "log"
    "net"
)

func flow_collector(packet []byte) {
    // Verifica se o tamanho do pacote é suficiente para conter o cabeçalho mínimo
    if len(packet) < 208 {
        fmt.Println("Pacote muito pequeno para ser um pacote sFlow válido")
        return
    }

    // Extrair os campos do cabeçalho de acordo com a estrutura fornecida
    datagramVersion := binary.BigEndian.Uint32(packet[0:4])
    ipVersion := binary.BigEndian.Uint32(packet[4:8])
    agentIP := net.IP(packet[8:12])
    subAgentID := binary.BigEndian.Uint32(packet[12:16])
    sequenceNumber := binary.BigEndian.Uint32(packet[16:20])
    sysUptime := binary.BigEndian.Uint32(packet[20:24])
    numSamples := binary.BigEndian.Uint32(packet[24:28])

	// Imprimir todos os valores dos campos
	fmt.Printf("-------------------------\n")
	fmt.Printf("Versão do Datagram: %d\n", datagramVersion)
	fmt.Printf("Versão do IP: %d\n", ipVersion)
	fmt.Printf("Endereço IP do Agente: %s\n", agentIP)
	fmt.Printf("ID do Sub-Agente: %d\n", subAgentID)
	fmt.Printf("Número de Sequência: %d\n", sequenceNumber)
	fmt.Printf("SysUptime (secs): %d\n", sysUptime/1000)
	fmt.Printf("Número de Amostras: %d\n", numSamples)

	// verificar o tipo de flow
	sflow_type := binary.BigEndian.Uint32(packet[28:32])
	fmt.Printf("Tipo de flow: %d\n", sflow_type)


	// se o flow type for igual a 4, então é um flow interval
	// se o flow type for igual a 3, então é um counter sample
	
	switch sflow_type {
		case 3:
			flowSample(packet)
		case 4:
			flowInterval(packet)
		default:
			fmt.Println("Tipo de flow desconhecido")
	}


}

func main() {
    addr := net.UDPAddr{
        Port: 6343, // Porta padrão para sFlow
        IP:   net.ParseIP("0.0.0.0"),
    }

    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        log.Fatalf("Erro ao iniciar o servidor UDP: %v", err)
    }
    defer conn.Close()

    fmt.Println("Servidor sFlow ouvindo na porta 6343...")

    buf := make([]byte, 2048) // Buffer para armazenar pacotes recebidos
    for {
        n, _, err := conn.ReadFromUDP(buf)
        if err != nil {
            log.Printf("Erro ao ler o pacote UDP: %v", err)
            continue
        }

        // Função para processar o pacote sFlow e imprimir os contadores
        flow_collector(buf[:n])
    }
}
