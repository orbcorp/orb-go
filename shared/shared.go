// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/tidwall/gjson"
)

type AmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string                     `json:"amount_discount,required"`
	DiscountType   AmountDiscountDiscountType `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []AmountDiscountFilter `json:"filters,nullable"`
	Reason  string                 `json:"reason,nullable"`
	JSON    amountDiscountJSON     `json:"-"`
}

// amountDiscountJSON contains the JSON metadata for the struct [AmountDiscount]
type amountDiscountJSON struct {
	AmountDiscount    apijson.Field
	DiscountType      apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AmountDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r amountDiscountJSON) RawJSON() string {
	return r.raw
}

func (r AmountDiscount) ImplementsDiscount() {}

func (r AmountDiscount) ImplementsInvoiceLevelDiscount() {}

func (r AmountDiscount) ImplementsCouponDiscount() {}

type AmountDiscountDiscountType string

const (
	AmountDiscountDiscountTypeAmount AmountDiscountDiscountType = "amount"
)

func (r AmountDiscountDiscountType) IsKnown() bool {
	switch r {
	case AmountDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type AmountDiscountFilter struct {
	// The property of the price to filter on.
	Field AmountDiscountFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator AmountDiscountFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                 `json:"values,required"`
	JSON   amountDiscountFilterJSON `json:"-"`
}

// amountDiscountFilterJSON contains the JSON metadata for the struct
// [AmountDiscountFilter]
type amountDiscountFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AmountDiscountFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r amountDiscountFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type AmountDiscountFiltersField string

const (
	AmountDiscountFiltersFieldPriceID       AmountDiscountFiltersField = "price_id"
	AmountDiscountFiltersFieldItemID        AmountDiscountFiltersField = "item_id"
	AmountDiscountFiltersFieldPriceType     AmountDiscountFiltersField = "price_type"
	AmountDiscountFiltersFieldCurrency      AmountDiscountFiltersField = "currency"
	AmountDiscountFiltersFieldPricingUnitID AmountDiscountFiltersField = "pricing_unit_id"
)

func (r AmountDiscountFiltersField) IsKnown() bool {
	switch r {
	case AmountDiscountFiltersFieldPriceID, AmountDiscountFiltersFieldItemID, AmountDiscountFiltersFieldPriceType, AmountDiscountFiltersFieldCurrency, AmountDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type AmountDiscountFiltersOperator string

const (
	AmountDiscountFiltersOperatorIncludes AmountDiscountFiltersOperator = "includes"
	AmountDiscountFiltersOperatorExcludes AmountDiscountFiltersOperator = "excludes"
)

func (r AmountDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case AmountDiscountFiltersOperatorIncludes, AmountDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

type AmountDiscountParam struct {
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string]                     `json:"amount_discount,required"`
	DiscountType   param.Field[AmountDiscountDiscountType] `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]AmountDiscountFilterParam] `json:"filters"`
	Reason  param.Field[string]                      `json:"reason"`
}

func (r AmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AmountDiscountParam) ImplementsDiscountUnionParam() {}

