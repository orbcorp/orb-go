// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
)

// A payment gateway connection is one Stripe account or one Adyen merchant account
// that a billing entity collects payments through.
//
// AccountPaymentGatewayConnectionService contains methods and other services that
// help with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPaymentGatewayConnectionService] method instead.
type AccountPaymentGatewayConnectionService struct {
	Options []option.RequestOption
}

// NewAccountPaymentGatewayConnectionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPaymentGatewayConnectionService(opts ...option.RequestOption) (r *AccountPaymentGatewayConnectionService) {
	r = &AccountPaymentGatewayConnectionService{}
	r.Options = opts
	return
}

// This endpoint returns the account's payment gateway connections with the billing
// entity each one collects for, ordered in descending order by creation time.
func (r *AccountPaymentGatewayConnectionService) List(ctx context.Context, query AccountPaymentGatewayConnectionListParams, opts ...option.RequestOption) (res *pagination.Page[PaymentGatewayConnection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account/payment_gateway_connections"
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

// This endpoint returns the account's payment gateway connections with the billing
// entity each one collects for, ordered in descending order by creation time.
func (r *AccountPaymentGatewayConnectionService) ListAutoPaging(ctx context.Context, query AccountPaymentGatewayConnectionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PaymentGatewayConnection] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint returns a payment gateway connection identified by its ID.
func (r *AccountPaymentGatewayConnectionService) Fetch(ctx context.Context, paymentGatewayConnectionID string, opts ...option.RequestOption) (res *PaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentGatewayConnectionID == "" {
		err = errors.New("missing required payment_gateway_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/payment_gateway_connections/%s", paymentGatewayConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A payment gateway connection is one Stripe account or one Adyen merchant account
// that a billing entity collects payments through.
type PaymentGatewayConnection struct {
	// The Orb-assigned unique identifier for the payment gateway connection.
	ID string `json:"id" api:"required"`
	// The ID of the billing entity that collects payments through this connection.
	BillingEntityID string `json:"billing_entity_id" api:"required"`
	// The time the connection was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the account in the external payment provider.
	ExternalAccountID string `json:"external_account_id" api:"required,nullable"`
	// Whether this is the billing entity's default payment gateway connection.
	IsDefault bool `json:"is_default" api:"required"`
	// The external payment provider of this connection.
	ProviderType PaymentGatewayConnectionProviderType `json:"provider_type" api:"required"`
	JSON         paymentGatewayConnectionJSON         `json:"-"`
}

// paymentGatewayConnectionJSON contains the JSON metadata for the struct
// [PaymentGatewayConnection]
type paymentGatewayConnectionJSON struct {
	ID                apijson.Field
	BillingEntityID   apijson.Field
	CreatedAt         apijson.Field
	ExternalAccountID apijson.Field
	IsDefault         apijson.Field
	ProviderType      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentGatewayConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGatewayConnectionJSON) RawJSON() string {
	return r.raw
}

// The external payment provider of this connection.
type PaymentGatewayConnectionProviderType string

const (
	PaymentGatewayConnectionProviderTypeStripe PaymentGatewayConnectionProviderType = "stripe"
	PaymentGatewayConnectionProviderTypeAdyen  PaymentGatewayConnectionProviderType = "adyen"
)

func (r PaymentGatewayConnectionProviderType) IsKnown() bool {
	switch r {
	case PaymentGatewayConnectionProviderTypeStripe, PaymentGatewayConnectionProviderTypeAdyen:
		return true
	}
	return false
}

type AccountPaymentGatewayConnectionListParams struct {
	// Only return connections of this billing entity.
	BillingEntityID param.Field[string] `query:"billing_entity_id"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// Only return connections of this payment provider.
	ProviderType param.Field[AccountPaymentGatewayConnectionListParamsProviderType] `query:"provider_type"`
}

// URLQuery serializes [AccountPaymentGatewayConnectionListParams]'s query
// parameters as `url.Values`.
func (r AccountPaymentGatewayConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only return connections of this payment provider.
type AccountPaymentGatewayConnectionListParamsProviderType string

const (
	AccountPaymentGatewayConnectionListParamsProviderTypeStripe AccountPaymentGatewayConnectionListParamsProviderType = "stripe"
	AccountPaymentGatewayConnectionListParamsProviderTypeAdyen  AccountPaymentGatewayConnectionListParamsProviderType = "adyen"
)

func (r AccountPaymentGatewayConnectionListParamsProviderType) IsKnown() bool {
	switch r {
	case AccountPaymentGatewayConnectionListParamsProviderTypeStripe, AccountPaymentGatewayConnectionListParamsProviderTypeAdyen:
		return true
	}
	return false
}
