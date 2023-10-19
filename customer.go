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

// CustomerService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCustomerService] method instead.
type CustomerService struct {
	Options             []option.RequestOption
	Costs               *CustomerCostService
	Usage               *CustomerUsageService
	Credits             *CustomerCreditService
	BalanceTransactions *CustomerBalanceTransactionService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
	r = &CustomerService{}
	r.Options = opts
	r.Costs = NewCustomerCostService(opts...)
	r.Usage = NewCustomerUsageService(opts...)
	r.Credits = NewCustomerCreditService(opts...)
	r.BalanceTransactions = NewCustomerBalanceTransactionService(opts...)
	return
}

// This operation is used to create an Orb customer, who is party to the core
// billing relationship. See [Customer](../guides/concepts#customer) for an
// overview of the customer resource.
//
// This endpoint is critical in the following Orb functionality:
//
//   - Automated charges can be configured by setting `payment_provider` and
//     `payment_provider_id` to automatically issue invoices
//   - [Customer ID Aliases](../guides/events-and-metrics/customer-aliases) can be
//     configured by setting `external_customer_id`
//   - [Timezone localization](../guides/product-catalog/timezones.md) can be
//     configured on a per-customer basis by setting the `timezone` parameter
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `payment_provider`,
// `payment_provider_id`, `name`, `email`, `email_delivery`, `tax_id`,
// `auto_collection`, `metadata`, `shipping_address`, `billing_address`, and
// `additional_emails` of an existing customer. "Other fields on a customer are
// currently immutable.
func (r *CustomerService) Update(ctx context.Context, customerID string, body CustomerUpdateParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all customers for an account. The list of
// customers is ordered starting from the most recently created customer. This
// endpoint follows Orb's
// [standardized pagination format](../reference/pagination).
//
// See [Customer](../guides/concepts#customer) for an overview of the customer
// model.
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *shared.Page[Customer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "customers"
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

// This endpoint returns a list of all customers for an account. The list of
// customers is ordered starting from the most recently created customer. This
// endpoint follows Orb's
// [standardized pagination format](../reference/pagination).
//
// See [Customer](../guides/concepts#customer) for an overview of the customer
// model.
func (r *CustomerService) ListAutoPaging(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) *shared.PageAutoPager[Customer] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This performs a deletion of this customer, its subscriptions, and its invoices.
// This operation is irreversible. Note that this is a _soft_ deletion, but the
// data will be inaccessible through the API and Orb dashboard. For hard-deletion,
// please reach out to the Orb team directly.
//
// **Note**: This operation happens asynchronously and can be expected to take a
// few minutes to propagate to related resources. However, querying for the
// customer on subsequent GET requests while deletion is in process will reflect
// its deletion with a `deleted: true` property. Once the customer deletion has
// been fully processed, the customer will not be returned in the API.
//
// On successful processing, this returns an empty dictionary (`{}`) in the API.
func (r *CustomerService) Delete(ctx context.Context, customerID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This endpoint is used to fetch customer details given an identifier. If the
// `Customer` is in the process of being deleted, only the properties `id` and
// `deleted: true` will be returned.
//
// See the [Customer resource](../guides/core-concepts.mdx#customer) for a full
// discussion of the Customer model.
func (r *CustomerService) Fetch(ctx context.Context, customerID string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint is used to fetch customer details given an `external_customer_id`
// (see [Customer ID Aliases](../guides/events-and-metrics/customer-aliases)).
//
// Note that the resource and semantics of this endpoint exactly mirror
// [Get Customer](fetch-customer).
func (r *CustomerService) FetchByExternalID(ctx context.Context, externalCustomerID string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint is used to update customer details given an `external_customer_id`
// (see [Customer ID Aliases](../guides/events-and-metrics/customer-aliases)). Note
// that the resource and semantics of this endpoint exactly mirror
// [Update Customer](update-customer).
func (r *CustomerService) UpdateByExternalID(ctx context.Context, id string, body CustomerUpdateByExternalIDParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A customer is a buyer of your products, and the other party to the billing
// relationship.
//
// In Orb, customers are assigned system generated identifiers automatically, but
// it's often desirable to have these match existing identifiers in your system. To
// avoid having to denormalize Orb ID information, you can pass in an
// `external_customer_id` with your own identifier. See
// [Customer ID Aliases](../guides/events-and-metrics/customer-aliases) for further
// information about how these aliases work in Orb.
//
// In addition to having an identifier in your system, a customer may exist in a
// payment provider solution like Stripe. Use the `payment_provider_id` and the
// `payment_provider` enum field to express this mapping.
//
// A customer also has a timezone (from the standard
// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
// your account's timezone. See
// [Timezone localization](../guides/product-catalog/timezones.md) for information
// on what this timezone parameter influences within Orb.
type Customer struct {
	ID               string   `json:"id,required"`
	AdditionalEmails []string `json:"additional_emails,required"`
	AutoCollection   bool     `json:"auto_collection,required"`
	// The customer's current balance in their currency.
	Balance        string                 `json:"balance,required"`
	BillingAddress CustomerBillingAddress `json:"billing_address,required,nullable"`
	CreatedAt      time.Time              `json:"created_at,required" format:"date-time"`
	Currency       string                 `json:"currency,required,nullable"`
	// A valid customer email, to be used for notifications. When Orb triggers payment
	// through a payment gateway, this email will be used for any automatically issued
	// receipts.
	Email         string `json:"email,required"`
	EmailDelivery bool   `json:"email_delivery,required"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID string            `json:"external_customer_id,required,nullable"`
	Metadata           map[string]string `json:"metadata,required"`
	// The full name of the customer
	Name string `json:"name,required"`
	// This is used for creating charges or invoices in an external system via Orb.
	// When not in test mode, the connection must first be configured in the Orb
	// webapp.
	PaymentProvider CustomerPaymentProvider `json:"payment_provider,required,nullable"`
	// The ID of this customer in an external payments solution, such as Stripe. This
	// is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID string                  `json:"payment_provider_id,required,nullable"`
	PortalURL         string                  `json:"portal_url,required,nullable"`
	ShippingAddress   CustomerShippingAddress `json:"shipping_address,required,nullable"`
	TaxID             CustomerTaxID           `json:"tax_id,required,nullable"`
	// A timezone identifier from the IANA timezone database, such as
	// "America/Los_Angeles". This "defaults to your account's timezone if not set.
	// This cannot be changed after customer creation.
	Timezone                    string                              `json:"timezone,required"`
	AccountingSyncConfiguration CustomerAccountingSyncConfiguration `json:"accounting_sync_configuration,nullable"`
	ReportingConfiguration      CustomerReportingConfiguration      `json:"reporting_configuration,nullable"`
	JSON                        customerJSON
}

// customerJSON contains the JSON metadata for the struct [Customer]
type customerJSON struct {
	ID                          apijson.Field
	AdditionalEmails            apijson.Field
	AutoCollection              apijson.Field
	Balance                     apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	Email                       apijson.Field
	EmailDelivery               apijson.Field
	ExternalCustomerID          apijson.Field
	Metadata                    apijson.Field
	Name                        apijson.Field
	PaymentProvider             apijson.Field
	PaymentProviderID           apijson.Field
	PortalURL                   apijson.Field
	ShippingAddress             apijson.Field
	TaxID                       apijson.Field
	Timezone                    apijson.Field
	AccountingSyncConfiguration apijson.Field
	ReportingConfiguration      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Customer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBillingAddress struct {
	City       string `json:"city,required,nullable"`
	Country    string `json:"country,required,nullable"`
	Line1      string `json:"line1,required,nullable"`
	Line2      string `json:"line2,required,nullable"`
	PostalCode string `json:"postal_code,required,nullable"`
	State      string `json:"state,required,nullable"`
	JSON       customerBillingAddressJSON
}

// customerBillingAddressJSON contains the JSON metadata for the struct
// [CustomerBillingAddress]
type customerBillingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is used for creating charges or invoices in an external system via Orb.
// When not in test mode, the connection must first be configured in the Orb
// webapp.
type CustomerPaymentProvider string

const (
	CustomerPaymentProviderQuickbooks    CustomerPaymentProvider = "quickbooks"
	CustomerPaymentProviderBillCom       CustomerPaymentProvider = "bill.com"
	CustomerPaymentProviderStripeCharge  CustomerPaymentProvider = "stripe_charge"
	CustomerPaymentProviderStripeInvoice CustomerPaymentProvider = "stripe_invoice"
	CustomerPaymentProviderNetsuite      CustomerPaymentProvider = "netsuite"
)

type CustomerShippingAddress struct {
	City       string `json:"city,required,nullable"`
	Country    string `json:"country,required,nullable"`
	Line1      string `json:"line1,required,nullable"`
	Line2      string `json:"line2,required,nullable"`
	PostalCode string `json:"postal_code,required,nullable"`
	State      string `json:"state,required,nullable"`
	JSON       customerShippingAddressJSON
}

// customerShippingAddressJSON contains the JSON metadata for the struct
// [CustomerShippingAddress]
type customerShippingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerTaxID struct {
	Country string `json:"country,required"`
	Type    string `json:"type,required"`
	Value   string `json:"value,required"`
	JSON    customerTaxIDJSON
}

// customerTaxIDJSON contains the JSON metadata for the struct [CustomerTaxID]
type customerTaxIDJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerTaxID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAccountingSyncConfiguration struct {
	AccountingProviders []CustomerAccountingSyncConfigurationAccountingProvider `json:"accounting_providers,required"`
	Excluded            bool                                                    `json:"excluded,required"`
	JSON                customerAccountingSyncConfigurationJSON
}

// customerAccountingSyncConfigurationJSON contains the JSON metadata for the
// struct [CustomerAccountingSyncConfiguration]
type customerAccountingSyncConfigurationJSON struct {
	AccountingProviders apijson.Field
	Excluded            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerAccountingSyncConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAccountingSyncConfigurationAccountingProvider struct {
	ExternalProviderID string                                                             `json:"external_provider_id,required,nullable"`
	ProviderType       CustomerAccountingSyncConfigurationAccountingProvidersProviderType `json:"provider_type,required"`
	JSON               customerAccountingSyncConfigurationAccountingProviderJSON
}

// customerAccountingSyncConfigurationAccountingProviderJSON contains the JSON
// metadata for the struct [CustomerAccountingSyncConfigurationAccountingProvider]
type customerAccountingSyncConfigurationAccountingProviderJSON struct {
	ExternalProviderID apijson.Field
	ProviderType       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerAccountingSyncConfigurationAccountingProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAccountingSyncConfigurationAccountingProvidersProviderType string

const (
	CustomerAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks CustomerAccountingSyncConfigurationAccountingProvidersProviderType = "quickbooks"
	CustomerAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite   CustomerAccountingSyncConfigurationAccountingProvidersProviderType = "netsuite"
)

type CustomerReportingConfiguration struct {
	Exempt bool `json:"exempt,required"`
	JSON   customerReportingConfigurationJSON
}

// customerReportingConfigurationJSON contains the JSON metadata for the struct
// [CustomerReportingConfiguration]
type customerReportingConfigurationJSON struct {
	Exempt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerReportingConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParams struct {
	// A valid customer email, to be used for notifications. When Orb triggers payment
	// through a payment gateway, this email will be used for any automatically issued
	// receipts.
	Email param.Field[string] `json:"email,required"`
	// The full name of the customer
	Name                        param.Field[string]                                       `json:"name,required"`
	AccountingSyncConfiguration param.Field[CustomerNewParamsAccountingSyncConfiguration] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool]                            `json:"auto_collection"`
	BillingAddress param.Field[CustomerNewParamsBillingAddress] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency      param.Field[string] `json:"currency"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// User-specified key value pairs, often useful for referencing internal resources
	// or IDs. Returned as-is in the customer resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// This is used for creating charges or invoices in an external system via Orb.
	// When not in test mode, the connection must first be configured in the Orb
	// webapp.
	PaymentProvider param.Field[CustomerNewParamsPaymentProvider] `json:"payment_provider"`
	// The ID of this customer in an external payments solution, such as Stripe. This
	// is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID      param.Field[string]                                  `json:"payment_provider_id"`
	ReportingConfiguration param.Field[CustomerNewParamsReportingConfiguration] `json:"reporting_configuration"`
	ShippingAddress        param.Field[CustomerNewParamsShippingAddress]        `json:"shipping_address"`
	TaxID                  param.Field[CustomerNewParamsTaxID]                  `json:"tax_id"`
	// A timezone identifier from the IANA timezone database, such as
	// `"America/Los_Angeles"`. This defaults to your account's timezone if not set.
	// This cannot be changed after customer creation.
	Timezone param.Field[string] `json:"timezone"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsAccountingSyncConfiguration struct {
	AccountingProviders param.Field[[]CustomerNewParamsAccountingSyncConfigurationAccountingProvider] `json:"accounting_providers"`
	Excluded            param.Field[bool]                                                             `json:"excluded"`
}

func (r CustomerNewParamsAccountingSyncConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsAccountingSyncConfigurationAccountingProvider struct {
	ExternalProviderID param.Field[string] `json:"external_provider_id,required"`
	ProviderType       param.Field[string] `json:"provider_type,required"`
}

func (r CustomerNewParamsAccountingSyncConfigurationAccountingProvider) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingAddress struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r CustomerNewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This is used for creating charges or invoices in an external system via Orb.
// When not in test mode, the connection must first be configured in the Orb
// webapp.
type CustomerNewParamsPaymentProvider string

const (
	CustomerNewParamsPaymentProviderQuickbooks    CustomerNewParamsPaymentProvider = "quickbooks"
	CustomerNewParamsPaymentProviderBillCom       CustomerNewParamsPaymentProvider = "bill.com"
	CustomerNewParamsPaymentProviderStripeCharge  CustomerNewParamsPaymentProvider = "stripe_charge"
	CustomerNewParamsPaymentProviderStripeInvoice CustomerNewParamsPaymentProvider = "stripe_invoice"
	CustomerNewParamsPaymentProviderNetsuite      CustomerNewParamsPaymentProvider = "netsuite"
)

type CustomerNewParamsReportingConfiguration struct {
	Exempt param.Field[bool] `json:"exempt,required"`
}

func (r CustomerNewParamsReportingConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsShippingAddress struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r CustomerNewParamsShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsTaxID struct {
	Country param.Field[string] `json:"country,required"`
	Type    param.Field[string] `json:"type,required"`
	Value   param.Field[string] `json:"value,required"`
}

func (r CustomerNewParamsTaxID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParams struct {
	AccountingSyncConfiguration param.Field[CustomerUpdateParamsAccountingSyncConfiguration] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool]                               `json:"auto_collection"`
	BillingAddress param.Field[CustomerUpdateParamsBillingAddress] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency param.Field[string] `json:"currency"`
	// A valid customer email, to be used for invoicing and notifications.
	Email         param.Field[string] `json:"email"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// User-specified key value pairs, often useful for referencing internal resources
	// or IDs. Returned as-is in the customer resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// The full name of the customer
	Name param.Field[string] `json:"name"`
	// This is used for creating charges or invoices in an external system via Orb.
	// When not in test mode:
	//
	//   - the connection must first be configured in the Orb webapp.
	//   - if the provider is an invoicing provider (`stripe_invoice`, `quickbooks`,
	//     `bill.com`, `netsuite`), any product mappings must first be configured with
	//     the Orb team.
	PaymentProvider param.Field[CustomerUpdateParamsPaymentProvider] `json:"payment_provider"`
	// The ID of this customer in an external payments solution, such as Stripe. This
	// is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID      param.Field[string]                                     `json:"payment_provider_id"`
	ReportingConfiguration param.Field[CustomerUpdateParamsReportingConfiguration] `json:"reporting_configuration"`
	ShippingAddress        param.Field[CustomerUpdateParamsShippingAddress]        `json:"shipping_address"`
	TaxID                  param.Field[CustomerUpdateParamsTaxID]                  `json:"tax_id"`
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsAccountingSyncConfiguration struct {
	AccountingProviders param.Field[[]CustomerUpdateParamsAccountingSyncConfigurationAccountingProvider] `json:"accounting_providers"`
	Excluded            param.Field[bool]                                                                `json:"excluded"`
}

func (r CustomerUpdateParamsAccountingSyncConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsAccountingSyncConfigurationAccountingProvider struct {
	ExternalProviderID param.Field[string] `json:"external_provider_id,required"`
	ProviderType       param.Field[string] `json:"provider_type,required"`
}

func (r CustomerUpdateParamsAccountingSyncConfigurationAccountingProvider) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsBillingAddress struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r CustomerUpdateParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This is used for creating charges or invoices in an external system via Orb.
// When not in test mode:
//
//   - the connection must first be configured in the Orb webapp.
//   - if the provider is an invoicing provider (`stripe_invoice`, `quickbooks`,
//     `bill.com`, `netsuite`), any product mappings must first be configured with
//     the Orb team.
type CustomerUpdateParamsPaymentProvider string

const (
	CustomerUpdateParamsPaymentProviderQuickbooks    CustomerUpdateParamsPaymentProvider = "quickbooks"
	CustomerUpdateParamsPaymentProviderBillCom       CustomerUpdateParamsPaymentProvider = "bill.com"
	CustomerUpdateParamsPaymentProviderStripeCharge  CustomerUpdateParamsPaymentProvider = "stripe_charge"
	CustomerUpdateParamsPaymentProviderStripeInvoice CustomerUpdateParamsPaymentProvider = "stripe_invoice"
	CustomerUpdateParamsPaymentProviderNetsuite      CustomerUpdateParamsPaymentProvider = "netsuite"
)

type CustomerUpdateParamsReportingConfiguration struct {
	Exempt param.Field[bool] `json:"exempt,required"`
}

func (r CustomerUpdateParamsReportingConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsShippingAddress struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r CustomerUpdateParamsShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsTaxID struct {
	Country param.Field[string] `json:"country,required"`
	Type    param.Field[string] `json:"type,required"`
	Value   param.Field[string] `json:"value,required"`
}

func (r CustomerUpdateParamsTaxID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerUpdateByExternalIDParams struct {
	AccountingSyncConfiguration param.Field[CustomerUpdateByExternalIDParamsAccountingSyncConfiguration] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool]                                           `json:"auto_collection"`
	BillingAddress param.Field[CustomerUpdateByExternalIDParamsBillingAddress] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency param.Field[string] `json:"currency"`
	// A valid customer email, to be used for invoicing and notifications.
	Email         param.Field[string] `json:"email"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// User-specified key value pairs, often useful for referencing internal resources
	// or IDs. Returned as-is in the customer resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// The full name of the customer
	Name param.Field[string] `json:"name"`
	// This is used for creating charges or invoices in an external system via Orb.
	// When not in test mode:
	//
	//   - the connection must first be configured in the Orb webapp.
	//   - if the provider is an invoicing provider (`stripe_invoice`, `quickbooks`,
	//     `bill.com`, `netsuite`), any product mappings must first be configured with
	//     the Orb team.
	PaymentProvider param.Field[CustomerUpdateByExternalIDParamsPaymentProvider] `json:"payment_provider"`
	// The ID of this customer in an external payments solution, such as Stripe. This
	// is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID      param.Field[string]                                                 `json:"payment_provider_id"`
	ReportingConfiguration param.Field[CustomerUpdateByExternalIDParamsReportingConfiguration] `json:"reporting_configuration"`
	ShippingAddress        param.Field[CustomerUpdateByExternalIDParamsShippingAddress]        `json:"shipping_address"`
	TaxID                  param.Field[CustomerUpdateByExternalIDParamsTaxID]                  `json:"tax_id"`
}

func (r CustomerUpdateByExternalIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateByExternalIDParamsAccountingSyncConfiguration struct {
	AccountingProviders param.Field[[]CustomerUpdateByExternalIDParamsAccountingSyncConfigurationAccountingProvider] `json:"accounting_providers"`
	Excluded            param.Field[bool]                                                                            `json:"excluded"`
}

func (r CustomerUpdateByExternalIDParamsAccountingSyncConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateByExternalIDParamsAccountingSyncConfigurationAccountingProvider struct {
	ExternalProviderID param.Field[string] `json:"external_provider_id,required"`
	ProviderType       param.Field[string] `json:"provider_type,required"`
}

func (r CustomerUpdateByExternalIDParamsAccountingSyncConfigurationAccountingProvider) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateByExternalIDParamsBillingAddress struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r CustomerUpdateByExternalIDParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This is used for creating charges or invoices in an external system via Orb.
// When not in test mode:
//
//   - the connection must first be configured in the Orb webapp.
//   - if the provider is an invoicing provider (`stripe_invoice`, `quickbooks`,
//     `bill.com`, `netsuite`), any product mappings must first be configured with
//     the Orb team.
type CustomerUpdateByExternalIDParamsPaymentProvider string

const (
	CustomerUpdateByExternalIDParamsPaymentProviderQuickbooks    CustomerUpdateByExternalIDParamsPaymentProvider = "quickbooks"
	CustomerUpdateByExternalIDParamsPaymentProviderBillCom       CustomerUpdateByExternalIDParamsPaymentProvider = "bill.com"
	CustomerUpdateByExternalIDParamsPaymentProviderStripeCharge  CustomerUpdateByExternalIDParamsPaymentProvider = "stripe_charge"
	CustomerUpdateByExternalIDParamsPaymentProviderStripeInvoice CustomerUpdateByExternalIDParamsPaymentProvider = "stripe_invoice"
	CustomerUpdateByExternalIDParamsPaymentProviderNetsuite      CustomerUpdateByExternalIDParamsPaymentProvider = "netsuite"
)

type CustomerUpdateByExternalIDParamsReportingConfiguration struct {
	Exempt param.Field[bool] `json:"exempt,required"`
}

func (r CustomerUpdateByExternalIDParamsReportingConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateByExternalIDParamsShippingAddress struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r CustomerUpdateByExternalIDParamsShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateByExternalIDParamsTaxID struct {
	Country param.Field[string] `json:"country,required"`
	Type    param.Field[string] `json:"type,required"`
	Value   param.Field[string] `json:"value,required"`
}

func (r CustomerUpdateByExternalIDParamsTaxID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
