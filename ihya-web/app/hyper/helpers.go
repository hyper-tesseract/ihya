package hyper

import (
	"fmt"
	"strings"
)

func FormatString(txt string, val string) string {
	if txt != "" && val != "" {
		return fmt.Sprintf(txt, val)
	}
	return ""
}

func SplitString(val string, splitter string) []string {
	if val == "" {
		return make([]string, 0)
	}
	return strings.Split(val, splitter)
}
