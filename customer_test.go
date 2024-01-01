// File generated from our OpenAPI spec by Stainless.

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
)

func TestCustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.New(context.TODO(), orb.CustomerNewParams{
		Email: orb.F("string"),
		Name:  orb.F("string"),
		AccountingSyncConfiguration: orb.F(orb.CustomerNewParamsAccountingSyncConfiguration{
			Excluded: orb.F(true),
			AccountingProviders: orb.F([]orb.CustomerNewParamsAccountingSyncConfigurationAccountingProvider{{
				ProviderType:       orb.F("string"),
				ExternalProviderID: orb.F("string"),
			}, {
				ProviderType:       orb.F("string"),
				ExternalProviderID: orb.F("string"),
			}, {
				ProviderType:       orb.F("string"),
				ExternalProviderID: orb.F("string"),
			}}),
		}),
		AdditionalEmails: orb.F([]string{"string", "string", "string"}),
		AutoCollection:   orb.F(true),
		BillingAddress: orb.F(orb.CustomerNewParamsBillingAddress{
			Line1:      orb.F("string"),
			Line2:      orb.F("string"),
			City:       orb.F("string"),
			State:      orb.F("string"),
			PostalCode: orb.F("string"),
			Country:    orb.F("string"),
		}),
		Currency:           orb.F("string"),
		EmailDelivery:      orb.F(true),
		ExternalCustomerID: orb.F("string"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		PaymentProvider:   orb.F(orb.CustomerNewParamsPaymentProviderQuickbooks),
		PaymentProviderID: orb.F("string"),
		ReportingConfiguration: orb.F(orb.CustomerNewParamsReportingConfiguration{
			Exempt: orb.F(true),
		}),
		ShippingAddress: orb.F(orb.CustomerNewParamsShippingAddress{
			Line1:      orb.F("string"),
			Line2:      orb.F("string"),
			City:       orb.F("string"),
			State:      orb.F("string"),
			PostalCode: orb.F("string"),
			Country:    orb.F("string"),
		}),
		TaxID: orb.F(orb.CustomerNewParamsTaxID{
			Country: orb.F(orb.CustomerNewParamsTaxIDCountryAd),
			Type:    orb.F(orb.CustomerNewParamsTaxIDTypeAdNrt),
			Value:   orb.F("string"),
		}),
		Timezone: orb.F("string"),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Update(
		context.TODO(),
		"string",
		orb.CustomerUpdateParams{
			AccountingSyncConfiguration: orb.F(orb.CustomerUpdateParamsAccountingSyncConfiguration{
				Excluded: orb.F(true),
				AccountingProviders: orb.F([]orb.CustomerUpdateParamsAccountingSyncConfigurationAccountingProvider{{
					ProviderType:       orb.F("string"),
					ExternalProviderID: orb.F("string"),
				}, {
					ProviderType:       orb.F("string"),
					ExternalProviderID: orb.F("string"),
				}, {
					ProviderType:       orb.F("string"),
					ExternalProviderID: orb.F("string"),
				}}),
			}),
			AdditionalEmails: orb.F([]string{"string"}),
			AutoCollection:   orb.F(true),
			BillingAddress: orb.F(orb.CustomerUpdateParamsBillingAddress{
				Line1:      orb.F("string"),
				Line2:      orb.F("string"),
				City:       orb.F("string"),
				State:      orb.F("string"),
				PostalCode: orb.F("string"),
				Country:    orb.F("string"),
			}),
			Currency:      orb.F("string"),
			Email:         orb.F("string"),
			EmailDelivery: orb.F(true),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			Name:              orb.F("string"),
			PaymentProvider:   orb.F(orb.CustomerUpdateParamsPaymentProviderQuickbooks),
			PaymentProviderID: orb.F("string"),
			ReportingConfiguration: orb.F(orb.CustomerUpdateParamsReportingConfiguration{
				Exempt: orb.F(true),
			}),
			ShippingAddress: orb.F(orb.CustomerUpdateParamsShippingAddress{
				Line1:      orb.F("string"),
				Line2:      orb.F("string"),
				City:       orb.F("string"),
				State:      orb.F("string"),
				PostalCode: orb.F("string"),
				Country:    orb.F("string"),
			}),
			TaxID: orb.F(orb.CustomerUpdateParamsTaxID{
				Country: orb.F(orb.CustomerUpdateParamsTaxIDCountryAd),
				Type:    orb.F(orb.CustomerUpdateParamsTaxIDTypeAdNrt),
				Value:   orb.F("string"),
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

func TestCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.List(context.TODO(), orb.CustomerListParams{
		CreatedAtGt:  orb.F(time.Now()),
		CreatedAtGte: orb.F(time.Now()),
		CreatedAtLt:  orb.F(time.Now()),
		CreatedAtLte: orb.F(time.Now()),
		Cursor:       orb.F("string"),
		Limit:        orb.F(int64(0)),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerDelete(t *testing.T) {
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
	err := client.Customers.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerFetch(t *testing.T) {
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
	_, err := client.Customers.Fetch(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerFetchByExternalID(t *testing.T) {
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
	_, err := client.Customers.FetchByExternalID(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerUpdateByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.UpdateByExternalID(
		context.TODO(),
		"string",
		orb.CustomerUpdateByExternalIDParams{
			AccountingSyncConfiguration: orb.F(orb.CustomerUpdateByExternalIDParamsAccountingSyncConfiguration{
				Excluded: orb.F(true),
				AccountingProviders: orb.F([]orb.CustomerUpdateByExternalIDParamsAccountingSyncConfigurationAccountingProvider{{
					ProviderType:       orb.F("string"),
					ExternalProviderID: orb.F("string"),
				}, {
					ProviderType:       orb.F("string"),
					ExternalProviderID: orb.F("string"),
				}, {
					ProviderType:       orb.F("string"),
					ExternalProviderID: orb.F("string"),
				}}),
			}),
			AdditionalEmails: orb.F([]string{"string"}),
			AutoCollection:   orb.F(true),
			BillingAddress: orb.F(orb.CustomerUpdateByExternalIDParamsBillingAddress{
				Line1:      orb.F("string"),
				Line2:      orb.F("string"),
				City:       orb.F("string"),
				State:      orb.F("string"),
				PostalCode: orb.F("string"),
				Country:    orb.F("string"),
			}),
			Currency:      orb.F("string"),
			Email:         orb.F("string"),
			EmailDelivery: orb.F(true),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			Name:              orb.F("string"),
			PaymentProvider:   orb.F(orb.CustomerUpdateByExternalIDParamsPaymentProviderQuickbooks),
			PaymentProviderID: orb.F("string"),
			ReportingConfiguration: orb.F(orb.CustomerUpdateByExternalIDParamsReportingConfiguration{
				Exempt: orb.F(true),
			}),
			ShippingAddress: orb.F(orb.CustomerUpdateByExternalIDParamsShippingAddress{
				Line1:      orb.F("string"),
				Line2:      orb.F("string"),
				City:       orb.F("string"),
				State:      orb.F("string"),
				PostalCode: orb.F("string"),
				Country:    orb.F("string"),
			}),
			TaxID: orb.F(orb.CustomerUpdateByExternalIDParamsTaxID{
				Country: orb.F(orb.CustomerUpdateByExternalIDParamsTaxIDCountryAd),
				Type:    orb.F(orb.CustomerUpdateByExternalIDParamsTaxIDTypeAdNrt),
				Value:   orb.F("string"),
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
