package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludo a la persona especificada.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}

	//Devuelve un saludo que incluye el nomber en un mensaje.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v¡ !Bienvenido!",
		"¡Que bueno verte, %v¡",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}

// Devielve un mapa de saludos asociando un nombre a un saludo.
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
