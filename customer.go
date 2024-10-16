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

// CustomerService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options             []option.RequestOption
	Costs               *CustomerCostService
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
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *pagination.Page[Customer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *CustomerService) ListAutoPaging(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Customer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This performs a deletion of this customer, its subscriptions, and its invoices,
// provided the customer does not have any issued invoices. Customers with issued
// invoices cannot be deleted. This operation is irreversible. Note that this is a
// _soft_ deletion, but the data will be inaccessible through the API and Orb
// dashboard. For a hard-deletion, please reach out to the Orb team directly.
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
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
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
	Email                  string `json:"email,required"`
	EmailDelivery          bool   `json:"email_delivery,required"`
	ExemptFromAutomatedTax bool   `json:"exempt_from_automated_tax,required,nullable"`
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
	// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
	// | France               | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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
	ExemptFromAutomatedTax      apijson.Field
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

func (r CustomerPaymentProvider) IsKnown() bool {
	switch r {
	case CustomerPaymentProviderQuickbooks, CustomerPaymentProviderBillCom, CustomerPaymentProviderStripeCharge, CustomerPaymentProviderStripeInvoice, CustomerPaymentProviderNetsuite:
		return true
	}
	return false
}

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
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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
	CustomerTaxIDCountryAr CustomerTaxIDCountry = "AR"
	CustomerTaxIDCountryAt CustomerTaxIDCountry = "AT"
	CustomerTaxIDCountryAu CustomerTaxIDCountry = "AU"
	CustomerTaxIDCountryBe CustomerTaxIDCountry = "BE"
	CustomerTaxIDCountryBg CustomerTaxIDCountry = "BG"
	CustomerTaxIDCountryBh CustomerTaxIDCountry = "BH"
	CustomerTaxIDCountryBo CustomerTaxIDCountry = "BO"
	CustomerTaxIDCountryBr CustomerTaxIDCountry = "BR"
	CustomerTaxIDCountryCa CustomerTaxIDCountry = "CA"
	CustomerTaxIDCountryCh CustomerTaxIDCountry = "CH"
	CustomerTaxIDCountryCl CustomerTaxIDCountry = "CL"
	CustomerTaxIDCountryCn CustomerTaxIDCountry = "CN"
	CustomerTaxIDCountryCo CustomerTaxIDCountry = "CO"
	CustomerTaxIDCountryCr CustomerTaxIDCountry = "CR"
	CustomerTaxIDCountryCy CustomerTaxIDCountry = "CY"
	CustomerTaxIDCountryCz CustomerTaxIDCountry = "CZ"
	CustomerTaxIDCountryDe CustomerTaxIDCountry = "DE"
	CustomerTaxIDCountryDk CustomerTaxIDCountry = "DK"
	CustomerTaxIDCountryEe CustomerTaxIDCountry = "EE"
	CustomerTaxIDCountryDo CustomerTaxIDCountry = "DO"
	CustomerTaxIDCountryEc CustomerTaxIDCountry = "EC"
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
	CustomerTaxIDCountryKz CustomerTaxIDCountry = "KZ"
	CustomerTaxIDCountryLi CustomerTaxIDCountry = "LI"
	CustomerTaxIDCountryLt CustomerTaxIDCountry = "LT"
	CustomerTaxIDCountryLu CustomerTaxIDCountry = "LU"
	CustomerTaxIDCountryLv CustomerTaxIDCountry = "LV"
	CustomerTaxIDCountryMt CustomerTaxIDCountry = "MT"
	CustomerTaxIDCountryMx CustomerTaxIDCountry = "MX"
	CustomerTaxIDCountryMy CustomerTaxIDCountry = "MY"
	CustomerTaxIDCountryNg CustomerTaxIDCountry = "NG"
	CustomerTaxIDCountryNl CustomerTaxIDCountry = "NL"
	CustomerTaxIDCountryNo CustomerTaxIDCountry = "NO"
	CustomerTaxIDCountryNz CustomerTaxIDCountry = "NZ"
	CustomerTaxIDCountryOm CustomerTaxIDCountry = "OM"
	CustomerTaxIDCountryPe CustomerTaxIDCountry = "PE"
	CustomerTaxIDCountryPh CustomerTaxIDCountry = "PH"
	CustomerTaxIDCountryPl CustomerTaxIDCountry = "PL"
	CustomerTaxIDCountryPt CustomerTaxIDCountry = "PT"
	CustomerTaxIDCountryRo CustomerTaxIDCountry = "RO"
	CustomerTaxIDCountryRs CustomerTaxIDCountry = "RS"
	CustomerTaxIDCountryRu CustomerTaxIDCountry = "RU"
	CustomerTaxIDCountrySa CustomerTaxIDCountry = "SA"
	CustomerTaxIDCountrySe CustomerTaxIDCountry = "SE"
	CustomerTaxIDCountrySg CustomerTaxIDCountry = "SG"
	CustomerTaxIDCountrySi CustomerTaxIDCountry = "SI"
	CustomerTaxIDCountrySk CustomerTaxIDCountry = "SK"
	CustomerTaxIDCountrySv CustomerTaxIDCountry = "SV"
	CustomerTaxIDCountryTh CustomerTaxIDCountry = "TH"
	CustomerTaxIDCountryTr CustomerTaxIDCountry = "TR"
	CustomerTaxIDCountryTw CustomerTaxIDCountry = "TW"
	CustomerTaxIDCountryUa CustomerTaxIDCountry = "UA"
	CustomerTaxIDCountryUs CustomerTaxIDCountry = "US"
	CustomerTaxIDCountryUy CustomerTaxIDCountry = "UY"
	CustomerTaxIDCountryVe CustomerTaxIDCountry = "VE"
	CustomerTaxIDCountryVn CustomerTaxIDCountry = "VN"
	CustomerTaxIDCountryZa CustomerTaxIDCountry = "ZA"
)

func (r CustomerTaxIDCountry) IsKnown() bool {
	switch r {
	case CustomerTaxIDCountryAd, CustomerTaxIDCountryAe, CustomerTaxIDCountryAr, CustomerTaxIDCountryAt, CustomerTaxIDCountryAu, CustomerTaxIDCountryBe, CustomerTaxIDCountryBg, CustomerTaxIDCountryBh, CustomerTaxIDCountryBo, CustomerTaxIDCountryBr, CustomerTaxIDCountryCa, CustomerTaxIDCountryCh, CustomerTaxIDCountryCl, CustomerTaxIDCountryCn, CustomerTaxIDCountryCo, CustomerTaxIDCountryCr, CustomerTaxIDCountryCy, CustomerTaxIDCountryCz, CustomerTaxIDCountryDe, CustomerTaxIDCountryDk, CustomerTaxIDCountryEe, CustomerTaxIDCountryDo, CustomerTaxIDCountryEc, CustomerTaxIDCountryEg, CustomerTaxIDCountryEs, CustomerTaxIDCountryEu, CustomerTaxIDCountryFi, CustomerTaxIDCountryFr, CustomerTaxIDCountryGB, CustomerTaxIDCountryGe, CustomerTaxIDCountryGr, CustomerTaxIDCountryHk, CustomerTaxIDCountryHr, CustomerTaxIDCountryHu, CustomerTaxIDCountryID, CustomerTaxIDCountryIe, CustomerTaxIDCountryIl, CustomerTaxIDCountryIn, CustomerTaxIDCountryIs, CustomerTaxIDCountryIt, CustomerTaxIDCountryJp, CustomerTaxIDCountryKe, CustomerTaxIDCountryKr, CustomerTaxIDCountryKz, CustomerTaxIDCountryLi, CustomerTaxIDCountryLt, CustomerTaxIDCountryLu, CustomerTaxIDCountryLv, CustomerTaxIDCountryMt, CustomerTaxIDCountryMx, CustomerTaxIDCountryMy, CustomerTaxIDCountryNg, CustomerTaxIDCountryNl, CustomerTaxIDCountryNo, CustomerTaxIDCountryNz, CustomerTaxIDCountryOm, CustomerTaxIDCountryPe, CustomerTaxIDCountryPh, CustomerTaxIDCountryPl, CustomerTaxIDCountryPt, CustomerTaxIDCountryRo, CustomerTaxIDCountryRs, CustomerTaxIDCountryRu, CustomerTaxIDCountrySa, CustomerTaxIDCountrySe, CustomerTaxIDCountrySg, CustomerTaxIDCountrySi, CustomerTaxIDCountrySk, CustomerTaxIDCountrySv, CustomerTaxIDCountryTh, CustomerTaxIDCountryTr, CustomerTaxIDCountryTw, CustomerTaxIDCountryUa, CustomerTaxIDCountryUs, CustomerTaxIDCountryUy, CustomerTaxIDCountryVe, CustomerTaxIDCountryVn, CustomerTaxIDCountryZa:
		return true
	}
	return false
}

type CustomerTaxIDType string

const (
	CustomerTaxIDTypeAdNrt    CustomerTaxIDType = "ad_nrt"
	CustomerTaxIDTypeAeTrn    CustomerTaxIDType = "ae_trn"
	CustomerTaxIDTypeArCuit   CustomerTaxIDType = "ar_cuit"
	CustomerTaxIDTypeEuVat    CustomerTaxIDType = "eu_vat"
	CustomerTaxIDTypeAuAbn    CustomerTaxIDType = "au_abn"
	CustomerTaxIDTypeAuArn    CustomerTaxIDType = "au_arn"
	CustomerTaxIDTypeBgUic    CustomerTaxIDType = "bg_uic"
	CustomerTaxIDTypeBhVat    CustomerTaxIDType = "bh_vat"
	CustomerTaxIDTypeBoTin    CustomerTaxIDType = "bo_tin"
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
	CustomerTaxIDTypeCnTin    CustomerTaxIDType = "cn_tin"
	CustomerTaxIDTypeCoNit    CustomerTaxIDType = "co_nit"
	CustomerTaxIDTypeCrTin    CustomerTaxIDType = "cr_tin"
	CustomerTaxIDTypeDoRcn    CustomerTaxIDType = "do_rcn"
	CustomerTaxIDTypeEcRuc    CustomerTaxIDType = "ec_ruc"
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
	CustomerTaxIDTypeKzBin    CustomerTaxIDType = "kz_bin"
	CustomerTaxIDTypeLiUid    CustomerTaxIDType = "li_uid"
	CustomerTaxIDTypeMxRfc    CustomerTaxIDType = "mx_rfc"
	CustomerTaxIDTypeMyFrp    CustomerTaxIDType = "my_frp"
	CustomerTaxIDTypeMyItn    CustomerTaxIDType = "my_itn"
	CustomerTaxIDTypeMySst    CustomerTaxIDType = "my_sst"
	CustomerTaxIDTypeNgTin    CustomerTaxIDType = "ng_tin"
	CustomerTaxIDTypeNoVat    CustomerTaxIDType = "no_vat"
	CustomerTaxIDTypeNoVoec   CustomerTaxIDType = "no_voec"
	CustomerTaxIDTypeNzGst    CustomerTaxIDType = "nz_gst"
	CustomerTaxIDTypeOmVat    CustomerTaxIDType = "om_vat"
	CustomerTaxIDTypePeRuc    CustomerTaxIDType = "pe_ruc"
	CustomerTaxIDTypePhTin    CustomerTaxIDType = "ph_tin"
	CustomerTaxIDTypeRoTin    CustomerTaxIDType = "ro_tin"
	CustomerTaxIDTypeRsPib    CustomerTaxIDType = "rs_pib"
	CustomerTaxIDTypeRuInn    CustomerTaxIDType = "ru_inn"
	CustomerTaxIDTypeRuKpp    CustomerTaxIDType = "ru_kpp"
	CustomerTaxIDTypeSaVat    CustomerTaxIDType = "sa_vat"
	CustomerTaxIDTypeSgGst    CustomerTaxIDType = "sg_gst"
	CustomerTaxIDTypeSgUen    CustomerTaxIDType = "sg_uen"
	CustomerTaxIDTypeSiTin    CustomerTaxIDType = "si_tin"
	CustomerTaxIDTypeSvNit    CustomerTaxIDType = "sv_nit"
	CustomerTaxIDTypeThVat    CustomerTaxIDType = "th_vat"
	CustomerTaxIDTypeTrTin    CustomerTaxIDType = "tr_tin"
	CustomerTaxIDTypeTwVat    CustomerTaxIDType = "tw_vat"
	CustomerTaxIDTypeUaVat    CustomerTaxIDType = "ua_vat"
	CustomerTaxIDTypeUsEin    CustomerTaxIDType = "us_ein"
	CustomerTaxIDTypeUyRuc    CustomerTaxIDType = "uy_ruc"
	CustomerTaxIDTypeVeRif    CustomerTaxIDType = "ve_rif"
	CustomerTaxIDTypeVnTin    CustomerTaxIDType = "vn_tin"
	CustomerTaxIDTypeZaVat    CustomerTaxIDType = "za_vat"
)

