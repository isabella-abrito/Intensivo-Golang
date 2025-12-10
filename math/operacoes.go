package math

import "errors"

func Soma(a int, b int) int {
	return a + b
}

func SomaError(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("total maior que 10") // traz uma mensagem de qual foi o erro
	}

	return res, nil //se não deu erro, retorna resultado e nulo para erro
}