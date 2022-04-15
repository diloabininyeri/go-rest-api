package functions

import "strconv"

func StrToInt(stringInt string) int {

	number, err := strconv.ParseUint(stringInt, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(number) //Convert uint64 To int
}
