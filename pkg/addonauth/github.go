package addonauth

import (
	"fmt"

	"go.yaml.in/yaml/v4"
)

type GitHubAuth struct {
	AuthType `yaml:",inline"`
	Remote   string `yaml:"remote"`
	Token    string `yaml:"token"`
}

var GithubAuthType AuthType = AuthType{Type: "github"}

func (gh *GitHubAuth) Match(downloadUrl string, headers map[string]string) (bool, error) {
	if !doesUrlMatchRemote(downloadUrl, gh.Remote) {
		return false, nil
	}

	headers["Authorization"] = "token " + gh.Token

	return true, nil
}

func (gh *GitHubAuth) String() string {
	return fmt.Sprintf("GitHub Auth for %s", gh.Remote)
}

func init() {
	TypeMap[GithubAuthType.Type] = func(value *yaml.Node) (error, AuthMatcher) {
		gitHubAuth := GitHubAuth{}
		gitHubAuth.Type = GithubAuthType.Type

		err := value.Decode(&gitHubAuth)
		if err != nil {
			return err, nil
		}

		if gitHubAuth.Remote == "" {
			gitHubAuth.Remote = "github.com"
		}

		err, gitHubAuth.Remote = normalizeRemote(gitHubAuth.Remote)
		if err != nil {
			return err, nil
		}

		return nil, &gitHubAuth
	}
}
