// File generated from our OpenAPI spec by Stainless.

package orb_test

import (
	"context"
	"os"
	"testing"

	"github.com/orbcorp/orb-go"
	"github.com/orbcorp/orb-go/internal/testutil"
	"github.com/orbcorp/orb-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := orb.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	customer, err := client.Customers.New(context.TODO(), orb.CustomerNewParams{
		Email: orb.F("string"),
		Name:  orb.F("string"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", customer)
}
