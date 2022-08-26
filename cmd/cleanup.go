package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/happyRip/ttn-snapshot-cleaner/pkg/dockerhub"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup ...",
	Short: "Deletes tags older than retention period",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dockerhub.InitDockerhub(*cmd.Flags())
		if err != nil {
			return err
		}
		if dryRun {
			tags, err := dockerhub.ListTagsBefore(time.Now().Add(-dockerhub.Config.Retention))
			if err != nil {
				return err
			}
			out, err := marshalTags(tags)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		}
		return dockerhub.PerformCleanup()
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
