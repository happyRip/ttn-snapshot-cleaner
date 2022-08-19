package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

var (
	logger  *log.Logger
	ctx     context.Context
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
			ctx := log.NewContext(context.Background(), logger)
			return nil
		},
	}
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
}
