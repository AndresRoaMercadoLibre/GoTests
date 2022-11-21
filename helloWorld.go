package main

import "errors"

func helloWorld(content string) (string, error) {
	if content != "hola" {
		return content, errors.New("Valor del argumento erroneo")
	}
	return "retorno", nil
}
