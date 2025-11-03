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
		"customer_id",
		orb.CustomerCreditLedgerListParams{
			CreatedAtGt:   orb.F(time.Now()),
			CreatedAtGte:  orb.F(time.Now()),
			CreatedAtLt:   orb.F(time.Now()),
			CreatedAtLte:  orb.F(time.Now()),
			Currency:      orb.F("currency"),
			Cursor:        orb.F("cursor"),
			EntryStatus:   orb.F(orb.CustomerCreditLedgerListParamsEntryStatusCommitted),
			EntryType:     orb.F(orb.CustomerCreditLedgerListParamsEntryTypeIncrement),
			Limit:         orb.F(int64(1)),
			MinimumAmount: orb.F("minimum_amount"),
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
		"customer_id",
		orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParams{
			Amount:        orb.F(0.000000),
			EntryType:     orb.F(orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement),
			Currency:      orb.F("currency"),
			Description:   orb.F("description"),
			EffectiveDate: orb.F(time.Now()),
			ExpiryDate:    orb.F(time.Now()),
			Filters: orb.F([]orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsFilter{{
				Field:    orb.F(orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsFiltersFieldItemID),
				Operator: orb.F(orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsFiltersOperatorIncludes),
				Values:   orb.F([]string{"string"}),
			}}),
			InvoiceSettings: orb.F(orb.CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings{
				AutoCollection:           orb.F(true),
				CustomDueDate:            orb.F(time.Now()),
				InvoiceDate:              orb.F(time.Now()),
				ItemID:                   orb.F("item_id"),
				Memo:                     orb.F("memo"),
				NetTerms:                 orb.F(int64(0)),
				RequireSuccessfulPayment: orb.F(true),
			}),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			PerUnitCostBasis: orb.F("per_unit_cost_basis"),
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
		"external_customer_id",
		orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParams{
			Amount:        orb.F(0.000000),
			EntryType:     orb.F(orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement),
			Currency:      orb.F("currency"),
			Description:   orb.F("description"),
			EffectiveDate: orb.F(time.Now()),
			ExpiryDate:    orb.F(time.Now()),
			Filters: orb.F([]orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsFilter{{
				Field:    orb.F(orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsFiltersFieldItemID),
				Operator: orb.F(orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsFiltersOperatorIncludes),
				Values:   orb.F([]string{"string"}),
			}}),
			InvoiceSettings: orb.F(orb.CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings{
				AutoCollection:           orb.F(true),
				CustomDueDate:            orb.F(time.Now()),
				InvoiceDate:              orb.F(time.Now()),
				ItemID:                   orb.F("item_id"),
				Memo:                     orb.F("memo"),
				NetTerms:                 orb.F(int64(0)),
				RequireSuccessfulPayment: orb.F(true),
			}),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			PerUnitCostBasis: orb.F("per_unit_cost_basis"),
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
		"external_customer_id",
		orb.CustomerCreditLedgerListByExternalIDParams{
			CreatedAtGt:   orb.F(time.Now()),
			CreatedAtGte:  orb.F(time.Now()),
			CreatedAtLt:   orb.F(time.Now()),
			CreatedAtLte:  orb.F(time.Now()),
			Currency:      orb.F("currency"),
			Cursor:        orb.F("cursor"),
			EntryStatus:   orb.F(orb.CustomerCreditLedgerListByExternalIDParamsEntryStatusCommitted),
			EntryType:     orb.F(orb.CustomerCreditLedgerListByExternalIDParamsEntryTypeIncrement),
			Limit:         orb.F(int64(1)),
			MinimumAmount: orb.F("minimum_amount"),
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
