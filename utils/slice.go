package utils

import (
	"fmt"
)

func modifySlice(x []int) []int {
	fmt.Println(x)
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 Подробнее                  ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║  Ввести число для слайса с конкретным id   ║")
	fmt.Println("║              Для Отмены: 101               ║")
	fmt.Println("║     Для добавления нового значения: 303    ║")
	fmt.Println("╚════════════════════════════════════════════╝")

	var actionAddValue int

	fmt.Print("\nВведите какой индекс вы бы хотели изменить: \n\t\t\t")
	fmt.Scan(&actionAddValue)

	if actionAddValue == 101 {
		AnimateLoading()
		ExitText()
		return x
	}

	if actionAddValue == 303 {
		AnimateLoading()
		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║                 Подробнее                  ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║          Ввести число для слайса           ║")
		fmt.Println("║              Для Отмены: 101               ║")
		fmt.Println("╚════════════════════════════════════════════╝")

		var newObject int
		fmt.Print("\nВведите новое значения для слайса: \n\t\t\t")
		fmt.Scan(&newObject)

		x = append(x, newObject)
		return x
	}

	if actionAddValue >= len(x) {
		fmt.Println("Не найден пожалуйста введите правильные данные!!!")
		return x
	}

	for i := 0; i < len(x); i++ {
		if actionAddValue == i {
			fmt.Println("╔════════════════════════════════════════════╗")
			fmt.Println("║                 Подробнее                  ║")
			fmt.Println("╠════════════════════════════════════════════╣")
			fmt.Println("║          Ввести число для слайса           ║")
			fmt.Println("║              Для Отмены: 101               ║")
			fmt.Println("╚════════════════════════════════════════════╝")

			var newValue int

			fmt.Print("\nВведите новое значения для элемента под индексом: \n\t\t\t")
			fmt.Scan(&newValue)

			x[i] = newValue

			if newValue == 101 {
				AnimateLoading()
				ExitText()
				return x
			}
		}
	}
	return x
}

func GetIdSlice(x map[string][]int) map[string][]int {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 Подробнее                  ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║          Ввести число для слайса           ║")
	fmt.Println("║              Для Отмены: 101               ║")
	fmt.Println("╚════════════════════════════════════════════╝")

	GetSlices(x)

	var sliceId string
	fmt.Print("\nнапишите номер слайса который хотите изменить: ")
	fmt.Scan(&sliceId)

	if sliceId == "101" {
		ExitText()
		return nil
	}

	sliceKey := x["Слайс №"+sliceId]
	if sliceKey != nil {
		newSliceObject := getNewSlice(sliceKey)
		x["Слайс №"+sliceId] = newSliceObject
		return x
	} else {
		fmt.Println("\tХммм такого слайса нет.")
		return x
	}
}

func getNewSlice(x []int) []int {
	copy_x := make([]int, len(x))
	copy(copy_x, x)

	for i := 0; i < len(x); i++ {
		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║                 Подробнее                  ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║          Ввести число для слайса           ║")
		fmt.Println("║              Для Отмены: 101               ║")
		fmt.Println("║     Для изменения конкретного id: 202      ║")
		fmt.Println("╚════════════════════════════════════════════╝")
		fmt.Print("\nВведите число в слайсе для индекса ", i, ": \n\t\t\t")
		fmt.Scan(&copy_x[i])

		if copy_x[i] == 101 {
			AnimateLoading()
			ExitText()
			return x
		}

		if copy_x[i] == 202 {
			AnimateLoading()
			x = modifySlice(x)
			return x
		}
	}

	AnimateLoading()

	return copy_x
}
