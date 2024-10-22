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
			Adjustment: orb.F[orb.SubscriptionNewParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
				AdjustmentType:     orb.F(orb.SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				PercentageDiscount: orb.F(0.000000),
			}),
			EndDate:        orb.F(time.Now()),
			PlanPhaseOrder: orb.F(int64(0)),
			StartDate:      orb.F(time.Now()),
		}, {
			Adjustment: orb.F[orb.SubscriptionNewParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
				AdjustmentType:     orb.F(orb.SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				PercentageDiscount: orb.F(0.000000),
			}),
			EndDate:        orb.F(time.Now()),
			PlanPhaseOrder: orb.F(int64(0)),
			StartDate:      orb.F(time.Now()),
		}, {
			Adjustment: orb.F[orb.SubscriptionNewParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
				AdjustmentType:     orb.F(orb.SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				PercentageDiscount: orb.F(0.000000),
			}),
			EndDate:        orb.F(time.Now()),
			PlanPhaseOrder: orb.F(int64(0)),
			StartDate:      orb.F(time.Now()),
		}}),
		AddPrices: orb.F([]orb.SubscriptionNewParamsAddPrice{{
			Discounts: orb.F([]orb.SubscriptionNewParamsAddPricesDiscount{{
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			EndDate:         orb.F(time.Now()),
			ExternalPriceID: orb.F("external_price_id"),
			MaximumAmount:   orb.F("1.23"),
			MinimumAmount:   orb.F("1.23"),
			PlanPhaseOrder:  orb.F(int64(0)),
			Price: orb.F[orb.SubscriptionNewParamsAddPricesPriceUnion](orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice{
				Cadence:   orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate:     orb.F(0.000000),
				Currency:           orb.F("currency"),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("invoice_grouping_key"),
				InvoicingCycleConfiguration: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
				}),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
				ReferenceID: orb.F("reference_id"),
			}),
			PriceID:   orb.F("h74gfhdjvn7ujokd"),
			StartDate: orb.F(time.Now()),
		}, {
			Discounts: orb.F([]orb.SubscriptionNewParamsAddPricesDiscount{{
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			EndDate:         orb.F(time.Now()),
			ExternalPriceID: orb.F("external_price_id"),
			MaximumAmount:   orb.F("1.23"),
			MinimumAmount:   orb.F("1.23"),
			PlanPhaseOrder:  orb.F(int64(0)),
			Price: orb.F[orb.SubscriptionNewParamsAddPricesPriceUnion](orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice{
				Cadence:   orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate:     orb.F(0.000000),
				Currency:           orb.F("currency"),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("invoice_grouping_key"),
				InvoicingCycleConfiguration: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
				}),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
				ReferenceID: orb.F("reference_id"),
			}),
			PriceID:   orb.F("h74gfhdjvn7ujokd"),
			StartDate: orb.F(time.Now()),
		}, {
			Discounts: orb.F([]orb.SubscriptionNewParamsAddPricesDiscount{{
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsAddPricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			EndDate:         orb.F(time.Now()),
			ExternalPriceID: orb.F("external_price_id"),
			MaximumAmount:   orb.F("1.23"),
			MinimumAmount:   orb.F("1.23"),
			PlanPhaseOrder:  orb.F(int64(0)),
			Price: orb.F[orb.SubscriptionNewParamsAddPricesPriceUnion](orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice{
				Cadence:   orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate:     orb.F(0.000000),
				Currency:           orb.F("currency"),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("invoice_grouping_key"),
				InvoicingCycleConfiguration: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
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
		BillingCycleAnchorConfiguration: orb.F(orb.SubscriptionNewParamsBillingCycleAnchorConfiguration{
			Day:   orb.F(int64(1)),
			Month: orb.F(int64(1)),
			Year:  orb.F(int64(0)),
		}),
		CouponRedemptionCode:           orb.F("coupon_redemption_code"),
		CreditsOverageRate:             orb.F(0.000000),
		CustomerID:                     orb.F("customer_id"),
		DefaultInvoiceMemo:             orb.F("default_invoice_memo"),
		EndDate:                        orb.F(time.Now()),
		ExternalCustomerID:             orb.F("external_customer_id"),
		ExternalMarketplace:            orb.F(orb.SubscriptionNewParamsExternalMarketplaceGoogle),
		ExternalMarketplaceReportingID: orb.F("external_marketplace_reporting_id"),
		ExternalPlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
		InitialPhaseOrder:              orb.F(int64(2)),
		InvoicingThreshold:             orb.F("10.00"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		NetTerms:               orb.F(int64(0)),
		PerCreditOverageAmount: orb.F(0.000000),
		PlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
		PlanVersionNumber:      orb.F(int64(0)),
		PriceOverrides:         orb.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
		RemoveAdjustments: orb.F([]orb.SubscriptionNewParamsRemoveAdjustment{{
			AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
		}, {
			AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
		}, {
			AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
		}}),
		RemovePrices: orb.F([]orb.SubscriptionNewParamsRemovePrice{{
			ExternalPriceID: orb.F("external_price_id"),
			PriceID:         orb.F("h74gfhdjvn7ujokd"),
		}, {
			ExternalPriceID: orb.F("external_price_id"),
			PriceID:         orb.F("h74gfhdjvn7ujokd"),
		}, {
			ExternalPriceID: orb.F("external_price_id"),
			PriceID:         orb.F("h74gfhdjvn7ujokd"),
		}}),
		ReplaceAdjustments: orb.F([]orb.SubscriptionNewParamsReplaceAdjustment{{
			Adjustment: orb.F[orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion](orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
				AdjustmentType:     orb.F(orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				PercentageDiscount: orb.F(0.000000),
			}),
			ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
		}, {
			Adjustment: orb.F[orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion](orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
				AdjustmentType:     orb.F(orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				PercentageDiscount: orb.F(0.000000),
			}),
			ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
		}, {
			Adjustment: orb.F[orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion](orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
				AdjustmentType:     orb.F(orb.SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
				AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
				PercentageDiscount: orb.F(0.000000),
			}),
			ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
		}}),
		ReplacePrices: orb.F([]orb.SubscriptionNewParamsReplacePrice{{
			ReplacesPriceID: orb.F("replaces_price_id"),
			Discounts: orb.F([]orb.SubscriptionNewParamsReplacePricesDiscount{{
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			ExternalPriceID:    orb.F("external_price_id"),
			FixedPriceQuantity: orb.F(2.000000),
			MaximumAmount:      orb.F("1.23"),
			MinimumAmount:      orb.F("1.23"),
			Price: orb.F[orb.SubscriptionNewParamsReplacePricesPriceUnion](orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice{
				Cadence:   orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate:     orb.F(0.000000),
				Currency:           orb.F("currency"),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("invoice_grouping_key"),
				InvoicingCycleConfiguration: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
				}),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
				ReferenceID: orb.F("reference_id"),
			}),
			PriceID: orb.F("h74gfhdjvn7ujokd"),
		}, {
			ReplacesPriceID: orb.F("replaces_price_id"),
			Discounts: orb.F([]orb.SubscriptionNewParamsReplacePricesDiscount{{
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			ExternalPriceID:    orb.F("external_price_id"),
			FixedPriceQuantity: orb.F(2.000000),
			MaximumAmount:      orb.F("1.23"),
			MinimumAmount:      orb.F("1.23"),
			Price: orb.F[orb.SubscriptionNewParamsReplacePricesPriceUnion](orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice{
				Cadence:   orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate:     orb.F(0.000000),
				Currency:           orb.F("currency"),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("invoice_grouping_key"),
				InvoicingCycleConfiguration: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
				}),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
				ReferenceID: orb.F("reference_id"),
			}),
			PriceID: orb.F("h74gfhdjvn7ujokd"),
		}, {
			ReplacesPriceID: orb.F("replaces_price_id"),
			Discounts: orb.F([]orb.SubscriptionNewParamsReplacePricesDiscount{{
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}, {
				DiscountType:       orb.F(orb.SubscriptionNewParamsReplacePricesDiscountsDiscountTypePercentage),
				AmountDiscount:     orb.F("amount_discount"),
				PercentageDiscount: orb.F(0.150000),
				UsageDiscount:      orb.F(0.000000),
			}}),
			ExternalPriceID:    orb.F("external_price_id"),
			FixedPriceQuantity: orb.F(2.000000),
			MaximumAmount:      orb.F("1.23"),
			MinimumAmount:      orb.F("1.23"),
			Price: orb.F[orb.SubscriptionNewParamsReplacePricesPriceUnion](orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice{
				Cadence:   orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit),
				Name:      orb.F("Annual fee"),
				UnitConfig: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig{
					UnitAmount: orb.F("unit_amount"),
				}),
				BillableMetricID: orb.F("billable_metric_id"),
				BilledInAdvance:  orb.F(true),
				BillingCycleConfiguration: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
				}),
				ConversionRate:     orb.F(0.000000),
				Currency:           orb.F("currency"),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(0.000000),
				InvoiceGroupingKey: orb.F("invoice_grouping_key"),
				InvoicingCycleConfiguration: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
					Duration:     orb.F(int64(0)),
					DurationUnit: orb.F(orb.SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
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
		CustomerID:         orb.F([]string{"string", "string", "string"}),
		ExternalCustomerID: orb.F("external_customer_id"),
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
		"subscription_id",
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
				AllocationPrice: orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceOneTime),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}}),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				ExternalPriceID: orb.F("external_price_id"),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				MaximumAmount: orb.F(0.000000),
				MinimumAmount: orb.F(0.000000),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice{
					Cadence:   orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual),
					Currency:  orb.F("currency"),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}, {
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddStartDateUnion](shared.UnionTime(time.Now())),
				AllocationPrice: orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceOneTime),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}}),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				ExternalPriceID: orb.F("external_price_id"),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				MaximumAmount: orb.F(0.000000),
				MinimumAmount: orb.F(0.000000),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice{
					Cadence:   orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual),
					Currency:  orb.F("currency"),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}, {
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddStartDateUnion](shared.UnionTime(time.Now())),
				AllocationPrice: orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPrice{
					Amount:                orb.F("10.00"),
					Cadence:               orb.F(orb.SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceOneTime),
					Currency:              orb.F("USD"),
					ExpiresAtEndOfCadence: orb.F(true),
				}),
				Discounts: orb.F([]orb.SubscriptionPriceIntervalsParamsAddDiscountUnion{orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}, orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams{
					AmountDiscount: orb.F(0.000000),
					DiscountType:   orb.F(orb.SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount),
				}}),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsAddEndDateUnion](shared.UnionTime(time.Now())),
				ExternalPriceID: orb.F("external_price_id"),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				MaximumAmount: orb.F(0.000000),
				MinimumAmount: orb.F(0.000000),
				Price: orb.F[orb.SubscriptionPriceIntervalsParamsAddPriceUnion](orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice{
					Cadence:   orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual),
					Currency:  orb.F("currency"),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
				}),
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}}),
			AddAdjustments: orb.F([]orb.SubscriptionPriceIntervalsParamsAddAdjustment{{
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}, {
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}, {
				Adjustment: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
				EndDate:   orb.F[orb.SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
			}}),
			Edit: orb.F([]orb.SubscriptionPriceIntervalsParamsEdit{{
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				BillingCycleDay: orb.F(int64(0)),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
			}, {
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				BillingCycleDay: orb.F(int64(0)),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
			}, {
				PriceIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				BillingCycleDay: orb.F(int64(0)),
				EndDate:         orb.F[orb.SubscriptionPriceIntervalsParamsEditEndDateUnion](shared.UnionTime(time.Now())),
				FixedFeeQuantityTransitions: orb.F([]orb.SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition{{
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}, {
					EffectiveDate: orb.F(time.Now()),
					Quantity:      orb.F(int64(5)),
				}}),
				StartDate: orb.F[orb.SubscriptionPriceIntervalsParamsEditStartDateUnion](shared.UnionTime(time.Now())),
			}}),
			EditAdjustments: orb.F([]orb.SubscriptionPriceIntervalsParamsEditAdjustment{{
				AdjustmentIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				EndDate:              orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
				StartDate:            orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
			}, {
				AdjustmentIntervalID: orb.F("sdfs6wdjvn7ujokd"),
				EndDate:              orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion](shared.UnionTime(time.Now())),
				StartDate:            orb.F[orb.SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion](shared.UnionTime(time.Now())),
			}, {
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
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				EndDate:        orb.F(time.Now()),
				PlanPhaseOrder: orb.F(int64(0)),
				StartDate:      orb.F(time.Now()),
			}, {
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				EndDate:        orb.F(time.Now()),
				PlanPhaseOrder: orb.F(int64(0)),
				StartDate:      orb.F(time.Now()),
			}, {
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion](orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				EndDate:        orb.F(time.Now()),
				PlanPhaseOrder: orb.F(int64(0)),
				StartDate:      orb.F(time.Now()),
			}}),
			AddPrices: orb.F([]orb.SubscriptionSchedulePlanChangeParamsAddPrice{{
				Discounts: orb.F([]orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscount{{
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				EndDate:         orb.F(time.Now()),
				ExternalPriceID: orb.F("external_price_id"),
				MaximumAmount:   orb.F("1.23"),
				MinimumAmount:   orb.F("1.23"),
				PlanPhaseOrder:  orb.F(int64(0)),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceUnion](orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPrice{
					Cadence:   orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					Currency:           orb.F("currency"),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PriceID:   orb.F("h74gfhdjvn7ujokd"),
				StartDate: orb.F(time.Now()),
			}, {
				Discounts: orb.F([]orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscount{{
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				EndDate:         orb.F(time.Now()),
				ExternalPriceID: orb.F("external_price_id"),
				MaximumAmount:   orb.F("1.23"),
				MinimumAmount:   orb.F("1.23"),
				PlanPhaseOrder:  orb.F(int64(0)),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceUnion](orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPrice{
					Cadence:   orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					Currency:           orb.F("currency"),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PriceID:   orb.F("h74gfhdjvn7ujokd"),
				StartDate: orb.F(time.Now()),
			}, {
				Discounts: orb.F([]orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscount{{
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				EndDate:         orb.F(time.Now()),
				ExternalPriceID: orb.F("external_price_id"),
				MaximumAmount:   orb.F("1.23"),
				MinimumAmount:   orb.F("1.23"),
				PlanPhaseOrder:  orb.F(int64(0)),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceUnion](orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPrice{
					Cadence:   orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					Currency:           orb.F("currency"),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
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
			ChangeDate:                     orb.F(time.Now()),
			CouponRedemptionCode:           orb.F("coupon_redemption_code"),
			CreditsOverageRate:             orb.F(0.000000),
			DefaultInvoiceMemo:             orb.F("default_invoice_memo"),
			ExternalPlanID:                 orb.F("ZMwNQefe7J3ecf7W"),
			InitialPhaseOrder:              orb.F(int64(2)),
			InvoicingThreshold:             orb.F("10.00"),
			NetTerms:                       orb.F(int64(0)),
			PerCreditOverageAmount:         orb.F(0.000000),
			PlanID:                         orb.F("ZMwNQefe7J3ecf7W"),
			PlanVersionNumber:              orb.F(int64(0)),
			PriceOverrides:                 orb.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
			RemoveAdjustments: orb.F([]orb.SubscriptionSchedulePlanChangeParamsRemoveAdjustment{{
				AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
			}, {
				AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
			}, {
				AdjustmentID: orb.F("h74gfhdjvn7ujokd"),
			}}),
			RemovePrices: orb.F([]orb.SubscriptionSchedulePlanChangeParamsRemovePrice{{
				ExternalPriceID: orb.F("external_price_id"),
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
			}, {
				ExternalPriceID: orb.F("external_price_id"),
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
			}, {
				ExternalPriceID: orb.F("external_price_id"),
				PriceID:         orb.F("h74gfhdjvn7ujokd"),
			}}),
			ReplaceAdjustments: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustment{{
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion](orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
			}, {
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion](orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
			}, {
				Adjustment: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion](orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount{
					AdjustmentType:     orb.F(orb.SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount),
					AppliesToPriceIDs:  orb.F([]string{"price_1", "price_2"}),
					PercentageDiscount: orb.F(0.000000),
				}),
				ReplacesAdjustmentID: orb.F("replaces_adjustment_id"),
			}}),
			ReplacePrices: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplacePrice{{
				ReplacesPriceID: orb.F("replaces_price_id"),
				Discounts: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscount{{
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(2.000000),
				MaximumAmount:      orb.F("1.23"),
				MinimumAmount:      orb.F("1.23"),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion](orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPrice{
					Cadence:   orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					Currency:           orb.F("currency"),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}, {
				ReplacesPriceID: orb.F("replaces_price_id"),
				Discounts: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscount{{
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(2.000000),
				MaximumAmount:      orb.F("1.23"),
				MinimumAmount:      orb.F("1.23"),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion](orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPrice{
					Cadence:   orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					Currency:           orb.F("currency"),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}, {
				ReplacesPriceID: orb.F("replaces_price_id"),
				Discounts: orb.F([]orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscount{{
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}, {
					DiscountType:       orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesDiscountsDiscountTypePercentage),
					AmountDiscount:     orb.F("amount_discount"),
					PercentageDiscount: orb.F(0.150000),
					UsageDiscount:      orb.F(0.000000),
				}}),
				ExternalPriceID:    orb.F("external_price_id"),
				FixedPriceQuantity: orb.F(2.000000),
				MaximumAmount:      orb.F("1.23"),
				MinimumAmount:      orb.F("1.23"),
				Price: orb.F[orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion](orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPrice{
					Cadence:   orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual),
					ItemID:    orb.F("item_id"),
					ModelType: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit),
					Name:      orb.F("Annual fee"),
					UnitConfig: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig{
						UnitAmount: orb.F("unit_amount"),
					}),
					BillableMetricID: orb.F("billable_metric_id"),
					BilledInAdvance:  orb.F(true),
					BillingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay),
					}),
					ConversionRate:     orb.F(0.000000),
					Currency:           orb.F("currency"),
					ExternalPriceID:    orb.F("external_price_id"),
					FixedPriceQuantity: orb.F(0.000000),
					InvoiceGroupingKey: orb.F("invoice_grouping_key"),
					InvoicingCycleConfiguration: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration{
						Duration:     orb.F(int64(0)),
						DurationUnit: orb.F(orb.SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay),
					}),
					Metadata: orb.F(map[string]string{
						"foo": "string",
					}),
					ReferenceID: orb.F("reference_id"),
				}),
				PriceID: orb.F("h74gfhdjvn7ujokd"),
			}}),
			TrialDurationDays: orb.F(int64(0)),
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
			PriceID:       orb.F("price_id"),
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
