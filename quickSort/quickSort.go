package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 7, 9, 3, 2, 10, 5, 8, 6}
	fmt.Println("Quick Sort")
	fmt.Printf("ori number : %v \n", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Printf("sorted number : %v \n", arr)
}

func quicksort(arr []int, left int, right int) {
	pivot := arr[(left+right)/2]
	if left < right {
		i,j := left,right
		for {
			for arr[i] < pivot { i++ }
			for arr[j] > pivot { j-- }
			if i >= j { break }
			t := arr[i]
			arr[i] = arr[j]
			arr[j] = t
		}
		quicksort(arr, left, i-1)
		quicksort(arr, j+1, right)
	}
}
