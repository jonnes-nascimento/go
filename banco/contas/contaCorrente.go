package contas

import "banco/clientes"

// ContaCorrente e um tipo que pode ser exportado
type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

// Sacar e um metodo de ContaCorrente
func (c *ContaCorrente) Sacar(valorSaque float64) string {

	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque

		return "Saque realizado com sucesso!"
	}

	return "saldo insuficiente."
}

// Depositar e um metodo de ContaCorrente
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do deposito menor que zero.", c.saldo
}

// Transferir e um metodo de ContaCorrente
func (c *ContaCorrente) Transferir(contaDestino *ContaCorrente, valorTransferencia float64) bool {

	if valorTransferencia > 0 && c.saldo > valorTransferencia {
		c.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)

		return true
	}

	return false
}

// ObterSaldo devolve o valor do saldo da conta corrente que invocar o metodo
func (c ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
