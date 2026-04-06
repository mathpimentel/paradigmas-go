package main

import (
	"fmt"
)

// Definindo a estrutura da transação (Paradigma Imperativo - Estrutura de Dados)
type Transacao struct {
	ID    int
	Valor float64
	Tipo  string // "Crédito" ou "Débito"
}

func main() {
	// 1. Entrada: Slice com 20 transações fixas
	transacoes := []Transacao{
		{1, 100.50, "Crédito"}, {2, 50.00, "Débito"}, {3, 200.00, "Crédito"},
		{4, 30.00, "Débito"}, {5, 150.00, "Crédito"}, {6, 10.00, "Débito"},
		{7, 500.00, "Crédito"}, {8, 100.00, "Débito"}, {9, 25.00, "Crédito"},
		{10, 5.00, "Débito"}, {11, 300.00, "Crédito"}, {12, 120.00, "Débito"},
		{13, 80.00, "Crédito"}, {14, 40.00, "Débito"}, {15, 90.00, "Crédito"},
		{16, 200.00, "Débito"}, {17, 1000.00, "Crédito"}, {18, 500.00, "Débito"},
		{19, 45.00, "Crédito"}, {20, 15.00, "Débito"},
	}

	// Canais para comunicação entre as Goroutines e a função principal
	canalCreditos := make(chan float64)
	canalDebitos := make(chan float64)

	// 2. Processamento Concorrente: Goroutine para CRÉDITOS
	go func() {
		somaCreditos := 0.0
		// Iteração imperativa (for + if)
		for i := 0; i < len(transacoes); i++ {
			if transacoes[i].Tipo == "Crédito" {
				somaCreditos += transacoes[i].Valor
			}
		}
		canalCreditos <- somaCreditos // Envia o resultado total
	}()

	// 3. Processamento Concorrente: Goroutine para DÉBITOS
	go func() {
		somaDebitos := 0.0
		// Iteração imperativa (for + if)
		for i := 0; i < len(transacoes); i++ {
			if transacoes[i].Tipo == "Débito" {
				somaDebitos += transacoes[i].Valor
			}
		}
		canalDebitos <- somaDebitos // Envia o resultado total
	}()

	// 4. Sincronização e Cálculo do Saldo Final na rotina principal
	totalCredito := <-canalCreditos
	totalDebito := <-canalDebitos
	saldoFinal := totalCredito - totalDebito

	// Exibição dos resultados
	fmt.Println("--- Relatório de Processamento Bancário ---")
	fmt.Printf("Total Processado em Créditos: R$ %.2f\n", totalCredito)
	fmt.Printf("Total Processado em Débitos:  R$ %.2f\n", totalDebito)
	fmt.Println("-------------------------------------------")
	fmt.Printf("SALDO FINAL DA CONTA:         R$ %.2f\n", saldoFinal)
}