package main

import (
	"errors"
	"fmt"
)

func validarNum(num int) (string, error) {
	if num == 0 {
		return "", errors.New("nome vazio")
	}

	mensagem := fmt.Sprintf("Você digitou: %v", num)

	return mensagem, nil
}

func main() {
	var numero int

	fmt.Println("Digite um número: ")
	_, err := fmt.Scan(&numero)
	if err != nil {
		fmt.Println("Erro ao ler o número:", err)
		return
	}

	msg, err := validarNum(numero)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(msg)
}
