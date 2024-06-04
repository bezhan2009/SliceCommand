package main

import (
	"KarimovbejanHW/hw_tasks/additions"
	"KarimovbejanHW/hw_tasks/utils"
	"fmt"
	"strconv"
)

func main() {
	var x = make(map[string][]int)
	menuType := 2
	sliceDefault := []int{1, 2, 3, 4, 5}
	x["Слайс №1"] = sliceDefault

	for {
		additions.MainMenu(menuType, true)

		var action int
		fmt.Print("\n\tВыбрать действия: \n\t\t\t")
		fmt.Scan(&action)

		if action == 0 {
			utils.ExitText()
			break
		}

		switch action {
		case 1:
			utils.AnimateLoading()
			checkResult := utils.CheckSlice(x)
			if checkResult {
				x = utils.GetIdSlice(x)
			} else {
				fmt.Println("\tУ вас пока нет слайсов!!!")
			}
		case 2:
			utils.AnimateLoading()
			newSlice, success := utils.AddSlice()
			if !success {
				utils.ExitText()
			} else {
				key := "Слайс №" + strconv.Itoa(len(x)+1)
				x[key] = newSlice
			}
		case 3:
			utils.AnimateLoading()
			utils.GetSlices(x)
			fmt.Print("\n")
		case 4:
			utils.AnimateLoading()
			checkResult := utils.CheckSlice(x)
			if checkResult {
				additions.RemoveSlice(x)
			} else {
				fmt.Println("\tУ вас пока нет слайсов!!!")
			}
		case 5:
			utils.AnimateLoading()
			checkResult := utils.CheckSlice(x)
			if checkResult {
				additions.MergeSlices(x)
				fmt.Print("\n")
			} else {
				fmt.Println("\tУ вас пока нет слайсов!!!")
			}
		case 6:
			utils.AnimateLoading()
			checkResult := utils.CheckSlice(x)
			if checkResult {
				var (
					sliceKey    string
					sliceValues []int
				)
				utils.GetSlices(x)
				fmt.Print("Введите номер слайса: ")
				if _, err := fmt.Scan(&sliceKey); err != nil {
					fmt.Println("\tОшибка при считывании первого номера слайса:", err)
					//return nil
				}

				key := "Слайс №" + sliceKey

				sliceValues = x[key]

				if sliceValues == nil {
					fmt.Println("\tСлайс не найден!!!")
					//return nil
				} else {
					additions.SortSlice(sliceValues)
				}
			} else {
				fmt.Println("\tУ вас пока нет слайсов!!!")
			}
		case 7:
			utils.AnimateLoading()
			checkResult := utils.CheckSlice(x)
			if checkResult {
				var (
					sliceKey       string
					numCycleShifts int
					sliceCycle     []int
				)
				utils.GetSlices(x)
				fmt.Print("Введите номер слайса: ")
				if _, err := fmt.Scan(&sliceKey); err != nil {
					fmt.Println("\tОшибка при считывании первого номера слайса:", err)
					//return nil
				}

				key := "Слайс №" + sliceKey

				sliceCycle = x[key]

				if sliceCycle == nil {
					fmt.Println("\tСлайс не найден!!!")
					//return nil
				}

				fmt.Print("\nВведите количество циклических сдвигов: ")
				if _, err := fmt.Scan(&numCycleShifts); err != nil {
					fmt.Println("\tОшибка при считывании циклических сдвигов:", err)
					//return nil
				}

				if sliceCycle != nil {
					cycleShift := additions.CycleShift(sliceCycle, numCycleShifts)
					x[key] = cycleShift
				}
			} else {
				fmt.Println("\tУ вас пока нет слайсов!!!")
			}
		case 8:
			utils.AnimateLoading()
			choosingMenu := additions.ChooseMenu()
			if choosingMenu != 10 {
				menuType = choosingMenu
			}
		case 9:
			utils.AnimateLoading()
			additions.AboutMe()
		default:
			utils.AnimateLoading()
			fmt.Println("\tНе опознанное действие!!!")
		}
	}
}

// version 1.3 :)
