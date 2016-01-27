package bre

import (
	"debug/dwarf"
	"debug/elf"
	"debug/macho"
	"debug/pe"
	"io"

	"github.com/ArtemKulyabin/bre/binaryx"
	"github.com/ArtemKulyabin/bre/binaryx/elfx"
	"github.com/ArtemKulyabin/bre/binaryx/machox"
	"github.com/ArtemKulyabin/bre/binaryx/pex"
)

func NewFile(r io.ReaderAt) (binaryx.File, error) {
	elfBinary, err := elf.NewFile(r)
	if err == nil {
		return newFile(&elfx.File{elfBinary})
	}
	machoBinary, err := macho.NewFile(r)
	if err == nil {
		return newFile(&machox.File{machoBinary})
	}
	peBinary, err := pe.NewFile(r)
	if err == nil {
		return newFile(&pex.File{peBinary})
	}
	return nil, err
}

func Open(name string) (binaryx.File, error) {
	elfBinary, err := elf.Open(name)
	if err == nil {
		return newFile(&elfx.File{elfBinary})
	}
	machoBinary, err := macho.Open(name)
	if err == nil {
		return newFile(&machox.File{machoBinary})
	}
	peBinary, err := pe.Open(name)
	if err == nil {
		return newFile(&pex.File{peBinary})
	}
	return nil, err
}

type file struct {
	arch              string
	os                string
	importedLibraries []string
	importedSymbols   []string
	dWARF             *dwarf.Data
	sections          []binaryx.Section
	ftype             binaryx.FileType
	format            binaryx.Format
	entry             uint64

	binaryFile binaryx.File
}

func newFile(binaryFile binaryx.File) (binaryx.File, error) {
	f := &file{}
	f.arch = binaryFile.Arch()
	f.os = binaryFile.Os()
	f.sections = binaryFile.Sections()
	f.ftype = binaryFile.Type()
	f.format = binaryFile.Format()
	f.entry = binaryFile.Entry()
	f.binaryFile = binaryFile
	return f, nil
}

func (f *file) Arch() string {
	return f.arch
}

func (f *file) Os() string {
	return f.os
}

func (f *file) ImportedLibraries() ([]string, error) {
	var err error
	if f.importedLibraries == nil {
		f.importedLibraries, err = f.binaryFile.ImportedLibraries()
		if err != nil {
			return nil, err
		}
	}
	return f.importedLibraries, nil
}

func (f *file) ImportedSymbols() ([]string, error) {
	var err error
	if f.importedSymbols == nil {
		f.importedSymbols, err = f.binaryFile.ImportedSymbols()
		if err != nil {
			return nil, err
		}
	}
	return f.importedSymbols, nil
}

func (f *file) DWARF() (*dwarf.Data, error) {
	var err error
	if f.dWARF == nil {
		f.dWARF, err = f.binaryFile.DWARF()
		if err != nil {
			return nil, err
		}
	}
	return f.dWARF, nil
}

func (f *file) Sections() []binaryx.Section {
	return f.sections
}

func (f *file) Type() binaryx.FileType {
	return f.ftype
}

func (f *file) Format() binaryx.Format {
	return f.format
}

func (f *file) Entry() uint64 {
	return f.entry
}

func (f *file) Close() error {
	return f.binaryFile.Close()
}
