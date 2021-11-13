package main

import "fmt"

func main() {
	var mid, low, data = 0,0,[10]int{1,2,3,4,5,6,7,8,9,10}
	high := len(data) -1
	key := 8
	for low <= high {
		mid = (low + high)/2
		if data[mid] == key {
			fmt.Printf("located %d",mid)
			break
		}

		if data[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}