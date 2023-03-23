package crypto

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func cpabeSetup() error {
	var out bytes.Buffer
	cmd := exec.Command("cpabe-setup")
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run cpabe-setup command: %v, output: %s", err, out.String())
	}

	return nil
