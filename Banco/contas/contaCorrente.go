package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64){
	podeSacar := valor <= c.saldo && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso.", c.saldo
	}else{
		return "saldo insufuciente ", c.saldo
	}

}

func(c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
	c.saldo += valorDoDeposito 
	return "Deposito realizado com sucesso.", c.saldo
	}else {
		return "Deposito falhou.", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0{
		c.saldo -= valorDaTransferencia 
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}else{
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64{
	return c.saldo
}

