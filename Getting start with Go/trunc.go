package main
import "fmt"

func main() {
	var x float32=0.0
	fmt.Printf("please enter a float number:")
	fmt.Scan(&x)
	y:= int(x)
	fmt.Printf("the number is %d\n",y)
}