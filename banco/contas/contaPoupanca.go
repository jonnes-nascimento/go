package contas

import "banco/clientes"

// ContaPoupanca e um tipo que pode ser exportado
type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

// Sacar e um metodo de ContaCorrente
func (c *ContaPoupanca) Sacar(valorSaque float64) string {

	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque

		return "Saque realizado com sucesso!"
	}

	return "saldo insuficiente."
}

// Depositar e um metodo de ContaCorrente
func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do deposito menor que zero.", c.saldo
}

// ObterSaldo devolve o valor do saldo da conta corrente que invocar o metodo
func (c ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
