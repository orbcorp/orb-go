// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// BetaService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaService] method instead.
type BetaService struct {
	Options        []option.RequestOption
	ExternalPlanID *BetaExternalPlanIDService
}

// NewBetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaService(opts ...option.RequestOption) (r *BetaService) {
	r = &BetaService{}
	r.Options = opts
	r.ExternalPlanID = NewBetaExternalPlanIDService(opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows the creation of a new plan version for an existing plan.
func (r *BetaService) NewPlanVersion(ctx context.Context, planID string, body BetaNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/versions", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *BetaService) FetchPlanVersion(ctx context.Context, planID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/versions/%s", planID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows setting the default version of a plan.
func (r *BetaService) SetDefaultPlanVersion(ctx context.Context, planID string, body BetaSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/set_default_version", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanVersion struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time               `json:"created_at,required" format:"date-time"`
	PlanPhases  []PlanVersionPhase      `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []Price         `json:"prices,required"`
	Version int64           `json:"version,required"`
	JSON    planVersionJSON `json:"-"`
}

// planVersionJSON contains the JSON metadata for the struct [PlanVersion]
type planVersionJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionJSON) RawJSON() string {
	return r.raw
}

type PlanVersionAdjustment struct {
	ID             string                               `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFilter],
	// [[]PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                   `json:"usage_discount"`
	JSON          planVersionAdjustmentJSON `json:"-"`
	union         PlanVersionAdjustmentsUnion
}

// planVersionAdjustmentJSON contains the JSON metadata for the struct
// [PlanVersionAdjustment]
type planVersionAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	AmountDiscount     apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	MinimumAmount      apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r planVersionAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanVersionAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanVersionAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanVersionAdjustmentsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhaseMinimumAdjustment],
// [PlanVersionAdjustmentsPlanPhaseMaximumAdjustment].
func (r PlanVersionAdjustment) AsUnion() PlanVersionAdjustmentsUnion {
	return r.union
}

// Union satisfied by [PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhaseMinimumAdjustment] or
// [PlanVersionAdjustmentsPlanPhaseMaximumAdjustment].
type PlanVersionAdjustmentsUnion interface {
	implementsPlanVersionAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanVersionAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionAdjustmentsPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionAdjustmentsPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                               `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                    `json:"usage_discount,required"`
	JSON          planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON contains the JSON
// metadata for the struct [PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment]
type planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustment) implementsPlanVersionAdjustment() {}

type PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                         `json:"values,required"`
	JSON   planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON contains the
// JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter]
type planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                      `json:"reason,required,nullable"`
	JSON   planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON contains the JSON
// metadata for the struct
// [PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment]
type planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustment) implementsPlanVersionAdjustment() {}

type PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                          `json:"values,required"`
	JSON   planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON contains the
// JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter]
type planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, PlanVersionAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                    `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                          `json:"reason,required,nullable"`
	JSON   planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON contains the
// JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment]
type planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustment) implementsPlanVersionAdjustment() {
}

type PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON contains
// the JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter]
type planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseMinimumAdjustment struct {
	ID             string                                                         `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                               `json:"reason,required,nullable"`
	JSON   planVersionAdjustmentsPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseMinimumAdjustmentJSON contains the JSON metadata
// for the struct [PlanVersionAdjustmentsPlanPhaseMinimumAdjustment]
type planVersionAdjustmentsPlanPhaseMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionAdjustmentsPlanPhaseMinimumAdjustment) implementsPlanVersionAdjustment() {}

type PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                   `json:"values,required"`
	JSON   planVersionAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON contains the JSON
// metadata for the struct [PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFilter]
type planVersionAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField string

const (
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID       PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID        PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType     PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency      PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "currency"
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID, PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID, PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType, PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency, PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, PlanVersionAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseMaximumAdjustment struct {
	ID             string                                                         `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                               `json:"reason,required,nullable"`
	JSON   planVersionAdjustmentsPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseMaximumAdjustmentJSON contains the JSON metadata
// for the struct [PlanVersionAdjustmentsPlanPhaseMaximumAdjustment]
type planVersionAdjustmentsPlanPhaseMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionAdjustmentsPlanPhaseMaximumAdjustment) implementsPlanVersionAdjustment() {}

type PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                   `json:"values,required"`
	JSON   planVersionAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON contains the JSON
// metadata for the struct [PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFilter]
type planVersionAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField string

