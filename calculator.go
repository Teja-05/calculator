package calculator

import (
	"errors"
	"fmt"
)

func Calculate(num1 float64, num2 float64, option int) (float64, error) {
	res := 0.0

	switch option {
	case 1:
		res = num1 + num2
	case 2:
		res = num1 - num2
	case 3:
		res = num1 * num2
	case 4:
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return 0, errors.New("division by zero error")
		} else {
			res = num1 / num2
		}
	default:
		fmt.Println("Enter a valid option")
		return 0, errors.New("invalid option")
	}
	return res, nil
}
