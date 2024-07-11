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

func TestCustomerCostListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Costs.List(
		context.TODO(),
		"customer_id",
		orb.CustomerCostListParams{
			TimeframeEnd:   orb.F(time.Now()),
			TimeframeStart: orb.F(time.Now()),
			ViewMode:       orb.F(orb.CustomerCostListParamsViewModePeriodic),
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

func TestCustomerCostListByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Costs.ListByExternalID(
		context.TODO(),
		"external_customer_id",
		orb.CustomerCostListByExternalIDParams{
			TimeframeEnd:   orb.F(time.Now()),
			TimeframeStart: orb.F(time.Now()),
			ViewMode:       orb.F(orb.CustomerCostListByExternalIDParamsViewModePeriodic),
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
