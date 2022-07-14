package google

import (
	"fmt"
)

func main() {
	var a, ctr int
	fmt.Scanln(&a)
TestCases:
	if a > 0 {
		ctr++
		var s1, s2 string
		fmt.Scanln(&s1)
		fmt.Scanln(&s2)
		fmt.Printf("Case #%d: %d\n", ctr, TransformString(s1, s2))
		a--
		goto TestCases
	}
}

func TransformString(s string, t string) int {
	var result int
OUTER:
	for _, a := range s {
		min := 27
		for _, b := range t {
			if a == b {
				continue OUTER
			}
			if a > b {
				a, b = b, a
			}
			diff := Min(int(b-a), int('z'-b+a-'a'+1))
			min = Min(min, diff)
		}
		result += int(min)
	}
	return result
}

func Min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
