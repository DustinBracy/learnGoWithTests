package property

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{19, "XIX"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{44, "XLIV"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2024, "MMXXIV"},
	{2023, "MMXXIII"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d is %q", c.Arabic, c.Roman), func(t *testing.T) {
			got := ConvertToRoman(c.Arabic)
			if got != c.Roman {
				t.Errorf("got %q want %q", got, c.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:10] {
		t.Run(fmt.Sprintf("%q is %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}
		})
	}
}

func BenchmarkConvertToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			ConvertToRoman(c.Arabic)
		}
	}
}
func BenchmarkConvertToRomanRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			ConvertToRomanRec(c.Arabic)
		}
	}
}
func BenchmarkConvertToArabic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			ConvertToArabic(c.Roman)
		}
	}
}

func BenchmarkConvertToArabicRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			ConvertToArabicRec(c.Roman)
		}
	}
}
