// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
)

// The Price resource represents a price that can be billed on a subscription,
// resulting in a charge on an invoice in the form of an invoice line item. Prices
// take a quantity and determine an amount to bill.
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given Price object. The model_type field
// determines the key for the configuration object that is present.
//
// For more on the types of prices, see
// [the core concepts documentation](/core-concepts#plan-and-price)
//
// PriceExternalPriceIDService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPriceExternalPriceIDService] method instead.
type PriceExternalPriceIDService struct {
	Options []option.RequestOption
}

// NewPriceExternalPriceIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPriceExternalPriceIDService(opts ...option.RequestOption) (r *PriceExternalPriceIDService) {
	r = &PriceExternalPriceIDService{}
	r.Options = opts
	return
}

// This endpoint allows you to update the `metadata` property on a price. If you
// pass null for the metadata value, it will clear any existing metadata for that
// price.
func (r *PriceExternalPriceIDService) Update(ctx context.Context, externalPriceID string, body PriceExternalPriceIDUpdateParams, opts ...option.RequestOption) (res *shared.Price, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPriceID == "" {
		err = errors.New("missing required external_price_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("prices/external_price_id/%s", externalPriceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// This endpoint returns a price given an external price id. See the
// [price creation API](/api-reference/price/create-price) for more information
// about external price aliases.
func (r *PriceExternalPriceIDService) Fetch(ctx context.Context, externalPriceID string, opts ...option.RequestOption) (res *shared.Price, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPriceID == "" {
		err = errors.New("missing required external_price_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("prices/external_price_id/%s", externalPriceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PriceExternalPriceIDUpdateParams struct {
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PriceExternalPriceIDUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
