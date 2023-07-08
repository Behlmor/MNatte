package main

import "fmt"

// Definición de la interfaz
type Animal interface {
	Sound() string
}

// Implementación de la interfaz para el tipo Dog
type Dog struct{}

func (d Dog) Sound() string {
	return "Woof!"
}

// Implementación de la interfaz para el tipo Cat
type Cat struct{}

func (c Cat) Sound() string {
	return "Meow!"
}

func main() {
	// Crear una instancia de Dog y Cat
	dog := Dog{}
	cat := Cat{}

	// Usar las instancias como objetos de tipo Animal
	animals := []Animal{dog, cat}

	// Iterar sobre los animales y llamar al método Sound
	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}
