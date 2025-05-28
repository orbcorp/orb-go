// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/orbcorp/orb-go"
	"github.com/orbcorp/orb-go/internal/testutil"
	"github.com/orbcorp/orb-go/option"
)

func TestBetaNewPlanVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.NewPlanVersion(
		context.TODO(),
		"plan_id",
		orb.BetaNewPlanVersionParams{
			Version: orb.F(int64(0)),
			AddAdjustments: orb.F([]orb.BetaNewPlanVersionParamsAddAdjustment{{
				Adjustment: orb.F[orb.BetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion](orb.BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			AddPrices: orb.F([]orb.BetaNewPlanVersionParamsAddPrice{{
				AllocationPrice: orb.F(orb.BetaNewPlanVersionParamsAddPricesAllocationPrice{
					Amount:   orb.F("10.00"),
					Cadence:  orb.F(orb.BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceMonthly),
					Currency: orb.F("USD"),
					CustomExpiration: orb.F(orb.BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitDay),
					}),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.BetaNewPlanVersionParamsAddPricesPriceUnion](orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
			}}),
			RemoveAdjustments: orb.F([]orb.BetaNewPlanVersionParamsRemoveAdjustment{{
				AdjustmentID:   orb.F("adjustment_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			RemovePrices: orb.F([]orb.BetaNewPlanVersionParamsRemovePrice{{
				PriceID:        orb.F("price_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			ReplaceAdjustments: orb.F([]orb.BetaNewPlanVersionParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion](orb.BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
				PlanPhaseOrder:       orb.F(int64(0)),
			}}),
			ReplacePrices: orb.F([]orb.BetaNewPlanVersionParamsReplacePrice{{
				ReplacesPriceID: orb.F("replaces_price_id"),
				AllocationPrice: orb.F(orb.BetaNewPlanVersionParamsReplacePricesAllocationPrice{
					Amount:   orb.F("10.00"),
					Cadence:  orb.F(orb.BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceMonthly),
					Currency: orb.F("USD"),
					CustomExpiration: orb.F(orb.BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitDay),
					}),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.BetaNewPlanVersionParamsReplacePricesPriceUnion](orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
			}}),
			SetAsDefault: orb.F(true),
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

func TestBetaFetchPlanVersion(t *testing.T) {
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
	_, err := client.Beta.FetchPlanVersion(
		context.TODO(),
		"plan_id",
		"version",
	)
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaSetDefaultPlanVersion(t *testing.T) {
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
	_, err := client.Beta.SetDefaultPlanVersion(
		context.TODO(),
		"plan_id",
		orb.BetaSetDefaultPlanVersionParams{
			Version: orb.F(int64(0)),
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
