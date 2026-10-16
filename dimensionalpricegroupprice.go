// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
)

// DimensionalPriceGroupPriceService contains methods and other services that help
// with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDimensionalPriceGroupPriceService] method instead.
type DimensionalPriceGroupPriceService struct {
	Options []option.RequestOption
}

// NewDimensionalPriceGroupPriceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDimensionalPriceGroupPriceService(opts ...option.RequestOption) (r *DimensionalPriceGroupPriceService) {
	r = &DimensionalPriceGroupPriceService{}
	r.Options = opts
	return
}

// Create prices in a dimensional price group. Every price shares `item_id`,
// `cadence`, `currency`, and `model_type`. Tiered prices also share tier
// boundaries; each row supplies its own amounts. The request is one transaction.
func (r *DimensionalPriceGroupPriceService) New(ctx context.Context, dimensionalPriceGroupID string, body DimensionalPriceGroupPriceNewParams, opts ...option.RequestOption) (res *DimensionalPriceGroupPriceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dimensionalPriceGroupID == "" {
		err = errors.New("missing required dimensional_price_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dimensional_price_groups/%s/prices", dimensionalPriceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint returns the prices in a dimensional price group. Subscription
// overrides are omitted. Archived prices are omitted unless `status` is `archived`
// or `all`.
func (r *DimensionalPriceGroupPriceService) List(ctx context.Context, dimensionalPriceGroupID string, query DimensionalPriceGroupPriceListParams, opts ...option.RequestOption) (res *pagination.Page[shared.Price], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if dimensionalPriceGroupID == "" {
		err = errors.New("missing required dimensional_price_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dimensional_price_groups/%s/prices", dimensionalPriceGroupID)
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

// This endpoint returns the prices in a dimensional price group. Subscription
// overrides are omitted. Archived prices are omitted unless `status` is `archived`
// or `all`.
func (r *DimensionalPriceGroupPriceService) ListAutoPaging(ctx context.Context, dimensionalPriceGroupID string, query DimensionalPriceGroupPriceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.Price] {
	return pagination.NewPageAutoPager(r.List(ctx, dimensionalPriceGroupID, query, opts...))
}

type DimensionalPriceGroupPriceNewResponse struct {
	// The prices created by this request, in request order.
	Data []shared.Price                            `json:"data" api:"required"`
	JSON dimensionalPriceGroupPriceNewResponseJSON `json:"-"`
}

// dimensionalPriceGroupPriceNewResponseJSON contains the JSON metadata for the
// struct [DimensionalPriceGroupPriceNewResponse]
type dimensionalPriceGroupPriceNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DimensionalPriceGroupPriceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dimensionalPriceGroupPriceNewResponseJSON) RawJSON() string {
	return r.raw
}

type DimensionalPriceGroupPriceNewParams struct {
	// Billing cadence for every price
	Cadence param.Field[DimensionalPriceGroupPriceNewParamsCadence] `json:"cadence" api:"required"`
	// An ISO 4217 currency string for every price.
	Currency param.Field[string] `json:"currency" api:"required"`
	// The item for every price. Must be the group's billable metric item.
	ItemID    param.Field[string]                                       `json:"item_id" api:"required"`
	ModelType param.Field[DimensionalPriceGroupPriceNewParamsModelType] `json:"model_type" api:"required"`
	Prices    param.Field[[]DimensionalPriceGroupPriceNewParamsPrice]   `json:"prices" api:"required"`
	// Shared tier boundaries. Required for tiered prices and omitted for unit prices.
	Tiers param.Field[[]DimensionalPriceGroupPriceNewParamsTier] `json:"tiers"`
}

func (r DimensionalPriceGroupPriceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Billing cadence for every price
type DimensionalPriceGroupPriceNewParamsCadence string

const (
	DimensionalPriceGroupPriceNewParamsCadenceOneTime    DimensionalPriceGroupPriceNewParamsCadence = "one_time"
	DimensionalPriceGroupPriceNewParamsCadenceMonthly    DimensionalPriceGroupPriceNewParamsCadence = "monthly"
	DimensionalPriceGroupPriceNewParamsCadenceQuarterly  DimensionalPriceGroupPriceNewParamsCadence = "quarterly"
	DimensionalPriceGroupPriceNewParamsCadenceSemiAnnual DimensionalPriceGroupPriceNewParamsCadence = "semi_annual"
	DimensionalPriceGroupPriceNewParamsCadenceAnnual     DimensionalPriceGroupPriceNewParamsCadence = "annual"
)

func (r DimensionalPriceGroupPriceNewParamsCadence) IsKnown() bool {
	switch r {
	case DimensionalPriceGroupPriceNewParamsCadenceOneTime, DimensionalPriceGroupPriceNewParamsCadenceMonthly, DimensionalPriceGroupPriceNewParamsCadenceQuarterly, DimensionalPriceGroupPriceNewParamsCadenceSemiAnnual, DimensionalPriceGroupPriceNewParamsCadenceAnnual:
		return true
	}
	return false
}

type DimensionalPriceGroupPriceNewParamsModelType string

const (
	DimensionalPriceGroupPriceNewParamsModelTypeUnit   DimensionalPriceGroupPriceNewParamsModelType = "unit"
	DimensionalPriceGroupPriceNewParamsModelTypeTiered DimensionalPriceGroupPriceNewParamsModelType = "tiered"
)

func (r DimensionalPriceGroupPriceNewParamsModelType) IsKnown() bool {
	switch r {
	case DimensionalPriceGroupPriceNewParamsModelTypeUnit, DimensionalPriceGroupPriceNewParamsModelTypeTiered:
		return true
	}
	return false
}

type DimensionalPriceGroupPriceNewParamsPrice struct {
	// One amount for a unit price. One amount per tier for a tiered price, in tier
	// order.
	Amounts param.Field[[]float64] `json:"amounts" api:"required"`
	// Dimension values, in the group's dimension order.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Price name. Defaults to the dimension values.
	Name param.Field[string] `json:"name"`
}

func (r DimensionalPriceGroupPriceNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DimensionalPriceGroupPriceNewParamsTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit" api:"required"`
	// Inclusive tier ending value. Null on the last tier.
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r DimensionalPriceGroupPriceNewParamsTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DimensionalPriceGroupPriceListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// active omits archived prices. archived returns only archived prices. all returns
	// both.
	Status param.Field[DimensionalPriceGroupPriceListParamsStatus] `query:"status"`
}

// URLQuery serializes [DimensionalPriceGroupPriceListParams]'s query parameters as
// `url.Values`.
func (r DimensionalPriceGroupPriceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// active omits archived prices. archived returns only archived prices. all returns
// both.
type DimensionalPriceGroupPriceListParamsStatus string

const (
	DimensionalPriceGroupPriceListParamsStatusActive   DimensionalPriceGroupPriceListParamsStatus = "active"
	DimensionalPriceGroupPriceListParamsStatusArchived DimensionalPriceGroupPriceListParamsStatus = "archived"
	DimensionalPriceGroupPriceListParamsStatusAll      DimensionalPriceGroupPriceListParamsStatus = "all"
)

func (r DimensionalPriceGroupPriceListParamsStatus) IsKnown() bool {
	switch r {
	case DimensionalPriceGroupPriceListParamsStatusActive, DimensionalPriceGroupPriceListParamsStatusArchived, DimensionalPriceGroupPriceListParamsStatusAll:
		return true
	}
	return false
}
