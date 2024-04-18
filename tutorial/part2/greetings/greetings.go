package greetings //Nome da pasta que você está (praticamente)

import (
	"errors"
	"fmt"
)

// Nome da função (Variavel Tipo) (Tipo de retorno)
func Hello(name string) (string, error) {

	//Valida se a variavel nome está vazia e retorna erro caso TRUE
	if name == "" {
		return "", errors.New("nome vazio")
	}

	//Mensagem que será retornada
	message := fmt.Sprintf("Olá, %v. Bem-vindo!", name)

	//nil = "no error"
	return message, nil
}

func Idade(idade int) string {
	message := fmt.Sprintf("Sua idade é: %v", idade)

	return message
}
