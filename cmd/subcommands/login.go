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
	"log"

	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/metadata"
	"github.com/itzmanish/slatomate/cmd/api"
	"github.com/itzmanish/slatomate/cmd/utils"
	"github.com/itzmanish/slatomate/proto/slatomate"
	"github.com/spf13/cobra"
)

// LoginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "To log into slatomate",
	Long:  `Login yourself to slatomate with login command`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Login(cmd, args)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var WhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Get the current logged in user",
	Long:  `Get the information about current logged in user`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Whoami()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Login(cmd *cobra.Command, args []string) error {
	emailContent := utils.PromptContent{ErrorMsg: "Username is required", Label: "Email", Type: utils.TextPrompt}
	passowrdContent := utils.PromptContent{ErrorMsg: "Password is required", Label: "Password", Type: utils.PasswordPrompt}
	email := utils.PromptGetInput(emailContent)
	password := utils.PromptGetInput(passowrdContent)
	users, err := api.APIClient.LoginUser(context.TODO(), &slatomate.User{Email: email, Password: password})
	if err != nil {
		return err
	}
	log.Println(users)

	return nil
}

func Whoami() error {
	api_key := "VzUacFTqgYkjzqqPzmDqCXxYRwcO"
	ctx := context.TODO()
	if ok := metadata.SetOutgoingContext(ctx, metadata.Metadata{"Authorization": "APIKEY " + api_key}); !ok {
		return errors.InternalServerError("APIKEY_SET_FAILED", "Unable to set apikey")
	}
	u, err := api.APIClient.GetUser(ctx, &slatomate.GetUserRequest{ApiKey: api_key})
	if err != nil {
		return err
	}
	log.Println(u)
	return nil
}
