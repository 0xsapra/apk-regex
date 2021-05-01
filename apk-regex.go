package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/0xsapra/apk-regex/extractor"
	"github.com/0xsapra/apk-regex/command/apktool"
	"github.com/ndelphit/apkurlgrep/directory"
	"os"
)

// most of code from https://github.com/ndelphit/apkurlgrep
func main() {

	parser := argparse.NewParser("apk-regex", "Apk-regex")
	regexFile := parser.String("r", "regex", &argparse.Options{Required: true, Help: "Regex config file."})
	apk := parser.String("a", "apk", &argparse.Options{Required: true, Help: "Input a path to APK file."})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(-1)
	}

	var regexes = extractor.ReadRegexFromFile(*regexFile)

	var baseApk = *apk
	var tempDir = directory.CreateTempDir()
	defer directory.RemoveTempDir(tempDir)
	apktool.RunApktool(baseApk, tempDir)
	extractor.Extract(tempDir, regexes)
}