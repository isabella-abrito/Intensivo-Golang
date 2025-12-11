package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string 
	Email string
	CPF   int
}

func (c Cliente) Exibe() { //metodo ligado à struct

	fmt.Println("Exibindo cliente pelo método: ", c.Nome)

}

type ClienteInternacional struct {
	//Nome string
	//Email string
	//CPF int <-- atributos apagados pq são iguais aos da struct Cliente
	Cliente        //pega os atributos da struct de cima
	Pais    string `json:"pais"` //tag: anotacao para quando rodar a funcao de JSON vai transformar Pais em pais
}

func main() {
	cliente := Cliente{
		Nome:  "Isabella",
		Email: "isabella@isabella.com",
		CPF:   12345678900,
	}

	cliente.Exibe() //chamando o metodo ligado à struct

	fmt.Println(cliente)

	cliente2 := Cliente{"Adriane", "adri@ane.com", 98765432100} //outra forma de preencher os dados da struct
	fmt.Println(cliente2)
	fmt.Printf("\nNome: %s, Email: %s, CPF: %d", cliente2.Nome, cliente2.Email, cliente2.CPF) // %s para string, %d para int

	cliente3 := ClienteInternacional{
		Cliente: Cliente{ //chama a primeira struct para colocar os atributos
			Nome:  "Zuri",
			Email: "zu@ri.com",
			CPF:   12457812220,
		},
		Pais: "Africa do Sul",
	}

	fmt.Printf("\nNome: %s, Email: %s, CPF: %d, Pais: %s", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais) // não precisa fazer cliente3.Cliente.Nome

	clienteJson, err := json.Marshal(cliente3) //Marshal pega dados em formato de bytes
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("\n\nImprimindo em formato JSON: ", string(clienteJson)) //pega os bytes e converte em string

	jsonCliente4 := `{"Nome":"Zuri","Email":"zu@ri.com","CPF":12457812220,"pais":"Africa do Sul"}`
	
	cliente4 := ClienteInternacional{} //em branco

	//quero pegar o json e colocar no cliente4

	json.Unmarshal([]byte(jsonCliente4), &cliente4) //unmarshal transforma de string para byte, &cliente4 altera o conteúdo no endereço de memória, 
	// se passasse apenas cliente4 iria alterar os dados apenas nesse escopo, mas a variavel cliente4 continuaria em branco

	fmt.Println(cliente4)
}
