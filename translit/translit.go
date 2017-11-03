package translit

import (
	"bytes"
	"log"
	"regexp"
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
	var str string
	reg, err := regexp.Compile("[^a-zA-Z0-9а-яА-Я.,]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(title, "-")
	s1, _ := encode(processedString)
	s2 := strings.Replace(s1, ".", "", -1)
	s3 := strings.Split(s2, "")
	if s3[0] == "-" {
		str = strings.Join(s3[1:], "")
	} else {
		str = strings.Join(s3, "")
	}
	return str
}

func Tsenc(text string) (string, int) {
	return encode(text)
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
