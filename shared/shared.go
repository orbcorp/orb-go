// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/tidwall/gjson"
)

type AddCreditLedgerEntryRequestParam struct {
	EntryType param.Field[AddCreditLedgerEntryRequestEntryType] `json:"entry_type,required"`
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount"`
	// The ID of the block affected by an expiration_change, used to differentiate
	// between multiple blocks with the same `expiry_date`.
	BlockID param.Field[string] `json:"block_id"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// An ISO 8601 format date that denotes when this credit balance should become
	// available for use.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date-time"`
	// An ISO 8601 format date that denotes when this credit balance should expire.
	ExpiryDate      param.Field[time.Time]   `json:"expiry_date" format:"date-time"`
	InvoiceSettings param.Field[interface{}] `json:"invoice_settings"`
	Metadata        param.Field[interface{}] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in the customer's
	// currency, a customer paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
	// A future date (specified in YYYY-MM-DD format) used for expiration change,
	// denoting when credits transferred (as part of a partial block expiration) should
	// expire.
	TargetExpiryDate param.Field[time.Time] `json:"target_expiry_date" format:"date"`
	// Can only be specified when `entry_type=void`. The reason for the void.
	VoidReason param.Field[AddCreditLedgerEntryRequestVoidReason] `json:"void_reason"`
}

func (r AddCreditLedgerEntryRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestParam) implementsAddCreditLedgerEntryRequestUnionParam() {}

// Satisfied by
// [shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam],
// [AddCreditLedgerEntryRequestParam].
type AddCreditLedgerEntryRequestUnionParam interface {
	implementsAddCreditLedgerEntryRequestUnionParam()
}

type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                        `json:"amount,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// An ISO 8601 format date that denotes when this credit balance should become
	// available for use.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date-time"`
	// An ISO 8601 format date that denotes when this credit balance should expire.
	ExpiryDate param.Field[time.Time] `json:"expiry_date" format:"date-time"`
	// Passing `invoice_settings` automatically generates an invoice for the newly
	// added credits. If `invoice_settings` is passed, you must specify
	// per_unit_cost_basis, as the calculation of the invoice total is done on that
	// basis.
	InvoiceSettings param.Field[AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam] `json:"invoice_settings"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in the customer's
	// currency, a customer paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
}

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType = "increment"
)

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement:
		return true
	}
	return false
}

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection param.Field[bool] `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo param.Field[string] `json:"memo"`
	// If true, the new credit block will require that the corresponding invoice is
	// paid before it can be drawn down from.
	RequireSuccessfulPayment param.Field[bool] `json:"require_successful_payment"`
}

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                        `json:"amount,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType = "decrement"
)

func (r AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam struct {
	EntryType param.Field[AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// An ISO 8601 format date that identifies the origination credit block to expire
	ExpiryDate param.Field[time.Time] `json:"expiry_date,required" format:"date-time"`
	// A future date (specified in YYYY-MM-DD format) used for expiration change,
	// denoting when credits transferred (as part of a partial block expiration) should
	// expire.
	TargetExpiryDate param.Field[time.Time] `json:"target_expiry_date,required" format:"date"`
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount"`
	// The ID of the block affected by an expiration_change, used to differentiate
	// between multiple blocks with the same `expiry_date`.
	BlockID param.Field[string] `json:"block_id"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = "expiration_change"
)

func (r AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to void.
	BlockID   param.Field[string]                                                                    `json:"block_id,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Can only be specified when `entry_type=void`. The reason for the void.
	VoidReason param.Field[AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason] `json:"void_reason"`
}

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType = "void"
)

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid:
		return true
	}
	return false
}

// Can only be specified when `entry_type=void`. The reason for the void.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason string

