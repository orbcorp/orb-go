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
		Email: orb.F("dev@stainlessapi.com"),
		Name:  orb.F("x"),
		AccountingSyncConfiguration: orb.F(orb.CustomerNewParamsAccountingSyncConfiguration{
			AccountingProviders: orb.F([]orb.CustomerNewParamsAccountingSyncConfigurationAccountingProvider{{
				ExternalProviderID: orb.F("external_provider_id"),
				ProviderType:       orb.F("provider_type"),
			}, {
				ExternalProviderID: orb.F("external_provider_id"),
				ProviderType:       orb.F("provider_type"),
			}, {
				ExternalProviderID: orb.F("external_provider_id"),
				ProviderType:       orb.F("provider_type"),
			}}),
			Excluded: orb.F(true),
		}),
		AdditionalEmails: orb.F([]string{"string", "string", "string"}),
		AutoCollection:   orb.F(true),
		BillingAddress: orb.F(orb.CustomerNewParamsBillingAddress{
			City:       orb.F("city"),
			Country:    orb.F("country"),
			Line1:      orb.F("line1"),
			Line2:      orb.F("line2"),
			PostalCode: orb.F("postal_code"),
			State:      orb.F("state"),
		}),
		Currency:           orb.F("currency"),
		EmailDelivery:      orb.F(true),
		ExternalCustomerID: orb.F("external_customer_id"),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		PaymentProvider:   orb.F(orb.CustomerNewParamsPaymentProviderQuickbooks),
		PaymentProviderID: orb.F("payment_provider_id"),
		ReportingConfiguration: orb.F(orb.CustomerNewParamsReportingConfiguration{
			Exempt: orb.F(true),
		}),
		ShippingAddress: orb.F(orb.CustomerNewParamsShippingAddress{
			City:       orb.F("city"),
			Country:    orb.F("country"),
			Line1:      orb.F("line1"),
			Line2:      orb.F("line2"),
			PostalCode: orb.F("postal_code"),
			State:      orb.F("state"),
		}),
		TaxConfiguration: orb.F[orb.CustomerNewParamsTaxConfigurationUnion](orb.CustomerNewParamsTaxConfigurationNewAvalaraTaxConfiguration{
			TaxExempt:        orb.F(true),
			TaxProvider:      orb.F(orb.CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara),
			TaxExemptionCode: orb.F("tax_exemption_code"),
		}),
		TaxID: orb.F(orb.CustomerNewParamsTaxID{
			Country: orb.F(orb.CustomerNewParamsTaxIDCountryAd),
			Type:    orb.F(orb.CustomerNewParamsTaxIDTypeAdNrt),
			Value:   orb.F("value"),
		}),
		Timezone: orb.F("timezone"),
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
		"customer_id",
		orb.CustomerUpdateParams{
			AccountingSyncConfiguration: orb.F(orb.CustomerUpdateParamsAccountingSyncConfiguration{
				AccountingProviders: orb.F([]orb.CustomerUpdateParamsAccountingSyncConfigurationAccountingProvider{{
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}, {
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}, {
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}}),
				Excluded: orb.F(true),
			}),
			AdditionalEmails: orb.F([]string{"string"}),
			AutoCollection:   orb.F(true),
			BillingAddress: orb.F(orb.CustomerUpdateParamsBillingAddress{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			Currency:           orb.F("currency"),
			Email:              orb.F("dev@stainlessapi.com"),
			EmailDelivery:      orb.F(true),
			ExternalCustomerID: orb.F("external_customer_id"),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			Name:              orb.F("name"),
			PaymentProvider:   orb.F(orb.CustomerUpdateParamsPaymentProviderQuickbooks),
			PaymentProviderID: orb.F("payment_provider_id"),
			ReportingConfiguration: orb.F(orb.CustomerUpdateParamsReportingConfiguration{
				Exempt: orb.F(true),
			}),
			ShippingAddress: orb.F(orb.CustomerUpdateParamsShippingAddress{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			TaxConfiguration: orb.F[orb.CustomerUpdateParamsTaxConfigurationUnion](orb.CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfiguration{
				TaxExempt:        orb.F(true),
				TaxProvider:      orb.F(orb.CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara),
				TaxExemptionCode: orb.F("tax_exemption_code"),
			}),
			TaxID: orb.F(orb.CustomerUpdateParamsTaxID{
				Country: orb.F(orb.CustomerUpdateParamsTaxIDCountryAd),
				Type:    orb.F(orb.CustomerUpdateParamsTaxIDTypeAdNrt),
				Value:   orb.F("value"),
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
		Cursor:       orb.F("cursor"),
		Limit:        orb.F(int64(1)),
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
	err := client.Customers.Delete(context.TODO(), "customer_id")
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
	_, err := client.Customers.Fetch(context.TODO(), "customer_id")
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
	_, err := client.Customers.FetchByExternalID(context.TODO(), "external_customer_id")
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
		"external_customer_id",
		orb.CustomerUpdateByExternalIDParams{
			AccountingSyncConfiguration: orb.F(orb.CustomerUpdateByExternalIDParamsAccountingSyncConfiguration{
				AccountingProviders: orb.F([]orb.CustomerUpdateByExternalIDParamsAccountingSyncConfigurationAccountingProvider{{
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}, {
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}, {
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}}),
				Excluded: orb.F(true),
			}),
			AdditionalEmails: orb.F([]string{"string"}),
			AutoCollection:   orb.F(true),
			BillingAddress: orb.F(orb.CustomerUpdateByExternalIDParamsBillingAddress{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			Currency:           orb.F("currency"),
			Email:              orb.F("dev@stainlessapi.com"),
			EmailDelivery:      orb.F(true),
			ExternalCustomerID: orb.F("external_customer_id"),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			Name:              orb.F("name"),
			PaymentProvider:   orb.F(orb.CustomerUpdateByExternalIDParamsPaymentProviderQuickbooks),
			PaymentProviderID: orb.F("payment_provider_id"),
			ReportingConfiguration: orb.F(orb.CustomerUpdateByExternalIDParamsReportingConfiguration{
				Exempt: orb.F(true),
			}),
			ShippingAddress: orb.F(orb.CustomerUpdateByExternalIDParamsShippingAddress{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			TaxConfiguration: orb.F[orb.CustomerUpdateByExternalIDParamsTaxConfigurationUnion](orb.CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfiguration{
				TaxExempt:        orb.F(true),
				TaxProvider:      orb.F(orb.CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara),
				TaxExemptionCode: orb.F("tax_exemption_code"),
			}),
			TaxID: orb.F(orb.CustomerUpdateByExternalIDParamsTaxID{
				Country: orb.F(orb.CustomerUpdateByExternalIDParamsTaxIDCountryAd),
				Type:    orb.F(orb.CustomerUpdateByExternalIDParamsTaxIDTypeAdNrt),
				Value:   orb.F("value"),
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
