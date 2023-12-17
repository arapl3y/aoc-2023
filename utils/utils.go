package utils

import (
	"strconv"
)

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func ConvertToInt(s string) int {
	intChar, err := strconv.Atoi(s)

	if err != nil {
		panic("could not convert numeric char to integer: " + s)
	}

	return intChar
}
