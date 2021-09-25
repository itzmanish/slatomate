/*
Copyright © 2021 Manish <itzmanish108@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package subcommands

import (
	"context"

	"github.com/fatih/color"
	"github.com/itzmanish/go-micro/v2/client"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/metadata"
	"github.com/itzmanish/slatomate/cmd/api"
	"github.com/itzmanish/slatomate/cmd/utils"
	v1 "github.com/itzmanish/slatomate/proto/gen/slatomate/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LoginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "To log into slatomate",
	Long:  `Login yourself to slatomate with login command`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Login(cmd, args)
		if err != nil {
			err := errors.FromError(err)
			if err.Code == 404 {
				color.Red("No user exists with this email")
			} else if err.Code == 500 {
				color.Red("Something is wrong with server.")
			} else {
				color.Red("Email/Password didn't match.")
			}
		}
	},
}

var WhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Get the current logged in user",
	Long:  `Get the information about current logged in user`,
	Run: func(cmd *cobra.Command, args []string) {
		Whoami()

	},
}

func Login(cmd *cobra.Command, args []string) error {
	emailContent := utils.PromptContent{ErrorMsg: "Username is required", Label: "Email", Type: utils.TextPrompt}
	passowrdContent := utils.PromptContent{ErrorMsg: "Password is required", Label: "Password", Type: utils.PasswordPrompt}
	email := utils.PromptGetInput(emailContent)
	password := utils.PromptGetInput(passowrdContent)
	user, err := api.APIClient.LoginUser(context.TODO(), &v1.User{Email: email, Password: password})
	if err != nil {
		return err
	}
	viper.Set("account.token", user.ApiKey)
	viper.Set("account.id", user.Id)
	viper.WriteConfig()
	color.Green("Successfully logged in")
	return nil
}

func Whoami() {
	auth_token, err := IsAuthorized()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	ctx := metadata.Set(context.TODO(), "Authorization", ("APIKEY " + auth_token))
	u, err := api.APIClient.GetUser(ctx, &v1.GetUserRequest{ApiKey: auth_token}, client.WithAddress(viper.GetString("service_host")))
	if err != nil {
		color.Red("Got error: %s", err.Error())
		return
	}
	color.Green("You are logged in as %s", u.Name)
}

func IsAuthorized() (string, error) {
	auth_token := viper.GetString("account.token")
	if len(auth_token) == 0 {
		return "", errors.Unauthorized("UNAUTHORIZED", "You are not logged in.")

	}
	return auth_token, nil
}

func GetAuthContext() (context.Context, error) {
	auth_token, err := IsAuthorized()
	if err != nil {
		return nil, err
	}
	return metadata.Set(context.TODO(), "Authorization", ("APIKEY " + auth_token)), nil
}
