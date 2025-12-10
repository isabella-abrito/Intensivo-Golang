package main

import (
	"fmt"
	"log"
	"net/http"
	"Aprendendo-Golang.com/math" //tem que importar o modulo externo para que a função seja reconhecida
) // para importar várias bibliotecas, usar essa estrutura no import

// var a string: um jeito de declarar uma variável com tipo explícito

func main() {
	a := "Isabella" // := declara e atribui valor à variável e o go "adivinha" qual é o tipo da variável
	// variável não pode mudar de tipo, se eu fizer a = 1 vai dar erro!

	b := 3.14
	c := false
	d := `Quero imprimir
	algo aqui
		quebrando a linha`

	fmt.Println(a) // Println printa e quebra a linha

	fmt.Printf("%v\n", b) // Printf é outro jeito de imprimir algo na tela, "%v": está passando o valor de b

	fmt.Printf("%v\n", c)
	fmt.Println(d)
	fmt.Printf("%T", d) // %T printa o tipo da variável

	fmt.Printf("\n\nValor da soma é: %v", math.Soma(5, 5)) 
	// todas as vezes que eu acessar um pacote externo, a Primeira Letra da estrutura deve ser Maiúscula
	// o math. busca a função no outro arquivo

	fmt.Println() //quebra de linha

	//tratamento de erros em golang: NÃO TEM TRY CATCH
	//simulando fazer uma requisicao qualquer:
	res, err := http.Get("http://google.com.br") //Get retorna uma resposta e um erro
	if err != nil {
		log.Fatal(err.Error()) //log com data e horario do erro, Error() vai trazer qual o tipo de erro aconteceu
	}
	fmt.Println(res.Header)

	// testando tratamento de erro na funcao SomaError quando o resultado é maior que 10
	res1, err1 := math.SomaError(7, 4)
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	fmt.Println(res1)

	// se eu não quiser tratar o erro na main:
	res2, _ := math.SomaError(7, 4) // o valor do erro vai para _ e é jogado fora
	fmt.Println(res2)
} 