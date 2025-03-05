// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// DimensionalPriceGroupExternalDimensionalPriceGroupIDService contains methods and
// other services that help with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDimensionalPriceGroupExternalDimensionalPriceGroupIDService] method
// instead.
type DimensionalPriceGroupExternalDimensionalPriceGroupIDService struct {
	Options []option.RequestOption
}

// NewDimensionalPriceGroupExternalDimensionalPriceGroupIDService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewDimensionalPriceGroupExternalDimensionalPriceGroupIDService(opts ...option.RequestOption) (r *DimensionalPriceGroupExternalDimensionalPriceGroupIDService) {
	r = &DimensionalPriceGroupExternalDimensionalPriceGroupIDService{}
	r.Options = opts
	return
}

// Fetch dimensional price group by external ID
func (r *DimensionalPriceGroupExternalDimensionalPriceGroupIDService) Get(ctx context.Context, externalDimensionalPriceGroupID string, opts ...option.RequestOption) (res *DimensionalPriceGroup, err error) {
	opts = append(r.Options[:], opts...)
	if externalDimensionalPriceGroupID == "" {
		err = errors.New("missing required external_dimensional_price_group_id parameter")
		return
	}
	path := fmt.Sprintf("dimensional_price_groups/external_dimensional_price_group_id/%s", externalDimensionalPriceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
