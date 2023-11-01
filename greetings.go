package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello, devuelve un saludo para una persona en especifico.

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	//message := fmt.Sprintf("Hola %v, Bienvenid@", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

// Funcion que recive varios nombres y devuelve para cada uno, un saludo.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message

	}
	return messages, nil

}

// randomFormat devuelve uno de varios mensajes de saludos aleatoriamente.
func randomFormat() string {
	formats := []string{
		"Hola %v, Bienvenid@",
		"Que gusto verte %v",
		"Saludos, %v",
	}
	return formats[rand.Intn(len(formats))]

}
