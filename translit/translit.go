package translit

import (
	"bytes"
	"strings"
)

func Encode(text string) string {
	if text == "" {
		return ""
	}
	text = strings.ToLower(text)
	var input = bytes.NewBufferString(text)
	var output = bytes.NewBuffer(nil)

	// Previous, next letter for special processor
	var rr string
	var ok bool

	for {
		r, _, err := input.ReadRune()
		if err != nil {
			break
		}

		rr, ok = encMap[string(r)]
		if ok {
			output.WriteString(rr)
			continue
		} else {
			output.WriteString(string(r))
			continue
		}
		rr, ok = encMap[string(r)]
		if ok {
			output.WriteString(rr)
		} else {
			output.WriteString(string(r))
			continue
		}
	}

	return output.String()
}
