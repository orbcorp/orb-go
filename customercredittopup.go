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
	"github.com/orbcorp/orb-go/shared"
)

// CustomerCreditTopUpService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCreditTopUpService] method instead.
type CustomerCreditTopUpService struct {
	Options []option.RequestOption
}

// NewCustomerCreditTopUpService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerCreditTopUpService(opts ...option.RequestOption) (r *CustomerCreditTopUpService) {
	r = &CustomerCreditTopUpService{}
	r.Options = opts
	return
}

// This endpoint allows you to create a new top-up for a specified customer's
// balance. While this top-up is active, the customer's balance will added in
// increments of the specified amount whenever the balance reaches the specified
// threshold.
//
// If a top-up already exists for this customer in the same currency, the existing
// top-up will be replaced.
func (r *CustomerCreditTopUpService) New(ctx context.Context, customerID string, body CustomerCreditTopUpNewParams, opts ...option.RequestOption) (res *shared.TopUpModel, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/credits/top_ups", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List top-ups
func (r *CustomerCreditTopUpService) List(ctx context.Context, customerID string, query CustomerCreditTopUpListParams, opts ...option.RequestOption) (res *pagination.Page[shared.TopUpModel], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/credits/top_ups", customerID)
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

// List top-ups
func (r *CustomerCreditTopUpService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditTopUpListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.TopUpModel] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// This deactivates the top-up and voids any invoices associated with pending
// credit blocks purchased through the top-up.
func (r *CustomerCreditTopUpService) Delete(ctx context.Context, customerID string, topUpID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if topUpID == "" {
		err = errors.New("missing required top_up_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/credits/top_ups/%s", customerID, topUpID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This endpoint allows you to create a new top-up for a specified customer's
// balance. While this top-up is active, the customer's balance will added in
// increments of the specified amount whenever the balance reaches the specified
// threshold.
//
// If a top-up already exists for this customer in the same currency, the existing
// top-up will be replaced.
func (r *CustomerCreditTopUpService) NewByExternalID(ctx context.Context, externalCustomerID string, body CustomerCreditTopUpNewByExternalIDParams, opts ...option.RequestOption) (res *shared.TopUpModel, err error) {
	opts = append(r.Options[:], opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/top_ups", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This deactivates the top-up and voids any invoices associated with pending
// credit blocks purchased through the top-up.
func (r *CustomerCreditTopUpService) DeleteByExternalID(ctx context.Context, externalCustomerID string, topUpID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
	if topUpID == "" {
		err = errors.New("missing required top_up_id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/top_ups/%s", externalCustomerID, topUpID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List top-ups by external ID
func (r *CustomerCreditTopUpService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditTopUpListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[shared.TopUpModel], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/top_ups", externalCustomerID)
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

// List top-ups by external ID
func (r *CustomerCreditTopUpService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditTopUpListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.TopUpModel] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

type CustomerCreditTopUpNewParams struct {
	// The amount to increment when the threshold is reached.
	Amount param.Field[string] `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings param.Field[CustomerCreditTopUpNewParamsInvoiceSettings] `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold param.Field[string] `json:"threshold,required"`
	// The date from which the top-up is active. If unspecified, the top-up is active
	// immediately.
	ActiveFrom param.Field[time.Time] `json:"active_from" format:"date-time"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter param.Field[int64] `json:"expires_after"`
	// The unit of expires_after.
	ExpiresAfterUnit param.Field[CustomerCreditTopUpNewParamsExpiresAfterUnit] `json:"expires_after_unit"`
}

func (r CustomerCreditTopUpNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for invoices generated by triggered top-ups.
type CustomerCreditTopUpNewParamsInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection param.Field[bool] `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo param.Field[string] `json:"memo"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment param.Field[bool] `json:"require_successful_payment"`
}

func (r CustomerCreditTopUpNewParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of expires_after.
type CustomerCreditTopUpNewParamsExpiresAfterUnit string

const (
	CustomerCreditTopUpNewParamsExpiresAfterUnitDay   CustomerCreditTopUpNewParamsExpiresAfterUnit = "day"
	CustomerCreditTopUpNewParamsExpiresAfterUnitMonth CustomerCreditTopUpNewParamsExpiresAfterUnit = "month"
)

func (r CustomerCreditTopUpNewParamsExpiresAfterUnit) IsKnown() bool {
	switch r {
	case CustomerCreditTopUpNewParamsExpiresAfterUnitDay, CustomerCreditTopUpNewParamsExpiresAfterUnitMonth:
		return true
	}
	return false
}

type CustomerCreditTopUpListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditTopUpListParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditTopUpListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditTopUpNewByExternalIDParams struct {
	// The amount to increment when the threshold is reached.
	Amount param.Field[string] `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings param.Field[CustomerCreditTopUpNewByExternalIDParamsInvoiceSettings] `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold param.Field[string] `json:"threshold,required"`
	// The date from which the top-up is active. If unspecified, the top-up is active
	// immediately.
	ActiveFrom param.Field[time.Time] `json:"active_from" format:"date-time"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter param.Field[int64] `json:"expires_after"`
	// The unit of expires_after.
	ExpiresAfterUnit param.Field[CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnit] `json:"expires_after_unit"`
}

func (r CustomerCreditTopUpNewByExternalIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for invoices generated by triggered top-ups.
type CustomerCreditTopUpNewByExternalIDParamsInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection param.Field[bool] `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo param.Field[string] `json:"memo"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment param.Field[bool] `json:"require_successful_payment"`
}

func (r CustomerCreditTopUpNewByExternalIDParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of expires_after.
type CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnit string

const (
	CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnitDay   CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnit = "day"
	CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnitMonth CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnit = "month"
)

func (r CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnit) IsKnown() bool {
	switch r {
	case CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnitDay, CustomerCreditTopUpNewByExternalIDParamsExpiresAfterUnitMonth:
		return true
	}
	return false
}

type CustomerCreditTopUpListByExternalIDParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditTopUpListByExternalIDParams]'s query
// parameters as `url.Values`.
func (r CustomerCreditTopUpListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
