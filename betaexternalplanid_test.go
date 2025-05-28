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

func TestBetaExternalPlanIDNewPlanVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.ExternalPlanID.NewPlanVersion(
		context.TODO(),
		"external_plan_id",
		orb.BetaExternalPlanIDNewPlanVersionParams{
			Version: orb.F(int64(0)),
			AddAdjustments: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsAddAdjustment{{
				Adjustment: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion](orb.BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			AddPrices: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsAddPrice{{
				AllocationPrice: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceMonthly),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion](orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
			}}),
			RemoveAdjustments: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsRemoveAdjustment{{
				AdjustmentID:   orb.F("adjustment_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			RemovePrices: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsRemovePrice{{
				PriceID:        orb.F("price_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			ReplaceAdjustments: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion](orb.BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
				PlanPhaseOrder:       orb.F(int64(0)),
			}}),
			ReplacePrices: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsReplacePrice{{
				ReplacesPriceID: orb.F("replaces_price_id"),
				AllocationPrice: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceMonthly),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion](orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
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

func TestBetaExternalPlanIDFetchPlanVersion(t *testing.T) {
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
	_, err := client.Beta.ExternalPlanID.FetchPlanVersion(
		context.TODO(),
		"external_plan_id",
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

func TestBetaExternalPlanIDSetDefaultPlanVersion(t *testing.T) {
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
	_, err := client.Beta.ExternalPlanID.SetDefaultPlanVersion(
		context.TODO(),
		"external_plan_id",
		orb.BetaExternalPlanIDSetDefaultPlanVersionParams{
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
