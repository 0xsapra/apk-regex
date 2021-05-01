package extractor

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Regex struct {
    Name   string `json:"name"`
    Regex   string `json:"regex"`
}

func ReadRegexFromFile(filePath string) []Regex {

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Print((err))
		os.Exit(-1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var regexes []Regex;
	json.Unmarshal([]byte(byteValue), &regexes)
	return regexes;
}