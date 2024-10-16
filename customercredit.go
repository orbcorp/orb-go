// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
)

// CustomerCreditService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCreditService] method instead.
type CustomerCreditService struct {
	Options []option.RequestOption
	Ledger  *CustomerCreditLedgerService
	TopUps  *CustomerCreditTopUpService
}

// NewCustomerCreditService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCreditService(opts ...option.RequestOption) (r *CustomerCreditService) {
	r = &CustomerCreditService{}
	r.Options = opts
	r.Ledger = NewCustomerCreditLedgerService(opts...)
	r.TopUps = NewCustomerCreditTopUpService(opts...)
	return
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) List(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditListByExternalIDResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
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
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditListByExternalIDResponse] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

type CustomerCreditListResponse struct {
	ID                    string                           `json:"id,required"`
	Balance               float64                          `json:"balance,required"`
	EffectiveDate         time.Time                        `json:"effective_date,required,nullable" format:"date-time"`
	ExpiryDate            time.Time                        `json:"expiry_date,required,nullable" format:"date-time"`
	MaximumInitialBalance float64                          `json:"maximum_initial_balance,required,nullable"`
	PerUnitCostBasis      string                           `json:"per_unit_cost_basis,required,nullable"`
	Status                CustomerCreditListResponseStatus `json:"status,required"`
	JSON                  customerCreditListResponseJSON   `json:"-"`
}

// customerCreditListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditListResponse]
type customerCreditListResponseJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	MaximumInitialBalance apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseStatus string

const (
	CustomerCreditListResponseStatusActive         CustomerCreditListResponseStatus = "active"
	CustomerCreditListResponseStatusPendingPayment CustomerCreditListResponseStatus = "pending_payment"
)

func (r CustomerCreditListResponseStatus) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseStatusActive, CustomerCreditListResponseStatusPendingPayment:
		return true
	}
	return false
}

type CustomerCreditListByExternalIDResponse struct {
	ID                    string                                       `json:"id,required"`
	Balance               float64                                      `json:"balance,required"`
	EffectiveDate         time.Time                                    `json:"effective_date,required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                    `json:"expiry_date,required,nullable" format:"date-time"`
	MaximumInitialBalance float64                                      `json:"maximum_initial_balance,required,nullable"`
	PerUnitCostBasis      string                                       `json:"per_unit_cost_basis,required,nullable"`
	Status                CustomerCreditListByExternalIDResponseStatus `json:"status,required"`
	JSON                  customerCreditListByExternalIDResponseJSON   `json:"-"`
}

// customerCreditListByExternalIDResponseJSON contains the JSON metadata for the
// struct [CustomerCreditListByExternalIDResponse]
type customerCreditListByExternalIDResponseJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	MaximumInitialBalance apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerCreditListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListByExternalIDResponseStatus string

const (
	CustomerCreditListByExternalIDResponseStatusActive         CustomerCreditListByExternalIDResponseStatus = "active"
	CustomerCreditListByExternalIDResponseStatusPendingPayment CustomerCreditListByExternalIDResponseStatus = "pending_payment"
)

func (r CustomerCreditListByExternalIDResponseStatus) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseStatusActive, CustomerCreditListByExternalIDResponseStatusPendingPayment:
		return true
	}
	return false
}

type CustomerCreditListParams struct {
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// Include all blocks, not just active ones.
	IncludeAllBlocks param.Field[bool] `query:"include_all_blocks"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditListParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditListByExternalIDParams struct {
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// Include all blocks, not just active ones.
	IncludeAllBlocks param.Field[bool] `query:"include_all_blocks"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditListByExternalIDParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
