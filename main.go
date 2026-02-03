package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	sample := "hallo dunia"
	vow, con := sortChar(sample)

	fmt.Printf("vowel: %s, consonants: %s\n", string(vow), string(con))

	n := minimumBuses(5, []int{1, 2, 4, 3, 3})
	if n < 0 {
		fmt.Printf("Input must be equal with count family\n")
	} else {
		fmt.Printf("minimumBuses %d\n", n)
	}
}

func sortChar(s string) ([]rune, []rune) {
	s = strings.ToLower(s)
	var consonants, vowels []rune

	for _, r := range s {
		if r == ' ' {
			continue
		}

		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels = append(vowels, r)
		default:
			consonants = append(consonants, r)
		}
	}

	return vowels, consonants
}

func minimumBuses(n1 int, n2 []int) int {
	maxPassenger := 4
	nBuses := 0

	if n1 != len(n2) {
		return -1
	}

	slices.Sort(n2)

	occupied := 0
	for i := range n2 {
		if (occupied + n2[i]) > maxPassenger {
			nBuses += 1
			occupied = n2[i]
		} else {
			occupied += n2[i]
		}
	}

	if occupied > 0 {
		nBuses += 1
	}

	return nBuses
}
