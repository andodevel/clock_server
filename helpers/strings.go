package helpers

import (
	"fmt"
	"strings"
)

// ToString ...Legendary ToString() function
func ToString(any interface{}) string {
	return fmt.Sprintf("%v", any)
}

// IsBlank ...
func IsBlank(value string) bool {
	return strings.Trim(value, " ") == ""
}

// IsEmpty ...
func IsEmpty(value string) bool {
	return "" == value
}
