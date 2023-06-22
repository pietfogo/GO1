package main

import "fmt"

func password2(senha string) {
	caractere := '*'
	quantidade := len(senha)
	for i := 0; i < quantidade; i++ {
		fmt.Print(string(caractere))
	}
	fmt.Println()
}

func main() {
	var estabelecimento = "Pietro's Pizzas"
	fmt.Printf("Aqui é do estabelecimento %v, por favor insira seu primeiro nome: ", estabelecimento)
	var firstName string
	fmt.Scan(&firstName)
	fmt.Printf("Prazer %v, se deseja consumir por favor preencha alguns dados listados abaixo:\n", firstName)
	var lastName string
	var email string
	var password string
	var leftRequestNumber int = 3
	var requestNumberClient int
	var decrescePedidos int
	fmt.Print("Seu sobrenome: ")
	fmt.Scan(&lastName)
	fmt.Print("Seu email: ")
	fmt.Scan(&email)
	fmt.Print("Sua senha: ")
	fmt.Scan(&password)
	fmt.Printf("Os dados foram inseridos corretamente?\nNome: %v\nSobrenome: %v\nEmail: %v\nSenha: ", firstName, lastName, email)
	password2(password)
	fmt.Printf("Você tem direito a %v pedidos, deseja fazer quantos pedidos? ", leftRequestNumber)
	fmt.Scan(&requestNumberClient)
	decrescePedidos = leftRequestNumber - requestNumberClient
	fmt.Printf("Você acaba de gastar %v pedidos e ainda lhe restam %v pedidos", requestNumberClient, decrescePedidos)
}
