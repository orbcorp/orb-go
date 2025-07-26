// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
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

// This endpoint can be used to update the `external_dimensional_price_group_id`
// and `metadata` of an existing dimensional price group. Other fields on a
// dimensional price group are currently immutable.
func (r *DimensionalPriceGroupExternalDimensionalPriceGroupIDService) Update(ctx context.Context, externalDimensionalPriceGroupID string, body DimensionalPriceGroupExternalDimensionalPriceGroupIDUpdateParams, opts ...option.RequestOption) (res *DimensionalPriceGroup, err error) {
	opts = append(r.Options[:], opts...)
	if externalDimensionalPriceGroupID == "" {
		err = errors.New("missing required external_dimensional_price_group_id parameter")
		return
	}
	path := fmt.Sprintf("dimensional_price_groups/external_dimensional_price_group_id/%s", externalDimensionalPriceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type DimensionalPriceGroupExternalDimensionalPriceGroupIDUpdateParams struct {
	// An optional user-defined ID for this dimensional price group resource, used
	// throughout the system as an alias for this dimensional price group. Use this
	// field to identify a dimensional price group by an existing identifier in your
	// system.
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r DimensionalPriceGroupExternalDimensionalPriceGroupIDUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
