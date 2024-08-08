package main

import (
	"flag"
	"fmt"
	add "life/addition1"
	"time"
)

// Основная функция программы.
func main() {
	// Инициализация флагов командной строки
	help := flag.Bool("help", false, "Show help message")
	verbose := flag.Bool("verbose", false, "Display detailed simulation information")
	delayMs := flag.Int("delay-ms", 2500, "Set animation speed in milliseconds")
	flag.Parse()
	validFlags := map[string]bool{
		"--help":     true,
		"--verbose":  true,
		"--delay-ms": true,
	}
	for _, arg := range flag.Args() {
		if arg[0] == '-' && !validFlags[arg] {
			fmt.Printf("flag provided but not defined: %s\n", arg)
			return
		}
	}
	if len(flag.Args()) > 0 {
		fmt.Printf("flag provided but not defined: %s\n", flag.Args()[0])

		return
	}
	// we need to check help to output information about programm, then we end our game
	if *help {
		fmt.Println("Usage: go run main.go [options]")
		fmt.Println("\nOptions:")
		fmt.Println("--help         : Show the help message and exit")
		fmt.Println("--verbose      : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
		fmt.Println("--delay-ms=X   : Set the animation speed in milliseconds. Default is 2500 milliseconds")
		return
	}
	// rows хранит количество строк в grid, cols хранит количество столбцов . grid - представляет собой игровое поле.
	var rows, cols int
	var grid [][]rune
	// Запрашиваем размер grid, затем если введены неправильные данные, выводим ошибку и завершаем программу
	fmt.Println("Enter grid size, (for example: 3 3; Remember input can't be lower than 3 3!):")
	var r, c int
	_, err := fmt.Scanf("%d %d\n", &r, &c)
	if err != nil || r < 3 || c < 3 {
		fmt.Println("Invalid grid size, please enter numbers >= 3.")
		return
	}
	rows, cols = r, c

	// здесь запрашиваем мертвые и живые клетки, если неправильно введены данные, выводим ошибку и завершаем программу
	grid = make([][]rune, rows)
	fmt.Printf("Enter grid configuration (use '#' for live cells and '.' for dead cells):\n")
	for i := 0; i < rows; i++ {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil || len(input) != cols {
			fmt.Println("Invalid row length, please try again.")
			return
		}
		grid[i] = []rune(input)
	}

	step := 0
	// Игровой цикл, который продолжается до тех пор, пока есть живые клетки в сетке (addition.Radiohead(grid) возвращает true).
	for {
		if !add.Radiohead(grid) {
			break
		}
		step++
		// условия если используем флаг verbose
		if *verbose {
			fmt.Printf("Step: %d\n", step)
			fmt.Printf("Grid Size: %dx%d\n", rows, cols)
			// deftones считает количество живых клеток
			liveCells := add.Deftones(grid)
			fmt.Printf("Live Cells: %d\n", liveCells)
			fmt.Printf("DelayMs: %dms\n", *delayMs)
		}
		// greenday выводит сетку
		add.Greenday(grid)
		time.Sleep(time.Duration(*delayMs) * time.Millisecond)
		// mindlessselfindulgence генерирует новое состояние сетки
		grid = add.MindlessSelfIndulgence(grid)
		fmt.Println()
	}
}
