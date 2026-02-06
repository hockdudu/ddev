package addonauth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.yaml.in/yaml/v4"
)

func TestUnmarshallTypes(t *testing.T) {
	t.Run("GitHubType", func(t *testing.T) {
		testCases := []string{
			"{type: github, token: example}",
		}

		for _, testCase := range testCases {
			var auth *AddonAuth
			err := yaml.Unmarshal([]byte(testCase), &auth)
			require.NoError(t, err)

			require.NotNil(t, auth.Matcher)
			gh, ok := auth.Matcher.(*GitHubAuth)
			require.True(t, ok)

			require.IsType(t, &GitHubAuth{}, gh)
			require.Equal(t, "github.com", gh.Remote)
		}
	})

	t.Run("GitLabTokenType", func(t *testing.T) {
		testCases := []string{
			"{type: gitlab-token, token: example}",
			"{type: gitlab-token, remote: gitlab.com, token: example}",
		}

		for _, testCase := range testCases {
			var auth *AddonAuth
			err := yaml.Unmarshal([]byte(testCase), &auth)
			require.NoError(t, err)

			require.NotNil(t, auth.Matcher)
			gl, ok := auth.Matcher.(*GitLabAuth)
			require.True(t, ok)

			require.IsType(t, &GitLabAuth{}, gl)
			require.Equal(t, "gitlab.com", gl.Remote)
		}
	})

	t.Run("HttpBasicType", func(t *testing.T) {
		testCases := []string{
			"{type: http-basic, remote: example.com, username: example, password: secret}",
		}

		for _, testCase := range testCases {
			var auth *AddonAuth
			err := yaml.Unmarshal([]byte(testCase), &auth)
			require.NoError(t, err)

			require.NotNil(t, auth.Matcher)
			require.IsType(t, &HttpBasicAuth{}, auth.Matcher)

			hb, ok := auth.Matcher.(*HttpBasicAuth)

			require.True(t, ok)
			require.Equal(t, "example", hb.Username)
			require.Equal(t, "secret", hb.Password)
		}
	})

	t.Run("CustomHeadersType", func(t *testing.T) {
		testCases := []string{
			"{type: custom-headers, remote: example.com, headers: {Authorization: Bearer example, X-Custom-Header: example}}",
		}

		for _, testCase := range testCases {
			var auth *AddonAuth
			err := yaml.Unmarshal([]byte(testCase), &auth)
			require.NoError(t, err)

			require.NotNil(t, auth.Matcher)
			require.IsType(t, &CustomHeadersAuth{}, auth.Matcher)
			ch, ok := auth.Matcher.(*CustomHeadersAuth)
			require.True(t, ok)
			require.Equal(t, "example.com", ch.Remote)
			require.Equal(t, "example", ch.Headers["X-Custom-Header"])
			require.Equal(t, "Bearer example", ch.Headers["Authorization"])
		}
	})

	t.Run("BearerType", func(t *testing.T) {
		testCases := []string{
			"{type: bearer, remote: example.com, token: example}",
		}

		for _, testCase := range testCases {
			var auth *AddonAuth
			err := yaml.Unmarshal([]byte(testCase), &auth)
			require.NoError(t, err)

			require.NotNil(t, auth.Matcher)
			require.IsType(t, &BearerAuth{}, auth.Matcher)
			require.Equal(t, "example.com", auth.Matcher.(*BearerAuth).Remote)
			require.Equal(t, "example", auth.Matcher.(*BearerAuth).Token)
		}
	})

}
