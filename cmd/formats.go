package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/docker/hub-tool/pkg/hub"
	"gopkg.in/yaml.v3"
)

type formats struct {
	json bool
	yaml bool
}

var format formats

func marshalTags(tags []hub.Tag) (string, error) {
	switch {
	case format.json:
		b, err := json.MarshalIndent(tags, "", "  ")
		if err != nil {
			return "", err
		}
		return string(b), nil
	case format.yaml:
		b, err := yaml.Marshal(tags)
		if err != nil {
			return "", err
		}
		return string(b), nil
	default:
		var s string
		for _, tag := range tags {
			s += fmt.Sprintln(tag.Name)
		}
		return s, nil
	}
}
