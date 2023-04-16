package types

type Version struct {
	Name string
}

var (
	Stable = Version{
		Name: "stable",
	}
	OpenBeta = Version{
		Name: "openbeta",
	}
	Repair = Version{
		Name: "repair",
	}
	Cleanup = Version{
		Name: "cleanup",
	}
)
