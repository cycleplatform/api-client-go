# Cycle API Go Client

_This is an auto-generated API client based on the [OpenAPI Spec for Cycle](https://github.com/cycleplatform/api-spec). Please do not open any PRs for the generated code under /client. If you have any questions on what changes are made in the latest version, please refer to the spec above._

## Setup

Install Ogen

`go install -v github.com/ogen-go/ogen/cmd/ogen@latest`

## Generating

### Updating the spec

`git submodule update --recursive --remote`

Using `npm`, run `(cd api-spec; npm run build:public && npm run build:internal)`

### Generating the client

`go generate ./...`

See [ogen](https://ogen.dev/docs/intro/) for usage.
