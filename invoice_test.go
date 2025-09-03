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
			EndDate:   orb.F(time.Now()),
			ItemID:    orb.F("4khy3nwzktxv7"),
			ModelType: orb.F(orb.InvoiceNewParamsLineItemsModelTypeUnit),
			Name:      orb.F("Line Item Name"),
			Quantity:  orb.F(1.000000),
			StartDate: orb.F(time.Now()),
			UnitConfig: orb.F(shared.UnitConfigParam{
				UnitAmount: orb.F("unit_amount"),
			}),
		}}),
		CustomerID: orb.F("4khy3nwzktxv7"),
		Discount: orb.F[shared.DiscountUnionParam](shared.PercentageDiscountParam{
			DiscountType:       orb.F(shared.PercentageDiscountDiscountTypePercentage),
			PercentageDiscount: orb.F(0.150000),
			AppliesToPriceIDs:  orb.F([]string{"h74gfhdjvn7ujokd", "7hfgtgjnbvc3ujkl"}),
			Filters: orb.F([]shared.TransformPriceFilterParam{{
				Field:    orb.F(shared.TransformPriceFilterFieldPriceID),
				Operator: orb.F(shared.TransformPriceFilterOperatorIncludes),
				Values:   orb.F([]string{"string"}),
			}}),
			Reason: orb.F("reason"),
		}),
		DueDate:            orb.F(time.Now()),
		ExternalCustomerID: orb.F("external-customer-id"),
		Memo:               orb.F("An optional memo for my invoice."),
		Metadata: orb.F(map[string]string{
			"foo": "string",
		}),
		NetTerms:      orb.F(int64(0)),
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

func TestInvoiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.Update(
		context.TODO(),
		"invoice_id",
		orb.InvoiceUpdateParams{
			DueDate: orb.F(time.Now()),
			Metadata: orb.F(map[string]string{
				"foo": "string",
			}),
			NetTerms: orb.F(int64(0)),
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
		Amount:             orb.F("amount"),
		AmountGt:           orb.F("amount[gt]"),
		AmountLt:           orb.F("amount[lt]"),
		Cursor:             orb.F("cursor"),
		CustomerID:         orb.F("customer_id"),
		DateType:           orb.F(orb.InvoiceListParamsDateTypeDueDate),
		DueDate:            orb.F(time.Now()),
		DueDateWindow:      orb.F("due_date_window"),
		DueDateGt:          orb.F(time.Now()),
		DueDateLt:          orb.F(time.Now()),
		ExternalCustomerID: orb.F("external_customer_id"),
		InvoiceDateGt:      orb.F(time.Now()),
		InvoiceDateGte:     orb.F(time.Now()),
		InvoiceDateLt:      orb.F(time.Now()),
		InvoiceDateLte:     orb.F(time.Now()),
		IsRecurring:        orb.F(true),
		Limit:              orb.F(int64(1)),
		Status:             orb.F([]orb.InvoiceListParamsStatus{orb.InvoiceListParamsStatusDraft}),
		SubscriptionID:     orb.F("subscription_id"),
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
	_, err := client.Invoices.Fetch(context.TODO(), "invoice_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceFetchUpcoming(t *testing.T) {
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
		SubscriptionID: orb.F("subscription_id"),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceIssueWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.Issue(
		context.TODO(),
		"invoice_id",
		orb.InvoiceIssueParams{
			Synchronous: orb.F(true),
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
		"invoice_id",
		orb.InvoiceMarkPaidParams{
			PaymentReceivedDate: orb.F(time.Now()),
			ExternalID:          orb.F("external_payment_id_123"),
			Notes:               orb.F("notes"),
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

func TestInvoicePay(t *testing.T) {
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
	_, err := client.Invoices.Pay(context.TODO(), "invoice_id")
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
	_, err := client.Invoices.Void(context.TODO(), "invoice_id")
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
