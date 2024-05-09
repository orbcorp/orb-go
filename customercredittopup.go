// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
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
func (r *CustomerCreditTopUpService) New(ctx context.Context, customerID string, body CustomerCreditTopUpNewParams, opts ...option.RequestOption) (res *CustomerCreditTopUpNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/credits/top_ups", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List top-ups
func (r *CustomerCreditTopUpService) List(ctx context.Context, customerID string, query CustomerCreditTopUpListParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditTopUpListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *CustomerCreditTopUpService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditTopUpListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditTopUpListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Delete top-up
func (r *CustomerCreditTopUpService) Delete(ctx context.Context, customerID string, topUpID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
func (r *CustomerCreditTopUpService) NewByExternalID(ctx context.Context, externalCustomerID string, body CustomerCreditTopUpNewByExternalIDParams, opts ...option.RequestOption) (res *CustomerCreditTopUpNewByExternalIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/top_ups", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete top-up by external ID
func (r *CustomerCreditTopUpService) DeleteByExternalID(ctx context.Context, externalCustomerID string, topUpID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/top_ups/%s", externalCustomerID, topUpID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List top-ups by external ID
func (r *CustomerCreditTopUpService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditTopUpListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditTopUpListByExternalIDResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *CustomerCreditTopUpService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditTopUpListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditTopUpListByExternalIDResponse] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

type CustomerCreditTopUpNewResponse struct {
	ID string `json:"id,required"`
	// The amount to increment when the threshold is reached.
	Amount string `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency string `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings CustomerCreditTopUpNewResponseInvoiceSettings `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis string `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold string `json:"threshold,required"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter int64 `json:"expires_after,nullable"`
	// The unit of expires_after.
	ExpiresAfterUnit CustomerCreditTopUpNewResponseExpiresAfterUnit `json:"expires_after_unit,nullable"`
	JSON             customerCreditTopUpNewResponseJSON             `json:"-"`
}

// customerCreditTopUpNewResponseJSON contains the JSON metadata for the struct
// [CustomerCreditTopUpNewResponse]
type customerCreditTopUpNewResponseJSON struct {
	ID               apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	InvoiceSettings  apijson.Field
	PerUnitCostBasis apijson.Field
	Threshold        apijson.Field
	ExpiresAfter     apijson.Field
	ExpiresAfterUnit apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditTopUpNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpNewResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for invoices generated by triggered top-ups.
type CustomerCreditTopUpNewResponseInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection bool `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms int64 `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo string `json:"memo,nullable"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment bool                                              `json:"require_successful_payment"`
	JSON                     customerCreditTopUpNewResponseInvoiceSettingsJSON `json:"-"`
}

// customerCreditTopUpNewResponseInvoiceSettingsJSON contains the JSON metadata for
// the struct [CustomerCreditTopUpNewResponseInvoiceSettings]
type customerCreditTopUpNewResponseInvoiceSettingsJSON struct {
	AutoCollection           apijson.Field
	NetTerms                 apijson.Field
	Memo                     apijson.Field
	RequireSuccessfulPayment apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CustomerCreditTopUpNewResponseInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpNewResponseInvoiceSettingsJSON) RawJSON() string {
	return r.raw
}

// The unit of expires_after.
type CustomerCreditTopUpNewResponseExpiresAfterUnit string

const (
	CustomerCreditTopUpNewResponseExpiresAfterUnitDay   CustomerCreditTopUpNewResponseExpiresAfterUnit = "day"
	CustomerCreditTopUpNewResponseExpiresAfterUnitMonth CustomerCreditTopUpNewResponseExpiresAfterUnit = "month"
)

func (r CustomerCreditTopUpNewResponseExpiresAfterUnit) IsKnown() bool {
	switch r {
	case CustomerCreditTopUpNewResponseExpiresAfterUnitDay, CustomerCreditTopUpNewResponseExpiresAfterUnitMonth:
		return true
	}
	return false
}

type CustomerCreditTopUpListResponse struct {
	ID string `json:"id,required"`
	// The amount to increment when the threshold is reached.
	Amount string `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency string `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings CustomerCreditTopUpListResponseInvoiceSettings `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis string `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold string `json:"threshold,required"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter int64 `json:"expires_after,nullable"`
	// The unit of expires_after.
	ExpiresAfterUnit CustomerCreditTopUpListResponseExpiresAfterUnit `json:"expires_after_unit,nullable"`
	JSON             customerCreditTopUpListResponseJSON             `json:"-"`
}

// customerCreditTopUpListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditTopUpListResponse]
type customerCreditTopUpListResponseJSON struct {
	ID               apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	InvoiceSettings  apijson.Field
	PerUnitCostBasis apijson.Field
	Threshold        apijson.Field
	ExpiresAfter     apijson.Field
	ExpiresAfterUnit apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditTopUpListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpListResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for invoices generated by triggered top-ups.
type CustomerCreditTopUpListResponseInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection bool `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms int64 `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo string `json:"memo,nullable"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment bool                                               `json:"require_successful_payment"`
	JSON                     customerCreditTopUpListResponseInvoiceSettingsJSON `json:"-"`
}

// customerCreditTopUpListResponseInvoiceSettingsJSON contains the JSON metadata
// for the struct [CustomerCreditTopUpListResponseInvoiceSettings]
type customerCreditTopUpListResponseInvoiceSettingsJSON struct {
	AutoCollection           apijson.Field
	NetTerms                 apijson.Field
	Memo                     apijson.Field
	RequireSuccessfulPayment apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CustomerCreditTopUpListResponseInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpListResponseInvoiceSettingsJSON) RawJSON() string {
	return r.raw
}

// The unit of expires_after.
type CustomerCreditTopUpListResponseExpiresAfterUnit string

const (
	CustomerCreditTopUpListResponseExpiresAfterUnitDay   CustomerCreditTopUpListResponseExpiresAfterUnit = "day"
	CustomerCreditTopUpListResponseExpiresAfterUnitMonth CustomerCreditTopUpListResponseExpiresAfterUnit = "month"
)

func (r CustomerCreditTopUpListResponseExpiresAfterUnit) IsKnown() bool {
	switch r {
	case CustomerCreditTopUpListResponseExpiresAfterUnitDay, CustomerCreditTopUpListResponseExpiresAfterUnitMonth:
		return true
	}
	return false
}

type CustomerCreditTopUpNewByExternalIDResponse struct {
	ID string `json:"id,required"`
	// The amount to increment when the threshold is reached.
	Amount string `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency string `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings CustomerCreditTopUpNewByExternalIDResponseInvoiceSettings `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis string `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold string `json:"threshold,required"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter int64 `json:"expires_after,nullable"`
	// The unit of expires_after.
	ExpiresAfterUnit CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnit `json:"expires_after_unit,nullable"`
	JSON             customerCreditTopUpNewByExternalIDResponseJSON             `json:"-"`
}

// customerCreditTopUpNewByExternalIDResponseJSON contains the JSON metadata for
// the struct [CustomerCreditTopUpNewByExternalIDResponse]
type customerCreditTopUpNewByExternalIDResponseJSON struct {
	ID               apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	InvoiceSettings  apijson.Field
	PerUnitCostBasis apijson.Field
	Threshold        apijson.Field
	ExpiresAfter     apijson.Field
	ExpiresAfterUnit apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditTopUpNewByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpNewByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for invoices generated by triggered top-ups.
type CustomerCreditTopUpNewByExternalIDResponseInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection bool `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms int64 `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo string `json:"memo,nullable"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment bool                                                          `json:"require_successful_payment"`
	JSON                     customerCreditTopUpNewByExternalIDResponseInvoiceSettingsJSON `json:"-"`
}

// customerCreditTopUpNewByExternalIDResponseInvoiceSettingsJSON contains the JSON
// metadata for the struct
// [CustomerCreditTopUpNewByExternalIDResponseInvoiceSettings]
type customerCreditTopUpNewByExternalIDResponseInvoiceSettingsJSON struct {
	AutoCollection           apijson.Field
	NetTerms                 apijson.Field
	Memo                     apijson.Field
	RequireSuccessfulPayment apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CustomerCreditTopUpNewByExternalIDResponseInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpNewByExternalIDResponseInvoiceSettingsJSON) RawJSON() string {
	return r.raw
}

// The unit of expires_after.
type CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnit string

const (
	CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnitDay   CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnit = "day"
	CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnitMonth CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnit = "month"
)

func (r CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnit) IsKnown() bool {
	switch r {
	case CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnitDay, CustomerCreditTopUpNewByExternalIDResponseExpiresAfterUnitMonth:
		return true
	}
	return false
}

type CustomerCreditTopUpListByExternalIDResponse struct {
	ID string `json:"id,required"`
	// The amount to increment when the threshold is reached.
	Amount string `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency string `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings CustomerCreditTopUpListByExternalIDResponseInvoiceSettings `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis string `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold string `json:"threshold,required"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter int64 `json:"expires_after,nullable"`
	// The unit of expires_after.
	ExpiresAfterUnit CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnit `json:"expires_after_unit,nullable"`
	JSON             customerCreditTopUpListByExternalIDResponseJSON             `json:"-"`
}

// customerCreditTopUpListByExternalIDResponseJSON contains the JSON metadata for
// the struct [CustomerCreditTopUpListByExternalIDResponse]
type customerCreditTopUpListByExternalIDResponseJSON struct {
	ID               apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	InvoiceSettings  apijson.Field
	PerUnitCostBasis apijson.Field
	Threshold        apijson.Field
	ExpiresAfter     apijson.Field
	ExpiresAfterUnit apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditTopUpListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpListByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for invoices generated by triggered top-ups.
type CustomerCreditTopUpListByExternalIDResponseInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection bool `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms int64 `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo string `json:"memo,nullable"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment bool                                                           `json:"require_successful_payment"`
	JSON                     customerCreditTopUpListByExternalIDResponseInvoiceSettingsJSON `json:"-"`
}

// customerCreditTopUpListByExternalIDResponseInvoiceSettingsJSON contains the JSON
// metadata for the struct
// [CustomerCreditTopUpListByExternalIDResponseInvoiceSettings]
type customerCreditTopUpListByExternalIDResponseInvoiceSettingsJSON struct {
	AutoCollection           apijson.Field
	NetTerms                 apijson.Field
	Memo                     apijson.Field
	RequireSuccessfulPayment apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CustomerCreditTopUpListByExternalIDResponseInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditTopUpListByExternalIDResponseInvoiceSettingsJSON) RawJSON() string {
	return r.raw
}

// The unit of expires_after.
type CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnit string

const (
	CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnitDay   CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnit = "day"
	CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnitMonth CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnit = "month"
)

func (r CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnit) IsKnown() bool {
	switch r {
	case CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnitDay, CustomerCreditTopUpListByExternalIDResponseExpiresAfterUnitMonth:
		return true
	}
	return false
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