type AmountDiscountFilterParam struct {
	// The property of the price to filter on.
	Field param.Field[AmountDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[AmountDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r AmountDiscountFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingCycleRelativeDate string

const (
	BillingCycleRelativeDateStartOfTerm BillingCycleRelativeDate = "start_of_term"
	BillingCycleRelativeDateEndOfTerm   BillingCycleRelativeDate = "end_of_term"
)

func (r BillingCycleRelativeDate) IsKnown() bool {
	switch r {
	case BillingCycleRelativeDateStartOfTerm, BillingCycleRelativeDateEndOfTerm:
		return true
	}
	return false
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddStartDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddEndDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion() {
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion() {
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditEndDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditStartDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion() {
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion() {
}

type Discount struct {
	DiscountType DiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	// This field can have the runtime type of [[]PercentageDiscountFilter],
	// [[]TrialDiscountFilter], [[]UsageDiscountFilter], [[]AmountDiscountFilter].
	Filters interface{} `json:"filters"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	Reason             string  `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64 `json:"trial_percentage_discount,nullable"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64      `json:"usage_discount"`
	JSON          discountJSON `json:"-"`
	union         DiscountUnion
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
	AppliesToPriceIDs       apijson.Field
	Filters                 apijson.Field
	PercentageDiscount      apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	UsageDiscount           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r discountJSON) RawJSON() string {
	return r.raw
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
	*r = Discount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DiscountUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [PercentageDiscount], [TrialDiscount],
// [UsageDiscount], [AmountDiscount].
func (r Discount) AsUnion() DiscountUnion {
	return r.union
}

// Union satisfied by [PercentageDiscount], [TrialDiscount], [UsageDiscount] or
// [AmountDiscount].
type DiscountUnion interface {
	ImplementsDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DiscountUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PercentageDiscount{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TrialDiscount{}),
			DiscriminatorValue: "trial",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UsageDiscount{}),
			DiscriminatorValue: "usage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscount{}),
			DiscriminatorValue: "amount",
		},
	)
}

type DiscountDiscountType string

const (
	DiscountDiscountTypePercentage DiscountDiscountType = "percentage"
	DiscountDiscountTypeTrial      DiscountDiscountType = "trial"
	DiscountDiscountTypeUsage      DiscountDiscountType = "usage"
	DiscountDiscountTypeAmount     DiscountDiscountType = "amount"
)

func (r DiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountDiscountTypePercentage, DiscountDiscountTypeTrial, DiscountDiscountTypeUsage, DiscountDiscountTypeAmount:
		return true
	}
	return false
}

type DiscountParam struct {
	DiscountType param.Field[DiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount    param.Field[string]      `json:"amount_discount"`
	AppliesToPriceIDs param.Field[interface{}] `json:"applies_to_price_ids"`
	Filters           param.Field[interface{}] `json:"filters"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	Reason             param.Field[string]  `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r DiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscountParam) ImplementsDiscountUnionParam() {}

// Satisfied by [shared.PercentageDiscountParam], [shared.TrialDiscountParam],
// [shared.UsageDiscountParam], [shared.AmountDiscountParam], [DiscountParam].
type DiscountUnionParam interface {
	ImplementsDiscountUnionParam()
}

type InvoiceLevelDiscount struct {
	DiscountType InvoiceLevelDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	// This field can have the runtime type of [[]PercentageDiscountFilter],
	// [[]AmountDiscountFilter], [[]TrialDiscountFilter].
	Filters interface{} `json:"filters"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	Reason             string  `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64                  `json:"trial_percentage_discount,nullable"`
	JSON                    invoiceLevelDiscountJSON `json:"-"`
	union                   InvoiceLevelDiscountUnion
}

// invoiceLevelDiscountJSON contains the JSON metadata for the struct
// [InvoiceLevelDiscount]
type invoiceLevelDiscountJSON struct {
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
	AppliesToPriceIDs       apijson.Field
	Filters                 apijson.Field
	PercentageDiscount      apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r invoiceLevelDiscountJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLevelDiscount) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLevelDiscount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLevelDiscountUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [PercentageDiscount], [AmountDiscount],
// [TrialDiscount].
func (r InvoiceLevelDiscount) AsUnion() InvoiceLevelDiscountUnion {
	return r.union
}

// Union satisfied by [PercentageDiscount], [AmountDiscount] or [TrialDiscount].
type InvoiceLevelDiscountUnion interface {
	ImplementsInvoiceLevelDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLevelDiscountUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PercentageDiscount{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscount{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TrialDiscount{}),
			DiscriminatorValue: "trial",
		},
	)
}

type InvoiceLevelDiscountDiscountType string

const (
	InvoiceLevelDiscountDiscountTypePercentage InvoiceLevelDiscountDiscountType = "percentage"
	InvoiceLevelDiscountDiscountTypeAmount     InvoiceLevelDiscountDiscountType = "amount"
	InvoiceLevelDiscountDiscountTypeTrial      InvoiceLevelDiscountDiscountType = "trial"
)

func (r InvoiceLevelDiscountDiscountType) IsKnown() bool {
	switch r {
	case InvoiceLevelDiscountDiscountTypePercentage, InvoiceLevelDiscountDiscountTypeAmount, InvoiceLevelDiscountDiscountTypeTrial:
		return true
	}
	return false
}

type PaginationMetadata struct {
	HasMore    bool                   `json:"has_more,required"`
	NextCursor string                 `json:"next_cursor,required,nullable"`
	JSON       paginationMetadataJSON `json:"-"`
}

// paginationMetadataJSON contains the JSON metadata for the struct
// [PaginationMetadata]
type paginationMetadataJSON struct {
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginationMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginationMetadataJSON) RawJSON() string {
	return r.raw
}

type PercentageDiscount struct {
	DiscountType PercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []PercentageDiscountFilter `json:"filters,nullable"`
	Reason  string                     `json:"reason,nullable"`
	JSON    percentageDiscountJSON     `json:"-"`
}

// percentageDiscountJSON contains the JSON metadata for the struct
// [PercentageDiscount]
type percentageDiscountJSON struct {
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PercentageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r PercentageDiscount) ImplementsDiscount() {}

func (r PercentageDiscount) ImplementsInvoiceLevelDiscount() {}

func (r PercentageDiscount) ImplementsCouponDiscount() {}

type PercentageDiscountDiscountType string

const (
	PercentageDiscountDiscountTypePercentage PercentageDiscountDiscountType = "percentage"
)

func (r PercentageDiscountDiscountType) IsKnown() bool {
	switch r {
	case PercentageDiscountDiscountTypePercentage:
		return true
	}
	return false
}

type PercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field PercentageDiscountFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PercentageDiscountFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                     `json:"values,required"`
	JSON   percentageDiscountFilterJSON `json:"-"`
}

// percentageDiscountFilterJSON contains the JSON metadata for the struct
// [PercentageDiscountFilter]
type percentageDiscountFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PercentageDiscountFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentageDiscountFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PercentageDiscountFiltersField string

const (
	PercentageDiscountFiltersFieldPriceID       PercentageDiscountFiltersField = "price_id"
	PercentageDiscountFiltersFieldItemID        PercentageDiscountFiltersField = "item_id"
	PercentageDiscountFiltersFieldPriceType     PercentageDiscountFiltersField = "price_type"
	PercentageDiscountFiltersFieldCurrency      PercentageDiscountFiltersField = "currency"
	PercentageDiscountFiltersFieldPricingUnitID PercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PercentageDiscountFiltersFieldPriceID, PercentageDiscountFiltersFieldItemID, PercentageDiscountFiltersFieldPriceType, PercentageDiscountFiltersFieldCurrency, PercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PercentageDiscountFiltersOperator string

const (
	PercentageDiscountFiltersOperatorIncludes PercentageDiscountFiltersOperator = "includes"
	PercentageDiscountFiltersOperatorExcludes PercentageDiscountFiltersOperator = "excludes"
)

func (r PercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PercentageDiscountFiltersOperatorIncludes, PercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

type PercentageDiscountParam struct {
	DiscountType param.Field[PercentageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]PercentageDiscountFilterParam] `json:"filters"`
	Reason  param.Field[string]                          `json:"reason"`
}

func (r PercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PercentageDiscountParam) ImplementsDiscountUnionParam() {}

type PercentageDiscountFilterParam struct {
	// The property of the price to filter on.
	Field param.Field[PercentageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PercentageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r PercentageDiscountFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TrialDiscount struct {
	DiscountType TrialDiscountDiscountType `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []TrialDiscountFilter `json:"filters,nullable"`
	Reason  string                `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64           `json:"trial_percentage_discount,nullable"`
	JSON                    trialDiscountJSON `json:"-"`
}

// trialDiscountJSON contains the JSON metadata for the struct [TrialDiscount]
type trialDiscountJSON struct {
	DiscountType            apijson.Field
	AppliesToPriceIDs       apijson.Field
	Filters                 apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TrialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trialDiscountJSON) RawJSON() string {
	return r.raw
}

func (r TrialDiscount) ImplementsDiscount() {}

func (r TrialDiscount) ImplementsInvoiceLevelDiscount() {}

type TrialDiscountDiscountType string

const (
	TrialDiscountDiscountTypeTrial TrialDiscountDiscountType = "trial"
)

func (r TrialDiscountDiscountType) IsKnown() bool {
	switch r {
	case TrialDiscountDiscountTypeTrial:
		return true
	}
	return false
}

type TrialDiscountFilter struct {
	// The property of the price to filter on.
	Field TrialDiscountFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator TrialDiscountFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                `json:"values,required"`
	JSON   trialDiscountFilterJSON `json:"-"`
}

// trialDiscountFilterJSON contains the JSON metadata for the struct
// [TrialDiscountFilter]
type trialDiscountFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrialDiscountFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trialDiscountFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type TrialDiscountFiltersField string

const (
	TrialDiscountFiltersFieldPriceID       TrialDiscountFiltersField = "price_id"
	TrialDiscountFiltersFieldItemID        TrialDiscountFiltersField = "item_id"
	TrialDiscountFiltersFieldPriceType     TrialDiscountFiltersField = "price_type"
	TrialDiscountFiltersFieldCurrency      TrialDiscountFiltersField = "currency"
	TrialDiscountFiltersFieldPricingUnitID TrialDiscountFiltersField = "pricing_unit_id"
)

func (r TrialDiscountFiltersField) IsKnown() bool {
	switch r {
	case TrialDiscountFiltersFieldPriceID, TrialDiscountFiltersFieldItemID, TrialDiscountFiltersFieldPriceType, TrialDiscountFiltersFieldCurrency, TrialDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type TrialDiscountFiltersOperator string

const (
	TrialDiscountFiltersOperatorIncludes TrialDiscountFiltersOperator = "includes"
	TrialDiscountFiltersOperatorExcludes TrialDiscountFiltersOperator = "excludes"
)

func (r TrialDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case TrialDiscountFiltersOperatorIncludes, TrialDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

type TrialDiscountParam struct {
	DiscountType param.Field[TrialDiscountDiscountType] `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]TrialDiscountFilterParam] `json:"filters"`
	Reason  param.Field[string]                     `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
}

func (r TrialDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrialDiscountParam) ImplementsDiscountUnionParam() {}

type TrialDiscountFilterParam struct {
	// The property of the price to filter on.
	Field param.Field[TrialDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[TrialDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r TrialDiscountFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageDiscount struct {
	DiscountType UsageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64 `json:"usage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []UsageDiscountFilter `json:"filters,nullable"`
	Reason  string                `json:"reason,nullable"`
	JSON    usageDiscountJSON     `json:"-"`
}

// usageDiscountJSON contains the JSON metadata for the struct [UsageDiscount]
type usageDiscountJSON struct {
	DiscountType      apijson.Field
	UsageDiscount     apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UsageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r UsageDiscount) ImplementsDiscount() {}

type UsageDiscountDiscountType string

const (
	UsageDiscountDiscountTypeUsage UsageDiscountDiscountType = "usage"
)

func (r UsageDiscountDiscountType) IsKnown() bool {
	switch r {
	case UsageDiscountDiscountTypeUsage:
		return true
	}
	return false
}

type UsageDiscountFilter struct {
	// The property of the price to filter on.
	Field UsageDiscountFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator UsageDiscountFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                `json:"values,required"`
	JSON   usageDiscountFilterJSON `json:"-"`
}

// usageDiscountFilterJSON contains the JSON metadata for the struct
// [UsageDiscountFilter]
type usageDiscountFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageDiscountFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDiscountFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type UsageDiscountFiltersField string

const (
	UsageDiscountFiltersFieldPriceID       UsageDiscountFiltersField = "price_id"
	UsageDiscountFiltersFieldItemID        UsageDiscountFiltersField = "item_id"
	UsageDiscountFiltersFieldPriceType     UsageDiscountFiltersField = "price_type"
	UsageDiscountFiltersFieldCurrency      UsageDiscountFiltersField = "currency"
	UsageDiscountFiltersFieldPricingUnitID UsageDiscountFiltersField = "pricing_unit_id"
)

func (r UsageDiscountFiltersField) IsKnown() bool {
	switch r {
	case UsageDiscountFiltersFieldPriceID, UsageDiscountFiltersFieldItemID, UsageDiscountFiltersFieldPriceType, UsageDiscountFiltersFieldCurrency, UsageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type UsageDiscountFiltersOperator string

const (
	UsageDiscountFiltersOperatorIncludes UsageDiscountFiltersOperator = "includes"
	UsageDiscountFiltersOperatorExcludes UsageDiscountFiltersOperator = "excludes"
)

func (r UsageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case UsageDiscountFiltersOperatorIncludes, UsageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

type UsageDiscountParam struct {
	DiscountType param.Field[UsageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]UsageDiscountFilterParam] `json:"filters"`
	Reason  param.Field[string]                     `json:"reason"`
}

func (r UsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDiscountParam) ImplementsDiscountUnionParam() {}

type UsageDiscountFilterParam struct {
	// The property of the price to filter on.
	Field param.Field[UsageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[UsageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r UsageDiscountFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
