package ldd

// https://www.freebsd.org/cgi/man.cgi?query=ldconfig

var ConfigFiles = []string{
	// Standard hints file for the a.out dynamic linker.
	"/var/run/ld.so.hints",
	// Standard hints file for the ELF dynamic linker.
	"/var/run/ld-elf.so.hints",
	// Conventional configuration file containing directory names for invocations
	// with -aout.
	"/etc/ld.so.conf",
	// Conventional configuration file containing directory names for invocations
	// with -elf.
	"/etc/ld-elf.so.conf",
	// Conventional configuration files containing directory names for invocations
	// with -32.
	"/var/run/ld-elf32.so.hints",
	"/var/run/ld32.so.hints",
}
