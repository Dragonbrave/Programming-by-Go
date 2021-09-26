package main


import "fmt"

func BubbleSort(arr[]int){
	for i:=0;i<len(arr);i++{
		for j := 0; j < len(arr) - i - 1; j++{
			if arr[j] > arr[j + 1]{
				Swap(&arr[j],&arr[j+1])
			}
		}
	}
}
func Swap(x *int ,y *int){
	temp:=0
	temp=*x
	*x=*y
	*y=temp
}
func main(){
	fmt.Printf("Enter 10 integers value: ")
	arr := make([]int, 10)
	for i := 0; i < 10; i++{
		fmt.Scan(&arr[i])
	}
	BubbleSort(arr)
	fmt.Println(arr)
}
