package addonauth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitHubAuthMatch(t *testing.T) {
	t.Run("Url", func(t *testing.T) {
		testCases := []struct {
			Url   string
			Match bool
		}{
			{Url: "https://github.com/ddev/ddev", Match: true},
			{Url: "https://github.com/test/asdf", Match: true},
			{Url: "https://gitlab.com/ddev/ddev", Match: false},
			{Url: "sftp://github.com/ddev/ddev", Match: false},
		}

		gh := GitHubAuth{
			Remote: "https://github.com",
		}
		h := make(map[string]string)
		for _, tc := range testCases {
			match, err := gh.Match(tc.Url, h)
			require.NoError(t, err)
			require.Equal(t, tc.Match, match)
		}
	})

	t.Run("Header Token", func(t *testing.T) {
		remote := "https://github.com/ddev/ddev"
		testCases := []struct {
			Token   string
			Headers map[string]string
		}{
			{Token: "example", Headers: map[string]string{
				"Authorization": "token example",
			}},
		}

		gh := GitHubAuth{
			Remote: "https://github.com",
		}
		for _, tc := range testCases {
			gh.Token = tc.Token

			headers := make(map[string]string)
			match, err := gh.Match(remote, headers)
			require.NoError(t, err)
			require.True(t, match)

			require.Equal(t, tc.Headers, headers)
		}
	})
}
