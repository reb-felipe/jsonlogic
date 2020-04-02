package jsonlogic

import (
	"bytes"
	"strings"
)

func substr(values interface{}) interface{} {
	parsed := values.([]interface{})

	runes := []rune(toString(parsed[0]))

	from := int(toNumber(parsed[1]))
	length := len(runes)

	if from < 0 {
		from = length + from
	}

	if len(parsed) == 3 {
		length = int(toNumber(parsed[2]))
	}

	var to int
	if length < 0 {
		length = len(runes) + length
		to = length
	} else {
		to = from + length
	}

	if from < 0 {
		from, to = to, len(runes)-from
	}

	if to > len(runes) {
		to = len(runes)
	}

	if from > len(runes) {
		from = len(runes)
	}

	return string(runes[from:to])
}

func concat(values interface{}) interface{} {
	if isString(values) {
		return values
	}

	var s bytes.Buffer
	for _, text := range values.([]interface{}) {
		s.WriteString(toString(text))
	}

	return strings.TrimSpace(s.String())
}