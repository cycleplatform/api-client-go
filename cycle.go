package cycle

import (
	"context"
	"net/http"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

// Function to combine multiple interceptors
func combineInterceptors(interceptors ...RequestEditorFn) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		for _, interceptor := range interceptors {
			if err := interceptor(ctx, req); err != nil {
				return err
			}
		}
		return nil
	}
}

type ClientConfig struct {
	APIKey  string
	HubID   string
	BaseURL *string
}

// generate a Cycle API client (with responses) that handles auth based on the provided token and hub
func NewAuthenticatedClient(config ClientConfig) (*ClientWithResponses, error) {
	if config.BaseURL == nil {
		baseUrl := "https://api.cycle.io"
		config.BaseURL = &baseUrl
	}

	authBearer, err := securityprovider.NewSecurityProviderBearerToken(config.APIKey)
	if err != nil {
		return nil, err
	}

	authHub, err := securityprovider.NewSecurityProviderApiKey("header", "X-Hub-Id", config.HubID)
	if err != nil {
		return nil, err
	}

	return NewClientWithResponses(*config.BaseURL, WithRequestEditorFn(
		combineInterceptors(
			authBearer.Intercept,
			authHub.Intercept,
		),
	))

}

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ./api-spec/dist/platform-3.0.3.json
