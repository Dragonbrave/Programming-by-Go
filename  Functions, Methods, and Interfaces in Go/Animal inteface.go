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
type Cow struct{
	//{food : "grass", locomotion : "walk", spoken : "moo"}
}
func (c *Cow)Eat() {
	fmt.Printf("grass")
}
func (c *Cow)Move() {
	fmt.Printf("walk")
}
func (c *Cow)Speak() {
	fmt.Printf("moo")
}
type Bird struct{
	//{food : "worms", locomotion : "fly", spoken : "poop"}
}
func (b *Bird)Eat() {
	fmt.Printf("worms")
}
func (b *Bird)Move() {
	fmt.Printf("fly")
}
func (b *Bird)Speak() {
	fmt.Printf("peep")
}
type Snake struct{
	//{food : "mice", locomotion : "slither", spoken : "hsss"}
}
func (s *Snake)Eat() {
	fmt.Printf("mice")
}
func (s *Snake)Move() {
	fmt.Printf("slither")
}
func (s *Snake)Speak() {
	fmt.Printf("hsss")
}
func main(){
	reader := bufio.NewReader(os.Stdin)

	animals := make(map[string]Animal)

	for{
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSpace(input), " ")

		switch s[0]{
		case "newanimal":
			switch s[2]{
			case "cow":
				animals[s[1]] = new(Cow)
			case "bird":
				animals[s[1]] = new(Bird)
			case "snake":
				animals[s[1]] = new(Snake)
			}
			fmt.Println("Created it!")
		case "query":
			obj, ok := animals[s[1]]

			if ok{
				switch s[2]{
				case "eat":
					obj.Eat()
				case "move":
					obj.Move()
				case "speak":
					obj.Speak()
				}
			}else{
				fmt.Println("Not found!")
			}
		}
	}
}