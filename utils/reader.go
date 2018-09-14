package utils

import (
	"bytes"
	"io/ioutil"
)

func GetContentFromFile(filename string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return data
}

func GetUpperCaseContentFromFile(filename string) []byte {
	content := GetContentFromFile(filename)
	var buf bytes.Buffer
	var messageByte byte
	for i := 0; i < len(content); i++ {
		messageByte = content[i]
		if 97 <= messageByte && messageByte < 123 {
			messageByte -= 32
		}
		if 65 <= messageByte && messageByte < 91 {
			buf.WriteByte(messageByte)
		}
	}
	return buf.Bytes()
}

//func GetUpperCaseStringFromFile(filename string) string {
//	contentString := string(GetContentFromFile(filename))
//	contentString = strings.TrimSpace(contentString)
//	reg, _ := regexp.Compile("[^a-zA-Z]+")
//	contentString = reg.ReplaceAllString(contentString, "")
//	return strings.ToUpper(contentString)
//}

