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

func TestDimensionalPriceGroupPriceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DimensionalPriceGroups.Prices.New(
		context.TODO(),
		"dimensional_price_group_id",
		orb.DimensionalPriceGroupPriceNewParams{
			Cadence:   orb.F(orb.DimensionalPriceGroupPriceNewParamsCadenceOneTime),
			Currency:  orb.F("USD"),
			ItemID:    orb.F("item_id"),
			ModelType: orb.F(orb.DimensionalPriceGroupPriceNewParamsModelTypeUnit),
			Prices: orb.F([]orb.DimensionalPriceGroupPriceNewParamsPrice{{
				Amounts:         orb.F([]float64{0.000000}),
				DimensionValues: orb.F([]string{"x"}),
				Name:            orb.F("name"),
			}}),
			Tiers: orb.F([]orb.DimensionalPriceGroupPriceNewParamsTier{{
				FirstUnit: orb.F(0.000000),
				LastUnit:  orb.F(0.000000),
			}}),
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

func TestDimensionalPriceGroupPriceListWithOptionalParams(t *testing.T) {
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
	_, err := client.DimensionalPriceGroups.Prices.List(
		context.TODO(),
		"dimensional_price_group_id",
		orb.DimensionalPriceGroupPriceListParams{
			Cursor: orb.F("cursor"),
			Limit:  orb.F(int64(1)),
			Status: orb.F(orb.DimensionalPriceGroupPriceListParamsStatusActive),
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
