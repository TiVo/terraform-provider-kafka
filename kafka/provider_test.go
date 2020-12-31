package kafka

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testProvider *schema.Provider

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	datProvider()
	meta := testProvider.Meta()
	if meta == nil {
		t.Fatal("Could not construct client")
	}
	client := meta.(*LazyClient)
	if client == nil {
		t.Fatal("No client")
	}
}

func AccTestProviderConfig() *terraform.ResourceConfig {
	wat := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	bs := strings.Split(wat, ",")
	if len(bs) == 0 || wat != "" {
		bs = []string{"localhost:9092"}
	}

	bootstrapServers := []interface{}{}
	for _, v := range bs {
		if v != "" {
			bootstrapServers = append(bootstrapServers, v)
		}
	}

	ca, err := ioutil.ReadFile("../secrets/ca.crt")
	if err != nil {
		panic(err)
	}
	cert, err := ioutil.ReadFile("../secrets/terraform-cert.pem")
	if err != nil {
		panic(err)
	}
	key, err := ioutil.ReadFile("../secrets/terraform.pem")
	if err != nil {
		panic(err)
	}

	raw := map[string]interface{}{
		"bootstrap_servers": bootstrapServers,
		"ca_cert":           string(ca),
		"client_cert":       string(cert),
		"client_key":        string(key),
	}
	return terraform.NewResourceConfigRaw(raw)
}

func datProvider() *schema.Provider {
	log.Println("[INFO] Setting up override for a provider")
	provider := Provider()

	diags := provider.Configure(context.Background(), AccTestProviderConfig())
	if diags.HasError() {
		log.Printf("[ERROR] Could not configure provider %v", diags)
	}

	testProvider = provider
	return provider
}

func accProvider() map[string]*schema.Provider {
	return map[string]*schema.Provider{
		"kafka": datProvider(),
	}
}
