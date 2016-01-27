package ldd

import "os/exec"

func GetDynLibs(name string) ([]string, error) {
	otoolPath, err := exec.LookPath("otool")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(otoolPath, "-L", name)
	return parseDynLibToolOutput(cmd)
}

func getDynLoader() string {
	return "/usr/lib/dyld"
}
