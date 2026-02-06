package cmd

import (
	"github.com/ddev/ddev/pkg/addonauth"
	"github.com/ddev/ddev/pkg/globalconfig"
	"github.com/ddev/ddev/pkg/util"
	"github.com/spf13/cobra"
)

var AuthAddOnCommand = &cobra.Command{
	Use:   "add-on",
	Short: "Add-on authentication commands",
	Long:  `Add-on authentication commands`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()
		util.CheckErr(err)
	},
}

var AuthAddOnGithub = &cobra.Command{
	Use:   "github",
	Short: "Add-on authentication for GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		token, _ := cmd.Flags().GetString("token")

		err = globalconfig.ReadAuthConfig()

		auth := addonauth.AddonAuth{
			Matcher: &addonauth.GitHubAuth{
				AuthType: addonauth.GithubAuthType,
				Remote:   "github.com",
				Token:    token,
			},
		}

		globalconfig.DdevAuth.Addon = append(globalconfig.DdevAuth.Addon, auth)

		err = globalconfig.WriteAuthConfig(&globalconfig.DdevAuth)
		util.CheckErr(err)
	},
}

func init() {
	AuthCmd.AddCommand(AuthAddOnCommand)

	AuthAddOnGithub.Flags().String("token", "", "GitHub personal access token")
	AuthAddOnCommand.AddCommand(AuthAddOnGithub)
}
