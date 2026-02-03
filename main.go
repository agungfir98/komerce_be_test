package main

import (
	"slices"
	"strings"
)

func main() {
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

	for i := range n2 {
		if n2[i] <= 0 {
			return -2
		}
		if n2[i] > maxPassenger {
			return -3
		}
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
