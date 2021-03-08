package main

import "fmt"

// Global "database"
var data = make(map[string]Animal)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define Types for animals
type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Define methods
func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

// Creates new record in data map
func AddAnimal(name, animalType string) {
	switch animalType {
	case "cow":
		data[name] = Cow{"grass", "walk", "moo"}
	case "bird":
		data[name] = Bird{"worms", "fly", "peep"}
	case "snake":
		data[name] = Snake{"mice", "slither", "hsss"}
	}
	fmt.Println("Created it!")
}

// Query "database"
func QueryAnimal(name, info string) {
	switch info {
	case "eat":
		data[name].Eat()
	case "move":
		data[name].Move()
	case "speak":
		data[name].Speak()
	}
}

func main() {
	var command, name, info string

	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s %s", &command, &name, &info)
		switch command {
		case "newanimal":
			AddAnimal(name, info)
		case "query":
			QueryAnimal(name, info)
		}
	}
}
