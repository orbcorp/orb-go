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

func TestInvoiceLineItemNew(t *testing.T) {
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
	_, err := client.InvoiceLineItems.New(context.TODO(), orb.InvoiceLineItemNewParams{
		Amount:    orb.F("12.00"),
		EndDate:   orb.F(time.Now()),
		InvoiceID: orb.F("4khy3nwzktxv7"),
		Name:      orb.F("Item Name"),
		Quantity:  orb.F(1.000000),
		StartDate: orb.F(time.Now()),
	})
	if err != nil {
		var apierr *orb.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
