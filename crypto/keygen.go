package crypto

import (
	"fmt"
	"os"
	"os/exec"
)

func cpabeKeygen(pubKeyFile, mskFile, attrFile, keyFile string) error {
	var out bytes.Buffer
	cmd := exec.Command("cpabe-keygen", "-o", keyFile, pubKeyFile, mskFile, attrFile)
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run cpabe-keygen command: %v, output: %s", err, out.String())
	}

	return nil
}
