// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// CouponSubscriptionService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCouponSubscriptionService] method instead.
type CouponSubscriptionService struct {
	Options []option.RequestOption
}

// NewCouponSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCouponSubscriptionService(opts ...option.RequestOption) (r *CouponSubscriptionService) {
	r = &CouponSubscriptionService{}
	r.Options = opts
	return
}

// This endpoint returns a list of all subscriptions that have redeemed a given
// coupon as a [paginated](../reference/pagination) list, ordered starting from the
// most recently created subscription. For a full discussion of the subscription
// resource, see [Subscription](../guides/concepts#subscription).
func (r *CouponSubscriptionService) List(ctx context.Context, couponID string, query CouponSubscriptionListParams, opts ...option.RequestOption) (res *pagination.Page[Subscription], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if couponID == "" {
		err = errors.New("missing required coupon_id parameter")
		return
	}
	path := fmt.Sprintf("coupons/%s/subscriptions", couponID)
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

// This endpoint returns a list of all subscriptions that have redeemed a given
// coupon as a [paginated](../reference/pagination) list, ordered starting from the
// most recently created subscription. For a full discussion of the subscription
// resource, see [Subscription](../guides/concepts#subscription).
func (r *CouponSubscriptionService) ListAutoPaging(ctx context.Context, couponID string, query CouponSubscriptionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Subscription] {
	return pagination.NewPageAutoPager(r.List(ctx, couponID, query, opts...))
}

type CouponSubscriptionListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CouponSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r CouponSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
