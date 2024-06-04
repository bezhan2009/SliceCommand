package additions

import (
	"KarimovbejanHW/hw_tasks/utils"
	"fmt"
	"strconv"
)

func MergeSlices(x map[string][]int) map[string][]int {
	if len(x) < 2 {
		fmt.Println("\tВам нужно минимум два слайса для этого действия")
		return nil
	}

	utils.GetSlices(x)

	var (
		first       string
		second      string
		firstSlice  []int
		secondSlice []int
	)

	fmt.Print("\n\tВведите первый номер слайса: ")
	if _, err := fmt.Scan(&first); err != nil {
		fmt.Println("\tОшибка при считывании первого номера слайса:", err)
		return nil
	}

	fmt.Print("\n\n\tВведите второй номер слайса: ")
	if _, err := fmt.Scan(&second); err != nil {
		fmt.Println("\tОшибка при считывании второго номера слайса:", err)
		return nil
	}

	firstSlice = x["Слайс №"+first]
	secondSlice = x["Слайс №"+second]

	if firstSlice == nil || secondSlice == nil {
		fmt.Println("\tСлайс не найден!!!")
		return nil
	}

	key := "Слайс №" + strconv.Itoa(len(x)+1)
	merged := utils.MergeSlices(firstSlice, secondSlice)

	if len(merged) == 0 {
		fmt.Println("\tРезультат объединения пустой")
		return nil
	}

	x[key] = Sort(merged)
	fmt.Println("\tУспешно!!!")
	return x
}
