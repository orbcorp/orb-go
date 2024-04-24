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

func TestPriceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.New(context.TODO(), orb.PriceNewParamsNewFloatingUnitPrice{
		Cadence:   orb.F(orb.PriceNewParamsNewFloatingUnitPriceCadenceAnnual),
		Currency:  orb.F("string"),
		ItemID:    orb.F("string"),
		ModelType: orb.F(orb.PriceNewParamsNewFloatingUnitPriceModelTypeUnit),
		Name:      orb.F("Annual fee"),
		UnitConfig: orb.F(orb.PriceNewParamsNewFloatingUnitPriceUnitConfig{
			UnitAmount: orb.F("string"),
		}),
		BillableMetricID:   orb.F("string"),
		BilledInAdvance:    orb.F(true),
		ExternalPriceID:    orb.F("string"),
		FixedPriceQuantity: orb.F(0.000000),
		InvoiceGroupingKey: orb.F("string"),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.List(context.TODO(), orb.PriceListParams{
		Cursor: orb.F("string"),
		Limit:  orb.F(int64(0)),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceEvaluateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.Evaluate(
		context.TODO(),
		"string",
		orb.PriceEvaluateParams{
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

func TestPriceFetch(t *testing.T) {
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
	_, err := client.Prices.Fetch(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
