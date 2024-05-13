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
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	az "github.com/microsoftgraph/msgraph-sdk-go-core/authentication"
	"github.com/rs/zerolog/log"
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

type CustomTokenCredential struct {
	Token string
	ExpiresOn time.Time
}

func (c *CustomTokenCredential) GetToken(ctx context.Context, options azpolicy.TokenRequestOptions) (azcore.AccessToken, error) {
	token := azcore.AccessToken{}
	token.Token = c.Token
	//token.ExpiresOn = c.ExpiresOn
	return token, nil
}

func NewGraphServiceClientWithToken(token string) (*msgraphsdk.GraphServiceClient, error) {
	validhosts := []string{"graph.microsoft.com", "graph.microsoft.us", "dod-graph.microsoft.us", "graph.microsoft.de", "microsoftgraph.chinacloudapi.cn", "canary.graph.microsoft.com"}
	scopes := []string{"https://graph.microsoft.com/.default"}
	var customCreds azcore.TokenCredential = &CustomTokenCredential{
		Token: token,
		//ExpiresOn: time.Now().AddDate(0, 0, 1),
	}

	auth, err := az.NewAzureIdentityAuthenticationProviderWithScopesAndValidHosts(customCreds, scopes, validhosts)
	if err != nil {
		return nil, err
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}

	client := msgraphsdk.NewGraphServiceClient(adapter)
	return client, nil
}

func GraphAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantId := os.Getenv("TenantId")
		clientId := os.Getenv("ClientId")
		clientSecret := os.Getenv("ClientSecret")
		refreshToken := os.Getenv("RefreshToken")
    if clientId == "" {
      log.Error().Msg("Client ID cannot be empty")
      return
    }

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
		graph, err := NewGraphServiceClientWithToken(tokenResponse.AccessToken)
		if err != nil {
			log.Error().Err(err).Msg("")
			return
		}

		ctx := context.WithValue(r.Context(), "graph", graph)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
