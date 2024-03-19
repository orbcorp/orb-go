// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
// `additional_emails` of an existing customer. Other fields on a customer are
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
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
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
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT number                                                                                     |
	// | France               | `eu_vat`     | European VAT number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
	// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT number                                                                                     |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
	// | Spain                | `eu_vat`     | European VAT number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	TaxID CustomerTaxID `json:"tax_id,required,nullable"`
	// A timezone identifier from the IANA timezone database, such as
	// "America/Los_Angeles". This "defaults to your account's timezone if not set.
	// This cannot be changed after customer creation.
	Timezone                    string                              `json:"timezone,required"`
	AccountingSyncConfiguration CustomerAccountingSyncConfiguration `json:"accounting_sync_configuration,nullable"`
	ReportingConfiguration      CustomerReportingConfiguration      `json:"reporting_configuration,nullable"`
	JSON                        customerJSON                        `json:"-"`
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

func (r customerJSON) RawJSON() string {
	return r.raw
}

type CustomerBillingAddress struct {
	City       string                     `json:"city,required,nullable"`
	Country    string                     `json:"country,required,nullable"`
	Line1      string                     `json:"line1,required,nullable"`
	Line2      string                     `json:"line2,required,nullable"`
	PostalCode string                     `json:"postal_code,required,nullable"`
	State      string                     `json:"state,required,nullable"`
	JSON       customerBillingAddressJSON `json:"-"`
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

func (r customerBillingAddressJSON) RawJSON() string {
	return r.raw
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
	City       string                      `json:"city,required,nullable"`
	Country    string                      `json:"country,required,nullable"`
	Line1      string                      `json:"line1,required,nullable"`
	Line2      string                      `json:"line2,required,nullable"`
	PostalCode string                      `json:"postal_code,required,nullable"`
	State      string                      `json:"state,required,nullable"`
	JSON       customerShippingAddressJSON `json:"-"`
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

func (r customerShippingAddressJSON) RawJSON() string {
	return r.raw
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT number                                                                                     |
// | France               | `eu_vat`     | European VAT number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT number                                                                                     |
// | Greece               | `eu_vat`     | European VAT number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
// | Malta                | `eu_vat `    | European VAT number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
// | Romania              | `eu_vat`     | European VAT number                                                                                     |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
// | Spain                | `eu_vat`     | European VAT number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
type CustomerTaxID struct {
	Country CustomerTaxIDCountry `json:"country,required"`
	Type    CustomerTaxIDType    `json:"type,required"`
	Value   string               `json:"value,required"`
	JSON    customerTaxIDJSON    `json:"-"`
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

func (r customerTaxIDJSON) RawJSON() string {
	return r.raw
}

type CustomerTaxIDCountry string

const (
	CustomerTaxIDCountryAd CustomerTaxIDCountry = "AD"
	CustomerTaxIDCountryAe CustomerTaxIDCountry = "AE"
	CustomerTaxIDCountryAt CustomerTaxIDCountry = "AT"
	CustomerTaxIDCountryAu CustomerTaxIDCountry = "AU"
	CustomerTaxIDCountryBe CustomerTaxIDCountry = "BE"
	CustomerTaxIDCountryBg CustomerTaxIDCountry = "BG"
	CustomerTaxIDCountryBr CustomerTaxIDCountry = "BR"
	CustomerTaxIDCountryCa CustomerTaxIDCountry = "CA"
	CustomerTaxIDCountryCh CustomerTaxIDCountry = "CH"
	CustomerTaxIDCountryCl CustomerTaxIDCountry = "CL"
	CustomerTaxIDCountryCy CustomerTaxIDCountry = "CY"
	CustomerTaxIDCountryCz CustomerTaxIDCountry = "CZ"
	CustomerTaxIDCountryDe CustomerTaxIDCountry = "DE"
	CustomerTaxIDCountryDk CustomerTaxIDCountry = "DK"
	CustomerTaxIDCountryEe CustomerTaxIDCountry = "EE"
	CustomerTaxIDCountryEg CustomerTaxIDCountry = "EG"
	CustomerTaxIDCountryEs CustomerTaxIDCountry = "ES"
	CustomerTaxIDCountryEu CustomerTaxIDCountry = "EU"
	CustomerTaxIDCountryFi CustomerTaxIDCountry = "FI"
	CustomerTaxIDCountryFr CustomerTaxIDCountry = "FR"
	CustomerTaxIDCountryGB CustomerTaxIDCountry = "GB"
	CustomerTaxIDCountryGe CustomerTaxIDCountry = "GE"
	CustomerTaxIDCountryGr CustomerTaxIDCountry = "GR"
	CustomerTaxIDCountryHk CustomerTaxIDCountry = "HK"
	CustomerTaxIDCountryHr CustomerTaxIDCountry = "HR"
	CustomerTaxIDCountryHu CustomerTaxIDCountry = "HU"
	CustomerTaxIDCountryID CustomerTaxIDCountry = "ID"
	CustomerTaxIDCountryIe CustomerTaxIDCountry = "IE"
	CustomerTaxIDCountryIl CustomerTaxIDCountry = "IL"
	CustomerTaxIDCountryIn CustomerTaxIDCountry = "IN"
	CustomerTaxIDCountryIs CustomerTaxIDCountry = "IS"
	CustomerTaxIDCountryIt CustomerTaxIDCountry = "IT"
	CustomerTaxIDCountryJp CustomerTaxIDCountry = "JP"
	CustomerTaxIDCountryKe CustomerTaxIDCountry = "KE"
	CustomerTaxIDCountryKr CustomerTaxIDCountry = "KR"
	CustomerTaxIDCountryLi CustomerTaxIDCountry = "LI"
	CustomerTaxIDCountryLt CustomerTaxIDCountry = "LT"
	CustomerTaxIDCountryLu CustomerTaxIDCountry = "LU"
	CustomerTaxIDCountryLv CustomerTaxIDCountry = "LV"
	CustomerTaxIDCountryMt CustomerTaxIDCountry = "MT"
	CustomerTaxIDCountryMx CustomerTaxIDCountry = "MX"
	CustomerTaxIDCountryMy CustomerTaxIDCountry = "MY"
	CustomerTaxIDCountryNl CustomerTaxIDCountry = "NL"
	CustomerTaxIDCountryNo CustomerTaxIDCountry = "NO"
	CustomerTaxIDCountryNz CustomerTaxIDCountry = "NZ"
	CustomerTaxIDCountryPh CustomerTaxIDCountry = "PH"
	CustomerTaxIDCountryPl CustomerTaxIDCountry = "PL"
	CustomerTaxIDCountryPt CustomerTaxIDCountry = "PT"
	CustomerTaxIDCountryRo CustomerTaxIDCountry = "RO"
	CustomerTaxIDCountryRu CustomerTaxIDCountry = "RU"
	CustomerTaxIDCountrySa CustomerTaxIDCountry = "SA"
	CustomerTaxIDCountrySe CustomerTaxIDCountry = "SE"
	CustomerTaxIDCountrySg CustomerTaxIDCountry = "SG"
	CustomerTaxIDCountrySi CustomerTaxIDCountry = "SI"
	CustomerTaxIDCountrySk CustomerTaxIDCountry = "SK"
	CustomerTaxIDCountryTh CustomerTaxIDCountry = "TH"
	CustomerTaxIDCountryTr CustomerTaxIDCountry = "TR"
	CustomerTaxIDCountryTw CustomerTaxIDCountry = "TW"
	CustomerTaxIDCountryUa CustomerTaxIDCountry = "UA"
	CustomerTaxIDCountryUs CustomerTaxIDCountry = "US"
	CustomerTaxIDCountryZa CustomerTaxIDCountry = "ZA"
)

type CustomerTaxIDType string

const (
	CustomerTaxIDTypeAdNrt    CustomerTaxIDType = "ad_nrt"
	CustomerTaxIDTypeAeTrn    CustomerTaxIDType = "ae_trn"
	CustomerTaxIDTypeEuVat    CustomerTaxIDType = "eu_vat"
	CustomerTaxIDTypeAuAbn    CustomerTaxIDType = "au_abn"
	CustomerTaxIDTypeAuArn    CustomerTaxIDType = "au_arn"
	CustomerTaxIDTypeBgUic    CustomerTaxIDType = "bg_uic"
	CustomerTaxIDTypeBrCnpj   CustomerTaxIDType = "br_cnpj"
	CustomerTaxIDTypeBrCpf    CustomerTaxIDType = "br_cpf"
	CustomerTaxIDTypeCaBn     CustomerTaxIDType = "ca_bn"
	CustomerTaxIDTypeCaGstHst CustomerTaxIDType = "ca_gst_hst"
	CustomerTaxIDTypeCaPstBc  CustomerTaxIDType = "ca_pst_bc"
	CustomerTaxIDTypeCaPstMB  CustomerTaxIDType = "ca_pst_mb"
	CustomerTaxIDTypeCaPstSk  CustomerTaxIDType = "ca_pst_sk"
	CustomerTaxIDTypeCaQst    CustomerTaxIDType = "ca_qst"
	CustomerTaxIDTypeChVat    CustomerTaxIDType = "ch_vat"
	CustomerTaxIDTypeClTin    CustomerTaxIDType = "cl_tin"
	CustomerTaxIDTypeEgTin    CustomerTaxIDType = "eg_tin"
	CustomerTaxIDTypeEsCif    CustomerTaxIDType = "es_cif"
	CustomerTaxIDTypeEuOssVat CustomerTaxIDType = "eu_oss_vat"
	CustomerTaxIDTypeGBVat    CustomerTaxIDType = "gb_vat"
	CustomerTaxIDTypeGeVat    CustomerTaxIDType = "ge_vat"
	CustomerTaxIDTypeHkBr     CustomerTaxIDType = "hk_br"
	CustomerTaxIDTypeHuTin    CustomerTaxIDType = "hu_tin"
	CustomerTaxIDTypeIDNpwp   CustomerTaxIDType = "id_npwp"
	CustomerTaxIDTypeIlVat    CustomerTaxIDType = "il_vat"
	CustomerTaxIDTypeInGst    CustomerTaxIDType = "in_gst"
	CustomerTaxIDTypeIsVat    CustomerTaxIDType = "is_vat"
	CustomerTaxIDTypeJpCn     CustomerTaxIDType = "jp_cn"
	CustomerTaxIDTypeJpRn     CustomerTaxIDType = "jp_rn"
	CustomerTaxIDTypeJpTrn    CustomerTaxIDType = "jp_trn"
	CustomerTaxIDTypeKePin    CustomerTaxIDType = "ke_pin"
	CustomerTaxIDTypeKrBrn    CustomerTaxIDType = "kr_brn"
	CustomerTaxIDTypeLiUid    CustomerTaxIDType = "li_uid"
	CustomerTaxIDTypeMxRfc    CustomerTaxIDType = "mx_rfc"
	CustomerTaxIDTypeMyFrp    CustomerTaxIDType = "my_frp"
	CustomerTaxIDTypeMyItn    CustomerTaxIDType = "my_itn"
	CustomerTaxIDTypeMySst    CustomerTaxIDType = "my_sst"
	CustomerTaxIDTypeNoVat    CustomerTaxIDType = "no_vat"
	CustomerTaxIDTypeNzGst    CustomerTaxIDType = "nz_gst"
	CustomerTaxIDTypePhTin    CustomerTaxIDType = "ph_tin"
	CustomerTaxIDTypeRuInn    CustomerTaxIDType = "ru_inn"
	CustomerTaxIDTypeRuKpp    CustomerTaxIDType = "ru_kpp"
	CustomerTaxIDTypeSaVat    CustomerTaxIDType = "sa_vat"
	CustomerTaxIDTypeSgGst    CustomerTaxIDType = "sg_gst"
	CustomerTaxIDTypeSgUen    CustomerTaxIDType = "sg_uen"
	CustomerTaxIDTypeSiTin    CustomerTaxIDType = "si_tin"
	CustomerTaxIDTypeThVat    CustomerTaxIDType = "th_vat"
	CustomerTaxIDTypeTrTin    CustomerTaxIDType = "tr_tin"
	CustomerTaxIDTypeTwVat    CustomerTaxIDType = "tw_vat"
	CustomerTaxIDTypeUaVat    CustomerTaxIDType = "ua_vat"
	CustomerTaxIDTypeUsEin    CustomerTaxIDType = "us_ein"
	CustomerTaxIDTypeZaVat    CustomerTaxIDType = "za_vat"
)

type CustomerAccountingSyncConfiguration struct {
	AccountingProviders []CustomerAccountingSyncConfigurationAccountingProvider `json:"accounting_providers,required"`
	Excluded            bool                                                    `json:"excluded,required"`
	JSON                customerAccountingSyncConfigurationJSON                 `json:"-"`
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

func (r customerAccountingSyncConfigurationJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountingSyncConfigurationAccountingProvider struct {
	ExternalProviderID string                                                             `json:"external_provider_id,required,nullable"`
	ProviderType       CustomerAccountingSyncConfigurationAccountingProvidersProviderType `json:"provider_type,required"`
	JSON               customerAccountingSyncConfigurationAccountingProviderJSON          `json:"-"`
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

func (r customerAccountingSyncConfigurationAccountingProviderJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountingSyncConfigurationAccountingProvidersProviderType string

const (
	CustomerAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks CustomerAccountingSyncConfigurationAccountingProvidersProviderType = "quickbooks"
	CustomerAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite   CustomerAccountingSyncConfigurationAccountingProvidersProviderType = "netsuite"
)

type CustomerReportingConfiguration struct {
	Exempt bool                               `json:"exempt,required"`
	JSON   customerReportingConfigurationJSON `json:"-"`
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

func (r customerReportingConfigurationJSON) RawJSON() string {
	return r.raw
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
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// This is used for creating charges or invoices in an external system via Orb.
	// When not in test mode, the connection must first be configured in the Orb
	// webapp.
	PaymentProvider param.Field[CustomerNewParamsPaymentProvider] `json:"payment_provider"`
	// The ID of this customer in an external payments solution, such as Stripe. This
	// is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID      param.Field[string]                                  `json:"payment_provider_id"`
	ReportingConfiguration param.Field[CustomerNewParamsReportingConfiguration] `json:"reporting_configuration"`
	ShippingAddress        param.Field[CustomerNewParamsShippingAddress]        `json:"shipping_address"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT number                                                                                     |
	// | France               | `eu_vat`     | European VAT number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
	// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT number                                                                                     |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
	// | Spain                | `eu_vat`     | European VAT number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	TaxID param.Field[CustomerNewParamsTaxID] `json:"tax_id"`
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

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT number                                                                                     |
// | France               | `eu_vat`     | European VAT number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT number                                                                                     |
// | Greece               | `eu_vat`     | European VAT number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
// | Malta                | `eu_vat `    | European VAT number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
// | Romania              | `eu_vat`     | European VAT number                                                                                     |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
// | Spain                | `eu_vat`     | European VAT number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
type CustomerNewParamsTaxID struct {
	Country param.Field[CustomerNewParamsTaxIDCountry] `json:"country,required"`
	Type    param.Field[CustomerNewParamsTaxIDType]    `json:"type,required"`
	Value   param.Field[string]                        `json:"value,required"`
}

func (r CustomerNewParamsTaxID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsTaxIDCountry string

const (
	CustomerNewParamsTaxIDCountryAd CustomerNewParamsTaxIDCountry = "AD"
	CustomerNewParamsTaxIDCountryAe CustomerNewParamsTaxIDCountry = "AE"
	CustomerNewParamsTaxIDCountryAt CustomerNewParamsTaxIDCountry = "AT"
	CustomerNewParamsTaxIDCountryAu CustomerNewParamsTaxIDCountry = "AU"
	CustomerNewParamsTaxIDCountryBe CustomerNewParamsTaxIDCountry = "BE"
	CustomerNewParamsTaxIDCountryBg CustomerNewParamsTaxIDCountry = "BG"
	CustomerNewParamsTaxIDCountryBr CustomerNewParamsTaxIDCountry = "BR"
	CustomerNewParamsTaxIDCountryCa CustomerNewParamsTaxIDCountry = "CA"
	CustomerNewParamsTaxIDCountryCh CustomerNewParamsTaxIDCountry = "CH"
	CustomerNewParamsTaxIDCountryCl CustomerNewParamsTaxIDCountry = "CL"
	CustomerNewParamsTaxIDCountryCy CustomerNewParamsTaxIDCountry = "CY"
	CustomerNewParamsTaxIDCountryCz CustomerNewParamsTaxIDCountry = "CZ"
	CustomerNewParamsTaxIDCountryDe CustomerNewParamsTaxIDCountry = "DE"
	CustomerNewParamsTaxIDCountryDk CustomerNewParamsTaxIDCountry = "DK"
	CustomerNewParamsTaxIDCountryEe CustomerNewParamsTaxIDCountry = "EE"
	CustomerNewParamsTaxIDCountryEg CustomerNewParamsTaxIDCountry = "EG"
	CustomerNewParamsTaxIDCountryEs CustomerNewParamsTaxIDCountry = "ES"
	CustomerNewParamsTaxIDCountryEu CustomerNewParamsTaxIDCountry = "EU"
	CustomerNewParamsTaxIDCountryFi CustomerNewParamsTaxIDCountry = "FI"
	CustomerNewParamsTaxIDCountryFr CustomerNewParamsTaxIDCountry = "FR"
	CustomerNewParamsTaxIDCountryGB CustomerNewParamsTaxIDCountry = "GB"
	CustomerNewParamsTaxIDCountryGe CustomerNewParamsTaxIDCountry = "GE"
	CustomerNewParamsTaxIDCountryGr CustomerNewParamsTaxIDCountry = "GR"
	CustomerNewParamsTaxIDCountryHk CustomerNewParamsTaxIDCountry = "HK"
	CustomerNewParamsTaxIDCountryHr CustomerNewParamsTaxIDCountry = "HR"
	CustomerNewParamsTaxIDCountryHu CustomerNewParamsTaxIDCountry = "HU"
	CustomerNewParamsTaxIDCountryID CustomerNewParamsTaxIDCountry = "ID"
	CustomerNewParamsTaxIDCountryIe CustomerNewParamsTaxIDCountry = "IE"
	CustomerNewParamsTaxIDCountryIl CustomerNewParamsTaxIDCountry = "IL"
	CustomerNewParamsTaxIDCountryIn CustomerNewParamsTaxIDCountry = "IN"
	CustomerNewParamsTaxIDCountryIs CustomerNewParamsTaxIDCountry = "IS"
	CustomerNewParamsTaxIDCountryIt CustomerNewParamsTaxIDCountry = "IT"
	CustomerNewParamsTaxIDCountryJp CustomerNewParamsTaxIDCountry = "JP"
	CustomerNewParamsTaxIDCountryKe CustomerNewParamsTaxIDCountry = "KE"
	CustomerNewParamsTaxIDCountryKr CustomerNewParamsTaxIDCountry = "KR"
	CustomerNewParamsTaxIDCountryLi CustomerNewParamsTaxIDCountry = "LI"
	CustomerNewParamsTaxIDCountryLt CustomerNewParamsTaxIDCountry = "LT"
	CustomerNewParamsTaxIDCountryLu CustomerNewParamsTaxIDCountry = "LU"
	CustomerNewParamsTaxIDCountryLv CustomerNewParamsTaxIDCountry = "LV"
	CustomerNewParamsTaxIDCountryMt CustomerNewParamsTaxIDCountry = "MT"
	CustomerNewParamsTaxIDCountryMx CustomerNewParamsTaxIDCountry = "MX"
	CustomerNewParamsTaxIDCountryMy CustomerNewParamsTaxIDCountry = "MY"
	CustomerNewParamsTaxIDCountryNl CustomerNewParamsTaxIDCountry = "NL"
	CustomerNewParamsTaxIDCountryNo CustomerNewParamsTaxIDCountry = "NO"
	CustomerNewParamsTaxIDCountryNz CustomerNewParamsTaxIDCountry = "NZ"
	CustomerNewParamsTaxIDCountryPh CustomerNewParamsTaxIDCountry = "PH"
	CustomerNewParamsTaxIDCountryPl CustomerNewParamsTaxIDCountry = "PL"
	CustomerNewParamsTaxIDCountryPt CustomerNewParamsTaxIDCountry = "PT"
	CustomerNewParamsTaxIDCountryRo CustomerNewParamsTaxIDCountry = "RO"
	CustomerNewParamsTaxIDCountryRu CustomerNewParamsTaxIDCountry = "RU"
	CustomerNewParamsTaxIDCountrySa CustomerNewParamsTaxIDCountry = "SA"
	CustomerNewParamsTaxIDCountrySe CustomerNewParamsTaxIDCountry = "SE"
	CustomerNewParamsTaxIDCountrySg CustomerNewParamsTaxIDCountry = "SG"
	CustomerNewParamsTaxIDCountrySi CustomerNewParamsTaxIDCountry = "SI"
	CustomerNewParamsTaxIDCountrySk CustomerNewParamsTaxIDCountry = "SK"
	CustomerNewParamsTaxIDCountryTh CustomerNewParamsTaxIDCountry = "TH"
	CustomerNewParamsTaxIDCountryTr CustomerNewParamsTaxIDCountry = "TR"
	CustomerNewParamsTaxIDCountryTw CustomerNewParamsTaxIDCountry = "TW"
	CustomerNewParamsTaxIDCountryUa CustomerNewParamsTaxIDCountry = "UA"
	CustomerNewParamsTaxIDCountryUs CustomerNewParamsTaxIDCountry = "US"
	CustomerNewParamsTaxIDCountryZa CustomerNewParamsTaxIDCountry = "ZA"
)

type CustomerNewParamsTaxIDType string

const (
	CustomerNewParamsTaxIDTypeAdNrt    CustomerNewParamsTaxIDType = "ad_nrt"
	CustomerNewParamsTaxIDTypeAeTrn    CustomerNewParamsTaxIDType = "ae_trn"
	CustomerNewParamsTaxIDTypeEuVat    CustomerNewParamsTaxIDType = "eu_vat"
	CustomerNewParamsTaxIDTypeAuAbn    CustomerNewParamsTaxIDType = "au_abn"
	CustomerNewParamsTaxIDTypeAuArn    CustomerNewParamsTaxIDType = "au_arn"
	CustomerNewParamsTaxIDTypeBgUic    CustomerNewParamsTaxIDType = "bg_uic"
	CustomerNewParamsTaxIDTypeBrCnpj   CustomerNewParamsTaxIDType = "br_cnpj"
	CustomerNewParamsTaxIDTypeBrCpf    CustomerNewParamsTaxIDType = "br_cpf"
	CustomerNewParamsTaxIDTypeCaBn     CustomerNewParamsTaxIDType = "ca_bn"
	CustomerNewParamsTaxIDTypeCaGstHst CustomerNewParamsTaxIDType = "ca_gst_hst"
	CustomerNewParamsTaxIDTypeCaPstBc  CustomerNewParamsTaxIDType = "ca_pst_bc"
	CustomerNewParamsTaxIDTypeCaPstMB  CustomerNewParamsTaxIDType = "ca_pst_mb"
	CustomerNewParamsTaxIDTypeCaPstSk  CustomerNewParamsTaxIDType = "ca_pst_sk"
	CustomerNewParamsTaxIDTypeCaQst    CustomerNewParamsTaxIDType = "ca_qst"
	CustomerNewParamsTaxIDTypeChVat    CustomerNewParamsTaxIDType = "ch_vat"
	CustomerNewParamsTaxIDTypeClTin    CustomerNewParamsTaxIDType = "cl_tin"
	CustomerNewParamsTaxIDTypeEgTin    CustomerNewParamsTaxIDType = "eg_tin"
	CustomerNewParamsTaxIDTypeEsCif    CustomerNewParamsTaxIDType = "es_cif"
	CustomerNewParamsTaxIDTypeEuOssVat CustomerNewParamsTaxIDType = "eu_oss_vat"
	CustomerNewParamsTaxIDTypeGBVat    CustomerNewParamsTaxIDType = "gb_vat"
	CustomerNewParamsTaxIDTypeGeVat    CustomerNewParamsTaxIDType = "ge_vat"
	CustomerNewParamsTaxIDTypeHkBr     CustomerNewParamsTaxIDType = "hk_br"
	CustomerNewParamsTaxIDTypeHuTin    CustomerNewParamsTaxIDType = "hu_tin"
	CustomerNewParamsTaxIDTypeIDNpwp   CustomerNewParamsTaxIDType = "id_npwp"
	CustomerNewParamsTaxIDTypeIlVat    CustomerNewParamsTaxIDType = "il_vat"
	CustomerNewParamsTaxIDTypeInGst    CustomerNewParamsTaxIDType = "in_gst"
	CustomerNewParamsTaxIDTypeIsVat    CustomerNewParamsTaxIDType = "is_vat"
	CustomerNewParamsTaxIDTypeJpCn     CustomerNewParamsTaxIDType = "jp_cn"
	CustomerNewParamsTaxIDTypeJpRn     CustomerNewParamsTaxIDType = "jp_rn"
	CustomerNewParamsTaxIDTypeJpTrn    CustomerNewParamsTaxIDType = "jp_trn"
	CustomerNewParamsTaxIDTypeKePin    CustomerNewParamsTaxIDType = "ke_pin"
	CustomerNewParamsTaxIDTypeKrBrn    CustomerNewParamsTaxIDType = "kr_brn"
	CustomerNewParamsTaxIDTypeLiUid    CustomerNewParamsTaxIDType = "li_uid"
	CustomerNewParamsTaxIDTypeMxRfc    CustomerNewParamsTaxIDType = "mx_rfc"
	CustomerNewParamsTaxIDTypeMyFrp    CustomerNewParamsTaxIDType = "my_frp"
	CustomerNewParamsTaxIDTypeMyItn    CustomerNewParamsTaxIDType = "my_itn"
	CustomerNewParamsTaxIDTypeMySst    CustomerNewParamsTaxIDType = "my_sst"
	CustomerNewParamsTaxIDTypeNoVat    CustomerNewParamsTaxIDType = "no_vat"
	CustomerNewParamsTaxIDTypeNzGst    CustomerNewParamsTaxIDType = "nz_gst"
	CustomerNewParamsTaxIDTypePhTin    CustomerNewParamsTaxIDType = "ph_tin"
	CustomerNewParamsTaxIDTypeRuInn    CustomerNewParamsTaxIDType = "ru_inn"
	CustomerNewParamsTaxIDTypeRuKpp    CustomerNewParamsTaxIDType = "ru_kpp"
	CustomerNewParamsTaxIDTypeSaVat    CustomerNewParamsTaxIDType = "sa_vat"
	CustomerNewParamsTaxIDTypeSgGst    CustomerNewParamsTaxIDType = "sg_gst"
	CustomerNewParamsTaxIDTypeSgUen    CustomerNewParamsTaxIDType = "sg_uen"
	CustomerNewParamsTaxIDTypeSiTin    CustomerNewParamsTaxIDType = "si_tin"
	CustomerNewParamsTaxIDTypeThVat    CustomerNewParamsTaxIDType = "th_vat"
	CustomerNewParamsTaxIDTypeTrTin    CustomerNewParamsTaxIDType = "tr_tin"
	CustomerNewParamsTaxIDTypeTwVat    CustomerNewParamsTaxIDType = "tw_vat"
	CustomerNewParamsTaxIDTypeUaVat    CustomerNewParamsTaxIDType = "ua_vat"
	CustomerNewParamsTaxIDTypeUsEin    CustomerNewParamsTaxIDType = "us_ein"
	CustomerNewParamsTaxIDTypeZaVat    CustomerNewParamsTaxIDType = "za_vat"
)

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
	// The external customer ID. This can only be set if empty and the customer has no
	// past or current subscriptions.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
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
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT number                                                                                     |
	// | France               | `eu_vat`     | European VAT number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
	// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT number                                                                                     |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
	// | Spain                | `eu_vat`     | European VAT number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	TaxID param.Field[CustomerUpdateParamsTaxID] `json:"tax_id"`
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

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT number                                                                                     |
// | France               | `eu_vat`     | European VAT number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT number                                                                                     |
// | Greece               | `eu_vat`     | European VAT number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
// | Malta                | `eu_vat `    | European VAT number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
// | Romania              | `eu_vat`     | European VAT number                                                                                     |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
// | Spain                | `eu_vat`     | European VAT number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
type CustomerUpdateParamsTaxID struct {
	Country param.Field[CustomerUpdateParamsTaxIDCountry] `json:"country,required"`
	Type    param.Field[CustomerUpdateParamsTaxIDType]    `json:"type,required"`
	Value   param.Field[string]                           `json:"value,required"`
}

func (r CustomerUpdateParamsTaxID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsTaxIDCountry string

const (
	CustomerUpdateParamsTaxIDCountryAd CustomerUpdateParamsTaxIDCountry = "AD"
	CustomerUpdateParamsTaxIDCountryAe CustomerUpdateParamsTaxIDCountry = "AE"
	CustomerUpdateParamsTaxIDCountryAt CustomerUpdateParamsTaxIDCountry = "AT"
	CustomerUpdateParamsTaxIDCountryAu CustomerUpdateParamsTaxIDCountry = "AU"
	CustomerUpdateParamsTaxIDCountryBe CustomerUpdateParamsTaxIDCountry = "BE"
	CustomerUpdateParamsTaxIDCountryBg CustomerUpdateParamsTaxIDCountry = "BG"
	CustomerUpdateParamsTaxIDCountryBr CustomerUpdateParamsTaxIDCountry = "BR"
	CustomerUpdateParamsTaxIDCountryCa CustomerUpdateParamsTaxIDCountry = "CA"
	CustomerUpdateParamsTaxIDCountryCh CustomerUpdateParamsTaxIDCountry = "CH"
	CustomerUpdateParamsTaxIDCountryCl CustomerUpdateParamsTaxIDCountry = "CL"
	CustomerUpdateParamsTaxIDCountryCy CustomerUpdateParamsTaxIDCountry = "CY"
	CustomerUpdateParamsTaxIDCountryCz CustomerUpdateParamsTaxIDCountry = "CZ"
	CustomerUpdateParamsTaxIDCountryDe CustomerUpdateParamsTaxIDCountry = "DE"
	CustomerUpdateParamsTaxIDCountryDk CustomerUpdateParamsTaxIDCountry = "DK"
	CustomerUpdateParamsTaxIDCountryEe CustomerUpdateParamsTaxIDCountry = "EE"
	CustomerUpdateParamsTaxIDCountryEg CustomerUpdateParamsTaxIDCountry = "EG"
	CustomerUpdateParamsTaxIDCountryEs CustomerUpdateParamsTaxIDCountry = "ES"
	CustomerUpdateParamsTaxIDCountryEu CustomerUpdateParamsTaxIDCountry = "EU"
	CustomerUpdateParamsTaxIDCountryFi CustomerUpdateParamsTaxIDCountry = "FI"
	CustomerUpdateParamsTaxIDCountryFr CustomerUpdateParamsTaxIDCountry = "FR"
	CustomerUpdateParamsTaxIDCountryGB CustomerUpdateParamsTaxIDCountry = "GB"
	CustomerUpdateParamsTaxIDCountryGe CustomerUpdateParamsTaxIDCountry = "GE"
	CustomerUpdateParamsTaxIDCountryGr CustomerUpdateParamsTaxIDCountry = "GR"
	CustomerUpdateParamsTaxIDCountryHk CustomerUpdateParamsTaxIDCountry = "HK"
	CustomerUpdateParamsTaxIDCountryHr CustomerUpdateParamsTaxIDCountry = "HR"
	CustomerUpdateParamsTaxIDCountryHu CustomerUpdateParamsTaxIDCountry = "HU"
	CustomerUpdateParamsTaxIDCountryID CustomerUpdateParamsTaxIDCountry = "ID"
	CustomerUpdateParamsTaxIDCountryIe CustomerUpdateParamsTaxIDCountry = "IE"
	CustomerUpdateParamsTaxIDCountryIl CustomerUpdateParamsTaxIDCountry = "IL"
	CustomerUpdateParamsTaxIDCountryIn CustomerUpdateParamsTaxIDCountry = "IN"
	CustomerUpdateParamsTaxIDCountryIs CustomerUpdateParamsTaxIDCountry = "IS"
	CustomerUpdateParamsTaxIDCountryIt CustomerUpdateParamsTaxIDCountry = "IT"
	CustomerUpdateParamsTaxIDCountryJp CustomerUpdateParamsTaxIDCountry = "JP"
	CustomerUpdateParamsTaxIDCountryKe CustomerUpdateParamsTaxIDCountry = "KE"
	CustomerUpdateParamsTaxIDCountryKr CustomerUpdateParamsTaxIDCountry = "KR"
	CustomerUpdateParamsTaxIDCountryLi CustomerUpdateParamsTaxIDCountry = "LI"
	CustomerUpdateParamsTaxIDCountryLt CustomerUpdateParamsTaxIDCountry = "LT"
	CustomerUpdateParamsTaxIDCountryLu CustomerUpdateParamsTaxIDCountry = "LU"
	CustomerUpdateParamsTaxIDCountryLv CustomerUpdateParamsTaxIDCountry = "LV"
	CustomerUpdateParamsTaxIDCountryMt CustomerUpdateParamsTaxIDCountry = "MT"
	CustomerUpdateParamsTaxIDCountryMx CustomerUpdateParamsTaxIDCountry = "MX"
	CustomerUpdateParamsTaxIDCountryMy CustomerUpdateParamsTaxIDCountry = "MY"
	CustomerUpdateParamsTaxIDCountryNl CustomerUpdateParamsTaxIDCountry = "NL"
	CustomerUpdateParamsTaxIDCountryNo CustomerUpdateParamsTaxIDCountry = "NO"
	CustomerUpdateParamsTaxIDCountryNz CustomerUpdateParamsTaxIDCountry = "NZ"
	CustomerUpdateParamsTaxIDCountryPh CustomerUpdateParamsTaxIDCountry = "PH"
	CustomerUpdateParamsTaxIDCountryPl CustomerUpdateParamsTaxIDCountry = "PL"
	CustomerUpdateParamsTaxIDCountryPt CustomerUpdateParamsTaxIDCountry = "PT"
	CustomerUpdateParamsTaxIDCountryRo CustomerUpdateParamsTaxIDCountry = "RO"
	CustomerUpdateParamsTaxIDCountryRu CustomerUpdateParamsTaxIDCountry = "RU"
	CustomerUpdateParamsTaxIDCountrySa CustomerUpdateParamsTaxIDCountry = "SA"
	CustomerUpdateParamsTaxIDCountrySe CustomerUpdateParamsTaxIDCountry = "SE"
	CustomerUpdateParamsTaxIDCountrySg CustomerUpdateParamsTaxIDCountry = "SG"
	CustomerUpdateParamsTaxIDCountrySi CustomerUpdateParamsTaxIDCountry = "SI"
	CustomerUpdateParamsTaxIDCountrySk CustomerUpdateParamsTaxIDCountry = "SK"
	CustomerUpdateParamsTaxIDCountryTh CustomerUpdateParamsTaxIDCountry = "TH"
	CustomerUpdateParamsTaxIDCountryTr CustomerUpdateParamsTaxIDCountry = "TR"
	CustomerUpdateParamsTaxIDCountryTw CustomerUpdateParamsTaxIDCountry = "TW"
	CustomerUpdateParamsTaxIDCountryUa CustomerUpdateParamsTaxIDCountry = "UA"
	CustomerUpdateParamsTaxIDCountryUs CustomerUpdateParamsTaxIDCountry = "US"
	CustomerUpdateParamsTaxIDCountryZa CustomerUpdateParamsTaxIDCountry = "ZA"
)

type CustomerUpdateParamsTaxIDType string

const (
	CustomerUpdateParamsTaxIDTypeAdNrt    CustomerUpdateParamsTaxIDType = "ad_nrt"
	CustomerUpdateParamsTaxIDTypeAeTrn    CustomerUpdateParamsTaxIDType = "ae_trn"
	CustomerUpdateParamsTaxIDTypeEuVat    CustomerUpdateParamsTaxIDType = "eu_vat"
	CustomerUpdateParamsTaxIDTypeAuAbn    CustomerUpdateParamsTaxIDType = "au_abn"
	CustomerUpdateParamsTaxIDTypeAuArn    CustomerUpdateParamsTaxIDType = "au_arn"
	CustomerUpdateParamsTaxIDTypeBgUic    CustomerUpdateParamsTaxIDType = "bg_uic"
	CustomerUpdateParamsTaxIDTypeBrCnpj   CustomerUpdateParamsTaxIDType = "br_cnpj"
	CustomerUpdateParamsTaxIDTypeBrCpf    CustomerUpdateParamsTaxIDType = "br_cpf"
	CustomerUpdateParamsTaxIDTypeCaBn     CustomerUpdateParamsTaxIDType = "ca_bn"
	CustomerUpdateParamsTaxIDTypeCaGstHst CustomerUpdateParamsTaxIDType = "ca_gst_hst"
	CustomerUpdateParamsTaxIDTypeCaPstBc  CustomerUpdateParamsTaxIDType = "ca_pst_bc"
	CustomerUpdateParamsTaxIDTypeCaPstMB  CustomerUpdateParamsTaxIDType = "ca_pst_mb"
	CustomerUpdateParamsTaxIDTypeCaPstSk  CustomerUpdateParamsTaxIDType = "ca_pst_sk"
	CustomerUpdateParamsTaxIDTypeCaQst    CustomerUpdateParamsTaxIDType = "ca_qst"
	CustomerUpdateParamsTaxIDTypeChVat    CustomerUpdateParamsTaxIDType = "ch_vat"
	CustomerUpdateParamsTaxIDTypeClTin    CustomerUpdateParamsTaxIDType = "cl_tin"
	CustomerUpdateParamsTaxIDTypeEgTin    CustomerUpdateParamsTaxIDType = "eg_tin"
	CustomerUpdateParamsTaxIDTypeEsCif    CustomerUpdateParamsTaxIDType = "es_cif"
	CustomerUpdateParamsTaxIDTypeEuOssVat CustomerUpdateParamsTaxIDType = "eu_oss_vat"
	CustomerUpdateParamsTaxIDTypeGBVat    CustomerUpdateParamsTaxIDType = "gb_vat"
	CustomerUpdateParamsTaxIDTypeGeVat    CustomerUpdateParamsTaxIDType = "ge_vat"
	CustomerUpdateParamsTaxIDTypeHkBr     CustomerUpdateParamsTaxIDType = "hk_br"
	CustomerUpdateParamsTaxIDTypeHuTin    CustomerUpdateParamsTaxIDType = "hu_tin"
	CustomerUpdateParamsTaxIDTypeIDNpwp   CustomerUpdateParamsTaxIDType = "id_npwp"
	CustomerUpdateParamsTaxIDTypeIlVat    CustomerUpdateParamsTaxIDType = "il_vat"
	CustomerUpdateParamsTaxIDTypeInGst    CustomerUpdateParamsTaxIDType = "in_gst"
	CustomerUpdateParamsTaxIDTypeIsVat    CustomerUpdateParamsTaxIDType = "is_vat"
	CustomerUpdateParamsTaxIDTypeJpCn     CustomerUpdateParamsTaxIDType = "jp_cn"
	CustomerUpdateParamsTaxIDTypeJpRn     CustomerUpdateParamsTaxIDType = "jp_rn"
	CustomerUpdateParamsTaxIDTypeJpTrn    CustomerUpdateParamsTaxIDType = "jp_trn"
	CustomerUpdateParamsTaxIDTypeKePin    CustomerUpdateParamsTaxIDType = "ke_pin"
	CustomerUpdateParamsTaxIDTypeKrBrn    CustomerUpdateParamsTaxIDType = "kr_brn"
	CustomerUpdateParamsTaxIDTypeLiUid    CustomerUpdateParamsTaxIDType = "li_uid"
	CustomerUpdateParamsTaxIDTypeMxRfc    CustomerUpdateParamsTaxIDType = "mx_rfc"
	CustomerUpdateParamsTaxIDTypeMyFrp    CustomerUpdateParamsTaxIDType = "my_frp"
	CustomerUpdateParamsTaxIDTypeMyItn    CustomerUpdateParamsTaxIDType = "my_itn"
	CustomerUpdateParamsTaxIDTypeMySst    CustomerUpdateParamsTaxIDType = "my_sst"
	CustomerUpdateParamsTaxIDTypeNoVat    CustomerUpdateParamsTaxIDType = "no_vat"
	CustomerUpdateParamsTaxIDTypeNzGst    CustomerUpdateParamsTaxIDType = "nz_gst"
	CustomerUpdateParamsTaxIDTypePhTin    CustomerUpdateParamsTaxIDType = "ph_tin"
	CustomerUpdateParamsTaxIDTypeRuInn    CustomerUpdateParamsTaxIDType = "ru_inn"
	CustomerUpdateParamsTaxIDTypeRuKpp    CustomerUpdateParamsTaxIDType = "ru_kpp"
	CustomerUpdateParamsTaxIDTypeSaVat    CustomerUpdateParamsTaxIDType = "sa_vat"
	CustomerUpdateParamsTaxIDTypeSgGst    CustomerUpdateParamsTaxIDType = "sg_gst"
	CustomerUpdateParamsTaxIDTypeSgUen    CustomerUpdateParamsTaxIDType = "sg_uen"
	CustomerUpdateParamsTaxIDTypeSiTin    CustomerUpdateParamsTaxIDType = "si_tin"
	CustomerUpdateParamsTaxIDTypeThVat    CustomerUpdateParamsTaxIDType = "th_vat"
	CustomerUpdateParamsTaxIDTypeTrTin    CustomerUpdateParamsTaxIDType = "tr_tin"
	CustomerUpdateParamsTaxIDTypeTwVat    CustomerUpdateParamsTaxIDType = "tw_vat"
	CustomerUpdateParamsTaxIDTypeUaVat    CustomerUpdateParamsTaxIDType = "ua_vat"
	CustomerUpdateParamsTaxIDTypeUsEin    CustomerUpdateParamsTaxIDType = "us_ein"
	CustomerUpdateParamsTaxIDTypeZaVat    CustomerUpdateParamsTaxIDType = "za_vat"
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
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
	// The external customer ID. This can only be set if empty and the customer has no
	// past or current subscriptions.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
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
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT number                                                                                     |
	// | France               | `eu_vat`     | European VAT number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
	// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT number                                                                                     |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
	// | Spain                | `eu_vat`     | European VAT number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	TaxID param.Field[CustomerUpdateByExternalIDParamsTaxID] `json:"tax_id"`
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

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT number                                                                                     |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT number                                                                                     |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | Croatia              | `eu_vat`     | European VAT number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT number                                                                                     |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | Estonia              | `eu_vat`     | European VAT number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT number                                                                                     |
// | France               | `eu_vat`     | European VAT number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT number                                                                                     |
// | Greece               | `eu_vat`     | European VAT number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary tax number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST number                                                                                    |
// | Malta                | `eu_vat `    | European VAT number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST number                                                                                  |
// | Norway               | `no_vat`     | Norwegian VAT number                                                                                    |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT number                                                                                     |
// | Romania              | `eu_vat`     | European VAT number                                                                                     |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia tax number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF number (previously Spanish CIF number)                                                      |
// | Spain                | `eu_vat`     | European VAT number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
type CustomerUpdateByExternalIDParamsTaxID struct {
	Country param.Field[CustomerUpdateByExternalIDParamsTaxIDCountry] `json:"country,required"`
	Type    param.Field[CustomerUpdateByExternalIDParamsTaxIDType]    `json:"type,required"`
	Value   param.Field[string]                                       `json:"value,required"`
}

func (r CustomerUpdateByExternalIDParamsTaxID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateByExternalIDParamsTaxIDCountry string

const (
	CustomerUpdateByExternalIDParamsTaxIDCountryAd CustomerUpdateByExternalIDParamsTaxIDCountry = "AD"
	CustomerUpdateByExternalIDParamsTaxIDCountryAe CustomerUpdateByExternalIDParamsTaxIDCountry = "AE"
	CustomerUpdateByExternalIDParamsTaxIDCountryAt CustomerUpdateByExternalIDParamsTaxIDCountry = "AT"
	CustomerUpdateByExternalIDParamsTaxIDCountryAu CustomerUpdateByExternalIDParamsTaxIDCountry = "AU"
	CustomerUpdateByExternalIDParamsTaxIDCountryBe CustomerUpdateByExternalIDParamsTaxIDCountry = "BE"
	CustomerUpdateByExternalIDParamsTaxIDCountryBg CustomerUpdateByExternalIDParamsTaxIDCountry = "BG"
	CustomerUpdateByExternalIDParamsTaxIDCountryBr CustomerUpdateByExternalIDParamsTaxIDCountry = "BR"
	CustomerUpdateByExternalIDParamsTaxIDCountryCa CustomerUpdateByExternalIDParamsTaxIDCountry = "CA"
	CustomerUpdateByExternalIDParamsTaxIDCountryCh CustomerUpdateByExternalIDParamsTaxIDCountry = "CH"
	CustomerUpdateByExternalIDParamsTaxIDCountryCl CustomerUpdateByExternalIDParamsTaxIDCountry = "CL"
	CustomerUpdateByExternalIDParamsTaxIDCountryCy CustomerUpdateByExternalIDParamsTaxIDCountry = "CY"
	CustomerUpdateByExternalIDParamsTaxIDCountryCz CustomerUpdateByExternalIDParamsTaxIDCountry = "CZ"
	CustomerUpdateByExternalIDParamsTaxIDCountryDe CustomerUpdateByExternalIDParamsTaxIDCountry = "DE"
	CustomerUpdateByExternalIDParamsTaxIDCountryDk CustomerUpdateByExternalIDParamsTaxIDCountry = "DK"
	CustomerUpdateByExternalIDParamsTaxIDCountryEe CustomerUpdateByExternalIDParamsTaxIDCountry = "EE"
	CustomerUpdateByExternalIDParamsTaxIDCountryEg CustomerUpdateByExternalIDParamsTaxIDCountry = "EG"
	CustomerUpdateByExternalIDParamsTaxIDCountryEs CustomerUpdateByExternalIDParamsTaxIDCountry = "ES"
	CustomerUpdateByExternalIDParamsTaxIDCountryEu CustomerUpdateByExternalIDParamsTaxIDCountry = "EU"
	CustomerUpdateByExternalIDParamsTaxIDCountryFi CustomerUpdateByExternalIDParamsTaxIDCountry = "FI"
	CustomerUpdateByExternalIDParamsTaxIDCountryFr CustomerUpdateByExternalIDParamsTaxIDCountry = "FR"
	CustomerUpdateByExternalIDParamsTaxIDCountryGB CustomerUpdateByExternalIDParamsTaxIDCountry = "GB"
	CustomerUpdateByExternalIDParamsTaxIDCountryGe CustomerUpdateByExternalIDParamsTaxIDCountry = "GE"
	CustomerUpdateByExternalIDParamsTaxIDCountryGr CustomerUpdateByExternalIDParamsTaxIDCountry = "GR"
	CustomerUpdateByExternalIDParamsTaxIDCountryHk CustomerUpdateByExternalIDParamsTaxIDCountry = "HK"
	CustomerUpdateByExternalIDParamsTaxIDCountryHr CustomerUpdateByExternalIDParamsTaxIDCountry = "HR"
	CustomerUpdateByExternalIDParamsTaxIDCountryHu CustomerUpdateByExternalIDParamsTaxIDCountry = "HU"
	CustomerUpdateByExternalIDParamsTaxIDCountryID CustomerUpdateByExternalIDParamsTaxIDCountry = "ID"
	CustomerUpdateByExternalIDParamsTaxIDCountryIe CustomerUpdateByExternalIDParamsTaxIDCountry = "IE"
	CustomerUpdateByExternalIDParamsTaxIDCountryIl CustomerUpdateByExternalIDParamsTaxIDCountry = "IL"
	CustomerUpdateByExternalIDParamsTaxIDCountryIn CustomerUpdateByExternalIDParamsTaxIDCountry = "IN"
	CustomerUpdateByExternalIDParamsTaxIDCountryIs CustomerUpdateByExternalIDParamsTaxIDCountry = "IS"
	CustomerUpdateByExternalIDParamsTaxIDCountryIt CustomerUpdateByExternalIDParamsTaxIDCountry = "IT"
	CustomerUpdateByExternalIDParamsTaxIDCountryJp CustomerUpdateByExternalIDParamsTaxIDCountry = "JP"
	CustomerUpdateByExternalIDParamsTaxIDCountryKe CustomerUpdateByExternalIDParamsTaxIDCountry = "KE"
	CustomerUpdateByExternalIDParamsTaxIDCountryKr CustomerUpdateByExternalIDParamsTaxIDCountry = "KR"
	CustomerUpdateByExternalIDParamsTaxIDCountryLi CustomerUpdateByExternalIDParamsTaxIDCountry = "LI"
	CustomerUpdateByExternalIDParamsTaxIDCountryLt CustomerUpdateByExternalIDParamsTaxIDCountry = "LT"
	CustomerUpdateByExternalIDParamsTaxIDCountryLu CustomerUpdateByExternalIDParamsTaxIDCountry = "LU"
	CustomerUpdateByExternalIDParamsTaxIDCountryLv CustomerUpdateByExternalIDParamsTaxIDCountry = "LV"
	CustomerUpdateByExternalIDParamsTaxIDCountryMt CustomerUpdateByExternalIDParamsTaxIDCountry = "MT"
	CustomerUpdateByExternalIDParamsTaxIDCountryMx CustomerUpdateByExternalIDParamsTaxIDCountry = "MX"
	CustomerUpdateByExternalIDParamsTaxIDCountryMy CustomerUpdateByExternalIDParamsTaxIDCountry = "MY"
	CustomerUpdateByExternalIDParamsTaxIDCountryNl CustomerUpdateByExternalIDParamsTaxIDCountry = "NL"
	CustomerUpdateByExternalIDParamsTaxIDCountryNo CustomerUpdateByExternalIDParamsTaxIDCountry = "NO"
	CustomerUpdateByExternalIDParamsTaxIDCountryNz CustomerUpdateByExternalIDParamsTaxIDCountry = "NZ"
	CustomerUpdateByExternalIDParamsTaxIDCountryPh CustomerUpdateByExternalIDParamsTaxIDCountry = "PH"
	CustomerUpdateByExternalIDParamsTaxIDCountryPl CustomerUpdateByExternalIDParamsTaxIDCountry = "PL"
	CustomerUpdateByExternalIDParamsTaxIDCountryPt CustomerUpdateByExternalIDParamsTaxIDCountry = "PT"
	CustomerUpdateByExternalIDParamsTaxIDCountryRo CustomerUpdateByExternalIDParamsTaxIDCountry = "RO"
	CustomerUpdateByExternalIDParamsTaxIDCountryRu CustomerUpdateByExternalIDParamsTaxIDCountry = "RU"
	CustomerUpdateByExternalIDParamsTaxIDCountrySa CustomerUpdateByExternalIDParamsTaxIDCountry = "SA"
	CustomerUpdateByExternalIDParamsTaxIDCountrySe CustomerUpdateByExternalIDParamsTaxIDCountry = "SE"
	CustomerUpdateByExternalIDParamsTaxIDCountrySg CustomerUpdateByExternalIDParamsTaxIDCountry = "SG"
	CustomerUpdateByExternalIDParamsTaxIDCountrySi CustomerUpdateByExternalIDParamsTaxIDCountry = "SI"
	CustomerUpdateByExternalIDParamsTaxIDCountrySk CustomerUpdateByExternalIDParamsTaxIDCountry = "SK"
	CustomerUpdateByExternalIDParamsTaxIDCountryTh CustomerUpdateByExternalIDParamsTaxIDCountry = "TH"
	CustomerUpdateByExternalIDParamsTaxIDCountryTr CustomerUpdateByExternalIDParamsTaxIDCountry = "TR"
	CustomerUpdateByExternalIDParamsTaxIDCountryTw CustomerUpdateByExternalIDParamsTaxIDCountry = "TW"
	CustomerUpdateByExternalIDParamsTaxIDCountryUa CustomerUpdateByExternalIDParamsTaxIDCountry = "UA"
	CustomerUpdateByExternalIDParamsTaxIDCountryUs CustomerUpdateByExternalIDParamsTaxIDCountry = "US"
	CustomerUpdateByExternalIDParamsTaxIDCountryZa CustomerUpdateByExternalIDParamsTaxIDCountry = "ZA"
)

type CustomerUpdateByExternalIDParamsTaxIDType string

const (
	CustomerUpdateByExternalIDParamsTaxIDTypeAdNrt    CustomerUpdateByExternalIDParamsTaxIDType = "ad_nrt"
	CustomerUpdateByExternalIDParamsTaxIDTypeAeTrn    CustomerUpdateByExternalIDParamsTaxIDType = "ae_trn"
	CustomerUpdateByExternalIDParamsTaxIDTypeEuVat    CustomerUpdateByExternalIDParamsTaxIDType = "eu_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeAuAbn    CustomerUpdateByExternalIDParamsTaxIDType = "au_abn"
	CustomerUpdateByExternalIDParamsTaxIDTypeAuArn    CustomerUpdateByExternalIDParamsTaxIDType = "au_arn"
	CustomerUpdateByExternalIDParamsTaxIDTypeBgUic    CustomerUpdateByExternalIDParamsTaxIDType = "bg_uic"
	CustomerUpdateByExternalIDParamsTaxIDTypeBrCnpj   CustomerUpdateByExternalIDParamsTaxIDType = "br_cnpj"
	CustomerUpdateByExternalIDParamsTaxIDTypeBrCpf    CustomerUpdateByExternalIDParamsTaxIDType = "br_cpf"
	CustomerUpdateByExternalIDParamsTaxIDTypeCaBn     CustomerUpdateByExternalIDParamsTaxIDType = "ca_bn"
	CustomerUpdateByExternalIDParamsTaxIDTypeCaGstHst CustomerUpdateByExternalIDParamsTaxIDType = "ca_gst_hst"
	CustomerUpdateByExternalIDParamsTaxIDTypeCaPstBc  CustomerUpdateByExternalIDParamsTaxIDType = "ca_pst_bc"
	CustomerUpdateByExternalIDParamsTaxIDTypeCaPstMB  CustomerUpdateByExternalIDParamsTaxIDType = "ca_pst_mb"
	CustomerUpdateByExternalIDParamsTaxIDTypeCaPstSk  CustomerUpdateByExternalIDParamsTaxIDType = "ca_pst_sk"
	CustomerUpdateByExternalIDParamsTaxIDTypeCaQst    CustomerUpdateByExternalIDParamsTaxIDType = "ca_qst"
	CustomerUpdateByExternalIDParamsTaxIDTypeChVat    CustomerUpdateByExternalIDParamsTaxIDType = "ch_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeClTin    CustomerUpdateByExternalIDParamsTaxIDType = "cl_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeEgTin    CustomerUpdateByExternalIDParamsTaxIDType = "eg_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeEsCif    CustomerUpdateByExternalIDParamsTaxIDType = "es_cif"
	CustomerUpdateByExternalIDParamsTaxIDTypeEuOssVat CustomerUpdateByExternalIDParamsTaxIDType = "eu_oss_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeGBVat    CustomerUpdateByExternalIDParamsTaxIDType = "gb_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeGeVat    CustomerUpdateByExternalIDParamsTaxIDType = "ge_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeHkBr     CustomerUpdateByExternalIDParamsTaxIDType = "hk_br"
	CustomerUpdateByExternalIDParamsTaxIDTypeHuTin    CustomerUpdateByExternalIDParamsTaxIDType = "hu_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeIDNpwp   CustomerUpdateByExternalIDParamsTaxIDType = "id_npwp"
	CustomerUpdateByExternalIDParamsTaxIDTypeIlVat    CustomerUpdateByExternalIDParamsTaxIDType = "il_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeInGst    CustomerUpdateByExternalIDParamsTaxIDType = "in_gst"
	CustomerUpdateByExternalIDParamsTaxIDTypeIsVat    CustomerUpdateByExternalIDParamsTaxIDType = "is_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeJpCn     CustomerUpdateByExternalIDParamsTaxIDType = "jp_cn"
	CustomerUpdateByExternalIDParamsTaxIDTypeJpRn     CustomerUpdateByExternalIDParamsTaxIDType = "jp_rn"
	CustomerUpdateByExternalIDParamsTaxIDTypeJpTrn    CustomerUpdateByExternalIDParamsTaxIDType = "jp_trn"
	CustomerUpdateByExternalIDParamsTaxIDTypeKePin    CustomerUpdateByExternalIDParamsTaxIDType = "ke_pin"
	CustomerUpdateByExternalIDParamsTaxIDTypeKrBrn    CustomerUpdateByExternalIDParamsTaxIDType = "kr_brn"
	CustomerUpdateByExternalIDParamsTaxIDTypeLiUid    CustomerUpdateByExternalIDParamsTaxIDType = "li_uid"
	CustomerUpdateByExternalIDParamsTaxIDTypeMxRfc    CustomerUpdateByExternalIDParamsTaxIDType = "mx_rfc"
	CustomerUpdateByExternalIDParamsTaxIDTypeMyFrp    CustomerUpdateByExternalIDParamsTaxIDType = "my_frp"
	CustomerUpdateByExternalIDParamsTaxIDTypeMyItn    CustomerUpdateByExternalIDParamsTaxIDType = "my_itn"
	CustomerUpdateByExternalIDParamsTaxIDTypeMySst    CustomerUpdateByExternalIDParamsTaxIDType = "my_sst"
	CustomerUpdateByExternalIDParamsTaxIDTypeNoVat    CustomerUpdateByExternalIDParamsTaxIDType = "no_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeNzGst    CustomerUpdateByExternalIDParamsTaxIDType = "nz_gst"
	CustomerUpdateByExternalIDParamsTaxIDTypePhTin    CustomerUpdateByExternalIDParamsTaxIDType = "ph_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeRuInn    CustomerUpdateByExternalIDParamsTaxIDType = "ru_inn"
	CustomerUpdateByExternalIDParamsTaxIDTypeRuKpp    CustomerUpdateByExternalIDParamsTaxIDType = "ru_kpp"
	CustomerUpdateByExternalIDParamsTaxIDTypeSaVat    CustomerUpdateByExternalIDParamsTaxIDType = "sa_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeSgGst    CustomerUpdateByExternalIDParamsTaxIDType = "sg_gst"
	CustomerUpdateByExternalIDParamsTaxIDTypeSgUen    CustomerUpdateByExternalIDParamsTaxIDType = "sg_uen"
	CustomerUpdateByExternalIDParamsTaxIDTypeSiTin    CustomerUpdateByExternalIDParamsTaxIDType = "si_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeThVat    CustomerUpdateByExternalIDParamsTaxIDType = "th_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeTrTin    CustomerUpdateByExternalIDParamsTaxIDType = "tr_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeTwVat    CustomerUpdateByExternalIDParamsTaxIDType = "tw_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeUaVat    CustomerUpdateByExternalIDParamsTaxIDType = "ua_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeUsEin    CustomerUpdateByExternalIDParamsTaxIDType = "us_ein"
	CustomerUpdateByExternalIDParamsTaxIDTypeZaVat    CustomerUpdateByExternalIDParamsTaxIDType = "za_vat"
)
