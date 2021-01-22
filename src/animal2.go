package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new
animal of one of the three types, or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the
types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals
and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound

cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program
should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or
a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the
new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new
animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the
name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or
“speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface
should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method
should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should
print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods
Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user
creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the
user issues a query command.
 */
func main() {
	reader := bufio.NewReader(os.Stdin)

	zoo := map[string]Animal {}

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error occurred while reading the user input %s \n", err)
		}

		args := strings.Fields(strings.TrimSpace(line))
		if len(args) != 3 {
			fmt.Println("You should enter 3 strings, query / newanimal command, animal name and information request or animal type")
			continue
		}

		OuterSwitch:
			switch args[0] {
			case "query":
				if animal, ok := zoo[args[1]]; ok {
					switch args[2] {
					case "eat":
						animal.Eat()
					case "move":
						animal.Move()
					case "speak":
						animal.Speak()
					default:
						fmt.Printf("Information for [%s] about [%s] is not present\n", args[1], args[2])
					}
				} else {
					fmt.Printf("Animal %s is no present in zoo\n", args[1])
				}
			case "newanimal":
				var newanimal Animal
				switch args[2] {
				case "cow":
					newanimal = Cow{}
				case "snake":
					newanimal = Snake{}
				case "bird":
					newanimal = Bird{}
				default:
					fmt.Printf("Uknown animal [%s] \n", args[2])
					break OuterSwitch
				}
				zoo[args[1]] = newanimal
				fmt.Println("Created it!")
			default:
				fmt.Printf("Bad command [%s]", args[0])
			}
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}

func (_ Cow) Eat() {
	fmt.Println("grass")
}

func (_ Cow) Move() {
	fmt.Println("walk")
}

func (_ Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {}

func (_ Bird) Eat() {
	fmt.Println("worms")
}

func (_ Bird) Move() {
	fmt.Println("fly")
}

func (_ Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {}

func (_ Snake) Eat() {
	fmt.Println("mice")
}

func (_ Snake) Move() {
	fmt.Println("slither")
}

func (_ Snake) Speak() {
	fmt.Println("hsss")
}