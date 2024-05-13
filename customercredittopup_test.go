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
		"string",
		orb.CustomerCreditTopUpNewParams{
			Amount:   orb.F("string"),
			Currency: orb.F("string"),
			InvoiceSettings: orb.F(orb.CustomerCreditTopUpNewParamsInvoiceSettings{
				AutoCollection:           orb.F(true),
				NetTerms:                 orb.F(int64(0)),
				Memo:                     orb.F("string"),
				RequireSuccessfulPayment: orb.F(true),
			}),
			PerUnitCostBasis: orb.F("string"),
			Threshold:        orb.F("string"),
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
		"string",
		orb.CustomerCreditTopUpListParams{
			Cursor: orb.F("string"),
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
		"string",
		"string",
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
		"string",
		orb.CustomerCreditTopUpNewByExternalIDParams{
			Amount:   orb.F("string"),
			Currency: orb.F("string"),
			InvoiceSettings: orb.F(orb.CustomerCreditTopUpNewByExternalIDParamsInvoiceSettings{
				AutoCollection:           orb.F(true),
				NetTerms:                 orb.F(int64(0)),
				Memo:                     orb.F("string"),
				RequireSuccessfulPayment: orb.F(true),
			}),
			PerUnitCostBasis: orb.F("string"),
			Threshold:        orb.F("string"),
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
		"string",
		"string",
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
		"string",
		orb.CustomerCreditTopUpListByExternalIDParams{
			Cursor: orb.F("string"),
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
