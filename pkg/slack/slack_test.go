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
	token, err := slack.GetOAuthV2Response(http.DefaultClient,
		"1001856848789.2347720282836",
		"13160e61a6257cdb40006d149231aabe",
		"1001856848789.2416145373863.f193daef6d76fe58dc9de0aa480bf01ab98384b97b9d2635c4879bbe33ebc371",
		"")
	log.Println(token, err)
}
