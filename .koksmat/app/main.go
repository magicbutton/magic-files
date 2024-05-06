package main

import (
	"runtime/debug"
	"strings"

	"github.com/magicbutton/magic-files/magicapp"
	"github.com/magicbutton/magic-files/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: magic-files
description: Describe the main purpose of this kitchen
---

# magic-files
`
	utils.Setup(".env")
	magicapp.RegisterServeCmd("magic-files", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.RegisterServiceCmd()
	magicapp.RegisterImportCmd()

	utils.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "magic-files", "")
}
