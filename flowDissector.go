// flow_counter.go
package main

import (
    "encoding/binary"
    "fmt"
    "net"
)

// Função auxiliar para "retirar" bytes do pacote
func popBytes(packet *[]byte, numBytes int) []byte {
    if len(*packet) < numBytes {
        return nil // Retorna nil se o pacote for menor que o número de bytes desejado
    }
    value := (*packet)[:numBytes]
    *packet = (*packet)[numBytes:]
    return value
}

// Funções para ler tipos específicos de dados
func popUint32(packet *[]byte) uint32 {
    bytes := popBytes(packet, 4)
    if bytes == nil {
        return 0
    }
    return binary.BigEndian.Uint32(bytes)
}

func popUint64(packet *[]byte) uint64 {
    bytes := popBytes(packet, 8)
    if bytes == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bytes)
}

func genericHeader(packet *[]byte) {

    // Extrair os campos do cabeçalho de acordo com a estrutura fornecida
    datagramVersion := popUint32(packet)
    ipVersion := popUint32(packet)

    // Obter o agente IP como um uint32 e converter para net.IP
    agentIPBytes := make([]byte, 4) // Criar um slice de 4 bytes
    binary.BigEndian.PutUint32(agentIPBytes, popUint32(packet)) // Colocar o uint32 no slice de bytes
    agentIP := net.IP(agentIPBytes)

    subAgentID := popUint32(packet)
    sequenceNumber := popUint32(packet)
    sysUptime := popUint32(packet)
    numSamples := popUint32(packet)

	// Imprimir todos os valores dos campos
	fmt.Printf("-------------------------\n")
	fmt.Printf("Versão do Datagram: %d\n", datagramVersion)
	fmt.Printf("Versão do IP: %d\n", ipVersion)
	fmt.Printf("Endereço IP do Agente: %s\n", agentIP)
	fmt.Printf("ID do Sub-Agente: %d\n", subAgentID)
	fmt.Printf("Número de Sequência: %d\n", sequenceNumber)
	fmt.Printf("SysUptime (secs): %d\n", sysUptime/1000)
	fmt.Printf("Número de Amostras: %d\n", numSamples)

} 

// Função para verificar o tipo de flow, e retonar um inteiro correspondente ao tipo de flow
func checkFlowType(packet *[]byte) int{
    // Verificar o tipo de flow
    flowType := popUint32(packet)
    fmt.Printf("Tipo de flow: %d\n", flowType)

    return int(flowType)
}


