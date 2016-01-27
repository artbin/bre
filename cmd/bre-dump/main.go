package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArtemKulyabin/bre"
	"github.com/ArtemKulyabin/bre/binaryx"
	"github.com/ArtemKulyabin/bre/disasm"
)

func main() {
	f, err := bre.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("format", f.Format())
	fmt.Println("os", f.Os())
	fmt.Println("arch", f.Arch())
	fmt.Println("type", f.Type())

	for _, section := range f.Sections() {
		if section.Perm()&binaryx.PermExecute != 0 {
			fmt.Println("Disassembly of section", section.Name()+":")

			code := section.Open()

			err = disasm.Disasm(code, os.Stdout, section.Addr(), f.Arch())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
