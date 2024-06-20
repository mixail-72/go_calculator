package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b, c, roman_result string   // первое число, второе число, оператор, результат в римской системе счисления
	var system_a, system_b, result int // переменные, введенные для опередения системы счисления + результат

	fmt.Println("Введите уравнение: ")
	equasion := bufio.NewReader(os.Stdin)
	line, _ := equasion.ReadString('\n')
	line = strings.ReplaceAll(line, " ", "") // убираем все пробелы

	a, b, c = extractor(line)

	a, system_a = converter(a, system_a) // конвертируем строку в число и проверяем, является оно римским или арабским

	b, system_b = converter(b, system_b) // конвертируем строку в число и проверяем, является оно римским или арабским

	if system_a == system_b && system_a == 0 {
		result = calculate(a, b, c)
		if result < 101 {
			fmt.Println("Result: ", result)
		} else {
			fmt.Println("Превышено верхнее значение")
			os.Exit(0)
		}
	} else if system_a == system_b && system_a == 1 {
		result = calculate(a, b, c)

		roman_result = convert_to_roman(result)

		if result < 101 {
			fmt.Println("result: ", roman_result)
		} else {
			fmt.Println("Превышено верхнее значение")
			os.Exit(0)
		}

	} else {
		fmt.Println("Используются разные системы счисления")
		os.Exit(0)
	}
}

func extractor(line string) (string, string, string) { // проверяет, правильно ли записано выражение

	var n1, n2, operator string

	TotalPlus := strings.Count(line, "+")        // считает все "+"
	num_in_line_Plus := strings.Index(line, "+") // вхождение "+"
	if TotalPlus > 0 {
		n1 = line[0:num_in_line_Plus]                          // обрезает строку с начала до пробела перед знаком ариф. д-я
		n2 = line[num_in_line_Plus+1 : len(line)-1]            // обрезает строку после пробела после знака ариф. д-я и до конца строки
		operator = line[num_in_line_Plus : num_in_line_Plus+1] // вырезает "+"
	}

	TotalMinus := strings.Count(line, "-")        // считает все "-"
	num_in_line_Minus := strings.Index(line, "-") // вхождение "-"
	if TotalMinus > 0 {
		n1 = line[0:num_in_line_Minus]
		n2 = line[num_in_line_Minus+1 : len(line)-1]
		operator = line[num_in_line_Minus : num_in_line_Minus+1]
	}

	TotalMultiply := strings.Count(line, "*")        // считает все "*"
	num_in_line_Multiply := strings.Index(line, "*") // вхождение "*"
	if TotalMultiply > 0 {
		n1 = line[0:num_in_line_Multiply]
		n2 = line[num_in_line_Multiply+1 : len(line)-1]
		operator = line[num_in_line_Multiply : num_in_line_Multiply+1]
	}

	TotalDivide := strings.Count(line, "/")        // считает все "/"
	num_in_line_Divide := strings.Index(line, "/") // вхождение "/"
	if TotalDivide > 0 {
		n1 = line[0:num_in_line_Divide]
		n2 = line[num_in_line_Divide+1 : len(line)-1]
		operator = line[num_in_line_Divide : num_in_line_Divide+1]
	}

	if TotalPlus+TotalMinus+TotalMultiply+TotalDivide > 1 { // если знаков ариф. д-я больше одного или их нет вообще, выполнение программы останавливается
		fmt.Println("Введены лишние операторы")
		os.Exit(0)
	} else if TotalPlus+TotalMinus+TotalMultiply+TotalDivide == 0 {
		fmt.Println("Отсутствует оператор", "+", "-", "*", "или /")
		os.Exit(0)
	}
	return n1, n2, operator // выдает число до оператора, после него и сам оператор
}

func converter(n string, is_roman int) (string, int) { //переводит римские числа от i до x в арабские

	switch n {
	case "i":
		n = "1"
		is_roman = 1
	case "ii":
		n = "2"
		is_roman = 1
	case "iii":
		n = "3"
		is_roman = 1
	case "iv":
		n = "4"
		is_roman = 1
	case "v":
		n = "5"
		is_roman = 1
	case "vi":
		n = "6"
		is_roman = 1
	case "vii":
		n = "7"
		is_roman = 1
	case "viii":
		n = "8"
		is_roman = 1
	case "ix":
		n = "9"
		is_roman = 1
	case "x":
		n = "10"
		is_roman = 1
	case "1":
		n = "1"
		is_roman = 0
	case "2":
		n = "2"
		is_roman = 0
	case "3":
		n = "3"
		is_roman = 0
	case "4":
		n = "4"
		is_roman = 0
	case "5":
		n = "5"
		is_roman = 0
	case "6":
		n = "6"
		is_roman = 0
	case "7":
		n = "7"
		is_roman = 0
	case "8":
		n = "8"
		is_roman = 0
	case "9":
		n = "9"
		is_roman = 0
	case "10":
		n = "10"
		is_roman = 0
	default:
		fmt.Println("Введено число, не удовлетворяющее требованиям")
		os.Exit(0)
	}
	return n, is_roman
}

func calculate(x string, y string, z string) int { // x - первое число, y - второе, z - знак ариф. д-я
	if z == "+" {
		x, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(y)
		if err != nil {
			log.Fatal(err)
		}
		return x + y
	} else if z == "-" {
		x, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(y)
		if err != nil {
			log.Fatal(err)
		}
		return x - y
	} else if z == "*" {
		x, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(y)
		if err != nil {
			log.Fatal(err)
		}
		return x * y
	} else if z == "/" {
		x, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(y)
		if err != nil {
			log.Fatal(err)
		}
		return x / y
	} else {
		return 0
	}
}

func convert_to_roman(number int) string {

	var x_num int = number / 10     // сколько десятков в числе
	var x_num_rem int = number % 10 // остаток от деления на 10 (есть i и v)
	var res, v_res, i_res string

	if number <= 0 { // если результат 0 или отриц., программа завершается
		fmt.Println("В римской системе нет 0 и отрицательных чисел")
		os.Exit(0)
	}

	if number > 0 && number < 49 {
		for i := 0; i < x_num; i++ {
			res += "x"
		}
		if x_num_rem == 4 {
			v_res = "iv"
		} else if x_num_rem == 5 || x_num_rem == 6 || x_num_rem == 7 || x_num_rem == 8 {
			v_res = "v"
		} else if x_num_rem == 9 {
			i_res = "ix"
		}
		if number%5 > 0 && number%5 < 4 {
			for i := 0; i < (number % 5); i++ {
				i_res += "i"
			}
		}
	}
	res = res + v_res + i_res

	if number >= 49 && number < 100 {
		if number == 49 {
			res = "il"
		} else {
			for i := 0; i < x_num-5; i++ {
				res += "x"
			}
			if x_num_rem == 4 {
				v_res = "iv"
			} else if x_num_rem == 5 || x_num_rem == 6 || x_num_rem == 7 || x_num_rem == 8 {
				v_res = "v"
			} else if x_num_rem == 9 {
				i_res = "ix"
			}
			if number%5 > 0 && number%5 < 4 {
				for i := 0; i < (number % 5); i++ {
					i_res += "i"
				}
			}
			res = "l" + res + v_res + i_res
		}
	} else if number == 100 {
		res = "c"
	}

	return res
}
