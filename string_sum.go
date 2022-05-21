package string_sum

import (
	"errors"
	"strconv"
	"strings"
)

func StringSum(input string) (output string, err error) {
	slice := strings.Split(input, "")
	if len(slice) < 3 || len(slice) > 4 {
		var errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
		return "", errorNotTwoOperands
	}
	var slice2 []string
	for i := 0; i < len(slice); i++ {
		if slice[i] == "-" {
			k := slice[i] + slice[i+1]
			i++
			slice2 = append(slice2, k)
		} else {
			slice2 = append(slice2, slice[i])
		}
	}
	sum := 0
	for _, n := range slice2 {
		k, _ := strconv.Atoi(n)
		sum = sum + k
	}
	if sum == 0 {
		var errorEmptyInput = errors.New("input is empty")
		output = ""
		return output, errorEmptyInput
	}
	return strconv.Itoa(sum), nil
}
