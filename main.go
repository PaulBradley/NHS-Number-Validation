package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"strconv"
	"strings"
)

//export isValidNHSNumber
func isValidNHSNumber(cValue *C.char) int {
	value := C.GoString(cValue)

	var intSUM int
	var intREM int
	var nhsNumber string
	var checkDigit string
	var digit [9]string

	intSUM = 0
	intREM = 0
	nhsNumber = strings.Replace(value, " ", "", -1)
	if len(nhsNumber) != 10 {
		return 0 // not valid
	}

	checkDigit = string(nhsNumber[9:])
	for i := 0; i <= 8; i++ {
		x, _ := strconv.Atoi(string(nhsNumber[i]))
		digit[i] = strconv.Itoa(x * (11 - (i + 1)))
	}
	for i := 0; i <= 8; i++ {
		x, _ := strconv.Atoi(string(digit[i]))
		intSUM = intSUM + x
	}
	intREM = intSUM % 11
	calculatedCheckDigit := 11 - int(intREM)
	if int(calculatedCheckDigit) == 11 {
		calculatedCheckDigit = 0
	}
	x, _ := strconv.Atoi(checkDigit)
	if int(calculatedCheckDigit) == x {
		return 1 // valid
	}
	return 0 // not valid
}

func main() {}
