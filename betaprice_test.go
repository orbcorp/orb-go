// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/orbcorp/orb-go"
	"github.com/orbcorp/orb-go/internal/testutil"
	"github.com/orbcorp/orb-go/option"
)

func TestBetaPriceEvaluateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Price.Evaluate(
		context.TODO(),
		"string",
		orb.BetaPriceEvaluateParams{
			TimeframeEnd:       orb.F(time.Now()),
			TimeframeStart:     orb.F(time.Now()),
			CustomerID:         orb.F("string"),
			ExternalCustomerID: orb.F("string"),
			Filter:             orb.F("my_numeric_property > 100 AND my_other_property = 'bar'"),
			GroupingKeys:       orb.F([]string{"case when my_event_type = 'foo' then true else false end"}),
		},
	)
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
