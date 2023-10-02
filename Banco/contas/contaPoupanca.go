package contas

import(
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valor float64) (string, float64){
	podeSacar := valor <= c.saldo && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso.", c.saldo
	}else{
		return "saldo insufuciente ", c.saldo
	}

}

func(c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
	c.saldo += valorDoDeposito 
	return "Deposito realizado com sucesso.", c.saldo
	}else {
		return "Deposito falhou.", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64{
	return c.saldo
}