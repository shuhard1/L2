package main

import (
	"fmt"
)

func StrUnpacking(s string) string {
	var result string = ""
	for v, r := range s {
		l := 1
		if r >= 48 && r <= 57 {
			l = int(r) - 48 - 1
			r = rune(s[v-1])
		}
		for i := 0; i < l; i++ {
			result += string(r)
		}
	}
	return result
}

func main() {
	str := "a4bc2d5e"
	fmt.Println(StrUnpacking(str))
}
