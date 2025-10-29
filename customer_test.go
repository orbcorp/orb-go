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
		Email: orb.F("dev@stainless.com"),
		Name:  orb.F("x"),
		AccountingSyncConfiguration: orb.F(orb.NewAccountingSyncConfigurationParam{
			AccountingProviders: orb.F([]orb.AccountingProviderConfigParam{{
				ExternalProviderID: orb.F("external_provider_id"),
				ProviderType:       orb.F("provider_type"),
			}}),
			Excluded: orb.F(true),
		}),
		AdditionalEmails: orb.F([]string{"dev@stainless.com"}),
		AutoCollection:   orb.F(true),
		AutoIssuance:     orb.F(true),
		BillingAddress: orb.F(orb.AddressInputParam{
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
		Hierarchy: orb.F(orb.CustomerHierarchyConfigParam{
			ChildCustomerIDs: orb.F([]string{"string"}),
			ParentCustomerID: orb.F("parent_customer_id"),
		}),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		PaymentProvider:   orb.F(orb.CustomerNewParamsPaymentProviderQuickbooks),
		PaymentProviderID: orb.F("payment_provider_id"),
		ReportingConfiguration: orb.F(orb.NewReportingConfigurationParam{
			Exempt: orb.F(true),
		}),
		ShippingAddress: orb.F(orb.AddressInputParam{
			City:       orb.F("city"),
			Country:    orb.F("country"),
			Line1:      orb.F("line1"),
			Line2:      orb.F("line2"),
			PostalCode: orb.F("postal_code"),
			State:      orb.F("state"),
		}),
		TaxConfiguration: orb.F[orb.CustomerNewParamsTaxConfigurationUnion](orb.NewAvalaraTaxConfigurationParam{
			TaxExempt:           orb.F(true),
			TaxProvider:         orb.F(orb.NewAvalaraTaxConfigurationTaxProviderAvalara),
			AutomaticTaxEnabled: orb.F(true),
			TaxExemptionCode:    orb.F("tax_exemption_code"),
		}),
		TaxID: orb.F(shared.CustomerTaxIDParam{
			Country: orb.F(shared.CustomerTaxIDCountryAd),
			Type:    orb.F(shared.CustomerTaxIDTypeAdNrt),
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
			AccountingSyncConfiguration: orb.F(orb.NewAccountingSyncConfigurationParam{
				AccountingProviders: orb.F([]orb.AccountingProviderConfigParam{{
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}}),
				Excluded: orb.F(true),
			}),
			AdditionalEmails:    orb.F([]string{"string"}),
			AutoCollection:      orb.F(true),
			AutoIssuance:        orb.F(true),
			AutomaticTaxEnabled: orb.F(true),
			BillingAddress: orb.F(orb.AddressInputParam{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			Currency:           orb.F("currency"),
			Email:              orb.F("dev@stainless.com"),
			EmailDelivery:      orb.F(true),
			ExternalCustomerID: orb.F("external_customer_id"),
			Hierarchy: orb.F(orb.CustomerHierarchyConfigParam{
				ChildCustomerIDs: orb.F([]string{"string"}),
				ParentCustomerID: orb.F("parent_customer_id"),
			}),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			Name:              orb.F("name"),
			PaymentProvider:   orb.F(orb.CustomerUpdateParamsPaymentProviderQuickbooks),
			PaymentProviderID: orb.F("payment_provider_id"),
			ReportingConfiguration: orb.F(orb.NewReportingConfigurationParam{
				Exempt: orb.F(true),
			}),
			ShippingAddress: orb.F(orb.AddressInputParam{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			TaxConfiguration: orb.F[orb.CustomerUpdateParamsTaxConfigurationUnion](orb.NewAvalaraTaxConfigurationParam{
				TaxExempt:           orb.F(true),
				TaxProvider:         orb.F(orb.NewAvalaraTaxConfigurationTaxProviderAvalara),
				AutomaticTaxEnabled: orb.F(true),
				TaxExemptionCode:    orb.F("tax_exemption_code"),
			}),
			TaxID: orb.F(shared.CustomerTaxIDParam{
				Country: orb.F(shared.CustomerTaxIDCountryAd),
				Type:    orb.F(shared.CustomerTaxIDTypeAdNrt),
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

func TestCustomerSyncPaymentMethodsFromGateway(t *testing.T) {
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
	err := client.Customers.SyncPaymentMethodsFromGateway(context.TODO(), "customer_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerSyncPaymentMethodsFromGatewayByExternalCustomerID(t *testing.T) {
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
	err := client.Customers.SyncPaymentMethodsFromGatewayByExternalCustomerID(context.TODO(), "external_customer_id")
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
			AccountingSyncConfiguration: orb.F(orb.NewAccountingSyncConfigurationParam{
				AccountingProviders: orb.F([]orb.AccountingProviderConfigParam{{
					ExternalProviderID: orb.F("external_provider_id"),
					ProviderType:       orb.F("provider_type"),
				}}),
				Excluded: orb.F(true),
			}),
			AdditionalEmails:    orb.F([]string{"string"}),
			AutoCollection:      orb.F(true),
			AutoIssuance:        orb.F(true),
			AutomaticTaxEnabled: orb.F(true),
			BillingAddress: orb.F(orb.AddressInputParam{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			Currency:           orb.F("currency"),
			Email:              orb.F("dev@stainless.com"),
			EmailDelivery:      orb.F(true),
			ExternalCustomerID: orb.F("external_customer_id"),
			Hierarchy: orb.F(orb.CustomerHierarchyConfigParam{
				ChildCustomerIDs: orb.F([]string{"string"}),
				ParentCustomerID: orb.F("parent_customer_id"),
			}),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			Name:              orb.F("name"),
			PaymentProvider:   orb.F(orb.CustomerUpdateByExternalIDParamsPaymentProviderQuickbooks),
			PaymentProviderID: orb.F("payment_provider_id"),
			ReportingConfiguration: orb.F(orb.NewReportingConfigurationParam{
				Exempt: orb.F(true),
			}),
			ShippingAddress: orb.F(orb.AddressInputParam{
				City:       orb.F("city"),
				Country:    orb.F("country"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			TaxConfiguration: orb.F[orb.CustomerUpdateByExternalIDParamsTaxConfigurationUnion](orb.NewAvalaraTaxConfigurationParam{
				TaxExempt:           orb.F(true),
				TaxProvider:         orb.F(orb.NewAvalaraTaxConfigurationTaxProviderAvalara),
				AutomaticTaxEnabled: orb.F(true),
				TaxExemptionCode:    orb.F("tax_exemption_code"),
			}),
			TaxID: orb.F(shared.CustomerTaxIDParam{
				Country: orb.F(shared.CustomerTaxIDCountryAd),
				Type:    orb.F(shared.CustomerTaxIDTypeAdNrt),
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
