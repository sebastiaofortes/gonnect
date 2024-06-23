package gonnect

import (
	"fmt"
	"reflect"
)

type Container struct {
	dependencies map[string]DependencyBean
}

func NewContainer() Container {
	return Container{}
}

func (c *Container) AddDependencies(deps []interface{}) {
	// Gera o array com as dependencias
	ReflectTypeArray := generateDependenciesArray(deps, false)
	c.checkingNameUnit(ReflectTypeArray)
	c.dependencies = ReflectTypeArray
}

func (c *Container) AddGlobalDependencies(deps []interface{}) {
	// Gera o array com as dependencias
	ReflectTypeArray := generateDependenciesArray(deps, true)
	c.checkingNameUnit(ReflectTypeArray)
	c.dependencies = ReflectTypeArray
}

func (f *Container) StartApp(startFunc interface{}) {

	fmt.Println("Inicianfo app.....")
	quantDep := len(f.dependencies)
	fmt.Println(quantDep, " dependencias registradas")

	fnType := reflect.TypeOf(startFunc)
	fnValue := reflect.ValueOf(startFunc)
	fnName := getFunctionName(fnValue)

	dep := DependencyBean{constructorType: fnType, fnValue: fnValue, Name: fnName}

	args := f.getDependencyConstructorArgs(dep)

	// Chamando o construtor e enviando os parametros encontrados
	dep.fnValue.Call(args)

	fmt.Println("............Iniciando aplicação................")
	fmt.Println()

	// Iterar sobre o array de funções para inspecionar os parâmetros

}

func (c *Container) getDependencyConstructorArgs(dependency DependencyBean) []reflect.Value {
	args := []reflect.Value{}
	var returnType reflect.Type
	fmt.Printf("Quantidade de parâmetros: %d\n", dependency.constructorType.NumIn())
	for i := 0; i < dependency.constructorType.NumIn(); i++ {
		paramType := dependency.constructorType.In(i)
		if dependency.constructorType.NumOut() == 1{
			returnType = dependency.constructorType.Out(0)
		} else{
			panic("as funções devem possuir apenas um tipo de retorno")
		}
		fmt.Printf("Parâmetro %d: %v\n", i, paramType)
		fmt.Println("Procurando funções com retorno ou que implementem do tipo:", paramType)
		// Procura na lista de um contrutuores um tipo igual ao do parametro

		injectableDependency := c.searchInjectableDependencies(paramType, returnType)

		if injectableDependency.IsFunction {
			argumants := c.getDependencyConstructorArgs(injectableDependency)
			resp := injectableDependency.fnValue.Call(argumants)
			fmt.Println("Adicionando contrutor ao slice de argumentos ***********")
			args = append(args, resp...)
			if injectableDependency.IsGlobal {
				// Change function dependency to object dependency
				injectableDependency.fnValue = resp[0]
				injectableDependency.IsFunction = false
				// Update the object in the dependencies list

				c.dependencies[injectableDependency.Name] = injectableDependency
			}
		} else {
			args = append(args, injectableDependency.fnValue)
		}
	}
	return args
}

func (c *Container) searchInjectableDependencies(paramType reflect.Type, returnType reflect.Type) DependencyBean {
	var dependenciesFound []DependencyBean
	var depFound DependencyBean
	if isInterface(paramType) {
		// Chama a função searchImplementations
		fmt.Println(paramType, " É uma interface ")
		dependenciesFound = c.searchImplementations(paramType)
	} else {
		fmt.Println("Não é uma interface")
		// Chama a função searchType
		dependenciesFound = c.searchTypes(paramType)
	}
	if len(dependenciesFound) > 1 {
		// O elemento 0 é o único já que os contrutores só tem um retorno
		disambiguation := searchDisambiguation(returnType, dependenciesFound)
		return disambiguation
	} else if len(dependenciesFound) == 0 {
		panic("nemhum construtor para o parametro foi encontrado")
	} else {
		depFound = dependenciesFound[0]
	}
	return depFound
}

func (f *Container) searchTypes(paramType reflect.Type) []DependencyBean {
	dependenciesFound := []DependencyBean{}
	for fnName, dependency := range f.dependencies {
		for i := 0; i < dependency.constructorType.NumOut(); i++ {
			returnType := dependency.constructorType.Out(i)
			//fmt.Printf("Retorno %d: %v\n", i, returnType)
			if returnType == paramType {
				fmt.Println("Encontrei a o retorno na função", fnName, " tipo ", returnType)
				dependenciesFound = append(dependenciesFound, dependency)
			}
		}
	}
	return dependenciesFound
}

func (f *Container) searchImplementations(paramType reflect.Type) []DependencyBean {
	dependenciesFound := []DependencyBean{}
	for fnName, dependency := range f.dependencies {
		for i := 0; i < dependency.constructorType.NumOut(); i++ {
			returnType := dependency.constructorType.Out(i)
			implements := implementsInterface(returnType, paramType)
			//fmt.Printf("Retorno %d: %v\n", i, returnType)
			if implements {
				fmt.Println("Encontrei a o retorno na função", fnName, " tipo ", returnType)
				dependenciesFound = append(dependenciesFound, dependency)
			}
		}
	}
	return dependenciesFound
}

func (c *Container) checkingNameUnit(reflectTypeArray map[string]DependencyBean) {
	for _, v := range reflectTypeArray {
		if _, exists := c.dependencies[v.Name]; exists {
			panic("Duplicate constructor registration")
		}
	}
}
