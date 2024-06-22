package main

import (
	"testing"

	"github.com/sebastiaofortes/gonnect"
	"github.com/stretchr/testify/assert"
)

func Test_interfaces_Bean_not_found(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{}
	a.AddDependencies(funcs)
	assert.Panics(t, func() { a.StartApp(NewMyDependencyObject) })
}

func Test_interfaces_Success(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{newMyImplementation}
	a.AddDependencies(funcs)
	assert.NotPanics(t, func() { a.StartApp(NewMyDependencyObject) })
}
