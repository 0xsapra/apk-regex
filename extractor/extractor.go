package extractor

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var founds []string
var regexesAll []Regex

// regexp from https://github.com/GerbenJavado/LinkFinder
const regexUrlsString = `(?:"|')(((?:[a-zA-Z]{1,10}://|//)[^"'/]{1,}\.[a-zA-Z]{2,}[^"']{0,})|((?:/|\.\./|\./)[^"'><,;| *()(%%$^/\\\[\]][^"'><,;|()]{1,})|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{1,}\.(?:[a-zA-Z]{1,4}|action)(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{3,}(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-]{1,}\.(?:php|asp|aspx|jsp|json|action|html|js|txt|xml)(?:[\?|#][^"|']{0,}|)))(?:"|')`

var regexpUrls = regexp.MustCompile(regexUrlsString)


func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func extractTextFromFile(path string) error {
	var textBytes, er = ioutil.ReadFile(path)
	if er != nil {
		fmt.Print(textBytes)
		panic(er)
	}

	for _, regexStruct := range regexesAll {
		
		regexCompiled := regexp.MustCompile(regexStruct.Regex)

		var indexes = regexCompiled.FindAllIndex(textBytes, -1)

		if len(indexes) != 0 {
			for _, k := range indexes {
				var foundStart = k[0]
				var foundEnd = k[1]
				var link = string(textBytes[foundStart:foundEnd])
				fmt.Println("Found ", regexStruct.Name, " -> ", link, " in ", path)
			}
		}
	}

	return nil
	
}

func doHashWalk(dirPath string) error {
	fullPath, err := filepath.Abs(dirPath)

	if err != nil {
		return err
	}

	callback := func(path string, fi os.FileInfo, err error) error {
		return hashFile(path, fi, err)
	}

	return filepath.Walk(fullPath, callback)
}

func hashFile(path string, fileInfo os.FileInfo, err error) error {
	var fileName = fileInfo.Name()

	if fileInfo.IsDir() {
		return nil
	}

	if SkipExtension(fileName) {
		return nil
	}

	if err != nil {
		return err
	}

	extractTextFromFile(path)

	return nil
}

func sortUrls(urls []string) ([]string, []string) {

	urls = unique(urls)

	var sortedUrls []string
	var sortedPaths []string

	for i := 1; i < len(urls); i++ {

		urls[i] = strings.ReplaceAll(urls[i], "'", "")
		urls[i] = strings.ReplaceAll(urls[i], "\"", "")

		if len(urls[i]) < 5 {
			continue
		}

		if urls[i][:4] == "http" || urls[i][:5] == "https" {
			sortedUrls = append(sortedUrls, urls[i])
			continue
		} else {
			sortedPaths = append(sortedPaths, urls[i])
		}

	}
	return sortedUrls, sortedPaths
}

func Extract(tempDir string, regexes []Regex) {
	
	regexesAll = regexes
	doHashWalk(tempDir)
}
