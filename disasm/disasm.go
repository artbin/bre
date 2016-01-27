package disasm

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/bnagy/gapstone"
)

func Disasm(code io.Reader, asm io.Writer, address uint64, arch string) (err error) {

	gapArch, gapMode := int(0), uint(0)

	switch arch {
	case "386":
		gapArch = gapstone.CS_ARCH_X86
		gapMode = gapstone.CS_MODE_32
	case "amd64":
		gapArch = gapstone.CS_ARCH_X86
		gapMode = gapstone.CS_MODE_64

	}

	engine, err := gapstone.New(gapArch, gapMode)
	if err != nil {
		return
	}
	defer engine.Close()

	err = engine.SetOption(gapstone.CS_OPT_SYNTAX, gapstone.CS_OPT_SYNTAX_ATT)
	if err != nil {
		return
	}

	codeBuf, err := ioutil.ReadAll(code)
	if err != nil {
		return
	}

	insns, err := engine.Disasm(codeBuf, address, 0)
	if err != nil {
		return
	}

	for _, insn := range insns {
		fmt.Fprintf(asm, "%x:\t% -32x\t%s\t\t%s\n", insn.Address, insn.Bytes, insn.Mnemonic, insn.OpStr)
	}

	return
}
