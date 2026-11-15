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

func TestPlanDraftDiscard(t *testing.T) {
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
	err := client.Plans.Draft.Discard(context.TODO(), "plan_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanDraftFetch(t *testing.T) {
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
	_, err := client.Plans.Draft.Fetch(context.TODO(), "plan_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanDraftPublishWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.Draft.Publish(
		context.TODO(),
		"plan_id",
		orb.PlanDraftPublishParams{
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

func TestPlanDraftSaveWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.Draft.Save(
		context.TODO(),
		"plan_id",
		orb.PlanDraftSaveParams{
			AddAdjustments: orb.F([]orb.PlanDraftSaveParamsAddAdjustment{{
				Adjustment: orb.F[orb.PlanDraftSaveParamsAddAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
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
			AddPrices: orb.F([]orb.PlanDraftSaveParamsAddPrice{{
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
					ItemID:        orb.F("item_id"),
					LicenseTypeID: orb.F("license_type_id"),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					PerUnitCostBasis: orb.F("per_unit_cost_basis"),
				}),
				LicenseAllocationPrice: orb.F(orb.PlanDraftSaveParamsAddPricesLicenseAllocationPrice{
					Cadence: orb.F(orb.PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceOneTime),
					ItemID:  orb.F("item_id"),
					LicenseAllocations: orb.F([]orb.PlanDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation{{
						Amount:          orb.F("amount"),
						Currency:        orb.F("currency"),
						WriteOffOverage: orb.F(true),
					}}),
					ModelType: orb.F(orb.PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(shared.UnitConfigParam{
						UnitAmount:    orb.F("unit_amount"),
						Prorated:      orb.F(true),
						ScalingFactor: orb.F(0.000000),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(shared.NewBillingCycleConfigurationParam{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(shared.NewBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					ConversionRateConfig: orb.F[orb.PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion](shared.UnitConversionRateConfigParam{
						ConversionRateType: orb.F(shared.UnitConversionRateConfigConversionRateTypeUnit),
						UnitConfig: orb.F(shared.ConversionRateUnitConfigParam{
							UnitAmount: orb.F("unit_amount"),
						}),
					}),
					Currency: orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(shared.NewDimensionalPriceConfigurationParam{
						DimensionValues:                 orb.F([]string{"x"}),
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
					LicenseTypeID: orb.F("license_type_id"),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.PlanDraftSaveParamsAddPricesPriceUnion](shared.NewPlanUnitPriceParam{
					Cadence:   orb.F(shared.NewPlanUnitPriceCadenceOneTime),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(shared.NewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(shared.UnitConfigParam{
						UnitAmount:    orb.F("unit_amount"),
						Prorated:      orb.F(true),
						ScalingFactor: orb.F(0.000000),
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
						DimensionValues:                 orb.F([]string{"x"}),
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
					LicenseTypeID: orb.F("license_type_id"),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
			}}),
			Description: orb.F("description"),
			RemoveAdjustments: orb.F([]orb.PlanDraftSaveParamsRemoveAdjustment{{
				AdjustmentID:   orb.F("adjustment_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			RemovePrices: orb.F([]orb.PlanDraftSaveParamsRemovePrice{{
				PriceID:        orb.F("price_id"),
				PlanPhaseOrder: orb.F(int64(0)),
			}}),
			ReplaceAdjustments: orb.F([]orb.PlanDraftSaveParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.PlanDraftSaveParamsReplaceAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
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
			ReplacePrices: orb.F([]orb.PlanDraftSaveParamsReplacePrice{{
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
					ItemID:        orb.F("item_id"),
					LicenseTypeID: orb.F("license_type_id"),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					PerUnitCostBasis: orb.F("per_unit_cost_basis"),
				}),
				LicenseAllocationPrice: orb.F(orb.PlanDraftSaveParamsReplacePricesLicenseAllocationPrice{
					Cadence: orb.F(orb.PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceOneTime),
					ItemID:  orb.F("item_id"),
					LicenseAllocations: orb.F([]orb.PlanDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation{{
						Amount:          orb.F("amount"),
						Currency:        orb.F("currency"),
						WriteOffOverage: orb.F(true),
					}}),
					ModelType: orb.F(orb.PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(shared.UnitConfigParam{
						UnitAmount:    orb.F("unit_amount"),
						Prorated:      orb.F(true),
						ScalingFactor: orb.F(0.000000),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(shared.NewBillingCycleConfigurationParam{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(shared.NewBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate: orb.F(0.000000),
					ConversionRateConfig: orb.F[orb.PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion](shared.UnitConversionRateConfigParam{
						ConversionRateType: orb.F(shared.UnitConversionRateConfigConversionRateTypeUnit),
						UnitConfig: orb.F(shared.ConversionRateUnitConfigParam{
							UnitAmount: orb.F("unit_amount"),
						}),
					}),
					Currency: orb.F("currency"),
					DimensionalPriceConfiguration: orb.F(shared.NewDimensionalPriceConfigurationParam{
						DimensionValues:                 orb.F([]string{"x"}),
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
					LicenseTypeID: orb.F("license_type_id"),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PlanPhaseOrder: orb.F(int64(0)),
				Price: orb.F[orb.PlanDraftSaveParamsReplacePricesPriceUnion](shared.NewPlanUnitPriceParam{
					Cadence:   orb.F(shared.NewPlanUnitPriceCadenceOneTime),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(shared.NewPlanUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(shared.UnitConfigParam{
						UnitAmount:    orb.F("unit_amount"),
						Prorated:      orb.F(true),
						ScalingFactor: orb.F(0.000000),
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
						DimensionValues:                 orb.F([]string{"x"}),
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
					LicenseTypeID: orb.F("license_type_id"),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
			}}),
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
