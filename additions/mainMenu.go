package additions

import (
	"KarimovbejanHW/hw_tasks/typeMenu"
	"fmt"
)

func MainMenu(menuType int, isUse bool) (int, bool) {
	if menuType == 1 {
		typeMenu.CircleMenu()
	} else if menuType == 2 {
		typeMenu.DefaultMenu()
	} else if menuType == 3 {
		typeMenu.FullMenu()
	} else if menuType == 4 {
		typeMenu.StrangeMenu()
	} else if menuType == 5 {
		typeMenu.SuperMenu()
	} else if menuType == 6 {
		typeMenu.UnusualMenu()
	}

	if !isUse {
		for {
			fmt.Println("╔════════════════════════════════════════════╗")
			fmt.Println("║                 Подробнее                  ║")
			fmt.Println("╠════════════════════════════════════════════╣")
			fmt.Println("║               Принять: 606                 ║")
			fmt.Println("║              Для Отмены: 101               ║")
			fmt.Println("╚════════════════════════════════════════════╝")

			var action int
			fmt.Print("\n\tВыбрать действия: \n\t\t\t")
			fmt.Scan(&action)

			if action == 101 {
				return menuType, false
			} else if action == 606 {
				return menuType, true
			} else {
				fmt.Println("\tНе опознанное действия!!!")
			}
		}
	}

	return menuType, false
}
