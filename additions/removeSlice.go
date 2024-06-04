package additions

import (
	"KarimovbejanHW/hw_tasks/utils"
	"fmt"
	"strconv"
)

func RemoveSlice(x map[string][]int) (map[string][]int, bool) {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 Подробнее                  ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║               Удалить слайс                ║")
	fmt.Println("║              Для Отмены: 101               ║")
	fmt.Println("║           Удалить один эл.: 808            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	utils.GetSlices(x)
	fmt.Printf("\n\t\tВведите номер слайса для его удаления: ")
	var sliceRemove int
	fmt.Scan(&sliceRemove)

	if sliceRemove == 101 {
		utils.ExitText()
		return nil, false
	} else if sliceRemove == 808 {
		utils.GetSlices(x)
		fmt.Printf("\n\tВведите номер слайса: ")
		var sliceRemoveEl int
		fmt.Scan(&sliceRemoveEl)
		keySliceRemoveEl := "Слайс №" + strconv.Itoa(sliceRemoveEl)
		sliceForDeletionElement := x[keySliceRemoveEl]
		slice := RemoveElement(sliceForDeletionElement)
		x[keySliceRemoveEl] = slice
	} else {
		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║                 Подробнее                  ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║            Потвердите действие             ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║               Принять: 606                 ║")
		fmt.Println("║               Отменить: 101                ║")
		fmt.Println("╚════════════════════════════════════════════╝")

		var action int
		fmt.Print("\n\tВыбрать действия: \n\t\t\t")
		fmt.Scan(&action)
		if action == 606 {
			key := "Слайс №" + strconv.Itoa(sliceRemove)

			var isDid bool

			x, isDid = utils.RemoveSlice(x, key)

			if isDid {
				fmt.Println("\tУспешно удален!!!")
				return x, true
			} else {
				return x, false
			}
		} else if action == 101 {
			utils.ExitText()
			return x, false
		} else {
			fmt.Println("\tНе понятное действие!!!")
		}
	}

	return x, false
}

func RemoveElement(slice []int) []int {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 Подробнее                  ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║          Удалить элемент слайса            ║")
	fmt.Println("║              Для Отмены: 101               ║")
	fmt.Println("╚════════════════════════════════════════════╝")

	fmt.Printf("\n\tВведите индекс элемента в слайсе %d для его удаления: ", slice)
	var elementRemove int
	fmt.Scan(&elementRemove)

	if elementRemove == 101 {
		utils.ExitText()
		return slice
	} else if elementRemove > len(slice) {
		fmt.Println(elementRemove)
		fmt.Println(len(slice))
		fmt.Println("\tЭлемент не найден!!!")
	} else {
		var isDelete bool
		slice, isDelete = utils.RemoveElement(slice, elementRemove)
		if isDelete {
			fmt.Println("\tУспешно удален!!!")
			return slice
		} else {
			fmt.Println("\tЭлемент не был найден!!!")
			return slice
		}
	}

	return slice
}
