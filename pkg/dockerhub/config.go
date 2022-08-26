package dockerhub

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
)

type config struct {
	ctx context.Context

	username   string
	password   string
	token      string
	repository string

	Retention time.Duration
}

func flagSet() *pflag.FlagSet {
	flags := &pflag.FlagSet{}
	flags.String(dockerhubFlag("username"), os.Getenv("DOCKER_HUB_USERNAME"), "Docker Hub User name")
	flags.String(dockerhubFlag("password"), os.Getenv("DOCKER_HUB_PASSWORD"), "Docker Hub User password")
	flags.String(dockerhubFlag("token"), os.Getenv("DOCKER_HUB_TOKEN"), "Docker Hub Token")
	flags.String(dockerhubFlag("repository"), os.Getenv("DOCKER_HUB_REPOSITORY"), "Docker Hub Repository")
	return flags
}

func getConfig(flags *pflag.FlagSet) (*config, error) {
	username, err := flags.GetString(dockerhubFlag("username"))
	if err != nil {
		return nil, err
	}
	password, err := flags.GetString(dockerhubFlag("password"))
	if err != nil {
		return nil, err
	}
	token, err := flags.GetString(dockerhubFlag("token"))
	if err != nil {
		return nil, err
	}
	repository, err := flags.GetString(dockerhubFlag("repository"))
	if err != nil {
		return nil, err
	}
	retention, err := func() (time.Duration, error) {
		r, err := flags.GetString("retention")
		if err != nil {
			return time.Duration(0), err
		}
		return parseRetentionString(r)
	}()
	if err != nil {
		return nil, err
	}
	return &config{
		ctx: context.Background(),

		username:   username,
		password:   password,
		token:      token,
		repository: repository,

		Retention: retention,
	}, nil
}

func parseRetentionString(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

func dockerhubFlag(name string) string {
	return fmt.Sprintf("docker-hub.%s", name)
}
