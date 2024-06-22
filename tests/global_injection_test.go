package main

import (
	"testing"

	"github.com/sebastiaofortes/gonnect"
	"github.com/stretchr/testify/assert"
)

func Test_Global_Injection_Bean_not_found(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{GlobalBeanString, globalBeanInt}
	a.AddGlobalDependencies(funcs)
	assert.Panics(t, func() { a.StartApp(GlobalInitializeAPP) })
}

func Test_Global_Injection_Success(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{globalBeanFloat32, GlobalBeanString, globalBeanInt}
	a.AddGlobalDependencies(funcs)
	assert.NotPanics(t, func() { a.StartApp(GlobalInitializeAPP) })
}
