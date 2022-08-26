package dockerhub

import (
	"time"

	"github.com/docker/hub-tool/pkg/hub"
	"github.com/spf13/pflag"
)

var (
	client *hub.Client
	Config *config
)

func InitDockerhub(flags pflag.FlagSet) error {
	flags.AddFlagSet(flagSet())
	var err error
	Config, err = getConfig(&flags)
	if err != nil {
		return err
	}
	client, err = hub.NewClient(
		hub.WithHubAccount(Config.username),
		hub.WithPassword(Config.password),
	)
	return err
}

func ListTags() ([]hub.Tag, error) {
	tags, _, err := client.GetTags(Config.repository)
	return tags, err
}

func ListTagsBefore(t time.Time) ([]hub.Tag, error) {
	tags, err := ListTags()
	if err != nil {
		return nil, err
	}
	for i, tag := range tags {
		if tag.LastUpdated.Before(t) {
			if i <= 0 {
				return nil, nil
			}
			return tags[i:], nil
		}
	}
	return tags, nil
}

func DeleteTags(tags ...hub.Tag) error {
	for _, tag := range tags {
		if err := client.RemoveTag(Config.repository, tag.Name); err != nil {
			return err
		}
	}
	return nil
}

func PerformCleanup() error {
	tags, err := ListTagsBefore(time.Now().Add(-Config.Retention))
	if err != nil {
		return err
	}
	return DeleteTags(tags...)
}
