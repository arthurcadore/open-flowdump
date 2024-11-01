package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// Estrutura para armazenar os dados do dispositivo
type SFlowData struct {
    SourceIP string
}

// Função para conectar ao MySQL
func connectToDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "sflowadmin:SflowAdmin123@tcp(127.0.0.1:3306)/sflowdb")
    if err != nil {
        return nil, err
    }

    // Testa a conexão
    if err = db.Ping(); err != nil {
        return nil, err
    }
    fmt.Println("Conectado ao MySQL!")
    return db, nil
}

// Função para inserir dados no banco
func insertSFlowData(db *sql.DB, data SFlowData) error {
    query := `INSERT INTO devices (ip_address) VALUES (?)`
    _, err := db.Exec(query, data.SourceIP)
    if err != nil {
        return err
    }
    return nil
}

func connect() {
    // Conecta ao banco
    db, err := connectToDB()
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }
    defer db.Close()

    // Exemplo de dados de sFlow
    sflowData := SFlowData{
        SourceIP: "192.168.1.1",
    }

    // Insere os dados no banco
    err = insertSFlowData(db, sflowData)
    if err != nil {
        log.Fatalf("Erro ao inserir dados de sFlow: %v", err)
    }

    fmt.Println("Dados de sFlow inseridos com sucesso!")
}
