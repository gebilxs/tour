package word

import (
	"strings"
	"unicode"
)

//全部转换为大写

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线转换成大写的驼峰

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//下划线转换成小写的驼峰 对前面的函数进行应用和处理即可

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰转下划线 rune int32
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))

	}
	return string(output)
}
