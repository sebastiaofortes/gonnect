package gonnect

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func isInterface(r reflect.Type) bool {
	return r.Kind() == reflect.Interface
}

func searchDisambiguation(returnType reflect.Type, dependenciesFound []DependencyBean) DependencyBean {

	fmt.Println("returnType: ", returnType)
	// Iterar sobre os campos da struct e ler os metadados
	numField := returnType.NumField()
	if numField == 0 {
		message := fmt.Sprintf("strct %v com mais de um contrutor e sem valores para desqualificar", returnType)
		panic(message)
	}
	for i := 0; i < numField; i++ {
		field := returnType.Field(i)
		fieldName := field.Name
		fmt.Println(fieldName)
		// obtem o metadado da tag
		tagValue := field.Tag.Get("sebas")
		fmt.Println("tagValue: ", tagValue)
		fmt.Println("valores de found :", dependenciesFound)
		for _, dependency := range dependenciesFound {
			fmt.Println("dentro de name: ", dependency.Name)
			nameParts := strings.Split(dependency.Name, ".")
			if nameParts[len(nameParts)-1] == tagValue {
				fmt.Println("Encontrado um METADADO COMPATIVEL")
				return dependency
			}
		}
	}
	panic("Mais de um construtor encontrado para um mesmo tipo, Nenhum METADADO encontrado para resolver a ambiguidade")
	return DependencyBean{}
}

// Função para verificar se uma struct implementa uma interface
func implementsInterface(structType reflect.Type, interfaceType reflect.Type) bool {
	return structType.Implements(interfaceType)
}

func generateDependenciesArray(funcs []interface{}, isGlobal bool) map[string]DependencyBean {
	ReflectTypeArray := make(map[string]DependencyBean)
	for _, fn := range funcs {
		fnType := reflect.TypeOf(fn)
		fnValue := reflect.ValueOf(fn)
		nameFunction := getFunctionName(fnValue)
		paramTypes := getParamTypes(fnType)
		returnType := getReturnType(fnType)
		ReflectTypeArray[nameFunction] = DependencyBean{constructorType: fnType, fnValue: fnValue, Name: nameFunction, IsGlobal: isGlobal, IsFunction: true, constructorReturn: returnType, ParamTypes: paramTypes}

	}
	return ReflectTypeArray
}

func getFunctionName(i reflect.Value) string {
	return runtime.FuncForPC(i.Pointer()).Name()
}

func getParamTypes(fnType reflect.Type) []reflect.Type {
	var paramTypes []reflect.Type
	for i := 0; i < fnType.NumIn(); i++ {
		paramTypes = append(paramTypes, fnType.In(i))
	}
	return paramTypes
}

func getReturnType(fnType reflect.Type) reflect.Type {
	if fnType.NumOut() == 1 {
		return fnType.Out(0)
	} else {
		message := fmt.Sprintf("Erro, a função %s deve possuir um único tipo de retrono \n", fnType.Name())
		panic(message)
	}
}
