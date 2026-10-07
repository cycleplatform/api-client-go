# Cycle API Go Client

<a href="https://cycle.io">
<picture class="red">
  <source media="(prefers-color-scheme: dark)" srcset="https://cycle.io/global/resources/images/logos/cycle-logo-white.svg">
  <source media="(prefers-color-scheme: light)" srcset="https://static.cycle.io/icons/logo/cycle-logo-fullcolor.svg">
  <img alt="cycle" width="300px" src="https://static.cycle.io/icons/logo/cycle-logo-fullcolor.svg">
</picture>
</a>

_This is an auto-generated API client based on the [OpenAPI Spec for Cycle](https://github.com/cycleplatform/api-spec). Please do not open any PRs for the generated code in `generated.go`. If you have any questions on what changes are made in the latest version, please refer to the spec above._

## Usage

`go get github.com/cycleplatform/api-client-go`

Create a client

```go
package main

import (
	"log" 
	"os"
	"context"
	"net/http"
	"fmt"
    "github.com/cycleplatform/api-client-go"
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

The client is generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) (pinned as a Go tool in `go.mod`) from the [api-spec](https://github.com/cycleplatform/api-spec) submodule. Requires Go (see `go.mod`) and Node.js.

### Updating the API spec

Update the API spec to the latest version:

`git submodule update --recursive --remote`

Install the spec's dependencies (first time only):

`(cd api-spec; npm ci)`

### Generating the client

`./generate.sh`

This bundles the spec, downconverts it from OpenAPI 3.1 to 3.0.3 (`api-spec/dist/platform-3.0.3.json`), and runs `go generate ./internal`, which writes `generated.go` using the options in `internal/config.yaml`.

### Releasing

The **Generate Client** GitHub Action (`.github/workflows/generate.yml`) does all of the above, commits the result, and publishes a release. Run it from the Actions tab and pick a release type:

- `patch` / `minor` / `major` bumps the latest `vX.Y.Z` tag and creates a GitHub release.
- `edge` moves the `edge` tag to the new commit and publishes it as a prerelease.

If the generated client hasn't changed since the last release (or the `edge` tag, for `edge`), no release is created unless `force` is checked. The api-spec repo can also trigger it with a `repository_dispatch` event of type `api-spec-updated`.
