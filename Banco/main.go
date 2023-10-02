package main

import (
	"banco/contas"
	"fmt"
	//"banco/clientes"
)


func main(){
	contaDoPec := contas.ContaCorrente{}
	contaDoPec.Depositar(300)
	PagarBoleto(&contaDoPec, 30)
	fmt.Println(contaDoPec.ObterSaldo())

	contaDoLeo := contas.ContaPoupanca{}
	contaDoLeo.Depositar(500)
	PagarBoleto(&contaDoLeo, 50)
	fmt.Println(contaDoLeo.ObterSaldo())



	// clienteBruno := clientes.Titular{"Bruno", "123.456.789-12", "Desenvolvedor Go"}
	// contaDoBruno := c.ContaCorrente{clienteBruno, 123, 123456, 100}
	// fmt.Println(contaDoBruno)
	// contaDoBruno.Depositar(100)


	/*contaDoPablo := c.ContaCorrente{}
	contaDoPablo.Titular = "Pablo"
	contaDoPablo.saldo = 500
	fmt.Println(contaDoPablo)

	fmt.Println(contaDoPablo.Sacar(200))
	fmt.Println(contaDoPablo.Depositar(1000))

	contaDoPayet := c.ContaCorrente{Titular: "Payet", saldo:2000}
	fmt.Println(contaDoPayet)

	status := contaDoPayet.Transferir(100, &contaDoPablo)
	fmt.Println(status)
	fmt.Println(contaDoPayet)
	fmt.Println(contaDoPablo)


	contaDoRossi := ContaCorrente{"Rossi", 310, 123457, 300}
	fmt.Println(contaDoRossi)
	

	Usando um ponteiro para criar uma conta
	var contaDoLeo *ContaCorrente
	contaDoLeo = new(ContaCorrente)
	contaDoLeo.Titular = "Leo"
	contaDoLeo.saldo = 500
	fmt.Println(*contaDoLeo)*/


}

func PagarBoleto(contas verificarConta, valorDoBoleto float64) {
	contas.Sacar(valorDoBoleto)
}  

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}