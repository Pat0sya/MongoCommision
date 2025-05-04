package main

import (
	"RBD_dev/config"
	"RBD_dev/services"
	"RBD_dev/utils"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// clearScreen очищает экран терминала (работает и в Windows, и в Unix-подобных системах)
func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// waitEnter ждет нажатия Enter для возврата в главное меню
func waitEnter(reader *bufio.Reader) {
	fmt.Println("\nНажмите Enter для возврата в меню...")
	reader.ReadString('\n')
}

func main() {
	config.ConnectDB()
	utils.InitCollections()

	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		myFigure := figure.NewFigure("Welcome", "larry3d", true)
		myFigure.Print()
		fmt.Println("╔═══════════════════════════╗")
		fmt.Println("║ 1. Абитуриенты            ║")
		fmt.Println("║ 2. Документы              ║")
		fmt.Println("║ 3. Факультеты             ║")
		fmt.Println("║ 4. Программы              ║")
		fmt.Println("║ 5. Заявления              ║")
		fmt.Println("║ 0. Выход                  ║")
		fmt.Println("╚═══════════════════════════╝")
		fmt.Print("Выберите раздел: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		clearScreen()
		switch choice {
		case "1":
			myFigure := figure.NewFigure("Applicants", "larry3d", true)
			myFigure.Print()
			fmt.Println("1. Добавить абитуриента")
			fmt.Println("2. Показать всех")
			fmt.Println("3. Удалить абитуриента")
			fmt.Print("Выберите действие: ")
			sub, _ := reader.ReadString('\n')
			switch strings.TrimSpace(sub) {
			case "1":
				services.AddApplicant(reader)
			case "2":
				services.ListApplicants()
			case "3":
				services.DeleteApplicant(reader)
			default:
				fmt.Println("Неверный выбор")
			}
			waitEnter(reader)

		case "2":
			myFigure := figure.NewFigure("Documents", "larry3d", true)
			myFigure.Print()
			fmt.Println("1. Добавить документ")
			fmt.Println("2. Показать документы")
			fmt.Println("3. Удалить документ")
			fmt.Print("Выберите действие: ")
			sub, _ := reader.ReadString('\n')
			switch strings.TrimSpace(sub) {
			case "1":
				services.AddDocument(reader)
			case "2":
				services.ShowDocuments()
			case "3":
				services.DeleteDocument(reader)
			default:
				fmt.Println("Неверный выбор")
			}
			waitEnter(reader)

		case "3":
			myFigure := figure.NewFigure("Faculty", "larry3d", true)
			myFigure.Print()
			fmt.Println("1. Добавить факультет")
			fmt.Println("2. Показать факультеты")
			fmt.Print("Выберите действие: ")
			sub, _ := reader.ReadString('\n')
			switch strings.TrimSpace(sub) {
			case "1":
				services.AddFaculty(reader)
			case "2":
				services.ListFaculties()
			default:
				fmt.Println("Неверный выбор")
			}
			waitEnter(reader)

		case "4":
			myFigure := figure.NewFigure("Programs", "larry3d", true)
			myFigure.Print()
			fmt.Println("1. Добавить программу")
			fmt.Println("2. Показать программы")
			fmt.Print("Выберите действие: ")
			sub, _ := reader.ReadString('\n')
			switch strings.TrimSpace(sub) {
			case "1":
				services.AddProgram(reader)
			case "2":
				services.ListPrograms()
			default:
				fmt.Println("Неверный выбор")
			}
			waitEnter(reader)

		case "5":
			myFigure := figure.NewFigure("Application", "larry3d", true)
			myFigure.Print()
			fmt.Println("1. Подать заявление")
			fmt.Println("2. Показать заявления")
			fmt.Print("Выберите действие: ")
			sub, _ := reader.ReadString('\n')
			switch strings.TrimSpace(sub) {
			case "1":
				services.AddApplication(reader)
			case "2":
				services.ListApplications()
			default:
				fmt.Println("Неверный выбор")
			}
			waitEnter(reader)

		case "0":
			fmt.Println("Выход...")
			return

		default:
			fmt.Println("Неверный выбор")
			waitEnter(reader)
		}
	}
}
