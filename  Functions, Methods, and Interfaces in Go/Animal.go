package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food string
	locomotion string
	noise      string
}
func (x *animal) Eat(){
	fmt.Println(x.food)
}

func (x *animal) Move(){
	fmt.Println(x.locomotion)
}

func (x *animal) Speak(){
	fmt.Println(x.noise)
}

func storeData(s []string){
	obj := new(animal)

	switch s[0]{
	case "cow":
		obj.food = "grass"
		obj.locomotion = "walk"
		obj.noise = "moo"

	case "bird":
		obj.food = "worms"
		obj.locomotion = "fly"
		obj.noise = "poop"

	case "snake":
		obj.food = "mice"
		obj.locomotion = "slither"
		obj.noise = "hsss"
	}

	switch s[1]{
	case "eat":
		obj.Eat()
	case "speak":
		obj.Speak()
	case "move":
		obj.Move()
	}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		a, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSpace(a), " ")

		storeData(s)
	}
}