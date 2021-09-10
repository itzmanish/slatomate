package slack

import (
	"log"
	"net/http"
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

var testSlackClient *SlackClient

func TestNewClient(t *testing.T) {
	testSlackClient = NewSlackClient("xoxp-1001856848789-2040990405655-2381966723552-c657a0c560681078666336217006abd2")
	assert.NotNil(t, testSlackClient)
}

func TestGetUserInfo(t *testing.T) {
	TestNewClient(t)

	user, err := testSlackClient.GetUserProfile(&slack.GetUserProfileParameters{})
	assert.Nil(t, err)
	assert.NotNil(t, user)
	log.Printf("%+v", user)
}
func TestSlackToken(t *testing.T) {
	redirect_uri := "https://localhost:8080/slatomate/authorizeOrganization?user_id=fe2e4817-cd86-427a-9522-b8ff52dab0a1&org_id=983fffac-62e9-4366-9ebc-9028f2dadf2f"
	token, err := slack.GetOAuthV2Response(http.DefaultClient,
		"1001856848789.2347720282836",
		"13160e61a6257cdb40006d149231aabe",
		"1001856848789.2469830765286.b8d670e72fa6f04ae0898717aa11921051b8343b493c68bf0d8cdf2f1735928e",
		redirect_uri)
	log.Println(token, err)
}
