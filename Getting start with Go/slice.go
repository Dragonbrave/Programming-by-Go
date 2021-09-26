package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	sli :=make([]int,3,4)
	w :=4
	for{
		fmt.Printf("Enter int value (X - quit): ")
		var n string
		fmt.Scan(&n)


		if n == "X" {
			break
		}else {
			t,_ := strconv.Atoi(n)//str in to integer
			sli = append(sli, t)
			w++
		}
		sort.Ints(sli[3:w])//为切片后面的内容排序

	    fmt.Printf("Contents of the slice are: ")
	    for i := 3; i < len(sli); i++{
		    fmt.Printf("%d ", sli[i])
	    }
	    fmt.Println()
	}

}
