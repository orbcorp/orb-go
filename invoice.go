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
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
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
func (r *InvoiceService) Update(ctx context.Context, invoiceID string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
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
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *pagination.Page[shared.InvoiceModel], err error) {
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
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.InvoiceModel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch an [`Invoice`](/core-concepts#invoice) given an
// identifier.
func (r *InvoiceService) Fetch(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
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
func (r *InvoiceService) Issue(ctx context.Context, invoiceID string, body InvoiceIssueParams, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
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
func (r *InvoiceService) MarkPaid(ctx context.Context, invoiceID string, body InvoiceMarkPaidParams, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
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
func (r *InvoiceService) Pay(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
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
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *shared.InvoiceModel, err error) {
	opts = append(r.Options[:], opts...)
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
	AmountDue      string                     `json:"amount_due,required"`
	AutoCollection shared.AutoCollectionModel `json:"auto_collection,required"`
	BillingAddress shared.AddressModel        `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []shared.CreditNoteSummaryModel `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                                   `json:"currency,required"`
	Customer                    shared.CustomerMinifiedModel             `json:"customer,required"`
	CustomerBalanceTransactions []shared.CustomerBalanceTransactionModel `json:"customer_balance_transactions,required"`
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
	CustomerTaxID shared.CustomerTaxIDModel `json:"customer_tax_id,required,nullable"`
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
	LineItems     []shared.InvoiceLineItemModel `json:"line_items,required"`
	Maximum       shared.MaximumModel           `json:"maximum,required,nullable"`
	MaximumAmount string                        `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string   `json:"metadata,required"`
	Minimum       shared.MinimumModel `json:"minimum,required,nullable"`
	MinimumAmount string              `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// A list of payment attempts associated with the invoice
	PaymentAttempts []shared.PaymentAttemptModel `json:"payment_attempts,required"`
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
	ShippingAddress  shared.AddressModel                `json:"shipping_address,required,nullable"`
	Status           InvoiceFetchUpcomingResponseStatus `json:"status,required"`
	Subscription     shared.SubscriptionMinifiedModel   `json:"subscription,required,nullable"`
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
	StartDate  param.Field[time.Time]                   `json:"start_date,required" format:"date"`
	UnitConfig param.Field[shared.UnitConfigModelParam] `json:"unit_config,required"`
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
