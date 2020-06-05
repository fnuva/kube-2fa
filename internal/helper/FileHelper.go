package helper

import (
	"io/ioutil"
	"log"
	"os"
)

func Write(content []byte) (string, error) {
	file, err := ioutil.TempFile("/tmp", "kube-2fa-deployment")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	_, err = file.Write(content)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}
func Read(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return content, nil
}
