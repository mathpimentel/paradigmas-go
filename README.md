# Processador Concorrente de Transações Bancárias 🏦

Este projeto é um protótipo de um sistema financeiro desenvolvido em **Go (Golang)** para fins acadêmicos. O objetivo é demonstrar a aplicação do **Paradigma Imperativo** em conjunto com os mecanismos de **Concorrência** nativos da linguagem.

## 🚀 Sobre o Programa
O programa processa um lote fixo de 20 transações financeiras simultaneamente. Ele divide o trabalho em duas unidades de execução independentes:
- **Goroutine A:** Filtra e soma todos os Créditos.
- **Goroutine B:** Filtra e soma todos os Débitos.

As rotinas se comunicam com a função principal através de **Channels**, e a sincronização é garantida por um **WaitGroup** (Barreira de Sincronização).

## 🛠️ Tecnologias Utilizadas
- **Linguagem:** Go 1.20+
- **Paradigma:** Imperativo
- **Bibliotecas Standard:** `fmt`, `sync`

## 🏁 Como Rodar o Projeto

### 1. Instalação do Go
Caso não tenha o Go instalado:
1. Acesse [go.dev/dl](https://go.dev/dl/).
2. Baixe o instalador correspondente ao seu Sistema Operacional (Windows, macOS ou Linux).
3. Siga os passos de instalação padrão ("Next, Next, Finish").
4. No terminal, verifique a instalação com o comando:
   ```bash
   go version

### 2. Executando o Código
Com o Go instalado, navegue até a pasta do projeto via terminal e execute:
   ```bash
   go run main.go
