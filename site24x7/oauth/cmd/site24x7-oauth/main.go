package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/sourcegraph/terraform-provider-site24x7/site24x7/oauth"
)

type oauthData struct {
	ClientId            string  `json:"CLIENT_ID"`
	ClientSecret        string  `json:"CLIENT_SECRET"`
	RefreshToken        string  `json:"REFRESH_TOKEN"`
}

var (
	clientId = flag.String("clientId", "", "(required) client id")
	clientSecret = flag.String("clientSecret", "", "(required) client secret")
	generateCode = flag.String("generateCode", "", "(required) generate code token")
)

func main() {
	flag.Parse()

	if *clientId == "" || *clientSecret == "" || *generateCode == "" {
		fmt.Fprintln(os.Stderr, "Follow the instructions at https://www.site24x7.com/help/api/index.html#authentication to obtain a client id, client secret and generate code")
		flag.PrintDefaults()
		os.Exit(2)
	}

	refreshToken, err := oauth.GenerateRefreshToken(*clientId, *clientSecret, *generateCode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	oad := &oauthData{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		RefreshToken: refreshToken,
	}

	contents, err := json.MarshalIndent(oad, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}

	fmt.Println(string(contents))
}
