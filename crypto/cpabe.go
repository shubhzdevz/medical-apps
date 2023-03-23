package crypto

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func cpabeEncrypt(data []byte, policy string) ([]byte, error) {
	tempFile, err := ioutil.TempFile("", "cpabe")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(data); err != nil {
		return nil, fmt.Errorf("failed to write data to temporary file: %v", err)
	}

	var out bytes.Buffer
	cmd := exec.Command("cpabe", "-e", policy, tempFile.Name())
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run cpabe command: %v, output: %s", err, out.String())
	}

	encryptedData, err := ioutil.ReadFile(tempFile.Name() + ".cpabe")
	if err != nil {
		return nil, fmt.Errorf("failed to read encrypted data file: %v", err)
	}

	return encryptedData, nil
}

func cpabeDecrypt(data []byte, key string) ([]byte, error) {
	tempFile, err := ioutil.TempFile("", "cpabe")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(data); err != nil {
		return nil, fmt.Errorf("failed to write data to temporary file: %v", err)
	}

	var out bytes.Buffer
	cmd := exec.Command("cpabe", "-d", key, tempFile.Name()+".cpabe")
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run cpabe command: %v, output: %s", err, out.String())
	}

	decryptedData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read decrypted data file: %v", err)
	}

	return decryptedData, nil
}
