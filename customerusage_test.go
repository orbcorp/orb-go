// File generated from our OpenAPI spec by Stainless.

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

func TestCustomerUsageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Usage.Update(
		context.TODO(),
		"string",
		orb.CustomerUsageUpdateParams{
			EventName:          orb.F("string"),
			Properties:         orb.F[any](map[string]interface{}{}),
			Timestamp:          orb.F(time.Now()),
			TimeframeEnd:       orb.F(time.Now()),
			TimeframeStart:     orb.F(time.Now()),
			CustomerID:         orb.F("string"),
			ExternalCustomerID: orb.F("string"),
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

func TestCustomerUsageUpdateByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Usage.UpdateByExternalID(
		context.TODO(),
		"string",
		orb.CustomerUsageUpdateByExternalIDParams{
			EventName:          orb.F("string"),
			Properties:         orb.F[any](map[string]interface{}{}),
			Timestamp:          orb.F(time.Now()),
			TimeframeEnd:       orb.F(time.Now()),
			TimeframeStart:     orb.F(time.Now()),
			CustomerID:         orb.F("string"),
			ExternalCustomerID: orb.F("string"),
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
