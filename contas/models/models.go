package models

type ContaCorrente struct {
	ID            int
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}
