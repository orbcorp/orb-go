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

func TestSubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.New(context.TODO(), orb.SubscriptionNewParams{
		AddAdjustments: orb.F([]orb.SubscriptionNewParamsAddAdjustment{{
			Adjustment: orb.F[orb.SubscriptionNewParamsAddAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
				AdjustmentType:     orb.F(shared.NewPercentageDiscountAdjustmentTypePercentageDiscount),
				PercentageDiscount: orb.F(0.000000),
				AppliesToAll:       orb.F(shared.NewPercentageDiscountAppliesToAllTrue),
				AppliesToItemIDs:   orb.F([]string{"item_1", "item_2"}),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				Currency:           orb.F("currency"),
				Filters: orb.F([]shared.TransformPriceFilterParam{{
					Field:    orb.F(shared.TransformPriceFilterFieldPriceID),
					Operator: orb.F(shared.TransformPriceFilterOperatorIncludes),
					Values:   orb.F([]string{"string"}),
				}}),
				IsInvoiceLevel: orb.F(true),
				PriceType:      orb.F(shared.NewPercentageDiscountPriceTypeUsage),
			}),
			EndDate:        orb.F(time.Now()),
			PlanPhaseOrder: orb.F(int64(0)),
			StartDate:      orb.F(time.Now()),
		}}),
		AddPrices: orb.F([]orb.SubscriptionNewParamsAddPrice{{
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
			Discounts: orb.F([]orb.DiscountOverrideParam{{
				DiscountType:       orb.F(orb.DiscountOverrideDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			EndDate:         orb.F(time.Now()),
			ExternalPriceID: orb.F("external_price_id"),
			MaximumAmount:   orb.F("1.23"),
			MinimumAmount:   orb.F("1.23"),
			PlanPhaseOrder:  orb.F(int64(0)),
			Price: orb.F[orb.SubscriptionNewParamsAddPricesPriceUnion](orb.NewSubscriptionUnitPriceParam{
				Cadence:   orb.F(orb.NewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.NewSubscriptionUnitPriceModelTypeUnit),
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
				ConversionRateConfig: orb.F[orb.NewSubscriptionUnitPriceConversionRateConfigUnionParam](shared.UnitConversionRateConfigParam{
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
			PriceID:   orb.F("h74gfhdjvn7ujokd"),
			StartDate: orb.F(time.Now()),
		}}),
		AlignBillingWithSubscriptionStartDate: orb.F(true),
		AutoCollection:                        orb.F(true),
		AwsRegion:                             orb.F("aws_region"),
		BillingCycleAnchorConfiguration: orb.F(shared.BillingCycleAnchorConfigurationParam{
			Day:   orb.F(int64(1)),
			Month: orb.F(int64(1)),
			Year:  orb.F(int64(0)),
		}),
		CouponRedemptionCode:           orb.F("coupon_redemption_code"),
		CreditsOverageRate:             orb.F(0.000000),
		Currency:                       orb.F("currency"),
		CustomerID:                     orb.F("customer_id"),
		DefaultInvoiceMemo:             orb.F("default_invoice_memo"),
		EndDate:                        orb.F(time.Now()),
		ExternalCustomerID:             orb.F("external_customer_id"),
		ExternalMarketplace:            orb.F(orb.SubscriptionNewParamsExternalMarketplaceGoogle),
		ExternalMarketplaceReportingID: orb.F("external_marketplace_reporting_id"),
		ExternalPlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
		Filter:                         orb.F("my_property > 100 AND my_other_property = 'bar'"),
		InitialPhaseOrder:              orb.F(int64(2)),
		InvoicingThreshold:             orb.F("10.00"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		Name:                   orb.F("name"),
		NetTerms:               orb.F(int64(0)),
		PerCreditOverageAmount: orb.F(0.000000),
		PlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
		PlanVersionNumber:      orb.F(int64(0)),
		PriceOverrides:         orb.F([]interface{}{map[string]interface{}{}}),
		RemoveAdjustments: orb.F([]orb.SubscriptionNewParamsRemoveAdjustment{{
			AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
		}}),
		RemovePrices: orb.F([]orb.SubscriptionNewParamsRemovePrice{{
			ExternalPriceID: orb.F("external_price_id"),
			PriceID:         orb.F("h74gfhdjvn7ujokd"),
		}}),
		ReplaceAdjustments: orb.F([]orb.SubscriptionNewParamsReplaceAdjustment{{
			Adjustment: orb.F[orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
				AdjustmentType:     orb.F(shared.NewPercentageDiscountAdjustmentTypePercentageDiscount),
				PercentageDiscount: orb.F(0.000000),
				AppliesToAll:       orb.F(shared.NewPercentageDiscountAppliesToAllTrue),
				AppliesToItemIDs:   orb.F([]string{"item_1", "item_2"}),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				Currency:           orb.F("currency"),
				Filters: orb.F([]shared.TransformPriceFilterParam{{
					Field:    orb.F(shared.TransformPriceFilterFieldPriceID),
					Operator: orb.F(shared.TransformPriceFilterOperatorIncludes),
					Values:   orb.F([]string{"string"}),
				}}),
				IsInvoiceLevel: orb.F(true),
				PriceType:      orb.F(shared.NewPercentageDiscountPriceTypeUsage),
			}),
			ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
		}}),
		ReplacePrices: orb.F([]orb.SubscriptionNewParamsReplacePrice{{
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
			}),
			Discounts: orb.F([]orb.DiscountOverrideParam{{
				DiscountType:       orb.F(orb.DiscountOverrideDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			ExternalPriceID:    orb.F("external_price_id"),
			FixedPriceQuantity: orb.F(2.000000),
			MaximumAmount:      orb.F("1.23"),
			MinimumAmount:      orb.F("1.23"),
			Price: orb.F[orb.SubscriptionNewParamsReplacePricesPriceUnion](orb.NewSubscriptionUnitPriceParam{
				Cadence:   orb.F(orb.NewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.NewSubscriptionUnitPriceModelTypeUnit),
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
				ConversionRateConfig: orb.F[orb.NewSubscriptionUnitPriceConversionRateConfigUnionParam](shared.UnitConversionRateConfigParam{
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
			PriceID: orb.F("h74gfhdjvn7ujokd"),
		}}),
		StartDate:         orb.F(time.Now()),
		TrialDurationDays: orb.F(int64(0)),
		UsageCustomerIDs:  orb.F([]string{"string"}),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.Update(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionUpdateParams{
			AutoCollection:     orb.F(true),
			DefaultInvoiceMemo: orb.F("default_invoice_memo"),
			InvoicingThreshold: orb.F("10.00"),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			NetTerms: orb.F(int64(0)),
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

func TestSubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.List(context.TODO(), orb.SubscriptionListParams{
		CreatedAtGt:        orb.F(time.Now()),
		CreatedAtGte:       orb.F(time.Now()),
		CreatedAtLt:        orb.F(time.Now()),
		CreatedAtLte:       orb.F(time.Now()),
		Cursor:             orb.F("cursor"),
		CustomerID:         orb.F([]string{"string"}),
		ExternalCustomerID: orb.F([]string{"string"}),
		ExternalPlanID:     orb.F("external_plan_id"),
		Limit:              orb.F(int64(1)),
		PlanID:             orb.F("plan_id"),
		Status:             orb.F(orb.SubscriptionListParamsStatusActive),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.Cancel(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionCancelParams{
			CancelOption:             orb.F(orb.SubscriptionCancelParamsCancelOptionEndOfSubscriptionTerm),
			AllowInvoiceCreditOrVoid: orb.F(true),
			CancellationDate:         orb.F(time.Now()),
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

func TestSubscriptionFetch(t *testing.T) {
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
	_, err := client.Subscriptions.Fetch(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionFetchCostsWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.FetchCosts(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionFetchCostsParams{
			Currency:       orb.F("currency"),
			TimeframeEnd:   orb.F(time.Now()),
			TimeframeStart: orb.F(time.Now()),
			ViewMode:       orb.F(orb.SubscriptionFetchCostsParamsViewModePeriodic),
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

func TestSubscriptionFetchScheduleWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.FetchSchedule(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionFetchScheduleParams{
			Cursor:       orb.F("cursor"),
			Limit:        orb.F(int64(1)),
			StartDateGt:  orb.F(time.Now()),
			StartDateGte: orb.F(time.Now()),
			StartDateLt:  orb.F(time.Now()),
			StartDateLte: orb.F(time.Now()),
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

func TestSubscriptionFetchUsageWithOptionalParams(t *testing.T) {
	t.Skip("Incorrect example breaks Prism")
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
	_, err := client.Subscriptions.FetchUsage(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionFetchUsageParams{
			BillableMetricID:     orb.F("billable_metric_id"),
			FirstDimensionKey:    orb.F("first_dimension_key"),
			FirstDimensionValue:  orb.F("first_dimension_value"),
			Granularity:          orb.F(orb.SubscriptionFetchUsageParamsGranularityDay),
			GroupBy:              orb.F("group_by"),
			SecondDimensionKey:   orb.F("second_dimension_key"),
			SecondDimensionValue: orb.F("second_dimension_value"),
			TimeframeEnd:         orb.F(time.Now()),
			TimeframeStart:       orb.F(time.Now()),
			ViewMode:             orb.F(orb.SubscriptionFetchUsageParamsViewModePeriodic),
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

func TestSubscriptionPriceIntervalsWithOptionalParams(t *testing.T) {
	t.Skip("Incorrect example breaks Prism")
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
	_, err := client.Subscriptions.PriceIntervals(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionPriceIntervalsParams{
			Add: orb.F([]orb.SubscriptionPriceIntervalsParamsAdd{{
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddStartDateUnion](shared.UnionTime(time.Now())),
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
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}}),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				ExternalPriceID: orb.F("external_price_id"),
				Filter:          orb.F("my_property > 100 AND my_other_property = 'bar'"),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				MaximumAmount: orb.F(0.000000),
				MinimumAmount: orb.F(0.000000),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](shared.NewFloatingUnitPriceParam{
					Cadence:   orb.F(shared.NewFloatingUnitPriceCadenceAnnual),
					Currency:  orb.F("currency"),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(shared.NewFloatingUnitPriceModelTypeUnit),
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
					ConversionRateConfig: orb.F[shared.NewFloatingUnitPriceConversionRateConfigUnionParam](shared.UnitConversionRateConfigParam{
						ConversionRateType: orb.F(shared.UnitConversionRateConfigConversionRateTypeUnit),
						UnitConfig: orb.F(shared.ConversionRateUnitConfigParam{
							UnitAmount: orb.F("unit_amount"),
						}),
					}),
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
				}),
				PriceID:          orb.F("h74gfhdjvn7ujokd"),
				UsageCustomerIDs: orb.F([]string{"string"}),
			}}),
			AddAdjustments: orb.F([]orb.SubscriptionPriceIntervalsParamsAddAdjustment{{
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
					AdjustmentType:     orb.F(shared.NewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToAll:       orb.F(shared.NewPercentageDiscountAppliesToAllTrue),
					AppliesToItemIDs:   orb.F([]string{"item_1", "item_2"}),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					Currency:           orb.F("currency"),
					Filters: orb.F([]shared.TransformPriceFilterParam{{
						Field:    orb.F(shared.TransformPriceFilterFieldPriceID),
						Operator: orb.F(shared.TransformPriceFilterOperatorIncludes),
						Values:   orb.F([]string{"string"}),
					}}),
					IsInvoiceLevel: orb.F(true),
					PriceType:      orb.F(shared.NewPercentageDiscountPriceTypeUsage),
				}),
				AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
				EndDate:      orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}}),
			AllowInvoiceCreditOrVoid: orb.F(true),
			Edit: orb.F([]orb.SubscriptionPriceIntervalsParamsEdit{{
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				BillingCycleDay: orb.F(int64(0)),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				Filter:          orb.F("my_property > 100 AND my_other_property = 'bar'"),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				StartDate:        orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
				UsageCustomerIDs: orb.F([]string{"string"}),
			}}),
			EditAdjustments: orb.F([]orb.SubscriptionPriceIntervalsParamsEditAdjustment{{
				AdjustmentIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				EndDate:              orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
				StartDate:            orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
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

func TestSubscriptionRedeemCouponWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.RedeemCoupon(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionRedeemCouponParams{
			ChangeOption:             orb.F(orb.SubscriptionRedeemCouponParamsChangeOptionRequestedDate),
			AllowInvoiceCreditOrVoid: orb.F(true),
			ChangeDate:               orb.F(time.Now()),
			CouponID:                 orb.F("coupon_id"),
			CouponRedemptionCode:     orb.F("coupon_redemption_code"),
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

func TestSubscriptionSchedulePlanChangeWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.SchedulePlanChange(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionSchedulePlanChangeParams{
			ChangeOption: orb.F(orb.SubscriptionSchedulePlanChangeParamsChangeOptionRequestedDate),
			AddAdjustments: orb.F([]orb.SubscriptionSchedulePlanChangeParamsAddAdjustment{{
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
					AdjustmentType:     orb.F(shared.NewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToAll:       orb.F(shared.NewPercentageDiscountAppliesToAllTrue),
					AppliesToItemIDs:   orb.F([]string{"item_1", "item_2"}),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					Currency:           orb.F("currency"),
					Filters: orb.F([]shared.TransformPriceFilterParam{{
						Field:    orb.F(shared.TransformPriceFilterFieldPriceID),
						Operator: orb.F(shared.TransformPriceFilterOperatorIncludes),
						Values:   orb.F([]string{"string"}),
					}}),
					IsInvoiceLevel: orb.F(true),
					PriceType:      orb.F(shared.NewPercentageDiscountPriceTypeUsage),
				}),
				EndDate:        orb.F(time.Now()),
				PlanPhaseOrder: orb.F(int64(0)),
				StartDate:      orb.F(time.Now()),
			}}),
			AddPrices: orb.F([]orb.SubscriptionSchedulePlanChangeParamsAddPrice{{
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
				Discounts: orb.F([]orb.DiscountOverrideParam{{
					DiscountType:       orb.F(orb.DiscountOverrideDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				EndDate:         orb.F(time.Now()),
				ExternalPriceID: orb.F("external_price_id"),
				MaximumAmount:   orb.F("1.23"),
				MinimumAmount:   orb.F("1.23"),
				PlanPhaseOrder:  orb.F(int64(0)),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceUnion](orb.NewSubscriptionUnitPriceParam{
					Cadence:   orb.F(orb.NewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.NewSubscriptionUnitPriceModelTypeUnit),
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
					ConversionRateConfig: orb.F[orb.NewSubscriptionUnitPriceConversionRateConfigUnionParam](shared.UnitConversionRateConfigParam{
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
				PriceID:   orb.F("h74gfhdjvn7ujokd"),
				StartDate: orb.F(time.Now()),
			}}),
			AlignBillingWithPlanChangeDate: orb.F(true),
			AutoCollection:                 orb.F(true),
			BillingCycleAlignment:          orb.F(orb.SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentUnchanged),
			BillingCycleAnchorConfiguration: orb.F(shared.BillingCycleAnchorConfigurationParam{
				Day:   orb.F(int64(1)),
				Month: orb.F(int64(1)),
				Year:  orb.F(int64(0)),
			}),
			ChangeDate:             orb.F(time.Now()),
			CouponRedemptionCode:   orb.F("coupon_redemption_code"),
			CreditsOverageRate:     orb.F(0.000000),
			DefaultInvoiceMemo:     orb.F("default_invoice_memo"),
			ExternalPlanID:         orb.F("ZMwNQefe7J3ecf7W"),
			Filter:                 orb.F("my_property > 100 AND my_other_property = 'bar'"),
			InitialPhaseOrder:      orb.F(int64(2)),
			InvoicingThreshold:     orb.F("10.00"),
			NetTerms:               orb.F(int64(0)),
			PerCreditOverageAmount: orb.F(0.000000),
			PlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
			PlanVersionNumber:      orb.F(int64(0)),
			PriceOverrides:         orb.F([]interface{}{map[string]interface{}{}}),
			RemoveAdjustments: orb.F([]orb.SubscriptionSchedulePlanChangeParamsRemoveAdjustment{{
				AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
			}}),
			RemovePrices: orb.F([]orb.SubscriptionSchedulePlanChangeParamsRemovePrice{{
				ExternalPriceID: orb.F("external_price_id"),
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
			}}),
			ReplaceAdjustments: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion](shared.NewPercentageDiscountParam{
					AdjustmentType:     orb.F(shared.NewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
					AppliesToAll:       orb.F(shared.NewPercentageDiscountAppliesToAllTrue),
					AppliesToItemIDs:   orb.F([]string{"item_1", "item_2"}),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					Currency:           orb.F("currency"),
					Filters: orb.F([]shared.TransformPriceFilterParam{{
						Field:    orb.F(shared.TransformPriceFilterFieldPriceID),
						Operator: orb.F(shared.TransformPriceFilterOperatorIncludes),
						Values:   orb.F([]string{"string"}),
					}}),
					IsInvoiceLevel: orb.F(true),
					PriceType:      orb.F(shared.NewPercentageDiscountPriceTypeUsage),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
			}}),
			ReplacePrices: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplacePrice{{
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
				}),
				Discounts: orb.F([]orb.DiscountOverrideParam{{
					DiscountType:       orb.F(orb.DiscountOverrideDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(2.000000),
				MaximumAmount:      orb.F("1.23"),
				MinimumAmount:      orb.F("1.23"),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion](orb.NewSubscriptionUnitPriceParam{
					Cadence:   orb.F(orb.NewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.NewSubscriptionUnitPriceModelTypeUnit),
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
					ConversionRateConfig: orb.F[orb.NewSubscriptionUnitPriceConversionRateConfigUnionParam](shared.UnitConversionRateConfigParam{
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
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}}),
			TrialDurationDays: orb.F(int64(0)),
			UsageCustomerIDs:  orb.F([]string{"string"}),
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

func TestSubscriptionTriggerPhaseWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.TriggerPhase(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionTriggerPhaseParams{
			AllowInvoiceCreditOrVoid: orb.F(true),
			EffectiveDate:            orb.F(time.Now()),
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

func TestSubscriptionUnscheduleCancellation(t *testing.T) {
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
	_, err := client.Subscriptions.UnscheduleCancellation(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUnscheduleFixedFeeQuantityUpdates(t *testing.T) {
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
	_, err := client.Subscriptions.UnscheduleFixedFeeQuantityUpdates(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionUnscheduleFixedFeeQuantityUpdatesParams{
			PriceID: orb.F("price_id"),
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

func TestSubscriptionUnschedulePendingPlanChanges(t *testing.T) {
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
	_, err := client.Subscriptions.UnschedulePendingPlanChanges(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUpdateFixedFeeQuantityWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.UpdateFixedFeeQuantity(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionUpdateFixedFeeQuantityParams{
			PriceID:                  orb.F("price_id"),
			Quantity:                 orb.F(0.000000),
			AllowInvoiceCreditOrVoid: orb.F(true),
			ChangeOption:             orb.F(orb.SubscriptionUpdateFixedFeeQuantityParamsChangeOptionImmediate),
			EffectiveDate:            orb.F(time.Now()),
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

func TestSubscriptionUpdateTrialWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.UpdateTrial(
		context.TODO(),
		"subscription_id",
		orb.SubscriptionUpdateTrialParams{
			TrialEndDate: orb.F[orb.SubscriptionUpdateTrialParamsTrialEndDateUnion](shared.UnionTime(time.Now())),
			Shift:        orb.F(true),
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
