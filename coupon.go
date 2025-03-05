// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
)

// CouponService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCouponService] method instead.
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
func (r *CouponService) New(ctx context.Context, body CouponNewParams, opts ...option.RequestOption) (res *shared.CouponModel, err error) {
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
func (r *CouponService) List(ctx context.Context, query CouponListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CouponModel], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *CouponService) ListAutoPaging(ctx context.Context, query CouponListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CouponModel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint allows a coupon to be archived. Archived coupons can no longer be
// redeemed, and will be hidden from lists of active coupons. Additionally, once a
// coupon is archived, its redemption code can be reused for a different coupon.
func (r *CouponService) Archive(ctx context.Context, couponID string, opts ...option.RequestOption) (res *shared.CouponModel, err error) {
	opts = append(r.Options[:], opts...)
	if couponID == "" {
		err = errors.New("missing required coupon_id parameter")
		return
	}
	path := fmt.Sprintf("coupons/%s/archive", couponID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint retrieves a coupon by its ID. To fetch coupons by their redemption
// code, use the [List coupons](list-coupons) endpoint with the redemption_code
// parameter.
func (r *CouponService) Fetch(ctx context.Context, couponID string, opts ...option.RequestOption) (res *shared.CouponModel, err error) {
	opts = append(r.Options[:], opts...)
	if couponID == "" {
		err = errors.New("missing required coupon_id parameter")
		return
	}
	path := fmt.Sprintf("coupons/%s", couponID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CouponNewParams struct {
	Discount param.Field[CouponNewParamsDiscountUnion] `json:"discount,required"`
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

type CouponNewParamsDiscount struct {
	DiscountType       param.Field[CouponNewParamsDiscountDiscountType] `json:"discount_type,required"`
	AmountDiscount     param.Field[string]                              `json:"amount_discount"`
	PercentageDiscount param.Field[float64]                             `json:"percentage_discount"`
}

func (r CouponNewParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CouponNewParamsDiscount) implementsCouponNewParamsDiscountUnion() {}

// Satisfied by [CouponNewParamsDiscountNewCouponPercentageDiscount],
// [CouponNewParamsDiscountNewCouponAmountDiscount], [CouponNewParamsDiscount].
type CouponNewParamsDiscountUnion interface {
	implementsCouponNewParamsDiscountUnion()
}

type CouponNewParamsDiscountNewCouponPercentageDiscount struct {
	DiscountType       param.Field[CouponNewParamsDiscountNewCouponPercentageDiscountDiscountType] `json:"discount_type,required"`
	PercentageDiscount param.Field[float64]                                                        `json:"percentage_discount,required"`
}

func (r CouponNewParamsDiscountNewCouponPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CouponNewParamsDiscountNewCouponPercentageDiscount) implementsCouponNewParamsDiscountUnion() {
}

type CouponNewParamsDiscountNewCouponPercentageDiscountDiscountType string

const (
	CouponNewParamsDiscountNewCouponPercentageDiscountDiscountTypePercentage CouponNewParamsDiscountNewCouponPercentageDiscountDiscountType = "percentage"
)

func (r CouponNewParamsDiscountNewCouponPercentageDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponNewParamsDiscountNewCouponPercentageDiscountDiscountTypePercentage:
		return true
	}
	return false
}

type CouponNewParamsDiscountNewCouponAmountDiscount struct {
	AmountDiscount param.Field[string]                                                     `json:"amount_discount,required"`
	DiscountType   param.Field[CouponNewParamsDiscountNewCouponAmountDiscountDiscountType] `json:"discount_type,required"`
}

func (r CouponNewParamsDiscountNewCouponAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CouponNewParamsDiscountNewCouponAmountDiscount) implementsCouponNewParamsDiscountUnion() {}

type CouponNewParamsDiscountNewCouponAmountDiscountDiscountType string

const (
	CouponNewParamsDiscountNewCouponAmountDiscountDiscountTypeAmount CouponNewParamsDiscountNewCouponAmountDiscountDiscountType = "amount"
)

func (r CouponNewParamsDiscountNewCouponAmountDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponNewParamsDiscountNewCouponAmountDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type CouponNewParamsDiscountDiscountType string

const (
	CouponNewParamsDiscountDiscountTypePercentage CouponNewParamsDiscountDiscountType = "percentage"
	CouponNewParamsDiscountDiscountTypeAmount     CouponNewParamsDiscountDiscountType = "amount"
)

func (r CouponNewParamsDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponNewParamsDiscountDiscountTypePercentage, CouponNewParamsDiscountDiscountTypeAmount:
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
