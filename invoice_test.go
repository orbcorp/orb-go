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

func TestInvoiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.New(context.TODO(), orb.InvoiceNewParams{
		Currency:    orb.F("USD"),
		InvoiceDate: orb.F(time.Now()),
		LineItems: orb.F([]orb.InvoiceNewParamsLineItem{{
			StartDate: orb.F(time.Now()),
			EndDate:   orb.F(time.Now()),
			Quantity:  orb.F(1.000000),
			Name:      orb.F("Line Item Name"),
			ItemID:    orb.F("4khy3nwzktxv7"),
			ModelType: orb.F(orb.InvoiceNewParamsLineItemsModelTypeUnit),
			UnitConfig: orb.F(orb.InvoiceNewParamsLineItemsUnitConfig{
				UnitAmount: orb.F("string"),
			}),
		}, {
			StartDate: orb.F(time.Now()),
			EndDate:   orb.F(time.Now()),
			Quantity:  orb.F(1.000000),
			Name:      orb.F("Line Item Name"),
			ItemID:    orb.F("4khy3nwzktxv7"),
			ModelType: orb.F(orb.InvoiceNewParamsLineItemsModelTypeUnit),
			UnitConfig: orb.F(orb.InvoiceNewParamsLineItemsUnitConfig{
				UnitAmount: orb.F("string"),
			}),
		}, {
			StartDate: orb.F(time.Now()),
			EndDate:   orb.F(time.Now()),
			Quantity:  orb.F(1.000000),
			Name:      orb.F("Line Item Name"),
			ItemID:    orb.F("4khy3nwzktxv7"),
			ModelType: orb.F(orb.InvoiceNewParamsLineItemsModelTypeUnit),
			UnitConfig: orb.F(orb.InvoiceNewParamsLineItemsUnitConfig{
				UnitAmount: orb.F("string"),
			}),
		}}),
		NetTerms:   orb.F(int64(0)),
		CustomerID: orb.F("4khy3nwzktxv7"),
		Discount: orb.F[shared.DiscountUnionParam](shared.DiscountPercentageDiscountParam{
			DiscountType:       orb.F(shared.DiscountPercentageDiscountDiscountTypePercentage),
			AppliesToPriceIDs:  orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
			Reason:             orb.F("string"),
			PercentageDiscount: orb.F(0.150000),
		}),
		ExternalCustomerID: orb.F("external-customer-id"),
		Memo:               orb.F("An optional memo for my invoice."),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		WillAutoIssue: orb.F(false),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.List(context.TODO(), orb.InvoiceListParams{
		Amount:             orb.F("string"),
		AmountGt:           orb.F("string"),
		AmountLt:           orb.F("string"),
		Cursor:             orb.F("string"),
		CustomerID:         orb.F("string"),
		DateType:           orb.F(orb.InvoiceListParamsDateTypeDueDate),
		DueDate:            orb.F(time.Now()),
		DueDateWindow:      orb.F("string"),
		DueDateGt:          orb.F(time.Now()),
		DueDateLt:          orb.F(time.Now()),
		ExternalCustomerID: orb.F("string"),
		InvoiceDateGt:      orb.F(time.Now()),
		InvoiceDateGte:     orb.F(time.Now()),
		InvoiceDateLt:      orb.F(time.Now()),
		InvoiceDateLte:     orb.F(time.Now()),
		IsRecurring:        orb.F(true),
		Limit:              orb.F(int64(1)),
		Status:             orb.F([]orb.InvoiceListParamsStatus{orb.InvoiceListParamsStatusDraft, orb.InvoiceListParamsStatusIssued, orb.InvoiceListParamsStatusPaid}),
		SubscriptionID:     orb.F("string"),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceFetch(t *testing.T) {
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
	_, err := client.Invoices.Fetch(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceFetchUpcomingWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.FetchUpcoming(context.TODO(), orb.InvoiceFetchUpcomingParams{
		SubscriptionID: orb.F("string"),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceIssue(t *testing.T) {
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
	_, err := client.Invoices.Issue(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceMarkPaidWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.MarkPaid(
		context.TODO(),
		"string",
		orb.InvoiceMarkPaidParams{
			PaymentReceivedDate: orb.F(time.Now()),
			ExternalID:          orb.F("external_payment_id_123"),
			Notes:               orb.F("string"),
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

func TestInvoiceVoid(t *testing.T) {
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
	_, err := client.Invoices.Void(context.TODO(), "string")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
