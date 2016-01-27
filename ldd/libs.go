// +build linux windows

package ldd

func getDynLibs(name string) (libs []string, err error) {
	dirs, err := GetDynLibDirs()
	if err != nil {
		return
	}

	deps := map[string]bool{}
	err = getDeps(name, dirs, deps)
	if err != nil {
		return
	}

	delete(deps, name)

	for lib := range deps {
		libs = append(libs, lib)
	}
	return
}