const (
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID       PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID        PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType     PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency      PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "currency"
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID, PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID, PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType, PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency, PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, PlanVersionAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionAdjustmentsAdjustmentType string

const (
	PlanVersionAdjustmentsAdjustmentTypeUsageDiscount      PlanVersionAdjustmentsAdjustmentType = "usage_discount"
	PlanVersionAdjustmentsAdjustmentTypeAmountDiscount     PlanVersionAdjustmentsAdjustmentType = "amount_discount"
	PlanVersionAdjustmentsAdjustmentTypePercentageDiscount PlanVersionAdjustmentsAdjustmentType = "percentage_discount"
	PlanVersionAdjustmentsAdjustmentTypeMinimum            PlanVersionAdjustmentsAdjustmentType = "minimum"
	PlanVersionAdjustmentsAdjustmentTypeMaximum            PlanVersionAdjustmentsAdjustmentType = "maximum"
)

func (r PlanVersionAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsAdjustmentTypeUsageDiscount, PlanVersionAdjustmentsAdjustmentTypeAmountDiscount, PlanVersionAdjustmentsAdjustmentTypePercentageDiscount, PlanVersionAdjustmentsAdjustmentTypeMinimum, PlanVersionAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionPhase struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                        `json:"duration,required,nullable"`
	DurationUnit PlanVersionPhaseDurationUnit `json:"duration_unit,required,nullable"`
	Name         string                       `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                `json:"order,required"`
	JSON  planVersionPhaseJSON `json:"-"`
}

// planVersionPhaseJSON contains the JSON metadata for the struct
// [PlanVersionPhase]
type planVersionPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanVersionPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionPhaseDurationUnit string

const (
	PlanVersionPhaseDurationUnitDaily      PlanVersionPhaseDurationUnit = "daily"
	PlanVersionPhaseDurationUnitMonthly    PlanVersionPhaseDurationUnit = "monthly"
	PlanVersionPhaseDurationUnitQuarterly  PlanVersionPhaseDurationUnit = "quarterly"
	PlanVersionPhaseDurationUnitSemiAnnual PlanVersionPhaseDurationUnit = "semi_annual"
	PlanVersionPhaseDurationUnitAnnual     PlanVersionPhaseDurationUnit = "annual"
)

func (r PlanVersionPhaseDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionPhaseDurationUnitDaily, PlanVersionPhaseDurationUnitMonthly, PlanVersionPhaseDurationUnitQuarterly, PlanVersionPhaseDurationUnitSemiAnnual, PlanVersionPhaseDurationUnitAnnual:
		return true
	}
	return false
}

type BetaNewPlanVersionParams struct {
	// New version number.
	Version param.Field[int64] `json:"version,required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]BetaNewPlanVersionParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]BetaNewPlanVersionParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]BetaNewPlanVersionParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]BetaNewPlanVersionParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]BetaNewPlanVersionParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]BetaNewPlanVersionParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r BetaNewPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                         `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                  `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                  `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string]      `json:"currency"`
	Filters  param.Field[interface{}] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType     param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                   `json:"usage_discount"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustment) implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount],
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount],
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount],
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum],
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum],
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustment].
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion interface {
	implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion()
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                             `json:"percentage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount) implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID       BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID        BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "item_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_type"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "currency"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "includes"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                        `json:"usage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount) implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID       BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID        BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "item_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_type"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "currency"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "includes"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                          `json:"amount_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount) implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID       BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID        BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "item_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_type"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "currency"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "includes"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum) implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID       BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "price_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldItemID        BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "item_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "price_type"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "currency"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldItemID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator = "includes"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                   `json:"maximum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum) implementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID       BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "price_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldItemID        BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "item_id"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "price_type"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "currency"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldItemID, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator = "includes"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[BetaNewPlanVersionParamsAddPricesAllocationPrice] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[BetaNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r BetaNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type BetaNewPlanVersionParamsAddPricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// The custom expiration for the allocation.
	CustomExpiration param.Field[BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration] `json:"custom_expiration"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period. Set to null if using custom_expiration.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence"`
}

func (r BetaNewPlanVersionParamsAddPricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type BetaNewPlanVersionParamsAddPricesAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The custom expiration for the allocation.
type BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration struct {
	Duration     param.Field[int64]                                                                        `json:"duration,required"`
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitDay   BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitMonth BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitDay, BetaNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitMonth:
		return true
	}
	return false
}

// The price to add to the plan
type BetaNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance           param.Field[bool]        `json:"billed_in_advance"`
	BillingCycleConfiguration param.Field[interface{}] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[interface{}] `json:"bps_config"`
	BulkBpsConfig             param.Field[interface{}] `json:"bulk_bps_config"`
	BulkConfig                param.Field[interface{}] `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}] `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                      param.Field[string]      `json:"currency"`
	DimensionalPriceConfiguration param.Field[interface{}] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig              param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey                    param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration           param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[interface{}] `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[interface{}] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}] `json:"metadata"`
	PackageConfig                         param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}] `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}] `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by [BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice],
// [BetaNewPlanVersionParamsAddPricesPrice].
type BetaNewPlanVersionParamsAddPricesPriceUnion interface {
	implementsBetaNewPlanVersionParamsAddPricesPriceUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                           `json:"name,required"`
	UnitConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                 `json:"name,required"`
	PackageConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelTypePackage BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType = "package"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                               `json:"item_id,required"`
	MatrixConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                               `json:"name,required"`
	TieredConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                     `json:"name,required"`
	TieredBpsConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                         `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelTypeBps BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                 `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                 `json:"name,required"`
	TieredPackageConfig param.Field[map[string]interface{}] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                 `json:"name,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                 `json:"name,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                 `json:"name,required"`
	TieredWithProrationConfig param.Field[map[string]interface{}] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	UnitWithProrationConfig param.Field[map[string]interface{}] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                     `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                              `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                             `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                           `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                           `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                        `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                           `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                           `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                         `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                           param.Field[string]                 `json:"name,required"`
	TieredPackageWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_package_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                           `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                 `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice) implementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnit                            BetaNewPlanVersionParamsAddPricesPriceModelType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePackage                         BetaNewPlanVersionParamsAddPricesPriceModelType = "package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTiered                          BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredBps                       BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_bps"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBps                             BetaNewPlanVersionParamsAddPricesPriceModelType = "bps"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkBps                         BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk_bps"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulk                            BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceModelTypeUnit, BetaNewPlanVersionParamsAddPricesPriceModelTypePackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrix, BetaNewPlanVersionParamsAddPricesPriceModelTypeTiered, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredBps, BetaNewPlanVersionParamsAddPricesPriceModelTypeBps, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkBps, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulk, BetaNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id,required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                             `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                      `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                      `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string]      `json:"currency"`
	Filters  param.Field[interface{}] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType     param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                       `json:"usage_discount"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment) implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount],
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount],
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount],
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum],
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum],
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment].
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion interface {
	implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion()
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                                 `json:"percentage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID       BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID        BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "item_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_type"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "currency"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "includes"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                            `json:"usage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID       BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID        BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "item_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_type"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "currency"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "includes"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                              `json:"amount_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID       BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID        BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "item_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_type"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "currency"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "includes"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum) implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID       BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "price_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldItemID        BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "item_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "price_type"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "currency"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldItemID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator = "includes"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                       `json:"maximum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType] `json:"price_type"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum) implementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID       BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "price_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldItemID        BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "item_id"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "price_type"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "currency"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "pricing_unit_id"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldItemID, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator = "includes"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator = "excludes"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[BetaNewPlanVersionParamsReplacePricesAllocationPrice] `json:"allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[BetaNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r BetaNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type BetaNewPlanVersionParamsReplacePricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// The custom expiration for the allocation.
	CustomExpiration param.Field[BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration] `json:"custom_expiration"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period. Set to null if using custom_expiration.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence"`
}

