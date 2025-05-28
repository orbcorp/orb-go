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

func TestPlanExternalPlanIDVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.ExternalPlanID.Versions.New(
		context.TODO(),
		"external_plan_id",
		orb.PlanExternalPlanIDVersionNewParams{
			Version: orb.F(int64(0)),
			AddAdjustments: orb.F([]orb.PlanExternalPlanIDVersionNewParamsAddAdjustment{{
				Adjustment: orb.F[orb.PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion](orb.PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			AddPrices: orb.F([]orb.PlanExternalPlanIDVersionNewParamsAddPrice{{
				AllocationPrice: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceMonthly),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceUnion](orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
			}}),
			RemoveAdjustments: orb.F([]orb.PlanExternalPlanIDVersionNewParamsRemoveAdjustment{{
				AdjustmentID:   orb.F("adjustment_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			RemovePrices: orb.F([]orb.PlanExternalPlanIDVersionNewParamsRemovePrice{{
				PriceID:        orb.F("price_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			ReplaceAdjustments: orb.F([]orb.PlanExternalPlanIDVersionNewParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion](orb.PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					IsInvoiceLevel:     orb.F(true),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
				PlanPhaseOrder:       orb.F(int64(0)),
			}}),
			ReplacePrices: orb.F([]orb.PlanExternalPlanIDVersionNewParamsReplacePrice{{
				ReplacesPriceID: orb.F("replaces_price_id"),
				AllocationPrice: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceMonthly),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion](orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPrice{
					Cadence:   orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					Currency:       orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration{
						DimensionValues:                 orb.F([]string{"string"}),
						DimensionalPriceGroupID:         orb.F("dimensional_price_group_id"),
						ExternalDimensionalPriceGroupID: orb.F("external_dimensional_price_group_id"),
					}),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("x"),
					InvoicingCycleConfiguration: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay),
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

func TestPlanExternalPlanIDVersionGet(t *testing.T) {
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
	_, err := client.Plans.ExternalPlanID.Versions.Get(
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
