package utils

import (
	"fmt"
	"time"
)

func AnimateLoading() {
	fmt.Print("\t Загрузка [")
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Print("█")
	}
	fmt.Print("]\n")

	fmt.Print("\r\n\t Загрузка завершена!\n")
}

func ExitText() {
	fmt.Println("╔═══════════════════════════════════════════╗")
	fmt.Println("║           Совершается Выход               ║")
	fmt.Println("╠═══════════════════════════════════════════╣")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("║             Выход совершен                ║")
	fmt.Println("║                Успешно!!!                 ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
}