func (r BetaNewPlanVersionParamsReplacePricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The custom expiration for the allocation.
type BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration struct {
	Duration     param.Field[int64]                                                                            `json:"duration,required"`
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitMonth:
		return true
	}
	return false
}

// The price to add to the plan
type BetaNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance           param.Field[bool]        `json:"billed_in_advance"`
	BillingCycleConfiguration param.Field[interface{}] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[interface{}] `json:"bps_config"`
	BulkBpsConfig             param.Field[interface{}] `json:"bulk_bps_config"`
	BulkConfig                param.Field[interface{}] `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}] `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                      param.Field[string]      `json:"currency"`
	DimensionalPriceConfiguration param.Field[interface{}] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig              param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey                    param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration           param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[interface{}] `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[interface{}] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}] `json:"metadata"`
	PackageConfig                         param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}] `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}] `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by [BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice],
// [BetaNewPlanVersionParamsReplacePricesPrice].
type BetaNewPlanVersionParamsReplacePricesPriceUnion interface {
	implementsBetaNewPlanVersionParamsReplacePricesPriceUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                               `json:"name,required"`
	UnitConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                     `json:"name,required"`
	PackageConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType = "package"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                   `json:"item_id,required"`
	MatrixConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                   `json:"name,required"`
	TieredConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                         `json:"name,required"`
	TieredBpsConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                 `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                 `json:"name,required"`
	TieredPackageConfig param.Field[map[string]interface{}] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                 `json:"name,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                 `json:"name,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                 `json:"name,required"`
	TieredWithProrationConfig param.Field[map[string]interface{}] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	UnitWithProrationConfig param.Field[map[string]interface{}] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                         `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                  `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                 `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                               `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                               `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                            `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                               `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                               `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                             `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                  `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                           param.Field[string]                 `json:"name,required"`
	TieredPackageWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_package_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                               `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                     `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice) implementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePackage                         BetaNewPlanVersionParamsReplacePricesPriceModelType = "package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredBps                       BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_bps"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBps                             BetaNewPlanVersionParamsReplacePricesPriceModelType = "bps"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkBps                         BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk_bps"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceModelTypePackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTiered, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredBps, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBps, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkBps, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulk, BetaNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type BetaSetDefaultPlanVersionParams struct {
	// Plan version to set as the default.
	Version param.Field[int64] `json:"version,required"`
}

func (r BetaSetDefaultPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
