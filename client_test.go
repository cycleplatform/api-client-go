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

	c, err := cycle.NewAuthenticatedClient(cycle.ClientConfig{
		APIKey: apiKey,
		HubID:  hubId,
	})

	if err != nil {
		log.Fatal(err)
	}

	{

		resp, err := c.GetVirtualMachinesWithResponse(context.TODO(), &cycle.GetVirtualMachinesParams{})
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode() != http.StatusOK {
			log.Fatalf("Expected HTTP 200 but received %d %s", resp.StatusCode(), *resp.JSONDefault.Error.Title)
		}

		for _, v := range resp.JSON200.Data {
			fmt.Printf("ID: %s - Name: %s\n", v.Id, v.Name)

			fmt.Println(v.Image.Discriminator())
			value, err := v.Image.ValueByDiscriminator()
			if err != nil {
				log.Fatal(err)
			}
			switch v := value.(type) {
			case cycle.VirtualMachineImageSourceBase:
				fmt.Println("Base image source:", v)
			case cycle.VirtualMachineImageSourceIpxe:
				fmt.Println("iPXE image source:", v)
			case cycle.VirtualMachineImageSourceUrl:
				fmt.Println("URL image source:", v)
			default:
				log.Fatalf("Unknown image source type %#v", v)
			}

		}

	}

	{
		name := "mattoni"
		resp, err := c.UpdateVirtualMachineWithResponse(context.TODO(), "67a18129133949a588564f12", cycle.UpdateVirtualMachineJSONRequestBody{
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
