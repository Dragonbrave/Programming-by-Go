package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string: ")
	word,_ := reader.ReadString('\n')

	word = strings.ToLower(word)
	b := strings.ContainsAny(word, "a")//contains函数是会返回一个新的字符串地
	if word[0]=='i' && word[len(word)-3]=='n'&& b {
		fmt.Printf("Found!")
	} else{
		fmt.Printf("Not Found!")
	}

}
