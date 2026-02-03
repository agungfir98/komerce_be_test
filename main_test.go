package main

import (
	"reflect"
	"testing"
)

func Test_sortChar(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		vowels     []rune
		consonants []rune
	}{
		{
			name:       "basic mixed string",
			s:          "hello world",
			vowels:     []rune{'e', 'o', 'o'},
			consonants: []rune{'h', 'l', 'l', 'w', 'r', 'l', 'd'},
		},
		{
			name:       "only vowels",
			s:          "aeiou",
			vowels:     []rune{'a', 'e', 'i', 'o', 'u'},
			consonants: nil,
		},
		{
			name:       "only consonants",
			s:          "bcdfg",
			vowels:     nil,
			consonants: []rune{'b', 'c', 'd', 'f', 'g'},
		},
		{
			name:       "empty string",
			s:          "",
			vowels:     nil,
			consonants: nil,
		},
		{
			name:       "spaces only",
			s:          "   ",
			vowels:     nil,
			consonants: nil,
		},
		{
			name:       "mixed with multiple spaces",
			s:          "a b c d e",
			vowels:     []rune{'a', 'e'},
			consonants: []rune{'b', 'c', 'd'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVowels, gotConsonants := sortChar(tt.s)
			if !reflect.DeepEqual(gotVowels, tt.vowels) {
				t.Errorf("sortChar() vowels = %v, want %v", gotVowels, tt.vowels)
			}
			if !reflect.DeepEqual(gotConsonants, tt.consonants) {
				t.Errorf("sortChar() consonants = %v, want %v", gotConsonants, tt.consonants)
			}
		})
	}
}

func Test_minimumBuses(t *testing.T) {
	tests := []struct {
		name    string
		n1      int
		n2      []int
		want    int
		wantErr bool
	}{
		{
			name:    "perfect fit - two families of 2 each",
			n1:      2,
			n2:      []int{2, 2},
			want:    1,
			wantErr: false,
		},
		{
			name:    "one family per bus - all size 4",
			n1:      3,
			n2:      []int{4, 4, 4},
			want:    3,
			wantErr: false,
		},
		{
			name:    "mixed sizes",
			n1:      4,
			n2:      []int{1, 2, 3, 1},
			want:    2,
			wantErr: false,
		},
		{
			name:    "small families",
			n1:      4,
			n2:      []int{1, 1, 1, 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "empty input",
			n1:      0,
			n2:      []int{},
			want:    0,
			wantErr: false,
		},
		{
			name:    "input mismatch - n1 doesn't match slice length",
			n1:      3,
			n2:      []int{1, 2},
			want:    -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumBuses(tt.n1, tt.n2)
			if tt.wantErr {
				if got > 0 {
					t.Errorf("minimumBuses() = %d, expected error", got)
				}
			} else {
				if got != tt.want {
					t.Errorf("minimumBuses() = %d, want %d", got, tt.want)
				}
			}
		})
	}
}
