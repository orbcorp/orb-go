// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
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
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to update the `metadata` property on an invoice. If you
// pass null for the metadata value, it will clear any existing metadata for that
// invoice.
//
// `metadata` can be modified regardless of invoice state.
func (r *InvoiceService) Update(ctx context.Context, invoiceID string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *pagination.Page[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Invoice] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch an [`Invoice`](/core-concepts#invoice) given an
// identifier.
func (r *InvoiceService) Fetch(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
func (r *InvoiceService) Issue(ctx context.Context, invoiceID string, body InvoiceIssueParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/issue", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set the `paid` status. This can
// only be done to invoices that are in the `issued` status.
func (r *InvoiceService) MarkPaid(ctx context.Context, invoiceID string, body InvoiceMarkPaidParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *InvoiceService) Pay(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/pay", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set the `void` status. This can
// only be done to invoices that are in the `issued` status.
//
// If the associated invoice has used the customer balance to change the amount
// due, the customer balance operation will be reverted. For example, if the
// invoice used $10 of customer balance, that amount will be added back to the
// customer balance upon voiding.
//
// If the invoice was used to purchase a credit block, but the invoice is not yet
// paid, the credit block will be voided. If the invoice was created due to a
// top-up, the top-up will be disabled.
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s/void", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
// representing the request for payment for a single subscription. This includes a
// set of line items, which correspond to prices in the subscription's plan and can
// represent fixed recurring fees or usage-based fees. They are generated at the
// end of a billing period, or as the result of an action, such as a cancellation.
type Invoice struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                `json:"amount_due,required"`
	AutoCollection InvoiceAutoCollection `json:"auto_collection,required"`
	BillingAddress InvoiceBillingAddress `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceCreditNote `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                              `json:"currency,required"`
	Customer                    InvoiceCustomer                     `json:"customer,required"`
	CustomerBalanceTransactions []InvoiceCustomerBalanceTransaction `json:"customer_balance_transactions,required"`
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
	CustomerTaxID InvoiceCustomerTaxID `json:"customer_tax_id,required,nullable"`
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
	// The scheduled date of the invoice
	InvoiceDate time.Time `json:"invoice_date,required" format:"date-time"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf    string               `json:"invoice_pdf,required,nullable"`
	InvoiceSource InvoiceInvoiceSource `json:"invoice_source,required"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at,required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at,required,nullable" format:"date-time"`
	// The breakdown of prices in this invoice.
	LineItems     []InvoiceLineItem `json:"line_items,required"`
	Maximum       InvoiceMaximum    `json:"maximum,required,nullable"`
	MaximumAmount string            `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string `json:"metadata,required"`
	Minimum       InvoiceMinimum    `json:"minimum,required,nullable"`
	MinimumAmount string            `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// A list of payment attempts associated with the invoice
	PaymentAttempts []InvoicePaymentAttempt `json:"payment_attempts,required"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at,required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time              `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  InvoiceShippingAddress `json:"shipping_address,required,nullable"`
	Status           InvoiceStatus          `json:"status,required"`
	Subscription     InvoiceSubscription    `json:"subscription,required,nullable"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal,required"`
	// If the invoice failed to sync, this will be the last time an external invoicing
	// provider sync was attempted. This field will always be `null` for invoices using
	// Orb Invoicing.
	SyncFailedAt time.Time `json:"sync_failed_at,required,nullable" format:"date-time"`
	// The total after any minimums and discounts have been applied.
	Total string `json:"total,required"`
	// If the invoice has a status of `void`, this gives a timestamp when the invoice
	// was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// This is true if the invoice will be automatically issued in the future, and
	// false otherwise.
	WillAutoIssue bool        `json:"will_auto_issue,required"`
	JSON          invoiceJSON `json:"-"`
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
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
	InvoiceDate                 apijson.Field
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
	Total                       apijson.Field
	VoidedAt                    apijson.Field
	WillAutoIssue               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceAutoCollection struct {
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
	PreviouslyAttemptedAt time.Time                 `json:"previously_attempted_at,required,nullable" format:"date-time"`
	JSON                  invoiceAutoCollectionJSON `json:"-"`
}

// invoiceAutoCollectionJSON contains the JSON metadata for the struct
// [InvoiceAutoCollection]
type invoiceAutoCollectionJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	NumAttempts           apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceAutoCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutoCollectionJSON) RawJSON() string {
	return r.raw
}

type InvoiceBillingAddress struct {
	City       string                    `json:"city,required,nullable"`
	Country    string                    `json:"country,required,nullable"`
	Line1      string                    `json:"line1,required,nullable"`
	Line2      string                    `json:"line2,required,nullable"`
	PostalCode string                    `json:"postal_code,required,nullable"`
	State      string                    `json:"state,required,nullable"`
	JSON       invoiceBillingAddressJSON `json:"-"`
}

// invoiceBillingAddressJSON contains the JSON metadata for the struct
// [InvoiceBillingAddress]
type invoiceBillingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceBillingAddressJSON) RawJSON() string {
	return r.raw
}

type InvoiceCreditNote struct {
	ID               string `json:"id,required"`
	CreditNoteNumber string `json:"credit_note_number,required"`
	// An optional memo supplied on the credit note.
	Memo   string `json:"memo,required,nullable"`
	Reason string `json:"reason,required"`
	Total  string `json:"total,required"`
	Type   string `json:"type,required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time             `json:"voided_at,required,nullable" format:"date-time"`
	JSON     invoiceCreditNoteJSON `json:"-"`
}

// invoiceCreditNoteJSON contains the JSON metadata for the struct
// [InvoiceCreditNote]
type invoiceCreditNoteJSON struct {
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

func (r *InvoiceCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCreditNoteJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomer struct {
	ID                 string              `json:"id,required"`
	ExternalCustomerID string              `json:"external_customer_id,required,nullable"`
	JSON               invoiceCustomerJSON `json:"-"`
}

// invoiceCustomerJSON contains the JSON metadata for the struct [InvoiceCustomer]
type invoiceCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                   `json:"id,required"`
	Action InvoiceCustomerBalanceTransactionsAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                    `json:"created_at,required" format:"date-time"`
	CreditNote InvoiceCustomerBalanceTransactionsCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                    `json:"ending_balance,required"`
	Invoice       InvoiceCustomerBalanceTransactionsInvoice `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                 `json:"starting_balance,required"`
	Type            InvoiceCustomerBalanceTransactionsType `json:"type,required"`
	JSON            invoiceCustomerBalanceTransactionJSON  `json:"-"`
}

// invoiceCustomerBalanceTransactionJSON contains the JSON metadata for the struct
// [InvoiceCustomerBalanceTransaction]
type invoiceCustomerBalanceTransactionJSON struct {
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

func (r *InvoiceCustomerBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerBalanceTransactionJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerBalanceTransactionsAction string

const (
	InvoiceCustomerBalanceTransactionsActionAppliedToInvoice     InvoiceCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceCustomerBalanceTransactionsActionManualAdjustment     InvoiceCustomerBalanceTransactionsAction = "manual_adjustment"
	InvoiceCustomerBalanceTransactionsActionProratedRefund       InvoiceCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceCustomerBalanceTransactionsActionRevertProratedRefund InvoiceCustomerBalanceTransactionsAction = "revert_prorated_refund"
	InvoiceCustomerBalanceTransactionsActionReturnFromVoiding    InvoiceCustomerBalanceTransactionsAction = "return_from_voiding"
	InvoiceCustomerBalanceTransactionsActionCreditNoteApplied    InvoiceCustomerBalanceTransactionsAction = "credit_note_applied"
	InvoiceCustomerBalanceTransactionsActionCreditNoteVoided     InvoiceCustomerBalanceTransactionsAction = "credit_note_voided"
	InvoiceCustomerBalanceTransactionsActionOverpaymentRefund    InvoiceCustomerBalanceTransactionsAction = "overpayment_refund"
	InvoiceCustomerBalanceTransactionsActionExternalPayment      InvoiceCustomerBalanceTransactionsAction = "external_payment"
)

func (r InvoiceCustomerBalanceTransactionsAction) IsKnown() bool {
	switch r {
	case InvoiceCustomerBalanceTransactionsActionAppliedToInvoice, InvoiceCustomerBalanceTransactionsActionManualAdjustment, InvoiceCustomerBalanceTransactionsActionProratedRefund, InvoiceCustomerBalanceTransactionsActionRevertProratedRefund, InvoiceCustomerBalanceTransactionsActionReturnFromVoiding, InvoiceCustomerBalanceTransactionsActionCreditNoteApplied, InvoiceCustomerBalanceTransactionsActionCreditNoteVoided, InvoiceCustomerBalanceTransactionsActionOverpaymentRefund, InvoiceCustomerBalanceTransactionsActionExternalPayment:
		return true
	}
	return false
}

type InvoiceCustomerBalanceTransactionsCreditNote struct {
	// The id of the Credit note
	ID   string                                           `json:"id,required"`
	JSON invoiceCustomerBalanceTransactionsCreditNoteJSON `json:"-"`
}

// invoiceCustomerBalanceTransactionsCreditNoteJSON contains the JSON metadata for
// the struct [InvoiceCustomerBalanceTransactionsCreditNote]
type invoiceCustomerBalanceTransactionsCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerBalanceTransactionsCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerBalanceTransactionsCreditNoteJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerBalanceTransactionsInvoice struct {
	// The Invoice id
	ID   string                                        `json:"id,required"`
	JSON invoiceCustomerBalanceTransactionsInvoiceJSON `json:"-"`
}

// invoiceCustomerBalanceTransactionsInvoiceJSON contains the JSON metadata for the
// struct [InvoiceCustomerBalanceTransactionsInvoice]
type invoiceCustomerBalanceTransactionsInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerBalanceTransactionsInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerBalanceTransactionsInvoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerBalanceTransactionsType string

const (
	InvoiceCustomerBalanceTransactionsTypeIncrement InvoiceCustomerBalanceTransactionsType = "increment"
	InvoiceCustomerBalanceTransactionsTypeDecrement InvoiceCustomerBalanceTransactionsType = "decrement"
)

func (r InvoiceCustomerBalanceTransactionsType) IsKnown() bool {
	switch r {
	case InvoiceCustomerBalanceTransactionsTypeIncrement, InvoiceCustomerBalanceTransactionsTypeDecrement:
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
type InvoiceCustomerTaxID struct {
	Country InvoiceCustomerTaxIDCountry `json:"country,required"`
	Type    InvoiceCustomerTaxIDType    `json:"type,required"`
	Value   string                      `json:"value,required"`
	JSON    invoiceCustomerTaxIDJSON    `json:"-"`
}

// invoiceCustomerTaxIDJSON contains the JSON metadata for the struct
// [InvoiceCustomerTaxID]
type invoiceCustomerTaxIDJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerTaxID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerTaxIDJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerTaxIDCountry string

const (
	InvoiceCustomerTaxIDCountryAd InvoiceCustomerTaxIDCountry = "AD"
	InvoiceCustomerTaxIDCountryAe InvoiceCustomerTaxIDCountry = "AE"
	InvoiceCustomerTaxIDCountryAr InvoiceCustomerTaxIDCountry = "AR"
	InvoiceCustomerTaxIDCountryAt InvoiceCustomerTaxIDCountry = "AT"
	InvoiceCustomerTaxIDCountryAu InvoiceCustomerTaxIDCountry = "AU"
	InvoiceCustomerTaxIDCountryBe InvoiceCustomerTaxIDCountry = "BE"
	InvoiceCustomerTaxIDCountryBg InvoiceCustomerTaxIDCountry = "BG"
	InvoiceCustomerTaxIDCountryBh InvoiceCustomerTaxIDCountry = "BH"
	InvoiceCustomerTaxIDCountryBo InvoiceCustomerTaxIDCountry = "BO"
	InvoiceCustomerTaxIDCountryBr InvoiceCustomerTaxIDCountry = "BR"
	InvoiceCustomerTaxIDCountryCa InvoiceCustomerTaxIDCountry = "CA"
	InvoiceCustomerTaxIDCountryCh InvoiceCustomerTaxIDCountry = "CH"
	InvoiceCustomerTaxIDCountryCl InvoiceCustomerTaxIDCountry = "CL"
	InvoiceCustomerTaxIDCountryCn InvoiceCustomerTaxIDCountry = "CN"
	InvoiceCustomerTaxIDCountryCo InvoiceCustomerTaxIDCountry = "CO"
	InvoiceCustomerTaxIDCountryCr InvoiceCustomerTaxIDCountry = "CR"
	InvoiceCustomerTaxIDCountryCy InvoiceCustomerTaxIDCountry = "CY"
	InvoiceCustomerTaxIDCountryCz InvoiceCustomerTaxIDCountry = "CZ"
	InvoiceCustomerTaxIDCountryDe InvoiceCustomerTaxIDCountry = "DE"
	InvoiceCustomerTaxIDCountryDk InvoiceCustomerTaxIDCountry = "DK"
	InvoiceCustomerTaxIDCountryEe InvoiceCustomerTaxIDCountry = "EE"
	InvoiceCustomerTaxIDCountryDo InvoiceCustomerTaxIDCountry = "DO"
	InvoiceCustomerTaxIDCountryEc InvoiceCustomerTaxIDCountry = "EC"
	InvoiceCustomerTaxIDCountryEg InvoiceCustomerTaxIDCountry = "EG"
	InvoiceCustomerTaxIDCountryEs InvoiceCustomerTaxIDCountry = "ES"
	InvoiceCustomerTaxIDCountryEu InvoiceCustomerTaxIDCountry = "EU"
	InvoiceCustomerTaxIDCountryFi InvoiceCustomerTaxIDCountry = "FI"
	InvoiceCustomerTaxIDCountryFr InvoiceCustomerTaxIDCountry = "FR"
	InvoiceCustomerTaxIDCountryGB InvoiceCustomerTaxIDCountry = "GB"
	InvoiceCustomerTaxIDCountryGe InvoiceCustomerTaxIDCountry = "GE"
	InvoiceCustomerTaxIDCountryGr InvoiceCustomerTaxIDCountry = "GR"
	InvoiceCustomerTaxIDCountryHk InvoiceCustomerTaxIDCountry = "HK"
	InvoiceCustomerTaxIDCountryHr InvoiceCustomerTaxIDCountry = "HR"
	InvoiceCustomerTaxIDCountryHu InvoiceCustomerTaxIDCountry = "HU"
	InvoiceCustomerTaxIDCountryID InvoiceCustomerTaxIDCountry = "ID"
	InvoiceCustomerTaxIDCountryIe InvoiceCustomerTaxIDCountry = "IE"
	InvoiceCustomerTaxIDCountryIl InvoiceCustomerTaxIDCountry = "IL"
	InvoiceCustomerTaxIDCountryIn InvoiceCustomerTaxIDCountry = "IN"
	InvoiceCustomerTaxIDCountryIs InvoiceCustomerTaxIDCountry = "IS"
	InvoiceCustomerTaxIDCountryIt InvoiceCustomerTaxIDCountry = "IT"
	InvoiceCustomerTaxIDCountryJp InvoiceCustomerTaxIDCountry = "JP"
	InvoiceCustomerTaxIDCountryKe InvoiceCustomerTaxIDCountry = "KE"
	InvoiceCustomerTaxIDCountryKr InvoiceCustomerTaxIDCountry = "KR"
	InvoiceCustomerTaxIDCountryKz InvoiceCustomerTaxIDCountry = "KZ"
	InvoiceCustomerTaxIDCountryLi InvoiceCustomerTaxIDCountry = "LI"
	InvoiceCustomerTaxIDCountryLt InvoiceCustomerTaxIDCountry = "LT"
	InvoiceCustomerTaxIDCountryLu InvoiceCustomerTaxIDCountry = "LU"
	InvoiceCustomerTaxIDCountryLv InvoiceCustomerTaxIDCountry = "LV"
	InvoiceCustomerTaxIDCountryMt InvoiceCustomerTaxIDCountry = "MT"
	InvoiceCustomerTaxIDCountryMx InvoiceCustomerTaxIDCountry = "MX"
	InvoiceCustomerTaxIDCountryMy InvoiceCustomerTaxIDCountry = "MY"
	InvoiceCustomerTaxIDCountryNg InvoiceCustomerTaxIDCountry = "NG"
	InvoiceCustomerTaxIDCountryNl InvoiceCustomerTaxIDCountry = "NL"
	InvoiceCustomerTaxIDCountryNo InvoiceCustomerTaxIDCountry = "NO"
	InvoiceCustomerTaxIDCountryNz InvoiceCustomerTaxIDCountry = "NZ"
	InvoiceCustomerTaxIDCountryOm InvoiceCustomerTaxIDCountry = "OM"
	InvoiceCustomerTaxIDCountryPe InvoiceCustomerTaxIDCountry = "PE"
	InvoiceCustomerTaxIDCountryPh InvoiceCustomerTaxIDCountry = "PH"
	InvoiceCustomerTaxIDCountryPl InvoiceCustomerTaxIDCountry = "PL"
	InvoiceCustomerTaxIDCountryPt InvoiceCustomerTaxIDCountry = "PT"
	InvoiceCustomerTaxIDCountryRo InvoiceCustomerTaxIDCountry = "RO"
	InvoiceCustomerTaxIDCountryRs InvoiceCustomerTaxIDCountry = "RS"
	InvoiceCustomerTaxIDCountryRu InvoiceCustomerTaxIDCountry = "RU"
	InvoiceCustomerTaxIDCountrySa InvoiceCustomerTaxIDCountry = "SA"
	InvoiceCustomerTaxIDCountrySe InvoiceCustomerTaxIDCountry = "SE"
	InvoiceCustomerTaxIDCountrySg InvoiceCustomerTaxIDCountry = "SG"
	InvoiceCustomerTaxIDCountrySi InvoiceCustomerTaxIDCountry = "SI"
	InvoiceCustomerTaxIDCountrySk InvoiceCustomerTaxIDCountry = "SK"
	InvoiceCustomerTaxIDCountrySv InvoiceCustomerTaxIDCountry = "SV"
	InvoiceCustomerTaxIDCountryTh InvoiceCustomerTaxIDCountry = "TH"
	InvoiceCustomerTaxIDCountryTr InvoiceCustomerTaxIDCountry = "TR"
	InvoiceCustomerTaxIDCountryTw InvoiceCustomerTaxIDCountry = "TW"
	InvoiceCustomerTaxIDCountryUa InvoiceCustomerTaxIDCountry = "UA"
	InvoiceCustomerTaxIDCountryUs InvoiceCustomerTaxIDCountry = "US"
	InvoiceCustomerTaxIDCountryUy InvoiceCustomerTaxIDCountry = "UY"
	InvoiceCustomerTaxIDCountryVe InvoiceCustomerTaxIDCountry = "VE"
	InvoiceCustomerTaxIDCountryVn InvoiceCustomerTaxIDCountry = "VN"
	InvoiceCustomerTaxIDCountryZa InvoiceCustomerTaxIDCountry = "ZA"
)

func (r InvoiceCustomerTaxIDCountry) IsKnown() bool {
	switch r {
	case InvoiceCustomerTaxIDCountryAd, InvoiceCustomerTaxIDCountryAe, InvoiceCustomerTaxIDCountryAr, InvoiceCustomerTaxIDCountryAt, InvoiceCustomerTaxIDCountryAu, InvoiceCustomerTaxIDCountryBe, InvoiceCustomerTaxIDCountryBg, InvoiceCustomerTaxIDCountryBh, InvoiceCustomerTaxIDCountryBo, InvoiceCustomerTaxIDCountryBr, InvoiceCustomerTaxIDCountryCa, InvoiceCustomerTaxIDCountryCh, InvoiceCustomerTaxIDCountryCl, InvoiceCustomerTaxIDCountryCn, InvoiceCustomerTaxIDCountryCo, InvoiceCustomerTaxIDCountryCr, InvoiceCustomerTaxIDCountryCy, InvoiceCustomerTaxIDCountryCz, InvoiceCustomerTaxIDCountryDe, InvoiceCustomerTaxIDCountryDk, InvoiceCustomerTaxIDCountryEe, InvoiceCustomerTaxIDCountryDo, InvoiceCustomerTaxIDCountryEc, InvoiceCustomerTaxIDCountryEg, InvoiceCustomerTaxIDCountryEs, InvoiceCustomerTaxIDCountryEu, InvoiceCustomerTaxIDCountryFi, InvoiceCustomerTaxIDCountryFr, InvoiceCustomerTaxIDCountryGB, InvoiceCustomerTaxIDCountryGe, InvoiceCustomerTaxIDCountryGr, InvoiceCustomerTaxIDCountryHk, InvoiceCustomerTaxIDCountryHr, InvoiceCustomerTaxIDCountryHu, InvoiceCustomerTaxIDCountryID, InvoiceCustomerTaxIDCountryIe, InvoiceCustomerTaxIDCountryIl, InvoiceCustomerTaxIDCountryIn, InvoiceCustomerTaxIDCountryIs, InvoiceCustomerTaxIDCountryIt, InvoiceCustomerTaxIDCountryJp, InvoiceCustomerTaxIDCountryKe, InvoiceCustomerTaxIDCountryKr, InvoiceCustomerTaxIDCountryKz, InvoiceCustomerTaxIDCountryLi, InvoiceCustomerTaxIDCountryLt, InvoiceCustomerTaxIDCountryLu, InvoiceCustomerTaxIDCountryLv, InvoiceCustomerTaxIDCountryMt, InvoiceCustomerTaxIDCountryMx, InvoiceCustomerTaxIDCountryMy, InvoiceCustomerTaxIDCountryNg, InvoiceCustomerTaxIDCountryNl, InvoiceCustomerTaxIDCountryNo, InvoiceCustomerTaxIDCountryNz, InvoiceCustomerTaxIDCountryOm, InvoiceCustomerTaxIDCountryPe, InvoiceCustomerTaxIDCountryPh, InvoiceCustomerTaxIDCountryPl, InvoiceCustomerTaxIDCountryPt, InvoiceCustomerTaxIDCountryRo, InvoiceCustomerTaxIDCountryRs, InvoiceCustomerTaxIDCountryRu, InvoiceCustomerTaxIDCountrySa, InvoiceCustomerTaxIDCountrySe, InvoiceCustomerTaxIDCountrySg, InvoiceCustomerTaxIDCountrySi, InvoiceCustomerTaxIDCountrySk, InvoiceCustomerTaxIDCountrySv, InvoiceCustomerTaxIDCountryTh, InvoiceCustomerTaxIDCountryTr, InvoiceCustomerTaxIDCountryTw, InvoiceCustomerTaxIDCountryUa, InvoiceCustomerTaxIDCountryUs, InvoiceCustomerTaxIDCountryUy, InvoiceCustomerTaxIDCountryVe, InvoiceCustomerTaxIDCountryVn, InvoiceCustomerTaxIDCountryZa:
		return true
	}
	return false
}

type InvoiceCustomerTaxIDType string

const (
	InvoiceCustomerTaxIDTypeAdNrt    InvoiceCustomerTaxIDType = "ad_nrt"
	InvoiceCustomerTaxIDTypeAeTrn    InvoiceCustomerTaxIDType = "ae_trn"
	InvoiceCustomerTaxIDTypeArCuit   InvoiceCustomerTaxIDType = "ar_cuit"
	InvoiceCustomerTaxIDTypeEuVat    InvoiceCustomerTaxIDType = "eu_vat"
	InvoiceCustomerTaxIDTypeAuAbn    InvoiceCustomerTaxIDType = "au_abn"
	InvoiceCustomerTaxIDTypeAuArn    InvoiceCustomerTaxIDType = "au_arn"
	InvoiceCustomerTaxIDTypeBgUic    InvoiceCustomerTaxIDType = "bg_uic"
	InvoiceCustomerTaxIDTypeBhVat    InvoiceCustomerTaxIDType = "bh_vat"
	InvoiceCustomerTaxIDTypeBoTin    InvoiceCustomerTaxIDType = "bo_tin"
	InvoiceCustomerTaxIDTypeBrCnpj   InvoiceCustomerTaxIDType = "br_cnpj"
	InvoiceCustomerTaxIDTypeBrCpf    InvoiceCustomerTaxIDType = "br_cpf"
	InvoiceCustomerTaxIDTypeCaBn     InvoiceCustomerTaxIDType = "ca_bn"
	InvoiceCustomerTaxIDTypeCaGstHst InvoiceCustomerTaxIDType = "ca_gst_hst"
	InvoiceCustomerTaxIDTypeCaPstBc  InvoiceCustomerTaxIDType = "ca_pst_bc"
	InvoiceCustomerTaxIDTypeCaPstMB  InvoiceCustomerTaxIDType = "ca_pst_mb"
	InvoiceCustomerTaxIDTypeCaPstSk  InvoiceCustomerTaxIDType = "ca_pst_sk"
	InvoiceCustomerTaxIDTypeCaQst    InvoiceCustomerTaxIDType = "ca_qst"
	InvoiceCustomerTaxIDTypeChVat    InvoiceCustomerTaxIDType = "ch_vat"
	InvoiceCustomerTaxIDTypeClTin    InvoiceCustomerTaxIDType = "cl_tin"
	InvoiceCustomerTaxIDTypeCnTin    InvoiceCustomerTaxIDType = "cn_tin"
	InvoiceCustomerTaxIDTypeCoNit    InvoiceCustomerTaxIDType = "co_nit"
	InvoiceCustomerTaxIDTypeCrTin    InvoiceCustomerTaxIDType = "cr_tin"
	InvoiceCustomerTaxIDTypeDoRcn    InvoiceCustomerTaxIDType = "do_rcn"
	InvoiceCustomerTaxIDTypeEcRuc    InvoiceCustomerTaxIDType = "ec_ruc"
	InvoiceCustomerTaxIDTypeEgTin    InvoiceCustomerTaxIDType = "eg_tin"
	InvoiceCustomerTaxIDTypeEsCif    InvoiceCustomerTaxIDType = "es_cif"
	InvoiceCustomerTaxIDTypeEuOssVat InvoiceCustomerTaxIDType = "eu_oss_vat"
	InvoiceCustomerTaxIDTypeGBVat    InvoiceCustomerTaxIDType = "gb_vat"
	InvoiceCustomerTaxIDTypeGeVat    InvoiceCustomerTaxIDType = "ge_vat"
	InvoiceCustomerTaxIDTypeHkBr     InvoiceCustomerTaxIDType = "hk_br"
	InvoiceCustomerTaxIDTypeHuTin    InvoiceCustomerTaxIDType = "hu_tin"
	InvoiceCustomerTaxIDTypeIDNpwp   InvoiceCustomerTaxIDType = "id_npwp"
	InvoiceCustomerTaxIDTypeIlVat    InvoiceCustomerTaxIDType = "il_vat"
	InvoiceCustomerTaxIDTypeInGst    InvoiceCustomerTaxIDType = "in_gst"
	InvoiceCustomerTaxIDTypeIsVat    InvoiceCustomerTaxIDType = "is_vat"
	InvoiceCustomerTaxIDTypeJpCn     InvoiceCustomerTaxIDType = "jp_cn"
	InvoiceCustomerTaxIDTypeJpRn     InvoiceCustomerTaxIDType = "jp_rn"
	InvoiceCustomerTaxIDTypeJpTrn    InvoiceCustomerTaxIDType = "jp_trn"
	InvoiceCustomerTaxIDTypeKePin    InvoiceCustomerTaxIDType = "ke_pin"
	InvoiceCustomerTaxIDTypeKrBrn    InvoiceCustomerTaxIDType = "kr_brn"
	InvoiceCustomerTaxIDTypeKzBin    InvoiceCustomerTaxIDType = "kz_bin"
	InvoiceCustomerTaxIDTypeLiUid    InvoiceCustomerTaxIDType = "li_uid"
	InvoiceCustomerTaxIDTypeMxRfc    InvoiceCustomerTaxIDType = "mx_rfc"
	InvoiceCustomerTaxIDTypeMyFrp    InvoiceCustomerTaxIDType = "my_frp"
	InvoiceCustomerTaxIDTypeMyItn    InvoiceCustomerTaxIDType = "my_itn"
	InvoiceCustomerTaxIDTypeMySst    InvoiceCustomerTaxIDType = "my_sst"
	InvoiceCustomerTaxIDTypeNgTin    InvoiceCustomerTaxIDType = "ng_tin"
	InvoiceCustomerTaxIDTypeNoVat    InvoiceCustomerTaxIDType = "no_vat"
	InvoiceCustomerTaxIDTypeNoVoec   InvoiceCustomerTaxIDType = "no_voec"
	InvoiceCustomerTaxIDTypeNzGst    InvoiceCustomerTaxIDType = "nz_gst"
	InvoiceCustomerTaxIDTypeOmVat    InvoiceCustomerTaxIDType = "om_vat"
	InvoiceCustomerTaxIDTypePeRuc    InvoiceCustomerTaxIDType = "pe_ruc"
	InvoiceCustomerTaxIDTypePhTin    InvoiceCustomerTaxIDType = "ph_tin"
	InvoiceCustomerTaxIDTypeRoTin    InvoiceCustomerTaxIDType = "ro_tin"
	InvoiceCustomerTaxIDTypeRsPib    InvoiceCustomerTaxIDType = "rs_pib"
	InvoiceCustomerTaxIDTypeRuInn    InvoiceCustomerTaxIDType = "ru_inn"
	InvoiceCustomerTaxIDTypeRuKpp    InvoiceCustomerTaxIDType = "ru_kpp"
	InvoiceCustomerTaxIDTypeSaVat    InvoiceCustomerTaxIDType = "sa_vat"
	InvoiceCustomerTaxIDTypeSgGst    InvoiceCustomerTaxIDType = "sg_gst"
	InvoiceCustomerTaxIDTypeSgUen    InvoiceCustomerTaxIDType = "sg_uen"
	InvoiceCustomerTaxIDTypeSiTin    InvoiceCustomerTaxIDType = "si_tin"
	InvoiceCustomerTaxIDTypeSvNit    InvoiceCustomerTaxIDType = "sv_nit"
	InvoiceCustomerTaxIDTypeThVat    InvoiceCustomerTaxIDType = "th_vat"
	InvoiceCustomerTaxIDTypeTrTin    InvoiceCustomerTaxIDType = "tr_tin"
	InvoiceCustomerTaxIDTypeTwVat    InvoiceCustomerTaxIDType = "tw_vat"
	InvoiceCustomerTaxIDTypeUaVat    InvoiceCustomerTaxIDType = "ua_vat"
	InvoiceCustomerTaxIDTypeUsEin    InvoiceCustomerTaxIDType = "us_ein"
	InvoiceCustomerTaxIDTypeUyRuc    InvoiceCustomerTaxIDType = "uy_ruc"
	InvoiceCustomerTaxIDTypeVeRif    InvoiceCustomerTaxIDType = "ve_rif"
	InvoiceCustomerTaxIDTypeVnTin    InvoiceCustomerTaxIDType = "vn_tin"
	InvoiceCustomerTaxIDTypeZaVat    InvoiceCustomerTaxIDType = "za_vat"
)

func (r InvoiceCustomerTaxIDType) IsKnown() bool {
	switch r {
	case InvoiceCustomerTaxIDTypeAdNrt, InvoiceCustomerTaxIDTypeAeTrn, InvoiceCustomerTaxIDTypeArCuit, InvoiceCustomerTaxIDTypeEuVat, InvoiceCustomerTaxIDTypeAuAbn, InvoiceCustomerTaxIDTypeAuArn, InvoiceCustomerTaxIDTypeBgUic, InvoiceCustomerTaxIDTypeBhVat, InvoiceCustomerTaxIDTypeBoTin, InvoiceCustomerTaxIDTypeBrCnpj, InvoiceCustomerTaxIDTypeBrCpf, InvoiceCustomerTaxIDTypeCaBn, InvoiceCustomerTaxIDTypeCaGstHst, InvoiceCustomerTaxIDTypeCaPstBc, InvoiceCustomerTaxIDTypeCaPstMB, InvoiceCustomerTaxIDTypeCaPstSk, InvoiceCustomerTaxIDTypeCaQst, InvoiceCustomerTaxIDTypeChVat, InvoiceCustomerTaxIDTypeClTin, InvoiceCustomerTaxIDTypeCnTin, InvoiceCustomerTaxIDTypeCoNit, InvoiceCustomerTaxIDTypeCrTin, InvoiceCustomerTaxIDTypeDoRcn, InvoiceCustomerTaxIDTypeEcRuc, InvoiceCustomerTaxIDTypeEgTin, InvoiceCustomerTaxIDTypeEsCif, InvoiceCustomerTaxIDTypeEuOssVat, InvoiceCustomerTaxIDTypeGBVat, InvoiceCustomerTaxIDTypeGeVat, InvoiceCustomerTaxIDTypeHkBr, InvoiceCustomerTaxIDTypeHuTin, InvoiceCustomerTaxIDTypeIDNpwp, InvoiceCustomerTaxIDTypeIlVat, InvoiceCustomerTaxIDTypeInGst, InvoiceCustomerTaxIDTypeIsVat, InvoiceCustomerTaxIDTypeJpCn, InvoiceCustomerTaxIDTypeJpRn, InvoiceCustomerTaxIDTypeJpTrn, InvoiceCustomerTaxIDTypeKePin, InvoiceCustomerTaxIDTypeKrBrn, InvoiceCustomerTaxIDTypeKzBin, InvoiceCustomerTaxIDTypeLiUid, InvoiceCustomerTaxIDTypeMxRfc, InvoiceCustomerTaxIDTypeMyFrp, InvoiceCustomerTaxIDTypeMyItn, InvoiceCustomerTaxIDTypeMySst, InvoiceCustomerTaxIDTypeNgTin, InvoiceCustomerTaxIDTypeNoVat, InvoiceCustomerTaxIDTypeNoVoec, InvoiceCustomerTaxIDTypeNzGst, InvoiceCustomerTaxIDTypeOmVat, InvoiceCustomerTaxIDTypePeRuc, InvoiceCustomerTaxIDTypePhTin, InvoiceCustomerTaxIDTypeRoTin, InvoiceCustomerTaxIDTypeRsPib, InvoiceCustomerTaxIDTypeRuInn, InvoiceCustomerTaxIDTypeRuKpp, InvoiceCustomerTaxIDTypeSaVat, InvoiceCustomerTaxIDTypeSgGst, InvoiceCustomerTaxIDTypeSgUen, InvoiceCustomerTaxIDTypeSiTin, InvoiceCustomerTaxIDTypeSvNit, InvoiceCustomerTaxIDTypeThVat, InvoiceCustomerTaxIDTypeTrTin, InvoiceCustomerTaxIDTypeTwVat, InvoiceCustomerTaxIDTypeUaVat, InvoiceCustomerTaxIDTypeUsEin, InvoiceCustomerTaxIDTypeUyRuc, InvoiceCustomerTaxIDTypeVeRif, InvoiceCustomerTaxIDTypeVnTin, InvoiceCustomerTaxIDTypeZaVat:
		return true
	}
	return false
}

type InvoiceInvoiceSource string

const (
	InvoiceInvoiceSourceSubscription InvoiceInvoiceSource = "subscription"
	InvoiceInvoiceSourcePartial      InvoiceInvoiceSource = "partial"
	InvoiceInvoiceSourceOneOff       InvoiceInvoiceSource = "one_off"
)

func (r InvoiceInvoiceSource) IsKnown() bool {
	switch r {
	case InvoiceInvoiceSourceSubscription, InvoiceInvoiceSourcePartial, InvoiceInvoiceSourceOneOff:
		return true
	}
	return false
}

type InvoiceLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The line amount after any adjustments and before overage conversion, credits and
	// partial invoicing.
	AdjustedSubtotal string `json:"adjusted_subtotal,required"`
	// All adjustments applied to the line item in the order they were applied based on
	// invoice calculations (ie. usage discounts -> amount discounts -> percentage
	// discounts -> minimums -> maximums).
	Adjustments []InvoiceLineItemsAdjustment `json:"adjustments,required"`
	// The final amount for a line item after all adjustments and pre paid credits have
	// been applied.
	Amount string `json:"amount,required"`
	// The number of prepaid credits applied.
	CreditsApplied string          `json:"credits_applied,required"`
	Discount       shared.Discount `json:"discount,required,nullable"`
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
	Maximum InvoiceLineItemsMaximum `json:"maximum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Minimum InvoiceLineItemsMinimum `json:"minimum,required,nullable"`
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
	Price Price `json:"price,required,nullable"`
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before before any adjustments.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceLineItemsTaxAmount `json:"tax_amounts,required"`
	// A list of customer ids that were used to calculate the usage for this line item.
	UsageCustomerIDs []string            `json:"usage_customer_ids,required,nullable"`
	JSON             invoiceLineItemJSON `json:"-"`
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
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

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsAdjustment struct {
	ID             string                                    `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
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
	UsageDiscount float64                        `json:"usage_discount"`
	JSON          invoiceLineItemsAdjustmentJSON `json:"-"`
	union         InvoiceLineItemsAdjustmentsUnion
}

// invoiceLineItemsAdjustmentJSON contains the JSON metadata for the struct
// [InvoiceLineItemsAdjustment]
type invoiceLineItemsAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
	Reason             apijson.Field
	AmountDiscount     apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	MinimumAmount      apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r invoiceLineItemsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemsAdjustmentsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment].
func (r InvoiceLineItemsAdjustment) AsUnion() InvoiceLineItemsAdjustmentsUnion {
	return r.union
}

// Union satisfied by [InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment],
// [InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment] or
// [InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment].
type InvoiceLineItemsAdjustmentsUnion interface {
	implementsInvoiceLineItemsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemsAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment struct {
	ID             string                                                                   `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                        `json:"usage_discount,required"`
	JSON          invoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment]
type invoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustment) implementsInvoiceLineItemsAdjustment() {
}

type InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment struct {
	ID             string                                                                    `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string                                                          `json:"reason,required,nullable"`
	JSON   invoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment]
type invoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustment) implementsInvoiceLineItemsAdjustment() {
}

type InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment struct {
	ID             string                                                                        `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The reason for the adjustment.
	Reason string                                                              `json:"reason,required,nullable"`
	JSON   invoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment]
type invoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment) implementsInvoiceLineItemsAdjustment() {
}

type InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment struct {
	ID             string                                                             `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                   `json:"reason,required,nullable"`
	JSON   invoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON `json:"-"`
}

// invoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON contains the JSON
// metadata for the struct [InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment]
type invoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustment) implementsInvoiceLineItemsAdjustment() {
}

type InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType = "minimum"
)

func (r InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment struct {
	ID             string                                                             `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                   `json:"reason,required,nullable"`
	JSON   invoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON `json:"-"`
}

// invoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON contains the JSON
// metadata for the struct [InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment]
type invoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustment) implementsInvoiceLineItemsAdjustment() {
}

type InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType = "maximum"
)

func (r InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceLineItemsAdjustmentsAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsAdjustmentTypeUsageDiscount      InvoiceLineItemsAdjustmentsAdjustmentType = "usage_discount"
	InvoiceLineItemsAdjustmentsAdjustmentTypeAmountDiscount     InvoiceLineItemsAdjustmentsAdjustmentType = "amount_discount"
	InvoiceLineItemsAdjustmentsAdjustmentTypePercentageDiscount InvoiceLineItemsAdjustmentsAdjustmentType = "percentage_discount"
	InvoiceLineItemsAdjustmentsAdjustmentTypeMinimum            InvoiceLineItemsAdjustmentsAdjustmentType = "minimum"
	InvoiceLineItemsAdjustmentsAdjustmentTypeMaximum            InvoiceLineItemsAdjustmentsAdjustmentType = "maximum"
)

func (r InvoiceLineItemsAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsAdjustmentTypeUsageDiscount, InvoiceLineItemsAdjustmentsAdjustmentTypeAmountDiscount, InvoiceLineItemsAdjustmentsAdjustmentTypePercentageDiscount, InvoiceLineItemsAdjustmentsAdjustmentTypeMinimum, InvoiceLineItemsAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

// This field is deprecated in favor of `adjustments`.
//
// Deprecated: deprecated
type InvoiceLineItemsMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                      `json:"maximum_amount,required"`
	JSON          invoiceLineItemsMaximumJSON `json:"-"`
}

// invoiceLineItemsMaximumJSON contains the JSON metadata for the struct
// [InvoiceLineItemsMaximum]
type invoiceLineItemsMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsMaximumJSON) RawJSON() string {
	return r.raw
}

// This field is deprecated in favor of `adjustments`.
//
// Deprecated: deprecated
type InvoiceLineItemsMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                      `json:"minimum_amount,required"`
	JSON          invoiceLineItemsMinimumJSON `json:"-"`
}

// invoiceLineItemsMinimumJSON contains the JSON metadata for the struct
// [InvoiceLineItemsMinimum]
type invoiceLineItemsMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsMinimumJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItem struct {
	// The total amount for this sub line item.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of
	// [InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping],
	// [InvoiceLineItemsSubLineItemsTierSubLineItemGrouping],
	// [InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping].
	Grouping interface{}                      `json:"grouping,required"`
	Name     string                           `json:"name,required"`
	Quantity float64                          `json:"quantity,required"`
	Type     InvoiceLineItemsSubLineItemsType `json:"type,required"`
	// This field can have the runtime type of
	// [InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig].
	MatrixConfig interface{} `json:"matrix_config"`
	// This field can have the runtime type of
	// [InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig].
	TierConfig interface{}                     `json:"tier_config"`
	JSON       invoiceLineItemsSubLineItemJSON `json:"-"`
	union      InvoiceLineItemsSubLineItemsUnion
}

// invoiceLineItemsSubLineItemJSON contains the JSON metadata for the struct
// [InvoiceLineItemsSubLineItem]
type invoiceLineItemsSubLineItemJSON struct {
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

func (r invoiceLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemsSubLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemsSubLineItemsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvoiceLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemsSubLineItemsTierSubLineItem],
// [InvoiceLineItemsSubLineItemsOtherSubLineItem].
func (r InvoiceLineItemsSubLineItem) AsUnion() InvoiceLineItemsSubLineItemsUnion {
	return r.union
}

// Union satisfied by [InvoiceLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemsSubLineItemsTierSubLineItem] or
// [InvoiceLineItemsSubLineItemsOtherSubLineItem].
type InvoiceLineItemsSubLineItemsUnion interface {
	implementsInvoiceLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemsSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsSubLineItemsMatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsSubLineItemsTierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemsSubLineItemsOtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceLineItemsSubLineItemsMatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                    `json:"amount,required"`
	Grouping     InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping     `json:"grouping,required,nullable"`
	MatrixConfig InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                                                    `json:"name,required"`
	Quantity     float64                                                   `json:"quantity,required"`
	Type         InvoiceLineItemsSubLineItemsMatrixSubLineItemType         `json:"type,required"`
	JSON         invoiceLineItemsSubLineItemsMatrixSubLineItemJSON         `json:"-"`
}

// invoiceLineItemsSubLineItemsMatrixSubLineItemJSON contains the JSON metadata for
// the struct [InvoiceLineItemsSubLineItemsMatrixSubLineItem]
type invoiceLineItemsSubLineItemsMatrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsMatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsMatrixSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsSubLineItemsMatrixSubLineItem) implementsInvoiceLineItemsSubLineItem() {}

type InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                    `json:"value,required,nullable"`
	JSON  invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping]
type invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string                                                      `json:"dimension_values,required"`
	JSON            invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig]
type invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItemsMatrixSubLineItemType string

const (
	InvoiceLineItemsSubLineItemsMatrixSubLineItemTypeMatrix InvoiceLineItemsSubLineItemsMatrixSubLineItemType = "matrix"
)

func (r InvoiceLineItemsSubLineItemsMatrixSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsSubLineItemsMatrixSubLineItemTypeMatrix:
		return true
	}
	return false
}

type InvoiceLineItemsSubLineItemsTierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string                                                `json:"amount,required"`
	Grouping   InvoiceLineItemsSubLineItemsTierSubLineItemGrouping   `json:"grouping,required,nullable"`
	Name       string                                                `json:"name,required"`
	Quantity   float64                                               `json:"quantity,required"`
	TierConfig InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig `json:"tier_config,required"`
	Type       InvoiceLineItemsSubLineItemsTierSubLineItemType       `json:"type,required"`
	JSON       invoiceLineItemsSubLineItemsTierSubLineItemJSON       `json:"-"`
}

// invoiceLineItemsSubLineItemsTierSubLineItemJSON contains the JSON metadata for
// the struct [InvoiceLineItemsSubLineItemsTierSubLineItem]
type invoiceLineItemsSubLineItemsTierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsTierSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsSubLineItemsTierSubLineItem) implementsInvoiceLineItemsSubLineItem() {}

type InvoiceLineItemsSubLineItemsTierSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                  `json:"value,required,nullable"`
	JSON  invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsTierSubLineItemGrouping]
type invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig struct {
	FirstUnit  float64                                                   `json:"first_unit,required"`
	LastUnit   float64                                                   `json:"last_unit,required,nullable"`
	UnitAmount string                                                    `json:"unit_amount,required"`
	JSON       invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig]
type invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItemsTierSubLineItemType string

const (
	InvoiceLineItemsSubLineItemsTierSubLineItemTypeTier InvoiceLineItemsSubLineItemsTierSubLineItemType = "tier"
)

func (r InvoiceLineItemsSubLineItemsTierSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsSubLineItemsTierSubLineItemTypeTier:
		return true
	}
	return false
}

type InvoiceLineItemsSubLineItemsOtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                                               `json:"amount,required"`
	Grouping InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping `json:"grouping,required,nullable"`
	Name     string                                               `json:"name,required"`
	Quantity float64                                              `json:"quantity,required"`
	Type     InvoiceLineItemsSubLineItemsOtherSubLineItemType     `json:"type,required"`
	JSON     invoiceLineItemsSubLineItemsOtherSubLineItemJSON     `json:"-"`
}

// invoiceLineItemsSubLineItemsOtherSubLineItemJSON contains the JSON metadata for
// the struct [InvoiceLineItemsSubLineItemsOtherSubLineItem]
type invoiceLineItemsSubLineItemsOtherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsOtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsOtherSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemsSubLineItemsOtherSubLineItem) implementsInvoiceLineItemsSubLineItem() {}

type InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                   `json:"value,required,nullable"`
	JSON  invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping]
type invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItemsOtherSubLineItemType string

const (
	InvoiceLineItemsSubLineItemsOtherSubLineItemTypeNull InvoiceLineItemsSubLineItemsOtherSubLineItemType = "'null'"
)

func (r InvoiceLineItemsSubLineItemsOtherSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsSubLineItemsOtherSubLineItemTypeNull:
		return true
	}
	return false
}

type InvoiceLineItemsSubLineItemsType string

const (
	InvoiceLineItemsSubLineItemsTypeMatrix InvoiceLineItemsSubLineItemsType = "matrix"
	InvoiceLineItemsSubLineItemsTypeTier   InvoiceLineItemsSubLineItemsType = "tier"
	InvoiceLineItemsSubLineItemsTypeNull   InvoiceLineItemsSubLineItemsType = "'null'"
)

func (r InvoiceLineItemsSubLineItemsType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsSubLineItemsTypeMatrix, InvoiceLineItemsSubLineItemsTypeTier, InvoiceLineItemsSubLineItemsTypeNull:
		return true
	}
	return false
}

type InvoiceLineItemsTaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string                        `json:"tax_rate_percentage,required,nullable"`
	JSON              invoiceLineItemsTaxAmountJSON `json:"-"`
}

// invoiceLineItemsTaxAmountJSON contains the JSON metadata for the struct
// [InvoiceLineItemsTaxAmount]
type invoiceLineItemsTaxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceLineItemsTaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsTaxAmountJSON) RawJSON() string {
	return r.raw
}

type InvoiceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string             `json:"maximum_amount,required"`
	JSON          invoiceMaximumJSON `json:"-"`
}

// invoiceMaximumJSON contains the JSON metadata for the struct [InvoiceMaximum]
type invoiceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceMaximumJSON) RawJSON() string {
	return r.raw
}

type InvoiceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string             `json:"minimum_amount,required"`
	JSON          invoiceMinimumJSON `json:"-"`
}

// invoiceMinimumJSON contains the JSON metadata for the struct [InvoiceMinimum]
type invoiceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceMinimumJSON) RawJSON() string {
	return r.raw
}

type InvoicePaymentAttempt struct {
	// The ID of the payment attempt.
	ID string `json:"id,required"`
	// The amount of the payment attempt.
	Amount string `json:"amount,required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider InvoicePaymentAttemptsPaymentProvider `json:"payment_provider,required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id,required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                      `json:"succeeded,required"`
	JSON      invoicePaymentAttemptJSON `json:"-"`
}

// invoicePaymentAttemptJSON contains the JSON metadata for the struct
// [InvoicePaymentAttempt]
type invoicePaymentAttemptJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoicePaymentAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentAttemptJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type InvoicePaymentAttemptsPaymentProvider string

const (
	InvoicePaymentAttemptsPaymentProviderStripe InvoicePaymentAttemptsPaymentProvider = "stripe"
)

func (r InvoicePaymentAttemptsPaymentProvider) IsKnown() bool {
	switch r {
	case InvoicePaymentAttemptsPaymentProviderStripe:
		return true
	}
	return false
}

type InvoiceShippingAddress struct {
	City       string                     `json:"city,required,nullable"`
	Country    string                     `json:"country,required,nullable"`
	Line1      string                     `json:"line1,required,nullable"`
	Line2      string                     `json:"line2,required,nullable"`
	PostalCode string                     `json:"postal_code,required,nullable"`
	State      string                     `json:"state,required,nullable"`
	JSON       invoiceShippingAddressJSON `json:"-"`
}

// invoiceShippingAddressJSON contains the JSON metadata for the struct
// [InvoiceShippingAddress]
type invoiceShippingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceShippingAddressJSON) RawJSON() string {
	return r.raw
}

type InvoiceStatus string

const (
	InvoiceStatusIssued InvoiceStatus = "issued"
	InvoiceStatusPaid   InvoiceStatus = "paid"
	InvoiceStatusSynced InvoiceStatus = "synced"
	InvoiceStatusVoid   InvoiceStatus = "void"
	InvoiceStatusDraft  InvoiceStatus = "draft"
)

func (r InvoiceStatus) IsKnown() bool {
	switch r {
	case InvoiceStatusIssued, InvoiceStatusPaid, InvoiceStatusSynced, InvoiceStatusVoid, InvoiceStatusDraft:
		return true
	}
	return false
}

type InvoiceSubscription struct {
	ID   string                  `json:"id,required"`
	JSON invoiceSubscriptionJSON `json:"-"`
}

// invoiceSubscriptionJSON contains the JSON metadata for the struct
// [InvoiceSubscription]
type invoiceSubscriptionJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceSubscriptionJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponse struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                                     `json:"amount_due,required"`
	AutoCollection InvoiceFetchUpcomingResponseAutoCollection `json:"auto_collection,required"`
	BillingAddress InvoiceFetchUpcomingResponseBillingAddress `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceFetchUpcomingResponseCreditNote `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                                                   `json:"currency,required"`
	Customer                    InvoiceFetchUpcomingResponseCustomer                     `json:"customer,required"`
	CustomerBalanceTransactions []InvoiceFetchUpcomingResponseCustomerBalanceTransaction `json:"customer_balance_transactions,required"`
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
	CustomerTaxID InvoiceFetchUpcomingResponseCustomerTaxID `json:"customer_tax_id,required,nullable"`
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
	Maximum       InvoiceFetchUpcomingResponseMaximum    `json:"maximum,required,nullable"`
	MaximumAmount string                                 `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string                   `json:"metadata,required"`
	Minimum       InvoiceFetchUpcomingResponseMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                              `json:"minimum_amount,required,nullable"`
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
	ScheduledIssueAt time.Time                                   `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  InvoiceFetchUpcomingResponseShippingAddress `json:"shipping_address,required,nullable"`
	Status           InvoiceFetchUpcomingResponseStatus          `json:"status,required"`
	Subscription     InvoiceFetchUpcomingResponseSubscription    `json:"subscription,required,nullable"`
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

type InvoiceFetchUpcomingResponseBillingAddress struct {
	City       string                                         `json:"city,required,nullable"`
	Country    string                                         `json:"country,required,nullable"`
	Line1      string                                         `json:"line1,required,nullable"`
	Line2      string                                         `json:"line2,required,nullable"`
	PostalCode string                                         `json:"postal_code,required,nullable"`
	State      string                                         `json:"state,required,nullable"`
	JSON       invoiceFetchUpcomingResponseBillingAddressJSON `json:"-"`
}

// invoiceFetchUpcomingResponseBillingAddressJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseBillingAddress]
type invoiceFetchUpcomingResponseBillingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseBillingAddressJSON) RawJSON() string {
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

type InvoiceFetchUpcomingResponseCustomer struct {
	ID                 string                                   `json:"id,required"`
	ExternalCustomerID string                                   `json:"external_customer_id,required,nullable"`
	JSON               invoiceFetchUpcomingResponseCustomerJSON `json:"-"`
}

// invoiceFetchUpcomingResponseCustomerJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseCustomer]
type invoiceFetchUpcomingResponseCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseCustomerJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                                        `json:"id,required"`
	Action InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                                         `json:"created_at,required" format:"date-time"`
	CreditNote InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                                         `json:"ending_balance,required"`
	Invoice       InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice `json:"invoice,required,nullable"`
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
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionAppliedToInvoice     InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionManualAdjustment     InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "manual_adjustment"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionProratedRefund       InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionRevertProratedRefund InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "revert_prorated_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionReturnFromVoiding    InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "return_from_voiding"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteApplied    InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "credit_note_applied"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteVoided     InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "credit_note_voided"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionOverpaymentRefund    InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "overpayment_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionExternalPayment      InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "external_payment"
)

func (r InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionAppliedToInvoice, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionManualAdjustment, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionProratedRefund, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionRevertProratedRefund, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionReturnFromVoiding, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteApplied, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionCreditNoteVoided, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionOverpaymentRefund, InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionExternalPayment:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote struct {
	// The id of the Credit note
	ID   string                                                                `json:"id,required"`
	JSON invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON `json:"-"`
}

// invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote]
type invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice struct {
	// The Invoice id
	ID   string                                                             `json:"id,required"`
	JSON invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON `json:"-"`
}

// invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON contains the
// JSON metadata for the struct
// [InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice]
type invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON) RawJSON() string {
	return r.raw
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
type InvoiceFetchUpcomingResponseCustomerTaxID struct {
	Country InvoiceFetchUpcomingResponseCustomerTaxIDCountry `json:"country,required"`
	Type    InvoiceFetchUpcomingResponseCustomerTaxIDType    `json:"type,required"`
	Value   string                                           `json:"value,required"`
	JSON    invoiceFetchUpcomingResponseCustomerTaxIDJSON    `json:"-"`
}

// invoiceFetchUpcomingResponseCustomerTaxIDJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseCustomerTaxID]
type invoiceFetchUpcomingResponseCustomerTaxIDJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerTaxID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseCustomerTaxIDJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseCustomerTaxIDCountry string

const (
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAd InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AD"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBg InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BG"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CA"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCn InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CN"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCy InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CY"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCz InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CZ"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryDe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "DE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryDk InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "DK"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryEe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "EE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryDo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "DO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryEc InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "EC"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryEg InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "EG"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryEs InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "ES"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryEu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "EU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryFi InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "FI"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryFr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "FR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryGB InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "GB"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryGe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "GE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryGr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "GR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryHk InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "HK"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryHr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "HR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryHu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "HU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryID InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "ID"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryIe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "IE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryIl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "IL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryIn InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "IN"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryIs InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "IS"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryIt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "IT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryJp InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "JP"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryKe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "KE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryKr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "KR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryKz InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "KZ"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLi InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LI"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLv InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LV"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryMt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "MT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryMx InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "MX"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryMy InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "MY"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNg InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NG"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNz InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NZ"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryOm InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "OM"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryRo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "RO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryRs InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "RS"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryRu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "RU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SA"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySg InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SG"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySi InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SI"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySk InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SK"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySv InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SV"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryTh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "TH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryTr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "TR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryTw InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "TW"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryUa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "UA"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryUs InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "US"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryUy InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "UY"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryVe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "VE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryVn InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "VN"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryZa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "ZA"
)

func (r InvoiceFetchUpcomingResponseCustomerTaxIDCountry) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseCustomerTaxIDCountryAd, InvoiceFetchUpcomingResponseCustomerTaxIDCountryAe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryAr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryAt, InvoiceFetchUpcomingResponseCustomerTaxIDCountryAu, InvoiceFetchUpcomingResponseCustomerTaxIDCountryBe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryBg, InvoiceFetchUpcomingResponseCustomerTaxIDCountryBh, InvoiceFetchUpcomingResponseCustomerTaxIDCountryBo, InvoiceFetchUpcomingResponseCustomerTaxIDCountryBr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCa, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCh, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCl, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCn, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCo, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCy, InvoiceFetchUpcomingResponseCustomerTaxIDCountryCz, InvoiceFetchUpcomingResponseCustomerTaxIDCountryDe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryDk, InvoiceFetchUpcomingResponseCustomerTaxIDCountryEe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryDo, InvoiceFetchUpcomingResponseCustomerTaxIDCountryEc, InvoiceFetchUpcomingResponseCustomerTaxIDCountryEg, InvoiceFetchUpcomingResponseCustomerTaxIDCountryEs, InvoiceFetchUpcomingResponseCustomerTaxIDCountryEu, InvoiceFetchUpcomingResponseCustomerTaxIDCountryFi, InvoiceFetchUpcomingResponseCustomerTaxIDCountryFr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryGB, InvoiceFetchUpcomingResponseCustomerTaxIDCountryGe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryGr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryHk, InvoiceFetchUpcomingResponseCustomerTaxIDCountryHr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryHu, InvoiceFetchUpcomingResponseCustomerTaxIDCountryID, InvoiceFetchUpcomingResponseCustomerTaxIDCountryIe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryIl, InvoiceFetchUpcomingResponseCustomerTaxIDCountryIn, InvoiceFetchUpcomingResponseCustomerTaxIDCountryIs, InvoiceFetchUpcomingResponseCustomerTaxIDCountryIt, InvoiceFetchUpcomingResponseCustomerTaxIDCountryJp, InvoiceFetchUpcomingResponseCustomerTaxIDCountryKe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryKr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryKz, InvoiceFetchUpcomingResponseCustomerTaxIDCountryLi, InvoiceFetchUpcomingResponseCustomerTaxIDCountryLt, InvoiceFetchUpcomingResponseCustomerTaxIDCountryLu, InvoiceFetchUpcomingResponseCustomerTaxIDCountryLv, InvoiceFetchUpcomingResponseCustomerTaxIDCountryMt, InvoiceFetchUpcomingResponseCustomerTaxIDCountryMx, InvoiceFetchUpcomingResponseCustomerTaxIDCountryMy, InvoiceFetchUpcomingResponseCustomerTaxIDCountryNg, InvoiceFetchUpcomingResponseCustomerTaxIDCountryNl, InvoiceFetchUpcomingResponseCustomerTaxIDCountryNo, InvoiceFetchUpcomingResponseCustomerTaxIDCountryNz, InvoiceFetchUpcomingResponseCustomerTaxIDCountryOm, InvoiceFetchUpcomingResponseCustomerTaxIDCountryPe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryPh, InvoiceFetchUpcomingResponseCustomerTaxIDCountryPl, InvoiceFetchUpcomingResponseCustomerTaxIDCountryPt, InvoiceFetchUpcomingResponseCustomerTaxIDCountryRo, InvoiceFetchUpcomingResponseCustomerTaxIDCountryRs, InvoiceFetchUpcomingResponseCustomerTaxIDCountryRu, InvoiceFetchUpcomingResponseCustomerTaxIDCountrySa, InvoiceFetchUpcomingResponseCustomerTaxIDCountrySe, InvoiceFetchUpcomingResponseCustomerTaxIDCountrySg, InvoiceFetchUpcomingResponseCustomerTaxIDCountrySi, InvoiceFetchUpcomingResponseCustomerTaxIDCountrySk, InvoiceFetchUpcomingResponseCustomerTaxIDCountrySv, InvoiceFetchUpcomingResponseCustomerTaxIDCountryTh, InvoiceFetchUpcomingResponseCustomerTaxIDCountryTr, InvoiceFetchUpcomingResponseCustomerTaxIDCountryTw, InvoiceFetchUpcomingResponseCustomerTaxIDCountryUa, InvoiceFetchUpcomingResponseCustomerTaxIDCountryUs, InvoiceFetchUpcomingResponseCustomerTaxIDCountryUy, InvoiceFetchUpcomingResponseCustomerTaxIDCountryVe, InvoiceFetchUpcomingResponseCustomerTaxIDCountryVn, InvoiceFetchUpcomingResponseCustomerTaxIDCountryZa:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseCustomerTaxIDType string

const (
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAdNrt    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ad_nrt"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAeTrn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ae_trn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeArCuit   InvoiceFetchUpcomingResponseCustomerTaxIDType = "ar_cuit"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeEuVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "eu_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAuAbn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "au_abn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAuArn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "au_arn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeBgUic    InvoiceFetchUpcomingResponseCustomerTaxIDType = "bg_uic"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeBhVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "bh_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeBoTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "bo_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeBrCnpj   InvoiceFetchUpcomingResponseCustomerTaxIDType = "br_cnpj"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeBrCpf    InvoiceFetchUpcomingResponseCustomerTaxIDType = "br_cpf"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaBn     InvoiceFetchUpcomingResponseCustomerTaxIDType = "ca_bn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaGstHst InvoiceFetchUpcomingResponseCustomerTaxIDType = "ca_gst_hst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaPstBc  InvoiceFetchUpcomingResponseCustomerTaxIDType = "ca_pst_bc"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaPstMB  InvoiceFetchUpcomingResponseCustomerTaxIDType = "ca_pst_mb"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaPstSk  InvoiceFetchUpcomingResponseCustomerTaxIDType = "ca_pst_sk"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaQst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ca_qst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeChVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ch_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeClTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "cl_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCnTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "cn_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCoNit    InvoiceFetchUpcomingResponseCustomerTaxIDType = "co_nit"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeCrTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "cr_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeDoRcn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "do_rcn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeEcRuc    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ec_ruc"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeEgTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "eg_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeEsCif    InvoiceFetchUpcomingResponseCustomerTaxIDType = "es_cif"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeEuOssVat InvoiceFetchUpcomingResponseCustomerTaxIDType = "eu_oss_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeGBVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "gb_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeGeVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ge_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeHkBr     InvoiceFetchUpcomingResponseCustomerTaxIDType = "hk_br"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeHuTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "hu_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeIDNpwp   InvoiceFetchUpcomingResponseCustomerTaxIDType = "id_npwp"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeIlVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "il_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeInGst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "in_gst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeIsVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "is_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeJpCn     InvoiceFetchUpcomingResponseCustomerTaxIDType = "jp_cn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeJpRn     InvoiceFetchUpcomingResponseCustomerTaxIDType = "jp_rn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeJpTrn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "jp_trn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeKePin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ke_pin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeKrBrn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "kr_brn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeKzBin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "kz_bin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeLiUid    InvoiceFetchUpcomingResponseCustomerTaxIDType = "li_uid"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMxRfc    InvoiceFetchUpcomingResponseCustomerTaxIDType = "mx_rfc"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMyFrp    InvoiceFetchUpcomingResponseCustomerTaxIDType = "my_frp"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMyItn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "my_itn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMySst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "my_sst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeNgTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ng_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeNoVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "no_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeNoVoec   InvoiceFetchUpcomingResponseCustomerTaxIDType = "no_voec"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeNzGst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "nz_gst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeOmVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "om_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypePeRuc    InvoiceFetchUpcomingResponseCustomerTaxIDType = "pe_ruc"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypePhTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ph_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeRoTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ro_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeRsPib    InvoiceFetchUpcomingResponseCustomerTaxIDType = "rs_pib"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeRuInn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ru_inn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeRuKpp    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ru_kpp"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSaVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sa_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSgGst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sg_gst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSgUen    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sg_uen"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSiTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "si_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSvNit    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sv_nit"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeThVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "th_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeTrTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "tr_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeTwVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "tw_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeUaVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ua_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeUsEin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "us_ein"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeUyRuc    InvoiceFetchUpcomingResponseCustomerTaxIDType = "uy_ruc"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeVeRif    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ve_rif"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeVnTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "vn_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeZaVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "za_vat"
)

func (r InvoiceFetchUpcomingResponseCustomerTaxIDType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseCustomerTaxIDTypeAdNrt, InvoiceFetchUpcomingResponseCustomerTaxIDTypeAeTrn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeArCuit, InvoiceFetchUpcomingResponseCustomerTaxIDTypeEuVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeAuAbn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeAuArn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeBgUic, InvoiceFetchUpcomingResponseCustomerTaxIDTypeBhVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeBoTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeBrCnpj, InvoiceFetchUpcomingResponseCustomerTaxIDTypeBrCpf, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaBn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaGstHst, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaPstBc, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaPstMB, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaPstSk, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCaQst, InvoiceFetchUpcomingResponseCustomerTaxIDTypeChVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeClTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCnTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCoNit, InvoiceFetchUpcomingResponseCustomerTaxIDTypeCrTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeDoRcn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeEcRuc, InvoiceFetchUpcomingResponseCustomerTaxIDTypeEgTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeEsCif, InvoiceFetchUpcomingResponseCustomerTaxIDTypeEuOssVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeGBVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeGeVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeHkBr, InvoiceFetchUpcomingResponseCustomerTaxIDTypeHuTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeIDNpwp, InvoiceFetchUpcomingResponseCustomerTaxIDTypeIlVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeInGst, InvoiceFetchUpcomingResponseCustomerTaxIDTypeIsVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeJpCn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeJpRn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeJpTrn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeKePin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeKrBrn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeKzBin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeLiUid, InvoiceFetchUpcomingResponseCustomerTaxIDTypeMxRfc, InvoiceFetchUpcomingResponseCustomerTaxIDTypeMyFrp, InvoiceFetchUpcomingResponseCustomerTaxIDTypeMyItn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeMySst, InvoiceFetchUpcomingResponseCustomerTaxIDTypeNgTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeNoVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeNoVoec, InvoiceFetchUpcomingResponseCustomerTaxIDTypeNzGst, InvoiceFetchUpcomingResponseCustomerTaxIDTypeOmVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypePeRuc, InvoiceFetchUpcomingResponseCustomerTaxIDTypePhTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeRoTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeRsPib, InvoiceFetchUpcomingResponseCustomerTaxIDTypeRuInn, InvoiceFetchUpcomingResponseCustomerTaxIDTypeRuKpp, InvoiceFetchUpcomingResponseCustomerTaxIDTypeSaVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeSgGst, InvoiceFetchUpcomingResponseCustomerTaxIDTypeSgUen, InvoiceFetchUpcomingResponseCustomerTaxIDTypeSiTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeSvNit, InvoiceFetchUpcomingResponseCustomerTaxIDTypeThVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeTrTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeTwVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeUaVat, InvoiceFetchUpcomingResponseCustomerTaxIDTypeUsEin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeUyRuc, InvoiceFetchUpcomingResponseCustomerTaxIDTypeVeRif, InvoiceFetchUpcomingResponseCustomerTaxIDTypeVnTin, InvoiceFetchUpcomingResponseCustomerTaxIDTypeZaVat:
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
	CreditsApplied string          `json:"credits_applied,required"`
	Discount       shared.Discount `json:"discount,required,nullable"`
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
	Maximum InvoiceFetchUpcomingResponseLineItemsMaximum `json:"maximum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Minimum InvoiceFetchUpcomingResponseLineItemsMinimum `json:"minimum,required,nullable"`
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
	Price Price `json:"price,required,nullable"`
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceFetchUpcomingResponseLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before before any adjustments.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceFetchUpcomingResponseLineItemsTaxAmount `json:"tax_amounts,required"`
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
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
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
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
	Reason             apijson.Field
	AmountDiscount     apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	MinimumAmount      apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment].
func (r InvoiceFetchUpcomingResponseLineItemsAdjustment) AsUnion() InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment],
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment] or
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment].
type InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion interface {
	implementsInvoiceFetchUpcomingResponseLineItemsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFetchUpcomingResponseLineItemsAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment struct {
	ID             string                                                                                        `json:"id,required"`
	AdjustmentType InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                             `json:"usage_discount,required"`
	JSON          invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment]
type invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustment) implementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType string

const (
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment struct {
	ID             string                                                                                         `json:"id,required"`
	AdjustmentType InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string                                                                               `json:"reason,required,nullable"`
	JSON   invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment]
type invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustment) implementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType string

const (
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment struct {
	ID             string                                                                                             `json:"id,required"`
	AdjustmentType InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The reason for the adjustment.
	Reason string                                                                                   `json:"reason,required,nullable"`
	JSON   invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment]
type invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustment) implementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType string

const (
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment]
type invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustment) implementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType string

const (
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType = "minimum"
)

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment]
type invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustment) implementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType string

const (
	InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType = "maximum"
)

func (r InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
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

// This field is deprecated in favor of `adjustments`.
//
// Deprecated: deprecated
type InvoiceFetchUpcomingResponseLineItemsMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                           `json:"maximum_amount,required"`
	JSON          invoiceFetchUpcomingResponseLineItemsMaximumJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsMaximumJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseLineItemsMaximum]
type invoiceFetchUpcomingResponseLineItemsMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsMaximumJSON) RawJSON() string {
	return r.raw
}

// This field is deprecated in favor of `adjustments`.
//
// Deprecated: deprecated
type InvoiceFetchUpcomingResponseLineItemsMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                           `json:"minimum_amount,required"`
	JSON          invoiceFetchUpcomingResponseLineItemsMinimumJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsMinimumJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseLineItemsMinimum]
type invoiceFetchUpcomingResponseLineItemsMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsMinimumJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItem struct {
	// The total amount for this sub line item.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of
	// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping],
	// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping],
	// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping].
	Grouping interface{}                                           `json:"grouping,required"`
	Name     string                                                `json:"name,required"`
	Quantity float64                                               `json:"quantity,required"`
	Type     InvoiceFetchUpcomingResponseLineItemsSubLineItemsType `json:"type,required"`
	// This field can have the runtime type of
	// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig].
	MatrixConfig interface{} `json:"matrix_config"`
	// This field can have the runtime type of
	// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig].
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
// Possible runtime types of the union are
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem],
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem].
func (r InvoiceFetchUpcomingResponseLineItemsSubLineItem) AsUnion() InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion {
	return r.union
}

// Union satisfied by
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem] or
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem].
type InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion interface {
	implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFetchUpcomingResponseLineItemsSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                                         `json:"amount,required"`
	Grouping     InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping     `json:"grouping,required,nullable"`
	MatrixConfig InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                                                                         `json:"name,required"`
	Quantity     float64                                                                        `json:"quantity,required"`
	Type         InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType         `json:"type,required"`
	JSON         invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON         `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem) implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                                         `json:"value,required,nullable"`
	JSON  invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string                                                                           `json:"dimension_values,required"`
	JSON            invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemTypeMatrix InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType = "matrix"
)

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemTypeMatrix:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string                                                                     `json:"amount,required"`
	Grouping   InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping   `json:"grouping,required,nullable"`
	Name       string                                                                     `json:"name,required"`
	Quantity   float64                                                                    `json:"quantity,required"`
	TierConfig InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig `json:"tier_config,required"`
	Type       InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType       `json:"type,required"`
	JSON       invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON       `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem) implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                                       `json:"value,required,nullable"`
	JSON  invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig struct {
	FirstUnit  float64                                                                        `json:"first_unit,required"`
	LastUnit   float64                                                                        `json:"last_unit,required,nullable"`
	UnitAmount string                                                                         `json:"unit_amount,required"`
	JSON       invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTypeTier InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType = "tier"
)

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTypeTier:
		return true
	}
	return false
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                                                                    `json:"amount,required"`
	Grouping InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping `json:"grouping,required,nullable"`
	Name     string                                                                    `json:"name,required"`
	Quantity float64                                                                   `json:"quantity,required"`
	Type     InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType     `json:"type,required"`
	JSON     invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON     `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem) implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                                        `json:"value,required,nullable"`
	JSON  invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemTypeNull InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType = "'null'"
)

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemTypeNull:
		return true
	}
	return false
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

