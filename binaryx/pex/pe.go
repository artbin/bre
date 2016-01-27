package pex

import (
	"debug/pe"

	"github.com/ArtemKulyabin/bre/binaryx"
)

type File struct {
	*pe.File
}

func (f *File) Format() binaryx.Format {
	return binaryx.FormatPE
}

func (f *File) Type() binaryx.FileType {
	if f.File.Characteristics&uint16(FlagDLL) != 0 {
		return binaryx.DynamicLibrary
	}
	if f.File.Characteristics&uint16(FlagExecutable) != 0 {
		return binaryx.Executable
	}
	return binaryx.Object
}

func (f *File) Arch() string {
	switch f.Machine {
	case pe.IMAGE_FILE_MACHINE_I386:
		return "386"
	case pe.IMAGE_FILE_MACHINE_AMD64:
		return "amd64"
	}
	return "unknown"
}

func (f *File) Os() string {
	return "windows"
}

func (f *File) Entry() uint64 {
	entry := uint64(0)
	if opthdr, ok := f.OptionalHeader.(*pe.OptionalHeader64); ok {
		entry = uint64(opthdr.AddressOfEntryPoint)
	} else {
		if opthdr, ok := f.OptionalHeader.(*pe.OptionalHeader32); ok {
			entry = uint64(opthdr.AddressOfEntryPoint)
		}
	}
	return f.getBase() + entry
}

func (f *File) Sections() []binaryx.Section {
	sections := []binaryx.Section{}
	for _, section := range f.File.Sections {
		sections = append(sections, &Section{section, f})
	}
	return sections
}

type Section struct {
	*pe.Section
	*File
}

func (s *Section) Name() string {
	return s.Section.Name
}

func (s *Section) Addr() uint64 {
	return s.getBase() + uint64(s.Section.VirtualAddress)
}

func (s *Section) Perm() binaryx.Perm {
	p := binaryx.Perm(0)
	if s.Section.Characteristics&uint32(SectFlagMemExec) != 0 {
		p |= binaryx.PermExecute
	}
	if s.Section.Characteristics&uint32(SectFlagMemWrite) != 0 {
		p |= binaryx.PermWrite
	}
	if s.Section.Characteristics&uint32(SectFlagMemRead) != 0 {
		p |= binaryx.PermRead
	}
	return p
}

func (f *File) getBase() uint64 {
	base := uint64(0)
	if opthdr, ok := f.OptionalHeader.(*pe.OptionalHeader64); ok {
		base = opthdr.ImageBase
	} else {
		if opthdr, ok := f.OptionalHeader.(*pe.OptionalHeader32); ok {
			base = uint64(opthdr.ImageBase)
		}
	}
	return base
}