func (r CustomerTaxIDType) IsKnown() bool {
	switch r {
	case CustomerTaxIDTypeAdNrt, CustomerTaxIDTypeAeTrn, CustomerTaxIDTypeArCuit, CustomerTaxIDTypeEuVat, CustomerTaxIDTypeAuAbn, CustomerTaxIDTypeAuArn, CustomerTaxIDTypeBgUic, CustomerTaxIDTypeBhVat, CustomerTaxIDTypeBoTin, CustomerTaxIDTypeBrCnpj, CustomerTaxIDTypeBrCpf, CustomerTaxIDTypeCaBn, CustomerTaxIDTypeCaGstHst, CustomerTaxIDTypeCaPstBc, CustomerTaxIDTypeCaPstMB, CustomerTaxIDTypeCaPstSk, CustomerTaxIDTypeCaQst, CustomerTaxIDTypeChVat, CustomerTaxIDTypeClTin, CustomerTaxIDTypeCnTin, CustomerTaxIDTypeCoNit, CustomerTaxIDTypeCrTin, CustomerTaxIDTypeDoRcn, CustomerTaxIDTypeEcRuc, CustomerTaxIDTypeEgTin, CustomerTaxIDTypeEsCif, CustomerTaxIDTypeEuOssVat, CustomerTaxIDTypeGBVat, CustomerTaxIDTypeGeVat, CustomerTaxIDTypeHkBr, CustomerTaxIDTypeHuTin, CustomerTaxIDTypeIDNpwp, CustomerTaxIDTypeIlVat, CustomerTaxIDTypeInGst, CustomerTaxIDTypeIsVat, CustomerTaxIDTypeJpCn, CustomerTaxIDTypeJpRn, CustomerTaxIDTypeJpTrn, CustomerTaxIDTypeKePin, CustomerTaxIDTypeKrBrn, CustomerTaxIDTypeKzBin, CustomerTaxIDTypeLiUid, CustomerTaxIDTypeMxRfc, CustomerTaxIDTypeMyFrp, CustomerTaxIDTypeMyItn, CustomerTaxIDTypeMySst, CustomerTaxIDTypeNgTin, CustomerTaxIDTypeNoVat, CustomerTaxIDTypeNoVoec, CustomerTaxIDTypeNzGst, CustomerTaxIDTypeOmVat, CustomerTaxIDTypePeRuc, CustomerTaxIDTypePhTin, CustomerTaxIDTypeRoTin, CustomerTaxIDTypeRsPib, CustomerTaxIDTypeRuInn, CustomerTaxIDTypeRuKpp, CustomerTaxIDTypeSaVat, CustomerTaxIDTypeSgGst, CustomerTaxIDTypeSgUen, CustomerTaxIDTypeSiTin, CustomerTaxIDTypeSvNit, CustomerTaxIDTypeThVat, CustomerTaxIDTypeTrTin, CustomerTaxIDTypeTwVat, CustomerTaxIDTypeUaVat, CustomerTaxIDTypeUsEin, CustomerTaxIDTypeUyRuc, CustomerTaxIDTypeVeRif, CustomerTaxIDTypeVnTin, CustomerTaxIDTypeZaVat:
		return true
	}
	return false
}

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

func (r CustomerAccountingSyncConfigurationAccountingProvidersProviderType) IsKnown() bool {
	switch r {
	case CustomerAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks, CustomerAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite:
		return true
	}
	return false
}

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
	Email param.Field[string] `json:"email,required" format:"email"`
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
	TaxConfiguration       param.Field[CustomerNewParamsTaxConfigurationUnion]  `json:"tax_configuration"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
	// | France               | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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

func (r CustomerNewParamsPaymentProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsPaymentProviderQuickbooks, CustomerNewParamsPaymentProviderBillCom, CustomerNewParamsPaymentProviderStripeCharge, CustomerNewParamsPaymentProviderStripeInvoice, CustomerNewParamsPaymentProviderNetsuite:
		return true
	}
	return false
}

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

type CustomerNewParamsTaxConfiguration struct {
	TaxExempt        param.Field[bool]                                         `json:"tax_exempt,required"`
	TaxProvider      param.Field[CustomerNewParamsTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                       `json:"tax_exemption_code"`
}

func (r CustomerNewParamsTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {}

// Satisfied by [CustomerNewParamsTaxConfigurationNewAvalaraTaxConfiguration],
// [CustomerNewParamsTaxConfigurationNewTaxJarConfiguration],
// [CustomerNewParamsTaxConfiguration].
type CustomerNewParamsTaxConfigurationUnion interface {
	implementsCustomerNewParamsTaxConfigurationUnion()
}

type CustomerNewParamsTaxConfigurationNewAvalaraTaxConfiguration struct {
	TaxExempt        param.Field[bool]                                                                   `json:"tax_exempt,required"`
	TaxProvider      param.Field[CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                                                 `json:"tax_exemption_code"`
}

func (r CustomerNewParamsTaxConfigurationNewAvalaraTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfigurationNewAvalaraTaxConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {
}

type CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider = "avalara"
)

func (r CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara:
		return true
	}
	return false
}

type CustomerNewParamsTaxConfigurationNewTaxJarConfiguration struct {
	TaxExempt   param.Field[bool]                                                               `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerNewParamsTaxConfigurationNewTaxJarConfigurationTaxProvider] `json:"tax_provider,required"`
}

func (r CustomerNewParamsTaxConfigurationNewTaxJarConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfigurationNewTaxJarConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {
}

type CustomerNewParamsTaxConfigurationNewTaxJarConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationNewTaxJarConfigurationTaxProviderTaxjar CustomerNewParamsTaxConfigurationNewTaxJarConfigurationTaxProvider = "taxjar"
)

func (r CustomerNewParamsTaxConfigurationNewTaxJarConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationNewTaxJarConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

type CustomerNewParamsTaxConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationTaxProviderAvalara CustomerNewParamsTaxConfigurationTaxProvider = "avalara"
	CustomerNewParamsTaxConfigurationTaxProviderTaxjar  CustomerNewParamsTaxConfigurationTaxProvider = "taxjar"
)

func (r CustomerNewParamsTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationTaxProviderAvalara, CustomerNewParamsTaxConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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
	CustomerNewParamsTaxIDCountryAr CustomerNewParamsTaxIDCountry = "AR"
	CustomerNewParamsTaxIDCountryAt CustomerNewParamsTaxIDCountry = "AT"
	CustomerNewParamsTaxIDCountryAu CustomerNewParamsTaxIDCountry = "AU"
	CustomerNewParamsTaxIDCountryBe CustomerNewParamsTaxIDCountry = "BE"
	CustomerNewParamsTaxIDCountryBg CustomerNewParamsTaxIDCountry = "BG"
	CustomerNewParamsTaxIDCountryBh CustomerNewParamsTaxIDCountry = "BH"
	CustomerNewParamsTaxIDCountryBo CustomerNewParamsTaxIDCountry = "BO"
	CustomerNewParamsTaxIDCountryBr CustomerNewParamsTaxIDCountry = "BR"
	CustomerNewParamsTaxIDCountryCa CustomerNewParamsTaxIDCountry = "CA"
	CustomerNewParamsTaxIDCountryCh CustomerNewParamsTaxIDCountry = "CH"
	CustomerNewParamsTaxIDCountryCl CustomerNewParamsTaxIDCountry = "CL"
	CustomerNewParamsTaxIDCountryCn CustomerNewParamsTaxIDCountry = "CN"
	CustomerNewParamsTaxIDCountryCo CustomerNewParamsTaxIDCountry = "CO"
	CustomerNewParamsTaxIDCountryCr CustomerNewParamsTaxIDCountry = "CR"
	CustomerNewParamsTaxIDCountryCy CustomerNewParamsTaxIDCountry = "CY"
	CustomerNewParamsTaxIDCountryCz CustomerNewParamsTaxIDCountry = "CZ"
	CustomerNewParamsTaxIDCountryDe CustomerNewParamsTaxIDCountry = "DE"
	CustomerNewParamsTaxIDCountryDk CustomerNewParamsTaxIDCountry = "DK"
	CustomerNewParamsTaxIDCountryEe CustomerNewParamsTaxIDCountry = "EE"
	CustomerNewParamsTaxIDCountryDo CustomerNewParamsTaxIDCountry = "DO"
	CustomerNewParamsTaxIDCountryEc CustomerNewParamsTaxIDCountry = "EC"
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
	CustomerNewParamsTaxIDCountryKz CustomerNewParamsTaxIDCountry = "KZ"
	CustomerNewParamsTaxIDCountryLi CustomerNewParamsTaxIDCountry = "LI"
	CustomerNewParamsTaxIDCountryLt CustomerNewParamsTaxIDCountry = "LT"
	CustomerNewParamsTaxIDCountryLu CustomerNewParamsTaxIDCountry = "LU"
	CustomerNewParamsTaxIDCountryLv CustomerNewParamsTaxIDCountry = "LV"
	CustomerNewParamsTaxIDCountryMt CustomerNewParamsTaxIDCountry = "MT"
	CustomerNewParamsTaxIDCountryMx CustomerNewParamsTaxIDCountry = "MX"
	CustomerNewParamsTaxIDCountryMy CustomerNewParamsTaxIDCountry = "MY"
	CustomerNewParamsTaxIDCountryNg CustomerNewParamsTaxIDCountry = "NG"
	CustomerNewParamsTaxIDCountryNl CustomerNewParamsTaxIDCountry = "NL"
	CustomerNewParamsTaxIDCountryNo CustomerNewParamsTaxIDCountry = "NO"
	CustomerNewParamsTaxIDCountryNz CustomerNewParamsTaxIDCountry = "NZ"
	CustomerNewParamsTaxIDCountryOm CustomerNewParamsTaxIDCountry = "OM"
	CustomerNewParamsTaxIDCountryPe CustomerNewParamsTaxIDCountry = "PE"
	CustomerNewParamsTaxIDCountryPh CustomerNewParamsTaxIDCountry = "PH"
	CustomerNewParamsTaxIDCountryPl CustomerNewParamsTaxIDCountry = "PL"
	CustomerNewParamsTaxIDCountryPt CustomerNewParamsTaxIDCountry = "PT"
	CustomerNewParamsTaxIDCountryRo CustomerNewParamsTaxIDCountry = "RO"
	CustomerNewParamsTaxIDCountryRs CustomerNewParamsTaxIDCountry = "RS"
	CustomerNewParamsTaxIDCountryRu CustomerNewParamsTaxIDCountry = "RU"
	CustomerNewParamsTaxIDCountrySa CustomerNewParamsTaxIDCountry = "SA"
	CustomerNewParamsTaxIDCountrySe CustomerNewParamsTaxIDCountry = "SE"
	CustomerNewParamsTaxIDCountrySg CustomerNewParamsTaxIDCountry = "SG"
	CustomerNewParamsTaxIDCountrySi CustomerNewParamsTaxIDCountry = "SI"
	CustomerNewParamsTaxIDCountrySk CustomerNewParamsTaxIDCountry = "SK"
	CustomerNewParamsTaxIDCountrySv CustomerNewParamsTaxIDCountry = "SV"
	CustomerNewParamsTaxIDCountryTh CustomerNewParamsTaxIDCountry = "TH"
	CustomerNewParamsTaxIDCountryTr CustomerNewParamsTaxIDCountry = "TR"
	CustomerNewParamsTaxIDCountryTw CustomerNewParamsTaxIDCountry = "TW"
	CustomerNewParamsTaxIDCountryUa CustomerNewParamsTaxIDCountry = "UA"
	CustomerNewParamsTaxIDCountryUs CustomerNewParamsTaxIDCountry = "US"
	CustomerNewParamsTaxIDCountryUy CustomerNewParamsTaxIDCountry = "UY"
	CustomerNewParamsTaxIDCountryVe CustomerNewParamsTaxIDCountry = "VE"
	CustomerNewParamsTaxIDCountryVn CustomerNewParamsTaxIDCountry = "VN"
	CustomerNewParamsTaxIDCountryZa CustomerNewParamsTaxIDCountry = "ZA"
)

