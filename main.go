package main

import "fmt"

// var a string: um jeito de declarar uma variável com tipo explícito

func main() {
	a := "Isabella" // := declara e atribui valor à variável e o go "adivinha" qual é o tipo da variável
	//variável não pode mudar de tipo, se eu fizer a = 1 vai dar erro!

	b := 3.14
	c := false 
	d := `Quero imprimir
	algo aqui
		quebrando a linha`

	fmt.Println(a) //Println printa e quebra a linha
	
	fmt.Printf("%v\n", b) // Printf é outro jeito de imprimir algo na tela, "%v": está passando o valor de b
	
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", d) // %T printa o tipo da variável

}
