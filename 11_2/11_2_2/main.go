package main

import (
	"fmt"
)

type notFoundError struct {
	target int
	slice  []int
}

func (nf notFoundError) Error() string {
	return fmt.Sprintf("число %d не найдено в срезе %v", nf.target, nf.slice)
}

func FindInSlice(slice []int, target int) (int, error) {
	for i, value := range slice {
		if value == target {
			return i, nil
		}

	}
	return -1, notFoundError{target: target, slice: slice}
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	targets := []int{30, 100, 20}

	for _, target := range targets {
		idx, err := FindInSlice(numbers, target)
		if err != nil {
			fmt.Println(err)
			if nfErr, ok := err.(notFoundError); ok {
				fmt.Printf(" Искали %d в срезе %v\n", nfErr.target, nfErr.slice)
			}
		} else {
			fmt.Printf("Найдено %d на позиции %d\n", target, idx)
		}
	}
}
