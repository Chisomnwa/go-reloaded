package functions

import (
	"strconv"
)

func Hexa_to_Decimal(s string, base int, bitSize int) (i int64, err error) {
	/*
	Converts a hexadecimal to decimal number.
	Args:
		s: the hexadecimal string
		base: the base value we wish to convert to
		bitSize: the bit size of the result we wish to obtain

	Outputs:
		i: the base 10 number
		err: error message in case of errors
	*/

	// Use the parseInt() function to convert
	decimal_num, err := strconv.ParseInt(s, 16, 64)
	
	// in case of any error
	if err != nil {
		return 0, err
	}

	return decimal_num, err
}
