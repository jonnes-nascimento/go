package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

// PagarBoleto implementa a funcionalida de pagamento de boletos independente do tipo de conta
func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {

	clienteJonnes := clientes.Titular{Nome: "Jonnes", Profissao: "Engenheiro", Cpf: "123.122.432-7"}
	clienteMi := clientes.Titular{Nome: "Io Mi", Profissao: "Gerente", Cpf: "231.145.631-9"}

	contaJonnes := contas.ContaCorrente{}
	contaJonnes.Titular = clienteJonnes
	contaJonnes.NumeroAgencia = 9673
	contaJonnes.NumeroConta = 16706
	contaJonnes.Depositar(1000)

	contaMi := contas.ContaPoupanca{Titular: clienteMi, NumeroAgencia: 123, NumeroConta: 43267, Operacao: 77}
	contaMi.Depositar(3000)

	PagarBoleto(&contaJonnes, 730)
	PagarBoleto(&contaMi, 200)

	fmt.Println(contaJonnes.ObterSaldo())
	fmt.Println(contaMi.ObterSaldo())
}