// Defina a função flowCounter
func flowInterval(packet *[]byte) {

    // Expanded counters sample
    sample_length := popUint32(packet)
    sequence_number := popUint32(packet)
    source_id_type := popUint32(packet)
    source_id_index := popUint32(packet)
    counters_records := popUint32(packet)

    // Generic Interface Counters
    flow_type := popUint32(packet)
    flow_data_length := popUint32(packet)
    if_id := popUint32(packet)
    if_type := popUint32(packet)
    if_speed := popUint64(packet)
    if_direction := popUint32(packet)
    if_status := popUint32(packet)

    // Interface Counters
    if_in_octets := popUint64(packet)
    if_in_pkt := popUint32(packet)
    if_in_multicast := popUint32(packet)
    if_in_broadcast := popUint32(packet)
    if_in_discards := popUint32(packet)
    if_in_errors := popUint32(packet)
    if_in_unknown := popUint32(packet)

    if_out_octets := popUint64(packet)
    if_out_pkt := popUint32(packet)
    if_out_multicast := popUint32(packet)
    if_out_broadcast := popUint32(packet)
    if_out_discards := popUint32(packet)
    if_out_errors := popUint32(packet)

    // Interface mode
    if_promiscuous := popUint32(packet)

    // Ethernet Counters
    eth_flow_type := popUint32(packet)
    eth_flow_data_length := popUint32(packet)
    align := popUint32(packet)
    fcs := popUint32(packet)
    single_collision := popUint32(packet)
    multiple_collision := popUint32(packet)
    sqe_test_errors := popUint32(packet)
    deferred_transmissions := popUint32(packet)
    late_collisions := popUint32(packet)
    excessive_collisions := popUint32(packet)
    internal_mac_transmit_errors := popUint32(packet)
    carrier_sense_errors := popUint32(packet)
    frame_too_longs := popUint32(packet)
    internal_mac_receive_errors := popUint32(packet)
    symbol_errors := popUint32(packet)

    // Verificar se os contadores de entrada e saída são diferentes de zero
    if if_in_octets != 0 || if_out_octets != 0 {

		// imprimir os valores genéricos: 
		fmt.Printf("Comprimento da captura: %d\n", sample_length)
		fmt.Printf("Sequence Number: %d\n", sequence_number)
		fmt.Printf("Source ID Type: %d\n", source_id_type)
		fmt.Printf("Source ID Index: %d\n", source_id_index)
		fmt.Printf("Counters Recorded: %d\n", counters_records)

		fmt.Printf("Tipo de flow: %d\n", flow_type)
		fmt.Printf("Comprimento da captura: %d\n", flow_data_length)
		fmt.Printf("ID da Interface: %d\n", if_id)
		fmt.Printf("Tipo de Interface: %d\n", if_type)
		fmt.Printf("Velocidade da Interface: %d\n", if_speed)
		fmt.Printf("Direção da Interface: %d\n", if_direction)
		fmt.Printf("Status da Interface: %d\n", if_status)
		fmt.Printf("ifPromiscuous: %d\n", if_promiscuous)


        // Imprimir contadores de interface
        fmt.Println("Contadores de Interface:")
        fmt.Printf("ifInOctets: %d\n", if_in_octets)
        fmt.Printf("ifInPackets: %d\n", if_in_pkt)
        fmt.Printf("ifInMulticast: %d\n", if_in_multicast)
        fmt.Printf("ifInBroadcast: %d\n", if_in_broadcast)
        fmt.Printf("ifInDiscards: %d\n", if_in_discards)
        fmt.Printf("ifInErrors: %d\n", if_in_errors)
        fmt.Printf("ifInUnknown: %d\n", if_in_unknown)
        fmt.Printf("ifOutOctets: %d\n", if_out_octets)
        fmt.Printf("ifOutPackets: %d\n", if_out_pkt)
        fmt.Printf("ifOutMulticast: %d\n", if_out_multicast)
        fmt.Printf("ifOutBroadcast: %d\n", if_out_broadcast)
        fmt.Printf("ifOutDiscards: %d\n", if_out_discards)
        fmt.Printf("ifOutErrors: %d\n", if_out_errors)

        // Imprimir contadores Ethernet
        fmt.Println("Contadores Ethernet:")
        fmt.Printf("ethFlowType: %d\n", eth_flow_type)
        fmt.Printf("ethFlowDataLength: %d\n", eth_flow_data_length)
        fmt.Printf("Align: %d\n", align)
        fmt.Printf("FCS: %d\n", fcs)
        fmt.Printf("Single Collisions: %d\n", single_collision)
        fmt.Printf("Multiple Collisions: %d\n", multiple_collision)
        fmt.Printf("SQE Test Errors: %d\n", sqe_test_errors)
        fmt.Printf("Deferred Transmissions: %d\n", deferred_transmissions)
        fmt.Printf("Late Collisions: %d\n", late_collisions)
        fmt.Printf("Excessive Collisions: %d\n", excessive_collisions)
        fmt.Printf("Internal MAC Transmit Errors: %d\n", internal_mac_transmit_errors)
        fmt.Printf("Carrier Sense Errors: %d\n", carrier_sense_errors)
        fmt.Printf("Frame Too Longs: %d\n", frame_too_longs)
        fmt.Printf("Internal MAC Receive Errors: %d\n", internal_mac_receive_errors)
        fmt.Printf("Symbol Errors: %d\n", symbol_errors)
    }
}


// Defina a função flowCounter
func flowSample(packet *[]byte, maxHeader int) {
}