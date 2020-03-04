package helpers

// TODO: Check wether these functions useful
import (
	"html/template"
	"reflect"
	"strings"
	"sync/atomic"
)

var templateOps uint64

type (
	// Map ...Alias for map[string]interface{}
	Map map[string]interface{}
)

// ParseHTMLTemplateFile ...
func ParseHTMLTemplateFile(templateName string, filename string, vars Map) (string, error) {
	text, err := LoadFile(filename)
	if err != nil {
		return "", err
	}

	return ParseHTMLTemplate(templateName, text, vars)
}

// ParseHTMLTemplate ...
func ParseHTMLTemplate(templateName string, text string, vars Map) (string, error) {
	if IsBlank(templateName) {
		atomic.AddUint64(&templateOps, 1)
		templateName = ToString(templateOps)
	}

	page := template.Must(template.New(templateName).Parse(text))
	bld := new(strings.Builder)
	err := page.Execute(bld, vars)
	return bld.String(), err
}

// AssignValueOrEmptyString ...
func AssignValueOrEmptyString(value interface{}) string {
	if IsZeroOfUnderlyingType(value) {
		return ""
	}
	return value.(string)
}

// IsZeroOfUnderlyingType ...check zero underlying
func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

// Isset ...check slices has specific index
func Isset(arr []interface{}, index int) bool {
	return (len(arr) > index)
}
