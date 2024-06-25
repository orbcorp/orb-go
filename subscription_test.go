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
		AlignBillingWithSubscriptionStartDate: orb.F(true),
		AutoCollection:                        orb.F(true),
		AwsRegion:                             orb.F("string"),
		CouponRedemptionCode:                  orb.F("string"),
		CreditsOverageRate:                    orb.F(0.000000),
		CustomerID:                            orb.F("string"),
		DefaultInvoiceMemo:                    orb.F("string"),
		EndDate:                               orb.F(time.Now()),
		ExternalCustomerID:                    orb.F("string"),
		ExternalMarketplace:                   orb.F(orb.SubscriptionNewParamsExternalMarketplaceGoogle),
		ExternalMarketplaceReportingID:        orb.F("string"),
		ExternalPlanID:                        orb.F("ZMwNQefe7J3ecf7W"),
		InitialPhaseOrder:                     orb.F(int64(0)),
		InvoicingThreshold:                    orb.F("string"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		NetTerms:               orb.F(int64(0)),
		PerCreditOverageAmount: orb.F(0.000000),
		PlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
		PriceOverrides: orb.F([]orb.SubscriptionNewParamsPriceOverrideUnion{orb.SubscriptionNewParamsPriceOverridesOverrideUnitPrice{
			ID:             orb.F("string"),
			ModelType:      orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelTypeUnit),
			MinimumAmount:  orb.F("1.23"),
			MaximumAmount:  orb.F("1.23"),
			Currency:       orb.F("string"),
			ConversionRate: orb.F(0.000000),
			Discount: orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscount{
				DiscountType:        orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage),
				PercentageDiscount:  orb.F(0.150000),
				TrialAmountDiscount: orb.F("string"),
				UsageDiscount:       orb.F(0.000000),
				AmountDiscount:      orb.F("string"),
				AppliesToPriceIDs:   orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
			}),
			FixedPriceQuantity: orb.F(2.000000),
			UnitConfig: orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceUnitConfig{
				UnitAmount: orb.F("string"),
			}),
		}, orb.SubscriptionNewParamsPriceOverridesOverrideUnitPrice{
			ID:             orb.F("string"),
			ModelType:      orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelTypeUnit),
			MinimumAmount:  orb.F("1.23"),
			MaximumAmount:  orb.F("1.23"),
			Currency:       orb.F("string"),
			ConversionRate: orb.F(0.000000),
			Discount: orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscount{
				DiscountType:        orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage),
				PercentageDiscount:  orb.F(0.150000),
				TrialAmountDiscount: orb.F("string"),
				UsageDiscount:       orb.F(0.000000),
				AmountDiscount:      orb.F("string"),
				AppliesToPriceIDs:   orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
			}),
			FixedPriceQuantity: orb.F(2.000000),
			UnitConfig: orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceUnitConfig{
				UnitAmount: orb.F("string"),
			}),
		}, orb.SubscriptionNewParamsPriceOverridesOverrideUnitPrice{
			ID:             orb.F("string"),
			ModelType:      orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelTypeUnit),
			MinimumAmount:  orb.F("1.23"),
			MaximumAmount:  orb.F("1.23"),
			Currency:       orb.F("string"),
			ConversionRate: orb.F(0.000000),
			Discount: orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscount{
				DiscountType:        orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage),
				PercentageDiscount:  orb.F(0.150000),
				TrialAmountDiscount: orb.F("string"),
				UsageDiscount:       orb.F(0.000000),
				AmountDiscount:      orb.F("string"),
				AppliesToPriceIDs:   orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
			}),
			FixedPriceQuantity: orb.F(2.000000),
			UnitConfig: orb.F(orb.SubscriptionNewParamsPriceOverridesOverrideUnitPriceUnitConfig{
				UnitAmount: orb.F("string"),
			}),
		}}),
		StartDate: orb.F(time.Now()),
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
		"string",
		orb.SubscriptionUpdateParams{
			AutoCollection:     orb.F(true),
			DefaultInvoiceMemo: orb.F("string"),
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
		Cursor:             orb.F("string"),
		CustomerID:         orb.F("string"),
		ExternalCustomerID: orb.F("string"),
		Limit:              orb.F(int64(1)),
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
		"string",
		orb.SubscriptionCancelParams{
			CancelOption:     orb.F(orb.SubscriptionCancelParamsCancelOptionEndOfSubscriptionTerm),
			CancellationDate: orb.F(time.Now()),
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
	_, err := client.Subscriptions.Fetch(context.TODO(), "string")
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
		"string",
		orb.SubscriptionFetchCostsParams{
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
		"string",
		orb.SubscriptionFetchScheduleParams{
			Cursor:       orb.F("string"),
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
		"string",
		orb.SubscriptionFetchUsageParams{
			BillableMetricID:     orb.F("string"),
			Cursor:               orb.F("string"),
			FirstDimensionKey:    orb.F("string"),
			FirstDimensionValue:  orb.F("string"),
			Granularity:          orb.F(orb.SubscriptionFetchUsageParamsGranularityDay),
			GroupBy:              orb.F("string"),
			Limit:                orb.F(int64(0)),
			SecondDimensionKey:   orb.F("string"),
			SecondDimensionValue: orb.F("string"),
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
		"string",
		orb.SubscriptionPriceIntervalsParams{
			Add: orb.F([]orb.SubscriptionPriceIntervalsParamsAdd{{
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
				ExternalPriceID: orb.F("external_price_id"),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice{
					ExternalPriceID:    orb.F("string"),
					Name:               orb.F("Annual fee"),
					BillableMetricID:   orb.F("string"),
					ItemID:             orb.F("string"),
					BilledInAdvance:    orb.F(true),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("string"),
					Cadence:            orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual),
					ConversionRate:     orb.F(0.000000),
					ModelType:          orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit),
					UnitConfig: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig{
						UnitAmount: orb.F("string"),
					}),
					Currency: orb.F("string"),
				}),
				AllocationPrice: orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPrice{
					Currency:              orb.F("USD"),
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceMonthly),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}}),
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}}),
				MinimumAmount: orb.F(0.000000),
				MaximumAmount: orb.F(0.000000),
			}, {
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
				ExternalPriceID: orb.F("external_price_id"),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice{
					ExternalPriceID:    orb.F("string"),
					Name:               orb.F("Annual fee"),
					BillableMetricID:   orb.F("string"),
					ItemID:             orb.F("string"),
					BilledInAdvance:    orb.F(true),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("string"),
					Cadence:            orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual),
					ConversionRate:     orb.F(0.000000),
					ModelType:          orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit),
					UnitConfig: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig{
						UnitAmount: orb.F("string"),
					}),
					Currency: orb.F("string"),
				}),
				AllocationPrice: orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPrice{
					Currency:              orb.F("USD"),
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceMonthly),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}}),
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}}),
				MinimumAmount: orb.F(0.000000),
				MaximumAmount: orb.F(0.000000),
			}, {
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
				ExternalPriceID: orb.F("external_price_id"),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice{
					ExternalPriceID:    orb.F("string"),
					Name:               orb.F("Annual fee"),
					BillableMetricID:   orb.F("string"),
					ItemID:             orb.F("string"),
					BilledInAdvance:    orb.F(true),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("string"),
					Cadence:            orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual),
					ConversionRate:     orb.F(0.000000),
					ModelType:          orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit),
					UnitConfig: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig{
						UnitAmount: orb.F("string"),
					}),
					Currency: orb.F("string"),
				}),
				AllocationPrice: orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPrice{
					Currency:              orb.F("USD"),
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceMonthly),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}}),
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
					AmountDiscount: orb.F(0.000000),
				}}),
				MinimumAmount: orb.F(0.000000),
				MaximumAmount: orb.F(0.000000),
			}}),
			AddAdjustments: orb.F([]orb.SubscriptionPriceIntervalsParamsAddAdjustment{{
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					AdjustmentType:     orb.F(orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}, {
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					AdjustmentType:     orb.F(orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}, {
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					AdjustmentType:     orb.F(orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					PercentageDiscount: orb.F(0.000000),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}}),
			Edit: orb.F([]orb.SubscriptionPriceIntervalsParamsEdit{{
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				StartDate:       orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}}),
				BillingCycleDay: orb.F(int64(0)),
			}, {
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				StartDate:       orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}}),
				BillingCycleDay: orb.F(int64(0)),
			}, {
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				StartDate:       orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}, {
					Quantity:      orb.F(int64(5)),
					EffectiveDate: orb.F(time.Now()),
				}}),
				BillingCycleDay: orb.F(int64(0)),
			}}),
			EditAdjustments: orb.F([]orb.SubscriptionPriceIntervalsParamsEditAdjustment{{
				AdjustmentIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				StartDate:            orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:              orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}, {
				AdjustmentIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				StartDate:            orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:              orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}, {
				AdjustmentIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				StartDate:            orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:              orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
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
		"string",
		orb.SubscriptionSchedulePlanChangeParams{
			ChangeOption:                   orb.F(orb.SubscriptionSchedulePlanChangeParamsChangeOptionRequestedDate),
			AlignBillingWithPlanChangeDate: orb.F(true),
			BillingCycleAlignment:          orb.F(orb.SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentUnchanged),
			ChangeDate:                     orb.F("2017-07-21T17:32:28Z"),
			CouponRedemptionCode:           orb.F("string"),
			CreditsOverageRate:             orb.F(0.000000),
			ExternalPlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
			InitialPhaseOrder:              orb.F(int64(2)),
			InvoicingThreshold:             orb.F("10.00"),
			PerCreditOverageAmount:         orb.F(0.000000),
			PlanID:                         orb.F("ZMwNQefe7J3ecf7W"),
			PriceOverrides: orb.F([]orb.SubscriptionSchedulePlanChangeParamsPriceOverrideUnion{orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice{
				ID:             orb.F("string"),
				ModelType:      orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelTypeUnit),
				MinimumAmount:  orb.F("1.23"),
				MaximumAmount:  orb.F("1.23"),
				Currency:       orb.F("string"),
				ConversionRate: orb.F(0.000000),
				Discount: orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscount{
					DiscountType:        orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage),
					PercentageDiscount:  orb.F(0.150000),
					TrialAmountDiscount: orb.F("string"),
					UsageDiscount:       orb.F(0.000000),
					AmountDiscount:      orb.F("string"),
					AppliesToPriceIDs:   orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
				}),
				FixedPriceQuantity: orb.F(2.000000),
				UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceUnitConfig{
					UnitAmount: orb.F("string"),
				}),
			}, orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice{
				ID:             orb.F("string"),
				ModelType:      orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelTypeUnit),
				MinimumAmount:  orb.F("1.23"),
				MaximumAmount:  orb.F("1.23"),
				Currency:       orb.F("string"),
				ConversionRate: orb.F(0.000000),
				Discount: orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscount{
					DiscountType:        orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage),
					PercentageDiscount:  orb.F(0.150000),
					TrialAmountDiscount: orb.F("string"),
					UsageDiscount:       orb.F(0.000000),
					AmountDiscount:      orb.F("string"),
					AppliesToPriceIDs:   orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
				}),
				FixedPriceQuantity: orb.F(2.000000),
				UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceUnitConfig{
					UnitAmount: orb.F("string"),
				}),
			}, orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice{
				ID:             orb.F("string"),
				ModelType:      orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelTypeUnit),
				MinimumAmount:  orb.F("1.23"),
				MaximumAmount:  orb.F("1.23"),
				Currency:       orb.F("string"),
				ConversionRate: orb.F(0.000000),
				Discount: orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscount{
					DiscountType:        orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage),
					PercentageDiscount:  orb.F(0.150000),
					TrialAmountDiscount: orb.F("string"),
					UsageDiscount:       orb.F(0.000000),
					AmountDiscount:      orb.F("string"),
					AppliesToPriceIDs:   orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
				}),
				FixedPriceQuantity: orb.F(2.000000),
				UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceUnitConfig{
					UnitAmount: orb.F("string"),
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
		"string",
		orb.SubscriptionTriggerPhaseParams{
			EffectiveDate: orb.F(time.Now()),
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
	_, err := client.Subscriptions.UnscheduleCancellation(context.TODO(), "string")
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
		"string",
		orb.SubscriptionUnscheduleFixedFeeQuantityUpdatesParams{
			PriceID: orb.F("string"),
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
	_, err := client.Subscriptions.UnschedulePendingPlanChanges(context.TODO(), "string")
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
		"string",
		orb.SubscriptionUpdateFixedFeeQuantityParams{
			PriceID:       orb.F("string"),
			Quantity:      orb.F(0.000000),
			ChangeOption:  orb.F(orb.SubscriptionUpdateFixedFeeQuantityParamsChangeOptionImmediate),
			EffectiveDate: orb.F(time.Now()),
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
