package elfx

import (
	"debug/elf"

	"github.com/ArtemKulyabin/bre/binaryx"
)

type File struct {
	*elf.File
}

func (f *File) Format() binaryx.Format {
	return binaryx.FormatELF
}

func (f *File) Type() binaryx.FileType {
	switch f.File.Type {
	case elf.ET_EXEC:
		return binaryx.Executable
	case elf.ET_DYN:
		return binaryx.DynamicLibrary
	case elf.ET_REL:
		return binaryx.Object
	}
	return 0
}

func (f *File) Arch() string {
	switch f.Machine {
	case elf.EM_386:
		return "386"
	case elf.EM_X86_64:
		return "amd64"
	case elf.EM_ARM:
		return "arm"
	case elf.EM_AARCH64:
		return "arm64"
	case elf.EM_PPC:
		return "ppc"
	case elf.EM_PPC64:
		if f.File.Data == elf.ELFDATA2LSB {
			return "ppc64le"
		}
		return "ppc64"
	}
	return "unknown"
}

func (f *File) Os() string {
	switch f.OSABI {
	case elf.ELFOSABI_NONE:
		return "sysv"
	case elf.ELFOSABI_LINUX:
		return "linux"
	case elf.ELFOSABI_NETBSD:
		return "netbsd"
	case elf.ELFOSABI_FREEBSD:
		return "freebsd"
	case elf.ELFOSABI_OPENBSD:
		return "openbsd"
	case elf.ELFOSABI_SOLARIS:
		return "solaris"
	}
	return "unknown"
}

func (f *File) Entry() uint64 {
	return f.File.Entry
}

func (f *File) ImportedSymbols() ([]string, error) {
	slice := []string{}
	symbols, err := f.File.ImportedSymbols()
	if err != nil {
		return nil, err
	}
	for _, symbol := range symbols {
		slice = append(slice, symbol.Name)
	}
	return slice, nil
}

func (f *File) Sections() []binaryx.Section {
	sections := []binaryx.Section{}
	for _, section := range f.File.Sections {
		sections = append(sections, &Section{section})
	}
	return sections
}

type Section struct {
	*elf.Section
}

func (s *Section) Name() string {
	return s.Section.Name
}

func (s *Section) Addr() uint64 {
	return s.Section.Addr
}

func (s *Section) Perm() binaryx.Perm {
	p := binaryx.Perm(0)
	if s.Section.Flags&elf.SHF_EXECINSTR != 0 {
		p |= binaryx.PermExecute
	}
	if s.Section.Flags&elf.SHF_WRITE != 0 {
		p |= binaryx.PermWrite
	}
	if s.Section.Flags&elf.SHF_ALLOC != 0 {
		p |= binaryx.PermRead
	}
	return p
}
