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
)

// A customer is a buyer of your products, and the other party to the billing
// relationship.
//
// In Orb, customers are assigned system generated identifiers automatically, but
// it's often desirable to have these match existing identifiers in your system. To
// avoid having to denormalize Orb ID information, you can pass in an
// `external_customer_id` with your own identifier. See
// [Customer ID Aliases](/events-and-metrics/customer-aliases) for further
// information about how these aliases work in Orb.
//
// In addition to having an identifier in your system, a customer may exist in a
// payment provider solution like Stripe. Use the `payment_provider_id` and the
// `payment_provider` enum field to express this mapping.
//
// A customer also has a timezone (from the standard
// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
// your account's timezone. See [Timezone localization](/essentials/timezones) for
// information on what this timezone parameter influences within Orb.
//
// CustomerPaymentGatewayConnectionService contains methods and other services that
// help with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerPaymentGatewayConnectionService] method instead.
type CustomerPaymentGatewayConnectionService struct {
	Options []option.RequestOption
}

// NewCustomerPaymentGatewayConnectionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomerPaymentGatewayConnectionService(opts ...option.RequestOption) (r *CustomerPaymentGatewayConnectionService) {
	r = &CustomerPaymentGatewayConnectionService{}
	r.Options = opts
	return
}

// This endpoint maps the customer to a payment gateway connection. Orb charges the
// customer for a billing entity's invoices through the customer's first mapping
// for that billing entity; `is_default` moves that to this mapping when the
// customer already has one.
func (r *CustomerPaymentGatewayConnectionService) New(ctx context.Context, customerID string, body CustomerPaymentGatewayConnectionNewParams, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/payment_gateway_connections", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint updates the customer's mapping to a payment gateway connection:
// the customer's ID in that account, and whether Orb charges the customer for the
// billing entity's invoices through it. Omitted fields keep their value.
func (r *CustomerPaymentGatewayConnectionService) Update(ctx context.Context, customerID string, paymentGatewayConnectionID string, body CustomerPaymentGatewayConnectionUpdateParams, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	if paymentGatewayConnectionID == "" {
		err = errors.New("missing required payment_gateway_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/payment_gateway_connections/%s", customerID, paymentGatewayConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// This endpoint returns the customer's mappings to payment gateway connections,
// for example the customer's Stripe ID in each connected Stripe account. A
// customer has at most one mapping per connection, so the list is not paginated.
func (r *CustomerPaymentGatewayConnectionService) List(ctx context.Context, customerID string, query CustomerPaymentGatewayConnectionListParams, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/payment_gateway_connections", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint removes the customer's mapping to a payment gateway connection and
// returns it. Active subscriptions and unpaid invoices do not block it. If the
// mapping was the customer's default for the billing entity, the customer's oldest
// remaining mapping on that entity becomes the default. With no mapping left, the
// next charge goes through the billing entity's default connection.
func (r *CustomerPaymentGatewayConnectionService) Delete(ctx context.Context, customerID string, paymentGatewayConnectionID string, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	if paymentGatewayConnectionID == "" {
		err = errors.New("missing required payment_gateway_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/payment_gateway_connections/%s", customerID, paymentGatewayConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// This endpoint maps the customer to a payment gateway connection. Orb charges the
// customer for a billing entity's invoices through the customer's first mapping
// for that billing entity; `is_default` moves that to this mapping when the
// customer already has one.
func (r *CustomerPaymentGatewayConnectionService) NewByExternalID(ctx context.Context, externalCustomerID string, body CustomerPaymentGatewayConnectionNewByExternalIDParams, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/payment_gateway_connections", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint removes the customer's mapping to a payment gateway connection and
// returns it. Active subscriptions and unpaid invoices do not block it. If the
// mapping was the customer's default for the billing entity, the customer's oldest
// remaining mapping on that entity becomes the default. With no mapping left, the
// next charge goes through the billing entity's default connection.
func (r *CustomerPaymentGatewayConnectionService) DeleteByExternalID(ctx context.Context, externalCustomerID string, paymentGatewayConnectionID string, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return nil, err
	}
	if paymentGatewayConnectionID == "" {
		err = errors.New("missing required payment_gateway_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/payment_gateway_connections/%s", externalCustomerID, paymentGatewayConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// This endpoint returns the customer's mappings to payment gateway connections,
// for example the customer's Stripe ID in each connected Stripe account. A
// customer has at most one mapping per connection, so the list is not paginated.
func (r *CustomerPaymentGatewayConnectionService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerPaymentGatewayConnectionListByExternalIDParams, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnectionListByExternalIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/payment_gateway_connections", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint updates the customer's mapping to a payment gateway connection:
// the customer's ID in that account, and whether Orb charges the customer for the
// billing entity's invoices through it. Omitted fields keep their value.
func (r *CustomerPaymentGatewayConnectionService) UpdateByExternalID(ctx context.Context, externalCustomerID string, paymentGatewayConnectionID string, body CustomerPaymentGatewayConnectionUpdateByExternalIDParams, opts ...option.RequestOption) (res *CustomerPaymentGatewayConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return nil, err
	}
	if paymentGatewayConnectionID == "" {
		err = errors.New("missing required payment_gateway_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/payment_gateway_connections/%s", externalCustomerID, paymentGatewayConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// A customer's mapping to one payment gateway connection.
type CustomerPaymentGatewayConnection struct {
	// The ID of the billing entity that collects payments through the connection.
	BillingEntityID string `json:"billing_entity_id" api:"required"`
	// The time the mapping was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the account in the external payment provider.
	ExternalAccountID string `json:"external_account_id" api:"required,nullable"`
	// Whether Orb charges the customer for that billing entity's invoices through this
	// mapping. On the default billing entity, it is the mapping the customer's
	// `payment_provider_id` reflects.
	IsDefault bool `json:"is_default" api:"required"`
	// The ID of the payment gateway connection the customer is mapped to.
	PaymentGatewayConnectionID string `json:"payment_gateway_connection_id" api:"required"`
	// The ID of the customer in the external payment provider, or null if none is set.
	ProviderCustomerID string `json:"provider_customer_id" api:"required,nullable"`
	// The external payment provider of the connection.
	ProviderType CustomerPaymentGatewayConnectionProviderType `json:"provider_type" api:"required"`
	JSON         customerPaymentGatewayConnectionJSON         `json:"-"`
}

// customerPaymentGatewayConnectionJSON contains the JSON metadata for the struct
// [CustomerPaymentGatewayConnection]
type customerPaymentGatewayConnectionJSON struct {
	BillingEntityID            apijson.Field
	CreatedAt                  apijson.Field
	ExternalAccountID          apijson.Field
	IsDefault                  apijson.Field
	PaymentGatewayConnectionID apijson.Field
	ProviderCustomerID         apijson.Field
	ProviderType               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *CustomerPaymentGatewayConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerPaymentGatewayConnectionJSON) RawJSON() string {
	return r.raw
}

// The external payment provider of the connection.
type CustomerPaymentGatewayConnectionProviderType string

const (
	CustomerPaymentGatewayConnectionProviderTypeStripe CustomerPaymentGatewayConnectionProviderType = "stripe"
	CustomerPaymentGatewayConnectionProviderTypeAdyen  CustomerPaymentGatewayConnectionProviderType = "adyen"
)

func (r CustomerPaymentGatewayConnectionProviderType) IsKnown() bool {
	switch r {
	case CustomerPaymentGatewayConnectionProviderTypeStripe, CustomerPaymentGatewayConnectionProviderTypeAdyen:
		return true
	}
	return false
}

type CustomerPaymentGatewayConnectionListResponse struct {
	CustomerID                string                                           `json:"customer_id" api:"required"`
	PaymentGatewayConnections []CustomerPaymentGatewayConnection               `json:"payment_gateway_connections" api:"required"`
	JSON                      customerPaymentGatewayConnectionListResponseJSON `json:"-"`
}

// customerPaymentGatewayConnectionListResponseJSON contains the JSON metadata for
// the struct [CustomerPaymentGatewayConnectionListResponse]
type customerPaymentGatewayConnectionListResponseJSON struct {
	CustomerID                apijson.Field
	PaymentGatewayConnections apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomerPaymentGatewayConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerPaymentGatewayConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerPaymentGatewayConnectionListByExternalIDResponse struct {
	CustomerID                string                                                       `json:"customer_id" api:"required"`
	PaymentGatewayConnections []CustomerPaymentGatewayConnection                           `json:"payment_gateway_connections" api:"required"`
	JSON                      customerPaymentGatewayConnectionListByExternalIDResponseJSON `json:"-"`
}

// customerPaymentGatewayConnectionListByExternalIDResponseJSON contains the JSON
// metadata for the struct
// [CustomerPaymentGatewayConnectionListByExternalIDResponse]
type customerPaymentGatewayConnectionListByExternalIDResponseJSON struct {
	CustomerID                apijson.Field
	PaymentGatewayConnections apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomerPaymentGatewayConnectionListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerPaymentGatewayConnectionListByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerPaymentGatewayConnectionNewParams struct {
	// The ID of the payment gateway connection to map the customer to.
	PaymentGatewayConnectionID param.Field[string] `json:"payment_gateway_connection_id" api:"required"`
	// Whether Orb charges the customer for the billing entity's invoices through this
	// mapping. The customer's first mapping for a billing entity is the default
	// regardless.
	IsDefault param.Field[bool] `json:"is_default"`
	// The ID of the customer in the external payment provider.
	ProviderCustomerID param.Field[string] `json:"provider_customer_id"`
}

func (r CustomerPaymentGatewayConnectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerPaymentGatewayConnectionUpdateParams struct {
	// Only `true` is accepted: it makes this the customer's default mapping for the
	// connection's billing entity. Omit it to keep the current default. To move the
	// default, set it on another mapping.
	IsDefault param.Field[CustomerPaymentGatewayConnectionUpdateParamsIsDefault] `json:"is_default"`
	// The ID of the customer in the external payment provider. Omit it to keep the
	// current one; `null` clears it, and Orb creates a Stripe customer at the next
	// payment through this connection.
	ProviderCustomerID param.Field[string] `json:"provider_customer_id"`
}

func (r CustomerPaymentGatewayConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only `true` is accepted: it makes this the customer's default mapping for the
// connection's billing entity. Omit it to keep the current default. To move the
// default, set it on another mapping.
type CustomerPaymentGatewayConnectionUpdateParamsIsDefault bool

const (
	CustomerPaymentGatewayConnectionUpdateParamsIsDefaultTrue CustomerPaymentGatewayConnectionUpdateParamsIsDefault = true
)

func (r CustomerPaymentGatewayConnectionUpdateParamsIsDefault) IsKnown() bool {
	switch r {
	case CustomerPaymentGatewayConnectionUpdateParamsIsDefaultTrue:
		return true
	}
	return false
}

type CustomerPaymentGatewayConnectionListParams struct {
	// Only return mappings to this billing entity's connections.
	BillingEntityID param.Field[string] `query:"billing_entity_id"`
	// Only return the mapping to this connection.
	PaymentGatewayConnectionID param.Field[string] `query:"payment_gateway_connection_id"`
}

// URLQuery serializes [CustomerPaymentGatewayConnectionListParams]'s query
// parameters as `url.Values`.
func (r CustomerPaymentGatewayConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerPaymentGatewayConnectionNewByExternalIDParams struct {
	// The ID of the payment gateway connection to map the customer to.
	PaymentGatewayConnectionID param.Field[string] `json:"payment_gateway_connection_id" api:"required"`
	// Whether Orb charges the customer for the billing entity's invoices through this
	// mapping. The customer's first mapping for a billing entity is the default
	// regardless.
	IsDefault param.Field[bool] `json:"is_default"`
	// The ID of the customer in the external payment provider.
	ProviderCustomerID param.Field[string] `json:"provider_customer_id"`
}

func (r CustomerPaymentGatewayConnectionNewByExternalIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerPaymentGatewayConnectionListByExternalIDParams struct {
	// Only return mappings to this billing entity's connections.
	BillingEntityID param.Field[string] `query:"billing_entity_id"`
	// Only return the mapping to this connection.
	PaymentGatewayConnectionID param.Field[string] `query:"payment_gateway_connection_id"`
}

// URLQuery serializes [CustomerPaymentGatewayConnectionListByExternalIDParams]'s
// query parameters as `url.Values`.
func (r CustomerPaymentGatewayConnectionListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerPaymentGatewayConnectionUpdateByExternalIDParams struct {
	// Only `true` is accepted: it makes this the customer's default mapping for the
	// connection's billing entity. Omit it to keep the current default. To move the
	// default, set it on another mapping.
	IsDefault param.Field[CustomerPaymentGatewayConnectionUpdateByExternalIDParamsIsDefault] `json:"is_default"`
	// The ID of the customer in the external payment provider. Omit it to keep the
	// current one; `null` clears it, and Orb creates a Stripe customer at the next
	// payment through this connection.
	ProviderCustomerID param.Field[string] `json:"provider_customer_id"`
}

func (r CustomerPaymentGatewayConnectionUpdateByExternalIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only `true` is accepted: it makes this the customer's default mapping for the
// connection's billing entity. Omit it to keep the current default. To move the
// default, set it on another mapping.
type CustomerPaymentGatewayConnectionUpdateByExternalIDParamsIsDefault bool

const (
	CustomerPaymentGatewayConnectionUpdateByExternalIDParamsIsDefaultTrue CustomerPaymentGatewayConnectionUpdateByExternalIDParamsIsDefault = true
)

func (r CustomerPaymentGatewayConnectionUpdateByExternalIDParamsIsDefault) IsKnown() bool {
	switch r {
	case CustomerPaymentGatewayConnectionUpdateByExternalIDParamsIsDefaultTrue:
		return true
	}
	return false
}
