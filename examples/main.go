package main

import "github.com/sebastiaofortes/gonnect"

func main() {
	// Criação de um array de funções de diferentes tipos
	dependencies := []interface{}{}

	app := gonnect.NewContainer()

	app.AddDependencies(dependencies)

	app.StartApp(InitializeAPP)
}

func InitializeAPP() {

}
