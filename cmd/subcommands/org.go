package subcommands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/itzmanish/go-micro/v2/client"
	"github.com/itzmanish/slatomate/cmd/api"
	"github.com/itzmanish/slatomate/cmd/utils"
	"github.com/itzmanish/slatomate/proto/slatomate"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
)

// OrgCmd represents the orgnisation command
var OrgCmd = &cobra.Command{
	Use:   "org",
	Short: "CRUD operations on organization",
	Long:  `CRUD operations on organization`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// OrgListCmd represents the orgnisation command
var orgListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available organizations",
	Long:  `List available organizations`,
	Run: func(cmd *cobra.Command, args []string) {
		listOrganization()

	},
}

var orgCreateCmd = &cobra.Command{
	Use:   "create [Org Name]",
	Short: "Create new organization",
	Long:  `Create new organization`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createOrganization(args)

	},
}

var orgGetCmd = &cobra.Command{
	Use:   "get [orgID]",
	Short: "Get an organization",
	Long:  `Get an organization`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getOrganization(args)

	},
}

var orgDeleteCmd = &cobra.Command{
	Use:   "delete [orgID/all]",
	Short: "Delete an organization",
	Long:  `Delete an organization`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteOrganization(args)
	},
}

var orgAuthorizeCmd = &cobra.Command{
	Use:   "authorize [orgID]",
	Short: "Authorize an organization",
	Long:  `Authorize an organization`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		authorizeOrganization(args)

	},
}

var orgSetCmd = &cobra.Command{
	Use:   "set [org name]",
	Short: "set an organization to work with",
	Long:  `set an organization to work with`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setOrganization(args)
	},
}

var orgCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "show current organization that you are working with",
	Long:  `show current organization that you are working with`,
	Run: func(cmd *cobra.Command, args []string) {
		currentOrg()
	},
}

func init() {
	OrgCmd.AddCommand(orgListCmd, orgCreateCmd, orgGetCmd, orgDeleteCmd, orgAuthorizeCmd, orgSetCmd, orgCurrentCmd)
}

func listOrganization() {
	ctx, err := GetAuthContext()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Prefix = "Getting organizations "
	s.Start()
	orgs, err := api.APIClient.GetAllOrganization(ctx, &slatomate.GetAllOrganizationRequest{}, client.WithAddress(viper.GetString("service_host")))
	s.Stop()
	if err != nil {
		color.Red("\n%v", err)
		os.Exit(1)
	}
	color.Green("\nTotal Organizations available: %v\n", orgs.Count)
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"ID", "Name", "Slack APIKey", "Created At"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgHiWhiteColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgHiWhiteColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgHiWhiteColor},
		tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgHiWhiteColor},
	)

	for _, org := range orgs.Organizations {
		table.Append([]string{org.Id, org.Name, org.SlackApikey, org.CreatedAt})
	}

	table.Render()
}

func createOrganization(args []string) {
	name := args[0]
	ctx, err := GetAuthContext()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Creating organization %s ", name)
	s.Start()
	org, err := api.APIClient.CreateOrganization(ctx, &slatomate.CreateOrganizationRequest{Name: name}, client.WithAddress(viper.GetString("service_host")))
	s.Stop()
	if err != nil {
		color.Red("\n%v", err)
		return
	}
	color.Green("\nOrg created: %v\n", org)
	authorizeOrganization([]string{org.Id})
	// color.Green("Now authorize this organization to integrate with slack. To do so execute slatomate org authorize %v", org.Id)
}

func getOrganization(args []string) {
	id := args[0]
	ctx, err := GetAuthContext()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Getting organization %s ", id)
	s.Start()
	org, err := api.APIClient.GetOrganization(ctx, &slatomate.GetOrganizationRequest{Id: id}, client.WithAddress(viper.GetString("service_host")))
	s.Stop()
	if err != nil {
		color.Red("\n%v", err)
		return
	}
	color.Green("\nOrganization info: %v", org)
}

func deleteOrganization(args []string) {
	id := args[0]
	ctx, err := GetAuthContext()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Deleting organization %s ", id)
	s.Start()
	if id == "all" {
		_, err = api.APIClient.DeleteAllOrganization(ctx, &emptypb.Empty{}, client.WithAddress(viper.GetString("service_host")))
	} else {
		_, err = api.APIClient.DeleteOrganization(ctx, &slatomate.DeleteOrganizationRequest{Id: id}, client.WithAddress(viper.GetString("service_host")))
	}
	s.Stop()
	if err != nil {
		color.Red("\n%v", err)
		return
	}
	color.Green("\nOrganization deleted: %v", id)
}

func authorizeOrganization(args []string) {
	id := args[0]
	userID := viper.GetString("account.id")
	if len(userID) == 0 {
		color.Red("You are not logged in.")
		return
	}
	url := viper.GetString("oauth_url") + fmt.Sprintf("&redirect_uri=https://localhost:8888/v1/oauth/slack/callback?user_id=%v&org_id=%v", userID, id)
	err := utils.Openbrowser(url)

	if err != nil {
		color.White("Unable to open browser. Please navigate to the following link in your browser: %s", url)
	} else {
		color.White("Please navigate to the following link in your browser: %s", url)
	}
	code, err := startOauthResponseServer()
	if err != nil {
		color.Red("%v", err)
		return
	}
	ctx, err := GetAuthContext()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Authorizing organization %s ", id)
	s.Start()
	_, err = api.APIClient.AuthorizeOrganization(ctx, &slatomate.AuthorizeOrganizationRequest{Code: code, OrgId: id}, client.WithAddress(viper.GetString("service_host")))
	s.Stop()
	if err != nil {
		color.Red("\n%v", err)
		return
	}
	color.Green("\nOrganization authorised successfully.")
}

func setOrganization(args []string) {
	orgName := args[0]
	ctx, err := GetAuthContext()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Prefix = fmt.Sprintf("Getting organization %s ", orgName)
	s.Start()
	org, err := api.APIClient.GetOrganization(ctx, &slatomate.GetOrganizationRequest{Name: orgName}, client.WithAddress(viper.GetString("service_host")))
	s.Stop()
	if err != nil {
		color.Red("\nUnable to set organization Error: %v", err)
		return
	}
	viper.Set("org.name", org.Name)
	viper.Set("org.id", org.Id)
	err = viper.WriteConfig()
	if err != nil {
		color.Red("\nSaving config error: %v", err)
		return
	}
	color.Green("\n%v is now your working organization.", org.Name)
}

func currentOrg() {
	orgName := viper.GetString("org.name")
	if len(orgName) == 0 {
		color.Red("Current active orgnization is not set.")
		return
	}
	color.Green("Your active orgnization is %v", orgName)
}

func startOauthResponseServer() (string, error) {
	m := http.NewServeMux()
	// Get a self signed certificate setting for server
	serverTLS, _, err := utils.GetSelfSignedCerts()

	if err != nil {
		return "", err
	}
	s := &http.Server{Addr: ":8888", Handler: m, TLSConfig: serverTLS}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var code string = ""
	var error string = ""

	m.HandleFunc("/v1/oauth/slack/callback", func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")
		error = r.URL.Query().Get("error")
		w.Write([]byte("Slatomate connected with your slack workspace! You can close this browser tab"))
		cancel()
	})

	go func() {
		if err := s.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	s.Shutdown(ctx)
	if len(error) != 0 {
		return code, errors.New(error)
	}

	return code, nil
}
