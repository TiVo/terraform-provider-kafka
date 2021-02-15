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
var testBootstrapServers []string = bootstrapServersFromEnv()

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

	ca, _ := ioutil.ReadFile("../secrets/ca.crt")
	cert, _ := ioutil.ReadFile("../secrets/terraform-cert.pem")
	key, _ := ioutil.ReadFile("../secrets/terraform.pem")

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

func bootstrapServersFromEnv() []string {
	fromEnv := strings.Split(os.Getenv("KAFKA_BOOTSTRAP_SERVER"), ",")
	fromEnv = nonEmptyAndTrimmed(fromEnv)

	if len(fromEnv) == 0 {
		fromEnv = []string{"localhost:9092"}
	}

	bootstrapServers := make([]string, 0)
	for _, bs := range fromEnv {
		if bs != "" {
			bootstrapServers = append(bootstrapServers, bs)
		}
	}

	return bootstrapServers
}

func nonEmptyAndTrimmed(bootstrapServers []string) []string {
	wellFormed := make([]string, 0)

	for _, bs := range bootstrapServers {
		trimmed := strings.TrimSpace(bs)
		if trimmed != "" {
			wellFormed = append(wellFormed, trimmed)
		}
	}

	return wellFormed
}