func (r CustomerNewParamsTaxIDCountry) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxIDCountryAd, CustomerNewParamsTaxIDCountryAe, CustomerNewParamsTaxIDCountryAr, CustomerNewParamsTaxIDCountryAt, CustomerNewParamsTaxIDCountryAu, CustomerNewParamsTaxIDCountryBe, CustomerNewParamsTaxIDCountryBg, CustomerNewParamsTaxIDCountryBh, CustomerNewParamsTaxIDCountryBo, CustomerNewParamsTaxIDCountryBr, CustomerNewParamsTaxIDCountryCa, CustomerNewParamsTaxIDCountryCh, CustomerNewParamsTaxIDCountryCl, CustomerNewParamsTaxIDCountryCn, CustomerNewParamsTaxIDCountryCo, CustomerNewParamsTaxIDCountryCr, CustomerNewParamsTaxIDCountryCy, CustomerNewParamsTaxIDCountryCz, CustomerNewParamsTaxIDCountryDe, CustomerNewParamsTaxIDCountryDk, CustomerNewParamsTaxIDCountryEe, CustomerNewParamsTaxIDCountryDo, CustomerNewParamsTaxIDCountryEc, CustomerNewParamsTaxIDCountryEg, CustomerNewParamsTaxIDCountryEs, CustomerNewParamsTaxIDCountryEu, CustomerNewParamsTaxIDCountryFi, CustomerNewParamsTaxIDCountryFr, CustomerNewParamsTaxIDCountryGB, CustomerNewParamsTaxIDCountryGe, CustomerNewParamsTaxIDCountryGr, CustomerNewParamsTaxIDCountryHk, CustomerNewParamsTaxIDCountryHr, CustomerNewParamsTaxIDCountryHu, CustomerNewParamsTaxIDCountryID, CustomerNewParamsTaxIDCountryIe, CustomerNewParamsTaxIDCountryIl, CustomerNewParamsTaxIDCountryIn, CustomerNewParamsTaxIDCountryIs, CustomerNewParamsTaxIDCountryIt, CustomerNewParamsTaxIDCountryJp, CustomerNewParamsTaxIDCountryKe, CustomerNewParamsTaxIDCountryKr, CustomerNewParamsTaxIDCountryKz, CustomerNewParamsTaxIDCountryLi, CustomerNewParamsTaxIDCountryLt, CustomerNewParamsTaxIDCountryLu, CustomerNewParamsTaxIDCountryLv, CustomerNewParamsTaxIDCountryMt, CustomerNewParamsTaxIDCountryMx, CustomerNewParamsTaxIDCountryMy, CustomerNewParamsTaxIDCountryNg, CustomerNewParamsTaxIDCountryNl, CustomerNewParamsTaxIDCountryNo, CustomerNewParamsTaxIDCountryNz, CustomerNewParamsTaxIDCountryOm, CustomerNewParamsTaxIDCountryPe, CustomerNewParamsTaxIDCountryPh, CustomerNewParamsTaxIDCountryPl, CustomerNewParamsTaxIDCountryPt, CustomerNewParamsTaxIDCountryRo, CustomerNewParamsTaxIDCountryRs, CustomerNewParamsTaxIDCountryRu, CustomerNewParamsTaxIDCountrySa, CustomerNewParamsTaxIDCountrySe, CustomerNewParamsTaxIDCountrySg, CustomerNewParamsTaxIDCountrySi, CustomerNewParamsTaxIDCountrySk, CustomerNewParamsTaxIDCountrySv, CustomerNewParamsTaxIDCountryTh, CustomerNewParamsTaxIDCountryTr, CustomerNewParamsTaxIDCountryTw, CustomerNewParamsTaxIDCountryUa, CustomerNewParamsTaxIDCountryUs, CustomerNewParamsTaxIDCountryUy, CustomerNewParamsTaxIDCountryVe, CustomerNewParamsTaxIDCountryVn, CustomerNewParamsTaxIDCountryZa:
		return true
	}
	return false
}

type CustomerNewParamsTaxIDType string

const (
	CustomerNewParamsTaxIDTypeAdNrt    CustomerNewParamsTaxIDType = "ad_nrt"
	CustomerNewParamsTaxIDTypeAeTrn    CustomerNewParamsTaxIDType = "ae_trn"
	CustomerNewParamsTaxIDTypeArCuit   CustomerNewParamsTaxIDType = "ar_cuit"
	CustomerNewParamsTaxIDTypeEuVat    CustomerNewParamsTaxIDType = "eu_vat"
	CustomerNewParamsTaxIDTypeAuAbn    CustomerNewParamsTaxIDType = "au_abn"
	CustomerNewParamsTaxIDTypeAuArn    CustomerNewParamsTaxIDType = "au_arn"
	CustomerNewParamsTaxIDTypeBgUic    CustomerNewParamsTaxIDType = "bg_uic"
	CustomerNewParamsTaxIDTypeBhVat    CustomerNewParamsTaxIDType = "bh_vat"
	CustomerNewParamsTaxIDTypeBoTin    CustomerNewParamsTaxIDType = "bo_tin"
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
	CustomerNewParamsTaxIDTypeCnTin    CustomerNewParamsTaxIDType = "cn_tin"
	CustomerNewParamsTaxIDTypeCoNit    CustomerNewParamsTaxIDType = "co_nit"
	CustomerNewParamsTaxIDTypeCrTin    CustomerNewParamsTaxIDType = "cr_tin"
	CustomerNewParamsTaxIDTypeDoRcn    CustomerNewParamsTaxIDType = "do_rcn"
	CustomerNewParamsTaxIDTypeEcRuc    CustomerNewParamsTaxIDType = "ec_ruc"
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
	CustomerNewParamsTaxIDTypeKzBin    CustomerNewParamsTaxIDType = "kz_bin"
	CustomerNewParamsTaxIDTypeLiUid    CustomerNewParamsTaxIDType = "li_uid"
	CustomerNewParamsTaxIDTypeMxRfc    CustomerNewParamsTaxIDType = "mx_rfc"
	CustomerNewParamsTaxIDTypeMyFrp    CustomerNewParamsTaxIDType = "my_frp"
	CustomerNewParamsTaxIDTypeMyItn    CustomerNewParamsTaxIDType = "my_itn"
	CustomerNewParamsTaxIDTypeMySst    CustomerNewParamsTaxIDType = "my_sst"
	CustomerNewParamsTaxIDTypeNgTin    CustomerNewParamsTaxIDType = "ng_tin"
	CustomerNewParamsTaxIDTypeNoVat    CustomerNewParamsTaxIDType = "no_vat"
	CustomerNewParamsTaxIDTypeNoVoec   CustomerNewParamsTaxIDType = "no_voec"
	CustomerNewParamsTaxIDTypeNzGst    CustomerNewParamsTaxIDType = "nz_gst"
	CustomerNewParamsTaxIDTypeOmVat    CustomerNewParamsTaxIDType = "om_vat"
	CustomerNewParamsTaxIDTypePeRuc    CustomerNewParamsTaxIDType = "pe_ruc"
	CustomerNewParamsTaxIDTypePhTin    CustomerNewParamsTaxIDType = "ph_tin"
	CustomerNewParamsTaxIDTypeRoTin    CustomerNewParamsTaxIDType = "ro_tin"
	CustomerNewParamsTaxIDTypeRsPib    CustomerNewParamsTaxIDType = "rs_pib"
	CustomerNewParamsTaxIDTypeRuInn    CustomerNewParamsTaxIDType = "ru_inn"
	CustomerNewParamsTaxIDTypeRuKpp    CustomerNewParamsTaxIDType = "ru_kpp"
	CustomerNewParamsTaxIDTypeSaVat    CustomerNewParamsTaxIDType = "sa_vat"
	CustomerNewParamsTaxIDTypeSgGst    CustomerNewParamsTaxIDType = "sg_gst"
	CustomerNewParamsTaxIDTypeSgUen    CustomerNewParamsTaxIDType = "sg_uen"
	CustomerNewParamsTaxIDTypeSiTin    CustomerNewParamsTaxIDType = "si_tin"
	CustomerNewParamsTaxIDTypeSvNit    CustomerNewParamsTaxIDType = "sv_nit"
	CustomerNewParamsTaxIDTypeThVat    CustomerNewParamsTaxIDType = "th_vat"
	CustomerNewParamsTaxIDTypeTrTin    CustomerNewParamsTaxIDType = "tr_tin"
	CustomerNewParamsTaxIDTypeTwVat    CustomerNewParamsTaxIDType = "tw_vat"
	CustomerNewParamsTaxIDTypeUaVat    CustomerNewParamsTaxIDType = "ua_vat"
	CustomerNewParamsTaxIDTypeUsEin    CustomerNewParamsTaxIDType = "us_ein"
	CustomerNewParamsTaxIDTypeUyRuc    CustomerNewParamsTaxIDType = "uy_ruc"
	CustomerNewParamsTaxIDTypeVeRif    CustomerNewParamsTaxIDType = "ve_rif"
	CustomerNewParamsTaxIDTypeVnTin    CustomerNewParamsTaxIDType = "vn_tin"
	CustomerNewParamsTaxIDTypeZaVat    CustomerNewParamsTaxIDType = "za_vat"
)

