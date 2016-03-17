package ldd

import "os/exec"

func GetDynLibs(path string) (libs []string, err error) {
	deps := map[string]bool{}
	var getdeps func(path string) error
	getdeps = func(path string) error {
		libs, err := getDynLibsOtool(path)
		if err != nil {
			return err
		}
		for _, lib := range libs {
			if _, ok := deps[lib]; !ok {
				deps[lib] = true
				return getdeps(lib)
			}
		}
		return nil
	}
	getdeps(path)
	for lib := range deps {
		libs = append(libs, lib)
	}
	return
}

func getDynLibsOtool(name string) ([]string, error) {
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
