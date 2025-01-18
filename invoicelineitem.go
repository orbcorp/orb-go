// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// InvoiceLineItemService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceLineItemService] method instead.
type InvoiceLineItemService struct {
	Options []option.RequestOption
}

// NewInvoiceLineItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvoiceLineItemService(opts ...option.RequestOption) (r *InvoiceLineItemService) {
	r = &InvoiceLineItemService{}
	r.Options = opts
	return
}

// This creates a one-off fixed fee invoice line item on an Invoice. This can only
// be done for invoices that are in a `draft` status.
func (r *InvoiceLineItemService) New(ctx context.Context, body InvoiceLineItemNewParams, opts ...option.RequestOption) (res *InvoiceLineItemNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "invoice_line_items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InvoiceLineItemNewResponse struct {
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
	Grouping      string                            `json:"grouping,required,nullable"`
	Maximum       InvoiceLineItemNewResponseMaximum `json:"maximum,required,nullable"`
	MaximumAmount string                            `json:"maximum_amount,required,nullable"`
	Minimum       InvoiceLineItemNewResponseMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                            `json:"minimum_amount,required,nullable"`
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
	// For more on the types of prices, see
	// [the core concepts documentation](/core-concepts#plan-and-price)
	Price    Price   `json:"price,required,nullable"`
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemNewResponseSubLineItem `json:"sub_line_items,required"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceLineItemNewResponseTaxAmount `json:"tax_amounts,required"`
	JSON       invoiceLineItemNewResponseJSON        `json:"-"`
}

// invoiceLineItemNewResponseJSON contains the JSON metadata for the struct
// [InvoiceLineItemNewResponse]
type invoiceLineItemNewResponseJSON struct {
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

func (r *InvoiceLineItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                `json:"maximum_amount,required"`
	JSON          invoiceLineItemNewResponseMaximumJSON `json:"-"`
}

// invoiceLineItemNewResponseMaximumJSON contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseMaximum]
type invoiceLineItemNewResponseMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseMaximumJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                `json:"minimum_amount,required"`
	JSON          invoiceLineItemNewResponseMinimumJSON `json:"-"`
}

// invoiceLineItemNewResponseMinimumJSON contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseMinimum]
type invoiceLineItemNewResponseMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseMinimumJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseSubLineItem struct {
	// The total amount for this sub line item.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of
	// [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGrouping],
	// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItemGrouping],
	// [InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemGrouping].
	Grouping interface{}                                `json:"grouping,required"`
	Name     string                                     `json:"name,required"`
	Quantity float64                                    `json:"quantity,required"`
	Type     InvoiceLineItemNewResponseSubLineItemsType `json:"type,required"`
	// This field can have the runtime type of
	// [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfig].
	MatrixConfig interface{} `json:"matrix_config"`
	// This field can have the runtime type of
	// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfig].
	TierConfig interface{}                               `json:"tier_config"`
	JSON       invoiceLineItemNewResponseSubLineItemJSON `json:"-"`
	union      InvoiceLineItemNewResponseSubLineItemsUnion
}

// invoiceLineItemNewResponseSubLineItemJSON contains the JSON metadata for the
// struct [InvoiceLineItemNewResponseSubLineItem]
type invoiceLineItemNewResponseSubLineItemJSON struct {
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

func (r invoiceLineItemNewResponseSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemNewResponseSubLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemNewResponseSubLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemNewResponseSubLineItemsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItem],
// [InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem].
func (r InvoiceLineItemNewResponseSubLineItem) AsUnion() InvoiceLineItemNewResponseSubLineItemsUnion {
	return r.union
}

// Union satisfied by [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItem] or
// [InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem].
type InvoiceLineItemNewResponseSubLineItemsUnion interface {
	implementsInvoiceLineItemNewResponseSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemNewResponseSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseSubLineItemsTierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                              `json:"amount,required"`
	Grouping     InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGrouping     `json:"grouping,required,nullable"`
	MatrixConfig InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                                                              `json:"name,required"`
	Quantity     float64                                                             `json:"quantity,required"`
	Type         InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemType         `json:"type,required"`
	JSON         invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemJSON         `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem]
type invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem) implementsInvoiceLineItemNewResponseSubLineItem() {
}

type InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                              `json:"value,required,nullable"`
	JSON  invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGroupingJSON `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGroupingJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGrouping]
type invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string                                                                `json:"dimension_values,required"`
	JSON            invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfigJSON `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfigJSON contains
// the JSON metadata for the struct
// [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfig]
type invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsMatrixSubLineItemMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemType string

const (
	InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemTypeMatrix InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemType = "matrix"
)

