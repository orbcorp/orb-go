// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
)

// CustomerCreditService contains methods and other services that help with
// interacting with the orb API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomerCreditService] method instead.
type CustomerCreditService struct {
	Options []option.RequestOption
	Ledger  *CustomerCreditLedgerService
}

// NewCustomerCreditService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCreditService(opts ...option.RequestOption) (r *CustomerCreditService) {
	r = &CustomerCreditService{}
	r.Options = opts
	r.Ledger = NewCustomerCreditLedgerService(opts...)
	return
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
func (r *CustomerCreditService) List(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) (res *shared.Page[CustomerCreditListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/%s/credits", customerID)
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

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
func (r *CustomerCreditService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerCreditListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
func (r *CustomerCreditService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) (res *shared.Page[CustomerCreditListByExternalIDResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s/credits", externalCustomerID)
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

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
func (r *CustomerCreditService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerCreditListByExternalIDResponse] {
	return shared.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

type CustomerCreditListResponse struct {
	ID               string    `json:"id,required"`
	Balance          float64   `json:"balance,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditListResponseJSON
}

// customerCreditListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditListResponse]
type customerCreditListResponseJSON struct {
	ID               apijson.Field
	Balance          apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditListByExternalIDResponse struct {
	ID               string    `json:"id,required"`
	Balance          float64   `json:"balance,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditListByExternalIDResponseJSON
}

// customerCreditListByExternalIDResponseJSON contains the JSON metadata for the
// struct [CustomerCreditListByExternalIDResponse]
type customerCreditListByExternalIDResponseJSON struct {
	ID               apijson.Field
	Balance          apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditListParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditListByExternalIDParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditListByExternalIDParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
