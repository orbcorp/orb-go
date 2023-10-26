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

func TestCustomerCreditLedgerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.Ledger.List(
		context.TODO(),
		"string",
		orb.CustomerCreditLedgerListParams{
			CreatedAtGt:   orb.F(time.Now()),
			CreatedAtGte:  orb.F(time.Now()),
			CreatedAtLt:   orb.F(time.Now()),
			CreatedAtLte:  orb.F(time.Now()),
			Currency:      orb.F("string"),
			Cursor:        orb.F("string"),
			EntryStatus:   orb.F(orb.CustomerCreditLedgerListParamsEntryStatusCommitted),
			EntryType:     orb.F(orb.CustomerCreditLedgerListParamsEntryTypeIncrement),
			Limit:         orb.F(int64(0)),
			MinimumAmount: orb.F("string"),
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

func TestCustomerCreditLedgerNewEntryWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.Ledger.NewEntry(
		context.TODO(),
		"string",
		orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParams{
			Amount:        orb.F(0.000000),
			EntryType:     orb.F(orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement),
			Description:   orb.F("string"),
			EffectiveDate: orb.F(time.Now()),
			ExpiryDate:    orb.F(time.Now()),
			InvoiceSettings: orb.F(orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings{
				AutoCollection: orb.F(true),
				NetTerms:       orb.F(int64(0)),
				Memo:           orb.F("string"),
			}),
			Metadata:         orb.F[any](map[string]interface{}{}),
			PerUnitCostBasis: orb.F("string"),
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

func TestCustomerCreditLedgerNewEntryByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.Ledger.NewEntryByExternalID(
		context.TODO(),
		"string",
		orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParams{
			Amount:        orb.F(0.000000),
			EntryType:     orb.F(orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement),
			Description:   orb.F("string"),
			EffectiveDate: orb.F(time.Now()),
			ExpiryDate:    orb.F(time.Now()),
			InvoiceSettings: orb.F(orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings{
				AutoCollection: orb.F(true),
				NetTerms:       orb.F(int64(0)),
				Memo:           orb.F("string"),
			}),
			Metadata:         orb.F[any](map[string]interface{}{}),
			PerUnitCostBasis: orb.F("string"),
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

func TestCustomerCreditLedgerListByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.Ledger.ListByExternalID(
		context.TODO(),
		"string",
		orb.CustomerCreditLedgerListByExternalIDParams{
			CreatedAtGt:   orb.F(time.Now()),
			CreatedAtGte:  orb.F(time.Now()),
			CreatedAtLt:   orb.F(time.Now()),
			CreatedAtLte:  orb.F(time.Now()),
			Currency:      orb.F("string"),
			Cursor:        orb.F("string"),
			EntryStatus:   orb.F(orb.CustomerCreditLedgerListByExternalIDParamsEntryStatusCommitted),
			EntryType:     orb.F(orb.CustomerCreditLedgerListByExternalIDParamsEntryTypeIncrement),
			Limit:         orb.F(int64(0)),
			MinimumAmount: orb.F("string"),
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
