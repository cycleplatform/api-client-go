package api

import "context"

type CycleApiSecurity struct {
	token       string
	activeHubId string
}

func (s *CycleApiSecurity) SetAuthToken(token string) {
	s.token = token
}

func (s *CycleApiSecurity) SetActiveHubId(activeHubId string) {
	s.activeHubId = activeHubId
}

func NewCycleApiSecurity(token string, activeHubId string) CycleApiSecurity {
	cas := CycleApiSecurity{
		token:       token,
		activeHubId: activeHubId,
	}

	return cas
}

func (s CycleApiSecurity) BearerAuth(ctx context.Context, operationName string) (BearerAuth, error) {
	var ba BearerAuth
	ba.SetToken(s.token)
	return ba, nil
}

func (s CycleApiSecurity) HubAuth(ctx context.Context, operationName string) (HubAuth, error) {
	var ha HubAuth
	ha.SetAPIKey(s.activeHubId)
	return ha, nil
}