func (r CustomerNewParamsTaxIDType) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxIDTypeAdNrt, CustomerNewParamsTaxIDTypeAeTrn, CustomerNewParamsTaxIDTypeArCuit, CustomerNewParamsTaxIDTypeEuVat, CustomerNewParamsTaxIDTypeAuAbn, CustomerNewParamsTaxIDTypeAuArn, CustomerNewParamsTaxIDTypeBgUic, CustomerNewParamsTaxIDTypeBhVat, CustomerNewParamsTaxIDTypeBoTin, CustomerNewParamsTaxIDTypeBrCnpj, CustomerNewParamsTaxIDTypeBrCpf, CustomerNewParamsTaxIDTypeCaBn, CustomerNewParamsTaxIDTypeCaGstHst, CustomerNewParamsTaxIDTypeCaPstBc, CustomerNewParamsTaxIDTypeCaPstMB, CustomerNewParamsTaxIDTypeCaPstSk, CustomerNewParamsTaxIDTypeCaQst, CustomerNewParamsTaxIDTypeChVat, CustomerNewParamsTaxIDTypeClTin, CustomerNewParamsTaxIDTypeCnTin, CustomerNewParamsTaxIDTypeCoNit, CustomerNewParamsTaxIDTypeCrTin, CustomerNewParamsTaxIDTypeDoRcn, CustomerNewParamsTaxIDTypeEcRuc, CustomerNewParamsTaxIDTypeEgTin, CustomerNewParamsTaxIDTypeEsCif, CustomerNewParamsTaxIDTypeEuOssVat, CustomerNewParamsTaxIDTypeGBVat, CustomerNewParamsTaxIDTypeGeVat, CustomerNewParamsTaxIDTypeHkBr, CustomerNewParamsTaxIDTypeHuTin, CustomerNewParamsTaxIDTypeIDNpwp, CustomerNewParamsTaxIDTypeIlVat, CustomerNewParamsTaxIDTypeInGst, CustomerNewParamsTaxIDTypeIsVat, CustomerNewParamsTaxIDTypeJpCn, CustomerNewParamsTaxIDTypeJpRn, CustomerNewParamsTaxIDTypeJpTrn, CustomerNewParamsTaxIDTypeKePin, CustomerNewParamsTaxIDTypeKrBrn, CustomerNewParamsTaxIDTypeKzBin, CustomerNewParamsTaxIDTypeLiUid, CustomerNewParamsTaxIDTypeMxRfc, CustomerNewParamsTaxIDTypeMyFrp, CustomerNewParamsTaxIDTypeMyItn, CustomerNewParamsTaxIDTypeMySst, CustomerNewParamsTaxIDTypeNgTin, CustomerNewParamsTaxIDTypeNoVat, CustomerNewParamsTaxIDTypeNoVoec, CustomerNewParamsTaxIDTypeNzGst, CustomerNewParamsTaxIDTypeOmVat, CustomerNewParamsTaxIDTypePeRuc, CustomerNewParamsTaxIDTypePhTin, CustomerNewParamsTaxIDTypeRoTin, CustomerNewParamsTaxIDTypeRsPib, CustomerNewParamsTaxIDTypeRuInn, CustomerNewParamsTaxIDTypeRuKpp, CustomerNewParamsTaxIDTypeSaVat, CustomerNewParamsTaxIDTypeSgGst, CustomerNewParamsTaxIDTypeSgUen, CustomerNewParamsTaxIDTypeSiTin, CustomerNewParamsTaxIDTypeSvNit, CustomerNewParamsTaxIDTypeThVat, CustomerNewParamsTaxIDTypeTrTin, CustomerNewParamsTaxIDTypeTwVat, CustomerNewParamsTaxIDTypeUaVat, CustomerNewParamsTaxIDTypeUsEin, CustomerNewParamsTaxIDTypeUyRuc, CustomerNewParamsTaxIDTypeVeRif, CustomerNewParamsTaxIDTypeVnTin, CustomerNewParamsTaxIDTypeZaVat:
		return true
	}
	return false
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
	Email         param.Field[string] `json:"email" format:"email"`
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
	TaxConfiguration       param.Field[CustomerUpdateParamsTaxConfigurationUnion]  `json:"tax_configuration"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
	// | France               | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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

func (r CustomerUpdateParamsPaymentProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsPaymentProviderQuickbooks, CustomerUpdateParamsPaymentProviderBillCom, CustomerUpdateParamsPaymentProviderStripeCharge, CustomerUpdateParamsPaymentProviderStripeInvoice, CustomerUpdateParamsPaymentProviderNetsuite:
		return true
	}
	return false
}

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

type CustomerUpdateParamsTaxConfiguration struct {
	TaxExempt        param.Field[bool]                                            `json:"tax_exempt,required"`
	TaxProvider      param.Field[CustomerUpdateParamsTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                          `json:"tax_exemption_code"`
}

func (r CustomerUpdateParamsTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {}

// Satisfied by [CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfiguration],
// [CustomerUpdateParamsTaxConfigurationNewTaxJarConfiguration],
// [CustomerUpdateParamsTaxConfiguration].
type CustomerUpdateParamsTaxConfigurationUnion interface {
	implementsCustomerUpdateParamsTaxConfigurationUnion()
}

type CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfiguration struct {
	TaxExempt        param.Field[bool]                                                                      `json:"tax_exempt,required"`
	TaxProvider      param.Field[CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                                                    `json:"tax_exemption_code"`
}

func (r CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {
}

type CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider = "avalara"
)

func (r CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara:
		return true
	}
	return false
}

type CustomerUpdateParamsTaxConfigurationNewTaxJarConfiguration struct {
	TaxExempt   param.Field[bool]                                                                  `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateParamsTaxConfigurationNewTaxJarConfigurationTaxProvider] `json:"tax_provider,required"`
}

func (r CustomerUpdateParamsTaxConfigurationNewTaxJarConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfigurationNewTaxJarConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {
}

type CustomerUpdateParamsTaxConfigurationNewTaxJarConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationNewTaxJarConfigurationTaxProviderTaxjar CustomerUpdateParamsTaxConfigurationNewTaxJarConfigurationTaxProvider = "taxjar"
)

func (r CustomerUpdateParamsTaxConfigurationNewTaxJarConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationNewTaxJarConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

type CustomerUpdateParamsTaxConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationTaxProviderAvalara CustomerUpdateParamsTaxConfigurationTaxProvider = "avalara"
	CustomerUpdateParamsTaxConfigurationTaxProviderTaxjar  CustomerUpdateParamsTaxConfigurationTaxProvider = "taxjar"
)

func (r CustomerUpdateParamsTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationTaxProviderAvalara, CustomerUpdateParamsTaxConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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
	CustomerUpdateParamsTaxIDCountryAr CustomerUpdateParamsTaxIDCountry = "AR"
	CustomerUpdateParamsTaxIDCountryAt CustomerUpdateParamsTaxIDCountry = "AT"
	CustomerUpdateParamsTaxIDCountryAu CustomerUpdateParamsTaxIDCountry = "AU"
	CustomerUpdateParamsTaxIDCountryBe CustomerUpdateParamsTaxIDCountry = "BE"
	CustomerUpdateParamsTaxIDCountryBg CustomerUpdateParamsTaxIDCountry = "BG"
	CustomerUpdateParamsTaxIDCountryBh CustomerUpdateParamsTaxIDCountry = "BH"
	CustomerUpdateParamsTaxIDCountryBo CustomerUpdateParamsTaxIDCountry = "BO"
	CustomerUpdateParamsTaxIDCountryBr CustomerUpdateParamsTaxIDCountry = "BR"
	CustomerUpdateParamsTaxIDCountryCa CustomerUpdateParamsTaxIDCountry = "CA"
	CustomerUpdateParamsTaxIDCountryCh CustomerUpdateParamsTaxIDCountry = "CH"
	CustomerUpdateParamsTaxIDCountryCl CustomerUpdateParamsTaxIDCountry = "CL"
	CustomerUpdateParamsTaxIDCountryCn CustomerUpdateParamsTaxIDCountry = "CN"
	CustomerUpdateParamsTaxIDCountryCo CustomerUpdateParamsTaxIDCountry = "CO"
	CustomerUpdateParamsTaxIDCountryCr CustomerUpdateParamsTaxIDCountry = "CR"
	CustomerUpdateParamsTaxIDCountryCy CustomerUpdateParamsTaxIDCountry = "CY"
	CustomerUpdateParamsTaxIDCountryCz CustomerUpdateParamsTaxIDCountry = "CZ"
	CustomerUpdateParamsTaxIDCountryDe CustomerUpdateParamsTaxIDCountry = "DE"
	CustomerUpdateParamsTaxIDCountryDk CustomerUpdateParamsTaxIDCountry = "DK"
	CustomerUpdateParamsTaxIDCountryEe CustomerUpdateParamsTaxIDCountry = "EE"
	CustomerUpdateParamsTaxIDCountryDo CustomerUpdateParamsTaxIDCountry = "DO"
	CustomerUpdateParamsTaxIDCountryEc CustomerUpdateParamsTaxIDCountry = "EC"
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
	CustomerUpdateParamsTaxIDCountryKz CustomerUpdateParamsTaxIDCountry = "KZ"
	CustomerUpdateParamsTaxIDCountryLi CustomerUpdateParamsTaxIDCountry = "LI"
	CustomerUpdateParamsTaxIDCountryLt CustomerUpdateParamsTaxIDCountry = "LT"
	CustomerUpdateParamsTaxIDCountryLu CustomerUpdateParamsTaxIDCountry = "LU"
	CustomerUpdateParamsTaxIDCountryLv CustomerUpdateParamsTaxIDCountry = "LV"
	CustomerUpdateParamsTaxIDCountryMt CustomerUpdateParamsTaxIDCountry = "MT"
	CustomerUpdateParamsTaxIDCountryMx CustomerUpdateParamsTaxIDCountry = "MX"
	CustomerUpdateParamsTaxIDCountryMy CustomerUpdateParamsTaxIDCountry = "MY"
	CustomerUpdateParamsTaxIDCountryNg CustomerUpdateParamsTaxIDCountry = "NG"
	CustomerUpdateParamsTaxIDCountryNl CustomerUpdateParamsTaxIDCountry = "NL"
	CustomerUpdateParamsTaxIDCountryNo CustomerUpdateParamsTaxIDCountry = "NO"
	CustomerUpdateParamsTaxIDCountryNz CustomerUpdateParamsTaxIDCountry = "NZ"
	CustomerUpdateParamsTaxIDCountryOm CustomerUpdateParamsTaxIDCountry = "OM"
	CustomerUpdateParamsTaxIDCountryPe CustomerUpdateParamsTaxIDCountry = "PE"
	CustomerUpdateParamsTaxIDCountryPh CustomerUpdateParamsTaxIDCountry = "PH"
	CustomerUpdateParamsTaxIDCountryPl CustomerUpdateParamsTaxIDCountry = "PL"
	CustomerUpdateParamsTaxIDCountryPt CustomerUpdateParamsTaxIDCountry = "PT"
	CustomerUpdateParamsTaxIDCountryRo CustomerUpdateParamsTaxIDCountry = "RO"
	CustomerUpdateParamsTaxIDCountryRs CustomerUpdateParamsTaxIDCountry = "RS"
	CustomerUpdateParamsTaxIDCountryRu CustomerUpdateParamsTaxIDCountry = "RU"
	CustomerUpdateParamsTaxIDCountrySa CustomerUpdateParamsTaxIDCountry = "SA"
	CustomerUpdateParamsTaxIDCountrySe CustomerUpdateParamsTaxIDCountry = "SE"
	CustomerUpdateParamsTaxIDCountrySg CustomerUpdateParamsTaxIDCountry = "SG"
	CustomerUpdateParamsTaxIDCountrySi CustomerUpdateParamsTaxIDCountry = "SI"
	CustomerUpdateParamsTaxIDCountrySk CustomerUpdateParamsTaxIDCountry = "SK"
	CustomerUpdateParamsTaxIDCountrySv CustomerUpdateParamsTaxIDCountry = "SV"
	CustomerUpdateParamsTaxIDCountryTh CustomerUpdateParamsTaxIDCountry = "TH"
	CustomerUpdateParamsTaxIDCountryTr CustomerUpdateParamsTaxIDCountry = "TR"
	CustomerUpdateParamsTaxIDCountryTw CustomerUpdateParamsTaxIDCountry = "TW"
	CustomerUpdateParamsTaxIDCountryUa CustomerUpdateParamsTaxIDCountry = "UA"
	CustomerUpdateParamsTaxIDCountryUs CustomerUpdateParamsTaxIDCountry = "US"
	CustomerUpdateParamsTaxIDCountryUy CustomerUpdateParamsTaxIDCountry = "UY"
	CustomerUpdateParamsTaxIDCountryVe CustomerUpdateParamsTaxIDCountry = "VE"
	CustomerUpdateParamsTaxIDCountryVn CustomerUpdateParamsTaxIDCountry = "VN"
	CustomerUpdateParamsTaxIDCountryZa CustomerUpdateParamsTaxIDCountry = "ZA"
)

