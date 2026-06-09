package iteration

import "strings"

func Repeat(character string, ntimes int) string {
	var repeated strings.Builder
	for range ntimes {
		repeated .WriteString(character)
	}
	return repeated.String() 
}