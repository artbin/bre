package ldd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func parseDynLibToolOutput(cmd *exec.Cmd) (libs []string, err error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	// Validate the input
	for scanner.Scan() {
		lib := scanner.Text()
		if _, err := os.Stat(lib); !os.IsNotExist(err) {
			libs = append(libs, lib)
		}
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("Invalid input: %s", err)
		return
	}
	return
}
