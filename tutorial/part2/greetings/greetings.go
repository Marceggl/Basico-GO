package greetings //Nome da pasta que você está (praticamente)

import (
	"errors"
	"fmt"
	"math/rand"
)

// Nome da função (Variavel Tipo) (Tipo de retorno)
func Hello(name string) (string, error) {

	//Valida se a variavel nome está vazia e retorna erro caso TRUE
	if name == "" {
		return "", errors.New("nome vazio")
	}

	//Criar uma mensagem usando random format
	message := fmt.Sprintf(randomFormat(), name)

	//nil = "no error"
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Olá, %v. Bem-Vindo!",
		"fala ai, bom te ver %v!",
		"Que a força esteja com você, %v!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func Idade(idade int) string {
	message := fmt.Sprintf("Sua idade é: %v", idade)

	return message
}