const (
	AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason = "refund"
)

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement or void operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to reverse a decrement from.
	BlockID   param.Field[string]                                                                         `json:"block_id,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType = "amendment"
)

func (r AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestEntryType string

const (
	AddCreditLedgerEntryRequestEntryTypeIncrement        AddCreditLedgerEntryRequestEntryType = "increment"
	AddCreditLedgerEntryRequestEntryTypeDecrement        AddCreditLedgerEntryRequestEntryType = "decrement"
	AddCreditLedgerEntryRequestEntryTypeExpirationChange AddCreditLedgerEntryRequestEntryType = "expiration_change"
	AddCreditLedgerEntryRequestEntryTypeVoid             AddCreditLedgerEntryRequestEntryType = "void"
	AddCreditLedgerEntryRequestEntryTypeAmendment        AddCreditLedgerEntryRequestEntryType = "amendment"
)

func (r AddCreditLedgerEntryRequestEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestEntryTypeIncrement, AddCreditLedgerEntryRequestEntryTypeDecrement, AddCreditLedgerEntryRequestEntryTypeExpirationChange, AddCreditLedgerEntryRequestEntryTypeVoid, AddCreditLedgerEntryRequestEntryTypeAmendment:
		return true
	}
	return false
}

// Can only be specified when `entry_type=void`. The reason for the void.
type AddCreditLedgerEntryRequestVoidReason string

const (
	AddCreditLedgerEntryRequestVoidReasonRefund AddCreditLedgerEntryRequestVoidReason = "refund"
)

func (r AddCreditLedgerEntryRequestVoidReason) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestVoidReasonRefund:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParams struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[AddSubscriptionAdjustmentParamsAdjustmentUnion] `json:"adjustment,required"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The start date of the adjustment interval. This is the date that the adjustment
	// will start affecting prices on the subscription. If null, the adjustment will
	// start when the phase or subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r AddSubscriptionAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type AddSubscriptionAdjustmentParamsAdjustment struct {
	AdjustmentType    param.Field[AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AppliesToPriceIDs param.Field[interface{}]                                             `json:"applies_to_price_ids,required"`
	AmountDiscount    param.Field[string]                                                  `json:"amount_discount"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	UsageDiscount      param.Field[float64] `json:"usage_discount"`
}

func (r AddSubscriptionAdjustmentParamsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionAdjustmentParamsAdjustment) implementsAddSubscriptionAdjustmentParamsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by
// [shared.AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount],
// [shared.AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount],
// [shared.AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount],
// [shared.AddSubscriptionAdjustmentParamsAdjustmentNewMinimum],
// [shared.AddSubscriptionAdjustmentParamsAdjustmentNewMaximum],
// [AddSubscriptionAdjustmentParamsAdjustment].
type AddSubscriptionAdjustmentParamsAdjustmentUnion interface {
	implementsAddSubscriptionAdjustmentParamsAdjustmentUnion()
}

type AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount struct {
	AdjustmentType param.Field[AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs  param.Field[[]string] `json:"applies_to_price_ids,required"`
	PercentageDiscount param.Field[float64]  `json:"percentage_discount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount) implementsAddSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	UsageDiscount     param.Field[float64]  `json:"usage_discount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount) implementsAddSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType string

