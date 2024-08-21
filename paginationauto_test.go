// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb_test

import (
	"context"
	"os"
	"testing"

	"github.com/orbcorp/orb-go"
	"github.com/orbcorp/orb-go/internal/testutil"
	"github.com/orbcorp/orb-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Coupons.ListAutoPaging(context.TODO(), orb.CouponListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		coupon := iter.Current()
		t.Logf("%+v\n", coupon.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
