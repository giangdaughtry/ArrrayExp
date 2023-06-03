package main

import (
	"fmt"
)

func main() {
	arr := [...]int{3, 6, 24, 7, 5, 35, 43, 564, 34}
	for i := 1; i < len(arr); i++ {
		max := arr[i]
		j := i - 1
		// for j := i - 1; j >= 0; j-- {
		// 	if arr[j] > max {
		// 		arr[j+1] = arr[j]
		// 		arr[j] = max
		// 	}

		for (j >= 0) && (arr[j] > max) {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = max
	}

	fmt.Println(arr)
}
