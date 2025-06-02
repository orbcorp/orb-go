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
	// The line amount after any adjustments and before overage conversion, credits and
	// partial invoicing.
	AdjustedSubtotal string `json:"adjusted_subtotal,required"`
	// All adjustments applied to the line item in the order they were applied based on
	// invoice calculations (ie. usage discounts -> amount discounts -> percentage
	// discounts -> minimums -> maximums).
	Adjustments []InvoiceLineItemNewResponseAdjustment `json:"adjustments,required"`
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
	Maximum InvoiceLineItemNewResponseMaximum `json:"maximum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Minimum InvoiceLineItemNewResponseMinimum `json:"minimum,required,nullable"`
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
	Price Price `json:"price,required"`
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemNewResponseSubLineItem `json:"sub_line_items,required"`
	// The line amount before before any adjustments.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceLineItemNewResponseTaxAmount `json:"tax_amounts,required"`
	// A list of customer ids that were used to calculate the usage for this line item.
	UsageCustomerIDs []string                       `json:"usage_customer_ids,required,nullable"`
	JSON             invoiceLineItemNewResponseJSON `json:"-"`
}

// invoiceLineItemNewResponseJSON contains the JSON metadata for the struct
// [InvoiceLineItemNewResponse]
type invoiceLineItemNewResponseJSON struct {
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

func (r *InvoiceLineItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemNewResponseAdjustment struct {
	ID             string                                              `json:"id,required"`
	AdjustmentType InvoiceLineItemNewResponseAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilter],
	// [[]InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilter],
	// [[]InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilter],
	// [[]InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilter],
	// [[]InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilter].
	Filters interface{} `json:"filters,required"`
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
	UsageDiscount float64                                  `json:"usage_discount"`
	JSON          invoiceLineItemNewResponseAdjustmentJSON `json:"-"`
	union         InvoiceLineItemNewResponseAdjustmentsUnion
}

// invoiceLineItemNewResponseAdjustmentJSON contains the JSON metadata for the
// struct [InvoiceLineItemNewResponseAdjustment]
type invoiceLineItemNewResponseAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
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

func (r invoiceLineItemNewResponseAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemNewResponseAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemNewResponseAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemNewResponseAdjustmentsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment].
func (r InvoiceLineItemNewResponseAdjustment) AsUnion() InvoiceLineItemNewResponseAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment],
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment] or
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment].
type InvoiceLineItemNewResponseAdjustmentsUnion interface {
	implementsInvoiceLineItemNewResponseAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemNewResponseAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment struct {
	ID             string                                                                             `json:"id,required"`
	AdjustmentType InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                  `json:"usage_discount,required"`
	JSON          invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment]
type invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustment) implementsInvoiceLineItemNewResponseAdjustment() {
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                       `json:"values,required"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilter]
type invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldPriceID       InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField = "price_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldItemID        InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField = "item_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldPriceType     InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField = "price_type"
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldCurrency      InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField = "currency"
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldPricingUnitID InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldPriceID, InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldItemID, InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldPriceType, InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldCurrency, InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperator string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperatorIncludes InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperator = "includes"
	InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperatorExcludes InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperatorIncludes, InvoiceLineItemNewResponseAdjustmentsMonetaryUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment struct {
	ID             string                                                                              `json:"id,required"`
	AdjustmentType InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string                                                                    `json:"reason,required,nullable"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment]
type invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustment) implementsInvoiceLineItemNewResponseAdjustment() {
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                        `json:"values,required"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilter]
type invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldPriceID       InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField = "price_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldItemID        InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField = "item_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldPriceType     InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField = "price_type"
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldCurrency      InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField = "currency"
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldPricingUnitID InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldPriceID, InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldItemID, InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldPriceType, InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldCurrency, InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperator string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperatorIncludes InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperator = "includes"
	InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperatorExcludes InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperatorIncludes, InvoiceLineItemNewResponseAdjustmentsMonetaryAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The reason for the adjustment.
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment]
type invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustment) implementsInvoiceLineItemNewResponseAdjustment() {
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                            `json:"values,required"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilter]
type invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldPriceID       InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField = "price_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldItemID        InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField = "item_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldPriceType     InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField = "price_type"
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldCurrency      InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField = "currency"
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldPricingUnitID InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldPriceID, InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldItemID, InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldPriceType, InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldCurrency, InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperator string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperatorIncludes InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperator = "includes"
	InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperatorExcludes InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperatorIncludes, InvoiceLineItemNewResponseAdjustmentsMonetaryPercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment struct {
	ID             string                                                                       `json:"id,required"`
	AdjustmentType InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                             `json:"reason,required,nullable"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment]
type invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustment) implementsInvoiceLineItemNewResponseAdjustment() {
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentAdjustmentType string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentAdjustmentType = "minimum"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                 `json:"values,required"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilter]
type invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldPriceID       InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField = "price_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldItemID        InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField = "item_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldPriceType     InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField = "price_type"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldCurrency      InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField = "currency"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldPricingUnitID InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldPriceID, InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldItemID, InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldPriceType, InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldCurrency, InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperator string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperatorIncludes InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperator = "includes"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperatorExcludes InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperatorIncludes, InvoiceLineItemNewResponseAdjustmentsMonetaryMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment struct {
	ID             string                                                                       `json:"id,required"`
	AdjustmentType InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                             `json:"reason,required,nullable"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment]
type invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustment) implementsInvoiceLineItemNewResponseAdjustment() {
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentAdjustmentType string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentAdjustmentType = "maximum"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                 `json:"values,required"`
	JSON   invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilter]
type invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldPriceID       InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField = "price_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldItemID        InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField = "item_id"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldPriceType     InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField = "price_type"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldCurrency      InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField = "currency"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldPricingUnitID InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldPriceID, InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldItemID, InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldPriceType, InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldCurrency, InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperator string

const (
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperatorIncludes InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperator = "includes"
	InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperatorExcludes InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperatorIncludes, InvoiceLineItemNewResponseAdjustmentsMonetaryMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type InvoiceLineItemNewResponseAdjustmentsAdjustmentType string

const (
	InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeUsageDiscount      InvoiceLineItemNewResponseAdjustmentsAdjustmentType = "usage_discount"
	InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeAmountDiscount     InvoiceLineItemNewResponseAdjustmentsAdjustmentType = "amount_discount"
	InvoiceLineItemNewResponseAdjustmentsAdjustmentTypePercentageDiscount InvoiceLineItemNewResponseAdjustmentsAdjustmentType = "percentage_discount"
	InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeMinimum            InvoiceLineItemNewResponseAdjustmentsAdjustmentType = "minimum"
	InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeMaximum            InvoiceLineItemNewResponseAdjustmentsAdjustmentType = "maximum"
)

func (r InvoiceLineItemNewResponseAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeUsageDiscount, InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeAmountDiscount, InvoiceLineItemNewResponseAdjustmentsAdjustmentTypePercentageDiscount, InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeMinimum, InvoiceLineItemNewResponseAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

// This field is deprecated in favor of `adjustments`.
//
// Deprecated: deprecated
type InvoiceLineItemNewResponseMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this maximum to.
	Filters []InvoiceLineItemNewResponseMaximumFilter `json:"filters,required"`
	// Maximum amount applied
	MaximumAmount string                                `json:"maximum_amount,required"`
	JSON          invoiceLineItemNewResponseMaximumJSON `json:"-"`
}

// invoiceLineItemNewResponseMaximumJSON contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseMaximum]
type invoiceLineItemNewResponseMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
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

type InvoiceLineItemNewResponseMaximumFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseMaximumFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseMaximumFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                    `json:"values,required"`
	JSON   invoiceLineItemNewResponseMaximumFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseMaximumFilterJSON contains the JSON metadata for the
// struct [InvoiceLineItemNewResponseMaximumFilter]
type invoiceLineItemNewResponseMaximumFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseMaximumFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseMaximumFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseMaximumFiltersField string

const (
	InvoiceLineItemNewResponseMaximumFiltersFieldPriceID       InvoiceLineItemNewResponseMaximumFiltersField = "price_id"
	InvoiceLineItemNewResponseMaximumFiltersFieldItemID        InvoiceLineItemNewResponseMaximumFiltersField = "item_id"
	InvoiceLineItemNewResponseMaximumFiltersFieldPriceType     InvoiceLineItemNewResponseMaximumFiltersField = "price_type"
	InvoiceLineItemNewResponseMaximumFiltersFieldCurrency      InvoiceLineItemNewResponseMaximumFiltersField = "currency"
	InvoiceLineItemNewResponseMaximumFiltersFieldPricingUnitID InvoiceLineItemNewResponseMaximumFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseMaximumFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseMaximumFiltersFieldPriceID, InvoiceLineItemNewResponseMaximumFiltersFieldItemID, InvoiceLineItemNewResponseMaximumFiltersFieldPriceType, InvoiceLineItemNewResponseMaximumFiltersFieldCurrency, InvoiceLineItemNewResponseMaximumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseMaximumFiltersOperator string

const (
	InvoiceLineItemNewResponseMaximumFiltersOperatorIncludes InvoiceLineItemNewResponseMaximumFiltersOperator = "includes"
	InvoiceLineItemNewResponseMaximumFiltersOperatorExcludes InvoiceLineItemNewResponseMaximumFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseMaximumFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseMaximumFiltersOperatorIncludes, InvoiceLineItemNewResponseMaximumFiltersOperatorExcludes:
		return true
	}
	return false
}

// This field is deprecated in favor of `adjustments`.
//
// Deprecated: deprecated
type InvoiceLineItemNewResponseMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this minimum to.
	Filters []InvoiceLineItemNewResponseMinimumFilter `json:"filters,required"`
	// Minimum amount applied
	MinimumAmount string                                `json:"minimum_amount,required"`
	JSON          invoiceLineItemNewResponseMinimumJSON `json:"-"`
}

