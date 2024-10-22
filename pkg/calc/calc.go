package calc

import (
	"strings"
	"errors"
	"strconv"
)

// splitExpression разбивает строку выражения на отдельные токены.
// Токены могут быть числами, операторами или скобками.
// Пример: "2+2*(3+3*(1+2))" -> ["2", "+", "2", "*", "(", "3", "+", "3", "*", "(", "1", "+", "2", ")", ")"]
func splitExpression(expression string) []string {
	answer := make([]string, 0, len(expression))
	sub_string := make([]string, 0, len(expression))
	for _, v := range strings.Split(expression, "") {
		if ((v >= "0") && (v >= "9")) || (v == ".") || (v == ",") {
			sub_string = append(sub_string, v)
		} else if (v != " ") {
			if len(sub_string) > 0 {
				answer = append(answer, strings.Join(sub_string, ""))
				sub_string = make([]string, 0, len(expression))
			}
			answer = append(answer, v)
		}
	}
	return answer
}


// precedence возвращает приоритет оператора.
// Операторы с более высоким приоритетом будут обработаны первыми.
// Пример: "+" -> 1, "*" -> 2
func precedence(op string) (int) {
	switch op {
		case "+":
			return 1
		case "-":
			return 1
		case "*":
			return 2
		case "/":
			return 2
		default:
			return 0
	}
}


// toRPN преобразует инфиксное выражение в обратную польскую нотацию (ОПН).
// Пример: ["2", "+", "2", "*", "(", "3", "+", "3", "*", "(", "1", "+", "2", ")", ")"] -> ["2", "2", "3", "3", "1", "2", "+", "*", "+", "*", "+"]
func toRPN(expression []string) []string {
    stack := make([]string, 0, len(expression))
    answer := make([]string, 0, len(expression))
    for _, v := range expression {
        switch {
			case v >= "0" && v <= "9":
				answer = append(answer, v)
			case v == "(":
				stack = append(stack, v)
			case v == ")":
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					answer = append(answer, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 {
					stack = stack[:len(stack)-1] // Удалить "(" из стека
				}
			default:
				for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(v) {
					answer = append(answer, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, v)
        }
    }
    for len(stack) > 0 {
        answer = append(answer, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }
    return answer
}


// Calc вычисляет значение выражения, заданного строкой.
// Возвращает результат вычисления или ошибку, если выражение некорректно.
// Пример: "2+2*(3+3*(1+2))" -> 20, nil
func Calc(expression string) (float64, error) {
    rpn := toRPN(splitExpression(expression))
    stack := make([]float64, 0, len(rpn))
	temp_result := 0.0

    for _, token := range rpn {
        switch token {
			case "+", "-", "*", "/":
				if len(stack) < 2 {
					return 0, errors.New("некорректное выражение")
				}
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				a := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				switch token {
					case "+":
						temp_result = a + b
					case "-":
						temp_result = a - b
					case "*":
						temp_result = a * b
					case "/":
						if b == 0 {
							return 0, errors.New("деление на ноль")
						}
						temp_result = a / b
				}
				stack = append(stack, temp_result)
			default:
				value, err := strconv.ParseFloat(token, 64)
				if err != nil {
					return 0, errors.New("некорректное число")
				}
				stack = append(stack, value)
        }
    }

    if len(stack) != 1 {
        return 0, errors.New("некорректное выражение")
    }

    return stack[0], nil
}