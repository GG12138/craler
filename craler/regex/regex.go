package regex

import (
	"fmt"
	"regexp"
)

func Regex(s string) {
	re := regexp.MustCompile(s)
	fmt.Println(re)
}