func (r CustomerUpdateParamsTaxIDCountry) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxIDCountryAd, CustomerUpdateParamsTaxIDCountryAe, CustomerUpdateParamsTaxIDCountryAr, CustomerUpdateParamsTaxIDCountryAt, CustomerUpdateParamsTaxIDCountryAu, CustomerUpdateParamsTaxIDCountryBe, CustomerUpdateParamsTaxIDCountryBg, CustomerUpdateParamsTaxIDCountryBh, CustomerUpdateParamsTaxIDCountryBo, CustomerUpdateParamsTaxIDCountryBr, CustomerUpdateParamsTaxIDCountryCa, CustomerUpdateParamsTaxIDCountryCh, CustomerUpdateParamsTaxIDCountryCl, CustomerUpdateParamsTaxIDCountryCn, CustomerUpdateParamsTaxIDCountryCo, CustomerUpdateParamsTaxIDCountryCr, CustomerUpdateParamsTaxIDCountryCy, CustomerUpdateParamsTaxIDCountryCz, CustomerUpdateParamsTaxIDCountryDe, CustomerUpdateParamsTaxIDCountryDk, CustomerUpdateParamsTaxIDCountryEe, CustomerUpdateParamsTaxIDCountryDo, CustomerUpdateParamsTaxIDCountryEc, CustomerUpdateParamsTaxIDCountryEg, CustomerUpdateParamsTaxIDCountryEs, CustomerUpdateParamsTaxIDCountryEu, CustomerUpdateParamsTaxIDCountryFi, CustomerUpdateParamsTaxIDCountryFr, CustomerUpdateParamsTaxIDCountryGB, CustomerUpdateParamsTaxIDCountryGe, CustomerUpdateParamsTaxIDCountryGr, CustomerUpdateParamsTaxIDCountryHk, CustomerUpdateParamsTaxIDCountryHr, CustomerUpdateParamsTaxIDCountryHu, CustomerUpdateParamsTaxIDCountryID, CustomerUpdateParamsTaxIDCountryIe, CustomerUpdateParamsTaxIDCountryIl, CustomerUpdateParamsTaxIDCountryIn, CustomerUpdateParamsTaxIDCountryIs, CustomerUpdateParamsTaxIDCountryIt, CustomerUpdateParamsTaxIDCountryJp, CustomerUpdateParamsTaxIDCountryKe, CustomerUpdateParamsTaxIDCountryKr, CustomerUpdateParamsTaxIDCountryKz, CustomerUpdateParamsTaxIDCountryLi, CustomerUpdateParamsTaxIDCountryLt, CustomerUpdateParamsTaxIDCountryLu, CustomerUpdateParamsTaxIDCountryLv, CustomerUpdateParamsTaxIDCountryMt, CustomerUpdateParamsTaxIDCountryMx, CustomerUpdateParamsTaxIDCountryMy, CustomerUpdateParamsTaxIDCountryNg, CustomerUpdateParamsTaxIDCountryNl, CustomerUpdateParamsTaxIDCountryNo, CustomerUpdateParamsTaxIDCountryNz, CustomerUpdateParamsTaxIDCountryOm, CustomerUpdateParamsTaxIDCountryPe, CustomerUpdateParamsTaxIDCountryPh, CustomerUpdateParamsTaxIDCountryPl, CustomerUpdateParamsTaxIDCountryPt, CustomerUpdateParamsTaxIDCountryRo, CustomerUpdateParamsTaxIDCountryRs, CustomerUpdateParamsTaxIDCountryRu, CustomerUpdateParamsTaxIDCountrySa, CustomerUpdateParamsTaxIDCountrySe, CustomerUpdateParamsTaxIDCountrySg, CustomerUpdateParamsTaxIDCountrySi, CustomerUpdateParamsTaxIDCountrySk, CustomerUpdateParamsTaxIDCountrySv, CustomerUpdateParamsTaxIDCountryTh, CustomerUpdateParamsTaxIDCountryTr, CustomerUpdateParamsTaxIDCountryTw, CustomerUpdateParamsTaxIDCountryUa, CustomerUpdateParamsTaxIDCountryUs, CustomerUpdateParamsTaxIDCountryUy, CustomerUpdateParamsTaxIDCountryVe, CustomerUpdateParamsTaxIDCountryVn, CustomerUpdateParamsTaxIDCountryZa:
		return true
	}
	return false
}

type CustomerUpdateParamsTaxIDType string

const (
	CustomerUpdateParamsTaxIDTypeAdNrt    CustomerUpdateParamsTaxIDType = "ad_nrt"
	CustomerUpdateParamsTaxIDTypeAeTrn    CustomerUpdateParamsTaxIDType = "ae_trn"
	CustomerUpdateParamsTaxIDTypeArCuit   CustomerUpdateParamsTaxIDType = "ar_cuit"
	CustomerUpdateParamsTaxIDTypeEuVat    CustomerUpdateParamsTaxIDType = "eu_vat"
	CustomerUpdateParamsTaxIDTypeAuAbn    CustomerUpdateParamsTaxIDType = "au_abn"
	CustomerUpdateParamsTaxIDTypeAuArn    CustomerUpdateParamsTaxIDType = "au_arn"
	CustomerUpdateParamsTaxIDTypeBgUic    CustomerUpdateParamsTaxIDType = "bg_uic"
	CustomerUpdateParamsTaxIDTypeBhVat    CustomerUpdateParamsTaxIDType = "bh_vat"
	CustomerUpdateParamsTaxIDTypeBoTin    CustomerUpdateParamsTaxIDType = "bo_tin"
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
	CustomerUpdateParamsTaxIDTypeCnTin    CustomerUpdateParamsTaxIDType = "cn_tin"
	CustomerUpdateParamsTaxIDTypeCoNit    CustomerUpdateParamsTaxIDType = "co_nit"
	CustomerUpdateParamsTaxIDTypeCrTin    CustomerUpdateParamsTaxIDType = "cr_tin"
	CustomerUpdateParamsTaxIDTypeDoRcn    CustomerUpdateParamsTaxIDType = "do_rcn"
	CustomerUpdateParamsTaxIDTypeEcRuc    CustomerUpdateParamsTaxIDType = "ec_ruc"
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
	CustomerUpdateParamsTaxIDTypeKzBin    CustomerUpdateParamsTaxIDType = "kz_bin"
	CustomerUpdateParamsTaxIDTypeLiUid    CustomerUpdateParamsTaxIDType = "li_uid"
	CustomerUpdateParamsTaxIDTypeMxRfc    CustomerUpdateParamsTaxIDType = "mx_rfc"
	CustomerUpdateParamsTaxIDTypeMyFrp    CustomerUpdateParamsTaxIDType = "my_frp"
	CustomerUpdateParamsTaxIDTypeMyItn    CustomerUpdateParamsTaxIDType = "my_itn"
	CustomerUpdateParamsTaxIDTypeMySst    CustomerUpdateParamsTaxIDType = "my_sst"
	CustomerUpdateParamsTaxIDTypeNgTin    CustomerUpdateParamsTaxIDType = "ng_tin"
	CustomerUpdateParamsTaxIDTypeNoVat    CustomerUpdateParamsTaxIDType = "no_vat"
	CustomerUpdateParamsTaxIDTypeNoVoec   CustomerUpdateParamsTaxIDType = "no_voec"
	CustomerUpdateParamsTaxIDTypeNzGst    CustomerUpdateParamsTaxIDType = "nz_gst"
	CustomerUpdateParamsTaxIDTypeOmVat    CustomerUpdateParamsTaxIDType = "om_vat"
	CustomerUpdateParamsTaxIDTypePeRuc    CustomerUpdateParamsTaxIDType = "pe_ruc"
	CustomerUpdateParamsTaxIDTypePhTin    CustomerUpdateParamsTaxIDType = "ph_tin"
	CustomerUpdateParamsTaxIDTypeRoTin    CustomerUpdateParamsTaxIDType = "ro_tin"
	CustomerUpdateParamsTaxIDTypeRsPib    CustomerUpdateParamsTaxIDType = "rs_pib"
	CustomerUpdateParamsTaxIDTypeRuInn    CustomerUpdateParamsTaxIDType = "ru_inn"
	CustomerUpdateParamsTaxIDTypeRuKpp    CustomerUpdateParamsTaxIDType = "ru_kpp"
	CustomerUpdateParamsTaxIDTypeSaVat    CustomerUpdateParamsTaxIDType = "sa_vat"
	CustomerUpdateParamsTaxIDTypeSgGst    CustomerUpdateParamsTaxIDType = "sg_gst"
	CustomerUpdateParamsTaxIDTypeSgUen    CustomerUpdateParamsTaxIDType = "sg_uen"
	CustomerUpdateParamsTaxIDTypeSiTin    CustomerUpdateParamsTaxIDType = "si_tin"
	CustomerUpdateParamsTaxIDTypeSvNit    CustomerUpdateParamsTaxIDType = "sv_nit"
	CustomerUpdateParamsTaxIDTypeThVat    CustomerUpdateParamsTaxIDType = "th_vat"
	CustomerUpdateParamsTaxIDTypeTrTin    CustomerUpdateParamsTaxIDType = "tr_tin"
	CustomerUpdateParamsTaxIDTypeTwVat    CustomerUpdateParamsTaxIDType = "tw_vat"
	CustomerUpdateParamsTaxIDTypeUaVat    CustomerUpdateParamsTaxIDType = "ua_vat"
	CustomerUpdateParamsTaxIDTypeUsEin    CustomerUpdateParamsTaxIDType = "us_ein"
	CustomerUpdateParamsTaxIDTypeUyRuc    CustomerUpdateParamsTaxIDType = "uy_ruc"
	CustomerUpdateParamsTaxIDTypeVeRif    CustomerUpdateParamsTaxIDType = "ve_rif"
	CustomerUpdateParamsTaxIDTypeVnTin    CustomerUpdateParamsTaxIDType = "vn_tin"
	CustomerUpdateParamsTaxIDTypeZaVat    CustomerUpdateParamsTaxIDType = "za_vat"
)

