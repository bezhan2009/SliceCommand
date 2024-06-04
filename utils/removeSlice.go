package utils

import "fmt"

func RemoveSlice(x map[string][]int, key string) (map[string][]int, bool) {
	if x[key] != nil {
		delete(x, key)
		return x, true
	} else {
		fmt.Println("Слайс не найден!!!")
		return x, false
	}
}

func RemoveElement(slice []int, index int) ([]int, bool) {
	if index < 0 || index >= len(slice) {
		return slice, false
	}
	return append(slice[:index], slice[index+1:]...), true
}