const (
	AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                   `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount) implementsAddSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType string

const (
	AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParamsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewMinimum) implementsAddSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType string

const (
	AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParamsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	MaximumAmount     param.Field[string]   `json:"maximum_amount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionAdjustmentParamsAdjustmentNewMaximum) implementsAddSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType string

const (
	AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType string

const (
	AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "percentage_discount"
	AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount      AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "usage_discount"
	AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount     AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "amount_discount"
	AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum            AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "minimum"
	AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum            AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "maximum"
)

func (r AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount, AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount, AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount, AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum, AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type AddSubscriptionPriceParams struct {
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[AddSubscriptionPriceParamsAllocationPrice] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for this
	// price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]AddSubscriptionPriceParamsDiscount] `json:"discounts"`
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription. If null, billing will end when the phase or
	// subscription ends.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// this price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// this price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[AddSubscriptionPriceParamsPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
	// The start date of the price interval. This is the date that the price will start
	// billing on the subscription. If null, billing will start when the phase or
	// subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r AddSubscriptionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new allocation price to create and add to the subscription.
type AddSubscriptionPriceParamsAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[AddSubscriptionPriceParamsAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r AddSubscriptionPriceParamsAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type AddSubscriptionPriceParamsAllocationPriceCadence string

const (
	AddSubscriptionPriceParamsAllocationPriceCadenceOneTime    AddSubscriptionPriceParamsAllocationPriceCadence = "one_time"
	AddSubscriptionPriceParamsAllocationPriceCadenceMonthly    AddSubscriptionPriceParamsAllocationPriceCadence = "monthly"
	AddSubscriptionPriceParamsAllocationPriceCadenceQuarterly  AddSubscriptionPriceParamsAllocationPriceCadence = "quarterly"
	AddSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual AddSubscriptionPriceParamsAllocationPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsAllocationPriceCadenceAnnual     AddSubscriptionPriceParamsAllocationPriceCadence = "annual"
	AddSubscriptionPriceParamsAllocationPriceCadenceCustom     AddSubscriptionPriceParamsAllocationPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsAllocationPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsAllocationPriceCadenceOneTime, AddSubscriptionPriceParamsAllocationPriceCadenceMonthly, AddSubscriptionPriceParamsAllocationPriceCadenceQuarterly, AddSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual, AddSubscriptionPriceParamsAllocationPriceCadenceAnnual, AddSubscriptionPriceParamsAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsDiscount struct {
	DiscountType param.Field[AddSubscriptionPriceParamsDiscountsDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r AddSubscriptionPriceParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsDiscountsDiscountType string

const (
	AddSubscriptionPriceParamsDiscountsDiscountTypePercentage AddSubscriptionPriceParamsDiscountsDiscountType = "percentage"
	AddSubscriptionPriceParamsDiscountsDiscountTypeUsage      AddSubscriptionPriceParamsDiscountsDiscountType = "usage"
	AddSubscriptionPriceParamsDiscountsDiscountTypeAmount     AddSubscriptionPriceParamsDiscountsDiscountType = "amount"
)

func (r AddSubscriptionPriceParamsDiscountsDiscountType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsDiscountsDiscountTypePercentage, AddSubscriptionPriceParamsDiscountsDiscountTypeUsage, AddSubscriptionPriceParamsDiscountsDiscountTypeAmount:
		return true
	}
	return false
}

// The definition of a new price to create and add to the subscription.
type AddSubscriptionPriceParamsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                   `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceModelType] `json:"model_type,required"`
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
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey          param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                param.Field[interface{}] `json:"matrix_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}] `json:"metadata"`
	PackageConfig               param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}] `json:"package_with_allocation_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]      `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r AddSubscriptionPriceParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPrice) implementsAddSubscriptionPriceParamsPriceUnion() {}

// The definition of a new price to create and add to the subscription.
//
// Satisfied by [shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice],
// [shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice],
// [AddSubscriptionPriceParamsPrice].
type AddSubscriptionPriceParamsPriceUnion interface {
	implementsAddSubscriptionPriceParamsPriceUnion()
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                           `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                            `json:"name,required"`
	UnitConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType = "unit"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                  `json:"name,required"`
	PackageConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType = "package"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                `json:"item_id,required"`
	MatrixConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType = "matrix"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                `json:"name,required"`
	TieredConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType = "tiered"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                      `json:"name,required"`
	TieredBpsConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType = "tiered_bps"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPrice struct {
	BpsConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType = "bps"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice struct {
	BulkBpsConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType = "bulk_bps"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPrice struct {
	BulkConfig param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                           `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType = "bulk"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                    `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType = "tiered_package"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                      `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                               `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                      `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                          `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                            `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                            `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                              `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                            `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                            `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                         `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice) implementsAddSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly  AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime    AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom     AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type AddSubscriptionPriceParamsPriceCadence string

const (
	AddSubscriptionPriceParamsPriceCadenceAnnual     AddSubscriptionPriceParamsPriceCadence = "annual"
	AddSubscriptionPriceParamsPriceCadenceSemiAnnual AddSubscriptionPriceParamsPriceCadence = "semi_annual"
	AddSubscriptionPriceParamsPriceCadenceMonthly    AddSubscriptionPriceParamsPriceCadence = "monthly"
	AddSubscriptionPriceParamsPriceCadenceQuarterly  AddSubscriptionPriceParamsPriceCadence = "quarterly"
	AddSubscriptionPriceParamsPriceCadenceOneTime    AddSubscriptionPriceParamsPriceCadence = "one_time"
	AddSubscriptionPriceParamsPriceCadenceCustom     AddSubscriptionPriceParamsPriceCadence = "custom"
)

func (r AddSubscriptionPriceParamsPriceCadence) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceCadenceAnnual, AddSubscriptionPriceParamsPriceCadenceSemiAnnual, AddSubscriptionPriceParamsPriceCadenceMonthly, AddSubscriptionPriceParamsPriceCadenceQuarterly, AddSubscriptionPriceParamsPriceCadenceOneTime, AddSubscriptionPriceParamsPriceCadenceCustom:
		return true
	}
	return false
}

type AddSubscriptionPriceParamsPriceModelType string

const (
	AddSubscriptionPriceParamsPriceModelTypeUnit                            AddSubscriptionPriceParamsPriceModelType = "unit"
	AddSubscriptionPriceParamsPriceModelTypePackage                         AddSubscriptionPriceParamsPriceModelType = "package"
	AddSubscriptionPriceParamsPriceModelTypeMatrix                          AddSubscriptionPriceParamsPriceModelType = "matrix"
	AddSubscriptionPriceParamsPriceModelTypeTiered                          AddSubscriptionPriceParamsPriceModelType = "tiered"
	AddSubscriptionPriceParamsPriceModelTypeTieredBps                       AddSubscriptionPriceParamsPriceModelType = "tiered_bps"
	AddSubscriptionPriceParamsPriceModelTypeBps                             AddSubscriptionPriceParamsPriceModelType = "bps"
	AddSubscriptionPriceParamsPriceModelTypeBulkBps                         AddSubscriptionPriceParamsPriceModelType = "bulk_bps"
	AddSubscriptionPriceParamsPriceModelTypeBulk                            AddSubscriptionPriceParamsPriceModelType = "bulk"
	AddSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount            AddSubscriptionPriceParamsPriceModelType = "threshold_total_amount"
	AddSubscriptionPriceParamsPriceModelTypeTieredPackage                   AddSubscriptionPriceParamsPriceModelType = "tiered_package"
	AddSubscriptionPriceParamsPriceModelTypeTieredWithMinimum               AddSubscriptionPriceParamsPriceModelType = "tiered_with_minimum"
	AddSubscriptionPriceParamsPriceModelTypeUnitWithPercent                 AddSubscriptionPriceParamsPriceModelType = "unit_with_percent"
	AddSubscriptionPriceParamsPriceModelTypePackageWithAllocation           AddSubscriptionPriceParamsPriceModelType = "package_with_allocation"
	AddSubscriptionPriceParamsPriceModelTypeTieredWithProration             AddSubscriptionPriceParamsPriceModelType = "tiered_with_proration"
	AddSubscriptionPriceParamsPriceModelTypeUnitWithProration               AddSubscriptionPriceParamsPriceModelType = "unit_with_proration"
	AddSubscriptionPriceParamsPriceModelTypeGroupedAllocation               AddSubscriptionPriceParamsPriceModelType = "grouped_allocation"
	AddSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum      AddSubscriptionPriceParamsPriceModelType = "grouped_with_prorated_minimum"
	AddSubscriptionPriceParamsPriceModelTypeBulkWithProration               AddSubscriptionPriceParamsPriceModelType = "bulk_with_proration"
	AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing   AddSubscriptionPriceParamsPriceModelType = "scalable_matrix_with_unit_pricing"
	AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing AddSubscriptionPriceParamsPriceModelType = "scalable_matrix_with_tiered_pricing"
	AddSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk           AddSubscriptionPriceParamsPriceModelType = "cumulative_grouped_bulk"
	AddSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage           AddSubscriptionPriceParamsPriceModelType = "max_group_tiered_package"
	AddSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum       AddSubscriptionPriceParamsPriceModelType = "grouped_with_metered_minimum"
	AddSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName           AddSubscriptionPriceParamsPriceModelType = "matrix_with_display_name"
	AddSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage            AddSubscriptionPriceParamsPriceModelType = "grouped_tiered_package"
)

func (r AddSubscriptionPriceParamsPriceModelType) IsKnown() bool {
	switch r {
	case AddSubscriptionPriceParamsPriceModelTypeUnit, AddSubscriptionPriceParamsPriceModelTypePackage, AddSubscriptionPriceParamsPriceModelTypeMatrix, AddSubscriptionPriceParamsPriceModelTypeTiered, AddSubscriptionPriceParamsPriceModelTypeTieredBps, AddSubscriptionPriceParamsPriceModelTypeBps, AddSubscriptionPriceParamsPriceModelTypeBulkBps, AddSubscriptionPriceParamsPriceModelTypeBulk, AddSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount, AddSubscriptionPriceParamsPriceModelTypeTieredPackage, AddSubscriptionPriceParamsPriceModelTypeTieredWithMinimum, AddSubscriptionPriceParamsPriceModelTypeUnitWithPercent, AddSubscriptionPriceParamsPriceModelTypePackageWithAllocation, AddSubscriptionPriceParamsPriceModelTypeTieredWithProration, AddSubscriptionPriceParamsPriceModelTypeUnitWithProration, AddSubscriptionPriceParamsPriceModelTypeGroupedAllocation, AddSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum, AddSubscriptionPriceParamsPriceModelTypeBulkWithProration, AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing, AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing, AddSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk, AddSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage, AddSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum, AddSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName, AddSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{}          `json:"applies_to_price_ids,required"`
	DiscountType      DiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
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
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
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
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.TrialDiscount], [shared.DiscountUsageDiscount], [shared.AmountDiscount].
func (r Discount) AsUnion() DiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount], [shared.TrialDiscount],
// [shared.DiscountUsageDiscount] or [shared.AmountDiscount].
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

func (r DiscountUsageDiscount) ImplementsDiscount() {}

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
	AppliesToPriceIDs param.Field[interface{}]          `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[DiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
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
// [shared.DiscountUsageDiscountParam], [shared.AmountDiscountParam],
// [DiscountParam].
type DiscountUnionParam interface {
	ImplementsDiscountUnionParam()
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

func (r DiscountUsageDiscountParam) ImplementsDiscountUnionParam() {}

type InvoiceLevelDiscount struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{}                      `json:"applies_to_price_ids,required"`
	DiscountType      InvoiceLevelDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
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
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
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
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.AmountDiscount], [shared.TrialDiscount].
func (r InvoiceLevelDiscount) AsUnion() InvoiceLevelDiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount], [shared.AmountDiscount] or
// [shared.TrialDiscount].
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

func (r PercentageDiscountParam) ImplementsDiscountUnionParam() {}

type RemoveSubscriptionAdjustmentParams struct {
	// The id of the adjustment to remove on the subscription.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
}

func (r RemoveSubscriptionAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RemoveSubscriptionPriceParams struct {
	// The external price id of the price to remove on the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The id of the price to remove on the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r RemoveSubscriptionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionAdjustmentParams struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the subscription.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
}

func (r ReplaceSubscriptionAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type ReplaceSubscriptionAdjustmentParamsAdjustment struct {
	AdjustmentType    param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AppliesToPriceIDs param.Field[interface{}]                                                 `json:"applies_to_price_ids,required"`
	AmountDiscount    param.Field[string]                                                      `json:"amount_discount"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	UsageDiscount      param.Field[float64] `json:"usage_discount"`
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustment) implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by
// [shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount],
// [shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount],
// [shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount],
// [shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimum],
// [shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximum],
// [ReplaceSubscriptionAdjustmentParamsAdjustment].
type ReplaceSubscriptionAdjustmentParamsAdjustmentUnion interface {
	implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion()
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount struct {
	AdjustmentType param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs  param.Field[[]string] `json:"applies_to_price_ids,required"`
	PercentageDiscount param.Field[float64]  `json:"percentage_discount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount) implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	UsageDiscount     param.Field[float64]  `json:"usage_discount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount) implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType string

const (
	ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                       `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount) implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType string

const (
	ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimum) implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType string

const (
	ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	MaximumAmount     param.Field[string]   `json:"maximum_amount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximum) implementsReplaceSubscriptionAdjustmentParamsAdjustmentUnion() {
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType string

const (
	ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType string

const (
	ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "percentage_discount"
	ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount      ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "usage_discount"
	ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount     ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "amount_discount"
	ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum            ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "minimum"
	ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum            ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType = "maximum"
)

func (r ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount, ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount, ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount, ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum, ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParams struct {
	// The id of the price on the plan to replace in the subscription.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[ReplaceSubscriptionPriceParamsAllocationPrice] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for the
	// replacement price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]ReplaceSubscriptionPriceParamsDiscount] `json:"discounts"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The new quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[ReplaceSubscriptionPriceParamsPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r ReplaceSubscriptionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new allocation price to create and add to the subscription.
type ReplaceSubscriptionPriceParamsAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[ReplaceSubscriptionPriceParamsAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r ReplaceSubscriptionPriceParamsAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type ReplaceSubscriptionPriceParamsAllocationPriceCadence string

const (
	ReplaceSubscriptionPriceParamsAllocationPriceCadenceOneTime    ReplaceSubscriptionPriceParamsAllocationPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsAllocationPriceCadenceMonthly    ReplaceSubscriptionPriceParamsAllocationPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsAllocationPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsAllocationPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsAllocationPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsAllocationPriceCadenceAnnual     ReplaceSubscriptionPriceParamsAllocationPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsAllocationPriceCadenceCustom     ReplaceSubscriptionPriceParamsAllocationPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsAllocationPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsAllocationPriceCadenceOneTime, ReplaceSubscriptionPriceParamsAllocationPriceCadenceMonthly, ReplaceSubscriptionPriceParamsAllocationPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsAllocationPriceCadenceAnnual, ReplaceSubscriptionPriceParamsAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsDiscount struct {
	DiscountType param.Field[ReplaceSubscriptionPriceParamsDiscountsDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r ReplaceSubscriptionPriceParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsDiscountsDiscountType string

const (
	ReplaceSubscriptionPriceParamsDiscountsDiscountTypePercentage ReplaceSubscriptionPriceParamsDiscountsDiscountType = "percentage"
	ReplaceSubscriptionPriceParamsDiscountsDiscountTypeUsage      ReplaceSubscriptionPriceParamsDiscountsDiscountType = "usage"
	ReplaceSubscriptionPriceParamsDiscountsDiscountTypeAmount     ReplaceSubscriptionPriceParamsDiscountsDiscountType = "amount"
)

func (r ReplaceSubscriptionPriceParamsDiscountsDiscountType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsDiscountsDiscountTypePercentage, ReplaceSubscriptionPriceParamsDiscountsDiscountTypeUsage, ReplaceSubscriptionPriceParamsDiscountsDiscountTypeAmount:
		return true
	}
	return false
}

// The definition of a new price to create and add to the subscription.
type ReplaceSubscriptionPriceParamsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceModelType] `json:"model_type,required"`
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
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey          param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                param.Field[interface{}] `json:"matrix_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}] `json:"metadata"`
	PackageConfig               param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}] `json:"package_with_allocation_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]      `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r ReplaceSubscriptionPriceParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {}

// The definition of a new price to create and add to the subscription.
//
// Satisfied by
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice],
// [shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice],
// [ReplaceSubscriptionPriceParamsPrice].
type ReplaceSubscriptionPriceParamsPriceUnion interface {
	implementsReplaceSubscriptionPriceParamsPriceUnion()
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                `json:"name,required"`
	UnitConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType = "unit"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                      `json:"name,required"`
	PackageConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType = "package"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                    `json:"item_id,required"`
	MatrixConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType = "matrix"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                    `json:"name,required"`
	TieredConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType = "tiered"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                    `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                          `json:"name,required"`
	TieredBpsConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType = "tiered_bps"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPrice struct {
	BpsConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType = "bps"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice struct {
	BulkBpsConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType = "bulk_bps"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPrice struct {
	BulkConfig param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType = "bulk"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType = "tiered_package"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                          `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                   `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                        `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                          `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                              `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                                `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                  `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                                `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                             `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice) implementsReplaceSubscriptionPriceParamsPriceUnion() {
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type ReplaceSubscriptionPriceParamsPriceCadence string

const (
	ReplaceSubscriptionPriceParamsPriceCadenceAnnual     ReplaceSubscriptionPriceParamsPriceCadence = "annual"
	ReplaceSubscriptionPriceParamsPriceCadenceSemiAnnual ReplaceSubscriptionPriceParamsPriceCadence = "semi_annual"
	ReplaceSubscriptionPriceParamsPriceCadenceMonthly    ReplaceSubscriptionPriceParamsPriceCadence = "monthly"
	ReplaceSubscriptionPriceParamsPriceCadenceQuarterly  ReplaceSubscriptionPriceParamsPriceCadence = "quarterly"
	ReplaceSubscriptionPriceParamsPriceCadenceOneTime    ReplaceSubscriptionPriceParamsPriceCadence = "one_time"
	ReplaceSubscriptionPriceParamsPriceCadenceCustom     ReplaceSubscriptionPriceParamsPriceCadence = "custom"
)

func (r ReplaceSubscriptionPriceParamsPriceCadence) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceCadenceAnnual, ReplaceSubscriptionPriceParamsPriceCadenceSemiAnnual, ReplaceSubscriptionPriceParamsPriceCadenceMonthly, ReplaceSubscriptionPriceParamsPriceCadenceQuarterly, ReplaceSubscriptionPriceParamsPriceCadenceOneTime, ReplaceSubscriptionPriceParamsPriceCadenceCustom:
		return true
	}
	return false
}

type ReplaceSubscriptionPriceParamsPriceModelType string

const (
	ReplaceSubscriptionPriceParamsPriceModelTypeUnit                            ReplaceSubscriptionPriceParamsPriceModelType = "unit"
	ReplaceSubscriptionPriceParamsPriceModelTypePackage                         ReplaceSubscriptionPriceParamsPriceModelType = "package"
	ReplaceSubscriptionPriceParamsPriceModelTypeMatrix                          ReplaceSubscriptionPriceParamsPriceModelType = "matrix"
	ReplaceSubscriptionPriceParamsPriceModelTypeTiered                          ReplaceSubscriptionPriceParamsPriceModelType = "tiered"
	ReplaceSubscriptionPriceParamsPriceModelTypeTieredBps                       ReplaceSubscriptionPriceParamsPriceModelType = "tiered_bps"
	ReplaceSubscriptionPriceParamsPriceModelTypeBps                             ReplaceSubscriptionPriceParamsPriceModelType = "bps"
	ReplaceSubscriptionPriceParamsPriceModelTypeBulkBps                         ReplaceSubscriptionPriceParamsPriceModelType = "bulk_bps"
	ReplaceSubscriptionPriceParamsPriceModelTypeBulk                            ReplaceSubscriptionPriceParamsPriceModelType = "bulk"
	ReplaceSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount            ReplaceSubscriptionPriceParamsPriceModelType = "threshold_total_amount"
	ReplaceSubscriptionPriceParamsPriceModelTypeTieredPackage                   ReplaceSubscriptionPriceParamsPriceModelType = "tiered_package"
	ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithMinimum               ReplaceSubscriptionPriceParamsPriceModelType = "tiered_with_minimum"
	ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithPercent                 ReplaceSubscriptionPriceParamsPriceModelType = "unit_with_percent"
	ReplaceSubscriptionPriceParamsPriceModelTypePackageWithAllocation           ReplaceSubscriptionPriceParamsPriceModelType = "package_with_allocation"
	ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithProration             ReplaceSubscriptionPriceParamsPriceModelType = "tiered_with_proration"
	ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithProration               ReplaceSubscriptionPriceParamsPriceModelType = "unit_with_proration"
	ReplaceSubscriptionPriceParamsPriceModelTypeGroupedAllocation               ReplaceSubscriptionPriceParamsPriceModelType = "grouped_allocation"
	ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum      ReplaceSubscriptionPriceParamsPriceModelType = "grouped_with_prorated_minimum"
	ReplaceSubscriptionPriceParamsPriceModelTypeBulkWithProration               ReplaceSubscriptionPriceParamsPriceModelType = "bulk_with_proration"
	ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing   ReplaceSubscriptionPriceParamsPriceModelType = "scalable_matrix_with_unit_pricing"
	ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing ReplaceSubscriptionPriceParamsPriceModelType = "scalable_matrix_with_tiered_pricing"
	ReplaceSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk           ReplaceSubscriptionPriceParamsPriceModelType = "cumulative_grouped_bulk"
	ReplaceSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage           ReplaceSubscriptionPriceParamsPriceModelType = "max_group_tiered_package"
	ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum       ReplaceSubscriptionPriceParamsPriceModelType = "grouped_with_metered_minimum"
	ReplaceSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName           ReplaceSubscriptionPriceParamsPriceModelType = "matrix_with_display_name"
	ReplaceSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage            ReplaceSubscriptionPriceParamsPriceModelType = "grouped_tiered_package"
)

func (r ReplaceSubscriptionPriceParamsPriceModelType) IsKnown() bool {
	switch r {
	case ReplaceSubscriptionPriceParamsPriceModelTypeUnit, ReplaceSubscriptionPriceParamsPriceModelTypePackage, ReplaceSubscriptionPriceParamsPriceModelTypeMatrix, ReplaceSubscriptionPriceParamsPriceModelTypeTiered, ReplaceSubscriptionPriceParamsPriceModelTypeTieredBps, ReplaceSubscriptionPriceParamsPriceModelTypeBps, ReplaceSubscriptionPriceParamsPriceModelTypeBulkBps, ReplaceSubscriptionPriceParamsPriceModelTypeBulk, ReplaceSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount, ReplaceSubscriptionPriceParamsPriceModelTypeTieredPackage, ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithMinimum, ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithPercent, ReplaceSubscriptionPriceParamsPriceModelTypePackageWithAllocation, ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithProration, ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithProration, ReplaceSubscriptionPriceParamsPriceModelTypeGroupedAllocation, ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum, ReplaceSubscriptionPriceParamsPriceModelTypeBulkWithProration, ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing, ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing, ReplaceSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk, ReplaceSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage, ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum, ReplaceSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName, ReplaceSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

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

func (r TrialDiscountParam) ImplementsDiscountUnionParam() {}
