package task2

import "strconv"

func ParseInt(num string) (int, error) {
	return strconv.Atoi(num)
}
