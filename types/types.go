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
	OpenAlpha = Version{
		Name: "openalpha",
	}
	Repair = Version{
		Name: "repair",
	}
	Cleanup = Version{
		Name: "cleanup",
	}
)
