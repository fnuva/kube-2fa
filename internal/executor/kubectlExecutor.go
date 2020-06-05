package executor

import (
	"bytes"
	"fmt"
	"kube-2fa/internal/helper"
	"os"
	"os/exec"
	"strings"
)

func Execute(args []string) error {
	cmd := exec.Command("kubectl", args...)
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = &errBuf
	cmd.Stdout = &outBuf
	err := cmd.Run()
	if err != nil {
		fmt.Println(errBuf.String())
		return err
	}
	fmt.Println(outBuf.String())
	return nil
}
func Apply(fileName string, customId string) error {
	content, _ := helper.Read(fileName)
	contentString := string(content)
	result := strings.ReplaceAll(contentString, "%AUTH_ID%", customId)
	newFile, err := helper.Write([]byte(result))
	if err != nil {
		return err
	}
	return Execute([]string{
		"apply", "-f", newFile})
}
