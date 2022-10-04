package pkg

import (
	"ohiecli/core"
	"ohiecli/util"

	"github.com/spf13/cobra"
)

func PackageRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"r", "destroy"},
		Short:   "Remove everything related to a package (volumes, configs, etc)",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := getConfigFromParams(cmd)
			util.PanicError(err)
			packageSpec, err := getPackageSpecFromParams(cmd, config)
			util.PanicError(err)
			packageSpec.DeployCommand = "destroy"
			packageSpec, err = loadInProfileParams(cmd, *config, *packageSpec)
			util.PanicError(err)

			err = core.LaunchPackage(*packageSpec, *config)
			util.PanicError(err)
		},
	}

	setPackageActionFlags(cmd)

	return cmd
}
