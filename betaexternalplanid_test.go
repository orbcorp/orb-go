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
	"github.com/orbcorp/orb-go/shared"
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
				Adjustment: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
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
			AddPrices: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsAddPrice{{
				AllocationPrice: orb.F(shared.NewAllocationPriceParam{
					Amount:   orb.F("10.00"),
					Cadence:  orb.F(shared.NewAllocationPriceCadenceMonthly),
					Currency: orb.F("USD"),
					CustomExpiration: orb.F(shared.CustomExpirationParam{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(shared.CustomExpirationDurationUnitDay),
					}),
					ExpiresAtEndOfCadence: orb.F(true),
					Filters: orb.F([]shared.NewAllocationPriceFilterParam{{
						Field:    orb.F(shared.NewAllocationPriceFiltersFieldItemID),
						Operator: orb.F(shared.NewAllocationPriceFiltersOperatorIncludes),
						Values:   orb.F([]string{"string"}),
					}}),
					ItemID:           orb.F("item_id"),
					PerUnitCostBasis: orb.F("per_unit_cost_basis"),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion](shared.NewPlanUnitPriceParam{
					Cadence:   orb.F(shared.NewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(shared.NewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(shared.UnitConfigParam{
						UnitAmount: orb.F("unit_amount"),
						Prorated:   orb.F(true),
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
			RemoveAdjustments: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsRemoveAdjustment{{
				AdjustmentID:   orb.F("adjustment_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			RemovePrices: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsRemovePrice{{
				PriceID:        orb.F("price_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			ReplaceAdjustments: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
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
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
				PlanPhaseOrder:       orb.F(int64(0)),
			}}),
			ReplacePrices: orb.F([]orb.BetaExternalPlanIDNewPlanVersionParamsReplacePrice{{
				ReplacesPriceID: orb.F("replaces_price_id"),
				AllocationPrice: orb.F(shared.NewAllocationPriceParam{
					Amount:   orb.F("10.00"),
					Cadence:  orb.F(shared.NewAllocationPriceCadenceMonthly),
					Currency: orb.F("USD"),
					CustomExpiration: orb.F(shared.CustomExpirationParam{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(shared.CustomExpirationDurationUnitDay),
					}),
					ExpiresAtEndOfCadence: orb.F(true),
					Filters: orb.F([]shared.NewAllocationPriceFilterParam{{
						Field:    orb.F(shared.NewAllocationPriceFiltersFieldItemID),
						Operator: orb.F(shared.NewAllocationPriceFiltersOperatorIncludes),
						Values:   orb.F([]string{"string"}),
					}}),
					ItemID:           orb.F("item_id"),
					PerUnitCostBasis: orb.F("per_unit_cost_basis"),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion](shared.NewPlanUnitPriceParam{
					Cadence:   orb.F(shared.NewPlanUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(shared.NewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(shared.UnitConfigParam{
						UnitAmount: orb.F("unit_amount"),
						Prorated:   orb.F(true),
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
