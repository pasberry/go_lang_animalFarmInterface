//Write a program which allows the user to create a set of animals and to get information about those animals.
//Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new
//animal of one of the three types, or the user can request information about an animal that he/she has already created.
//Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type,
//but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of
//animals and their associated data.
//
//  Animal | Food eaten	| Locomotion method	| Spoken sound
// -----------------------------------------------------------------------
//  cow	   | grass	    | walk	            | moo
//  bird   | worms	    | fly	            | peep
//  snake  | mice	    | slither	        | hsss

//Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
//Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a
//new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal”
//command or a “query” command.

//Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
//The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the
//new animal, either “cow”, “bird”, or “snake”. Your program should process each newanimal command by creating the new
//animal and printing “Created it!” on the screen.

//Each “query” command must be a single line containing 3 strings. The first string is “query”.
//The second string is the name of the animal. The third string is the name of the information requested about the
//animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the
//requested data.

//Define an interface type called Animal which describes the methods of an animal.
//Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments
//and return no values. The Eat() method should print the animal’s food, the Move() method should print the
//animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
//Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak()
//so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal,
//create an object of the appropriate type. Your program should call the appropriate method when the user issues a
//query command.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {

	Eat()
	Move()
	Speak()
}

type Cow struct {
	food , locomotion, noise string
}

func (animal *Cow) Eat()  {
	fmt.Println(animal.food)
}

func (animal *Cow) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Cow) Speak() {
	fmt.Println(animal.noise)
}

type Bird struct {
	food , locomotion, noise string
}

func (animal *Bird) Eat()  {
	fmt.Println(animal.food)
}

func (animal *Bird) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Bird) Speak() {
	fmt.Println(animal.noise)
}

type Snake struct {
	food , locomotion, noise string
}

func (animal *Snake) Eat()  {
	fmt.Println(animal.food)
}

func (animal *Snake) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Snake) Speak() {
	fmt.Println(animal.noise)
}

var cow  = Cow{"grass", "walk", "moo"}
var bird  = Bird{"worms", "fly", "peep"}
var snake  = Snake{"mice", "slither", "hsss"}

var animalFarm = make(map[string]Animal)

func main() {

	fmt.Println("!!! Welcome to the animal program !!!")
	fmt.Println("Type exit to end the program. ")
	fmt.Print(">> ")

	for {
		
		inputScanner := setUpConsoleScanner()
		userCommands := getUserInput(inputScanner)
		processCommands(userCommands)
		fmt.Print(">> ")
	}

}

func processCommands(commands []string) {

	if commands[0] == "exit" {
		os.Exit(0)
	}

	switch commands[0] {

	case "newanimal":
		createNewAnimal(commands[1], commands[2])
	case "query":
		findNewAnimal(commands[1], commands[2])
	default :
		fmt.Println("Sorry , I do not recognize that command.")
	}
}

func findNewAnimal(name, animalAction string) {

	animal, ok  := animalFarm[name]

	if ok {
		processAnimalAction(animal , animalAction)
	}
	if ! ok {
		fmt.Println("Sorry, I can't find that animal.")
	}
}

func processAnimalAction(animal Animal, action string) {
	
	switch action {
	
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak" :
		animal.Speak()
	default:
		fmt.Println("Sorry this animal does not know this trick.")
	}
}

func createNewAnimal(name, animalType string) {

	switch animalType {

	case "cow" :
		animalFarm[name] = &cow
	case "bird" :
		animalFarm[name] = &bird
	case "snake" :
		animalFarm[name] = &snake
	default:
		fmt.Println("Sorry I don't know that animal.")
	}
}

func setUpConsoleScanner() *bufio.Scanner {

	inputScanner := bufio.NewScanner(os.Stdin)
	return inputScanner
}

func getUserInput(inputScanner *bufio.Scanner) []string {
	inputScanner.Scan()
	return strings.Split(strings.TrimSpace(inputScanner.Text()), " ")
}