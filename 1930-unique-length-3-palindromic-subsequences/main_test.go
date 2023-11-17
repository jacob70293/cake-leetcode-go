package _1930_unique_length_3_palindromic_subsequences

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{""}, 0},
		{"aba", args{"aba"}, 1},
		{"abca", args{"abca"}, 2},
		{"aabca", args{"aabca"}, 3},
		{"adc", args{"adc"}, 0},
		{"adc", args{"adc"}, 0},
		{"bbcbaba", args{"bbcbaba"}, 4},
		{"case_4", args{"tlpjzdmtwderpkpmgoyrcxttiheassztncqvnfjeyxxp"}, 161},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.args.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
