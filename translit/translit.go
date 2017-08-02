package translit

import (
	"bytes"
	"strings"
)

func EncodeTags(tag []string) []string {
	var arrEncTag []string
	for _, val := range tag {
		str, count := encode(val)
		if count > 0 {
			str = "ru--" + str
		}
		arrEncTag = append(arrEncTag, str)
	}
	return arrEncTag
}

func EncodeTag(tag string) string {
	str, count := encode(tag)
	if count > 0 {
		str = "ru--" + str
	}
	return str
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
