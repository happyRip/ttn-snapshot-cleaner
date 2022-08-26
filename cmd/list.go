package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/happyRip/ttn-snapshot-cleaner/pkg/dockerhub"
)

var listCmd = &cobra.Command{
	Use:   "list ...",
	Short: "List Docker Hub tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dockerhub.InitDockerhub(*cmd.Flags())
		if err != nil {
			return err
		}
		tags, err := dockerhub.ListTags()
		if err != nil {
			return err
		}
		out, err := marshalTags(tags)
		if err != nil {
			return err
		}
		fmt.Println(out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
