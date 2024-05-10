package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/microsoft/kiota-abstractions-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/rs/zerolog/log"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

type CustomRequestAdapter struct {
	underlyingAdapter abstractions.RequestAdapter
	token             string
}

func NewCustomRequestAdapter(adapter abstractions.RequestAdapter, token string) *CustomRequestAdapter {
	return &CustomRequestAdapter{
		underlyingAdapter: adapter,
		token:             token,
	}
}

func GraphAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantId := os.Getenv("TenantID")
		clientId := os.Getenv("ClientId")
		clientSecret := os.Getenv("ClientSecret")
		refreshToken := os.Getenv("RefreshToken")

		form := url.Values{}
		form.Set("client_id", clientId)
		form.Set("scope", "https://graph.microsoft.com/.default")
		form.Set("refresh_token", refreshToken)
		form.Set("grant_type", "refresh_token")
		form.Set("client_secret", clientSecret)

		req, err := http.NewRequest("POST", fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantId), strings.NewReader(form.Encode()))
		if err != nil {
			log.Error().Err(err)
			return
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Error().Err(err)
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error().Err(err)
			return
		}

		var tokenResponse TokenResponse
		json.Unmarshal(body, &tokenResponse)

		id, err := azidentity.NewGraphServiceClient()

		adapter := NewCustomRequestAdapter()
		graph := msgraphsdk.NewGraphServiceClient(adapter)

		//pass the graph client in context
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
