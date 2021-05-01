package apktool

import (
	"os"
	"os/exec"
	"fmt"
)

func RunApktool(apk string, tempDir string) string {
	path, err := os.Getwd()
	cmd := exec.Command("java", "-jar", string(path) + "/bin/apktool_2.5.0.jar","d", apk, "-o", tempDir, "-fq")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return string(output)
}