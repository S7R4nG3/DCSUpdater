package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

const registryKey string = `SOFTWARE\Eagle Dynamics\`

func main() {
	k, _ := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.ALL_ACCESS)
	defer k.Close()

	subkeys, _ := k.ReadSubKeyNames(0)
	if len(subkeys) > 1 {
		fmt.Println("Found multiple installs...")
	} else {
		installKey := registryKey + subkeys[0]
		k, _ = registry.OpenKey(registry.CURRENT_USER, installKey, registry.ALL_ACCESS)
	}
	fmt.Println(k.GetStringValue("Path"))
}
