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
		Currency: orb.F("currency"),
		Name:     orb.F("name"),
		Prices: orb.F([]orb.PlanNewParamsPriceUnion{orb.PlanNewParamsPricesNewPlanUnitPrice{
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			ExternalPriceID:    orb.F("external_price_id"),
			Name:               orb.F("Annual fee"),
			BillableMetricID:   orb.F("billable_metric_id"),
			ItemID:             orb.F("item_id"),
			BilledInAdvance:    orb.F(true),
			FixedPriceQuantity: orb.F(0.000000),
			InvoiceGroupingKey: orb.F("invoice_grouping_key"),
			Cadence:            orb.F(orb.PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual),
			ConversionRate:     orb.F(0.000000),
			ModelType:          orb.F(orb.PlanNewParamsPricesNewPlanUnitPriceModelTypeUnit),
			UnitConfig: orb.F(orb.PlanNewParamsPricesNewPlanUnitPriceUnitConfig{
				UnitAmount: orb.F("unit_amount"),
			}),
			Currency: orb.F("currency"),
		}}),
		DefaultInvoiceMemo: orb.F("default_invoice_memo"),
		ExternalPlanID:     orb.F("external_plan_id"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		NetTerms: orb.F(int64(0)),
		Status:   orb.F(orb.PlanNewParamsStatusActive),
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
		"plan_id",
		orb.PlanUpdateParams{
			ExternalPlanID: orb.F("external_plan_id"),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
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
		Cursor:       orb.F("cursor"),
		Limit:        orb.F(int64(1)),
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
	_, err := client.Plans.Fetch(context.TODO(), "plan_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
