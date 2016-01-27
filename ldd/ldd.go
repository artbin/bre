// +build !linux,!windows,!darwin

package ldd

import "os/exec"

func GetDynLibs(name string) ([]string, error) {
	lddPath, err := exec.LookPath("ldd")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(lddPath, name)
	return parseDynLibToolOutput(cmd)
}
