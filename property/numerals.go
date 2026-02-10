package property

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(n uint16) string {
	var result strings.Builder

	for _, numeral := range romanNumerals {
		for n >= uint16(numeral.Value) {
			result.WriteString(numeral.Symbol)
			n -= uint16(numeral.Value)
		}
	}

	return result.String()
}

func ConvertToRomanRec(n uint16) string {
	if n == 0 {
		return ""
	}

	for _, numeral := range romanNumerals {
		if n >= uint16(numeral.Value) {
			return numeral.Symbol + ConvertToRomanRec(n-uint16(numeral.Value))
		}
	}
	return ""
}

func ConvertToArabic(roman string) uint16 {
	arabic := uint16(0)

	for _, numeral := range romanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += uint16(numeral.Value)
			roman = strings.TrimPrefix(roman, numeral.Symbol)

		}
	}
	return arabic
}

func ConvertToArabicRec(roman string) uint16 {
	if roman == "" {
		return 0
	}

	for _, numeral := range romanNumerals {
		if strings.HasPrefix(roman, numeral.Symbol) {
			return uint16(numeral.Value) + ConvertToArabicRec(strings.TrimPrefix(roman, numeral.Symbol))
		}
	}
	return 0
}
