package binaryx

import (
	"debug/dwarf"
	"io"
)

type File interface {
	Arch() string
	Os() string
	ImportedLibraries() ([]string, error)
	ImportedSymbols() ([]string, error)
	DWARF() (*dwarf.Data, error)
	Sections() []Section
	Type() FileType
	Format() Format
	Entry() uint64
	Close() error
}

type Section interface {
	SectionHeader
	Data() ([]byte, error)
	Open() io.ReadSeeker
	io.ReaderAt
}

type SectionHeader interface {
	Name() string
	Addr() uint64
	Perm() Perm
}

// Format specifies the file format of a binary executable.
type Format uint8

// File formats.
const (
	// FormatELF represents the Executable and Linkable Format (ELF).
	FormatELF Format = 1 + iota
	// FormatMACHO represents the Mach-O (MACHO) file format.
	FormatMACHO
	// FormatPE represents the Portable Executable (PE) file format.
	FormatPE
)

func (f Format) String() string {
	switch f {
	case FormatELF:
		return "elf"
	case FormatMACHO:
		return "macho"
	case FormatPE:
		return "pe"
	}
	return ""
}

// A Type is the binary file type, e.g. an object file, executable, or dynamic library.
type FileType uint8

const (
	Executable FileType = 1 + iota
	DynamicLibrary
	Object
)

func (t FileType) String() string {
	switch t {
	case Executable:
		return "executable"
	case DynamicLibrary:
		return "dynamic library"
	case Object:
		return "object"
	}
	return ""
}

// Perm specifies the access permissions of a segment or section in memory.
type Perm uint8

// Access permissions.
const (
	// PermExecute specifies that the memory is executable.
	PermExecute Perm = 1 << iota
	// PermWrite specifies that the memory is writeable.
	PermWrite
	// PermRead specifies that the memory is readable.
	PermRead
)
