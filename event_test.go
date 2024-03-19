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

func TestEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Update(
		context.TODO(),
		"string",
		orb.EventUpdateParams{
			EventName:          orb.F("string"),
			Properties:         orb.F[any](map[string]interface{}{}),
			Timestamp:          orb.F(time.Now()),
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

func TestEventDeprecate(t *testing.T) {
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
	_, err := client.Events.Deprecate(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventIngestWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Ingest(context.TODO(), orb.EventIngestParams{
		Events: orb.F([]orb.EventIngestParamsEvent{{
			CustomerID:         orb.F("string"),
			ExternalCustomerID: orb.F("string"),
			EventName:          orb.F("string"),
			Timestamp:          orb.F(time.Now()),
			Properties:         orb.F[any](map[string]interface{}{}),
			IdempotencyKey:     orb.F("string"),
		}, {
			CustomerID:         orb.F("string"),
			ExternalCustomerID: orb.F("string"),
			EventName:          orb.F("string"),
			Timestamp:          orb.F(time.Now()),
			Properties:         orb.F[any](map[string]interface{}{}),
			IdempotencyKey:     orb.F("string"),
		}, {
			CustomerID:         orb.F("string"),
			ExternalCustomerID: orb.F("string"),
			EventName:          orb.F("string"),
			Timestamp:          orb.F(time.Now()),
			Properties:         orb.F[any](map[string]interface{}{}),
			IdempotencyKey:     orb.F("string"),
		}}),
		BackfillID: orb.F("string"),
		Debug:      orb.F(true),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Search(context.TODO(), orb.EventSearchParams{
		EventIDs:       orb.F([]string{"string"}),
		TimeframeEnd:   orb.F(time.Now()),
		TimeframeStart: orb.F(time.Now()),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
