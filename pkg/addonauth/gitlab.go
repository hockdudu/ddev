package addonauth

import (
	"fmt"

	"go.yaml.in/yaml/v4"
)

type GitLabAuth struct {
	AuthType `yaml:",inline"`
	Remote   string `yaml:"remote,omitempty"`
	Token    string `yaml:"token"`
}

var GitLabAuthType AuthType = AuthType{Type: "gitlab-token"}

func (gla *GitLabAuth) Match(downloadUrl string, headers map[string]string) (bool, error) {
	if !doesUrlMatchRemote(downloadUrl, gla.Remote) {
		return false, nil
	}

	headers["Authorization"] = "Bearer " + gla.Token

	return true, nil
}

func (gla *GitLabAuth) String() string {
	return fmt.Sprintf("GitLab Auth for %s", gla.Remote)
}

func init() {
	// TODO: Normalize names? How is it called on Composer?
	TypeMap[GitLabAuthType.Type] = func(value *yaml.Node) (error, AuthMatcher) {
		gitLabAuth := GitLabAuth{}
		gitLabAuth.Type = GitLabAuthType.Type

		err := value.Decode(&gitLabAuth)
		if err != nil {
			return err, nil
		}

		if gitLabAuth.Remote == "" {
			gitLabAuth.Remote = "gitlab.com"
		}

		err, gitLabAuth.Remote = normalizeRemote(gitLabAuth.Remote)
		if err != nil {
			return err, nil
		}

		return nil, &gitLabAuth
	}
}
