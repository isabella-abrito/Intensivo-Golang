package main

import (
	"fmt"
	"log"
	"net/http"
	"Aprendendo-Golang.com/math" //tem que importar o modulo externo para que a função seja reconhecida
) // para importar várias bibliotecas, usar essa estrutura no import

// var a string: um jeito de declarar uma variável com tipo explícito

// CRIANDO UMA STRUCT
type Carro struct {
	Nome string //atributos aqui
}

func (c Carro) andar() { // variavel c esta fazendo referencia à estrutura.  andar() é um método!
	fmt.Println(c.Nome, "andou")
}

func pont(a *int) int {
	return *a //retorna o que tá guardado dentro daquele endereço de memória
}

func main() {

	//testando funcao anonima q acessar a struct
	carro := Carro{
		Nome: "BMW",
	} //atribuindo nome ao meu carro
	//chamo o metodo andar atrelado ao carro que acabei de nomear
	carro.andar()

	//ENTENDENDO PONTEIROS!!
	A := 10
	fmt.Println(&A) // & indica o endereço de memoria da variavel

	var ponteiro *int = &A //guardando o endereço de memoria de A dentro da variavel ponteiro q é um ponteiro inteiro
	fmt.Println(ponteiro)  // printando desse jeito, vai trazer o endereço de memória armazenado na variável ponteiro
	fmt.Println(*ponteiro) // printando desse jeito, vai pegar o endereço de memória que está apontando e ver qual valor está armazenado lá

	variavel := 100000000
	fmt.Println(pont(&variavel)) //pega o endereço da variavel, mas como o retorno da funcao eh o que esta guardado dentro dela, o retorno é o valor guardado lá
	


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

	// se eu não quiser tratar o erro na main:
	res2, _ := math.SomaError(7, 4) // o valor do erro vai para _ e é jogado fora
	fmt.Println(res2)

	// testando o retorno da funcao: string e inteiro
	string, resultado := math.SomaString(20, 10)
	fmt.Println(string, resultado)
	// ou posso apenas printar chamando a funcao dentro do argumento:
	fmt.Println(math.SomaString(10, 20))

	// testando funcao SomaTudo, passando quantos valores eu quiser
	fmt.Println(math.SomaTudo(3, 5, 4, 6, 7, 8, 280))

	// trabalhando com funcao anonima
	//a funcao retorna uma funcao - ACHEI MT COMPLICADO ESSA PARTE!
	result := func(x ...int) func() int {
		resultado := 0
		for _, v := range x {
			resultado += v
		}
		return func() int {
			return resultado
		}
	}

	// como a funcao anonimia retorna outra funcao, tem q colocar o () dps de passar os argumentos
	fmt.Println(result(45, 5, 5)())

	//----------
	// testando tratamento de erro na funcao SomaError quando o resultado é maior que 10
	res1, err1 := math.SomaError(7, 4)
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	fmt.Println(res1)

}
