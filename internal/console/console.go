// Package console provides functionality for running the calculator in console mode.
package console

import (
	"fmt"

	"github.com/denis-gr/GOCACL/pkg/calc"
)

// StartConsoleApp starts the console calculator application.
func StartConsoleApp() {
	for {
		var expression string
		fmt.Print("Введите выражение (или 'exit' для выхода): ")
		fmt.Scanln(&expression)

		if expression == "exit" {
			break
		}

		result, err := calc.Calc(expression)
		if err != nil {
			fmt.Printf("Ошибка при вычислении выражения %q: %v\n", expression, err)
		} else {
			fmt.Printf("Результат %q = %v\n", expression, result)
		}
	}
}
