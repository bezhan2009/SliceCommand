package utils

import (
	"fmt"
	"sort"
)

func AddSlice() ([]int, bool) {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 Подробнее                  ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║            Создать новый слайс             ║")
	fmt.Println("║              Для Отмены: 101               ║")
	fmt.Println("╚════════════════════════════════════════════╝")

	var newLenSlice int
	fmt.Print("Введите сколько значений вы хотите в слайсе (По умолчанию 5): ")
	fmt.Scan(&newLenSlice)
	if newLenSlice == 101 {
		return nil, false
	}

	if newLenSlice == 0 {
		newLenSlice = 5
	}

	newSlice := make([]int, newLenSlice)
	for i := 0; i < newLenSlice; i++ {
		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║                 Подробнее                  ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║            Создать новый слайс             ║")
		fmt.Println("║Для Отмены (удаления созданного слайса): 101║")
		fmt.Println("║         Пропустить эту часть: 909          ║")
		fmt.Println("╚════════════════════════════════════════════╝")
		fmt.Printf("\n\t\tВведите данные для нового слайса (id: %d): ", i)
		var newValue int
		fmt.Scan(&newValue)

		if newValue == 101 {
			return nil, false
		} else if newValue == 909 {
			fmt.Println("\tНовый слайс: ", newSlice)
			return newSlice, true
		} else {
			newSlice[i] = newValue
		}
	}

	fmt.Println("Новый слайс: ", newSlice)
	return newSlice, true
}

func GetSlices(x map[string][]int) {
	keys := make([]string, 0, len(x))
	for key := range x {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if len(keys) > 0 {
		fmt.Println("Слайсы:")
		for _, key := range keys {
			fmt.Printf("\t%s %v\n", key, x[key])
		}
	} else {
		fmt.Println("\n\tУ вас пока нет слайсов")
	}
}
