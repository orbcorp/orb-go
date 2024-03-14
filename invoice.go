// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// InvoiceService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewInvoiceService] method instead.
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

// This endpoint returns a list of all [`Invoice`](../guides/concepts#invoice)s for
// an account in a list format.
//
// The list of invoices is ordered starting from the most recently issued invoice
// date. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist.
//
// By default, this only returns invoices that are `issued`, `paid`, or `synced`.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *shared.Page[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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

// This endpoint returns a list of all [`Invoice`](../guides/concepts#invoice)s for
// an account in a list format.
//
// The list of invoices is ordered starting from the most recently issued invoice
// date. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist.
//
// By default, this only returns invoices that are `issued`, `paid`, or `synced`.
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *shared.PageAutoPager[Invoice] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch an [`Invoice`](../guides/concepts#invoice) given
// an identifier.
//
// This endpoint supports returning ApiCachedUsageData (see
// api/cache_control_utils.py)
func (r *InvoiceService) Fetch(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint can be used to fetch the upcoming
// [invoice](../guides/concepts#invoice) for the current billing period given a
// subscription.
//
// This endpoint supports returning ApiCachedUsageData (see
// api/cache_control_utils.py)
func (r *InvoiceService) FetchUpcoming(ctx context.Context, query InvoiceFetchUpcomingParams, opts ...option.RequestOption) (res *InvoiceFetchUpcomingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "invoices/upcoming"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint allows an eligible invoice to be issued manually. This is only
// possible with invoices where status is `draft`, `will_auto_issue` is true, and
// an `eligible_to_issue_at` is a time in the past. Issuing an invoice could
// possibly trigger side effects, some of which could be customer-visible (e.g.
// sending emails, auto-collecting payment, syncing the invoice to external
// providers, etc).
func (r *InvoiceService) Issue(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s/issue", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set the `paid` status. This can
// only be done to invoices that are in the `issued` status.
func (r *InvoiceService) MarkPaid(ctx context.Context, invoiceID string, body InvoiceMarkPaidParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s/mark_paid", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set the `void` status. This can
// only be done to invoices that are in the `issued` status.
//
// If the associated invoice has used the customer balance to change the amount
// due, the customer balance operation will be reverted. For example, if the
// invoice used $10 of customer balance, that amount will be added back to the
// customer balance upon voiding.
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s/void", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// An [`Invoice`](../guides/concepts#invoice) is a fundamental billing entity,
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
	CustomerTaxID InvoiceCustomerTaxID `json:"customer_tax_id,required,nullable"`
	Discount      shared.Discount      `json:"discount,required,nullable"`
	Discounts     []shared.Discount    `json:"discounts,required"`
	// When the invoice payment is due.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at,required,nullable" format:"date-time"`
	// A URL for the invoice portal.
	HostedInvoiceURL string `json:"hosted_invoice_url,required,nullable"`
	// The scheduled date of the invoice
	InvoiceDate time.Time `json:"invoice_date,required" format:"date-time"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf string `json:"invoice_pdf,required,nullable"`
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
)

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
	InvoiceCustomerTaxIDCountryAt InvoiceCustomerTaxIDCountry = "AT"
	InvoiceCustomerTaxIDCountryAu InvoiceCustomerTaxIDCountry = "AU"
	InvoiceCustomerTaxIDCountryBe InvoiceCustomerTaxIDCountry = "BE"
	InvoiceCustomerTaxIDCountryBg InvoiceCustomerTaxIDCountry = "BG"
	InvoiceCustomerTaxIDCountryBr InvoiceCustomerTaxIDCountry = "BR"
	InvoiceCustomerTaxIDCountryCa InvoiceCustomerTaxIDCountry = "CA"
	InvoiceCustomerTaxIDCountryCh InvoiceCustomerTaxIDCountry = "CH"
	InvoiceCustomerTaxIDCountryCl InvoiceCustomerTaxIDCountry = "CL"
	InvoiceCustomerTaxIDCountryCy InvoiceCustomerTaxIDCountry = "CY"
	InvoiceCustomerTaxIDCountryCz InvoiceCustomerTaxIDCountry = "CZ"
	InvoiceCustomerTaxIDCountryDe InvoiceCustomerTaxIDCountry = "DE"
	InvoiceCustomerTaxIDCountryDk InvoiceCustomerTaxIDCountry = "DK"
	InvoiceCustomerTaxIDCountryEe InvoiceCustomerTaxIDCountry = "EE"
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
	InvoiceCustomerTaxIDCountryLi InvoiceCustomerTaxIDCountry = "LI"
	InvoiceCustomerTaxIDCountryLt InvoiceCustomerTaxIDCountry = "LT"
	InvoiceCustomerTaxIDCountryLu InvoiceCustomerTaxIDCountry = "LU"
	InvoiceCustomerTaxIDCountryLv InvoiceCustomerTaxIDCountry = "LV"
	InvoiceCustomerTaxIDCountryMt InvoiceCustomerTaxIDCountry = "MT"
	InvoiceCustomerTaxIDCountryMx InvoiceCustomerTaxIDCountry = "MX"
	InvoiceCustomerTaxIDCountryMy InvoiceCustomerTaxIDCountry = "MY"
	InvoiceCustomerTaxIDCountryNl InvoiceCustomerTaxIDCountry = "NL"
	InvoiceCustomerTaxIDCountryNo InvoiceCustomerTaxIDCountry = "NO"
	InvoiceCustomerTaxIDCountryNz InvoiceCustomerTaxIDCountry = "NZ"
	InvoiceCustomerTaxIDCountryPh InvoiceCustomerTaxIDCountry = "PH"
	InvoiceCustomerTaxIDCountryPl InvoiceCustomerTaxIDCountry = "PL"
	InvoiceCustomerTaxIDCountryPt InvoiceCustomerTaxIDCountry = "PT"
	InvoiceCustomerTaxIDCountryRo InvoiceCustomerTaxIDCountry = "RO"
	InvoiceCustomerTaxIDCountryRu InvoiceCustomerTaxIDCountry = "RU"
	InvoiceCustomerTaxIDCountrySa InvoiceCustomerTaxIDCountry = "SA"
	InvoiceCustomerTaxIDCountrySe InvoiceCustomerTaxIDCountry = "SE"
	InvoiceCustomerTaxIDCountrySg InvoiceCustomerTaxIDCountry = "SG"
	InvoiceCustomerTaxIDCountrySi InvoiceCustomerTaxIDCountry = "SI"
	InvoiceCustomerTaxIDCountrySk InvoiceCustomerTaxIDCountry = "SK"
	InvoiceCustomerTaxIDCountryTh InvoiceCustomerTaxIDCountry = "TH"
	InvoiceCustomerTaxIDCountryTr InvoiceCustomerTaxIDCountry = "TR"
	InvoiceCustomerTaxIDCountryTw InvoiceCustomerTaxIDCountry = "TW"
	InvoiceCustomerTaxIDCountryUa InvoiceCustomerTaxIDCountry = "UA"
	InvoiceCustomerTaxIDCountryUs InvoiceCustomerTaxIDCountry = "US"
	InvoiceCustomerTaxIDCountryZa InvoiceCustomerTaxIDCountry = "ZA"
)

type InvoiceCustomerTaxIDType string

const (
	InvoiceCustomerTaxIDTypeAdNrt    InvoiceCustomerTaxIDType = "ad_nrt"
	InvoiceCustomerTaxIDTypeAeTrn    InvoiceCustomerTaxIDType = "ae_trn"
	InvoiceCustomerTaxIDTypeEuVat    InvoiceCustomerTaxIDType = "eu_vat"
	InvoiceCustomerTaxIDTypeAuAbn    InvoiceCustomerTaxIDType = "au_abn"
	InvoiceCustomerTaxIDTypeAuArn    InvoiceCustomerTaxIDType = "au_arn"
	InvoiceCustomerTaxIDTypeBgUic    InvoiceCustomerTaxIDType = "bg_uic"
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
	InvoiceCustomerTaxIDTypeLiUid    InvoiceCustomerTaxIDType = "li_uid"
	InvoiceCustomerTaxIDTypeMxRfc    InvoiceCustomerTaxIDType = "mx_rfc"
	InvoiceCustomerTaxIDTypeMyFrp    InvoiceCustomerTaxIDType = "my_frp"
	InvoiceCustomerTaxIDTypeMyItn    InvoiceCustomerTaxIDType = "my_itn"
	InvoiceCustomerTaxIDTypeMySst    InvoiceCustomerTaxIDType = "my_sst"
	InvoiceCustomerTaxIDTypeNoVat    InvoiceCustomerTaxIDType = "no_vat"
	InvoiceCustomerTaxIDTypeNzGst    InvoiceCustomerTaxIDType = "nz_gst"
	InvoiceCustomerTaxIDTypePhTin    InvoiceCustomerTaxIDType = "ph_tin"
	InvoiceCustomerTaxIDTypeRuInn    InvoiceCustomerTaxIDType = "ru_inn"
	InvoiceCustomerTaxIDTypeRuKpp    InvoiceCustomerTaxIDType = "ru_kpp"
	InvoiceCustomerTaxIDTypeSaVat    InvoiceCustomerTaxIDType = "sa_vat"
	InvoiceCustomerTaxIDTypeSgGst    InvoiceCustomerTaxIDType = "sg_gst"
	InvoiceCustomerTaxIDTypeSgUen    InvoiceCustomerTaxIDType = "sg_uen"
	InvoiceCustomerTaxIDTypeSiTin    InvoiceCustomerTaxIDType = "si_tin"
	InvoiceCustomerTaxIDTypeThVat    InvoiceCustomerTaxIDType = "th_vat"
	InvoiceCustomerTaxIDTypeTrTin    InvoiceCustomerTaxIDType = "tr_tin"
	InvoiceCustomerTaxIDTypeTwVat    InvoiceCustomerTaxIDType = "tw_vat"
	InvoiceCustomerTaxIDTypeUaVat    InvoiceCustomerTaxIDType = "ua_vat"
	InvoiceCustomerTaxIDTypeUsEin    InvoiceCustomerTaxIDType = "us_ein"
	InvoiceCustomerTaxIDTypeZaVat    InvoiceCustomerTaxIDType = "za_vat"
)

type InvoiceLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The final amount after any discounts or minimums.
	Amount   string          `json:"amount,required"`
	Discount shared.Discount `json:"discount,required,nullable"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping      string                  `json:"grouping,required,nullable"`
	Maximum       InvoiceLineItemsMaximum `json:"maximum,required,nullable"`
	MaximumAmount string                  `json:"maximum_amount,required,nullable"`
	Minimum       InvoiceLineItemsMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                  `json:"minimum_amount,required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name,required"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// ## Unit pricing
	//
	// With unit pricing, each unit costs a fixed amount.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	        "unit_amount": "0.50"
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered pricing
	//
	// In tiered pricing, the cost of a given unit depends on the tier range that it
	// falls into, where each tier range is defined by an upper and lower bound. For
	// example, the first ten units may cost $0.50 each and all units thereafter may
	// cost $0.10 each.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "tiered",
	//	    "tiered_config": {
	//	        "tiers": [
	//	            {
	//	                "first_unit": 1,
	//	                "last_unit": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "first_unit": 11,
	//	                "last_unit": null,
	//	                "unit_amount": "0.10"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//
	// ```
	//
	// ## Bulk pricing
	//
	// Bulk pricing applies when the number of units determine the cost of all units.
	// For example, if you've bought less than 10 units, they may each be $0.50 for a
	// total of $5.00. Once you've bought more than 10 units, all units may now be
	// priced at $0.40 (i.e. 101 units total would be $40.40).
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bulk",
	//	    "bulk_config": {
	//	        "tiers": [
	//	            {
	//	                "maximum_units": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "maximum_units": 1000,
	//	                "unit_amount": "0.40"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Package pricing
	//
	// Package pricing defines the size or granularity of a unit for billing purposes.
	// For example, if the package size is set to 5, then 4 units will be billed as 5
	// and 6 units will be billed at 10.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "package",
	//	    "package_config": {
	//	        "package_amount": "0.80",
	//	        "package_size": 10
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## BPS pricing
	//
	// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
	// percent (the number of basis points to charge), as well as a cap per event to
	// assess. For example, this would allow you to assess a fee of 0.25% on every
	// payment you process, with a maximum charge of $25 per payment.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bps",
	//	    "bps_config": {
	//	       "bps": 125,
	//	       "per_unit_maximum": "11.00"
	//	    }
	//	    ...
	//	 }
	//
	// ```
	//
	// ## Bulk BPS pricing
	//
	// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
	// total quantity across all events. Similar to bulk pricing, the BPS parameters of
	// a given event depends on the tier range that the billing period falls into. Each
	// tier range is defined by an upper bound. For example, after $1.5M of payment
	// volume is reached, each individual payment may have a lower cap or a smaller
	// take-rate.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "bulk_bps",
	//	    "bulk_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered BPS pricing
	//
	// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an
	// event's applicable parameter is a function of its marginal addition to the
	// period total. Similar to tiered pricing, the BPS parameters of a given event
	// depends on the tier range that it falls into, where each tier range is defined
	// by an upper and lower bound. For example, the first few payments may have a 0.8
	// BPS take-rate and all payments after a specific volume may incur a take-rate of
	// 0.5 BPS each.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "tiered_bps",
	//	    "tiered_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "minimum_amount": "0",
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "minimum_amount": "1000000.00",
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Matrix pricing
	//
	// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix.
	// `dimensions` defines the two event property values evaluated in this pricing
	// model. In a one-dimensional matrix, the second value is `null`. Every
	// configuration has a list of `matrix_values` which give the unit prices for
	// specified property values. In a one-dimensional matrix, the matrix values will
	// have `dimension_values` where the second value of the pair is null. If an event
	// does not match any of the dimension values in the matrix, it will resort to the
	// `default_unit_amount`.
	//
	// ```json
	//
	//	{
	//	    "model_type": "matrix"
	//	    "matrix_config": {
	//	        "default_unit_amount": "3.00",
	//	        "dimensions": [
	//	            "cluster_name",
	//	            "region"
	//	        ],
	//	        "matrix_values": [
	//	            {
	//	                "dimension_values": [
	//	                    "alpha",
	//	                    "west"
	//	                ],
	//	                "unit_amount": "2.00"
	//	            },
	//	            ...
	//	        ]
	//	    }
	//	}
	//
	// ```
	//
	// ### Fixed fees
	//
	// Fixed fees are prices that are applied independent of usage quantities, and
	// follow unit pricing. They also have an additional parameter
	// `fixed_price_quantity`. If the Price represents a fixed cost, this represents
	// the quantity of units applied.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "id": "price_id",
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	       "unit_amount": "2.00"
	//	    },
	//	    "fixed_price_quantity": 3.0
	//	    ...
	//	}
	//
	// ```
	Price    Price   `json:"price,required,nullable"`
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceLineItemsTaxAmount `json:"tax_amounts,required"`
	JSON       invoiceLineItemJSON         `json:"-"`
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	ID            apijson.Field
	Amount        apijson.Field
	Discount      apijson.Field
	EndDate       apijson.Field
	Grouping      apijson.Field
	Maximum       apijson.Field
	MaximumAmount apijson.Field
	Minimum       apijson.Field
	MinimumAmount apijson.Field
	Name          apijson.Field
	Price         apijson.Field
	Quantity      apijson.Field
	StartDate     apijson.Field
	SubLineItems  apijson.Field
	Subtotal      apijson.Field
	TaxAmounts    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemJSON) RawJSON() string {
	return r.raw
}

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

// Union satisfied by [InvoiceLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemsSubLineItemsTierSubLineItem] or
// [InvoiceLineItemsSubLineItemsOtherSubLineItem].
type InvoiceLineItemsSubLineItem interface {
	implementsInvoiceLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemsSubLineItem)(nil)).Elem(),
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
	CustomerTaxID InvoiceFetchUpcomingResponseCustomerTaxID `json:"customer_tax_id,required,nullable"`
	Discount      shared.Discount                           `json:"discount,required,nullable"`
	Discounts     []shared.Discount                         `json:"discounts,required"`
	// When the invoice payment is due.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at,required,nullable" format:"date-time"`
	// A URL for the invoice portal.
	HostedInvoiceURL string `json:"hosted_invoice_url,required,nullable"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf string `json:"invoice_pdf,required,nullable"`
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
)

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
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryAu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "AU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBg InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BG"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryBr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "BR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CA"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCy InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CY"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryCz InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "CZ"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryDe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "DE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryDk InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "DK"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryEe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "EE"
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
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLi InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LI"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryLv InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "LV"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryMt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "MT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryMx InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "MX"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryMy InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "MY"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryNz InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "NZ"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPl InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PL"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryPt InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "PT"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryRo InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "RO"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryRu InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "RU"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SA"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySe InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SE"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySg InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SG"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySi InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SI"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountrySk InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "SK"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryTh InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "TH"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryTr InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "TR"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryTw InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "TW"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryUa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "UA"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryUs InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "US"
	InvoiceFetchUpcomingResponseCustomerTaxIDCountryZa InvoiceFetchUpcomingResponseCustomerTaxIDCountry = "ZA"
)

type InvoiceFetchUpcomingResponseCustomerTaxIDType string

const (
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAdNrt    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ad_nrt"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAeTrn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ae_trn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeEuVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "eu_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAuAbn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "au_abn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeAuArn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "au_arn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeBgUic    InvoiceFetchUpcomingResponseCustomerTaxIDType = "bg_uic"
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
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeLiUid    InvoiceFetchUpcomingResponseCustomerTaxIDType = "li_uid"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMxRfc    InvoiceFetchUpcomingResponseCustomerTaxIDType = "mx_rfc"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMyFrp    InvoiceFetchUpcomingResponseCustomerTaxIDType = "my_frp"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMyItn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "my_itn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeMySst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "my_sst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeNoVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "no_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeNzGst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "nz_gst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypePhTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ph_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeRuInn    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ru_inn"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeRuKpp    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ru_kpp"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSaVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sa_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSgGst    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sg_gst"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSgUen    InvoiceFetchUpcomingResponseCustomerTaxIDType = "sg_uen"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeSiTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "si_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeThVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "th_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeTrTin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "tr_tin"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeTwVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "tw_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeUaVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "ua_vat"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeUsEin    InvoiceFetchUpcomingResponseCustomerTaxIDType = "us_ein"
	InvoiceFetchUpcomingResponseCustomerTaxIDTypeZaVat    InvoiceFetchUpcomingResponseCustomerTaxIDType = "za_vat"
)

type InvoiceFetchUpcomingResponseLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The final amount after any discounts or minimums.
	Amount   string          `json:"amount,required"`
	Discount shared.Discount `json:"discount,required,nullable"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping      string                                       `json:"grouping,required,nullable"`
	Maximum       InvoiceFetchUpcomingResponseLineItemsMaximum `json:"maximum,required,nullable"`
	MaximumAmount string                                       `json:"maximum_amount,required,nullable"`
	Minimum       InvoiceFetchUpcomingResponseLineItemsMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                                       `json:"minimum_amount,required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name,required"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// ## Unit pricing
	//
	// With unit pricing, each unit costs a fixed amount.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	        "unit_amount": "0.50"
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered pricing
	//
	// In tiered pricing, the cost of a given unit depends on the tier range that it
	// falls into, where each tier range is defined by an upper and lower bound. For
	// example, the first ten units may cost $0.50 each and all units thereafter may
	// cost $0.10 each.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "tiered",
	//	    "tiered_config": {
	//	        "tiers": [
	//	            {
	//	                "first_unit": 1,
	//	                "last_unit": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "first_unit": 11,
	//	                "last_unit": null,
	//	                "unit_amount": "0.10"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//
	// ```
	//
	// ## Bulk pricing
	//
	// Bulk pricing applies when the number of units determine the cost of all units.
	// For example, if you've bought less than 10 units, they may each be $0.50 for a
	// total of $5.00. Once you've bought more than 10 units, all units may now be
	// priced at $0.40 (i.e. 101 units total would be $40.40).
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bulk",
	//	    "bulk_config": {
	//	        "tiers": [
	//	            {
	//	                "maximum_units": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "maximum_units": 1000,
	//	                "unit_amount": "0.40"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Package pricing
	//
	// Package pricing defines the size or granularity of a unit for billing purposes.
	// For example, if the package size is set to 5, then 4 units will be billed as 5
	// and 6 units will be billed at 10.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "package",
	//	    "package_config": {
	//	        "package_amount": "0.80",
	//	        "package_size": 10
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## BPS pricing
	//
	// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
	// percent (the number of basis points to charge), as well as a cap per event to
	// assess. For example, this would allow you to assess a fee of 0.25% on every
	// payment you process, with a maximum charge of $25 per payment.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bps",
	//	    "bps_config": {
	//	       "bps": 125,
	//	       "per_unit_maximum": "11.00"
	//	    }
	//	    ...
	//	 }
	//
	// ```
	//
	// ## Bulk BPS pricing
	//
	// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
	// total quantity across all events. Similar to bulk pricing, the BPS parameters of
	// a given event depends on the tier range that the billing period falls into. Each
	// tier range is defined by an upper bound. For example, after $1.5M of payment
	// volume is reached, each individual payment may have a lower cap or a smaller
	// take-rate.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "bulk_bps",
	//	    "bulk_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered BPS pricing
	//
	// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an
	// event's applicable parameter is a function of its marginal addition to the
	// period total. Similar to tiered pricing, the BPS parameters of a given event
	// depends on the tier range that it falls into, where each tier range is defined
	// by an upper and lower bound. For example, the first few payments may have a 0.8
	// BPS take-rate and all payments after a specific volume may incur a take-rate of
	// 0.5 BPS each.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "tiered_bps",
	//	    "tiered_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "minimum_amount": "0",
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "minimum_amount": "1000000.00",
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Matrix pricing
	//
	// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix.
	// `dimensions` defines the two event property values evaluated in this pricing
	// model. In a one-dimensional matrix, the second value is `null`. Every
	// configuration has a list of `matrix_values` which give the unit prices for
	// specified property values. In a one-dimensional matrix, the matrix values will
	// have `dimension_values` where the second value of the pair is null. If an event
	// does not match any of the dimension values in the matrix, it will resort to the
	// `default_unit_amount`.
	//
	// ```json
	//
	//	{
	//	    "model_type": "matrix"
	//	    "matrix_config": {
	//	        "default_unit_amount": "3.00",
	//	        "dimensions": [
	//	            "cluster_name",
	//	            "region"
	//	        ],
	//	        "matrix_values": [
	//	            {
	//	                "dimension_values": [
	//	                    "alpha",
	//	                    "west"
	//	                ],
	//	                "unit_amount": "2.00"
	//	            },
	//	            ...
	//	        ]
	//	    }
	//	}
	//
	// ```
	//
	// ### Fixed fees
	//
	// Fixed fees are prices that are applied independent of usage quantities, and
	// follow unit pricing. They also have an additional parameter
	// `fixed_price_quantity`. If the Price represents a fixed cost, this represents
	// the quantity of units applied.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "id": "price_id",
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	       "unit_amount": "2.00"
	//	    },
	//	    "fixed_price_quantity": 3.0
	//	    ...
	//	}
	//
	// ```
	Price    Price   `json:"price,required,nullable"`
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceFetchUpcomingResponseLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceFetchUpcomingResponseLineItemsTaxAmount `json:"tax_amounts,required"`
	JSON       invoiceFetchUpcomingResponseLineItemJSON         `json:"-"`
}

// invoiceFetchUpcomingResponseLineItemJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseLineItem]
type invoiceFetchUpcomingResponseLineItemJSON struct {
	ID            apijson.Field
	Amount        apijson.Field
	Discount      apijson.Field
	EndDate       apijson.Field
	Grouping      apijson.Field
	Maximum       apijson.Field
	MaximumAmount apijson.Field
	Minimum       apijson.Field
	MinimumAmount apijson.Field
	Name          apijson.Field
	Price         apijson.Field
	Quantity      apijson.Field
	StartDate     apijson.Field
	SubLineItems  apijson.Field
	Subtotal      apijson.Field
	TaxAmounts    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFetchUpcomingResponseLineItemJSON) RawJSON() string {
	return r.raw
}

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

// Union satisfied by
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem] or
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem].
type InvoiceFetchUpcomingResponseLineItemsSubLineItem interface {
	implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFetchUpcomingResponseLineItemsSubLineItem)(nil)).Elem(),
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
	// The `external_customer_id` of the `Customer` to create this invoice for. One of
	// `customer_id` and `external_customer_id` are required.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// An optional memo to attach to the invoice.
	Memo param.Field[string] `json:"memo"`
	// When true, this invoice will automatically be issued upon creation. When false,
	// the resulting invoice will require manual review to issue. Defaulted to false.
	WillAutoIssue param.Field[bool] `json:"will_auto_issue"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsLineItem struct {
	// A date string to specify the line item's end date in the customer's timezone.
	EndDate param.Field[time.Time] `json:"end_date,required" format:"date"`
	ItemID param.Field[string] `json:"item_id,required"`
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

type InvoiceNewParamsLineItemsUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r InvoiceNewParamsLineItemsUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceListParams struct {
	Amount   param.Field[string] `query:"amount"`
	AmountGt param.Field[string] `query:"amount[gt]"`
	AmountLt param.Field[string] `query:"amount[lt]"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor             param.Field[string]                    `query:"cursor"`
	CustomerID         param.Field[string]                    `query:"customer_id"`
	DateType           param.Field[InvoiceListParamsDateType] `query:"date_type"`
	DueDate            param.Field[time.Time]                 `query:"due_date" format:"date"`
	DueDateWindow      param.Field[string]                    `query:"due_date_window"`
	DueDateGt          param.Field[time.Time]                 `query:"due_date[gt]" format:"date"`
	DueDateLt          param.Field[time.Time]                 `query:"due_date[lt]" format:"date"`
	ExternalCustomerID param.Field[string]                    `query:"external_customer_id"`
	InvoiceDateGt      param.Field[time.Time]                 `query:"invoice_date[gt]" format:"date-time"`
	InvoiceDateGte     param.Field[time.Time]                 `query:"invoice_date[gte]" format:"date-time"`
	InvoiceDateLt      param.Field[time.Time]                 `query:"invoice_date[lt]" format:"date-time"`
	InvoiceDateLte     param.Field[time.Time]                 `query:"invoice_date[lte]" format:"date-time"`
	IsRecurring        param.Field[bool]                      `query:"is_recurring"`
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

type InvoiceListParamsStatus string

const (
	InvoiceListParamsStatusDraft  InvoiceListParamsStatus = "draft"
	InvoiceListParamsStatusIssued InvoiceListParamsStatus = "issued"
	InvoiceListParamsStatusPaid   InvoiceListParamsStatus = "paid"
	InvoiceListParamsStatusSynced InvoiceListParamsStatus = "synced"
	InvoiceListParamsStatusVoid   InvoiceListParamsStatus = "void"
)

type InvoiceFetchUpcomingParams struct {
	SubscriptionID param.Field[string] `query:"subscription_id"`
}

// URLQuery serializes [InvoiceFetchUpcomingParams]'s query parameters as
// `url.Values`.
func (r InvoiceFetchUpcomingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceMarkPaidParams struct {
	// An optional external ID to associate with the payment.
	ExternalID param.Field[string] `json:"external_id,required"`
	// An optional note to associate with the payment.
	Notes param.Field[string] `json:"notes,required"`
	// A date string to specify the date of the payment.
	PaymentReceivedDate param.Field[time.Time] `json:"payment_received_date,required" format:"date"`
}

func (r InvoiceMarkPaidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
