package jpera

import (
	"regexp"
	"strconv"
)

var (
	startYear = map[string]int{
		"明治": 1868,
		"大正": 1912,
		"昭和": 1926,
		"平成": 1989,
		"令和": 2019,
	}
	regex = regexp.MustCompile(`(令和|平成|昭和|大正|明治)\s*(元|\d+)`)
)

// ReplaceJpEra は文字列中の元号表記を西暦に直した文字列を返します
func ReplaceJpEra(text string) string {
	return regex.ReplaceAllStringFunc(text, replaceJpEraCallback)
}

func replaceJpEraCallback(s string) string {
	sub := regex.FindStringSubmatch(s)
	year := 1
	if sub[2] != "元" {
		y, _ := strconv.Atoi(sub[2])
		year = y
	}
	return strconv.Itoa(startYear[sub[1]] + year - 1)
}
