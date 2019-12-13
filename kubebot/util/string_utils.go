package util

import "strings"

func StripUserId(value string) string {
	var replacer = strings.NewReplacer("<", "", ">", "")
	return replacer.Replace(value)
}
