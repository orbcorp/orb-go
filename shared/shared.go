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
	AppliesToPriceIDs []string           `json:"applies_to_price_ids,nullable"`
	Reason            string             `json:"reason,nullable"`
	JSON              amountDiscountJSON `json:"-"`
}

// amountDiscountJSON contains the JSON metadata for the struct [AmountDiscount]
type amountDiscountJSON struct {
	AmountDiscount    apijson.Field
	DiscountType      apijson.Field
	AppliesToPriceIDs apijson.Field
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

type AmountDiscountParam struct {
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string]                     `json:"amount_discount,required"`
	DiscountType   param.Field[AmountDiscountDiscountType] `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	Reason            param.Field[string]   `json:"reason"`
}

func (r AmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AmountDiscountParam) ImplementsDiscountUnionParam() {}

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
	AppliesToPriceIDs []string               `json:"applies_to_price_ids,nullable"`
	Reason            string                 `json:"reason,nullable"`
	JSON              percentageDiscountJSON `json:"-"`
}

// percentageDiscountJSON contains the JSON metadata for the struct
// [PercentageDiscount]
type percentageDiscountJSON struct {
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPriceIDs  apijson.Field
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

type PercentageDiscountParam struct {
	DiscountType param.Field[PercentageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	Reason            param.Field[string]   `json:"reason"`
}

func (r PercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PercentageDiscountParam) ImplementsDiscountUnionParam() {}

type TrialDiscount struct {
	DiscountType TrialDiscountDiscountType `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	Reason            string   `json:"reason,nullable"`
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

type TrialDiscountParam struct {
	DiscountType param.Field[TrialDiscountDiscountType] `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	Reason            param.Field[string]   `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
}

func (r TrialDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrialDiscountParam) ImplementsDiscountUnionParam() {}

type UsageDiscount struct {
	DiscountType UsageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64 `json:"usage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string          `json:"applies_to_price_ids,nullable"`
	Reason            string            `json:"reason,nullable"`
	JSON              usageDiscountJSON `json:"-"`
}

// usageDiscountJSON contains the JSON metadata for the struct [UsageDiscount]
type usageDiscountJSON struct {
	DiscountType      apijson.Field
	UsageDiscount     apijson.Field
	AppliesToPriceIDs apijson.Field
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

type UsageDiscountParam struct {
	DiscountType param.Field[UsageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	Reason            param.Field[string]   `json:"reason"`
}

func (r UsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDiscountParam) ImplementsDiscountUnionParam() {}
