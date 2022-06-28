package main

import "fmt"

func main() {
	//swap, check order, one-round bubbling, repeated n times
	numbers := []int{6, 1, 4, 2, 9}

	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				swap(numbers, i, i+1)

			}
		}
	}

	fmt.Println(numbers)

}

func swap(number []int, i, j int) {

	temp := number[i]
	number[i] = number[j]
	number[j] = temp
}
