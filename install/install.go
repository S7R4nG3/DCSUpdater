package install

import (
	"log"

	"github.com/S7R4nG3/DCSUpdater/utils"
	"golang.org/x/sys/windows/registry"
)

type Installation struct {
	RegistryRootKey string
	InstallationKey string
	InstallPath     string
}

func (i Installation) Find() {
	k, err := registry.OpenKey(registry.CURRENT_USER, i.RegistryRootKey, registry.ALL_ACCESS)
	utils.Check(err)
	defer k.Close()

	sk, err := k.ReadSubKeyNames(0)
	utils.Check(err)

	if len(sk) > 1 {
		log.Println("Found multiple installs...")
	} else {
		installKey := i.RegistryRootKey + sk[0]
		k, _ = registry.OpenKey(registry.CURRENT_USER, installKey, registry.ALL_ACCESS)
	}
	path, _, err := k.GetStringValue("Path")
	utils.Check(err)
	log.Println(path)
}
