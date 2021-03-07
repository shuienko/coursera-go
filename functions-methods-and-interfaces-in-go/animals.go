package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var name, info string
	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s", &name, &info)

		switch name {
		case "cow":
			{
				if info == "eat" {
					cow.Eat()
				}
				if info == "move" {
					cow.Move()
				}
				if info == "speak" {
					cow.Speak()
				}
			}
		case "bird":
			{
				if info == "eat" {
					bird.Eat()
				}
				if info == "move" {
					bird.Move()
				}
				if info == "speak" {
					bird.Speak()
				}
			}
		case "snake":
			{
				if info == "eat" {
					snake.Eat()
				}
				if info == "move" {
					snake.Move()
				}
				if info == "speak" {
					snake.Speak()
				}
			}
		}
	}
}
