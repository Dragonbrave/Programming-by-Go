package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main(){
	fmt.Printf("Enter 12 integers value: ")
	arr := make([]int, 12)
	for i := 0; i < 12; i++{
		fmt.Scan(&arr[i])
	}
	const pair = 4
	subArray := int(math.Max(math.Ceil(float64(12)/float64(pair)), 1))
	var wg sync.WaitGroup

	for i := 0; i < pair; i++{
		start := int(math.Min(float64(i*subArray), float64(len(arr))))
		end := int(math.Min(float64((i+1)*subArray), float64(len(arr))))

		wg.Add(1)
		go func(arrA []int){
			fmt.Println("Now sorting:", arrA)
			sort.Ints(arrA)
			fmt.Println("sorted:", arrA)
			wg.Done()
		}(arr[start:end])
	}

	wg.Wait()
	sort.Ints(arr)
	fmt.Println("Sorted Array:", arr)
}
