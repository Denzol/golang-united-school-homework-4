package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	var slice1 []string
	var slice2 []string
	var number1 []string
	var number2 []string
	var number3 []string
	slice1 = strings.Split(input, "")

	//fmt.Println("slice1", slice1)

	for _, a := range slice1 {
		if a != " " {
			slice2 = append(slice2, a)
		}
	}

	//fmt.Println("slice2", slice2)

	for i := 0; i < len(slice2); i++ {
		if slice2[i] == "+" || slice2[i] == "-" {
			number1 = append(number1, slice2[0])
		}
		if slice2[i] != "+" && slice2[i] != "-" {
			number1 = append(number1, slice2[i])
		}
		if i < len(slice2)-1 && (slice2[i+1] == "+" || slice2[i+1] == "-") {
			break
		}
	}
	//fmt.Println("number1", number1)

	if len(number1) != len(slice2) {
		for i := len(number1); i <= len(slice2)-1; i++ {
			number2 = append(number2, slice2[i])
			if i < len(slice2)-1 && (slice2[i+1] == "+" || slice2[i+1] == "-") {
				break
			}
		}
	}

	//fmt.Println("number2", number2)

	if len(number1)+len(number2) != len(slice2) {
		for i := len(number1) + len(number2); i <= len(slice2)-1; i++ {
			number3 = append(number3, slice2[i])
		}
	}

	//fmt.Println("number3", number3)

	if len(number1) == 0 && len(number2) == 0 && len(number3) == 0 {
		err := fmt.Errorf("bad token %w", errorEmptyInput)
		return "", err
	}

	number1ToString := strings.Join(number1, "") // преобразуем слайс с первым операндом в стринг
	//fmt.Println("number1ToString", number1ToString)
	number2ToString := strings.Join(number2, "") // преобразуем слайс со вторым операндом в стринг
	//fmt.Println("number2ToString", number2ToString)
	number3ToString := strings.Join(number3, "") // преобразуем слайс со вторым операндом в стринг
	//fmt.Println("number3ToString", number3ToString)

	stringToInt1, err1 := strconv.Atoi(number1ToString)
	if err1 != nil {
		err1 = fmt.Errorf("incorrect input %w", err1)
		return "", err1
	}

	//	fmt.Println("stringToInt1", stringToInt1)

	stringToInt2, err2 := strconv.Atoi(number2ToString)
	if err2 != nil && len(number2) == 0 {
		err2 = fmt.Errorf("bad token %w", errorNotTwoOperands)
		return "", err2
	} else if err2 != nil {
		err2 = fmt.Errorf("bad token %w", err2)
		return "", err2
	}

	//fmt.Println("stringToInt2", stringToInt2)

	stringToInt3, err3 := strconv.Atoi(number3ToString)
	if err3 != nil {
		err3 = fmt.Errorf("%w", err3)
	} else if err3 == nil {
		err3 = fmt.Errorf("bad token %w", errorNotTwoOperands)
		return "", err3
	}

	//fmt.Println("stringToInt3", stringToInt3)

	if err1 != nil || err2 != nil {
		err = fmt.Errorf("incorrect input %w", err)
		return "", err
	}
	e := stringToInt1 + stringToInt2
	return strconv.Itoa(e), nil

}
