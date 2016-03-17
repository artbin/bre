package ldd

import "os/exec"

func GetDynLibs(name string) ([]string, error) {
	lddPath, err := exec.LookPath("ldd")
	if err != nil {
		return getDynLibs(name)
	}
	cmd := exec.Command(lddPath, name)
	return parseDynLibToolOutput(cmd)
}

func getDynLoader() string {
	return "/lib/x86_64-linux-gnu/ld-2.19.so"
}
