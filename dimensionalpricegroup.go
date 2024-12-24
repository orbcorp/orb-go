// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
)

// DimensionalPriceGroupService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDimensionalPriceGroupService] method instead.
type DimensionalPriceGroupService struct {
	Options                         []option.RequestOption
	ExternalDimensionalPriceGroupID *DimensionalPriceGroupExternalDimensionalPriceGroupIDService
}

// NewDimensionalPriceGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDimensionalPriceGroupService(opts ...option.RequestOption) (r *DimensionalPriceGroupService) {
	r = &DimensionalPriceGroupService{}
	r.Options = opts
	r.ExternalDimensionalPriceGroupID = NewDimensionalPriceGroupExternalDimensionalPriceGroupIDService(opts...)
	return
}

// A dimensional price group is used to partition the result of a billable metric
// by a set of dimensions. Prices in a price group must specify the parition used
// to derive their usage.
//
// For example, suppose we have a billable metric that measures the number of
// widgets used and we want to charge differently depending on the color of the
// widget. We can create a price group with a dimension "color" and two prices: one
// that charges $10 per red widget and one that charges $20 per blue widget.
func (r *DimensionalPriceGroupService) New(ctx context.Context, body DimensionalPriceGroupNewParams, opts ...option.RequestOption) (res *DimensionalPriceGroup, err error) {
	opts = append(r.Options[:], opts...)
	path := "dimensional_price_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch dimensional price group
func (r *DimensionalPriceGroupService) Get(ctx context.Context, dimensionalPriceGroupID string, opts ...option.RequestOption) (res *DimensionalPriceGroup, err error) {
	opts = append(r.Options[:], opts...)
	if dimensionalPriceGroupID == "" {
		err = errors.New("missing required dimensional_price_group_id parameter")
		return
	}
	path := fmt.Sprintf("dimensional_price_groups/%s", dimensionalPriceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List dimensional price groups
func (r *DimensionalPriceGroupService) List(ctx context.Context, query DimensionalPriceGroupListParams, opts ...option.RequestOption) (res *pagination.Page[DimensionalPriceGroup], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "dimensional_price_groups"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List dimensional price groups
func (r *DimensionalPriceGroupService) ListAutoPaging(ctx context.Context, query DimensionalPriceGroupListParams, opts ...option.RequestOption) *pagination.PageAutoPager[DimensionalPriceGroup] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// A dimensional price group is used to partition the result of a billable metric
// by a set of dimensions. Prices in a price group must specify the parition used
// to derive their usage.
type DimensionalPriceGroup struct {
	ID string `json:"id,required"`
	// The billable metric associated with this dimensional price group. All prices
	// associated with this dimensional price group will be computed using this
	// billable metric.
	BillableMetricID string `json:"billable_metric_id,required"`
	// The dimensions that this dimensional price group is defined over
	Dimensions []string `json:"dimensions,required"`
	// An alias for the dimensional price group
	ExternalDimensionalPriceGroupID string `json:"external_dimensional_price_group_id,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The name of the dimensional price group
	Name string                    `json:"name,required"`
	JSON dimensionalPriceGroupJSON `json:"-"`
}

// dimensionalPriceGroupJSON contains the JSON metadata for the struct
// [DimensionalPriceGroup]
type dimensionalPriceGroupJSON struct {
	ID                              apijson.Field
	BillableMetricID                apijson.Field
	Dimensions                      apijson.Field
	ExternalDimensionalPriceGroupID apijson.Field
	Metadata                        apijson.Field
	Name                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *DimensionalPriceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dimensionalPriceGroupJSON) RawJSON() string {
	return r.raw
}

type DimensionalPriceGroups struct {
	Data               []DimensionalPriceGroup    `json:"data,required"`
	PaginationMetadata shared.PaginationMetadata  `json:"pagination_metadata,required"`
	JSON               dimensionalPriceGroupsJSON `json:"-"`
}

// dimensionalPriceGroupsJSON contains the JSON metadata for the struct
// [DimensionalPriceGroups]
type dimensionalPriceGroupsJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DimensionalPriceGroups) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dimensionalPriceGroupsJSON) RawJSON() string {
	return r.raw
}

type DimensionalPriceGroupNewParams struct {
	BillableMetricID param.Field[string] `json:"billable_metric_id,required"`
	// The set of keys (in order) used to disambiguate prices in the group.
	Dimensions                      param.Field[[]string] `json:"dimensions,required"`
	Name                            param.Field[string]   `json:"name,required"`
	ExternalDimensionalPriceGroupID param.Field[string]   `json:"external_dimensional_price_group_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r DimensionalPriceGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DimensionalPriceGroupListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [DimensionalPriceGroupListParams]'s query parameters as
// `url.Values`.
func (r DimensionalPriceGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
