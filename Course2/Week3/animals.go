package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	animalMap := map[string]Animal{"cow": cow, "bird": bird, "snake": snake}

	for {
		fmt.Print("> ")
		var animal string
		var query string
		fmt.Scanln(&animal)
		fmt.Scanln(&query)
		a := animalMap[animal]
		if query == "at" {
			a.Eat()
		} else if query == "ove" {
			a.Move()
		} else if query == "peak" {
			a.Speak()
		}
	}
}
