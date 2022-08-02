package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

const (
	special   = `_!@#$%^&`
	lowerCase = `abcdefghijklmnopqrstuvwxyz`
	upperCase = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digits    = `0123456789`
	attempt   = 5
	maxlength = 15
	minlength = 8
)

// Вводим пароль
func inputPassword(p *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*p = scanner.Text()
	fmt.Print("=============================================================\n")
}

// Проверяем пароль на корректность
func valid(p *string, hasReplace, hasLower, hasUpper, hasDigits, length, invalid *bool) {

	chars := map[rune]rune{}
	for _, value := range *p {
		chars[value] = 0
	}
	fmt.Println(len(*p))
	if len(*p) < minlength || len(*p) > maxlength {
		*length = false
	}

	for value := range chars {
		switch {
		case strings.ContainsRune(special, value):
			*hasReplace = true
		case strings.ContainsRune(lowerCase, value):
			*hasLower = true
		case strings.ContainsRune(upperCase, value):
			*hasUpper = true
		case strings.ContainsRune(digits, value):
			*hasDigits = true
		default:
			*invalid = true
		}
	}
}

// Выводим сообщения об ошибках во время ввода.
func messageError(hasReplace, hasLower, hasUpper, hasDigits, length, invalid *bool) {

	errMessage := make([]string, 0, 6)

	if *invalid {
		errMessage = append(errMessage, "Используются недопустимые символы!")
	} else {
		if !*hasReplace {
			errMessage = append(errMessage, "Используйте спецсимволы!")
		}
		if !*hasLower {
			errMessage = append(errMessage, "Используйте буквы в нижнем регистре!")
		}
		if !*hasUpper {
			errMessage = append(errMessage, "Используйте буквы в верхнем регистре!")
		}
		if !*hasDigits {
			errMessage = append(errMessage, "Используйте цифры!")
		}
		if !*length {
			errMessage = append(errMessage, "Длина пароля не соответствует заданным параметрам!")
		}
	}
	for _, elem := range errMessage {
		fmt.Println(elem)
	}
}

func main() {
	fmt.Print("Введите пароль, состоящий из: цифр, букв латинского алфавита(строчные и прописные) и специальных символов\n" +
		"Длина пароля должна быть от 8(мин) до 15(макс) символов. Максимальное количество попыток ввода пароля - 5\n")
	var password string
	for count := 1; count <= attempt; count++ {

		hasReplace := false
		hasLower := false
		hasUpper := false
		hasDigits := false
		invalid := false
		length := true

		inputPassword(&password)

		valid(&password, &hasReplace, &hasLower, &hasUpper, &hasDigits, &length, &invalid)

		if !hasReplace || !hasLower || !hasUpper || !hasDigits || invalid || !length {
			messageError(&hasReplace, &hasLower, &hasUpper, &hasDigits, &length, &invalid)
			fmt.Printf("Оставшиеся попытки ввода %d \n", attempt-count)
		} else {
			fmt.Println("Пароль введён верно!")
			break
		}
	}
}
