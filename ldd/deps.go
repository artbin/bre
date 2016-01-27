package ldd

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ArtemKulyabin/bre"
)

func getDeps(name string, dirs []string, deps map[string]bool) (err error) {
	f, err := bre.Open(name)
	if err != nil {
		return
	}
	defer f.Close()

	libs, err := f.ImportedLibraries()
	if err != nil {
		return
	}

	if runtime.GOOS == "windows" {
		libsMap := map[string]bool{}
		syms, err := f.ImportedSymbols()
		if err != nil {
			return err
		}
		for _, sym := range syms {
			symInLib := strings.Split(sym, ":")
			libsMap[symInLib[1]] = true
		}
		for lib := range libsMap {
			libs = append(libs, lib)
		}
	}

	deps[name] = true

	for _, lib := range libs {
		for _, dir := range dirs {
			path := filepath.Join(dir, lib)
			if fi, err := os.Stat(path); !os.IsNotExist(err) && !fi.IsDir() {
				err = getDeps(path, dirs, deps)
				if err != nil {
					return err
				}
			}
		}
	}
	return
}
