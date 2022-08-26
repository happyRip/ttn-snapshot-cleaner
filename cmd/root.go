package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/rpclog"
)

var (
	logger  *log.Logger
	rootCmd = &cobra.Command{
		Use:   "ttn-snapshot-cleaner",
		Short: "Cleanup old snapshots",

		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logLevel := log.InfoLevel
			if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
				logLevel = log.DebugLevel
			}
			logHandler, err := log.NewZap("console")
			if err != nil {
				return err
			}
			logger := log.NewLogger(logHandler, log.WithLevel(logLevel))
			rpclog.ReplaceGrpcLogger(logger)
			return nil
		},
	}
	dryRun bool
)

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		printStack(os.Stderr, err)
		return 1
	}
	return 0
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose output")
	rootCmd.PersistentFlags().String("retention", os.Getenv("CLEANER_RETENTION"), "Snapshot retention in days. Snapshots past retentions will be removed when using a `delete` command.")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "When set no snapshots will be deleted.")
	rootCmd.PersistentFlags().BoolVar(&format.json, "json", false, "Print output in json format")
	rootCmd.PersistentFlags().BoolVar(&format.yaml, "yaml", false, "Print output in yaml format")
	rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")
}