func (r InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemTypeMatrix:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseSubLineItemsTierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string                                                          `json:"amount,required"`
	Grouping   InvoiceLineItemNewResponseSubLineItemsTierSubLineItemGrouping   `json:"grouping,required,nullable"`
	Name       string                                                          `json:"name,required"`
	Quantity   float64                                                         `json:"quantity,required"`
	TierConfig InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfig `json:"tier_config,required"`
	Type       InvoiceLineItemNewResponseSubLineItemsTierSubLineItemType       `json:"type,required"`
	JSON       invoiceLineItemNewResponseSubLineItemsTierSubLineItemJSON       `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsTierSubLineItemJSON contains the JSON
// metadata for the struct [InvoiceLineItemNewResponseSubLineItemsTierSubLineItem]
type invoiceLineItemNewResponseSubLineItemsTierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsTierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsTierSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseSubLineItemsTierSubLineItem) implementsInvoiceLineItemNewResponseSubLineItem() {
}

type InvoiceLineItemNewResponseSubLineItemsTierSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                            `json:"value,required,nullable"`
	JSON  invoiceLineItemNewResponseSubLineItemsTierSubLineItemGroupingJSON `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsTierSubLineItemGroupingJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItemGrouping]
type invoiceLineItemNewResponseSubLineItemsTierSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsTierSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsTierSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfig struct {
	FirstUnit  float64                                                             `json:"first_unit,required"`
	LastUnit   float64                                                             `json:"last_unit,required,nullable"`
	UnitAmount string                                                              `json:"unit_amount,required"`
	JSON       invoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfigJSON `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfigJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfig]
type invoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsTierSubLineItemTierConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseSubLineItemsTierSubLineItemType string

const (
	InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTypeTier InvoiceLineItemNewResponseSubLineItemsTierSubLineItemType = "tier"
)

func (r InvoiceLineItemNewResponseSubLineItemsTierSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTypeTier:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                                                         `json:"amount,required"`
	Grouping InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemGrouping `json:"grouping,required,nullable"`
	Name     string                                                         `json:"name,required"`
	Quantity float64                                                        `json:"quantity,required"`
	Type     InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemType     `json:"type,required"`
	JSON     invoiceLineItemNewResponseSubLineItemsOtherSubLineItemJSON     `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsOtherSubLineItemJSON contains the JSON
// metadata for the struct [InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem]
type invoiceLineItemNewResponseSubLineItemsOtherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsOtherSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem) implementsInvoiceLineItemNewResponseSubLineItem() {
}

type InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                                                             `json:"value,required,nullable"`
	JSON  invoiceLineItemNewResponseSubLineItemsOtherSubLineItemGroupingJSON `json:"-"`
}

// invoiceLineItemNewResponseSubLineItemsOtherSubLineItemGroupingJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemGrouping]
type invoiceLineItemNewResponseSubLineItemsOtherSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseSubLineItemsOtherSubLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemType string

const (
	InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemTypeNull InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemType = "'null'"
)

func (r InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemTypeNull:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseSubLineItemsType string

const (
	InvoiceLineItemNewResponseSubLineItemsTypeMatrix InvoiceLineItemNewResponseSubLineItemsType = "matrix"
	InvoiceLineItemNewResponseSubLineItemsTypeTier   InvoiceLineItemNewResponseSubLineItemsType = "tier"
	InvoiceLineItemNewResponseSubLineItemsTypeNull   InvoiceLineItemNewResponseSubLineItemsType = "'null'"
)

func (r InvoiceLineItemNewResponseSubLineItemsType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseSubLineItemsTypeMatrix, InvoiceLineItemNewResponseSubLineItemsTypeTier, InvoiceLineItemNewResponseSubLineItemsTypeNull:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseTaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string                                  `json:"tax_rate_percentage,required,nullable"`
	JSON              invoiceLineItemNewResponseTaxAmountJSON `json:"-"`
}

// invoiceLineItemNewResponseTaxAmountJSON contains the JSON metadata for the
// struct [InvoiceLineItemNewResponseTaxAmount]
type invoiceLineItemNewResponseTaxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseTaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseTaxAmountJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewParams struct {
	// The total amount in the invoice's currency to add to the line item.
	Amount param.Field[string] `json:"amount,required"`
	// A date string to specify the line item's end date in the customer's timezone.
	EndDate param.Field[time.Time] `json:"end_date,required" format:"date"`
	// The id of the Invoice to add this line item.
	InvoiceID param.Field[string] `json:"invoice_id,required"`
	// The item name associated with this line item. If an item with the same name
	// exists in Orb, that item will be associated with the line item.
	Name param.Field[string] `json:"name,required"`
	// The number of units on the line item
	Quantity param.Field[float64] `json:"quantity,required"`
	// A date string to specify the line item's start date in the customer's timezone.
	StartDate param.Field[time.Time] `json:"start_date,required" format:"date"`
}

func (r InvoiceLineItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
