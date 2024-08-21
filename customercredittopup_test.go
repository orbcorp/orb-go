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

func TestCustomerCreditTopUpNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.TopUps.New(
		context.TODO(),
		"customer_id",
		orb.CustomerCreditTopUpNewParams{
			Amount:   orb.F("amount"),
			Currency: orb.F("currency"),
			InvoiceSettings: orb.F(orb.CustomerCreditTopUpNewParamsInvoiceSettings{
				AutoCollection:           orb.F(true),
				NetTerms:                 orb.F(int64(0)),
				Memo:                     orb.F("memo"),
				RequireSuccessfulPayment: orb.F(true),
			}),
			PerUnitCostBasis: orb.F("per_unit_cost_basis"),
			Threshold:        orb.F("threshold"),
			ExpiresAfter:     orb.F(int64(0)),
			ExpiresAfterUnit: orb.F(orb.CustomerCreditTopUpNewParamsExpiresAfterUnitDay),
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

func TestCustomerCreditTopUpListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.TopUps.List(
		context.TODO(),
		"customer_id",
		orb.CustomerCreditTopUpListParams{
			Cursor: orb.F("cursor"),
			Limit:  orb.F(int64(1)),
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

func TestCustomerCreditTopUpDelete(t *testing.T) {
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
	err := client.Customers.Credits.TopUps.Delete(
		context.TODO(),
		"customer_id",
		"top_up_id",
	)
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerCreditTopUpNewByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.TopUps.NewByExternalID(
		context.TODO(),
		"external_customer_id",
		orb.CustomerCreditTopUpNewByExternalIDParams{
			Amount:   orb.F("amount"),
			Currency: orb.F("currency"),
			InvoiceSettings: orb.F(orb.CustomerCreditTopUpNewByExternalIDParamsInvoiceSettings{
				AutoCollection:           orb.F(true),
				NetTerms:                 orb.F(int64(0)),
				Memo:                     orb.F("memo"),
				RequireSuccessfulPayment: orb.F(true),
			}),
			PerUnitCostBasis: orb.F("per_unit_cost_basis"),
			Threshold:        orb.F("threshold"),
			ExpiresAfter:     orb.F(int64(0)),
			ExpiresAfterUnit: orb.F(orb.CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnitDay),
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

func TestCustomerCreditTopUpDeleteByExternalID(t *testing.T) {
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
	err := client.Customers.Credits.TopUps.DeleteByExternalID(
		context.TODO(),
		"external_customer_id",
		"top_up_id",
	)
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerCreditTopUpListByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.TopUps.ListByExternalID(
		context.TODO(),
		"external_customer_id",
		orb.CustomerCreditTopUpListByExternalIDParams{
			Cursor: orb.F("cursor"),
			Limit:  orb.F(int64(1)),
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
