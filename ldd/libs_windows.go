package ldd

func GetDynLibs(name string) ([]string, error) {
	return getDynLibs(name)
}

func getDynLoader() string {
	return ""
}
