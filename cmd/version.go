package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack-migrate/pkg/version"
)

func printVar(k, v string) {
	fmt.Printf("%-20s %s\n", k+":", v)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s: %s\n", cmd.Root().Short, cmd.Root().Name())
		printVar("Version", version.Version)
		if version.BuildDate != "" {
			printVar("Build date", version.BuildDate)
		}
		if version.GitCommit != "" {
			printVar("Git commit", version.GitCommit)
		}
		printVar("Go version", runtime.Version())
		printVar("OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
