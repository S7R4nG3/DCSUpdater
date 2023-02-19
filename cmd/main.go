package main

import "github.com/S7R4nG3/DCSUpdater/install"

const rootRegistryKey string = `SOFTWARE\Eagle Dynamics\`

func main() {
	dcs := install.Installation{
		RegistryRootKey: rootRegistryKey,
	}
	dcs.Find()
}
