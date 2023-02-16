package wrapper

import(
	"strings"
)

func WrapString(str string, n int) string {
	var count int = 0
	var s string = ""
	input := strings.Fields(str)
	for i := 0; i < len(input); i++ {
		count += 1
		if count == n-1 {
			s += input[i]+"\n"
			count = 0
		} else {
			s += input[i]+" "
		}
	}

	return s
}