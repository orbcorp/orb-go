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
	"github.com/orbcorp/orb-go/shared"
)

// A billing entity is a legal entity an account invoices from, with its own legal
// name, address and tax registrations.
//
// AccountBillingEntityService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBillingEntityService] method instead.
type AccountBillingEntityService struct {
	Options []option.RequestOption
}

// NewAccountBillingEntityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBillingEntityService(opts ...option.RequestOption) (r *AccountBillingEntityService) {
	r = &AccountBillingEntityService{}
	r.Options = opts
	return
}

// This endpoint creates a billing entity: a legal entity the account invoices
// from, with its own legal name, address, tax registrations, and invoice settings.
// `external_billing_entity_id` cannot be changed after creation. This endpoint
// requires an admin API key.
func (r *AccountBillingEntityService) New(ctx context.Context, body AccountBillingEntityNewParams, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "account/billing_entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint updates a billing entity, including the default one. Omitted
// fields keep their value. A change to the name, legal name, address, or tax
// registrations writes a new version of the billing entity; invoices already
// issued keep the version they were issued with. In `invoice_settings`, null
// clears a field so it inherits the account's setting; the default billing entity
// accepts only `memo` and `net_terms`. An archived billing entity cannot be
// updated until it is unarchived. This endpoint requires an admin API key.
//
// `payment_gateway_connections` is the full set of connections the billing entity
// collects payments through; a connection left out is disconnected from the
// billing entity. If a payment through a connection left out is still in progress,
// the whole update is rejected and can be retried once the payment finishes.
func (r *AccountBillingEntityService) Update(ctx context.Context, billingEntityID string, body AccountBillingEntityUpdateParams, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if billingEntityID == "" {
		err = errors.New("missing required billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/%s", billingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// This endpoint returns the account's billing entities, including the default
// billing entity (`is_default: true`), ordered in descending order by creation
// time.
func (r *AccountBillingEntityService) List(ctx context.Context, query AccountBillingEntityListParams, opts ...option.RequestOption) (res *pagination.Page[BillingEntity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account/billing_entities"
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

// This endpoint returns the account's billing entities, including the default
// billing entity (`is_default: true`), ordered in descending order by creation
// time.
func (r *AccountBillingEntityService) ListAutoPaging(ctx context.Context, query AccountBillingEntityListParams, opts ...option.RequestOption) *pagination.PageAutoPager[BillingEntity] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint archives a billing entity. An archived billing entity cannot be
// updated; invoices already issued keep it. The default billing entity cannot be
// archived. Archiving a billing entity that is already archived returns it
// unchanged. This endpoint requires an admin API key.
func (r *AccountBillingEntityService) Archive(ctx context.Context, billingEntityID string, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if billingEntityID == "" {
		err = errors.New("missing required billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/%s/archive", billingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This endpoint archives a billing entity. An archived billing entity cannot be
// updated; invoices already issued keep it. The default billing entity cannot be
// archived. Archiving a billing entity that is already archived returns it
// unchanged. This endpoint requires an admin API key.
func (r *AccountBillingEntityService) ArchiveByExternalID(ctx context.Context, externalBillingEntityID string, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalBillingEntityID == "" {
		err = errors.New("missing required external_billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/external_billing_entity_id/%s/archive", externalBillingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This endpoint returns a billing entity identified by its ID, or by the
// `external_billing_entity_id` set when it was created.
func (r *AccountBillingEntityService) Fetch(ctx context.Context, billingEntityID string, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if billingEntityID == "" {
		err = errors.New("missing required billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/%s", billingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint returns a billing entity identified by its ID, or by the
// `external_billing_entity_id` set when it was created.
func (r *AccountBillingEntityService) FetchByExternalID(ctx context.Context, externalBillingEntityID string, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalBillingEntityID == "" {
		err = errors.New("missing required external_billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/external_billing_entity_id/%s", externalBillingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint restores an archived billing entity to active. Unarchiving an
// active billing entity returns it unchanged. This endpoint requires an admin API
// key.
func (r *AccountBillingEntityService) Unarchive(ctx context.Context, billingEntityID string, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if billingEntityID == "" {
		err = errors.New("missing required billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/%s/unarchive", billingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This endpoint restores an archived billing entity to active. Unarchiving an
// active billing entity returns it unchanged. This endpoint requires an admin API
// key.
func (r *AccountBillingEntityService) UnarchiveByExternalID(ctx context.Context, externalBillingEntityID string, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalBillingEntityID == "" {
		err = errors.New("missing required external_billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/external_billing_entity_id/%s/unarchive", externalBillingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This endpoint updates a billing entity, including the default one. Omitted
// fields keep their value. A change to the name, legal name, address, or tax
// registrations writes a new version of the billing entity; invoices already
// issued keep the version they were issued with. In `invoice_settings`, null
// clears a field so it inherits the account's setting; the default billing entity
// accepts only `memo` and `net_terms`. An archived billing entity cannot be
// updated until it is unarchived. This endpoint requires an admin API key.
//
// `payment_gateway_connections` is the full set of connections the billing entity
// collects payments through; a connection left out is disconnected from the
// billing entity. If a payment through a connection left out is still in progress,
// the whole update is rejected and can be retried once the payment finishes.
func (r *AccountBillingEntityService) UpdateByExternalID(ctx context.Context, externalBillingEntityID string, body AccountBillingEntityUpdateByExternalIDParams, opts ...option.RequestOption) (res *BillingEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalBillingEntityID == "" {
		err = errors.New("missing required external_billing_entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/billing_entities/external_billing_entity_id/%s", externalBillingEntityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// A billing entity is a legal entity an account invoices from, with its own legal
// name, address and tax registrations.
type BillingEntity struct {
	// The Orb-assigned unique identifier for the billing entity.
	ID string `json:"id" api:"required"`
	// The address printed on invoices.
	Address shared.Address `json:"address" api:"required,nullable"`
	// The time the billing entity was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the connection Orb charges when the customer is not mapped to any of
	// the billing entity's connections.
	DefaultPaymentGatewayConnectionID string `json:"default_payment_gateway_connection_id" api:"required,nullable"`
	// An optional user-defined ID for this billing entity, unique within the account.
	// This cannot be changed after creation.
	ExternalBillingEntityID string `json:"external_billing_entity_id" api:"required,nullable"`
	// The billing entity's own invoice settings. Each null field inherits the
	// account's setting.
	InvoiceSettings BillingEntityInvoiceSettings `json:"invoice_settings" api:"required"`
	// Whether this is the account's default billing entity.
	IsDefault bool `json:"is_default" api:"required"`
	// The legal name printed on invoices.
	LegalName string `json:"legal_name" api:"required"`
	// An internal name for the billing entity.
	Name string `json:"name" api:"required"`
	// The payment gateway connections the billing entity collects payments through.
	PaymentGatewayConnections []BillingEntityPaymentGatewayConnection `json:"payment_gateway_connections" api:"required"`
	Status                    BillingEntityStatus                     `json:"status" api:"required"`
	TaxRegistrations          []BillingEntityTaxRegistration          `json:"tax_registrations" api:"required"`
	JSON                      billingEntityJSON                       `json:"-"`
}

// billingEntityJSON contains the JSON metadata for the struct [BillingEntity]
type billingEntityJSON struct {
	ID                                apijson.Field
	Address                           apijson.Field
	CreatedAt                         apijson.Field
	DefaultPaymentGatewayConnectionID apijson.Field
	ExternalBillingEntityID           apijson.Field
	InvoiceSettings                   apijson.Field
	IsDefault                         apijson.Field
	LegalName                         apijson.Field
	Name                              apijson.Field
	PaymentGatewayConnections         apijson.Field
	Status                            apijson.Field
	TaxRegistrations                  apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *BillingEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingEntityJSON) RawJSON() string {
	return r.raw
}

// The billing entity's own invoice settings. Each null field inherits the
// account's setting.
type BillingEntityInvoiceSettings struct {
	// Whether Orb charges the customer's payment method automatically.
	AutoCollection bool `json:"auto_collection" api:"required,nullable"`
	// Whether Orb calculates tax through the tax provider.
	AutomaticTaxEnabled bool `json:"automatic_tax_enabled" api:"required,nullable"`
	// Whether Orb emails invoices to customers.
	EmailDelivery bool `json:"email_delivery" api:"required,nullable"`
	// Whether invoices stay in draft until issued by hand.
	ManualIssuance bool `json:"manual_issuance" api:"required,nullable"`
	// The memo printed on the entity's invoices.
	Memo string `json:"memo" api:"required,nullable"`
	// The number of days until payment is due, counting the start date as day 1, so
	// net 30 from October 1 is due October 30. The account's
	// [due date calculation method](/invoicing/payments#configuring-net-terms-and-collection-behavior)
	// sets whether the start date is the invoice date or the issuance date.
	NetTerms int64                            `json:"net_terms" api:"required,nullable"`
	JSON     billingEntityInvoiceSettingsJSON `json:"-"`
}

// billingEntityInvoiceSettingsJSON contains the JSON metadata for the struct
// [BillingEntityInvoiceSettings]
type billingEntityInvoiceSettingsJSON struct {
	AutoCollection      apijson.Field
	AutomaticTaxEnabled apijson.Field
	EmailDelivery       apijson.Field
	ManualIssuance      apijson.Field
	Memo                apijson.Field
	NetTerms            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BillingEntityInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingEntityInvoiceSettingsJSON) RawJSON() string {
	return r.raw
}

type BillingEntityPaymentGatewayConnection struct {
	// Whether this is the billing entity's default payment gateway connection.
	IsDefault bool `json:"is_default" api:"required"`
	// The ID of the billing entity's payment gateway connection.
	PaymentGatewayConnectionID string `json:"payment_gateway_connection_id" api:"required"`
	// The connection's payment provider.
	PaymentProviderType BillingEntityPaymentGatewayConnectionsPaymentProviderType `json:"payment_provider_type" api:"required"`
	JSON                billingEntityPaymentGatewayConnectionJSON                 `json:"-"`
}

// billingEntityPaymentGatewayConnectionJSON contains the JSON metadata for the
// struct [BillingEntityPaymentGatewayConnection]
type billingEntityPaymentGatewayConnectionJSON struct {
	IsDefault                  apijson.Field
	PaymentGatewayConnectionID apijson.Field
	PaymentProviderType        apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *BillingEntityPaymentGatewayConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingEntityPaymentGatewayConnectionJSON) RawJSON() string {
	return r.raw
}

// The connection's payment provider.
type BillingEntityPaymentGatewayConnectionsPaymentProviderType string

const (
	BillingEntityPaymentGatewayConnectionsPaymentProviderTypeStripe BillingEntityPaymentGatewayConnectionsPaymentProviderType = "stripe"
	BillingEntityPaymentGatewayConnectionsPaymentProviderTypeAdyen  BillingEntityPaymentGatewayConnectionsPaymentProviderType = "adyen"
)

func (r BillingEntityPaymentGatewayConnectionsPaymentProviderType) IsKnown() bool {
	switch r {
	case BillingEntityPaymentGatewayConnectionsPaymentProviderTypeStripe, BillingEntityPaymentGatewayConnectionsPaymentProviderTypeAdyen:
		return true
	}
	return false
}

type BillingEntityStatus string

const (
	BillingEntityStatusActive   BillingEntityStatus = "active"
	BillingEntityStatusArchived BillingEntityStatus = "archived"
)

func (r BillingEntityStatus) IsKnown() bool {
	switch r {
	case BillingEntityStatusActive, BillingEntityStatusArchived:
		return true
	}
	return false
}

type BillingEntityTaxRegistration struct {
	// The jurisdiction the registration applies to. If null, it applies to every
	// customer.
	AppliesTo string `json:"applies_to" api:"required,nullable"`
	// The tax registration type, for example `DE VAT`.
	Code string `json:"code" api:"required"`
	// The registration number.
	Value string                           `json:"value" api:"required"`
	JSON  billingEntityTaxRegistrationJSON `json:"-"`
}

// billingEntityTaxRegistrationJSON contains the JSON metadata for the struct
// [BillingEntityTaxRegistration]
type billingEntityTaxRegistrationJSON struct {
	AppliesTo   apijson.Field
	Code        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingEntityTaxRegistration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingEntityTaxRegistrationJSON) RawJSON() string {
	return r.raw
}

type AccountBillingEntityNewParams struct {
	// The address printed on invoices.
	Address param.Field[AddressInputParam] `json:"address" api:"required"`
	// The legal name printed on invoices.
	LegalName param.Field[string] `json:"legal_name" api:"required"`
	// An internal name for the billing entity.
	Name param.Field[string] `json:"name" api:"required"`
	// An optional user-defined ID for this billing entity, unique within the account.
	// This cannot be changed after creation.
	ExternalBillingEntityID param.Field[string] `json:"external_billing_entity_id"`
	// Settings that override the account's for this billing entity.
	InvoiceSettings  param.Field[AccountBillingEntityNewParamsInvoiceSettings]   `json:"invoice_settings"`
	TaxRegistrations param.Field[[]AccountBillingEntityNewParamsTaxRegistration] `json:"tax_registrations"`
}

func (r AccountBillingEntityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings that override the account's for this billing entity.
type AccountBillingEntityNewParamsInvoiceSettings struct {
	// Whether Orb charges the customer's payment method automatically.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Whether Orb calculates tax through the tax provider.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
	// Whether Orb emails invoices to customers.
	EmailDelivery param.Field[bool] `json:"email_delivery"`
	// Whether invoices stay in draft until issued by hand.
	ManualIssuance param.Field[bool] `json:"manual_issuance"`
	// The memo printed on the entity's invoices.
	Memo param.Field[string] `json:"memo"`
	// The number of days until payment is due, counting the start date as day 1, so
	// net 30 from October 1 is due October 30. The account's
	// [due date calculation method](/invoicing/payments#configuring-net-terms-and-collection-behavior)
	// sets whether the start date is the invoice date or the issuance date.
	NetTerms param.Field[int64] `json:"net_terms"`
}

func (r AccountBillingEntityNewParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBillingEntityNewParamsTaxRegistration struct {
	// The tax registration type, for example `DE VAT`.
	Code param.Field[string] `json:"code" api:"required"`
	// The registration number.
	Value param.Field[string] `json:"value" api:"required"`
	// The jurisdiction the registration applies to. If null, it applies to every
	// customer.
	AppliesTo param.Field[string] `json:"applies_to"`
}

func (r AccountBillingEntityNewParamsTaxRegistration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBillingEntityUpdateParams struct {
	// Replaces the whole address. Omit it to keep the current one.
	Address param.Field[AddressInputParam] `json:"address"`
	// The ID of the connection Orb charges when the customer is not mapped to any of
	// the billing entity's connections. Required when `payment_gateway_connections`
	// has more than one entry.
	DefaultPaymentGatewayConnectionID param.Field[string] `json:"default_payment_gateway_connection_id"`
	// The invoice settings to change. Omitted fields keep their value; null clears a
	// field so it inherits the account's setting.
	InvoiceSettings param.Field[AccountBillingEntityUpdateParamsInvoiceSettings] `json:"invoice_settings"`
	// The legal name printed on invoices.
	LegalName param.Field[string] `json:"legal_name"`
	// An internal name for the billing entity.
	Name param.Field[string] `json:"name"`
	// The full set of payment gateway connections the billing entity collects payments
	// through. Omit it to keep the current ones. Naming another billing entity's
	// connection gives this billing entity a connection of its own through the same
	// Stripe account or Adyen merchant account.
	PaymentGatewayConnections param.Field[[]AccountBillingEntityUpdateParamsPaymentGatewayConnection] `json:"payment_gateway_connections"`
	// The full list of registrations. Omit it to keep the current ones.
	TaxRegistrations param.Field[[]AccountBillingEntityUpdateParamsTaxRegistration] `json:"tax_registrations"`
}

func (r AccountBillingEntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice settings to change. Omitted fields keep their value; null clears a
// field so it inherits the account's setting.
type AccountBillingEntityUpdateParamsInvoiceSettings struct {
	// Whether Orb charges the customer's payment method automatically.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Whether Orb calculates tax through the tax provider.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
	// Whether Orb emails invoices to customers.
	EmailDelivery param.Field[bool] `json:"email_delivery"`
	// Whether invoices stay in draft until issued by hand.
	ManualIssuance param.Field[bool] `json:"manual_issuance"`
	// The memo printed on the entity's invoices.
	Memo param.Field[string] `json:"memo"`
	// The number of days until payment is due, counting the start date as day 1, so
	// net 30 from October 1 is due October 30. The account's
	// [due date calculation method](/invoicing/payments#configuring-net-terms-and-collection-behavior)
	// sets whether the start date is the invoice date or the issuance date.
	NetTerms param.Field[int64] `json:"net_terms"`
}

func (r AccountBillingEntityUpdateParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBillingEntityUpdateParamsPaymentGatewayConnection struct {
	// The ID of a payment gateway connection of the account.
	PaymentGatewayConnectionID param.Field[string] `json:"payment_gateway_connection_id" api:"required"`
	// The connection's payment provider.
	PaymentProviderType param.Field[AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderType] `json:"payment_provider_type" api:"required"`
}

func (r AccountBillingEntityUpdateParamsPaymentGatewayConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The connection's payment provider.
type AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderType string

const (
	AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderTypeStripe AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderType = "stripe"
	AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderTypeAdyen  AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderType = "adyen"
)

func (r AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderType) IsKnown() bool {
	switch r {
	case AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderTypeStripe, AccountBillingEntityUpdateParamsPaymentGatewayConnectionsPaymentProviderTypeAdyen:
		return true
	}
	return false
}

type AccountBillingEntityUpdateParamsTaxRegistration struct {
	// The tax registration type, for example `DE VAT`.
	Code param.Field[string] `json:"code" api:"required"`
	// The registration number.
	Value param.Field[string] `json:"value" api:"required"`
	// The jurisdiction the registration applies to. If null, it applies to every
	// customer.
	AppliesTo param.Field[string] `json:"applies_to"`
}

func (r AccountBillingEntityUpdateParamsTaxRegistration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBillingEntityListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// Only return billing entities with this status. Omit to return both.
	Status param.Field[AccountBillingEntityListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountBillingEntityListParams]'s query parameters as
// `url.Values`.
func (r AccountBillingEntityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only return billing entities with this status. Omit to return both.
type AccountBillingEntityListParamsStatus string

const (
	AccountBillingEntityListParamsStatusActive   AccountBillingEntityListParamsStatus = "active"
	AccountBillingEntityListParamsStatusArchived AccountBillingEntityListParamsStatus = "archived"
)

func (r AccountBillingEntityListParamsStatus) IsKnown() bool {
	switch r {
	case AccountBillingEntityListParamsStatusActive, AccountBillingEntityListParamsStatusArchived:
		return true
	}
	return false
}

type AccountBillingEntityUpdateByExternalIDParams struct {
	// Replaces the whole address. Omit it to keep the current one.
	Address param.Field[AddressInputParam] `json:"address"`
	// The ID of the connection Orb charges when the customer is not mapped to any of
	// the billing entity's connections. Required when `payment_gateway_connections`
	// has more than one entry.
	DefaultPaymentGatewayConnectionID param.Field[string] `json:"default_payment_gateway_connection_id"`
	// The invoice settings to change. Omitted fields keep their value; null clears a
	// field so it inherits the account's setting.
	InvoiceSettings param.Field[AccountBillingEntityUpdateByExternalIDParamsInvoiceSettings] `json:"invoice_settings"`
	// The legal name printed on invoices.
	LegalName param.Field[string] `json:"legal_name"`
	// An internal name for the billing entity.
	Name param.Field[string] `json:"name"`
	// The full set of payment gateway connections the billing entity collects payments
	// through. Omit it to keep the current ones. Naming another billing entity's
	// connection gives this billing entity a connection of its own through the same
	// Stripe account or Adyen merchant account.
	PaymentGatewayConnections param.Field[[]AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnection] `json:"payment_gateway_connections"`
	// The full list of registrations. Omit it to keep the current ones.
	TaxRegistrations param.Field[[]AccountBillingEntityUpdateByExternalIDParamsTaxRegistration] `json:"tax_registrations"`
}

func (r AccountBillingEntityUpdateByExternalIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice settings to change. Omitted fields keep their value; null clears a
// field so it inherits the account's setting.
type AccountBillingEntityUpdateByExternalIDParamsInvoiceSettings struct {
	// Whether Orb charges the customer's payment method automatically.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Whether Orb calculates tax through the tax provider.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
	// Whether Orb emails invoices to customers.
	EmailDelivery param.Field[bool] `json:"email_delivery"`
	// Whether invoices stay in draft until issued by hand.
	ManualIssuance param.Field[bool] `json:"manual_issuance"`
	// The memo printed on the entity's invoices.
	Memo param.Field[string] `json:"memo"`
	// The number of days until payment is due, counting the start date as day 1, so
	// net 30 from October 1 is due October 30. The account's
	// [due date calculation method](/invoicing/payments#configuring-net-terms-and-collection-behavior)
	// sets whether the start date is the invoice date or the issuance date.
	NetTerms param.Field[int64] `json:"net_terms"`
}

func (r AccountBillingEntityUpdateByExternalIDParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnection struct {
	// The ID of a payment gateway connection of the account.
	PaymentGatewayConnectionID param.Field[string] `json:"payment_gateway_connection_id" api:"required"`
	// The connection's payment provider.
	PaymentProviderType param.Field[AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderType] `json:"payment_provider_type" api:"required"`
}

func (r AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The connection's payment provider.
type AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderType string

const (
	AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderTypeStripe AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderType = "stripe"
	AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderTypeAdyen  AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderType = "adyen"
)

func (r AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderType) IsKnown() bool {
	switch r {
	case AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderTypeStripe, AccountBillingEntityUpdateByExternalIDParamsPaymentGatewayConnectionsPaymentProviderTypeAdyen:
		return true
	}
	return false
}

type AccountBillingEntityUpdateByExternalIDParamsTaxRegistration struct {
	// The tax registration type, for example `DE VAT`.
	Code param.Field[string] `json:"code" api:"required"`
	// The registration number.
	Value param.Field[string] `json:"value" api:"required"`
	// The jurisdiction the registration applies to. If null, it applies to every
	// customer.
	AppliesTo param.Field[string] `json:"applies_to"`
}

func (r AccountBillingEntityUpdateByExternalIDParamsTaxRegistration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
