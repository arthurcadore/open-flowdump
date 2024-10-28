package main

import (
    "encoding/binary"
    "fmt"
    "log"
    "net"
)

func flow_counter(packet []byte) {

} 

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
			flow_counter(packet)
		case 4:
			flow_interval(packet)
		default:
			fmt.Println("Tipo de flow desconhecido")
	}


}

func flow_interval(packet []byte) {

    // Expanded counters sample
    sample_length := binary.BigEndian.Uint32(packet[32:36])
    sequence_number := binary.BigEndian.Uint32(packet[36:40])
    source_id_type := binary.BigEndian.Uint32(packet[40:44])
    source_id_index := binary.BigEndian.Uint32(packet[44:48])
    counters_records := binary.BigEndian.Uint32(packet[48:52])

    // Generic Interface Counters
    flow_type := binary.BigEndian.Uint32(packet[52:56])
    flow_data_length := binary.BigEndian.Uint32(packet[56:60])
    if_id := binary.BigEndian.Uint32(packet[60:64])
    if_type := binary.BigEndian.Uint32(packet[64:68])
    if_speed := binary.BigEndian.Uint64(packet[68:76])
    if_direction := binary.BigEndian.Uint32(packet[76:80])
    if_status := binary.BigEndian.Uint32(packet[80:84])

    // Interface Counters
    if_in_octets := binary.BigEndian.Uint64(packet[84:92])
    if_in_pkt := binary.BigEndian.Uint32(packet[92:96])
    if_in_multicast := binary.BigEndian.Uint32(packet[96:100])
    if_in_broadcast := binary.BigEndian.Uint32(packet[100:104])
    if_in_discards := binary.BigEndian.Uint32(packet[104:108])
    if_in_errors := binary.BigEndian.Uint32(packet[108:112])
    if_in_unknown := binary.BigEndian.Uint32(packet[112:116])

    if_out_octets := binary.BigEndian.Uint64(packet[116:124])
    if_out_pkt := binary.BigEndian.Uint32(packet[124:128])
    if_out_multicast := binary.BigEndian.Uint32(packet[128:132])
    if_out_broadcast := binary.BigEndian.Uint32(packet[132:136])
    if_out_discards := binary.BigEndian.Uint32(packet[136:140])
    if_out_errors := binary.BigEndian.Uint32(packet[140:144])

    // Interface mode
    if_promiscuous := binary.BigEndian.Uint32(packet[144:148])

    // Ethernet Counters
    eth_flow_type := binary.BigEndian.Uint32(packet[148:152])
    eth_flow_data_length := binary.BigEndian.Uint32(packet[152:156])
    align := binary.BigEndian.Uint32(packet[156:160])
    fcs := binary.BigEndian.Uint32(packet[160:164])
    single_collision := binary.BigEndian.Uint32(packet[164:168])
    multiple_collision := binary.BigEndian.Uint32(packet[168:172])
    sqe_test_errors := binary.BigEndian.Uint32(packet[172:176])
    deferred_transmissions := binary.BigEndian.Uint32(packet[176:180])
    late_collisions := binary.BigEndian.Uint32(packet[180:184])
    excessive_collisions := binary.BigEndian.Uint32(packet[184:188])
    internal_mac_transmit_errors := binary.BigEndian.Uint32(packet[188:192])
    carrier_sense_errors := binary.BigEndian.Uint32(packet[192:196])
    frame_too_longs := binary.BigEndian.Uint32(packet[196:200])
    internal_mac_receive_errors := binary.BigEndian.Uint32(packet[200:204])
    symbol_errors := binary.BigEndian.Uint32(packet[204:208])

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

        // Adicione mais impressões conforme necessário para outros campos
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
