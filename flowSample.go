// flow_counter.go
package main

import (
    "net"
    "fmt"
    "encoding/binary"
)



// Defina a função flowCounter
func flowSample(packet []byte) {
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

    fmt.Printf("Datagram Version: %d, IP Version: %d, Agent IP: %s\n", datagramVersion, ipVersion, agentIP)
    fmt.Printf("Sub-Agent ID: %d, Sequence Number: %d, Sys Uptime: %d, Number of Samples: %d\n", subAgentID, sequenceNumber, sysUptime, numSamples)
}
