package main

import (
	"fmt"
	"sync"
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

	// Canais para comunicação entre as Goroutines
	canalCreditos := make(chan float64)
	canalDebitos := make(chan float64)

	// WaitGroup para sincronizar o encerramento das Goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// 2. Processamento Concorrente - Abordagem Imperativa (uso de loops e condicionais)

	// Goroutine para Créditos
	go func() {
		defer wg.Done()
		var soma float64 = 0
		for i := 0; i < len(transacoes); i++ {
			if transacoes[i].Tipo == "Crédito" {
				soma += transacoes[i].Valor
			}
		}
		canalCreditos <- soma
	}()

	// Goroutine para Débitos
	go func() {
		defer wg.Done()
		var subtracao float64 = 0
		for i := 0; i < len(transacoes); i++ {
			if transacoes[i].Tipo == "Débito" {
				subtracao += transacoes[i].Valor
			}
		}
		canalDebitos <- subtracao
	}()

	// Rotina principal aguarda o processamento e fecha os canais
	go func() {
		wg.Wait()
		close(canalCreditos)
		close(canalDebitos)
	}()

	// 3. Resultado Final: Coleta dos canais e cálculo do saldo
	totalCredito := <-canalCreditos
	totalDebito := <-canalDebitos
	saldoFinal := totalCredito - totalDebito

	// Exibição dos resultados
	fmt.Println("--- Relatório de Processamento Bancário ---")
	fmt.Printf("Total Processado em Créditos: R$ %.2f\n", totalCredito)
	fmt.Printf("Total Processado em Débitos:  R$ %.2f\n", totalDebito)
	fmt.Println("-------------------------------------------")
	if saldoFinal < 0 {
		fmt.Printf("Atenção: Conta negativada! Saldo: R$ %.2f\n", saldoFinal)
	} else {
		fmt.Printf("Saldo Final: R$ %.2f\n", saldoFinal)
	}
}