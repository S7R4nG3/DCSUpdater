package install

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

type Installation struct {
	RegistryRootKey string
	InstallationKey string
	UpdaterPath     string
	InstallPath     string
}

func (i *Installation) Find() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, i.RegistryRootKey, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	sk, err := k.ReadSubKeyNames(0)
	if err != nil {
		return err
	}

	if len(sk) > 1 {
		log.Println("Found multiple installs...")
	} else {
		i.InstallationKey = i.RegistryRootKey + sk[0]
		k, _ = registry.OpenKey(registry.CURRENT_USER, i.InstallationKey, registry.ALL_ACCESS)
	}
	i.UpdaterPath, _, err = k.GetStringValue("Path")
	i.InstallPath = i.UpdaterPath
	if err != nil {
		return err
	}
	i.UpdaterPath += `\bin\dcs_updater.exe`
	return nil
}
