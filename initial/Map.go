package initial

import (
	"strings"
)

func Create_the_map(array []string, arg string) string {
	myMap := make(map[rune][]string)
	var q rune = 32
	for i := 1; i < len(array); i += 9 {
		myMap[q] = array[i : i+8]
		q++
	}
	str := ""
	k := strings.ReplaceAll(arg, "\\n", "\n")
	rk := strings.Split(k, "\n")
	for _, rn := range rk {
		if rn == "" {
			str += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for d := 0; d < len(rn); d++ {
					str += myMap[rune(rn[d])][i]
				}
				str += "\n"
			}
		}
	}
	return str
}
