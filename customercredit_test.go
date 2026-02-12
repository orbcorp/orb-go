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

func TestCustomerCreditListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.List(
		context.TODO(),
		"customer_id",
		orb.CustomerCreditListParams{
			Currency:         orb.F("currency"),
			Cursor:           orb.F("cursor"),
			EffectiveDateGt:  orb.F(time.Now()),
			EffectiveDateGte: orb.F(time.Now()),
			EffectiveDateLt:  orb.F(time.Now()),
			EffectiveDateLte: orb.F(time.Now()),
			IncludeAllBlocks: orb.F(true),
			Limit:            orb.F(int64(1)),
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

func TestCustomerCreditListByExternalIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Credits.ListByExternalID(
		context.TODO(),
		"external_customer_id",
		orb.CustomerCreditListByExternalIDParams{
			Currency:         orb.F("currency"),
			Cursor:           orb.F("cursor"),
			EffectiveDateGt:  orb.F(time.Now()),
			EffectiveDateGte: orb.F(time.Now()),
			EffectiveDateLt:  orb.F(time.Now()),
			EffectiveDateLte: orb.F(time.Now()),
			IncludeAllBlocks: orb.F(true),
			Limit:            orb.F(int64(1)),
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
