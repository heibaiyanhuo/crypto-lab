package utils

import "io/ioutil"

func GetContentFromFile(filename string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return data
}
