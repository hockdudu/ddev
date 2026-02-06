package addonauth

import (
	"errors"

	"go.yaml.in/yaml/v4"
)

type CustomHeadersAuth struct {
	AuthType `yaml:",inline"`
	Remote   string            `yaml:"remote,omitempty"`
	Headers  map[string]string `yaml:"headers"`
}

var CustomHeadersAuthType AuthType = AuthType{Type: "custom-headers"}

func (customHeadersAuth *CustomHeadersAuth) Match(downloadUrl string, headers map[string]string) (bool, error) {
	if !doesUrlMatchRemote(downloadUrl, customHeadersAuth.Remote) {
		return false, nil
	}

	for name, value := range customHeadersAuth.Headers {
		headers[name] = value
	}

	return true, nil
}

func (customHeadersAuth *CustomHeadersAuth) String() string {
	return "Custom Headers Auth for " + customHeadersAuth.Remote
}

func init() {
	TypeMap[CustomHeadersAuthType.Type] = func(value *yaml.Node) (AuthMatcher, error) {
		customHeadersAuth := CustomHeadersAuth{}
		customHeadersAuth.Type = CustomHeadersAuthType.Type

		err := value.Decode(&customHeadersAuth)
		if err != nil {
			return nil, err
		}

		if customHeadersAuth.Remote == "" {
			return nil, errors.New("custom headers auth remote is missing")
		}

		if len(customHeadersAuth.Headers) == 0 {
			return nil, errors.New("custom headers auth headers list is empty")
		}

		customHeadersAuth.Remote, err = normalizeRemote(customHeadersAuth.Remote)
		if err != nil {
			return nil, err
		}

		return &customHeadersAuth, nil
	}
}