func (r CustomerUpdateParamsTaxIDType) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxIDTypeAdNrt, CustomerUpdateParamsTaxIDTypeAeTrn, CustomerUpdateParamsTaxIDTypeArCuit, CustomerUpdateParamsTaxIDTypeEuVat, CustomerUpdateParamsTaxIDTypeAuAbn, CustomerUpdateParamsTaxIDTypeAuArn, CustomerUpdateParamsTaxIDTypeBgUic, CustomerUpdateParamsTaxIDTypeBhVat, CustomerUpdateParamsTaxIDTypeBoTin, CustomerUpdateParamsTaxIDTypeBrCnpj, CustomerUpdateParamsTaxIDTypeBrCpf, CustomerUpdateParamsTaxIDTypeCaBn, CustomerUpdateParamsTaxIDTypeCaGstHst, CustomerUpdateParamsTaxIDTypeCaPstBc, CustomerUpdateParamsTaxIDTypeCaPstMB, CustomerUpdateParamsTaxIDTypeCaPstSk, CustomerUpdateParamsTaxIDTypeCaQst, CustomerUpdateParamsTaxIDTypeChVat, CustomerUpdateParamsTaxIDTypeClTin, CustomerUpdateParamsTaxIDTypeCnTin, CustomerUpdateParamsTaxIDTypeCoNit, CustomerUpdateParamsTaxIDTypeCrTin, CustomerUpdateParamsTaxIDTypeDoRcn, CustomerUpdateParamsTaxIDTypeEcRuc, CustomerUpdateParamsTaxIDTypeEgTin, CustomerUpdateParamsTaxIDTypeEsCif, CustomerUpdateParamsTaxIDTypeEuOssVat, CustomerUpdateParamsTaxIDTypeGBVat, CustomerUpdateParamsTaxIDTypeGeVat, CustomerUpdateParamsTaxIDTypeHkBr, CustomerUpdateParamsTaxIDTypeHuTin, CustomerUpdateParamsTaxIDTypeIDNpwp, CustomerUpdateParamsTaxIDTypeIlVat, CustomerUpdateParamsTaxIDTypeInGst, CustomerUpdateParamsTaxIDTypeIsVat, CustomerUpdateParamsTaxIDTypeJpCn, CustomerUpdateParamsTaxIDTypeJpRn, CustomerUpdateParamsTaxIDTypeJpTrn, CustomerUpdateParamsTaxIDTypeKePin, CustomerUpdateParamsTaxIDTypeKrBrn, CustomerUpdateParamsTaxIDTypeKzBin, CustomerUpdateParamsTaxIDTypeLiUid, CustomerUpdateParamsTaxIDTypeMxRfc, CustomerUpdateParamsTaxIDTypeMyFrp, CustomerUpdateParamsTaxIDTypeMyItn, CustomerUpdateParamsTaxIDTypeMySst, CustomerUpdateParamsTaxIDTypeNgTin, CustomerUpdateParamsTaxIDTypeNoVat, CustomerUpdateParamsTaxIDTypeNoVoec, CustomerUpdateParamsTaxIDTypeNzGst, CustomerUpdateParamsTaxIDTypeOmVat, CustomerUpdateParamsTaxIDTypePeRuc, CustomerUpdateParamsTaxIDTypePhTin, CustomerUpdateParamsTaxIDTypeRoTin, CustomerUpdateParamsTaxIDTypeRsPib, CustomerUpdateParamsTaxIDTypeRuInn, CustomerUpdateParamsTaxIDTypeRuKpp, CustomerUpdateParamsTaxIDTypeSaVat, CustomerUpdateParamsTaxIDTypeSgGst, CustomerUpdateParamsTaxIDTypeSgUen, CustomerUpdateParamsTaxIDTypeSiTin, CustomerUpdateParamsTaxIDTypeSvNit, CustomerUpdateParamsTaxIDTypeThVat, CustomerUpdateParamsTaxIDTypeTrTin, CustomerUpdateParamsTaxIDTypeTwVat, CustomerUpdateParamsTaxIDTypeUaVat, CustomerUpdateParamsTaxIDTypeUsEin, CustomerUpdateParamsTaxIDTypeUyRuc, CustomerUpdateParamsTaxIDTypeVeRif, CustomerUpdateParamsTaxIDTypeVnTin, CustomerUpdateParamsTaxIDTypeZaVat:
		return true
	}
	return false
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
	Email         param.Field[string] `json:"email" format:"email"`
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
	TaxConfiguration       param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationUnion]  `json:"tax_configuration"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
	// | France               | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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

func (r CustomerUpdateByExternalIDParamsPaymentProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsPaymentProviderQuickbooks, CustomerUpdateByExternalIDParamsPaymentProviderBillCom, CustomerUpdateByExternalIDParamsPaymentProviderStripeCharge, CustomerUpdateByExternalIDParamsPaymentProviderStripeInvoice, CustomerUpdateByExternalIDParamsPaymentProviderNetsuite:
		return true
	}
	return false
}

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

type CustomerUpdateByExternalIDParamsTaxConfiguration struct {
	TaxExempt        param.Field[bool]                                                        `json:"tax_exempt,required"`
	TaxProvider      param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                                      `json:"tax_exemption_code"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

