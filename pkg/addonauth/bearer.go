package addonauth

import (
	"errors"
	"fmt"

	"go.yaml.in/yaml/v4"
)

type BearerAuth struct {
	AuthType `yaml:",inline"`
	Remote   string `yaml:"remote,omitempty"`
	Token    string `yaml:"token"`
}

var BearerAuthType AuthType = AuthType{Type: "bearer"}

func (bearerAuth *BearerAuth) Match(downloadUrl string, headers map[string]string) (bool, error) {
	if !doesUrlMatchRemote(downloadUrl, bearerAuth.Remote) {
		return false, nil
	}

	headers["Authorization"] = "Bearer " + bearerAuth.Token

	return true, nil
}

func (bearerAuth *BearerAuth) String() string {
	return fmt.Sprintf("Bearer Auth for %s", bearerAuth.Remote)
}

func init() {
	TypeMap[BearerAuthType.Type] = func(value *yaml.Node) (error, AuthMatcher) {
		bearerAuth := BearerAuth{}
		bearerAuth.Type = BearerAuthType.Type

		err := value.Decode(&bearerAuth)
		if err != nil {
			return err, nil
		}

		if bearerAuth.Remote == "" {
			return errors.New("bearer auth remote is missing"), nil
		}

		err, bearerAuth.Remote = normalizeRemote(bearerAuth.Remote)
		if err != nil {
			return err, nil
		}

		if bearerAuth.Token == "" {
			return errors.New("bearer token is missing"), nil
		}

		return nil, &bearerAuth
	}
}
