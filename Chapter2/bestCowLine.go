package main

import (
	"fmt"
)

func solve4(s string) {
	a, b := 0, len(s)-1
	t := ""
	for a <= b {
		var left bool = false

		for i := 0; a+i <= b; i++ {
			if s[a+i] < s[b-i] {
				left = true
				break
			} else if s[a+i] > s[b-i] {
				left = false
				break
			}
		}

		if left {
			t += string(s[0])
			s = s[1:]

		} else {
			t += string(s[b])
			s = s[:b]
		}

		b --

	}

	fmt.Println(t)
}

func main() {
	var s string
	fmt.Println("Please enter a string:")
	fmt.Scan(&s)

	solve4(s)
}
