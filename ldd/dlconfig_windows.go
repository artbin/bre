package ldd

import (
	"os"
	"strings"
	"syscall"

	"github.com/ArtemKulyabin/yax/osx"
	"github.com/ArtemKulyabin/yax/syscallx"
)

// Search Path Used by Windows to Locate a DLL - https://msdn.microsoft.com/en-us/library/7d83bc18.aspx
func GetDynLibDirs() (dirs []string, err error) {
	// The directory where the executable module for the current process is located.
	path, err := osx.Executable()
	if err != nil {
		return
	}
	dirs = append(dirs, path)
	// The current directory.
	path, err = os.Getwd()
	if err != nil {
		return
	}
	dirs = append(dirs, path)
	// The Windows system directory. The GetSystemDirectory function retrieves the path of this directory.
	buf := make([]uint16, syscall.MAX_PATH)
	_, err = syscallx.GetSystemDirectory(&buf[0], syscall.MAX_PATH)
	if err != nil {
		return
	}
	path = syscall.UTF16ToString(buf)
	dirs = append(dirs, path)
	// The Windows directory. The GetWindowsDirectory function retrieves the path of this directory.
	_, err = syscallx.GetWindowsDirectory(&buf[0], syscall.MAX_PATH)
	if err != nil {
		return
	}
	path = syscall.UTF16ToString(buf)
	dirs = append(dirs, path)
	// The directories listed in the PATH environment variable.
	path = os.Getenv("PATH")
	dirs = append(dirs, strings.Split(path, string(os.PathListSeparator))...)
	return
}
