package main

import (
    "fmt"

    "gocalc/pkg/calc"
)

func main() {
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