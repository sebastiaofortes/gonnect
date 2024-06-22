package main

import (
	"testing"

	"github.com/sebastiaofortes/gonnect"
	"github.com/stretchr/testify/assert"
)

func Test_Duplicate_constructor(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{beanInt, beanFloat32}
	assert.Panics(t, func() { 
		a.AddDependencies(funcs)
		a.AddGlobalDependencies(funcs)
		a.StartApp(InitializeAPP) 
	})
}