// invoiceLineItemNewResponseMinimumJSON contains the JSON metadata for the struct
// [InvoiceLineItemNewResponseMinimum]
type invoiceLineItemNewResponseMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
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

type InvoiceLineItemNewResponseMinimumFilter struct {
	// The property of the price to filter on.
	Field InvoiceLineItemNewResponseMinimumFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceLineItemNewResponseMinimumFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                    `json:"values,required"`
	JSON   invoiceLineItemNewResponseMinimumFilterJSON `json:"-"`
}

// invoiceLineItemNewResponseMinimumFilterJSON contains the JSON metadata for the
// struct [InvoiceLineItemNewResponseMinimumFilter]
type invoiceLineItemNewResponseMinimumFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemNewResponseMinimumFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemNewResponseMinimumFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceLineItemNewResponseMinimumFiltersField string

const (
	InvoiceLineItemNewResponseMinimumFiltersFieldPriceID       InvoiceLineItemNewResponseMinimumFiltersField = "price_id"
	InvoiceLineItemNewResponseMinimumFiltersFieldItemID        InvoiceLineItemNewResponseMinimumFiltersField = "item_id"
	InvoiceLineItemNewResponseMinimumFiltersFieldPriceType     InvoiceLineItemNewResponseMinimumFiltersField = "price_type"
	InvoiceLineItemNewResponseMinimumFiltersFieldCurrency      InvoiceLineItemNewResponseMinimumFiltersField = "currency"
	InvoiceLineItemNewResponseMinimumFiltersFieldPricingUnitID InvoiceLineItemNewResponseMinimumFiltersField = "pricing_unit_id"
)

func (r InvoiceLineItemNewResponseMinimumFiltersField) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseMinimumFiltersFieldPriceID, InvoiceLineItemNewResponseMinimumFiltersFieldItemID, InvoiceLineItemNewResponseMinimumFiltersFieldPriceType, InvoiceLineItemNewResponseMinimumFiltersFieldCurrency, InvoiceLineItemNewResponseMinimumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceLineItemNewResponseMinimumFiltersOperator string

const (
	InvoiceLineItemNewResponseMinimumFiltersOperatorIncludes InvoiceLineItemNewResponseMinimumFiltersOperator = "includes"
	InvoiceLineItemNewResponseMinimumFiltersOperatorExcludes InvoiceLineItemNewResponseMinimumFiltersOperator = "excludes"
)

func (r InvoiceLineItemNewResponseMinimumFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceLineItemNewResponseMinimumFiltersOperatorIncludes, InvoiceLineItemNewResponseMinimumFiltersOperatorExcludes:
		return true
	}
	return false
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
