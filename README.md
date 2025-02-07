# Cycle API Go Client

_This is an auto-generated API client based on the [OpenAPI Spec for Cycle](https://github.com/cycleplatform/api-spec). Please do not open any PRs for the generated code in `generated.go`. If you have any questions on what changes are made in the latest version, please refer to the spec above._


## Usage

`go get github.com/cycleplatform/api-client-go`

Create a client

```go
package main

import (
    cycle "github.com/cycleplatform/api-client-go"
)

func main() {
    apiKey := os.Getenv("CYCLE_API_KEY")
	if apiKey == "" {
		log.Fatal("missing env var CYCLE_API_KEY")
	}

	hubId := os.Getenv("CYCLE_HUB_ID")
	if hubId == "" {
		log.Fatal("missing env var CYCLE_HUB_ID")
	}

	c, err := cycle.NewAuthenticatedClient(cycle.ClientConfig{
		APIKey: apiKey,
		HubID:  hubId,
	})

	if err != nil {
		log.Fatal(err)
	}

    // Get list of environments
    resp, err := c.GetEnvironmentsWithResponse(context.TODO(), &cycle.GetEnvironmentsParams{})
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode() != http.StatusOK {
        log.Fatalf("Expected HTTP 200 but received %d %s", resp.StatusCode(), *resp.JSONDefault.Error.Title)
    }

    for _, v := range resp.JSON200.Data {
        fmt.Printf("ID: %s - Name: %s\n", v.Id, v.Name)
    }
}
```

## Development

### Updating the API spec

Update the API spec to the latest version:

`git submodule update --recursive --remote`

Using `npm`, run `(cd api-spec; npm run build:platform && npm run downconvert:platform)`

### Generating the client

`go generate ./...`

See [ogen](https://ogen.dev/docs/intro/) for usage.
