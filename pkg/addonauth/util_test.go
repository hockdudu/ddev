package addonauth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoteNormalize(t *testing.T) {
	t.Run("Normalize", func(t *testing.T) {
		testCases := []struct {
			Remote     string
			Normalized string
		}{
			{Remote: "gitlab.com", Normalized: "gitlab.com"},
			{Remote: "GitLab.Com", Normalized: "gitlab.com"},
			{Remote: "example.com:1234", Normalized: "example.com:1234"},
			{Remote: "example.com:1234", Normalized: "example.com:1234"},
		}

		for _, tc := range testCases {
			err, normalized := normalizeRemote(tc.Remote)
			require.NoError(t, err)
			require.Equal(t, tc.Normalized, normalized, "Expected remote \"%s\" to be normalized as \"%s\", got \"%s\" instead", tc.Remote, tc.Normalized, normalized)
		}
	})
}

func TestUrlMatch(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		remote := "gitlab.com"
		testCases := []struct {
			Url   string
			Match bool
		}{
			{Url: "https://gitlab.com/ddev/ddev", Match: true},
			{Url: "https://GitLab.com/ddev/ddev", Match: true},
			{Url: "HTTPS://GITLAB.COM/DDEV/DDEV", Match: true},
			{Url: "https://www.gitlab.com/ddev/ddev", Match: true},
			{Url: "https://github.com/ddev/ddev", Match: false},
			{Url: "sftp://gitlab.com/ddev/ddev", Match: true},
		}

		for _, tc := range testCases {
			match := doesUrlMatchRemote(tc.Url, remote)

			expectString := "Expected URL \"%s\" to be matched by remote \"%s\""
			if !tc.Match {
				expectString = "Expected URL \"%s\" to not be matched by remote \"%s\""
			}
			require.Equal(t, tc.Match, match, expectString, tc.Url, remote)
		}
	})
}
