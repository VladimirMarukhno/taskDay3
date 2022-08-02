package main

import (
	"fmt"
	"strconv"
)

/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/

const alphabet = "abcdefghijklmnopqrstuvwxyz "

func codeToString(code string, output []rune) {
	myMap := map[string]rune{}

	for key, value := range alphabet {
		tmp := strconv.Itoa(key)

		if key < 10 {
			myMap["0"+tmp] = value
		} else {
			myMap[tmp] = value
		}
	}

	for i := 0; i < len(code); i += 2 {
		tmp := string(code[i]) + string(code[i+1])
		output = append(output, myMap[tmp])
	}

	fmt.Println(string(output))
}

func main() {
	var output []rune

	code := "220411112603141304"

	codeToString(code, output)
	//myMap := map[string]rune{}
	//
	//for key, value := range alphabet {
	//	tmp := strconv.Itoa(key)
	//
	//	if key < 10 {
	//		myMap["0"+tmp] = value
	//	} else {
	//		myMap[tmp] = value
	//	}
	//}
	//
	//for i := 0; i < len(code); i += 2 {
	//	tmp := string(code[i]) + string(code[i+1])
	//	output = append(output, myMap[tmp])
	//}
	//
	//fmt.Println(string(output))
}
