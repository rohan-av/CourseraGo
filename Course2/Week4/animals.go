package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	nameMap := map[string]Animal{}

	for {
		fmt.Print("> ")
		var query string
		var name string
		var typeInfo string
		fmt.Scan(&query, &name, &typeInfo)
		if query == "newanimal" {
			var newAnimal Animal
			switch typeInfo {
			case "cow":
				newAnimal = Cow{name}
				nameMap[name] = newAnimal
			case "bird":
				newAnimal = Bird{name}
				nameMap[name] = newAnimal
			case "snake":
				newAnimal = Snake{name}
				nameMap[name] = newAnimal
			}
			fmt.Println("Created it!")
		} else if query == "query" {
			a, p := nameMap[name]
			if !p {
				fmt.Println("No such animal found!")
			} else {
				switch typeInfo {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			}
		}

	}
}
