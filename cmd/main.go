package main

import (
	"github.com/S7R4nG3/DCSUpdater/install"
	"github.com/S7R4nG3/DCSUpdater/prompts"
)

const rootRegistryKey string = `SOFTWARE\Eagle Dynamics\`

func main() {
	dcs := install.Installation{
		RegistryRootKey: rootRegistryKey,
	}
	dcs.Find()
	ui := prompts.UI{
		UpdaterPath: dcs.UpdaterPath,
		InstallPath: dcs.InstallPath,
	}
	ui.Prompt()
}
