package machox

import (
	"debug/macho"

	"github.com/ArtemKulyabin/bre/binaryx"
)

type File struct {
	*macho.File
}

func (f *File) Format() binaryx.Format {
	return binaryx.FormatMACHO
}

func (f *File) Type() binaryx.FileType {
	switch f.File.Type {
	case macho.TypeExec:
		return binaryx.Executable
	case macho.TypeDylib:
		return binaryx.DynamicLibrary
	case macho.TypeObj:
		return binaryx.Object
	}
	return 0
}

func (f *File) Arch() string {
	switch f.Cpu {
	case macho.Cpu386:
		return "386"
	case macho.CpuAmd64:
		return "amd64"
	case macho.CpuArm:
		return "arm"
	case macho.CpuPpc:
		return "ppc"
	case macho.CpuPpc64:
		return "ppc64"
	}
	return "unknown"
}

func (f *File) Os() string {
	return "darwin"
}

func (f *File) Entry() uint64 {
	return 0 // TODO: Compute entry point
}

func (f *File) Sections() []binaryx.Section {
	sections := []binaryx.Section{}
	for _, section := range f.File.Sections {
		sections = append(sections, &Section{section})
	}
	return sections
}

type Section struct {
	*macho.Section
}

func (s *Section) Name() string {
	return s.Section.Name
}

func (s *Section) Addr() uint64 {
	return s.Section.Addr
}

func (s *Section) Perm() binaryx.Perm {
	p := binaryx.Perm(0)
	if s.Section.Flags&SectFlagPureInstructions != 0 {
		p |= binaryx.PermExecute
	}
	// TODO: Add PermWrite and PermRead detection
	return p
}
