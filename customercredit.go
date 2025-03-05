// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
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
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) List(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CustomerCreditBalancesModelData], err error) {
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
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CustomerCreditBalancesModelData] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[shared.CustomerCreditBalancesModelData], err error) {
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
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
func (r *CustomerCreditService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CustomerCreditBalancesModelData] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

type CustomerCreditListParams struct {
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// If set to True, all expired and depleted blocks, as well as active block will be
	// returned.
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
	// If set to True, all expired and depleted blocks, as well as active block will be
	// returned.
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