type InvoiceFetchUpcomingResponseLineItemsTaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string                                             `json:"tax_rate_percentage,required,nullable"`
	JSON              invoiceFetchUpcomingResponseLineItemsTaxAmountJSON `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemsTaxAmountJSON contains the JSON metadata
// for the struct [InvoiceFetchUpcomingResponseLineItemsTaxAmount]
type invoiceFetchUpcomingResponseLineItemsTaxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsTaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemsTaxAmountJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                  `json:"maximum_amount,required"`
	JSON          invoiceFetchUpcomingResponseMaximumJSON `json:"-"`
}

// invoiceFetchUpcomingResponseMaximumJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseMaximum]
type invoiceFetchUpcomingResponseMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseMaximumJSON) RawJSON() string {
	return r.raw
}

type InvoiceFetchUpcomingResponseMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                  `json:"minimum_amount,required"`
	JSON          invoiceFetchUpcomingResponseMinimumJSON `json:"-"`
}

// invoiceFetchUpcomingResponseMinimumJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseMinimum]
type invoiceFetchUpcomingResponseMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseMinimumJSON) RawJSON() string {
	return r.raw
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

type InvoiceFetchUpcomingResponseShippingAddress struct {
	City       string                                          `json:"city,required,nullable"`
	Country    string                                          `json:"country,required,nullable"`
	Line1      string                                          `json:"line1,required,nullable"`
	Line2      string                                          `json:"line2,required,nullable"`
	PostalCode string                                          `json:"postal_code,required,nullable"`
	State      string                                          `json:"state,required,nullable"`
	JSON       invoiceFetchUpcomingResponseShippingAddressJSON `json:"-"`
}

// invoiceFetchUpcomingResponseShippingAddressJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseShippingAddress]
type invoiceFetchUpcomingResponseShippingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseShippingAddressJSON) RawJSON() string {
	return r.raw
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

type InvoiceFetchUpcomingResponseSubscription struct {
	ID   string                                       `json:"id,required"`
	JSON invoiceFetchUpcomingResponseSubscriptionJSON `json:"-"`
}

// invoiceFetchUpcomingResponseSubscriptionJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseSubscription]
type invoiceFetchUpcomingResponseSubscriptionJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

type InvoiceNewParams struct {
	// An ISO 4217 currency string. Must be the same as the customer's currency if it
	// is set.
	Currency param.Field[string] `json:"currency,required"`
	// Optional invoice date to set. Must be in the past, if not set, `invoice_date` is
	// set to the current time in the customer's timezone.
	InvoiceDate param.Field[time.Time]                  `json:"invoice_date,required" format:"date-time"`
	LineItems   param.Field[[]InvoiceNewParamsLineItem] `json:"line_items,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of '0' here represents that the
	// invoice is due on issue, whereas a value of 30 represents that the customer has
	// 30 days to pay the invoice.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// The id of the `Customer` to create this invoice for. One of `customer_id` and
	// `external_customer_id` are required.
	CustomerID param.Field[string] `json:"customer_id"`
	// An optional discount to attach to the invoice.
	Discount param.Field[shared.DiscountUnionParam] `json:"discount"`
	// The `external_customer_id` of the `Customer` to create this invoice for. One of
	// `customer_id` and `external_customer_id` are required.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// An optional memo to attach to the invoice.
	Memo param.Field[string] `json:"memo"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// When true, this invoice will automatically be issued upon creation. When false,
	// the resulting invoice will require manual review to issue. Defaulted to false.
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
	StartDate  param.Field[time.Time]                           `json:"start_date,required" format:"date"`
	UnitConfig param.Field[InvoiceNewParamsLineItemsUnitConfig] `json:"unit_config,required"`
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

type InvoiceNewParamsLineItemsUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r InvoiceNewParamsLineItemsUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateParams struct {
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
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
