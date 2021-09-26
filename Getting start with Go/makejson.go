package main
import (
	"encoding/json"
	"fmt"
	"log"
)
func main(){
	var name,address string
	fmt.Printf("please enter name:")
	fmt.Scan(&name)
	fmt.Printf("please enter address:")
	fmt.Scan(&address)
	person :=make(map[string]string)
	person[name]=address
	data,err:=json.Marshal(person)

	if err != nil {
		log.Println("ERROR:", err)
	}


	fmt.Println(string(data))

}
