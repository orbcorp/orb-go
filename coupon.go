// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// CouponService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCouponService] method instead.
type CouponService struct {
	Options       []option.RequestOption
	Subscriptions *CouponSubscriptionService
}

// NewCouponService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCouponService(opts ...option.RequestOption) (r *CouponService) {
	r = &CouponService{}
	r.Options = opts
	r.Subscriptions = NewCouponSubscriptionService(opts...)
	return
}

// This endpoint allows the creation of coupons, which can then be redeemed at
// subscription creation or plan change.
func (r *CouponService) New(ctx context.Context, body CouponNewParams, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = append(r.Options[:], opts...)
	path := "coupons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all coupons for an account in a list format.
//
// The list of coupons is ordered starting from the most recently created coupon.
// The response also includes `pagination_metadata`, which lets the caller retrieve
// the next page of results if they exist. More information about pagination can be
// found in the Pagination-metadata schema.
func (r *CouponService) List(ctx context.Context, query CouponListParams, opts ...option.RequestOption) (res *shared.Page[Coupon], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "coupons"
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

// This endpoint returns a list of all coupons for an account in a list format.
//
// The list of coupons is ordered starting from the most recently created coupon.
// The response also includes `pagination_metadata`, which lets the caller retrieve
// the next page of results if they exist. More information about pagination can be
// found in the Pagination-metadata schema.
func (r *CouponService) ListAutoPaging(ctx context.Context, query CouponListParams, opts ...option.RequestOption) *shared.PageAutoPager[Coupon] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint allows a coupon to be archived. Archived coupons can no longer be
// redeemed, and will be hidden from lists of active coupons. Additionally, once a
// coupon is archived, its redemption code can be reused for a different coupon.
func (r *CouponService) Archive(ctx context.Context, couponID string, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("coupons/%s/archive", couponID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint retrieves a coupon by its ID. To fetch coupons by their redemption
// code, use the [List coupons](list-coupons) endpoint with the redemption_code
// parameter.
func (r *CouponService) Fetch(ctx context.Context, couponID string, opts ...option.RequestOption) (res *Coupon, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("coupons/%s", couponID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A coupon represents a reusable discount configuration, and have an attached
// redemption code that can be issued to your end users. Coupons are most often
// used in self-serve signup or upgrade flows in your checkout experience or
// billing portal.
//
// To redeem a coupon, pass the `redemption_code` property in the
// [create subscription](create-subscription.api.mdx) or
// [schedule plan change](schedule-plan-change.api.mdx) request.
type Coupon struct {
	// Also referred to as coupon_id in this documentation.
	ID string `json:"id,required"`
	// An archived coupon can no longer be redeemed. Active coupons will have a value
	// of null for `archived_at`; this field will be non-null for archived coupons.
	ArchivedAt time.Time      `json:"archived_at,required,nullable" format:"date-time"`
	Discount   CouponDiscount `json:"discount,required"`
	// This allows for a coupon's discount to apply for a limited time (determined in
	// months); a `null` value here means "unlimited time".
	DurationInMonths int64 `json:"duration_in_months,required,nullable"`
	// The maximum number of redemptions allowed for this coupon before it is
	// exhausted; `null` here means "unlimited".
	MaxRedemptions int64 `json:"max_redemptions,required,nullable"`
	// This string can be used to redeem this coupon for a given subscription.
	RedemptionCode string `json:"redemption_code,required"`
	// The number of times this coupon has been redeemed.
	TimesRedeemed int64      `json:"times_redeemed,required"`
	JSON          couponJSON `json:"-"`
}

// couponJSON contains the JSON metadata for the struct [Coupon]
type couponJSON struct {
	ID               apijson.Field
	ArchivedAt       apijson.Field
	Discount         apijson.Field
	DurationInMonths apijson.Field
	MaxRedemptions   apijson.Field
	RedemptionCode   apijson.Field
	TimesRedeemed    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Coupon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r couponJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [CouponDiscountPercentageDiscount] or
// [CouponDiscountAmountDiscount].
type CouponDiscount interface {
	implementsCouponDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CouponDiscount)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CouponDiscountPercentageDiscount{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CouponDiscountAmountDiscount{}),
			DiscriminatorValue: "amount",
		},
	)
}

type CouponDiscountPercentageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                                     `json:"applies_to_price_ids,required"`
	DiscountType      CouponDiscountPercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64                              `json:"percentage_discount,required"`
	Reason             string                               `json:"reason,nullable"`
	JSON               couponDiscountPercentageDiscountJSON `json:"-"`
}

// couponDiscountPercentageDiscountJSON contains the JSON metadata for the struct
// [CouponDiscountPercentageDiscount]
type couponDiscountPercentageDiscountJSON struct {
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CouponDiscountPercentageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r couponDiscountPercentageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r CouponDiscountPercentageDiscount) implementsCouponDiscount() {}

type CouponDiscountPercentageDiscountDiscountType string

const (
	CouponDiscountPercentageDiscountDiscountTypePercentage CouponDiscountPercentageDiscountDiscountType = "percentage"
)

func (r CouponDiscountPercentageDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponDiscountPercentageDiscountDiscountTypePercentage:
		return true
	}
	return false
}

type CouponDiscountAmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                                 `json:"applies_to_price_ids,required"`
	DiscountType      CouponDiscountAmountDiscountDiscountType `json:"discount_type,required"`
	Reason            string                                   `json:"reason,nullable"`
	JSON              couponDiscountAmountDiscountJSON         `json:"-"`
}

// couponDiscountAmountDiscountJSON contains the JSON metadata for the struct
// [CouponDiscountAmountDiscount]
type couponDiscountAmountDiscountJSON struct {
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CouponDiscountAmountDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r couponDiscountAmountDiscountJSON) RawJSON() string {
	return r.raw
}

func (r CouponDiscountAmountDiscount) implementsCouponDiscount() {}

type CouponDiscountAmountDiscountDiscountType string

const (
	CouponDiscountAmountDiscountDiscountTypeAmount CouponDiscountAmountDiscountDiscountType = "amount"
)

func (r CouponDiscountAmountDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponDiscountAmountDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type CouponNewParams struct {
	Discount param.Field[CouponNewParamsDiscount] `json:"discount,required"`
	// This string can be used to redeem this coupon for a given subscription.
	RedemptionCode param.Field[string] `json:"redemption_code,required"`
	// This allows for a coupon's discount to apply for a limited time (determined in
	// months); a `null` value here means "unlimited time".
	DurationInMonths param.Field[int64] `json:"duration_in_months"`
	// The maximum number of redemptions allowed for this coupon before it is
	// exhausted;`null` here means "unlimited".
	MaxRedemptions param.Field[int64] `json:"max_redemptions"`
}

func (r CouponNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [CouponNewParamsDiscountPercentageDiscount],
// [CouponNewParamsDiscountAmountDiscount].
type CouponNewParamsDiscount interface {
	implementsCouponNewParamsDiscount()
}

type CouponNewParamsDiscountPercentageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                                              `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[CouponNewParamsDiscountPercentageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
	Reason             param.Field[string]  `json:"reason"`
}

func (r CouponNewParamsDiscountPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CouponNewParamsDiscountPercentageDiscount) implementsCouponNewParamsDiscount() {}

type CouponNewParamsDiscountPercentageDiscountDiscountType string

const (
	CouponNewParamsDiscountPercentageDiscountDiscountTypePercentage CouponNewParamsDiscountPercentageDiscountDiscountType = "percentage"
)

func (r CouponNewParamsDiscountPercentageDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponNewParamsDiscountPercentageDiscountDiscountTypePercentage:
		return true
	}
	return false
}

type CouponNewParamsDiscountAmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                                          `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[CouponNewParamsDiscountAmountDiscountDiscountType] `json:"discount_type,required"`
	Reason            param.Field[string]                                            `json:"reason"`
}

func (r CouponNewParamsDiscountAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CouponNewParamsDiscountAmountDiscount) implementsCouponNewParamsDiscount() {}

type CouponNewParamsDiscountAmountDiscountDiscountType string

const (
	CouponNewParamsDiscountAmountDiscountDiscountTypeAmount CouponNewParamsDiscountAmountDiscountDiscountType = "amount"
)

func (r CouponNewParamsDiscountAmountDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponNewParamsDiscountAmountDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type CouponListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// Filter to coupons matching this redemption code.
	RedemptionCode param.Field[string] `query:"redemption_code"`
	// Show archived coupons as well (by default, this endpoint only returns active
	// coupons).
	ShowArchived param.Field[bool] `query:"show_archived"`
}

// URLQuery serializes [CouponListParams]'s query parameters as `url.Values`.
func (r CouponListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
