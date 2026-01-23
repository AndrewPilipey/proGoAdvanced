package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := []string{"Apple", "Banana", "apple"}
	findIndexesOfProducts("apple", list, false)
	findIndexesOfProducts("apple", list, true)
}

func findIndexesOfProducts(value string, list []string, ignoreRegister bool) {
	foundItemsIgnore := []int{}
	foundItemsStrict := []int{}

	if ignoreRegister {
		for i := range list {
			if strings.ToLower(list[i]) == strings.ToLower(value) {
				foundItemsIgnore = append(foundItemsIgnore, i)
			}
		}
		var strFoundItemsIgnore []string
		for _, num := range foundItemsIgnore {
			strFoundItemsIgnore = append(strFoundItemsIgnore, strconv.Itoa(num))
		}
		resultIgnore := strings.Join(strFoundItemsIgnore, ", ")
		fmt.Println(resultIgnore)
	} else {
		for i := range list {
			if list[i] == value {
				foundItemsStrict = append(foundItemsStrict, i)
			}
		}
		var strFoundItemsStrict []string
		for _, num := range foundItemsStrict {
			strFoundItemsStrict = append(strFoundItemsStrict, strconv.Itoa(num))
		}
		resultStrict := strings.Join(strFoundItemsStrict, ", ")
		fmt.Println(resultStrict)
	}
}
