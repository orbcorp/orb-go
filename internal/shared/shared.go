// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"reflect"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/tidwall/gjson"
)

// Union satisfied by [shared.DiscountPercentageDiscount],
// [shared.DiscountTrialDiscount], [shared.DiscountUsageDiscount] or
// [shared.DiscountAmountDiscount].
type Discount interface {
	implementsSharedDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Discount)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DiscountPercentageDiscount{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DiscountTrialDiscount{}),
			DiscriminatorValue: "trial",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DiscountUsageDiscount{}),
			DiscriminatorValue: "usage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DiscountAmountDiscount{}),
			DiscriminatorValue: "amount",
		},
	)
}

type DiscountPercentageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                               `json:"applies_to_price_ids,required"`
	DiscountType      DiscountPercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64                        `json:"percentage_discount,required"`
	Reason             string                         `json:"reason,nullable"`
	JSON               discountPercentageDiscountJSON `json:"-"`
}

// discountPercentageDiscountJSON contains the JSON metadata for the struct
// [DiscountPercentageDiscount]
type discountPercentageDiscountJSON struct {
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DiscountPercentageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DiscountPercentageDiscount) implementsSharedDiscount() {}

type DiscountPercentageDiscountDiscountType string

const (
	DiscountPercentageDiscountDiscountTypePercentage DiscountPercentageDiscountDiscountType = "percentage"
)

type DiscountTrialDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                          `json:"applies_to_price_ids,required"`
	DiscountType      DiscountTrialDiscountDiscountType `json:"discount_type,required"`
	Reason            string                            `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64                   `json:"trial_percentage_discount,nullable"`
	JSON                    discountTrialDiscountJSON `json:"-"`
}

// discountTrialDiscountJSON contains the JSON metadata for the struct
// [DiscountTrialDiscount]
type discountTrialDiscountJSON struct {
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DiscountTrialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DiscountTrialDiscount) implementsSharedDiscount() {}

type DiscountTrialDiscountDiscountType string

const (
	DiscountTrialDiscountDiscountTypeTrial DiscountTrialDiscountDiscountType = "trial"
)

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

func (r DiscountUsageDiscount) implementsSharedDiscount() {}

type DiscountUsageDiscountDiscountType string

const (
	DiscountUsageDiscountDiscountTypeUsage DiscountUsageDiscountDiscountType = "usage"
)

type DiscountAmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                           `json:"applies_to_price_ids,required"`
	DiscountType      DiscountAmountDiscountDiscountType `json:"discount_type,required"`
	Reason            string                             `json:"reason,nullable"`
	JSON              discountAmountDiscountJSON         `json:"-"`
}

// discountAmountDiscountJSON contains the JSON metadata for the struct
// [DiscountAmountDiscount]
type discountAmountDiscountJSON struct {
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DiscountAmountDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DiscountAmountDiscount) implementsSharedDiscount() {}

type DiscountAmountDiscountDiscountType string

const (
	DiscountAmountDiscountDiscountTypeAmount DiscountAmountDiscountDiscountType = "amount"
)
