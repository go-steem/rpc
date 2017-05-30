package translit

import (
	"bytes"
	"strings"
)

func Transliteration(text string) string {
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

		rr, ok = extMap[string(r)]
		if ok {
			output.WriteString(rr)
			continue
		}
		rr, ok = extMap[string(r)]
		if ok {
			output.WriteString(rr)
		}
	}
	return output.String()
}
