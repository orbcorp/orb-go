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

func TestLicenseExternalLicenseGetUsageWithOptionalParams(t *testing.T) {
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
	_, err := client.Licenses.ExternalLicenses.GetUsage(
		context.TODO(),
		"external_license_id",
		orb.LicenseExternalLicenseGetUsageParams{
			LicenseTypeID:  orb.F("license_type_id"),
			SubscriptionID: orb.F("subscription_id"),
			Cursor:         orb.F("cursor"),
			EndDate:        orb.F(time.Now()),
			GroupBy:        orb.F([]string{"string"}),
			Limit:          orb.F(int64(1)),
			StartDate:      orb.F(time.Now()),
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
