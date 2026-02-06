package addonauth

import (
	"fmt"

	"go.yaml.in/yaml/v4"
)

type AuthType struct {
	Type string `yaml:"type"`
}

type AddonAuth struct {
	Matcher AuthMatcher
}

type AuthMatcher interface {
	fmt.Stringer
	Match(downloadUrl string, headers map[string]string) (bool, error)
}

var (
	TypeMap = map[string]func(value *yaml.Node) (AuthMatcher, error){}
)

func (auth AddonAuth) UnmarshalYAML(value *yaml.Node) (err error) {
	authType := AuthType{}
	err = value.Decode(&authType)
	if err != nil {
		return err
	}

	matcherFunc := TypeMap[authType.Type]
	if matcherFunc == nil {
		return fmt.Errorf("unknown auth type: %s", authType.Type)
	}

	matcher, err := matcherFunc(value)
	if err != nil {
		return err
	}

	auth.Matcher = matcher

	return nil
}

func (auth AddonAuth) MarshalYAML() (interface{}, error) {
	return auth.Matcher, nil
}
