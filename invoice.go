// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// InvoiceService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	return
}

// This endpoint is used to create a one-off invoice for a customer.
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to update the `metadata`, `net_terms`, and `due_date`
// properties on an invoice. If you pass null for the metadata value, it will clear
// any existing metadata for that invoice.
//
// `metadata` can be modified regardless of invoice state. `net_terms` and
// `due_date` can only be modified if the invoice is in a `draft` state.
func (r *InvoiceService) Update(ctx context.Context, invoiceID string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all [`Invoice`](/core-concepts#invoice)s for an
// account in a list format.
//
// The list of invoices is ordered starting from the most recently issued invoice
// date. The response also includes
// [`pagination_metadata`](/api-reference/pagination), which lets the caller
// retrieve the next page of results if they exist.
//
// By default, this only returns invoices that are `issued`, `paid`, or `synced`.
//
// When fetching any `draft` invoices, this returns the last-computed invoice
// values for each draft invoice, which may not always be up-to-date since Orb
// regularly refreshes invoices asynchronously.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *pagination.Page[shared.Invoice], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "invoices"
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

// This endpoint returns a list of all [`Invoice`](/core-concepts#invoice)s for an
// account in a list format.
//
// The list of invoices is ordered starting from the most recently issued invoice
// date. The response also includes
// [`pagination_metadata`](/api-reference/pagination), which lets the caller
// retrieve the next page of results if they exist.
//
// By default, this only returns invoices that are `issued`, `paid`, or `synced`.
//
// When fetching any `draft` invoices, this returns the last-computed invoice
// values for each draft invoice, which may not always be up-to-date since Orb
// regularly refreshes invoices asynchronously.
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.Invoice] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch an [`Invoice`](/core-concepts#invoice) given an
// identifier.
func (r *InvoiceService) Fetch(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint can be used to fetch the upcoming
// [invoice](/core-concepts#invoice) for the current billing period given a
// subscription.
func (r *InvoiceService) FetchUpcoming(ctx context.Context, query InvoiceFetchUpcomingParams, opts ...option.RequestOption) (res *InvoiceFetchUpcomingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "invoices/upcoming"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint allows an eligible invoice to be issued manually. This is only
// possible with invoices where status is `draft`, `will_auto_issue` is false, and
// an `eligible_to_issue_at` is a time in the past. Issuing an invoice could
// possibly trigger side effects, some of which could be customer-visible (e.g.
// sending emails, auto-collecting payment, syncing the invoice to external
// providers, etc).
func (r *InvoiceService) Issue(ctx context.Context, invoiceID string, body InvoiceIssueParams, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/issue", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set to the `paid` status. This
// can only be done to invoices that are in the `issued` or `synced` status.
func (r *InvoiceService) MarkPaid(ctx context.Context, invoiceID string, body InvoiceMarkPaidParams, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/mark_paid", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint collects payment for an invoice using the customer's default
// payment method. This action can only be taken on invoices with status "issued".
func (r *InvoiceService) Pay(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/pay", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set to the `void` status. This
// can only be done to invoices that are in the `issued` status.
//
// If the associated invoice has used the customer balance to change the amount
// due, the customer balance operation will be reverted. For example, if the
// invoice used \$10 of customer balance, that amount will be added back to the
// customer balance upon voiding.
//
// If the invoice was used to purchase a credit block, but the invoice is not yet
// paid, the credit block will be voided. If the invoice was created due to a
// top-up, the top-up will be disabled.
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *shared.Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/void", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type InvoiceFetchUpcomingResponse struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                                     `json:"amount_due,required"`
	AutoCollection InvoiceFetchUpcomingResponseAutoCollection `json:"auto_collection,required"`
	BillingAddress shared.Address                             `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceFetchUpcomingResponseCreditNote `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                                                   `json:"currency,required"`
	Customer                    shared.CustomerMinified                                  `json:"customer,required"`
	CustomerBalanceTransactions []InvoiceFetchUpcomingResponseCustomerBalanceTransaction `json:"customer_balance_transactions,required"`
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
	CustomerTaxID shared.CustomerTaxID `json:"customer_tax_id,required,nullable"`
	// This field is deprecated in favor of `discounts`. If a `discounts` list is
	// provided, the first discount in the list will be returned. If the list is empty,
	// `None` will be returned.
	//
	// Deprecated: deprecated
	Discount  interface{}                   `json:"discount,required"`
	Discounts []shared.InvoiceLevelDiscount `json:"discounts,required"`
	// When the invoice payment is due. The due date is null if the invoice is not yet
	// finalized.
	DueDate time.Time `json:"due_date,required,nullable" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at,required,nullable" format:"date-time"`
	// A URL for the customer-facing invoice portal. This URL expires 30 days after the
	// invoice's due date, or 60 days after being re-generated through the UI.
	HostedInvoiceURL string `json:"hosted_invoice_url,required,nullable"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf    string                                    `json:"invoice_pdf,required,nullable"`
	InvoiceSource InvoiceFetchUpcomingResponseInvoiceSource `json:"invoice_source,required"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at,required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at,required,nullable" format:"date-time"`
	// The breakdown of prices in this invoice.
	LineItems     []InvoiceFetchUpcomingResponseLineItem `json:"line_items,required"`
	Maximum       shared.Maximum                         `json:"maximum,required,nullable"`
	MaximumAmount string                                 `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string `json:"metadata,required"`
	Minimum       shared.Minimum    `json:"minimum,required,nullable"`
	MinimumAmount string            `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// A list of payment attempts associated with the invoice
	PaymentAttempts []InvoiceFetchUpcomingResponsePaymentAttempt `json:"payment_attempts,required"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at,required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time                          `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  shared.Address                     `json:"shipping_address,required,nullable"`
	Status           InvoiceFetchUpcomingResponseStatus `json:"status,required"`
	Subscription     shared.SubscriptionMinified        `json:"subscription,required,nullable"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal,required"`
	// If the invoice failed to sync, this will be the last time an external invoicing
	// provider sync was attempted. This field will always be `null` for invoices using
	// Orb Invoicing.
	SyncFailedAt time.Time `json:"sync_failed_at,required,nullable" format:"date-time"`
	// The scheduled date of the invoice
	TargetDate time.Time `json:"target_date,required" format:"date-time"`
	// The total after any minimums and discounts have been applied.
	Total string `json:"total,required"`
	// If the invoice has a status of `void`, this gives a timestamp when the invoice
	// was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// This is true if the invoice will be automatically issued in the future, and
	// false otherwise.
	WillAutoIssue bool                             `json:"will_auto_issue,required"`
	JSON          invoiceFetchUpcomingResponseJSON `json:"-"`
}

// invoiceFetchUpcomingResponseJSON contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponse]
type invoiceFetchUpcomingResponseJSON struct {
	ID                          apijson.Field
	AmountDue                   apijson.Field
	AutoCollection              apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	CreditNotes                 apijson.Field
	Currency                    apijson.Field
	Customer                    apijson.Field
	CustomerBalanceTransactions apijson.Field
	CustomerTaxID               apijson.Field
	Discount                    apijson.Field
	Discounts                   apijson.Field
	DueDate                     apijson.Field
	EligibleToIssueAt           apijson.Field
	HostedInvoiceURL            apijson.Field
	InvoiceNumber               apijson.Field
	InvoicePdf                  apijson.Field
	InvoiceSource               apijson.Field
	IssueFailedAt               apijson.Field
	IssuedAt                    apijson.Field
	LineItems                   apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Memo                        apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	PaidAt                      apijson.Field
	PaymentAttempts             apijson.Field
	PaymentFailedAt             apijson.Field
	PaymentStartedAt            apijson.Field
	ScheduledIssueAt            apijson.Field
	ShippingAddress             apijson.Field
	Status                      apijson.Field
	Subscription                apijson.Field
	Subtotal                    apijson.Field
	SyncFailedAt                apijson.Field
	TargetDate                  apijson.Field
	Total                       apijson.Field
	VoidedAt                    apijson.Field
	WillAutoIssue               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseAutoCollection struct {
	// True only if auto-collection is enabled for this invoice.
	Enabled bool `json:"enabled,required,nullable"`
	// If the invoice is scheduled for auto-collection, this field will reflect when
	// the next attempt will occur. If dunning has been exhausted, or auto-collection
	// is not enabled for this invoice, this field will be `null`.
	NextAttemptAt time.Time `json:"next_attempt_at,required,nullable" format:"date-time"`
	// Number of auto-collection payment attempts.
	NumAttempts int64 `json:"num_attempts,required,nullable"`
	// If Orb has ever attempted payment auto-collection for this invoice, this field
	// will reflect when that attempt occurred. In conjunction with `next_attempt_at`,
	// this can be used to tell whether the invoice is currently in dunning (that is,
	// `previously_attempted_at` is non-null, and `next_attempt_time` is non-null), or
	// if dunning has been exhausted (`previously_attempted_at` is non-null, but
	// `next_attempt_time` is null).
	PreviouslyAttemptedAt time.Time                                      `json:"previously_attempted_at,required,nullable" format:"date-time"`
	JSON                  invoiceFetchUpcomingResponseAutoCollectionJSON `json:"-"`
}

// invoiceFetchUpcomingResponseAutoCollectionJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseAutoCollection]
type invoiceFetchUpcomingResponseAutoCollectionJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	NumAttempts           apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseAutoCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseAutoCollectionJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseCreditNote struct {
	ID               string `json:"id,required"`
	CreditNoteNumber string `json:"credit_note_number,required"`
	// An optional memo supplied on the credit note.
	Memo   string `json:"memo,required,nullable"`
	Reason string `json:"reason,required"`
	Total  string `json:"total,required"`
	Type   string `json:"type,required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time                                  `json:"voided_at,required,nullable" format:"date-time"`
	JSON     invoiceFetchUpcomingResponseCreditNoteJSON `json:"-"`
}

// invoiceFetchUpcomingResponseCreditNoteJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseCreditNote]
type invoiceFetchUpcomingResponseCreditNoteJSON struct {
	ID               apijson.Field
	CreditNoteNumber apijson.Field
	Memo             apijson.Field
	Reason           apijson.Field
	Total            apijson.Field
	Type             apijson.Field
	VoidedAt         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseCreditNoteJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                                        `json:"id,required"`
	Action InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time             `json:"created_at,required" format:"date-time"`
	CreditNote shared.CreditNoteTiny `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string             `json:"ending_balance,required"`
	Invoice       shared.InvoiceTiny `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                                      `json:"starting_balance,required"`
	Type            InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType `json:"type,required"`
	JSON            invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON  `json:"-"`
}

// invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON contains the JSON
// metadata for the struct [InvoiceFetchUpcomingResponseCustomerBalanceTransaction]
type invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON struct {
	ID              apijson.Field
	Action          apijson.Field
	Amount          apijson.Field
	CreatedAt       apijson.Field
	CreditNote      apijson.Field
	Description     apijson.Field
	EndingBalance   apijson.Field
	Invoice         apijson.Field
	StartingBalance apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction string

const (
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionAppliedToInvoice      InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionManualAdjustment      InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "manual_adjustment"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionProratedRefund        InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionRevertProratedRefund  InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "revert_prorated_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionReturnFromVoiding     InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "return_from_voiding"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteApplied     InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "credit_note_applied"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteVoided      InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "credit_note_voided"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionOverpaymentRefund     InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "overpayment_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionExternalPayment       InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "external_payment"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionSmallInvoiceCarryover InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "small_invoice_carryover"
)

func (r InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionAppliedToInvoice, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionManualAdjustment, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionProratedRefund, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionRevertProratedRefund, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionReturnFromVoiding, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteApplied, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteVoided, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionOverpaymentRefund, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionExternalPayment, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionSmallInvoiceCarryover:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType string

const (
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsTypeIncrement InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType = "increment"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsTypeDecrement InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType = "decrement"
)

func (r InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseCustomerBalanceTransactionsTypeIncrement, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsTypeDecrement:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseInvoiceSource string

const (
	InvoiceFetchUpcomingResponseInvoiceSourceSubscription InvoiceFetchUpcomingResponseInvoiceSource = "subscription"
	InvoiceFetchUpcomingResponseInvoiceSourcePartial      InvoiceFetchUpcomingResponseInvoiceSource = "partial"
	InvoiceFetchUpcomingResponseInvoiceSourceOneOff       InvoiceFetchUpcomingResponseInvoiceSource = "one_off"
)

func (r InvoiceFetchUpcomingResponseInvoiceSource) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseInvoiceSourceSubscription, InvoiceFetchUpcomingResponseInvoiceSourcePartial, InvoiceFetchUpcomingResponseInvoiceSourceOneOff:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The line amount after any adjustments and before overage conversion, credits and
	// partial invoicing.
	AdjustedSubtotal string `json:"adjusted_subtotal,required"`
	// All adjustments applied to the line item in the order they were applied based on
	// invoice calculations (ie. usage discounts -> amount discounts -> percentage
	// discounts -> minimums -> maximums).
	Adjustments []InvoiceFetchUpcomingResponseLineItemsAdjustment `json:"adjustments,required"`
	// The final amount for a line item after all adjustments and pre paid credits have
	// been applied.
	Amount string `json:"amount,required"`
	// The number of prepaid credits applied.
	CreditsApplied string `json:"credits_applied,required"`
	// This field is deprecated in favor of `adjustments`
	//
	// Deprecated: deprecated
	Discount shared.Discount `json:"discount,required,nullable"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// An additional filter that was used to calculate the usage for this line item.
	Filter string `json:"filter,required,nullable"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping string `json:"grouping,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Maximum shared.Maximum `json:"maximum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Minimum shared.Minimum `json:"minimum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MinimumAmount string `json:"minimum_amount,required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name,required"`
	// Any amount applied from a partial invoice
	PartiallyInvoicedAmount string `json:"partially_invoiced_amount,required"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// For more on the types of prices, see
	// [the core concepts documentation](/core-concepts#plan-and-price)
	Price shared.Price `json:"price,required"`
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceFetchUpcomingResponseLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before any adjustments.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []shared.TaxAmount `json:"tax_amounts,required"`
	// A list of customer ids that were used to calculate the usage for this line item.
	UsageCustomerIDs []string                                 `json:"usage_customer_ids,required,nullable"`
	JSON             invoiceFetchUpcomingResponseLineItemJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseLineItem]
type invoiceFetchUpcomingResponseLineItemJSON struct {
	ID                      apijson.Field
	AdjustedSubtotal        apijson.Field
	Adjustments             apijson.Field
	Amount                  apijson.Field
	CreditsApplied          apijson.Field
	Discount                apijson.Field
	EndDate                 apijson.Field
	Filter                  apijson.Field
	Grouping                apijson.Field
	Maximum                 apijson.Field
	MaximumAmount           apijson.Field
	Minimum                 apijson.Field
	MinimumAmount           apijson.Field
	Name                    apijson.Field
	PartiallyInvoicedAmount apijson.Field
	Price                   apijson.Field
	Quantity                apijson.Field
	StartDate               apijson.Field
	SubLineItems            apijson.Field
	Subtotal                apijson.Field
	TaxAmounts              apijson.Field
	UsageCustomerIDs        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsAdjustment struct {
	ID             string                                                         `json:"id,required"`
	AdjustmentType InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]shared.TransformPriceFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id,required,nullable"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                             `json:"usage_discount"`
	JSON          invoiceFetchUpcomingResponseLineItemsAdjustmentJSON `json:"-"`
	union         InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion
}

// invoiceFetchUpcomingResponseLineItemsAdjustmentJSON contains the JSON metadata
// for the struct [InvoiceFetchUpcomingResponseLineItemsAdjustment]
type invoiceFetchUpcomingResponseLineItemsAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	AmountDiscount       apijson.Field
	ItemID               apijson.Field
	MaximumAmount        apijson.Field
	MinimumAmount        apijson.Field
	PercentageDiscount   apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r invoiceFetchUpcomingResponseLineItemsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceFetchUpcomingResponseLineItemsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceFetchUpcomingResponseLineItemsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.MonetaryUsageDiscountAdjustment],
// [shared.MonetaryAmountDiscountAdjustment],
// [shared.MonetaryPercentageDiscountAdjustment],
// [shared.MonetaryMinimumAdjustment], [shared.MonetaryMaximumAdjustment].
func (r InvoiceFetchUpcomingResponseLineItemsAdjustment) AsUnion() InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.MonetaryUsageDiscountAdjustment],
// [shared.MonetaryAmountDiscountAdjustment],
// [shared.MonetaryPercentageDiscountAdjustment],
// [shared.MonetaryMinimumAdjustment] or [shared.MonetaryMaximumAdjustment].
type InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion interface {
	ImplementsInvoiceFetchUpcomingResponseLineItemsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType string

const (
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeUsageDiscount      InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType = "usage_discount"
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeAmountDiscount     InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType = "amount_discount"
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypePercentageDiscount InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType = "percentage_discount"
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeMinimum            InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType = "minimum"
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeMaximum            InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType = "maximum"
)

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeUsageDiscount, InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeAmountDiscount, InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypePercentageDiscount, InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeMinimum, InvoiceFetchUpcomingResponseLineItemsAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                `json:"amount,required"`
	Grouping     shared.SubLineItemGrouping                            `json:"grouping,required,nullable"`
	Name         string                                                `json:"name,required"`
	Quantity     float64                                               `json:"quantity,required"`
	Type         InvoiceFetchUpcomingResponseLineItemsSubLineItemsType `json:"type,required"`
	MatrixConfig shared.SubLineItemMatrixConfig                        `json:"matrix_config"`
	// This field can have the runtime type of [shared.TierSubLineItemTierConfig].
	TierConfig interface{}                                          `json:"tier_config"`
	JSON       invoiceFetchUpcomingResponseLineItemsSubLineItemJSON `json:"-"`
	union      InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemJSON contains the JSON metadata
// for the struct [InvoiceFetchUpcomingResponseLineItemsSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	MatrixConfig apijson.Field
	TierConfig   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceFetchUpcomingResponseLineItemsSubLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.MatrixSubLineItem],
// [shared.TierSubLineItem], [shared.OtherSubLineItem].
func (r InvoiceFetchUpcomingResponseLineItemsSubLineItem) AsUnion() InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion {
	return r.union
}

// Union satisfied by [shared.MatrixSubLineItem], [shared.TierSubLineItem] or
// [shared.OtherSubLineItem].
type InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion interface {
	ImplementsInvoiceFetchUpcomingResponseLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.TierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.OtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsTypeMatrix InvoiceFetchUpcomingResponseLineItemsSubLineItemsType = "matrix"
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsTypeTier   InvoiceFetchUpcomingResponseLineItemsSubLineItemsType = "tier"
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsTypeNull   InvoiceFetchUpcomingResponseLineItemsSubLineItemsType = "'null'"
)

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsSubLineItemsTypeMatrix, InvoiceFetchUpcomingResponseLineItemsSubLineItemsTypeTier, InvoiceFetchUpcomingResponseLineItemsSubLineItemsTypeNull:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponsePaymentAttempt struct {
	// The ID of the payment attempt.
	ID string `json:"id,required"`
	// The amount of the payment attempt.
	Amount string `json:"amount,required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider InvoiceFetchUpcomingResponsePaymentAttemptsPaymentProvider `json:"payment_provider,required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id,required,nullable"`
	// URL to the downloadable PDF version of the receipt. This field will be `null`
	// for payment attempts that did not succeed.
	ReceiptPdf string `json:"receipt_pdf,required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                                           `json:"succeeded,required"`
	JSON      invoiceFetchUpcomingResponsePaymentAttemptJSON `json:"-"`
}

// invoiceFetchUpcomingResponsePaymentAttemptJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponsePaymentAttempt]
type invoiceFetchUpcomingResponsePaymentAttemptJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	ReceiptPdf        apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponsePaymentAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponsePaymentAttemptJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type InvoiceFetchUpcomingResponsePaymentAttemptsPaymentProvider string

const (
	InvoiceFetchUpcomingResponsePaymentAttemptsPaymentProviderStripe InvoiceFetchUpcomingResponsePaymentAttemptsPaymentProvider = "stripe"
)

func (r InvoiceFetchUpcomingResponsePaymentAttemptsPaymentProvider) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponsePaymentAttemptsPaymentProviderStripe:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseStatus string

const (
	InvoiceFetchUpcomingResponseStatusIssued InvoiceFetchUpcomingResponseStatus = "issued"
	InvoiceFetchUpcomingResponseStatusPaid   InvoiceFetchUpcomingResponseStatus = "paid"
	InvoiceFetchUpcomingResponseStatusSynced InvoiceFetchUpcomingResponseStatus = "synced"
	InvoiceFetchUpcomingResponseStatusVoid   InvoiceFetchUpcomingResponseStatus = "void"
	InvoiceFetchUpcomingResponseStatusDraft  InvoiceFetchUpcomingResponseStatus = "draft"
)

func (r InvoiceFetchUpcomingResponseStatus) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseStatusIssued, InvoiceFetchUpcomingResponseStatusPaid, InvoiceFetchUpcomingResponseStatusSynced, InvoiceFetchUpcomingResponseStatusVoid, InvoiceFetchUpcomingResponseStatusDraft:
		return true
	}
	return false
}

type InvoiceNewParams struct {
	// An ISO 4217 currency string. Must be the same as the customer's currency if it
	// is set.
	Currency param.Field[string] `json:"currency,required"`
	// Optional invoice date to set. Must be in the past, if not set, `invoice_date` is
	// set to the current time in the customer's timezone.
	InvoiceDate param.Field[time.Time]                  `json:"invoice_date,required" format:"date-time"`
	LineItems   param.Field[[]InvoiceNewParamsLineItem] `json:"line_items,required"`
	// The id of the `Customer` to create this invoice for. One of `customer_id` and
	// `external_customer_id` are required.
	CustomerID param.Field[string] `json:"customer_id"`
	// An optional discount to attach to the invoice.
	Discount param.Field[shared.DiscountUnionParam] `json:"discount"`
	// An optional custom due date for the invoice. If not set, the due date will be
	// calculated based on the `net_terms` value.
	DueDate param.Field[time.Time] `json:"due_date" format:"date-time"`
	// The `external_customer_id` of the `Customer` to create this invoice for. One of
	// `customer_id` and `external_customer_id` are required.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// An optional memo to attach to the invoice. If no memo is provided, we will
	// attach the default memo
	Memo param.Field[string] `json:"memo"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The net terms determines the due date of the invoice. Due date is calculated
	// based on the invoice or issuance date, depending on the account's configured due
	// date calculation method. A value of '0' here represents that the invoice is due
	// on issue, whereas a value of '30' represents that the customer has 30 days to
	// pay the invoice. Do not set this field if you want to set a custom due date.
	NetTerms param.Field[int64] `json:"net_terms"`
	// When true, this invoice will be submitted for issuance upon creation. When
	// false, the resulting invoice will require manual review to issue. Defaulted to
	// false.
	WillAutoIssue param.Field[bool] `json:"will_auto_issue"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsLineItem struct {
	// A date string to specify the line item's end date in the customer's timezone.
	EndDate   param.Field[time.Time]                          `json:"end_date,required" format:"date"`
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[InvoiceNewParamsLineItemsModelType] `json:"model_type,required"`
	// The name of the line item.
	Name param.Field[string] `json:"name,required"`
	// The number of units on the line item
	Quantity param.Field[float64] `json:"quantity,required"`
	// A date string to specify the line item's start date in the customer's timezone.
	StartDate param.Field[time.Time] `json:"start_date,required" format:"date"`
	// Configuration for unit pricing
	UnitConfig param.Field[shared.UnitConfigParam] `json:"unit_config,required"`
}

func (r InvoiceNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsLineItemsModelType string

const (
	InvoiceNewParamsLineItemsModelTypeUnit InvoiceNewParamsLineItemsModelType = "unit"
)

func (r InvoiceNewParamsLineItemsModelType) IsKnown() bool {
	switch r {
	case InvoiceNewParamsLineItemsModelTypeUnit:
		return true
	}
	return false
}

type InvoiceUpdateParams struct {
	// An optional custom due date for the invoice. If not set, the due date will be
	// calculated based on the `net_terms` value.
	DueDate param.Field[time.Time] `json:"due_date" format:"date-time"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The net terms determines the due date of the invoice. Due date is calculated
	// based on the invoice or issuance date, depending on the account's configured due
	// date calculation method. A value of '0' here represents that the invoice is due
	// on issue, whereas a value of '30' represents that the customer has 30 days to
	// pay the invoice. Do not set this field if you want to set a custom due date.
	NetTerms param.Field[int64] `json:"net_terms"`
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceListParams struct {
	Amount   param.Field[string] `query:"amount"`
	AmountGt param.Field[string] `query:"amount[gt]"`
	AmountLt param.Field[string] `query:"amount[lt]"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor     param.Field[string]                    `query:"cursor"`
	CustomerID param.Field[string]                    `query:"customer_id"`
	DateType   param.Field[InvoiceListParamsDateType] `query:"date_type"`
	DueDate    param.Field[time.Time]                 `query:"due_date" format:"date"`
	// Filters invoices by their due dates within a specific time range in the past.
	// Specify the range as a number followed by 'd' (days) or 'm' (months). For
	// example, '7d' filters invoices due in the last 7 days, and '2m' filters those
	// due in the last 2 months.
	DueDateWindow      param.Field[string]    `query:"due_date_window"`
	DueDateGt          param.Field[time.Time] `query:"due_date[gt]" format:"date"`
	DueDateLt          param.Field[time.Time] `query:"due_date[lt]" format:"date"`
	ExternalCustomerID param.Field[string]    `query:"external_customer_id"`
	InvoiceDateGt      param.Field[time.Time] `query:"invoice_date[gt]" format:"date-time"`
	InvoiceDateGte     param.Field[time.Time] `query:"invoice_date[gte]" format:"date-time"`
	InvoiceDateLt      param.Field[time.Time] `query:"invoice_date[lt]" format:"date-time"`
	InvoiceDateLte     param.Field[time.Time] `query:"invoice_date[lte]" format:"date-time"`
	IsRecurring        param.Field[bool]      `query:"is_recurring"`
	// The number of items to fetch. Defaults to 20.
	Limit          param.Field[int64]                     `query:"limit"`
	Status         param.Field[[]InvoiceListParamsStatus] `query:"status"`
	SubscriptionID param.Field[string]                    `query:"subscription_id"`
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceListParamsDateType string

const (
	InvoiceListParamsDateTypeDueDate     InvoiceListParamsDateType = "due_date"
	InvoiceListParamsDateTypeInvoiceDate InvoiceListParamsDateType = "invoice_date"
)

func (r InvoiceListParamsDateType) IsKnown() bool {
	switch r {
	case InvoiceListParamsDateTypeDueDate, InvoiceListParamsDateTypeInvoiceDate:
		return true
	}
	return false
}

type InvoiceListParamsStatus string

const (
	InvoiceListParamsStatusDraft  InvoiceListParamsStatus = "draft"
	InvoiceListParamsStatusIssued InvoiceListParamsStatus = "issued"
	InvoiceListParamsStatusPaid   InvoiceListParamsStatus = "paid"
	InvoiceListParamsStatusSynced InvoiceListParamsStatus = "synced"
	InvoiceListParamsStatusVoid   InvoiceListParamsStatus = "void"
)

func (r InvoiceListParamsStatus) IsKnown() bool {
	switch r {
	case InvoiceListParamsStatusDraft, InvoiceListParamsStatusIssued, InvoiceListParamsStatusPaid, InvoiceListParamsStatusSynced, InvoiceListParamsStatusVoid:
		return true
	}
	return false
}

type InvoiceFetchUpcomingParams struct {
	SubscriptionID param.Field[string] `query:"subscription_id,required"`
}

// URLQuery serializes [InvoiceFetchUpcomingParams]'s query parameters as
// `url.Values`.
func (r InvoiceFetchUpcomingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceIssueParams struct {
	// If true, the invoice will be issued synchronously. If false, the invoice will be
	// issued asynchronously. The synchronous option is only available for invoices
	// that have no usage fees. If the invoice is configured to sync to an external
	// provider, a successful response from this endpoint guarantees the invoice is
	// present in the provider.
	Synchronous param.Field[bool] `json:"synchronous"`
}

func (r InvoiceIssueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceMarkPaidParams struct {
	// A date string to specify the date of the payment.
	PaymentReceivedDate param.Field[time.Time] `json:"payment_received_date,required" format:"date"`
	// An optional external ID to associate with the payment.
	ExternalID param.Field[string] `json:"external_id"`
	// An optional note to associate with the payment.
	Notes param.Field[string] `json:"notes"`
}

func (r InvoiceMarkPaidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
