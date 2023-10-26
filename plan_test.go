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

func TestPlanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.New(context.TODO(), orb.PlanNewParams{
		Currency: orb.F("string"),
		Name:     orb.F("string"),
		Prices: orb.F([]orb.PlanNewParamsPrice{orb.PlanNewParamsPricesNewPlanUnitPrice(orb.PlanNewParamsPricesNewPlanUnitPrice{
			ExternalPriceID:    orb.F("string"),
			Name:               orb.F("Annual fee"),
			BillableMetricID:   orb.F("string"),
			ItemID:             orb.F("string"),
			BilledInAdvance:    orb.F(true),
			FixedPriceQuantity: orb.F(0.000000),
			InvoiceGroupingKey: orb.F("string"),
			Cadence:            orb.F(orb.PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual),
			ModelType:          orb.F(orb.PlanNewParamsPricesNewPlanUnitPriceModelTypeUnit),
			UnitConfig: orb.F(orb.PlanNewParamsPricesNewPlanUnitPriceUnitConfig{
				UnitAmount:    orb.F("string"),
				ScalingFactor: orb.F(0.000000),
			}),
		})}),
		DefaultInvoiceMemo: orb.F("string"),
		ExternalPlanID:     orb.F("string"),
		Metadata:           orb.F[any](map[string]interface{}{}),
		NetTerms:           orb.F(int64(0)),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.Update(
		context.TODO(),
		"string",
		orb.PlanUpdateParams{
			ExternalPlanID: orb.F("string"),
			Metadata:       orb.F[any](map[string]interface{}{}),
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

func TestPlanListWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.List(context.TODO(), orb.PlanListParams{
		CreatedAtGt:  orb.F(time.Now()),
		CreatedAtGte: orb.F(time.Now()),
		CreatedAtLt:  orb.F(time.Now()),
		CreatedAtLte: orb.F(time.Now()),
		Cursor:       orb.F("string"),
		Limit:        orb.F(int64(0)),
		Status:       orb.F(orb.PlanListParamsStatusActive),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanFetch(t *testing.T) {
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
	_, err := client.Plans.Fetch(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
