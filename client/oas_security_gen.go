// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// BearerAuth provides bearerAuth security value.
	BearerAuth(ctx context.Context, operationName string) (BearerAuth, error)
	// HubAuth provides hubAuth security value.
	// Defines the scope of the request to a specific Hub.
	HubAuth(ctx context.Context, operationName string) (HubAuth, error)
}

func (s *Client) securityBearerAuth(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.BearerAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"BearerAuth\"")
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return nil
}
func (s *Client) securityHubAuth(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.HubAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"HubAuth\"")
	}
	req.Header.Set("X-Hub-Id", t.APIKey)
	return nil
}
