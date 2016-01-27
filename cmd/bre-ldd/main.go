package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArtemKulyabin/bre/ldd"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bre-ldd"
	app.Version = "0.0.1"
	app.Usage = "List dynamic dependencies - prints the dynamic libraries required by each program or dynamic library"
	app.Action = lddcmd
	app.Run(os.Args)
}

func lddcmd(c *cli.Context) {
	for _, name := range c.Args() {
		fmt.Println(name, ":")
		libs, err := ldd.GetDynLibs(name)
		if err != nil {
			log.Println("\t", err)
		}
		if len(libs) == 0 {
			fmt.Println("\t", "not a dynamic executable")
			continue
		}
		for _, lib := range libs {
			fmt.Println("\t", lib)
		}
	}
}
