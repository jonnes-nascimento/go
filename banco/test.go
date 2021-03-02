package main

import (
	"fmt"
)

func test() {

	fmt.Println("Teste da funcao variadica somando (1,2,3,4,5):", variadicFunctionSum(1, 2, 3, 4, 5))
}

func variadicFunctionSum(numeros ...int) int {

	var resultadoSoma int

	for _, num := range numeros {
		resultadoSoma += num
	}

	return resultadoSoma
}
