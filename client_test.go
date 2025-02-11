package cycle_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	cycle "github.com/cycleplatform/api-client-go"
)

func TestClient_canCall(t *testing.T) {
	apiKey := os.Getenv("CYCLE_API_KEY")
	if apiKey == "" {
		log.Fatal("missing env var CYCLE_API_KEY")
	}

	hubId := os.Getenv("CYCLE_HUB_ID")
	if hubId == "" {
		log.Fatal("missing env var CYCLE_HUB_ID")
	}

	baseUrl := os.Getenv("CYCLE_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://api.cycle.io"
	}

	c, err := cycle.NewAuthenticatedClient(cycle.ClientConfig{
		BaseURL: &baseUrl,
		APIKey:  apiKey,
		HubID:   hubId,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Starting a container
	{
		var task cycle.ContainerTask
		err := task.FromContainerStartAction(cycle.ContainerStartAction{Action: cycle.ContainerStartActionActionStart})
		if err != nil {
			log.Fatal(err)
		}

		res, err := c.CreateContainerJobWithResponse(context.TODO(), "678739ac492d4c2033df7a7c", task)
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode() != http.StatusAccepted {
			log.Fatalf("expected status code 202, got %v (%s)", res.StatusCode(), *res.JSONDefault.Error.Title)
		}

		fmt.Printf("Started container - Job ID %s\n", res.JSON202.Data.Job.Id)
	}

	// Listing environments with discriminated union LB service
	{
		resp, err := c.GetEnvironmentsWithResponse(context.TODO(), &cycle.GetEnvironmentsParams{})
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode() != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d %s", resp.StatusCode(), *resp.JSONDefault.Error.Title)
		}

		for _, v := range resp.JSON200.Data {
			fmt.Printf("Environment ID: %s - Name: %s\n", v.Id, v.Name)
			if v.Services.Loadbalancer == nil {
				fmt.Printf("  No Loadbalancer\n")
				continue
			}
			if v.Services.Loadbalancer.Config == nil {
				fmt.Printf("  No Loadbalancer Config Set\n")
				continue
			}

			d, err := v.Services.Loadbalancer.Config.Discriminator()
			if err != nil {
				log.Printf("  Error discrimining loadbalancer: %v\n", err)
				continue
			}

			fmt.Printf("discriminator type: %s\n", d)

			value, err := v.Services.Loadbalancer.Config.ValueByDiscriminator()
			if err != nil {
				fmt.Printf("  Error getting loadbalancer value: %v\n", err)
				continue
			}

			switch v := value.(type) {
			case cycle.V1LbType:
				fmt.Println("Using V1 Loadbalancer")
			case cycle.HaProxyLbType:
				fmt.Println("Using HAProxy Loadbalancer")
			case cycle.DefaultLbType:
				fmt.Println("Using Default Loadbalancer")
			default:
				log.Fatalf("Unknown load balancer type %#v", v)
			}

		}
	}

	// Updating a container name
	{
		name := "updated container name"
		resp, err := c.UpdateContainerWithResponse(context.TODO(), "67a18129133949a588564f12", cycle.UpdateContainerJSONRequestBody{
			Name: &name,
		})
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode() != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d %s", resp.StatusCode(), *resp.JSONDefault.Error.Title)
		}

		fmt.Printf("Name updated to %s\n", resp.JSON200.Data.Name)
	}
}
