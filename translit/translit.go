package translit

import (
	"bytes"
	"strings"
)

func EncodeTag(arrTag []string) []string {
	var ans []string
	for _, val := range arrTag {
		str, count := encode(val)
		if count > 0 {
			str = "ru--" + str
		}
		ans = append(ans, str)
	}
	return ans
}

func EncodeTitle(title string) string {
	str, _ := encode(title)
	return str
}

func encode(text string) (string, int) {
	if text == "" {
		return "", 0
	}
	text = strings.ToLower(text)
	var input = bytes.NewBufferString(text)
	var output = bytes.NewBuffer(nil)

	// Previous, next letter for special processor
	var rr string
	var ok bool

	i := 0

	for {
		r, _, err := input.ReadRune()
		if err != nil {
			break
		}

		rr, ok = encMap[string(r)]
		if ok {
			output.WriteString(rr)
			i++
			continue
		} else {
			output.WriteString(string(r))
			continue
		}
		rr, ok = encMap[string(r)]
		if ok {
			output.WriteString(rr)
			i++
		} else {
			output.WriteString(string(r))
			continue
		}
	}

	return output.String(), i
}
