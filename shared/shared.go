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
	AmountDiscount string `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                   `json:"applies_to_price_ids,required"`
	DiscountType      AmountDiscountDiscountType `json:"discount_type,required"`
	Reason            string                     `json:"reason,nullable"`
	JSON              amountDiscountJSON         `json:"-"`
}

// amountDiscountJSON contains the JSON metadata for the struct [AmountDiscount]
type amountDiscountJSON struct {
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
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

func (r AmountDiscount) ImplementsSharedDiscount() {}

func (r AmountDiscount) ImplementsSharedInvoiceLevelDiscount() {}

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
	AmountDiscount param.Field[string] `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                   `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[AmountDiscountDiscountType] `json:"discount_type,required"`
	Reason            param.Field[string]                     `json:"reason"`
}

func (r AmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AmountDiscountParam) ImplementsSharedDiscountUnionParam() {}

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
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	Reason            string      `json:"reason,nullable"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64 `json:"trial_percentage_discount,nullable"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64 `json:"usage_discount"`
	// Only available if discount_type is `amount`.
	AmountDiscount string       `json:"amount_discount"`
	JSON           discountJSON `json:"-"`
	union          DiscountUnion
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	DiscountType            apijson.Field
	AppliesToPriceIDs       apijson.Field
	Reason                  apijson.Field
	PercentageDiscount      apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	UsageDiscount           apijson.Field
	AmountDiscount          apijson.Field
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
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.TrialDiscount], [shared.DiscountUsageDiscount], [shared.AmountDiscount].
func (r Discount) AsUnion() DiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount], [shared.TrialDiscount],
// [shared.DiscountUsageDiscount] or [shared.AmountDiscount].
type DiscountUnion interface {
	ImplementsSharedDiscount()
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
			Type:               reflect.TypeOf(DiscountUsageDiscount{}),
			DiscriminatorValue: "usage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscount{}),
			DiscriminatorValue: "amount",
		},
	)
}

type DiscountUsageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                          `json:"applies_to_price_ids,required"`
	DiscountType      DiscountUsageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                   `json:"usage_discount,required"`
	Reason        string                    `json:"reason,nullable"`
	JSON          discountUsageDiscountJSON `json:"-"`
}

// discountUsageDiscountJSON contains the JSON metadata for the struct
// [DiscountUsageDiscount]
type discountUsageDiscountJSON struct {
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
	UsageDiscount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DiscountUsageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountUsageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r DiscountUsageDiscount) ImplementsSharedDiscount() {}

type DiscountUsageDiscountDiscountType string

const (
	DiscountUsageDiscountDiscountTypeUsage DiscountUsageDiscountDiscountType = "usage"
)

func (r DiscountUsageDiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountUsageDiscountDiscountTypeUsage:
		return true
	}
	return false
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
	DiscountType      param.Field[DiscountDiscountType] `json:"discount_type,required"`
	AppliesToPriceIDs param.Field[interface{}]          `json:"applies_to_price_ids"`
	Reason            param.Field[string]               `json:"reason"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
}

func (r DiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscountParam) ImplementsSharedDiscountUnionParam() {}

// Satisfied by [shared.PercentageDiscountParam], [shared.TrialDiscountParam],
// [shared.DiscountUsageDiscountParam], [shared.AmountDiscountParam],
// [DiscountParam].
type DiscountUnionParam interface {
	ImplementsSharedDiscountUnionParam()
}

type DiscountUsageDiscountParam struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                          `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[DiscountUsageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount,required"`
	Reason        param.Field[string]  `json:"reason"`
}

func (r DiscountUsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscountUsageDiscountParam) ImplementsSharedDiscountUnionParam() {}

type InvoiceLevelDiscount struct {
	DiscountType InvoiceLevelDiscountDiscountType `json:"discount_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	Reason            string      `json:"reason,nullable"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
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
	AppliesToPriceIDs       apijson.Field
	Reason                  apijson.Field
	PercentageDiscount      apijson.Field
	AmountDiscount          apijson.Field
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
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.AmountDiscount], [shared.TrialDiscount].
func (r InvoiceLevelDiscount) AsUnion() InvoiceLevelDiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount], [shared.AmountDiscount] or
// [shared.TrialDiscount].
type InvoiceLevelDiscountUnion interface {
	ImplementsSharedInvoiceLevelDiscount()
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
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                       `json:"applies_to_price_ids,required"`
	DiscountType      PercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64                `json:"percentage_discount,required"`
	Reason             string                 `json:"reason,nullable"`
	JSON               percentageDiscountJSON `json:"-"`
}

// percentageDiscountJSON contains the JSON metadata for the struct
// [PercentageDiscount]
type percentageDiscountJSON struct {
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
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

func (r PercentageDiscount) ImplementsSharedDiscount() {}

func (r PercentageDiscount) ImplementsSharedInvoiceLevelDiscount() {}

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
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                       `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[PercentageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
	Reason             param.Field[string]  `json:"reason"`
}

func (r PercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PercentageDiscountParam) ImplementsSharedDiscountUnionParam() {}

type TrialDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                  `json:"applies_to_price_ids,required"`
	DiscountType      TrialDiscountDiscountType `json:"discount_type,required"`
	Reason            string                    `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64           `json:"trial_percentage_discount,nullable"`
	JSON                    trialDiscountJSON `json:"-"`
}

// trialDiscountJSON contains the JSON metadata for the struct [TrialDiscount]
type trialDiscountJSON struct {
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
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

func (r TrialDiscount) ImplementsSharedDiscount() {}

func (r TrialDiscount) ImplementsSharedInvoiceLevelDiscount() {}

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
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                  `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[TrialDiscountDiscountType] `json:"discount_type,required"`
	Reason            param.Field[string]                    `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
}

func (r TrialDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrialDiscountParam) ImplementsSharedDiscountUnionParam() {}
