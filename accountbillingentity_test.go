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

func TestAccountBillingEntityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.BillingEntities.New(context.TODO(), orb.AccountBillingEntityNewParams{
		Address: orb.F(orb.AddressInputParam{
			Country:    orb.F("x"),
			City:       orb.F("city"),
			Line1:      orb.F("line1"),
			Line2:      orb.F("line2"),
			PostalCode: orb.F("postal_code"),
			State:      orb.F("state"),
		}),
		LegalName:               orb.F("x"),
		Name:                    orb.F("x"),
		ExternalBillingEntityID: orb.F("x"),
		InvoiceSettings: orb.F(orb.AccountBillingEntityNewParamsInvoiceSettings{
			AutoCollection:      orb.F(true),
			AutomaticTaxEnabled: orb.F(true),
			EmailDelivery:       orb.F(true),
			ManualIssuance:      orb.F(true),
			Memo:                orb.F("memo"),
			NetTerms:            orb.F(int64(0)),
		}),
		TaxRegistrations: orb.F([]orb.AccountBillingEntityNewParamsTaxRegistration{{
			Code:      orb.F("x"),
			Value:     orb.F("x"),
			AppliesTo: orb.F("applies_to"),
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

func TestAccountBillingEntityUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.BillingEntities.Update(
		context.TODO(),
		"billing_entity_id",
		orb.AccountBillingEntityUpdateParams{
			Address: orb.F(orb.AddressInputParam{
				Country:    orb.F("x"),
				City:       orb.F("city"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			DefaultPaymentGatewayConnectionID: orb.F("default_payment_gateway_connection_id"),
			InvoiceSettings: orb.F(orb.AccountBillingEntityUpdateParamsInvoiceSettings{
				AutoCollection:      orb.F(true),
				AutomaticTaxEnabled: orb.F(true),
				EmailDelivery:       orb.F(true),
				ManualIssuance:      orb.F(true),
				Memo:                orb.F("memo"),
				NetTerms:            orb.F(int64(0)),
			}),
			LegalName: orb.F("x"),
			Name:      orb.F("x"),
			PaymentGatewayConnections: orb.F([]orb.AccountBillingEntityUpdateParamsPaymentGatewayConnection{{
				PaymentGatewayConnectionID: orb.F("payment_gateway_connection_id"),
				PaymentProviderType:        orb.F(orb.AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderTypeStripe),
			}}),
			TaxRegistrations: orb.F([]orb.AccountBillingEntityUpdateParamsTaxRegistration{{
				Code:      orb.F("x"),
				Value:     orb.F("x"),
				AppliesTo: orb.F("applies_to"),
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

func TestAccountBillingEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.BillingEntities.List(context.TODO(), orb.AccountBillingEntityListParams{
		Cursor: orb.F("cursor"),
		Limit:  orb.F(int64(1)),
		Status: orb.F(orb.AccountBillingEntityListParamsStatusActive),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityArchive(t *testing.T) {
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
	_, err := client.Account.BillingEntities.Archive(context.TODO(), "billing_entity_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityArchiveByExternalID(t *testing.T) {
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
	_, err := client.Account.BillingEntities.ArchiveByExternalID(context.TODO(), "external_billing_entity_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityFetch(t *testing.T) {
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
	_, err := client.Account.BillingEntities.Fetch(context.TODO(), "billing_entity_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityFetchByExternalID(t *testing.T) {
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
	_, err := client.Account.BillingEntities.FetchByExternalID(context.TODO(), "external_billing_entity_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityUnarchive(t *testing.T) {
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
	_, err := client.Account.BillingEntities.Unarchive(context.TODO(), "billing_entity_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityUnarchiveByExternalID(t *testing.T) {
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
	_, err := client.Account.BillingEntities.UnarchiveByExternalID(context.TODO(), "external_billing_entity_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBillingEntityUpdateByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.BillingEntities.UpdateByExternalID(
		context.TODO(),
		"external_billing_entity_id",
		orb.AccountBillingEntityUpdateByExternalIDParams{
			Address: orb.F(orb.AddressInputParam{
				Country:    orb.F("x"),
				City:       orb.F("city"),
				Line1:      orb.F("line1"),
				Line2:      orb.F("line2"),
				PostalCode: orb.F("postal_code"),
				State:      orb.F("state"),
			}),
			DefaultPaymentGatewayConnectionID: orb.F("default_payment_gateway_connection_id"),
			InvoiceSettings: orb.F(orb.AccountBillingEntityUpdateByExternalIDParamsInvoiceSettings{
				AutoCollection:      orb.F(true),
				AutomaticTaxEnabled: orb.F(true),
				EmailDelivery:       orb.F(true),
				ManualIssuance:      orb.F(true),
				Memo:                orb.F("memo"),
				NetTerms:            orb.F(int64(0)),
			}),
			LegalName: orb.F("x"),
			Name:      orb.F("x"),
			PaymentGatewayConnections: orb.F([]orb.AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnection{{
				PaymentGatewayConnectionID: orb.F("payment_gateway_connection_id"),
				PaymentProviderType:        orb.F(orb.AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderTypeStripe),
			}}),
			TaxRegistrations: orb.F([]orb.AccountBillingEntityUpdateByExternalIDParamsTaxRegistration{{
				Code:      orb.F("x"),
				Value:     orb.F("x"),
				AppliesTo: orb.F("applies_to"),
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
