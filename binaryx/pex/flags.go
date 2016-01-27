package pex

// Flag is a bitfield which specifies the characteristics of an executable.
type Flag uint16

// Executable characteristics.
const (
	// FlagNoReloc indicates that the file contains no relocation information.
	FlagNoReloc Flag = 1 << iota
	// FlagExecutable indicates that the file is executable.
	FlagExecutable
	// FlagNoLineNums indicates that the file contains no line numbers.
	FlagNoLineNums
	// FlagNoSymTbl indicates that the file contains no symbol table.
	FlagNoSymTbl
	_ // obsolete.
	// FlagLargeAddr indicates that the application can handle addresses larger
	// than 2 GB.
	FlagLargeAddr
	_ // obsolete.
	_ // obsolete.
	// Flag32bit indicates that the computer supports 32-bit words.
	Flag32bit
	// FlagNoDebug indicates that the file contains no debugging information. It
	// may be present in a separate file.
	FlagNoDebug
	// FlagUSBCopyToSwap copies the file from a usb device to the swap before
	// running it.
	FlagUSBCopyToSwap
	// FlagNetCopyToSwap copies the file from the network to the swap before
	// running it.
	FlagNetCopyToSwap
	// FlagSystem indicates that the file is a system file.
	FlagSystem
	// FlagDLL indicates that the file is a dynamic link library (DLL).
	FlagDLL
	// FlagUniprocessor indicates that the file should only be run on a
	// uniprocessor computer.
	FlagUniprocessor
)

// SectFlag is a bitfield which specifies the characteristics of a section.
type SectFlag uint32

const (
	// SectFlagCode indicates that the section contains executable code.
	SectFlagCode SectFlag = 0x00000020
	// SectFlagData indicates that the section contains initialized data.
	SectFlagData SectFlag = 0x00000040
	// SectFlagBSS indicates that the section contains uninitialized data.
	SectFlagBSS SectFlag = 0x00000080
	// SectFlagLinkInfo indicates that the section contains comments or other
	// information. Only valid for object files.
	SectFlagLinkInfo SectFlag = 0x00000200
	// SectFlagLinkRemove indicates that the section will not become part of the
	// image. Only valid for object files.
	SectFlagLinkRemove SectFlag = 0x00000800
	// SectFlagLinkCOMDAT indicates that the section contains COMDAT data. Only
	// valid for object files.
	SectFlagLinkCOMDAT SectFlag = 0x00001000
	// SectFlagDeferSpecExc resets speculative exception handling bits in the TLB
	// entries for this section.
	SectFlagDeferSpecExc SectFlag = 0x00004000
	// SectFlagGPRef indicates that the section contains data referenced through
	// the global pointer.
	SectFlagGPRef SectFlag = 0x00008000
	// SectFlagObjAlign1 aligns data on a 1-byte boundary. Only valid for object
	// files.
	SectFlagObjAlign1 SectFlag = 0x00100000
	// SectFlagObjAlign2 aligns data on a 2-byte boundary. Only valid for object
	// files.
	SectFlagObjAlign2 SectFlag = 0x00200000
	// SectFlagObjAlign4 aligns data on a 4-byte boundary. Only valid for object
	// files.
	SectFlagObjAlign4 SectFlag = 0x00300000
	// SectFlagObjAlign8 aligns data on a 8-byte boundary. Only valid for object
	// files.
	SectFlagObjAlign8 SectFlag = 0x00400000
	// SectFlagObjAlign16 aligns data on a 16-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign16 SectFlag = 0x00500000
	// SectFlagObjAlign32 aligns data on a 32-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign32 SectFlag = 0x00600000
	// SectFlagObjAlign64 aligns data on a 64-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign64 SectFlag = 0x00700000
	// SectFlagObjAlign128 aligns data on a 128-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign128 SectFlag = 0x00800000
	// SectFlagObjAlign256 aligns data on a 256-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign256 SectFlag = 0x00900000
	// SectFlagObjAlign512 aligns data on a 512-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign512 SectFlag = 0x00A00000
	// SectFlagObjAlign1024 aligns data on a 1024-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign1024 SectFlag = 0x00B00000
	// SectFlagObjAlign2048 aligns data on a 2048-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign2048 SectFlag = 0x00C00000
	// SectFlagObjAlign4096 aligns data on a 4096-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign4096 SectFlag = 0x00D00000
	// SectFlagObjAlign8192 aligns data on a 8192-byte boundary. Only valid for
	// object files.
	SectFlagObjAlign8192 SectFlag = 0x00E00000
	// SectFlagRelocsOverflow indicates that there are more relocations than can
	// be represented by the 16-bit value in the section header. If the value of
	// Relocs in the section header is 0xFFFF, the actual relocation count is
	// stored in the RelAddr field of the first relocation.
	SectFlagRelocsOverflow SectFlag = 0x01000000
	// SectFlagMemDiscard indicates that the section memory can be discarded as
	// needed.
	SectFlagMemDiscard SectFlag = 0x02000000
	// SectFlagMemNoCache indicates that the section memory cannot be cached.
	SectFlagMemNoCache SectFlag = 0x04000000
	// SectFlagMemNoPage indicates that the section memory cannot be paged.
	SectFlagMemNoPage SectFlag = 0x08000000
	// SectFlagMemShared indicates that the section memory can be shared.
	SectFlagMemShared SectFlag = 0x10000000
	// SectFlagMemExec indicates that the section memory can be executed.
	SectFlagMemExec SectFlag = 0x20000000
	// SectFlagMemRead indicates that the section memory can be read.
	SectFlagMemRead SectFlag = 0x40000000
	// SectFlagMemWrite indicates that the section memory can be written to.
	SectFlagMemWrite SectFlag = 0x80000000
)
