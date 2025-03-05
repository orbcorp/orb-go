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
// billing relationship. See [Customer](/core-concepts##customer) for an overview
// of the customer resource.
//
// This endpoint is critical in the following Orb functionality:
//
//   - Automated charges can be configured by setting `payment_provider` and
//     `payment_provider_id` to automatically issue invoices
//   - [Customer ID Aliases](/events-and-metrics/customer-aliases) can be configured
//     by setting `external_customer_id`
//   - [Timezone localization](/essentials/timezones) can be configured on a
//     per-customer basis by setting the `timezone` parameter
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *shared.CustomerModel, err error) {
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
func (r *CustomerService) Update(ctx context.Context, customerID string, body CustomerUpdateParams, opts ...option.RequestOption) (res *shared.CustomerModel, err error) {
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
// [standardized pagination format](/api-reference/pagination).
//
// See [Customer](/core-concepts##customer) for an overview of the customer model.
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CustomerModel], err error) {
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
// [standardized pagination format](/api-reference/pagination).
//
// See [Customer](/core-concepts##customer) for an overview of the customer model.
func (r *CustomerService) ListAutoPaging(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CustomerModel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This performs a deletion of this customer, its subscriptions, and its invoices,
// provided the customer does not have any issued invoices. Customers with issued
// invoices cannot be deleted. This operation is irreversible. Note that this is a
// _soft_ deletion, but the data will be inaccessible through the API and Orb
// dashboard.
//
// For a hard-deletion, please reach out to the Orb team directly.
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
// See the [Customer resource](/core-concepts#customer) for a full discussion of
// the Customer model.
func (r *CustomerService) Fetch(ctx context.Context, customerID string, opts ...option.RequestOption) (res *shared.CustomerModel, err error) {
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
// (see [Customer ID Aliases](/events-and-metrics/customer-aliases)).
//
// Note that the resource and semantics of this endpoint exactly mirror
// [Get Customer](fetch-customer).
func (r *CustomerService) FetchByExternalID(ctx context.Context, externalCustomerID string, opts ...option.RequestOption) (res *shared.CustomerModel, err error) {
	opts = append(r.Options[:], opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sync Orb's payment methods for the customer with their gateway.
//
// This method can be called before taking an action that may cause the customer to
// be charged, ensuring that the most up-to-date payment method is charged.
//
// **Note**: This functionality is currently only available for Stripe.
func (r *CustomerService) SyncPaymentMethodsFromGateway(ctx context.Context, externalCustomerID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/sync_payment_methods_from_gateway", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Sync Orb's payment methods for the customer with their gateway.
//
// This method can be called before taking an action that may cause the customer to
// be charged, ensuring that the most up-to-date payment method is charged.
//
// **Note**: This functionality is currently only available for Stripe.
func (r *CustomerService) SyncPaymentMethodsFromGatewayByExternalCustomerID(ctx context.Context, customerID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/sync_payment_methods_from_gateway", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This endpoint is used to update customer details given an `external_customer_id`
// (see [Customer ID Aliases](/events-and-metrics/customer-aliases)). Note that the
// resource and semantics of this endpoint exactly mirror
// [Update Customer](update-customer).
func (r *CustomerService) UpdateByExternalID(ctx context.Context, id string, body CustomerUpdateByExternalIDParams, opts ...option.RequestOption) (res *shared.CustomerModel, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CustomerNewParams struct {
	// A valid customer email, to be used for notifications. When Orb triggers payment
	// through a payment gateway, this email will be used for any automatically issued
	// receipts.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The full name of the customer
	Name                        param.Field[string]                                          `json:"name,required"`
	AccountingSyncConfiguration param.Field[shared.NewAccountingSyncConfigurationModelParam] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool]                          `json:"auto_collection"`
	BillingAddress param.Field[shared.AddressInputModelParam] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency      param.Field[string] `json:"currency"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// The hierarchical relationships for this customer.
	Hierarchy param.Field[shared.CustomerHierarchyConfigModelParam] `json:"hierarchy"`
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
	PaymentProviderID      param.Field[string]                                     `json:"payment_provider_id"`
	ReportingConfiguration param.Field[shared.NewReportingConfigurationModelParam] `json:"reporting_configuration"`
	ShippingAddress        param.Field[shared.AddressInputModelParam]              `json:"shipping_address"`
	TaxConfiguration       param.Field[shared.NewTaxConfigurationModelUnionParam]  `json:"tax_configuration"`
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
	TaxID param.Field[shared.CustomerTaxIDModelParam] `json:"tax_id"`
	// A timezone identifier from the IANA timezone database, such as
	// `"America/Los_Angeles"`. This defaults to your account's timezone if not set.
	// This cannot be changed after customer creation.
	Timezone param.Field[string] `json:"timezone"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
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

type CustomerUpdateParams struct {
	AccountingSyncConfiguration param.Field[shared.NewAccountingSyncConfigurationModelParam] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool]                          `json:"auto_collection"`
	BillingAddress param.Field[shared.AddressInputModelParam] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency param.Field[string] `json:"currency"`
	// A valid customer email, to be used for invoicing and notifications.
	Email         param.Field[string] `json:"email" format:"email"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// The external customer ID. This can only be set if empty and the customer has no
	// past or current subscriptions.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// The hierarchical relationships for this customer.
	Hierarchy param.Field[shared.CustomerHierarchyConfigModelParam] `json:"hierarchy"`
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
	ReportingConfiguration param.Field[shared.NewReportingConfigurationModelParam] `json:"reporting_configuration"`
	ShippingAddress        param.Field[shared.AddressInputModelParam]              `json:"shipping_address"`
	TaxConfiguration       param.Field[shared.NewTaxConfigurationModelUnionParam]  `json:"tax_configuration"`
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
	TaxID param.Field[shared.CustomerTaxIDModelParam] `json:"tax_id"`
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
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
	AccountingSyncConfiguration param.Field[shared.NewAccountingSyncConfigurationModelParam] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool]                          `json:"auto_collection"`
	BillingAddress param.Field[shared.AddressInputModelParam] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency param.Field[string] `json:"currency"`
	// A valid customer email, to be used for invoicing and notifications.
	Email         param.Field[string] `json:"email" format:"email"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// The external customer ID. This can only be set if empty and the customer has no
	// past or current subscriptions.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// The hierarchical relationships for this customer.
	Hierarchy param.Field[shared.CustomerHierarchyConfigModelParam] `json:"hierarchy"`
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
	PaymentProviderID      param.Field[string]                                     `json:"payment_provider_id"`
	ReportingConfiguration param.Field[shared.NewReportingConfigurationModelParam] `json:"reporting_configuration"`
	ShippingAddress        param.Field[shared.AddressInputModelParam]              `json:"shipping_address"`
	TaxConfiguration       param.Field[shared.NewTaxConfigurationModelUnionParam]  `json:"tax_configuration"`
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
	TaxID param.Field[shared.CustomerTaxIDModelParam] `json:"tax_id"`
}

func (r CustomerUpdateByExternalIDParams) MarshalJSON() (data []byte, err error) {
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
