package addonauth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitLabAuthMatch(t *testing.T) {
	t.Run("Url", func(t *testing.T) {
		testCases := []struct {
			Url   string
			Match bool
		}{
			{Url: "https://gitlab.com/ddev/ddev", Match: true},
			{Url: "https://www.gitlab.com/ddev/ddev", Match: true},
			{Url: "https://github.com/ddev/ddev", Match: false},
			{Url: "sftp://gitlab.com/ddev/ddev", Match: false},
		}

		gl := GitLabAuth{
			Remote: "https://gitlab.com",
		}
		h := make(map[string]string)
		for _, tc := range testCases {
			match, err := gl.Match(tc.Url, h)
			require.NoError(t, err)
			require.Equal(t, tc.Match, match)
		}
	})

	t.Run("Header Token", func(t *testing.T) {
		remote := "https://gitlab.com/ddev/ddev"
		testCases := []struct {
			Token   string
			Headers map[string]string
		}{
			{Token: "glpat-test", Headers: map[string]string{
				"Authorization": "Bearer glpat-test",
			}},
		}

		gl := GitLabAuth{
			Remote: "https://gitlab.com",
		}
		for _, tc := range testCases {
			gl.Token = tc.Token

			headers := make(map[string]string)
			match, err := gl.Match(remote, headers)
			require.NoError(t, err)
			require.True(t, match)

			require.Equal(t, tc.Headers, headers)
		}
	})
}
