// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

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

// This endpoint returns a price given an external price id. See the
// [price creation API](../reference/create-price) for more information about
// external price aliases.
func (r *PriceExternalPriceIDService) Fetch(ctx context.Context, externalPriceID string, opts ...option.RequestOption) (res *Price, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("prices/external_price_id/%s", externalPriceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
