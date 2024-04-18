package greetings //Nome da pasta que você está (praticamente)

import "fmt"

// Nome da função (Variavel Tipo) Tipo de retorno
func Hello(name string) string {

	//Mensagem que será retornada
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
