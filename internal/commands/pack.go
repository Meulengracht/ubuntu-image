package commands

// PackOpts holds all flags that are specific to the pack command
type PackOpts struct {
	ArtifactType string `long:"artifact-type" description:"Type of the resulting disk image file." required:"true" default:"raw" choice:"raw"`
	GadgetDir    string `long:"gadget-dir" description:"Directory containing the gadget.yaml and the gadget tree." required:"true"`
	RootfsDir    string `long:"rootfs-dir" description:"Directory containing the rootfs" required:"true"`
}

type PackCommand struct {
	PackOptsPassed PackOpts `required:"true"`
}