// Satisfied by
// [CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfiguration],
// [CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfiguration],
// [CustomerUpdateByExternalIDParamsTaxConfiguration].
type CustomerUpdateByExternalIDParamsTaxConfigurationUnion interface {
	implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion()
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfiguration struct {
	TaxExempt        param.Field[bool]                                                                                  `json:"tax_exempt,required"`
	TaxProvider      param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                                                                `json:"tax_exemption_code"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider = "avalara"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationNewAvalaraTaxConfigurationTaxProviderAvalara:
		return true
	}
	return false
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfiguration struct {
	TaxExempt   param.Field[bool]                                                                              `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfigurationTaxProvider] `json:"tax_provider,required"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfigurationTaxProviderTaxjar CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfigurationTaxProvider = "taxjar"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationNewTaxJarConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

type CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderAvalara CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "avalara"
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderTaxjar  CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "taxjar"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderAvalara, CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
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
	CustomerUpdateByExternalIDParamsTaxIDCountryAr CustomerUpdateByExternalIDParamsTaxIDCountry = "AR"
	CustomerUpdateByExternalIDParamsTaxIDCountryAt CustomerUpdateByExternalIDParamsTaxIDCountry = "AT"
	CustomerUpdateByExternalIDParamsTaxIDCountryAu CustomerUpdateByExternalIDParamsTaxIDCountry = "AU"
	CustomerUpdateByExternalIDParamsTaxIDCountryBe CustomerUpdateByExternalIDParamsTaxIDCountry = "BE"
	CustomerUpdateByExternalIDParamsTaxIDCountryBg CustomerUpdateByExternalIDParamsTaxIDCountry = "BG"
	CustomerUpdateByExternalIDParamsTaxIDCountryBh CustomerUpdateByExternalIDParamsTaxIDCountry = "BH"
	CustomerUpdateByExternalIDParamsTaxIDCountryBo CustomerUpdateByExternalIDParamsTaxIDCountry = "BO"
	CustomerUpdateByExternalIDParamsTaxIDCountryBr CustomerUpdateByExternalIDParamsTaxIDCountry = "BR"
	CustomerUpdateByExternalIDParamsTaxIDCountryCa CustomerUpdateByExternalIDParamsTaxIDCountry = "CA"
	CustomerUpdateByExternalIDParamsTaxIDCountryCh CustomerUpdateByExternalIDParamsTaxIDCountry = "CH"
	CustomerUpdateByExternalIDParamsTaxIDCountryCl CustomerUpdateByExternalIDParamsTaxIDCountry = "CL"
	CustomerUpdateByExternalIDParamsTaxIDCountryCn CustomerUpdateByExternalIDParamsTaxIDCountry = "CN"
	CustomerUpdateByExternalIDParamsTaxIDCountryCo CustomerUpdateByExternalIDParamsTaxIDCountry = "CO"
	CustomerUpdateByExternalIDParamsTaxIDCountryCr CustomerUpdateByExternalIDParamsTaxIDCountry = "CR"
	CustomerUpdateByExternalIDParamsTaxIDCountryCy CustomerUpdateByExternalIDParamsTaxIDCountry = "CY"
	CustomerUpdateByExternalIDParamsTaxIDCountryCz CustomerUpdateByExternalIDParamsTaxIDCountry = "CZ"
	CustomerUpdateByExternalIDParamsTaxIDCountryDe CustomerUpdateByExternalIDParamsTaxIDCountry = "DE"
	CustomerUpdateByExternalIDParamsTaxIDCountryDk CustomerUpdateByExternalIDParamsTaxIDCountry = "DK"
	CustomerUpdateByExternalIDParamsTaxIDCountryEe CustomerUpdateByExternalIDParamsTaxIDCountry = "EE"
	CustomerUpdateByExternalIDParamsTaxIDCountryDo CustomerUpdateByExternalIDParamsTaxIDCountry = "DO"
	CustomerUpdateByExternalIDParamsTaxIDCountryEc CustomerUpdateByExternalIDParamsTaxIDCountry = "EC"
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
	CustomerUpdateByExternalIDParamsTaxIDCountryKz CustomerUpdateByExternalIDParamsTaxIDCountry = "KZ"
	CustomerUpdateByExternalIDParamsTaxIDCountryLi CustomerUpdateByExternalIDParamsTaxIDCountry = "LI"
	CustomerUpdateByExternalIDParamsTaxIDCountryLt CustomerUpdateByExternalIDParamsTaxIDCountry = "LT"
	CustomerUpdateByExternalIDParamsTaxIDCountryLu CustomerUpdateByExternalIDParamsTaxIDCountry = "LU"
	CustomerUpdateByExternalIDParamsTaxIDCountryLv CustomerUpdateByExternalIDParamsTaxIDCountry = "LV"
	CustomerUpdateByExternalIDParamsTaxIDCountryMt CustomerUpdateByExternalIDParamsTaxIDCountry = "MT"
	CustomerUpdateByExternalIDParamsTaxIDCountryMx CustomerUpdateByExternalIDParamsTaxIDCountry = "MX"
	CustomerUpdateByExternalIDParamsTaxIDCountryMy CustomerUpdateByExternalIDParamsTaxIDCountry = "MY"
	CustomerUpdateByExternalIDParamsTaxIDCountryNg CustomerUpdateByExternalIDParamsTaxIDCountry = "NG"
	CustomerUpdateByExternalIDParamsTaxIDCountryNl CustomerUpdateByExternalIDParamsTaxIDCountry = "NL"
	CustomerUpdateByExternalIDParamsTaxIDCountryNo CustomerUpdateByExternalIDParamsTaxIDCountry = "NO"
	CustomerUpdateByExternalIDParamsTaxIDCountryNz CustomerUpdateByExternalIDParamsTaxIDCountry = "NZ"
	CustomerUpdateByExternalIDParamsTaxIDCountryOm CustomerUpdateByExternalIDParamsTaxIDCountry = "OM"
	CustomerUpdateByExternalIDParamsTaxIDCountryPe CustomerUpdateByExternalIDParamsTaxIDCountry = "PE"
	CustomerUpdateByExternalIDParamsTaxIDCountryPh CustomerUpdateByExternalIDParamsTaxIDCountry = "PH"
	CustomerUpdateByExternalIDParamsTaxIDCountryPl CustomerUpdateByExternalIDParamsTaxIDCountry = "PL"
	CustomerUpdateByExternalIDParamsTaxIDCountryPt CustomerUpdateByExternalIDParamsTaxIDCountry = "PT"
	CustomerUpdateByExternalIDParamsTaxIDCountryRo CustomerUpdateByExternalIDParamsTaxIDCountry = "RO"
	CustomerUpdateByExternalIDParamsTaxIDCountryRs CustomerUpdateByExternalIDParamsTaxIDCountry = "RS"
	CustomerUpdateByExternalIDParamsTaxIDCountryRu CustomerUpdateByExternalIDParamsTaxIDCountry = "RU"
	CustomerUpdateByExternalIDParamsTaxIDCountrySa CustomerUpdateByExternalIDParamsTaxIDCountry = "SA"
	CustomerUpdateByExternalIDParamsTaxIDCountrySe CustomerUpdateByExternalIDParamsTaxIDCountry = "SE"
	CustomerUpdateByExternalIDParamsTaxIDCountrySg CustomerUpdateByExternalIDParamsTaxIDCountry = "SG"
	CustomerUpdateByExternalIDParamsTaxIDCountrySi CustomerUpdateByExternalIDParamsTaxIDCountry = "SI"
	CustomerUpdateByExternalIDParamsTaxIDCountrySk CustomerUpdateByExternalIDParamsTaxIDCountry = "SK"
	CustomerUpdateByExternalIDParamsTaxIDCountrySv CustomerUpdateByExternalIDParamsTaxIDCountry = "SV"
	CustomerUpdateByExternalIDParamsTaxIDCountryTh CustomerUpdateByExternalIDParamsTaxIDCountry = "TH"
	CustomerUpdateByExternalIDParamsTaxIDCountryTr CustomerUpdateByExternalIDParamsTaxIDCountry = "TR"
	CustomerUpdateByExternalIDParamsTaxIDCountryTw CustomerUpdateByExternalIDParamsTaxIDCountry = "TW"
	CustomerUpdateByExternalIDParamsTaxIDCountryUa CustomerUpdateByExternalIDParamsTaxIDCountry = "UA"
	CustomerUpdateByExternalIDParamsTaxIDCountryUs CustomerUpdateByExternalIDParamsTaxIDCountry = "US"
	CustomerUpdateByExternalIDParamsTaxIDCountryUy CustomerUpdateByExternalIDParamsTaxIDCountry = "UY"
	CustomerUpdateByExternalIDParamsTaxIDCountryVe CustomerUpdateByExternalIDParamsTaxIDCountry = "VE"
	CustomerUpdateByExternalIDParamsTaxIDCountryVn CustomerUpdateByExternalIDParamsTaxIDCountry = "VN"
	CustomerUpdateByExternalIDParamsTaxIDCountryZa CustomerUpdateByExternalIDParamsTaxIDCountry = "ZA"
)

func (r CustomerUpdateByExternalIDParamsTaxIDCountry) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxIDCountryAd, CustomerUpdateByExternalIDParamsTaxIDCountryAe, CustomerUpdateByExternalIDParamsTaxIDCountryAr, CustomerUpdateByExternalIDParamsTaxIDCountryAt, CustomerUpdateByExternalIDParamsTaxIDCountryAu, CustomerUpdateByExternalIDParamsTaxIDCountryBe, CustomerUpdateByExternalIDParamsTaxIDCountryBg, CustomerUpdateByExternalIDParamsTaxIDCountryBh, CustomerUpdateByExternalIDParamsTaxIDCountryBo, CustomerUpdateByExternalIDParamsTaxIDCountryBr, CustomerUpdateByExternalIDParamsTaxIDCountryCa, CustomerUpdateByExternalIDParamsTaxIDCountryCh, CustomerUpdateByExternalIDParamsTaxIDCountryCl, CustomerUpdateByExternalIDParamsTaxIDCountryCn, CustomerUpdateByExternalIDParamsTaxIDCountryCo, CustomerUpdateByExternalIDParamsTaxIDCountryCr, CustomerUpdateByExternalIDParamsTaxIDCountryCy, CustomerUpdateByExternalIDParamsTaxIDCountryCz, CustomerUpdateByExternalIDParamsTaxIDCountryDe, CustomerUpdateByExternalIDParamsTaxIDCountryDk, CustomerUpdateByExternalIDParamsTaxIDCountryEe, CustomerUpdateByExternalIDParamsTaxIDCountryDo, CustomerUpdateByExternalIDParamsTaxIDCountryEc, CustomerUpdateByExternalIDParamsTaxIDCountryEg, CustomerUpdateByExternalIDParamsTaxIDCountryEs, CustomerUpdateByExternalIDParamsTaxIDCountryEu, CustomerUpdateByExternalIDParamsTaxIDCountryFi, CustomerUpdateByExternalIDParamsTaxIDCountryFr, CustomerUpdateByExternalIDParamsTaxIDCountryGB, CustomerUpdateByExternalIDParamsTaxIDCountryGe, CustomerUpdateByExternalIDParamsTaxIDCountryGr, CustomerUpdateByExternalIDParamsTaxIDCountryHk, CustomerUpdateByExternalIDParamsTaxIDCountryHr, CustomerUpdateByExternalIDParamsTaxIDCountryHu, CustomerUpdateByExternalIDParamsTaxIDCountryID, CustomerUpdateByExternalIDParamsTaxIDCountryIe, CustomerUpdateByExternalIDParamsTaxIDCountryIl, CustomerUpdateByExternalIDParamsTaxIDCountryIn, CustomerUpdateByExternalIDParamsTaxIDCountryIs, CustomerUpdateByExternalIDParamsTaxIDCountryIt, CustomerUpdateByExternalIDParamsTaxIDCountryJp, CustomerUpdateByExternalIDParamsTaxIDCountryKe, CustomerUpdateByExternalIDParamsTaxIDCountryKr, CustomerUpdateByExternalIDParamsTaxIDCountryKz, CustomerUpdateByExternalIDParamsTaxIDCountryLi, CustomerUpdateByExternalIDParamsTaxIDCountryLt, CustomerUpdateByExternalIDParamsTaxIDCountryLu, CustomerUpdateByExternalIDParamsTaxIDCountryLv, CustomerUpdateByExternalIDParamsTaxIDCountryMt, CustomerUpdateByExternalIDParamsTaxIDCountryMx, CustomerUpdateByExternalIDParamsTaxIDCountryMy, CustomerUpdateByExternalIDParamsTaxIDCountryNg, CustomerUpdateByExternalIDParamsTaxIDCountryNl, CustomerUpdateByExternalIDParamsTaxIDCountryNo, CustomerUpdateByExternalIDParamsTaxIDCountryNz, CustomerUpdateByExternalIDParamsTaxIDCountryOm, CustomerUpdateByExternalIDParamsTaxIDCountryPe, CustomerUpdateByExternalIDParamsTaxIDCountryPh, CustomerUpdateByExternalIDParamsTaxIDCountryPl, CustomerUpdateByExternalIDParamsTaxIDCountryPt, CustomerUpdateByExternalIDParamsTaxIDCountryRo, CustomerUpdateByExternalIDParamsTaxIDCountryRs, CustomerUpdateByExternalIDParamsTaxIDCountryRu, CustomerUpdateByExternalIDParamsTaxIDCountrySa, CustomerUpdateByExternalIDParamsTaxIDCountrySe, CustomerUpdateByExternalIDParamsTaxIDCountrySg, CustomerUpdateByExternalIDParamsTaxIDCountrySi, CustomerUpdateByExternalIDParamsTaxIDCountrySk, CustomerUpdateByExternalIDParamsTaxIDCountrySv, CustomerUpdateByExternalIDParamsTaxIDCountryTh, CustomerUpdateByExternalIDParamsTaxIDCountryTr, CustomerUpdateByExternalIDParamsTaxIDCountryTw, CustomerUpdateByExternalIDParamsTaxIDCountryUa, CustomerUpdateByExternalIDParamsTaxIDCountryUs, CustomerUpdateByExternalIDParamsTaxIDCountryUy, CustomerUpdateByExternalIDParamsTaxIDCountryVe, CustomerUpdateByExternalIDParamsTaxIDCountryVn, CustomerUpdateByExternalIDParamsTaxIDCountryZa:
		return true
	}
	return false
}

type CustomerUpdateByExternalIDParamsTaxIDType string

const (
	CustomerUpdateByExternalIDParamsTaxIDTypeAdNrt    CustomerUpdateByExternalIDParamsTaxIDType = "ad_nrt"
	CustomerUpdateByExternalIDParamsTaxIDTypeAeTrn    CustomerUpdateByExternalIDParamsTaxIDType = "ae_trn"
	CustomerUpdateByExternalIDParamsTaxIDTypeArCuit   CustomerUpdateByExternalIDParamsTaxIDType = "ar_cuit"
	CustomerUpdateByExternalIDParamsTaxIDTypeEuVat    CustomerUpdateByExternalIDParamsTaxIDType = "eu_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeAuAbn    CustomerUpdateByExternalIDParamsTaxIDType = "au_abn"
	CustomerUpdateByExternalIDParamsTaxIDTypeAuArn    CustomerUpdateByExternalIDParamsTaxIDType = "au_arn"
	CustomerUpdateByExternalIDParamsTaxIDTypeBgUic    CustomerUpdateByExternalIDParamsTaxIDType = "bg_uic"
	CustomerUpdateByExternalIDParamsTaxIDTypeBhVat    CustomerUpdateByExternalIDParamsTaxIDType = "bh_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeBoTin    CustomerUpdateByExternalIDParamsTaxIDType = "bo_tin"
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
	CustomerUpdateByExternalIDParamsTaxIDTypeCnTin    CustomerUpdateByExternalIDParamsTaxIDType = "cn_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeCoNit    CustomerUpdateByExternalIDParamsTaxIDType = "co_nit"
	CustomerUpdateByExternalIDParamsTaxIDTypeCrTin    CustomerUpdateByExternalIDParamsTaxIDType = "cr_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeDoRcn    CustomerUpdateByExternalIDParamsTaxIDType = "do_rcn"
	CustomerUpdateByExternalIDParamsTaxIDTypeEcRuc    CustomerUpdateByExternalIDParamsTaxIDType = "ec_ruc"
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
	CustomerUpdateByExternalIDParamsTaxIDTypeKzBin    CustomerUpdateByExternalIDParamsTaxIDType = "kz_bin"
	CustomerUpdateByExternalIDParamsTaxIDTypeLiUid    CustomerUpdateByExternalIDParamsTaxIDType = "li_uid"
	CustomerUpdateByExternalIDParamsTaxIDTypeMxRfc    CustomerUpdateByExternalIDParamsTaxIDType = "mx_rfc"
	CustomerUpdateByExternalIDParamsTaxIDTypeMyFrp    CustomerUpdateByExternalIDParamsTaxIDType = "my_frp"
	CustomerUpdateByExternalIDParamsTaxIDTypeMyItn    CustomerUpdateByExternalIDParamsTaxIDType = "my_itn"
	CustomerUpdateByExternalIDParamsTaxIDTypeMySst    CustomerUpdateByExternalIDParamsTaxIDType = "my_sst"
	CustomerUpdateByExternalIDParamsTaxIDTypeNgTin    CustomerUpdateByExternalIDParamsTaxIDType = "ng_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeNoVat    CustomerUpdateByExternalIDParamsTaxIDType = "no_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeNoVoec   CustomerUpdateByExternalIDParamsTaxIDType = "no_voec"
	CustomerUpdateByExternalIDParamsTaxIDTypeNzGst    CustomerUpdateByExternalIDParamsTaxIDType = "nz_gst"
	CustomerUpdateByExternalIDParamsTaxIDTypeOmVat    CustomerUpdateByExternalIDParamsTaxIDType = "om_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypePeRuc    CustomerUpdateByExternalIDParamsTaxIDType = "pe_ruc"
	CustomerUpdateByExternalIDParamsTaxIDTypePhTin    CustomerUpdateByExternalIDParamsTaxIDType = "ph_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeRoTin    CustomerUpdateByExternalIDParamsTaxIDType = "ro_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeRsPib    CustomerUpdateByExternalIDParamsTaxIDType = "rs_pib"
	CustomerUpdateByExternalIDParamsTaxIDTypeRuInn    CustomerUpdateByExternalIDParamsTaxIDType = "ru_inn"
	CustomerUpdateByExternalIDParamsTaxIDTypeRuKpp    CustomerUpdateByExternalIDParamsTaxIDType = "ru_kpp"
	CustomerUpdateByExternalIDParamsTaxIDTypeSaVat    CustomerUpdateByExternalIDParamsTaxIDType = "sa_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeSgGst    CustomerUpdateByExternalIDParamsTaxIDType = "sg_gst"
	CustomerUpdateByExternalIDParamsTaxIDTypeSgUen    CustomerUpdateByExternalIDParamsTaxIDType = "sg_uen"
	CustomerUpdateByExternalIDParamsTaxIDTypeSiTin    CustomerUpdateByExternalIDParamsTaxIDType = "si_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeSvNit    CustomerUpdateByExternalIDParamsTaxIDType = "sv_nit"
	CustomerUpdateByExternalIDParamsTaxIDTypeThVat    CustomerUpdateByExternalIDParamsTaxIDType = "th_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeTrTin    CustomerUpdateByExternalIDParamsTaxIDType = "tr_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeTwVat    CustomerUpdateByExternalIDParamsTaxIDType = "tw_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeUaVat    CustomerUpdateByExternalIDParamsTaxIDType = "ua_vat"
	CustomerUpdateByExternalIDParamsTaxIDTypeUsEin    CustomerUpdateByExternalIDParamsTaxIDType = "us_ein"
	CustomerUpdateByExternalIDParamsTaxIDTypeUyRuc    CustomerUpdateByExternalIDParamsTaxIDType = "uy_ruc"
	CustomerUpdateByExternalIDParamsTaxIDTypeVeRif    CustomerUpdateByExternalIDParamsTaxIDType = "ve_rif"
	CustomerUpdateByExternalIDParamsTaxIDTypeVnTin    CustomerUpdateByExternalIDParamsTaxIDType = "vn_tin"
	CustomerUpdateByExternalIDParamsTaxIDTypeZaVat    CustomerUpdateByExternalIDParamsTaxIDType = "za_vat"
)

func (r CustomerUpdateByExternalIDParamsTaxIDType) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxIDTypeAdNrt, CustomerUpdateByExternalIDParamsTaxIDTypeAeTrn, CustomerUpdateByExternalIDParamsTaxIDTypeArCuit, CustomerUpdateByExternalIDParamsTaxIDTypeEuVat, CustomerUpdateByExternalIDParamsTaxIDTypeAuAbn, CustomerUpdateByExternalIDParamsTaxIDTypeAuArn, CustomerUpdateByExternalIDParamsTaxIDTypeBgUic, CustomerUpdateByExternalIDParamsTaxIDTypeBhVat, CustomerUpdateByExternalIDParamsTaxIDTypeBoTin, CustomerUpdateByExternalIDParamsTaxIDTypeBrCnpj, CustomerUpdateByExternalIDParamsTaxIDTypeBrCpf, CustomerUpdateByExternalIDParamsTaxIDTypeCaBn, CustomerUpdateByExternalIDParamsTaxIDTypeCaGstHst, CustomerUpdateByExternalIDParamsTaxIDTypeCaPstBc, CustomerUpdateByExternalIDParamsTaxIDTypeCaPstMB, CustomerUpdateByExternalIDParamsTaxIDTypeCaPstSk, CustomerUpdateByExternalIDParamsTaxIDTypeCaQst, CustomerUpdateByExternalIDParamsTaxIDTypeChVat, CustomerUpdateByExternalIDParamsTaxIDTypeClTin, CustomerUpdateByExternalIDParamsTaxIDTypeCnTin, CustomerUpdateByExternalIDParamsTaxIDTypeCoNit, CustomerUpdateByExternalIDParamsTaxIDTypeCrTin, CustomerUpdateByExternalIDParamsTaxIDTypeDoRcn, CustomerUpdateByExternalIDParamsTaxIDTypeEcRuc, CustomerUpdateByExternalIDParamsTaxIDTypeEgTin, CustomerUpdateByExternalIDParamsTaxIDTypeEsCif, CustomerUpdateByExternalIDParamsTaxIDTypeEuOssVat, CustomerUpdateByExternalIDParamsTaxIDTypeGBVat, CustomerUpdateByExternalIDParamsTaxIDTypeGeVat, CustomerUpdateByExternalIDParamsTaxIDTypeHkBr, CustomerUpdateByExternalIDParamsTaxIDTypeHuTin, CustomerUpdateByExternalIDParamsTaxIDTypeIDNpwp, CustomerUpdateByExternalIDParamsTaxIDTypeIlVat, CustomerUpdateByExternalIDParamsTaxIDTypeInGst, CustomerUpdateByExternalIDParamsTaxIDTypeIsVat, CustomerUpdateByExternalIDParamsTaxIDTypeJpCn, CustomerUpdateByExternalIDParamsTaxIDTypeJpRn, CustomerUpdateByExternalIDParamsTaxIDTypeJpTrn, CustomerUpdateByExternalIDParamsTaxIDTypeKePin, CustomerUpdateByExternalIDParamsTaxIDTypeKrBrn, CustomerUpdateByExternalIDParamsTaxIDTypeKzBin, CustomerUpdateByExternalIDParamsTaxIDTypeLiUid, CustomerUpdateByExternalIDParamsTaxIDTypeMxRfc, CustomerUpdateByExternalIDParamsTaxIDTypeMyFrp, CustomerUpdateByExternalIDParamsTaxIDTypeMyItn, CustomerUpdateByExternalIDParamsTaxIDTypeMySst, CustomerUpdateByExternalIDParamsTaxIDTypeNgTin, CustomerUpdateByExternalIDParamsTaxIDTypeNoVat, CustomerUpdateByExternalIDParamsTaxIDTypeNoVoec, CustomerUpdateByExternalIDParamsTaxIDTypeNzGst, CustomerUpdateByExternalIDParamsTaxIDTypeOmVat, CustomerUpdateByExternalIDParamsTaxIDTypePeRuc, CustomerUpdateByExternalIDParamsTaxIDTypePhTin, CustomerUpdateByExternalIDParamsTaxIDTypeRoTin, CustomerUpdateByExternalIDParamsTaxIDTypeRsPib, CustomerUpdateByExternalIDParamsTaxIDTypeRuInn, CustomerUpdateByExternalIDParamsTaxIDTypeRuKpp, CustomerUpdateByExternalIDParamsTaxIDTypeSaVat, CustomerUpdateByExternalIDParamsTaxIDTypeSgGst, CustomerUpdateByExternalIDParamsTaxIDTypeSgUen, CustomerUpdateByExternalIDParamsTaxIDTypeSiTin, CustomerUpdateByExternalIDParamsTaxIDTypeSvNit, CustomerUpdateByExternalIDParamsTaxIDTypeThVat, CustomerUpdateByExternalIDParamsTaxIDTypeTrTin, CustomerUpdateByExternalIDParamsTaxIDTypeTwVat, CustomerUpdateByExternalIDParamsTaxIDTypeUaVat, CustomerUpdateByExternalIDParamsTaxIDTypeUsEin, CustomerUpdateByExternalIDParamsTaxIDTypeUyRuc, CustomerUpdateByExternalIDParamsTaxIDTypeVeRif, CustomerUpdateByExternalIDParamsTaxIDTypeVnTin, CustomerUpdateByExternalIDParamsTaxIDTypeZaVat:
		return true
	}
	return false
}
