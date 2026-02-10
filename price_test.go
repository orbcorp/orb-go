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

func TestPriceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.New(context.TODO(), orb.PriceNewParamsNewFloatingUnitPrice{
		Cadence:   orb.F(orb.PriceNewParamsNewFloatingUnitPriceCadenceAnnual),
		Currency:  orb.F("currency"),
		ItemID:    orb.F("item_id"),
		ModelType: orb.F(orb.PriceNewParamsNewFloatingUnitPriceModelTypeUnit),
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
		ConversionRateConfig: orb.F[orb.PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion](shared.UnitConversionRateConfigParam{
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
		LicenseTypeID: orb.F("license_type_id"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.Update(
		context.TODO(),
		"price_id",
		orb.PriceUpdateParams{
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

func TestPriceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.List(context.TODO(), orb.PriceListParams{
		Cursor: orb.F("cursor"),
		Limit:  orb.F(int64(1)),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceEvaluateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.Evaluate(
		context.TODO(),
		"price_id",
		orb.PriceEvaluateParams{
			TimeframeEnd:       orb.F(time.Now()),
			TimeframeStart:     orb.F(time.Now()),
			CustomerID:         orb.F("customer_id"),
			ExternalCustomerID: orb.F("external_customer_id"),
			Filter:             orb.F("my_numeric_property > 100 AND my_other_property = 'bar'"),
			GroupingKeys:       orb.F([]string{"case when my_event_type = 'foo' then true else false end"}),
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

func TestPriceEvaluateMultipleWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.EvaluateMultiple(context.TODO(), orb.PriceEvaluateMultipleParams{
		TimeframeEnd:       orb.F(time.Now()),
		TimeframeStart:     orb.F(time.Now()),
		CustomerID:         orb.F("customer_id"),
		ExternalCustomerID: orb.F("external_customer_id"),
		PriceEvaluations: orb.F([]orb.PriceEvaluateMultipleParamsPriceEvaluation{{
			ExternalPriceID: orb.F("external_price_id"),
			Filter:          orb.F("my_numeric_property > 100 AND my_other_property = 'bar'"),
			GroupingKeys:    orb.F([]string{"case when my_event_type = 'foo' then true else false end"}),
			Price: orb.F[orb.PriceEvaluateMultipleParamsPriceEvaluationsPriceUnion](shared.NewFloatingUnitPriceParam{
				Cadence:   orb.F(shared.NewFloatingUnitPriceCadenceAnnual),
				Currency:  orb.F("currency"),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(shared.NewFloatingUnitPriceModelTypeUnit),
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
				LicenseTypeID: orb.F("license_type_id"),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
			}),
			PriceID: orb.F("price_id"),
		}}),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceEvaluatePreviewEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Prices.EvaluatePreviewEvents(context.TODO(), orb.PriceEvaluatePreviewEventsParams{
		TimeframeEnd:   orb.F(time.Now()),
		TimeframeStart: orb.F(time.Now()),
		CustomerID:     orb.F("customer_id"),
		Events: orb.F([]orb.PriceEvaluatePreviewEventsParamsEvent{{
			EventName: orb.F("event_name"),
			Properties: orb.F(map[string]interface{}{
				"foo": "bar",
			}),
			Timestamp:          orb.F(time.Now()),
			CustomerID:         orb.F("customer_id"),
			ExternalCustomerID: orb.F("external_customer_id"),
		}}),
		ExternalCustomerID: orb.F("external_customer_id"),
		PriceEvaluations: orb.F([]orb.PriceEvaluatePreviewEventsParamsPriceEvaluation{{
			ExternalPriceID: orb.F("external_price_id"),
			Filter:          orb.F("my_numeric_property > 100 AND my_other_property = 'bar'"),
			GroupingKeys:    orb.F([]string{"case when my_event_type = 'foo' then true else false end"}),
			Price: orb.F[orb.PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion](shared.NewFloatingUnitPriceParam{
				Cadence:   orb.F(shared.NewFloatingUnitPriceCadenceAnnual),
				Currency:  orb.F("currency"),
				ItemID:    orb.F("item_id"),
				ModelType: orb.F(shared.NewFloatingUnitPriceModelTypeUnit),
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
				LicenseTypeID: orb.F("license_type_id"),
				Metadata: orb.F(map[string]string{
					"foo": "string",
				}),
			}),
			PriceID: orb.F("price_id"),
		}}),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceFetch(t *testing.T) {
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
	_, err := client.Prices.Fetch(context.TODO(), "price_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
