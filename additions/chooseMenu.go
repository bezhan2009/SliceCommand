package additions

import (
	"KarimovbejanHW/hw_tasks/utils"
	"fmt"
)

func ChooseMenu() int {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 Подробнее                  ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║            Выбрать новое меню              ║")
	fmt.Println("╠────────────────────────────────────────────╣")
	fmt.Println("║                 Типы меню                  ║")
	fmt.Println("╠────────────────────────────────────────────╣")
	fmt.Println("║      1. Circle Menu                        ║")
	fmt.Println("║      2. Default Menu                       ║")
	fmt.Println("║      3. Full Menu                          ║")
	fmt.Println("║      4. Strange Menu                       ║")
	fmt.Println("║      5. Super Menu                         ║")
	fmt.Println("║      6. Unusual Menu                       ║")
	fmt.Println("╚════════════════════════════════════════════╝")

	var typeMenu int
	fmt.Print("Какое меню хотите рассмотреть: ")
	fmt.Scan(&typeMenu)

	if typeMenu > 6 {
		fmt.Println("Нет такого типа меню")
		return 10
	} else {
		menuTypeChoosing, chosen := MainMenu(typeMenu, false)
		if chosen == false {
			utils.ExitText()
		} else {
			fmt.Println("\tПоздравляю с новым видом меню!!!")
			return menuTypeChoosing
		}
	}

	return 10
}
