package services

import (
	"alura/contas/models"
	"errors"
)

func DebitarConta(conta *models.ContaCorrente, valor float64) error {
	if valor <= 0 {
		return errors.New("valor deve ser positivo")
	}
	if conta.Saldo < valor {
		return errors.New("saldo insuficiente")
	}
	conta.Saldo -= valor
	return nil
}

func CreditarConta(conta *models.ContaCorrente, valor float64) {
	conta.Saldo += valor
}

func TransferirEntreContas(contaOrigem, contaDestino *models.ContaCorrente, valor float64) error {
	if err := DebitarConta(contaOrigem, valor); err != nil {
		return err
	}
	CreditarConta(contaDestino, valor)
	return nil
}
