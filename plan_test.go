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
	"github.com/orbcorp/orb-go/shared"
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
		Prices: orb.F([]orb.PlanNewParamsPrice{{
			AllocationPrice: orb.F(shared.NewAllocationPriceParam{
				Amount:   orb.F("10.00"),
				Cadence:  orb.F(shared.NewAllocationPriceCadenceMonthly),
				Currency: orb.F("USD"),
				CustomExpiration: orb.F(shared.CustomExpirationParam{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(shared.CustomExpirationDurationUnitDay),
				}),
				ExpiresAtEndOfCadence: orb.F(true),
			}),
			PlanPhaseOrder: orb.F(int64(0)),
			Price: orb.F[orb.PlanNewParamsPricesPriceUnion](shared.NewPlanUnitPriceParam{
				Cadence:   orb.F(shared.NewPlanUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(shared.NewPlanUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(shared.UnitConfigParam{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(shared.NewBillingCycleConfigurationParam{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(shared.NewBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate: orb.F(0.000000),
				ConversionRateConfig: orb.F[shared.NewPlanUnitPriceConversionRateConfigUnionParam](shared.UnitConversionRateConfigParam{
					ConversionRateType: orb.F(shared.UnitConversionRateConfigConversionRateTypeUnit),
					UnitConfig: orb.F(shared.ConversionRateUnitConfigParam{
						UnitAmount: orb.F("unit_amount"),
					}),
				}),
				Currency: orb.F("currency"),
				DimensionalPriceConfiguration: orb.F(shared.NewDimensionalPriceConfigurationParam{
					DimensionValues:                 orb.F([]string{"string"}),
					DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
					ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
				}),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("x"),
				InvoicingCycleConfiguration: orb.F(shared.NewBillingCycleConfigurationParam{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(shared.NewBillingCycleConfigurationDurationUnitDay),
				}),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
				ReferenceID: orb.F("reference_id"),
			}),
		}}),
		Adjustments: orb.F([]orb.PlanNewParamsAdjustment{{
			Adjustment: orb.F[orb.PlanNewParamsAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
				AdjustmentType:     orb.F(shared.NewPercentageDiscountAdjustmentTypePercentageDiscount),
				PercentageDiscount: orb.F(0.000000),
				AppliesToAll:       orb.F(shared.NewPercentageDiscountAppliesToAllTrue),
				AppliesToItemIDs:   orb.F([]string{"item_1", "item_2"}),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				Currency:           orb.F("currency"),
				Filters: orb.F([]shared.NewPercentageDiscountFilterParam{{
					Field:    orb.F(shared.NewPercentageDiscountFiltersFieldPriceID),
					Operator: orb.F(shared.NewPercentageDiscountFiltersOperatorIncludes),
					Values:   orb.F([]string{"string"}),
				}}),
				IsInvoiceLevel: orb.F(true),
				PriceType:      orb.F(shared.NewPercentageDiscountPriceTypeUsage),
			}),
			PlanPhaseOrder: orb.F(int64(0)),
		}}),
		DefaultInvoiceMemo: orb.F("default_invoice_memo"),
		ExternalPlanID:     orb.F("external_plan_id"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		NetTerms: orb.F(int64(0)),
		PlanPhases: orb.F([]orb.PlanNewParamsPlanPhase{{
			Order:                          orb.F(int64(0)),
			AlignBillingWithPhaseStartDate: orb.F(true),
			Duration:                       orb.F(int64(1)),
			DurationUnit:                   orb.F(orb.PlanNewParamsPlanPhasesDurationUnitDaily),
		}}),
		Status: orb.F(orb.PlanNewParamsStatusActive),
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
