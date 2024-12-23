# Спринт 0-1. Калькулятор

## Описание
1. Реализовать функцию `func Calc(expression string) (float64, error)`.  
`expression` - строка-выражение, состоящее из односимвольных идентификаторов и знаков арифметических действий.
2. Написать сервис для решения арифметических выражений, использующий формат json

## Входящие данные
- Цифры (рациональные)
- Операции: `+`, `-`, `*`, `/`
- Операции приоритезации: `(` и `)`

## Пример использования (импорт gocacl)
```go
result, err := Calc("3 + 2 * (1 + 2)")
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Результат:", result)
}
```

## Запуск программы (консольный режим)
Для запуска программы выполните следующую команду:
```shell
export GOPATH := $(shell pwd) # Only For Linux
setx GOPATH "%cd%" # Only For Windows
go run cmd\calc\console.go
```

## Запуск программы (режим cервиса)
Для запуска программы выполните следующую команду:
```shell
export GOPATH := $(shell pwd) # Only For Linux
setx GOPATH "%cd%" # Only For Windows
go run cmd\server\server.go --ipPort :8123 # Listen on 8123 port
```

## Запуск программы в docker (режим cервиса)
Для запуска программы выполните следующую команду:
```shell
docker build --rm -f "dockerfile" -t gocacl:latest "." 
docker run --rm -d -p 8080:8080/tcp gocacl:latest
```

## Примеры (режим cервиса)
```shell
>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2+2*2\" }"
{"result":6}

>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2+2*2\" }"
{"result":6}

C:\Users\denis\YandexDisk\Рабочий стол\GOCACL>curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2/2\" }" 
{"result":1}

>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2/0\" }"
{"error":"деление на ноль"}

>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2+2*\" }" 
{"error":"некорректное выражение"}

>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"\" }"
{"error":"некорректное выражение"}

>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \" }"
{"error":"unexpected EOF"}

>>> curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"Python идет за тобой\" }"
{"error":"Горфер паникует"}
```

## Запуск тестов
Для запуска тестов выполните следующую команду:
```shell
export GOPATH := $(shell pwd) # Only For Linux
setx GOPATH "%cd%" # Only For Windows
go test ./... --cover
```

## Контакты
https://t.me/denisgrigoriev04
