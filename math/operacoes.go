package math

import "errors"

func Soma(a int, b int) int { // tem que explicitar o tipo de retorno
	return a + b
}

func SomaError(x int, y int) (int, error) { 
	res := x + y
	if res > 10 {
		return 0, errors.New("total maior que 10") // traz uma mensagem de qual foi o erro
	}

	return res, nil //se não deu erro, retorna resultado e nulo para erro
}

func SomaString (a int, b int) (string, int) {
	return "O resultado da soma é: ", a + b
}

// outro jeito de mostrar o retorno da funcao:
func SomaRetorno (a int, b int) (resultado int) {
	resultado = a + b
	return // não precisa passar nada nesse return pq no cabeçalho da funcao ja estou dizendo que vai retornar o resultado
}

func SomaTudo (x ...int) int { // colocar ... é como passar um array de inteiros 
	resultado := 0
	for _, v := range x { // _ é a key, mas não precisa ser explicita; v é o valor q ta entrando
		resultado += v 
	}
	return resultado
}