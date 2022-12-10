package module01

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) string {
	var sb strings.Builder

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			sb.WriteString("FizzBuzz ")
		} else if i%3 == 0 {
			sb.WriteString("Fizz ")
		} else if i%5 == 0 {
			sb.WriteString("Buzz ")
		} else {
			sb.WriteString(strconv.Itoa(i) + " ")
		}
	}
	return strings.TrimSuffix(sb.String(), " ")
}
