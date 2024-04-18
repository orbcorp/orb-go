// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/tidwall/gjson"
)

type Discount struct {
	DiscountType      DiscountDiscountType `json:"discount_type,required"`
	AppliesToPriceIDs interface{}          `json:"applies_to_price_ids"`
	Reason            string               `json:"reason,nullable"`
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Discount) AsUnion() DiscountUnion {
	return r.union
}

// Union satisfied by [shared.DiscountPercentageDiscount],
// [shared.DiscountTrialDiscount], [shared.DiscountUsageDiscount] or
// [shared.DiscountAmountDiscount].
type DiscountUnion interface {
	implementsSharedDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DiscountUnion)(nil)).Elem(),
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

func (r discountPercentageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r DiscountPercentageDiscount) implementsSharedDiscount() {}

type DiscountPercentageDiscountDiscountType string

const (
	DiscountPercentageDiscountDiscountTypePercentage DiscountPercentageDiscountDiscountType = "percentage"
)

func (r DiscountPercentageDiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountPercentageDiscountDiscountTypePercentage:
		return true
	}
	return false
}

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

func (r discountTrialDiscountJSON) RawJSON() string {
	return r.raw
}

func (r DiscountTrialDiscount) implementsSharedDiscount() {}

type DiscountTrialDiscountDiscountType string

const (
	DiscountTrialDiscountDiscountTypeTrial DiscountTrialDiscountDiscountType = "trial"
)

func (r DiscountTrialDiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountTrialDiscountDiscountTypeTrial:
		return true
	}
	return false
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

func (r DiscountUsageDiscount) implementsSharedDiscount() {}

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

func (r discountAmountDiscountJSON) RawJSON() string {
	return r.raw
}

func (r DiscountAmountDiscount) implementsSharedDiscount() {}

type DiscountAmountDiscountDiscountType string

const (
	DiscountAmountDiscountDiscountTypeAmount DiscountAmountDiscountDiscountType = "amount"
)

func (r DiscountAmountDiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountAmountDiscountDiscountTypeAmount:
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
