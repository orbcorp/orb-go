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

func TestPlanVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.Versions.New(
		context.TODO(),
		"plan_id",
		orb.PlanVersionNewParams{
			Version: orb.F(int64(0)),
			AddAdjustments: orb.F([]orb.PlanVersionNewParamsAddAdjustment{{
				Adjustment: orb.F[orb.PlanVersionNewParamsAddAdjustmentsAdjustmentUnion](orb.PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			AddPrices: orb.F([]orb.PlanVersionNewParamsAddPrice{{
				AllocationPrice: orb.F(orb.PlanVersionNewParamsAddPricesAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.PlanVersionNewParamsAddPricesAllocationPriceCadenceMonthly),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.PlanVersionNewParamsAddPricesPriceUnion](orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
			}}),
			RemoveAdjustments: orb.F([]orb.PlanVersionNewParamsRemoveAdjustment{{
				AdjustmentID:   orb.F("adjustment_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			RemovePrices: orb.F([]orb.PlanVersionNewParamsRemovePrice{{
				PriceID:        orb.F("price_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			ReplaceAdjustments: orb.F([]orb.PlanVersionNewParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.PlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion](orb.PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
				PlanPhaseOrder:       orb.F(int64(0)),
			}}),
			ReplacePrices: orb.F([]orb.PlanVersionNewParamsReplacePrice{{
				ReplacesPriceID: orb.F("replaces_price_id"),
				AllocationPrice: orb.F(orb.PlanVersionNewParamsReplacePricesAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.PlanVersionNewParamsReplacePricesAllocationPriceCadenceMonthly),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.PlanVersionNewParamsReplacePricesPriceUnion](orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
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

func TestPlanVersionGet(t *testing.T) {
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
	_, err := client.Plans.Versions.Get(
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
