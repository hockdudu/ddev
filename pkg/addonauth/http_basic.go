package addonauth

import (
	"errors"
	"fmt"

	"encoding/base64"

	"go.yaml.in/yaml/v4"
)

type HttpBasicAuth struct {
	AuthType `yaml:",inline"`
	Remote   string `yaml:"remote,omitempty"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var HttpBasicAuthType AuthType = AuthType{Type: "http-basic"}

func (basicAuth *HttpBasicAuth) Match(downloadUrl string, headers map[string]string) (bool, error) {
	if !doesUrlMatchRemote(downloadUrl, basicAuth.Remote) {
		return false, nil
	}

	encodedLogin := base64.StdEncoding.EncodeToString([]byte(basicAuth.Username + ":" + basicAuth.Password))

	headers["Authorization"] = "Basic " + encodedLogin

	return true, nil
}

func (basicAuth *HttpBasicAuth) String() string {
	return fmt.Sprintf("Basic Auth for %s with username %s", basicAuth.Remote, basicAuth.Username)
}

func init() {
	TypeMap[HttpBasicAuthType.Type] = func(value *yaml.Node) (AuthMatcher, error) {
		basicAuth := HttpBasicAuth{}
		basicAuth.Type = HttpBasicAuthType.Type

		err := value.Decode(&basicAuth)
		if err != nil {
			return nil, err
		}

		if basicAuth.Remote == "" {
			return nil, errors.New("basic auth remote is missing")
		}

		basicAuth.Remote, err = normalizeRemote(basicAuth.Remote)
		if err != nil {
			return nil, err
		}

		return &basicAuth, nil
	}
}
