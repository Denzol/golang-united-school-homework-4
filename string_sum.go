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
	slice1              []string
	slice2              []string
	input3              []string
)

func StringSum(input string) (output string, err error) {
	input2 := strings.Split(input, "")
	for _, n := range input2 {
		if n != " " {
			input3 = append(input3, n)
		}
	}
	for i, k := range input3 {
		if i == 0 || k != "+" && k != "-" {
			slice1 = append(slice1, k)
		} else if k == "+" || k == "-" {
			slice2 = input3[len(slice1):]
			break
		}
	}

	sliceToString1 := strings.Join(slice1, "")        // преобразуем слайс с первым операндом в стринг
	sliceToString2 := strings.Join(slice2, "")        // преобразуем слайс со вторым операндом в стринг
	stringToInt1, err := strconv.Atoi(sliceToString1) // преобразуем стринг с первым операндом в инт
	stringToInt2, err := strconv.Atoi(sliceToString2) // преобразуем стринг со вторым операндом в инт
	if len(slice1) == 0 && len(slice2) == 0 {
		err = fmt.Errorf("bad token %w", errorEmptyInput)
		return "", err
	} else if err != nil && len(slice1) != 0 && stringToInt1 == 0 {
		err = fmt.Errorf("incorrect input %w", err)
		return "", err
	}
	if err != nil && stringToInt2 == 0 && stringToInt1 != 0 {
		err = fmt.Errorf("bad token %w", errorNotTwoOperands)
		return "", err
	}
	e := stringToInt1 + stringToInt2
	output, err = strconv.Itoa(e) // преобразуем результат в стринг
	return output, err
}
