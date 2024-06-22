package main

import (
	"testing"

	"github.com/sebastiaofortes/gonnect"
	"github.com/stretchr/testify/assert"
)

func Test_metadata_Bean_not_found(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{}
	a.AddDependencies(funcs)
	assert.Panics(t, func() { a.StartApp(NewBeanWithMetadata) })
}

func Test_metadata_Success(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{NewBeanDependency1}
	a.AddDependencies(funcs)
	assert.NotPanics(t, func() { a.StartApp(NewBeanWithMetadata) })
}

func Test_Success_Disambiguation(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{NewBeanDependency1, NewBeanDependency2}
	a.AddDependencies(funcs)
	assert.NotPanics(t, func() { a.StartApp(NewBeanWithMetadata) })
}

func Test_Disambiguation_Tag_not_found(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{NewBeanDependency1, NewBeanDependency3}
	a.AddDependencies(funcs)
	assert.Panics(t, func() { a.StartApp(NewBeanWithMetadata) })
}

func Test_Disambiguation_Not_Tag(t *testing.T) {
	a := gonnect.NewContainer()
	// Criação de um array de funções de diferentes tipos
	funcs := []interface{}{NewBeanDependency1, NewBeanDependency2, NewBeanDependency3}
	a.AddDependencies(funcs)
	assert.Panics(t, func() { a.StartApp(NewBeanWithoutMetadata) })
}
