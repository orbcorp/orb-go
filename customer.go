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
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *pagination.Page[Customer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *CustomerService) ListAutoPaging(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Customer] {
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
// its deletion.
func (r *CustomerService) Delete(ctx context.Context, customerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
func (r *CustomerService) Fetch(ctx context.Context, customerID string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *CustomerService) FetchByExternalID(ctx context.Context, externalCustomerID string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *CustomerService) SyncPaymentMethodsFromGateway(ctx context.Context, customerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/sync_payment_methods_from_gateway", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Sync Orb's payment methods for the customer with their gateway.
//
// This method can be called before taking an action that may cause the customer to
// be charged, ensuring that the most up-to-date payment method is charged.
//
// **Note**: This functionality is currently only available for Stripe.
func (r *CustomerService) SyncPaymentMethodsFromGatewayByExternalCustomerID(ctx context.Context, externalCustomerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/sync_payment_methods_from_gateway", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This endpoint is used to update customer details given an `external_customer_id`
// (see [Customer ID Aliases](/events-and-metrics/customer-aliases)). Note that the
// resource and semantics of this endpoint exactly mirror
// [Update Customer](update-customer).
func (r *CustomerService) UpdateByExternalID(ctx context.Context, id string, body CustomerUpdateByExternalIDParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountingProviderConfigParam struct {
	ExternalProviderID param.Field[string] `json:"external_provider_id,required"`
	ProviderType       param.Field[string] `json:"provider_type,required"`
}

func (r AccountingProviderConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressInputParam struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r AddressInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
type Customer struct {
	ID               string   `json:"id,required"`
	AdditionalEmails []string `json:"additional_emails,required"`
	AutoCollection   bool     `json:"auto_collection,required"`
	// Whether invoices for this customer should be automatically issued. If true,
	// invoices will be automatically issued. If false, invoices will require manual
	// approval. If null, inherits the account-level setting.
	AutoIssuance bool `json:"auto_issuance,required,nullable"`
	// The customer's current balance in their currency.
	Balance        string         `json:"balance,required"`
	BillingAddress shared.Address `json:"billing_address,required,nullable"`
	CreatedAt      time.Time      `json:"created_at,required" format:"date-time"`
	Currency       string         `json:"currency,required,nullable"`
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
	// The hierarchical relationships for this customer.
	Hierarchy CustomerHierarchy `json:"hierarchy,required"`
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
	PaymentProviderID string         `json:"payment_provider_id,required,nullable"`
	PortalURL         string         `json:"portal_url,required,nullable"`
	ShippingAddress   shared.Address `json:"shipping_address,required,nullable"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country                | Type         | Description                                                                                             |
	// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
	// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
	// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
	// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
	// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
	// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
	// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
	// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
	// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
	// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
	// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
	// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
	// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
	// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
	// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
	// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
	// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
	// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
	// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
	// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
	// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                  | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
	// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
	// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
	// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
	// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
	// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
	// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
	// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
	// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
	// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
	// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
	// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
	// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
	// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
	// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
	// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
	// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States          | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
	// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
	// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
	// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
	TaxID shared.CustomerTaxID `json:"tax_id,required,nullable"`
	// A timezone identifier from the IANA timezone database, such as
	// "America/Los_Angeles". This "defaults to your account's timezone if not set.
	// This cannot be changed after customer creation.
	Timezone                    string                              `json:"timezone,required"`
	AccountingSyncConfiguration CustomerAccountingSyncConfiguration `json:"accounting_sync_configuration,nullable"`
	// Whether automatic tax calculation is enabled for this customer. This field is
	// nullable for backwards compatibility but will always return a boolean value.
	AutomaticTaxEnabled    bool                           `json:"automatic_tax_enabled,nullable"`
	ReportingConfiguration CustomerReportingConfiguration `json:"reporting_configuration,nullable"`
	JSON                   customerJSON                   `json:"-"`
}

// customerJSON contains the JSON metadata for the struct [Customer]
type customerJSON struct {
	ID                          apijson.Field
	AdditionalEmails            apijson.Field
	AutoCollection              apijson.Field
	AutoIssuance                apijson.Field
	Balance                     apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	Email                       apijson.Field
	EmailDelivery               apijson.Field
	ExemptFromAutomatedTax      apijson.Field
	ExternalCustomerID          apijson.Field
	Hierarchy                   apijson.Field
	Metadata                    apijson.Field
	Name                        apijson.Field
	PaymentProvider             apijson.Field
	PaymentProviderID           apijson.Field
	PortalURL                   apijson.Field
	ShippingAddress             apijson.Field
	TaxID                       apijson.Field
	Timezone                    apijson.Field
	AccountingSyncConfiguration apijson.Field
	AutomaticTaxEnabled         apijson.Field
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

// The hierarchical relationships for this customer.
type CustomerHierarchy struct {
	Children []shared.CustomerMinified `json:"children,required"`
	Parent   shared.CustomerMinified   `json:"parent,required,nullable"`
	JSON     customerHierarchyJSON     `json:"-"`
}

// customerHierarchyJSON contains the JSON metadata for the struct
// [CustomerHierarchy]
type customerHierarchyJSON struct {
	Children    apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerHierarchy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerHierarchyJSON) RawJSON() string {
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

type CustomerHierarchyConfigParam struct {
	// A list of child customer IDs to add to the hierarchy. The desired child
	// customers must not already be part of another hierarchy.
	ChildCustomerIDs param.Field[[]string] `json:"child_customer_ids"`
	// The ID of the parent customer in the hierarchy. The desired parent customer must
	// not be a child of another customer.
	ParentCustomerID param.Field[string] `json:"parent_customer_id"`
}

func (r CustomerHierarchyConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewAccountingSyncConfigurationParam struct {
	AccountingProviders param.Field[[]AccountingProviderConfigParam] `json:"accounting_providers"`
	Excluded            param.Field[bool]                            `json:"excluded"`
}

func (r NewAccountingSyncConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewAvalaraTaxConfigurationParam struct {
	TaxExempt   param.Field[bool]                                  `json:"tax_exempt,required"`
	TaxProvider param.Field[NewAvalaraTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool]   `json:"automatic_tax_enabled"`
	TaxExemptionCode    param.Field[string] `json:"tax_exemption_code"`
}

func (r NewAvalaraTaxConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAvalaraTaxConfigurationParam) implementsCustomerNewParamsTaxConfigurationUnion() {}

func (r NewAvalaraTaxConfigurationParam) implementsCustomerUpdateParamsTaxConfigurationUnion() {}

func (r NewAvalaraTaxConfigurationParam) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type NewAvalaraTaxConfigurationTaxProvider string

const (
	NewAvalaraTaxConfigurationTaxProviderAvalara NewAvalaraTaxConfigurationTaxProvider = "avalara"
)

func (r NewAvalaraTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case NewAvalaraTaxConfigurationTaxProviderAvalara:
		return true
	}
	return false
}

type NewReportingConfigurationParam struct {
	Exempt param.Field[bool] `json:"exempt,required"`
}

func (r NewReportingConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSphereConfigurationParam struct {
	TaxExempt   param.Field[bool]                              `json:"tax_exempt,required"`
	TaxProvider param.Field[NewSphereConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r NewSphereConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSphereConfigurationParam) implementsCustomerNewParamsTaxConfigurationUnion() {}

func (r NewSphereConfigurationParam) implementsCustomerUpdateParamsTaxConfigurationUnion() {}

func (r NewSphereConfigurationParam) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type NewSphereConfigurationTaxProvider string

const (
	NewSphereConfigurationTaxProviderSphere NewSphereConfigurationTaxProvider = "sphere"
)

func (r NewSphereConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case NewSphereConfigurationTaxProviderSphere:
		return true
	}
	return false
}

type NewTaxJarConfigurationParam struct {
	TaxExempt   param.Field[bool]                              `json:"tax_exempt,required"`
	TaxProvider param.Field[NewTaxJarConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r NewTaxJarConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewTaxJarConfigurationParam) implementsCustomerNewParamsTaxConfigurationUnion() {}

func (r NewTaxJarConfigurationParam) implementsCustomerUpdateParamsTaxConfigurationUnion() {}

func (r NewTaxJarConfigurationParam) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type NewTaxJarConfigurationTaxProvider string

const (
	NewTaxJarConfigurationTaxProviderTaxjar NewTaxJarConfigurationTaxProvider = "taxjar"
)

func (r NewTaxJarConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case NewTaxJarConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

type CustomerNewParams struct {
	// A valid customer email, to be used for notifications. When Orb triggers payment
	// through a payment gateway, this email will be used for any automatically issued
	// receipts.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The full name of the customer
	Name                        param.Field[string]                              `json:"name,required"`
	AccountingSyncConfiguration param.Field[NewAccountingSyncConfigurationParam] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications. The total number of email
	// addresses (including the primary email) cannot exceed 50.
	AdditionalEmails param.Field[[]string] `json:"additional_emails" format:"email"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Used to determine if invoices for this customer will be automatically issued. If
	// true, invoices will be automatically issued. If false, invoices will require
	// manual approval. If `null` is specified, the customer's auto issuance setting
	// will be inherited from the account-level setting.
	AutoIssuance   param.Field[bool]              `json:"auto_issuance"`
	BillingAddress param.Field[AddressInputParam] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency      param.Field[string] `json:"currency"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// The hierarchical relationships for this customer.
	Hierarchy param.Field[CustomerHierarchyConfigParam] `json:"hierarchy"`
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
	PaymentProviderID      param.Field[string]                                 `json:"payment_provider_id"`
	ReportingConfiguration param.Field[NewReportingConfigurationParam]         `json:"reporting_configuration"`
	ShippingAddress        param.Field[AddressInputParam]                      `json:"shipping_address"`
	TaxConfiguration       param.Field[CustomerNewParamsTaxConfigurationUnion] `json:"tax_configuration"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country                | Type         | Description                                                                                             |
	// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
	// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
	// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
	// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
	// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
	// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
	// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
	// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
	// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
	// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
	// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
	// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
	// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
	// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
	// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
	// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
	// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
	// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
	// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
	// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
	// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                  | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
	// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
	// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
	// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
	// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
	// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
	// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
	// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
	// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
	// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
	// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
	// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
	// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
	// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
	// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
	// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
	// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States          | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
	// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
	// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
	// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
	TaxID param.Field[shared.CustomerTaxIDParam] `json:"tax_id"`
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

type CustomerNewParamsTaxConfiguration struct {
	TaxExempt   param.Field[bool]                                         `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerNewParamsTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool]   `json:"automatic_tax_enabled"`
	TaxExemptionCode    param.Field[string] `json:"tax_exemption_code"`
}

func (r CustomerNewParamsTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {}

// Satisfied by [NewAvalaraTaxConfigurationParam], [NewTaxJarConfigurationParam],
// [NewSphereConfigurationParam],
// [CustomerNewParamsTaxConfigurationNewNumeralConfiguration],
// [CustomerNewParamsTaxConfigurationNewAnrokConfiguration],
// [CustomerNewParamsTaxConfigurationNewStripeTaxConfiguration],
// [CustomerNewParamsTaxConfiguration].
type CustomerNewParamsTaxConfigurationUnion interface {
	implementsCustomerNewParamsTaxConfigurationUnion()
}

type CustomerNewParamsTaxConfigurationNewNumeralConfiguration struct {
	TaxExempt   param.Field[bool]                                                                `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerNewParamsTaxConfigurationNewNumeralConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerNewParamsTaxConfigurationNewNumeralConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfigurationNewNumeralConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {
}

type CustomerNewParamsTaxConfigurationNewNumeralConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationNewNumeralConfigurationTaxProviderNumeral CustomerNewParamsTaxConfigurationNewNumeralConfigurationTaxProvider = "numeral"
)

func (r CustomerNewParamsTaxConfigurationNewNumeralConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationNewNumeralConfigurationTaxProviderNumeral:
		return true
	}
	return false
}

type CustomerNewParamsTaxConfigurationNewAnrokConfiguration struct {
	TaxExempt   param.Field[bool]                                                              `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerNewParamsTaxConfigurationNewAnrokConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerNewParamsTaxConfigurationNewAnrokConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfigurationNewAnrokConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {
}

type CustomerNewParamsTaxConfigurationNewAnrokConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationNewAnrokConfigurationTaxProviderAnrok CustomerNewParamsTaxConfigurationNewAnrokConfigurationTaxProvider = "anrok"
)

func (r CustomerNewParamsTaxConfigurationNewAnrokConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationNewAnrokConfigurationTaxProviderAnrok:
		return true
	}
	return false
}

type CustomerNewParamsTaxConfigurationNewStripeTaxConfiguration struct {
	TaxExempt   param.Field[bool]                                                                  `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerNewParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerNewParamsTaxConfigurationNewStripeTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerNewParamsTaxConfigurationNewStripeTaxConfiguration) implementsCustomerNewParamsTaxConfigurationUnion() {
}

type CustomerNewParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationNewStripeTaxConfigurationTaxProviderStripe CustomerNewParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider = "stripe"
)

func (r CustomerNewParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationNewStripeTaxConfigurationTaxProviderStripe:
		return true
	}
	return false
}

type CustomerNewParamsTaxConfigurationTaxProvider string

const (
	CustomerNewParamsTaxConfigurationTaxProviderAvalara CustomerNewParamsTaxConfigurationTaxProvider = "avalara"
	CustomerNewParamsTaxConfigurationTaxProviderTaxjar  CustomerNewParamsTaxConfigurationTaxProvider = "taxjar"
	CustomerNewParamsTaxConfigurationTaxProviderSphere  CustomerNewParamsTaxConfigurationTaxProvider = "sphere"
	CustomerNewParamsTaxConfigurationTaxProviderNumeral CustomerNewParamsTaxConfigurationTaxProvider = "numeral"
	CustomerNewParamsTaxConfigurationTaxProviderAnrok   CustomerNewParamsTaxConfigurationTaxProvider = "anrok"
	CustomerNewParamsTaxConfigurationTaxProviderStripe  CustomerNewParamsTaxConfigurationTaxProvider = "stripe"
)

func (r CustomerNewParamsTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerNewParamsTaxConfigurationTaxProviderAvalara, CustomerNewParamsTaxConfigurationTaxProviderTaxjar, CustomerNewParamsTaxConfigurationTaxProviderSphere, CustomerNewParamsTaxConfigurationTaxProviderNumeral, CustomerNewParamsTaxConfigurationTaxProviderAnrok, CustomerNewParamsTaxConfigurationTaxProviderStripe:
		return true
	}
	return false
}

type CustomerUpdateParams struct {
	AccountingSyncConfiguration param.Field[NewAccountingSyncConfigurationParam] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications. The total number of email
	// addresses (including the primary email) cannot exceed 50.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Used to determine if invoices for this customer will be automatically issued. If
	// true, invoices will be automatically issued. If false, invoices will require
	// manual approval.If `null` is specified, the customer's auto issuance setting
	// will be inherited from the account-level setting.
	AutoIssuance   param.Field[bool]              `json:"auto_issuance"`
	BillingAddress param.Field[AddressInputParam] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency param.Field[string] `json:"currency"`
	// A valid customer email, to be used for invoicing and notifications.
	Email         param.Field[string] `json:"email" format:"email"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// The external customer ID. This can only be set if the customer has no existing
	// external customer ID. Since this action may change usage quantities for all
	// existing subscriptions, it is disallowed if the customer has issued invoices
	// with usage line items and subject to the same restrictions as backdated
	// subscription creation.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// The hierarchical relationships for this customer.
	Hierarchy param.Field[CustomerHierarchyConfigParam] `json:"hierarchy"`
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
	PaymentProviderID      param.Field[string]                                    `json:"payment_provider_id"`
	ReportingConfiguration param.Field[NewReportingConfigurationParam]            `json:"reporting_configuration"`
	ShippingAddress        param.Field[AddressInputParam]                         `json:"shipping_address"`
	TaxConfiguration       param.Field[CustomerUpdateParamsTaxConfigurationUnion] `json:"tax_configuration"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country                | Type         | Description                                                                                             |
	// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
	// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
	// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
	// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
	// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
	// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
	// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
	// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
	// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
	// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
	// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
	// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
	// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
	// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
	// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
	// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
	// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
	// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
	// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
	// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
	// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                  | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
	// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
	// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
	// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
	// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
	// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
	// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
	// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
	// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
	// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
	// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
	// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
	// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
	// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
	// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
	// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
	// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States          | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
	// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
	// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
	// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
	TaxID param.Field[shared.CustomerTaxIDParam] `json:"tax_id"`
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

type CustomerUpdateParamsTaxConfiguration struct {
	TaxExempt   param.Field[bool]                                            `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateParamsTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool]   `json:"automatic_tax_enabled"`
	TaxExemptionCode    param.Field[string] `json:"tax_exemption_code"`
}

func (r CustomerUpdateParamsTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {}

// Satisfied by [NewAvalaraTaxConfigurationParam], [NewTaxJarConfigurationParam],
// [NewSphereConfigurationParam],
// [CustomerUpdateParamsTaxConfigurationNewNumeralConfiguration],
// [CustomerUpdateParamsTaxConfigurationNewAnrokConfiguration],
// [CustomerUpdateParamsTaxConfigurationNewStripeTaxConfiguration],
// [CustomerUpdateParamsTaxConfiguration].
type CustomerUpdateParamsTaxConfigurationUnion interface {
	implementsCustomerUpdateParamsTaxConfigurationUnion()
}

type CustomerUpdateParamsTaxConfigurationNewNumeralConfiguration struct {
	TaxExempt   param.Field[bool]                                                                   `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateParamsTaxConfigurationNewNumeralConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerUpdateParamsTaxConfigurationNewNumeralConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfigurationNewNumeralConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {
}

type CustomerUpdateParamsTaxConfigurationNewNumeralConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationNewNumeralConfigurationTaxProviderNumeral CustomerUpdateParamsTaxConfigurationNewNumeralConfigurationTaxProvider = "numeral"
)

func (r CustomerUpdateParamsTaxConfigurationNewNumeralConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationNewNumeralConfigurationTaxProviderNumeral:
		return true
	}
	return false
}

type CustomerUpdateParamsTaxConfigurationNewAnrokConfiguration struct {
	TaxExempt   param.Field[bool]                                                                 `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateParamsTaxConfigurationNewAnrokConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerUpdateParamsTaxConfigurationNewAnrokConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfigurationNewAnrokConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {
}

type CustomerUpdateParamsTaxConfigurationNewAnrokConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationNewAnrokConfigurationTaxProviderAnrok CustomerUpdateParamsTaxConfigurationNewAnrokConfigurationTaxProvider = "anrok"
)

func (r CustomerUpdateParamsTaxConfigurationNewAnrokConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationNewAnrokConfigurationTaxProviderAnrok:
		return true
	}
	return false
}

type CustomerUpdateParamsTaxConfigurationNewStripeTaxConfiguration struct {
	TaxExempt   param.Field[bool]                                                                     `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerUpdateParamsTaxConfigurationNewStripeTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateParamsTaxConfigurationNewStripeTaxConfiguration) implementsCustomerUpdateParamsTaxConfigurationUnion() {
}

type CustomerUpdateParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationNewStripeTaxConfigurationTaxProviderStripe CustomerUpdateParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider = "stripe"
)

func (r CustomerUpdateParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationNewStripeTaxConfigurationTaxProviderStripe:
		return true
	}
	return false
}

type CustomerUpdateParamsTaxConfigurationTaxProvider string

const (
	CustomerUpdateParamsTaxConfigurationTaxProviderAvalara CustomerUpdateParamsTaxConfigurationTaxProvider = "avalara"
	CustomerUpdateParamsTaxConfigurationTaxProviderTaxjar  CustomerUpdateParamsTaxConfigurationTaxProvider = "taxjar"
	CustomerUpdateParamsTaxConfigurationTaxProviderSphere  CustomerUpdateParamsTaxConfigurationTaxProvider = "sphere"
	CustomerUpdateParamsTaxConfigurationTaxProviderNumeral CustomerUpdateParamsTaxConfigurationTaxProvider = "numeral"
	CustomerUpdateParamsTaxConfigurationTaxProviderAnrok   CustomerUpdateParamsTaxConfigurationTaxProvider = "anrok"
	CustomerUpdateParamsTaxConfigurationTaxProviderStripe  CustomerUpdateParamsTaxConfigurationTaxProvider = "stripe"
)

func (r CustomerUpdateParamsTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsTaxConfigurationTaxProviderAvalara, CustomerUpdateParamsTaxConfigurationTaxProviderTaxjar, CustomerUpdateParamsTaxConfigurationTaxProviderSphere, CustomerUpdateParamsTaxConfigurationTaxProviderNumeral, CustomerUpdateParamsTaxConfigurationTaxProviderAnrok, CustomerUpdateParamsTaxConfigurationTaxProviderStripe:
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
	AccountingSyncConfiguration param.Field[NewAccountingSyncConfigurationParam] `json:"accounting_sync_configuration"`
	// Additional email addresses for this customer. If populated, these email
	// addresses will be CC'd for customer communications. The total number of email
	// addresses (including the primary email) cannot exceed 50.
	AdditionalEmails param.Field[[]string] `json:"additional_emails"`
	// Used to determine if invoices for this customer will automatically attempt to
	// charge a saved payment method, if available. This parameter defaults to `True`
	// when a payment provider is provided on customer creation.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Used to determine if invoices for this customer will be automatically issued. If
	// true, invoices will be automatically issued. If false, invoices will require
	// manual approval.If `null` is specified, the customer's auto issuance setting
	// will be inherited from the account-level setting.
	AutoIssuance   param.Field[bool]              `json:"auto_issuance"`
	BillingAddress param.Field[AddressInputParam] `json:"billing_address"`
	// An ISO 4217 currency string used for the customer's invoices and balance. If not
	// set at creation time, will be set at subscription creation time.
	Currency param.Field[string] `json:"currency"`
	// A valid customer email, to be used for invoicing and notifications.
	Email         param.Field[string] `json:"email" format:"email"`
	EmailDelivery param.Field[bool]   `json:"email_delivery"`
	// The external customer ID. This can only be set if the customer has no existing
	// external customer ID. Since this action may change usage quantities for all
	// existing subscriptions, it is disallowed if the customer has issued invoices
	// with usage line items and subject to the same restrictions as backdated
	// subscription creation.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// The hierarchical relationships for this customer.
	Hierarchy param.Field[CustomerHierarchyConfigParam] `json:"hierarchy"`
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
	PaymentProviderID      param.Field[string]                                                `json:"payment_provider_id"`
	ReportingConfiguration param.Field[NewReportingConfigurationParam]                        `json:"reporting_configuration"`
	ShippingAddress        param.Field[AddressInputParam]                                     `json:"shipping_address"`
	TaxConfiguration       param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationUnion] `json:"tax_configuration"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country                | Type         | Description                                                                                             |
	// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
	// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
	// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
	// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
	// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
	// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
	// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
	// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
	// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
	// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
	// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
	// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
	// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
	// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
	// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
	// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
	// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
	// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
	// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
	// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
	// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                  | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
	// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
	// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
	// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
	// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
	// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
	// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
	// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
	// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
	// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
	// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
	// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
	// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
	// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
	// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
	// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
	// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States          | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
	// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
	// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
	// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
	TaxID param.Field[shared.CustomerTaxIDParam] `json:"tax_id"`
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

type CustomerUpdateByExternalIDParamsTaxConfiguration struct {
	TaxExempt   param.Field[bool]                                                        `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool]   `json:"automatic_tax_enabled"`
	TaxExemptionCode    param.Field[string] `json:"tax_exemption_code"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

// Satisfied by [NewAvalaraTaxConfigurationParam], [NewTaxJarConfigurationParam],
// [NewSphereConfigurationParam],
// [CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfiguration],
// [CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfiguration],
// [CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfiguration],
// [CustomerUpdateByExternalIDParamsTaxConfiguration].
type CustomerUpdateByExternalIDParamsTaxConfigurationUnion interface {
	implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion()
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfiguration struct {
	TaxExempt   param.Field[bool]                                                                               `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfigurationTaxProviderNumeral CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfigurationTaxProvider = "numeral"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationNewNumeralConfigurationTaxProviderNumeral:
		return true
	}
	return false
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfiguration struct {
	TaxExempt   param.Field[bool]                                                                             `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfigurationTaxProviderAnrok CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfigurationTaxProvider = "anrok"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationNewAnrokConfigurationTaxProviderAnrok:
		return true
	}
	return false
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfiguration struct {
	TaxExempt   param.Field[bool]                                                                                 `json:"tax_exempt,required"`
	TaxProvider param.Field[CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	// Whether to automatically calculate tax for this customer. When null, inherits
	// from account-level setting. When true or false, overrides the account setting.
	AutomaticTaxEnabled param.Field[bool] `json:"automatic_tax_enabled"`
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfiguration) implementsCustomerUpdateByExternalIDParamsTaxConfigurationUnion() {
}

type CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfigurationTaxProviderStripe CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider = "stripe"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationNewStripeTaxConfigurationTaxProviderStripe:
		return true
	}
	return false
}

type CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider string

const (
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderAvalara CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "avalara"
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderTaxjar  CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "taxjar"
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderSphere  CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "sphere"
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderNumeral CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "numeral"
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderAnrok   CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "anrok"
	CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderStripe  CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider = "stripe"
)

func (r CustomerUpdateByExternalIDParamsTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderAvalara, CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderTaxjar, CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderSphere, CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderNumeral, CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderAnrok, CustomerUpdateByExternalIDParamsTaxConfigurationTaxProviderStripe:
		return true
	}
	return false
}
