// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"fmt"
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

// BetaPriceService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewBetaPriceService] method instead.
type BetaPriceService struct {
	Options []option.RequestOption
}

// NewBetaPriceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaPriceService(opts ...option.RequestOption) (r *BetaPriceService) {
	r = &BetaPriceService{}
	r.Options = opts
	return
}

// This endpoint is used to evaluate the output of a price for a given customer and
// time range. It enables filtering and grouping the output using
// [computed properties](../guides/extensibility/advanced-metrics#computed-properties),
// supporting the following workflows:
//
// 1. Showing detailed usage and costs to the end customer.
// 2. Auditing subtotals on invoice line items.
//
// For these workflows, the expressiveness of computed properties in both the
// filters and grouping is critical. For example, if you'd like to show your
// customer their usage grouped by hour and another property, you can do so with
// the following `grouping_keys`:
// `["hour_floor_timestamp_millis(timestamp_millis)", "my_property"]`. If you'd
// like to examine a customer's usage for a specific property value, you can do so
// with the following `filter`:
// `my_property = 'foo' AND my_other_property = 'bar'`.
//
// By default, the start of the time range must be no more than 100 days ago and
// the length of the results must be no greater than 1000. Note that this is a POST
// endpoint rather than a GET endpoint because it employs a JSON body rather than
// query parameters.
func (r *BetaPriceService) Evaluate(ctx context.Context, priceID string, body BetaPriceEvaluateParams, opts ...option.RequestOption) (res *BetaPriceEvaluateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("prices/%s/evaluate", priceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvaluatePriceGroup struct {
	// The price's output for the group
	Amount string `json:"amount,required"`
	// The values for the group in the order specified by `grouping_keys`
	GroupingValues []EvaluatePriceGroupGroupingValue `json:"grouping_values,required"`
	// The price's usage quantity for the group
	Quantity float64                `json:"quantity,required"`
	JSON     evaluatePriceGroupJSON `json:"-"`
}

// evaluatePriceGroupJSON contains the JSON metadata for the struct
// [EvaluatePriceGroup]
type evaluatePriceGroupJSON struct {
	Amount         apijson.Field
	GroupingValues apijson.Field
	Quantity       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvaluatePriceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluatePriceGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type EvaluatePriceGroupGroupingValue interface {
	ImplementsEvaluatePriceGroupGroupingValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluatePriceGroupGroupingValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type BetaPriceEvaluateResponse struct {
	Data []EvaluatePriceGroup          `json:"data,required"`
	JSON betaPriceEvaluateResponseJSON `json:"-"`
}

// betaPriceEvaluateResponseJSON contains the JSON metadata for the struct
// [BetaPriceEvaluateResponse]
type betaPriceEvaluateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaPriceEvaluateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaPriceEvaluateResponseJSON) RawJSON() string {
	return r.raw
}

type BetaPriceEvaluateParams struct {
	// The exclusive upper bound for event timestamps
	TimeframeEnd param.Field[time.Time] `json:"timeframe_end,required" format:"date-time"`
	// The inclusive lower bound for event timestamps
	TimeframeStart param.Field[time.Time] `json:"timeframe_start,required" format:"date-time"`
	// The ID of the customer to which this evaluation is scoped.
	CustomerID param.Field[string] `json:"customer_id"`
	// The external customer ID of the customer to which this evaluation is scoped.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// A boolean
	// [computed property](../guides/extensibility/advanced-metrics#computed-properties)
	// used to filter the underlying billable metric
	Filter param.Field[string] `json:"filter"`
	// Properties (or
	// [computed properties](../guides/extensibility/advanced-metrics#computed-properties))
	// used to group the underlying billable metric
	GroupingKeys param.Field[[]string] `json:"grouping_keys"`
}

func (r BetaPriceEvaluateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
