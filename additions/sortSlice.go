package additions

import (
	"KarimovbejanHW/hw_tasks/utils"
	"fmt"
)

func SortSlice(x []int) []int {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       Как нам сортировать?       ║")
	fmt.Println("╠══════════════════════════════════╣")
	fmt.Println("║ 1. Сортировать по возрастанию    ║")
	fmt.Println("║ 2. Сортировать по убыванию       ║")
	fmt.Println("║ 0. Отмена                        ║")
	fmt.Println("╚══════════════════════════════════╝")

	var actionForSort int
	var isSort bool = false
	fmt.Print("\n\tВыбрать действия: \n\t\t\t")
	fmt.Scan(&actionForSort)

	if actionForSort == 0 {
		utils.ExitText()
		return x
	}

	switch actionForSort {
	case 1:
		isSort = true
		fmt.Println("Сортируем слайс по возрастанию")
		utils.AnimateLoading()
		fmt.Println("\t Отсортировано!!!")
	case 2:
		isSort = false
		fmt.Println("Сортируем слайс по убыванию")
		utils.AnimateLoading()
		fmt.Println("\t Отсортировано!!!")
	default:
		fmt.Println("Не опознанное действие!!!")
		return x
	}

	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] && isSort == true {
				x[i], x[j] = x[j], x[i]
			} else if x[i] < x[j] && isSort != true {
				x[i], x[j] = x[j], x[i]
			}
		}
	}

	return x
}

func Sort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}
