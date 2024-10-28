# Sistema de Gerenciamento de Contas

Este projeto é um sistema de gerenciamento de contas bancárias implementado em Go. Ele fornece uma API RESTful para criar, listar, debitar, creditar e transferir valores entre contas.

## Estrutura do Projeto

A estrutura do projeto é a seguinte:

/ │ ├── main.go # Ponto de entrada da aplicação ├── go.mod # Gerenciador de módulos do Go ├── contas/ # Módulo de contas │ ├── controllers/ # Controladores da API │ │ └── controller.go # Lógica dos controladores │ ├── models/ # Modelos de dados │ │ └── models.go # Definições de estruturas de dados │ ├── services/ # Lógica de negócios │ │ └── service.go # Funções de serviço │ └── routes/ # Rotas da API │ └── routes.go # Definição das rotas


## Funcionalidades

- **Criar Conta**: Permite criar uma nova conta com um titular, número da agência, número da conta e saldo inicial.
- **Listar Contas**: Exibe todas as contas existentes.
- **Debitar**: Permite debitar um valor de uma conta existente.
- **Creditar**: Permite creditar um valor em uma conta existente.
- **Transferir**: Permite transferir um valor entre duas contas.

## Tecnologias Utilizadas

- Go (Golang)
- Gorilla Mux (para gerenciamento de rotas)

