// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// InvoiceLineItemService contains methods and other services that help with
// interacting with the orb API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewInvoiceLineItemService] method
// instead.
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

// Union satisfied by [InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemNewResponseSubLineItemsTierSubLineItem] or
// [InvoiceLineItemNewResponseSubLineItemsOtherSubLineItem].
type InvoiceLineItemNewResponseSubLineItem interface {
	implementsInvoiceLineItemNewResponseSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemNewResponseSubLineItem)(nil)).Elem(),
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

type InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemType string

const (
	InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemTypeMatrix InvoiceLineItemNewResponseSubLineItemsMatrixSubLineItemType = "matrix"
)

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

type InvoiceLineItemNewResponseSubLineItemsTierSubLineItemType string

const (
	InvoiceLineItemNewResponseSubLineItemsTierSubLineItemTypeTier InvoiceLineItemNewResponseSubLineItemsTierSubLineItemType = "tier"
)

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

type InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemType string

const (
	InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemTypeNull InvoiceLineItemNewResponseSubLineItemsOtherSubLineItemType = "'null'"
)

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
