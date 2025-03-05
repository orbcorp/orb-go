// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"github.com/orbcorp/orb-go/internal/apierror"
	"github.com/orbcorp/orb-go/shared"
)

type Error = apierror.Error
type ErrorStatus = apierror.ErrorStatus

const ErrorStatus500 = apierror.ErrorStatus500
const ErrorStatus429 = apierror.ErrorStatus429
const ErrorStatus413 = apierror.ErrorStatus413
const ErrorStatus409 = apierror.ErrorStatus409
const ErrorStatus404 = apierror.ErrorStatus404
const ErrorStatus400 = apierror.ErrorStatus400
const ErrorStatus401 = apierror.ErrorStatus401

type ErrorType = apierror.ErrorType

const ErrorTypeOrbInternalServerError = apierror.ErrorTypeOrbInternalServerError
const ErrorTypeTooManyRequests = apierror.ErrorTypeTooManyRequests
const ErrorTypeResourceTooLarge = apierror.ErrorTypeResourceTooLarge
const ErrorTypeRequestTooLarge = apierror.ErrorTypeRequestTooLarge
const ErrorTypeResourceConflict = apierror.ErrorTypeResourceConflict
const ErrorTypeURLNotFound = apierror.ErrorTypeURLNotFound
const ErrorTypeResourceNotFound = apierror.ErrorTypeResourceNotFound
const ErrorTypeFeatureNotAvailable = apierror.ErrorTypeFeatureNotAvailable
const ErrorTypeOrbAuthenticationError = apierror.ErrorTypeOrbAuthenticationError
const ErrorTypeRequestValidationError = apierror.ErrorTypeRequestValidationError
const ErrorTypeDuplicateResourceCreation = apierror.ErrorTypeDuplicateResourceCreation
const ErrorTypeConstraintViolation = apierror.ErrorTypeConstraintViolation

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestUnionParam = shared.AddCreditLedgerEntryRequestUnionParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
//
// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement = shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange = shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid

// Can only be specified when `entry_type=void`. The reason for the void.
//
// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment = shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestEntryType = shared.AddCreditLedgerEntryRequestEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeIncrement = shared.AddCreditLedgerEntryRequestEntryTypeIncrement

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeDecrement = shared.AddCreditLedgerEntryRequestEntryTypeDecrement

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeExpirationChange = shared.AddCreditLedgerEntryRequestEntryTypeExpirationChange

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeVoid = shared.AddCreditLedgerEntryRequestEntryTypeVoid

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeAmendment = shared.AddCreditLedgerEntryRequestEntryTypeAmendment

// Can only be specified when `entry_type=void`. The reason for the void.
//
// This is an alias to an internal type.
type AddCreditLedgerEntryRequestVoidReason = shared.AddCreditLedgerEntryRequestVoidReason

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestVoidReasonRefund = shared.AddCreditLedgerEntryRequestVoidReasonRefund

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParams = shared.AddSubscriptionAdjustmentParams

// This is an alias to an internal type.
type AddSubscriptionPriceParams = shared.AddSubscriptionPriceParams

// This is an alias to an internal type.
type AddressInputModelParam = shared.AddressInputModelParam

// This is an alias to an internal type.
type AddressModel = shared.AddressModel

// This is an alias to an internal type.
type AdjustmentIntervalModel = shared.AdjustmentIntervalModel

// This is an alias to an internal type.
type AdjustmentModel = shared.AdjustmentModel

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseUsageDiscountAdjustment = shared.AdjustmentModelPlanPhaseUsageDiscountAdjustment

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentType = shared.AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount = shared.AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseAmountDiscountAdjustment = shared.AdjustmentModelPlanPhaseAmountDiscountAdjustment

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentType = shared.AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount = shared.AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type AdjustmentModelPlanPhasePercentageDiscountAdjustment = shared.AdjustmentModelPlanPhasePercentageDiscountAdjustment

// This is an alias to an internal type.
type AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentType = shared.AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount = shared.AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseMinimumAdjustment = shared.AdjustmentModelPlanPhaseMinimumAdjustment

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentType = shared.AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentType

// This is an alias to an internal value.
const AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum = shared.AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseMaximumAdjustment = shared.AdjustmentModelPlanPhaseMaximumAdjustment

// This is an alias to an internal type.
type AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentType = shared.AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentType

// This is an alias to an internal value.
const AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum = shared.AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type AdjustmentModelAdjustmentType = shared.AdjustmentModelAdjustmentType

// This is an alias to an internal value.
const AdjustmentModelAdjustmentTypeUsageDiscount = shared.AdjustmentModelAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const AdjustmentModelAdjustmentTypeAmountDiscount = shared.AdjustmentModelAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const AdjustmentModelAdjustmentTypePercentageDiscount = shared.AdjustmentModelAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const AdjustmentModelAdjustmentTypeMinimum = shared.AdjustmentModelAdjustmentTypeMinimum

// This is an alias to an internal value.
const AdjustmentModelAdjustmentTypeMaximum = shared.AdjustmentModelAdjustmentTypeMaximum

// This is an alias to an internal type.
type AffectedBlockModel = shared.AffectedBlockModel

// This is an alias to an internal type.
type AggregatedCostModel = shared.AggregatedCostModel

// This is an alias to an internal type.
type AggregatedCostModelPerPriceCost = shared.AggregatedCostModelPerPriceCost

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
//
// This is an alias to an internal type.
type AlertModel = shared.AlertModel

// The metric the alert applies to.
//
// This is an alias to an internal type.
type AlertModelMetric = shared.AlertModelMetric

// The plan the alert applies to.
//
// This is an alias to an internal type.
type AlertModelPlan = shared.AlertModelPlan

// The type of alert. This must be a valid alert type.
//
// This is an alias to an internal type.
type AlertModelType = shared.AlertModelType

// This is an alias to an internal value.
const AlertModelTypeUsageExceeded = shared.AlertModelTypeUsageExceeded

// This is an alias to an internal value.
const AlertModelTypeCostExceeded = shared.AlertModelTypeCostExceeded

// This is an alias to an internal value.
const AlertModelTypeCreditBalanceDepleted = shared.AlertModelTypeCreditBalanceDepleted

// This is an alias to an internal value.
const AlertModelTypeCreditBalanceDropped = shared.AlertModelTypeCreditBalanceDropped

// This is an alias to an internal value.
const AlertModelTypeCreditBalanceRecovered = shared.AlertModelTypeCreditBalanceRecovered

// This is an alias to an internal type.
type AllocationModel = shared.AllocationModel

// This is an alias to an internal type.
type AmountDiscount = shared.AmountDiscount

// This is an alias to an internal type.
type AmountDiscountDiscountType = shared.AmountDiscountDiscountType

// This is an alias to an internal value.
const AmountDiscountDiscountTypeAmount = shared.AmountDiscountDiscountTypeAmount

// This is an alias to an internal type.
type AmountDiscountParam = shared.AmountDiscountParam

// This is an alias to an internal type.
type AmountDiscountIntervalModel = shared.AmountDiscountIntervalModel

// This is an alias to an internal type.
type AmountDiscountIntervalModelDiscountType = shared.AmountDiscountIntervalModelDiscountType

// This is an alias to an internal value.
const AmountDiscountIntervalModelDiscountTypeAmount = shared.AmountDiscountIntervalModelDiscountTypeAmount

// This is an alias to an internal type.
type AutoCollectionModel = shared.AutoCollectionModel

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
//
// This is an alias to an internal type.
type BackfillModel = shared.BackfillModel

// The status of the backfill.
//
// This is an alias to an internal type.
type BackfillModelStatus = shared.BackfillModelStatus

// This is an alias to an internal value.
const BackfillModelStatusPending = shared.BackfillModelStatusPending

// This is an alias to an internal value.
const BackfillModelStatusReflected = shared.BackfillModelStatusReflected

// This is an alias to an internal value.
const BackfillModelStatusPendingRevert = shared.BackfillModelStatusPendingRevert

// This is an alias to an internal value.
const BackfillModelStatusReverted = shared.BackfillModelStatusReverted

// The Metric resource represents a calculation of a quantity based on events.
// Metrics are defined by the query that transforms raw usage events into
// meaningful values for your customers.
//
// This is an alias to an internal type.
type BillableMetricModel = shared.BillableMetricModel

// This is an alias to an internal type.
type BillableMetricModelStatus = shared.BillableMetricModelStatus

// This is an alias to an internal value.
const BillableMetricModelStatusActive = shared.BillableMetricModelStatusActive

// This is an alias to an internal value.
const BillableMetricModelStatusDraft = shared.BillableMetricModelStatusDraft

// This is an alias to an internal value.
const BillableMetricModelStatusArchived = shared.BillableMetricModelStatusArchived

// This is an alias to an internal type.
type BillableMetricSimpleModel = shared.BillableMetricSimpleModel

// This is an alias to an internal type.
type BillableMetricTinyModel = shared.BillableMetricTinyModel

// This is an alias to an internal type.
type BillingCycleAnchorConfigurationModel = shared.BillingCycleAnchorConfigurationModel

// This is an alias to an internal type.
type BillingCycleAnchorConfigurationModelParam = shared.BillingCycleAnchorConfigurationModelParam

// This is an alias to an internal type.
type BillingCycleConfigurationModel = shared.BillingCycleConfigurationModel

// This is an alias to an internal type.
type BillingCycleConfigurationModelDurationUnit = shared.BillingCycleConfigurationModelDurationUnit

// This is an alias to an internal value.
const BillingCycleConfigurationModelDurationUnitDay = shared.BillingCycleConfigurationModelDurationUnitDay

// This is an alias to an internal value.
const BillingCycleConfigurationModelDurationUnitMonth = shared.BillingCycleConfigurationModelDurationUnitMonth

// This is an alias to an internal type.
type BillingCycleRelativeDate = shared.BillingCycleRelativeDate

// This is an alias to an internal value.
const BillingCycleRelativeDateStartOfTerm = shared.BillingCycleRelativeDateStartOfTerm

// This is an alias to an internal value.
const BillingCycleRelativeDateEndOfTerm = shared.BillingCycleRelativeDateEndOfTerm

// This is an alias to an internal type.
type BpsConfigModel = shared.BpsConfigModel

// This is an alias to an internal type.
type BpsConfigModelParam = shared.BpsConfigModelParam

// This is an alias to an internal type.
type BulkBpsConfigModel = shared.BulkBpsConfigModel

// This is an alias to an internal type.
type BulkBpsConfigModelTier = shared.BulkBpsConfigModelTier

// This is an alias to an internal type.
type BulkBpsConfigModelParam = shared.BulkBpsConfigModelParam

// This is an alias to an internal type.
type BulkBpsConfigModelTierParam = shared.BulkBpsConfigModelTierParam

// This is an alias to an internal type.
type BulkConfigModel = shared.BulkConfigModel

// This is an alias to an internal type.
type BulkConfigModelTier = shared.BulkConfigModelTier

// This is an alias to an internal type.
type BulkConfigModelParam = shared.BulkConfigModelParam

// This is an alias to an internal type.
type BulkConfigModelTierParam = shared.BulkConfigModelTierParam

// A coupon represents a reusable discount configuration that can be applied either
// as a fixed or percentage amount to an invoice or subscription. Coupons are
// activated using a redemption code, which applies the discount to a subscription
// or invoice. The duration of a coupon determines how long it remains available
// for use by end users.
//
// This is an alias to an internal type.
type CouponModel = shared.CouponModel

// This is an alias to an internal type.
type CouponModelDiscount = shared.CouponModelDiscount

// This is an alias to an internal type.
type CouponModelDiscountDiscountType = shared.CouponModelDiscountDiscountType

// This is an alias to an internal value.
const CouponModelDiscountDiscountTypePercentage = shared.CouponModelDiscountDiscountTypePercentage

// This is an alias to an internal value.
const CouponModelDiscountDiscountTypeAmount = shared.CouponModelDiscountDiscountTypeAmount

// This is an alias to an internal type.
type CouponRedemptionModel = shared.CouponRedemptionModel

// This is an alias to an internal type.
type CreditLedgerEntriesModel = shared.CreditLedgerEntriesModel

// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
// credits within Orb.
//
// This is an alias to an internal type.
type CreditLedgerEntryModel = shared.CreditLedgerEntryModel

// This is an alias to an internal type.
type CreditLedgerEntryModelIncrementLedgerEntry = shared.CreditLedgerEntryModelIncrementLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelIncrementLedgerEntryEntryStatus = shared.CreditLedgerEntryModelIncrementLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelIncrementLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelIncrementLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelIncrementLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelIncrementLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelIncrementLedgerEntryEntryType = shared.CreditLedgerEntryModelIncrementLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelIncrementLedgerEntryEntryTypeIncrement = shared.CreditLedgerEntryModelIncrementLedgerEntryEntryTypeIncrement

// This is an alias to an internal type.
type CreditLedgerEntryModelDecrementLedgerEntry = shared.CreditLedgerEntryModelDecrementLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelDecrementLedgerEntryEntryStatus = shared.CreditLedgerEntryModelDecrementLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelDecrementLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelDecrementLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelDecrementLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelDecrementLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelDecrementLedgerEntryEntryType = shared.CreditLedgerEntryModelDecrementLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelDecrementLedgerEntryEntryTypeDecrement = shared.CreditLedgerEntryModelDecrementLedgerEntryEntryTypeDecrement

// This is an alias to an internal type.
type CreditLedgerEntryModelExpirationChangeLedgerEntry = shared.CreditLedgerEntryModelExpirationChangeLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus = shared.CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelExpirationChangeLedgerEntryEntryType = shared.CreditLedgerEntryModelExpirationChangeLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelExpirationChangeLedgerEntryEntryTypeExpirationChange = shared.CreditLedgerEntryModelExpirationChangeLedgerEntryEntryTypeExpirationChange

// This is an alias to an internal type.
type CreditLedgerEntryModelCreditBlockExpiryLedgerEntry = shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus = shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryType = shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry = shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry

// This is an alias to an internal type.
type CreditLedgerEntryModelVoidLedgerEntry = shared.CreditLedgerEntryModelVoidLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelVoidLedgerEntryEntryStatus = shared.CreditLedgerEntryModelVoidLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelVoidLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelVoidLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelVoidLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelVoidLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelVoidLedgerEntryEntryType = shared.CreditLedgerEntryModelVoidLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelVoidLedgerEntryEntryTypeVoid = shared.CreditLedgerEntryModelVoidLedgerEntryEntryTypeVoid

// This is an alias to an internal type.
type CreditLedgerEntryModelVoidInitiatedLedgerEntry = shared.CreditLedgerEntryModelVoidInitiatedLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus = shared.CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryType = shared.CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryTypeVoidInitiated = shared.CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryTypeVoidInitiated

// This is an alias to an internal type.
type CreditLedgerEntryModelAmendmentLedgerEntry = shared.CreditLedgerEntryModelAmendmentLedgerEntry

// This is an alias to an internal type.
type CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus = shared.CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusCommitted = shared.CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusPending = shared.CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelAmendmentLedgerEntryEntryType = shared.CreditLedgerEntryModelAmendmentLedgerEntryEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelAmendmentLedgerEntryEntryTypeAmendment = shared.CreditLedgerEntryModelAmendmentLedgerEntryEntryTypeAmendment

// This is an alias to an internal type.
type CreditLedgerEntryModelEntryStatus = shared.CreditLedgerEntryModelEntryStatus

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryStatusCommitted = shared.CreditLedgerEntryModelEntryStatusCommitted

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryStatusPending = shared.CreditLedgerEntryModelEntryStatusPending

// This is an alias to an internal type.
type CreditLedgerEntryModelEntryType = shared.CreditLedgerEntryModelEntryType

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeIncrement = shared.CreditLedgerEntryModelEntryTypeIncrement

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeDecrement = shared.CreditLedgerEntryModelEntryTypeDecrement

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeExpirationChange = shared.CreditLedgerEntryModelEntryTypeExpirationChange

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeCreditBlockExpiry = shared.CreditLedgerEntryModelEntryTypeCreditBlockExpiry

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeVoid = shared.CreditLedgerEntryModelEntryTypeVoid

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeVoidInitiated = shared.CreditLedgerEntryModelEntryTypeVoidInitiated

// This is an alias to an internal value.
const CreditLedgerEntryModelEntryTypeAmendment = shared.CreditLedgerEntryModelEntryTypeAmendment

// This is an alias to an internal type.
type CreditNoteDiscountModel = shared.CreditNoteDiscountModel

// This is an alias to an internal type.
type CreditNoteDiscountModelDiscountType = shared.CreditNoteDiscountModelDiscountType

// This is an alias to an internal value.
const CreditNoteDiscountModelDiscountTypePercentage = shared.CreditNoteDiscountModelDiscountTypePercentage

// This is an alias to an internal type.
type CreditNoteDiscountModelAppliesToPrice = shared.CreditNoteDiscountModelAppliesToPrice

// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
// been applied to a particular invoice.
//
// This is an alias to an internal type.
type CreditNoteModel = shared.CreditNoteModel

// This is an alias to an internal type.
type CreditNoteModelLineItem = shared.CreditNoteModelLineItem

// This is an alias to an internal type.
type CreditNoteModelLineItemsDiscount = shared.CreditNoteModelLineItemsDiscount

// This is an alias to an internal type.
type CreditNoteModelLineItemsDiscountsDiscountType = shared.CreditNoteModelLineItemsDiscountsDiscountType

// This is an alias to an internal value.
const CreditNoteModelLineItemsDiscountsDiscountTypePercentage = shared.CreditNoteModelLineItemsDiscountsDiscountTypePercentage

// This is an alias to an internal value.
const CreditNoteModelLineItemsDiscountsDiscountTypeAmount = shared.CreditNoteModelLineItemsDiscountsDiscountTypeAmount

// This is an alias to an internal type.
type CreditNoteModelReason = shared.CreditNoteModelReason

// This is an alias to an internal value.
const CreditNoteModelReasonDuplicate = shared.CreditNoteModelReasonDuplicate

// This is an alias to an internal value.
const CreditNoteModelReasonFraudulent = shared.CreditNoteModelReasonFraudulent

// This is an alias to an internal value.
const CreditNoteModelReasonOrderChange = shared.CreditNoteModelReasonOrderChange

// This is an alias to an internal value.
const CreditNoteModelReasonProductUnsatisfactory = shared.CreditNoteModelReasonProductUnsatisfactory

// This is an alias to an internal type.
type CreditNoteModelType = shared.CreditNoteModelType

// This is an alias to an internal value.
const CreditNoteModelTypeRefund = shared.CreditNoteModelTypeRefund

// This is an alias to an internal value.
const CreditNoteModelTypeAdjustment = shared.CreditNoteModelTypeAdjustment

// This is an alias to an internal type.
type CreditNoteSummaryModel = shared.CreditNoteSummaryModel

// This is an alias to an internal type.
type CustomRatingFunctionConfigModel = shared.CustomRatingFunctionConfigModel

// This is an alias to an internal type.
type CustomRatingFunctionConfigModelParam = shared.CustomRatingFunctionConfigModelParam

// This is an alias to an internal type.
type CustomerBalanceTransactionModel = shared.CustomerBalanceTransactionModel

// This is an alias to an internal type.
type CustomerBalanceTransactionModelAction = shared.CustomerBalanceTransactionModelAction

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionAppliedToInvoice = shared.CustomerBalanceTransactionModelActionAppliedToInvoice

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionManualAdjustment = shared.CustomerBalanceTransactionModelActionManualAdjustment

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionProratedRefund = shared.CustomerBalanceTransactionModelActionProratedRefund

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionRevertProratedRefund = shared.CustomerBalanceTransactionModelActionRevertProratedRefund

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionReturnFromVoiding = shared.CustomerBalanceTransactionModelActionReturnFromVoiding

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionCreditNoteApplied = shared.CustomerBalanceTransactionModelActionCreditNoteApplied

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionCreditNoteVoided = shared.CustomerBalanceTransactionModelActionCreditNoteVoided

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionOverpaymentRefund = shared.CustomerBalanceTransactionModelActionOverpaymentRefund

// This is an alias to an internal value.
const CustomerBalanceTransactionModelActionExternalPayment = shared.CustomerBalanceTransactionModelActionExternalPayment

// This is an alias to an internal type.
type CustomerBalanceTransactionModelCreditNote = shared.CustomerBalanceTransactionModelCreditNote

// This is an alias to an internal type.
type CustomerBalanceTransactionModelInvoice = shared.CustomerBalanceTransactionModelInvoice

// This is an alias to an internal type.
type CustomerBalanceTransactionModelType = shared.CustomerBalanceTransactionModelType

// This is an alias to an internal value.
const CustomerBalanceTransactionModelTypeIncrement = shared.CustomerBalanceTransactionModelTypeIncrement

// This is an alias to an internal value.
const CustomerBalanceTransactionModelTypeDecrement = shared.CustomerBalanceTransactionModelTypeDecrement

// This is an alias to an internal type.
type CustomerCostsModel = shared.CustomerCostsModel

// This is an alias to an internal type.
type CustomerCreditBalancesModel = shared.CustomerCreditBalancesModel

// This is an alias to an internal type.
type CustomerCreditBalancesModelData = shared.CustomerCreditBalancesModelData

// This is an alias to an internal type.
type CustomerCreditBalancesModelDataStatus = shared.CustomerCreditBalancesModelDataStatus

// This is an alias to an internal value.
const CustomerCreditBalancesModelDataStatusActive = shared.CustomerCreditBalancesModelDataStatusActive

// This is an alias to an internal value.
const CustomerCreditBalancesModelDataStatusPendingPayment = shared.CustomerCreditBalancesModelDataStatusPendingPayment

// This is an alias to an internal type.
type CustomerHierarchyConfigModelParam = shared.CustomerHierarchyConfigModelParam

// This is an alias to an internal type.
type CustomerMinifiedModel = shared.CustomerMinifiedModel

// A customer is a buyer of your products, and the other party to the billing
// relationship.
//
// In Orb, customers are assigned system generated identifiers automatically, but
// it's often desirable to have these match existing identifiers in your system. To
// avoid having to denormalize Orb ID information, you can pass in an
// `external_customer_id` with your own identifier. See
// [Customer ID Aliases](/events-and-metrics/customer-aliases) for further
// information about how these aliases work in Orb.
//
// In addition to having an identifier in your system, a customer may exist in a
// payment provider solution like Stripe. Use the `payment_provider_id` and the
// `payment_provider` enum field to express this mapping.
//
// A customer also has a timezone (from the standard
// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
// your account's timezone. See [Timezone localization](/essentials/timezones) for
// information on what this timezone parameter influences within Orb.
//
// This is an alias to an internal type.
type CustomerModel = shared.CustomerModel

// The hierarchical relationships for this customer.
//
// This is an alias to an internal type.
type CustomerModelHierarchy = shared.CustomerModelHierarchy

// This is used for creating charges or invoices in an external system via Orb.
// When not in test mode, the connection must first be configured in the Orb
// webapp.
//
// This is an alias to an internal type.
type CustomerModelPaymentProvider = shared.CustomerModelPaymentProvider

// This is an alias to an internal value.
const CustomerModelPaymentProviderQuickbooks = shared.CustomerModelPaymentProviderQuickbooks

// This is an alias to an internal value.
const CustomerModelPaymentProviderBillCom = shared.CustomerModelPaymentProviderBillCom

// This is an alias to an internal value.
const CustomerModelPaymentProviderStripeCharge = shared.CustomerModelPaymentProviderStripeCharge

// This is an alias to an internal value.
const CustomerModelPaymentProviderStripeInvoice = shared.CustomerModelPaymentProviderStripeInvoice

// This is an alias to an internal value.
const CustomerModelPaymentProviderNetsuite = shared.CustomerModelPaymentProviderNetsuite

// This is an alias to an internal type.
type CustomerModelAccountingSyncConfiguration = shared.CustomerModelAccountingSyncConfiguration

// This is an alias to an internal type.
type CustomerModelAccountingSyncConfigurationAccountingProvider = shared.CustomerModelAccountingSyncConfigurationAccountingProvider

// This is an alias to an internal type.
type CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType = shared.CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType

// This is an alias to an internal value.
const CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks = shared.CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks

// This is an alias to an internal value.
const CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite = shared.CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite

// This is an alias to an internal type.
type CustomerModelReportingConfiguration = shared.CustomerModelReportingConfiguration

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
//
// This is an alias to an internal type.
type CustomerTaxIDModel = shared.CustomerTaxIDModel

// This is an alias to an internal type.
type CustomerTaxIDModelCountry = shared.CustomerTaxIDModelCountry

// This is an alias to an internal value.
const CustomerTaxIDModelCountryAd = shared.CustomerTaxIDModelCountryAd

// This is an alias to an internal value.
const CustomerTaxIDModelCountryAe = shared.CustomerTaxIDModelCountryAe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryAr = shared.CustomerTaxIDModelCountryAr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryAt = shared.CustomerTaxIDModelCountryAt

// This is an alias to an internal value.
const CustomerTaxIDModelCountryAu = shared.CustomerTaxIDModelCountryAu

// This is an alias to an internal value.
const CustomerTaxIDModelCountryBe = shared.CustomerTaxIDModelCountryBe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryBg = shared.CustomerTaxIDModelCountryBg

// This is an alias to an internal value.
const CustomerTaxIDModelCountryBh = shared.CustomerTaxIDModelCountryBh

// This is an alias to an internal value.
const CustomerTaxIDModelCountryBo = shared.CustomerTaxIDModelCountryBo

// This is an alias to an internal value.
const CustomerTaxIDModelCountryBr = shared.CustomerTaxIDModelCountryBr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCa = shared.CustomerTaxIDModelCountryCa

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCh = shared.CustomerTaxIDModelCountryCh

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCl = shared.CustomerTaxIDModelCountryCl

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCn = shared.CustomerTaxIDModelCountryCn

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCo = shared.CustomerTaxIDModelCountryCo

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCr = shared.CustomerTaxIDModelCountryCr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCy = shared.CustomerTaxIDModelCountryCy

// This is an alias to an internal value.
const CustomerTaxIDModelCountryCz = shared.CustomerTaxIDModelCountryCz

// This is an alias to an internal value.
const CustomerTaxIDModelCountryDe = shared.CustomerTaxIDModelCountryDe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryDk = shared.CustomerTaxIDModelCountryDk

// This is an alias to an internal value.
const CustomerTaxIDModelCountryEe = shared.CustomerTaxIDModelCountryEe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryDo = shared.CustomerTaxIDModelCountryDo

// This is an alias to an internal value.
const CustomerTaxIDModelCountryEc = shared.CustomerTaxIDModelCountryEc

// This is an alias to an internal value.
const CustomerTaxIDModelCountryEg = shared.CustomerTaxIDModelCountryEg

// This is an alias to an internal value.
const CustomerTaxIDModelCountryEs = shared.CustomerTaxIDModelCountryEs

// This is an alias to an internal value.
const CustomerTaxIDModelCountryEu = shared.CustomerTaxIDModelCountryEu

// This is an alias to an internal value.
const CustomerTaxIDModelCountryFi = shared.CustomerTaxIDModelCountryFi

// This is an alias to an internal value.
const CustomerTaxIDModelCountryFr = shared.CustomerTaxIDModelCountryFr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryGB = shared.CustomerTaxIDModelCountryGB

// This is an alias to an internal value.
const CustomerTaxIDModelCountryGe = shared.CustomerTaxIDModelCountryGe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryGr = shared.CustomerTaxIDModelCountryGr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryHk = shared.CustomerTaxIDModelCountryHk

// This is an alias to an internal value.
const CustomerTaxIDModelCountryHr = shared.CustomerTaxIDModelCountryHr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryHu = shared.CustomerTaxIDModelCountryHu

// This is an alias to an internal value.
const CustomerTaxIDModelCountryID = shared.CustomerTaxIDModelCountryID

// This is an alias to an internal value.
const CustomerTaxIDModelCountryIe = shared.CustomerTaxIDModelCountryIe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryIl = shared.CustomerTaxIDModelCountryIl

// This is an alias to an internal value.
const CustomerTaxIDModelCountryIn = shared.CustomerTaxIDModelCountryIn

// This is an alias to an internal value.
const CustomerTaxIDModelCountryIs = shared.CustomerTaxIDModelCountryIs

// This is an alias to an internal value.
const CustomerTaxIDModelCountryIt = shared.CustomerTaxIDModelCountryIt

// This is an alias to an internal value.
const CustomerTaxIDModelCountryJp = shared.CustomerTaxIDModelCountryJp

// This is an alias to an internal value.
const CustomerTaxIDModelCountryKe = shared.CustomerTaxIDModelCountryKe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryKr = shared.CustomerTaxIDModelCountryKr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryKz = shared.CustomerTaxIDModelCountryKz

// This is an alias to an internal value.
const CustomerTaxIDModelCountryLi = shared.CustomerTaxIDModelCountryLi

// This is an alias to an internal value.
const CustomerTaxIDModelCountryLt = shared.CustomerTaxIDModelCountryLt

// This is an alias to an internal value.
const CustomerTaxIDModelCountryLu = shared.CustomerTaxIDModelCountryLu

// This is an alias to an internal value.
const CustomerTaxIDModelCountryLv = shared.CustomerTaxIDModelCountryLv

// This is an alias to an internal value.
const CustomerTaxIDModelCountryMt = shared.CustomerTaxIDModelCountryMt

// This is an alias to an internal value.
const CustomerTaxIDModelCountryMx = shared.CustomerTaxIDModelCountryMx

// This is an alias to an internal value.
const CustomerTaxIDModelCountryMy = shared.CustomerTaxIDModelCountryMy

// This is an alias to an internal value.
const CustomerTaxIDModelCountryNg = shared.CustomerTaxIDModelCountryNg

// This is an alias to an internal value.
const CustomerTaxIDModelCountryNl = shared.CustomerTaxIDModelCountryNl

// This is an alias to an internal value.
const CustomerTaxIDModelCountryNo = shared.CustomerTaxIDModelCountryNo

// This is an alias to an internal value.
const CustomerTaxIDModelCountryNz = shared.CustomerTaxIDModelCountryNz

// This is an alias to an internal value.
const CustomerTaxIDModelCountryOm = shared.CustomerTaxIDModelCountryOm

// This is an alias to an internal value.
const CustomerTaxIDModelCountryPe = shared.CustomerTaxIDModelCountryPe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryPh = shared.CustomerTaxIDModelCountryPh

// This is an alias to an internal value.
const CustomerTaxIDModelCountryPl = shared.CustomerTaxIDModelCountryPl

// This is an alias to an internal value.
const CustomerTaxIDModelCountryPt = shared.CustomerTaxIDModelCountryPt

// This is an alias to an internal value.
const CustomerTaxIDModelCountryRo = shared.CustomerTaxIDModelCountryRo

// This is an alias to an internal value.
const CustomerTaxIDModelCountryRs = shared.CustomerTaxIDModelCountryRs

// This is an alias to an internal value.
const CustomerTaxIDModelCountryRu = shared.CustomerTaxIDModelCountryRu

// This is an alias to an internal value.
const CustomerTaxIDModelCountrySa = shared.CustomerTaxIDModelCountrySa

// This is an alias to an internal value.
const CustomerTaxIDModelCountrySe = shared.CustomerTaxIDModelCountrySe

// This is an alias to an internal value.
const CustomerTaxIDModelCountrySg = shared.CustomerTaxIDModelCountrySg

// This is an alias to an internal value.
const CustomerTaxIDModelCountrySi = shared.CustomerTaxIDModelCountrySi

// This is an alias to an internal value.
const CustomerTaxIDModelCountrySk = shared.CustomerTaxIDModelCountrySk

// This is an alias to an internal value.
const CustomerTaxIDModelCountrySv = shared.CustomerTaxIDModelCountrySv

// This is an alias to an internal value.
const CustomerTaxIDModelCountryTh = shared.CustomerTaxIDModelCountryTh

// This is an alias to an internal value.
const CustomerTaxIDModelCountryTr = shared.CustomerTaxIDModelCountryTr

// This is an alias to an internal value.
const CustomerTaxIDModelCountryTw = shared.CustomerTaxIDModelCountryTw

// This is an alias to an internal value.
const CustomerTaxIDModelCountryUa = shared.CustomerTaxIDModelCountryUa

// This is an alias to an internal value.
const CustomerTaxIDModelCountryUs = shared.CustomerTaxIDModelCountryUs

// This is an alias to an internal value.
const CustomerTaxIDModelCountryUy = shared.CustomerTaxIDModelCountryUy

// This is an alias to an internal value.
const CustomerTaxIDModelCountryVe = shared.CustomerTaxIDModelCountryVe

// This is an alias to an internal value.
const CustomerTaxIDModelCountryVn = shared.CustomerTaxIDModelCountryVn

// This is an alias to an internal value.
const CustomerTaxIDModelCountryZa = shared.CustomerTaxIDModelCountryZa

// This is an alias to an internal type.
type CustomerTaxIDModelType = shared.CustomerTaxIDModelType

// This is an alias to an internal value.
const CustomerTaxIDModelTypeAdNrt = shared.CustomerTaxIDModelTypeAdNrt

// This is an alias to an internal value.
const CustomerTaxIDModelTypeAeTrn = shared.CustomerTaxIDModelTypeAeTrn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeArCuit = shared.CustomerTaxIDModelTypeArCuit

// This is an alias to an internal value.
const CustomerTaxIDModelTypeEuVat = shared.CustomerTaxIDModelTypeEuVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeAuAbn = shared.CustomerTaxIDModelTypeAuAbn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeAuArn = shared.CustomerTaxIDModelTypeAuArn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeBgUic = shared.CustomerTaxIDModelTypeBgUic

// This is an alias to an internal value.
const CustomerTaxIDModelTypeBhVat = shared.CustomerTaxIDModelTypeBhVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeBoTin = shared.CustomerTaxIDModelTypeBoTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeBrCnpj = shared.CustomerTaxIDModelTypeBrCnpj

// This is an alias to an internal value.
const CustomerTaxIDModelTypeBrCpf = shared.CustomerTaxIDModelTypeBrCpf

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCaBn = shared.CustomerTaxIDModelTypeCaBn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCaGstHst = shared.CustomerTaxIDModelTypeCaGstHst

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCaPstBc = shared.CustomerTaxIDModelTypeCaPstBc

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCaPstMB = shared.CustomerTaxIDModelTypeCaPstMB

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCaPstSk = shared.CustomerTaxIDModelTypeCaPstSk

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCaQst = shared.CustomerTaxIDModelTypeCaQst

// This is an alias to an internal value.
const CustomerTaxIDModelTypeChVat = shared.CustomerTaxIDModelTypeChVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeClTin = shared.CustomerTaxIDModelTypeClTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCnTin = shared.CustomerTaxIDModelTypeCnTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCoNit = shared.CustomerTaxIDModelTypeCoNit

// This is an alias to an internal value.
const CustomerTaxIDModelTypeCrTin = shared.CustomerTaxIDModelTypeCrTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeDoRcn = shared.CustomerTaxIDModelTypeDoRcn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeEcRuc = shared.CustomerTaxIDModelTypeEcRuc

// This is an alias to an internal value.
const CustomerTaxIDModelTypeEgTin = shared.CustomerTaxIDModelTypeEgTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeEsCif = shared.CustomerTaxIDModelTypeEsCif

// This is an alias to an internal value.
const CustomerTaxIDModelTypeEuOssVat = shared.CustomerTaxIDModelTypeEuOssVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeGBVat = shared.CustomerTaxIDModelTypeGBVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeGeVat = shared.CustomerTaxIDModelTypeGeVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeHkBr = shared.CustomerTaxIDModelTypeHkBr

// This is an alias to an internal value.
const CustomerTaxIDModelTypeHuTin = shared.CustomerTaxIDModelTypeHuTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeIDNpwp = shared.CustomerTaxIDModelTypeIDNpwp

// This is an alias to an internal value.
const CustomerTaxIDModelTypeIlVat = shared.CustomerTaxIDModelTypeIlVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeInGst = shared.CustomerTaxIDModelTypeInGst

// This is an alias to an internal value.
const CustomerTaxIDModelTypeIsVat = shared.CustomerTaxIDModelTypeIsVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeJpCn = shared.CustomerTaxIDModelTypeJpCn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeJpRn = shared.CustomerTaxIDModelTypeJpRn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeJpTrn = shared.CustomerTaxIDModelTypeJpTrn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeKePin = shared.CustomerTaxIDModelTypeKePin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeKrBrn = shared.CustomerTaxIDModelTypeKrBrn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeKzBin = shared.CustomerTaxIDModelTypeKzBin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeLiUid = shared.CustomerTaxIDModelTypeLiUid

// This is an alias to an internal value.
const CustomerTaxIDModelTypeMxRfc = shared.CustomerTaxIDModelTypeMxRfc

// This is an alias to an internal value.
const CustomerTaxIDModelTypeMyFrp = shared.CustomerTaxIDModelTypeMyFrp

// This is an alias to an internal value.
const CustomerTaxIDModelTypeMyItn = shared.CustomerTaxIDModelTypeMyItn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeMySst = shared.CustomerTaxIDModelTypeMySst

// This is an alias to an internal value.
const CustomerTaxIDModelTypeNgTin = shared.CustomerTaxIDModelTypeNgTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeNoVat = shared.CustomerTaxIDModelTypeNoVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeNoVoec = shared.CustomerTaxIDModelTypeNoVoec

// This is an alias to an internal value.
const CustomerTaxIDModelTypeNzGst = shared.CustomerTaxIDModelTypeNzGst

// This is an alias to an internal value.
const CustomerTaxIDModelTypeOmVat = shared.CustomerTaxIDModelTypeOmVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypePeRuc = shared.CustomerTaxIDModelTypePeRuc

// This is an alias to an internal value.
const CustomerTaxIDModelTypePhTin = shared.CustomerTaxIDModelTypePhTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeRoTin = shared.CustomerTaxIDModelTypeRoTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeRsPib = shared.CustomerTaxIDModelTypeRsPib

// This is an alias to an internal value.
const CustomerTaxIDModelTypeRuInn = shared.CustomerTaxIDModelTypeRuInn

// This is an alias to an internal value.
const CustomerTaxIDModelTypeRuKpp = shared.CustomerTaxIDModelTypeRuKpp

// This is an alias to an internal value.
const CustomerTaxIDModelTypeSaVat = shared.CustomerTaxIDModelTypeSaVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeSgGst = shared.CustomerTaxIDModelTypeSgGst

// This is an alias to an internal value.
const CustomerTaxIDModelTypeSgUen = shared.CustomerTaxIDModelTypeSgUen

// This is an alias to an internal value.
const CustomerTaxIDModelTypeSiTin = shared.CustomerTaxIDModelTypeSiTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeSvNit = shared.CustomerTaxIDModelTypeSvNit

// This is an alias to an internal value.
const CustomerTaxIDModelTypeThVat = shared.CustomerTaxIDModelTypeThVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeTrTin = shared.CustomerTaxIDModelTypeTrTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeTwVat = shared.CustomerTaxIDModelTypeTwVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeUaVat = shared.CustomerTaxIDModelTypeUaVat

// This is an alias to an internal value.
const CustomerTaxIDModelTypeUsEin = shared.CustomerTaxIDModelTypeUsEin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeUyRuc = shared.CustomerTaxIDModelTypeUyRuc

// This is an alias to an internal value.
const CustomerTaxIDModelTypeVeRif = shared.CustomerTaxIDModelTypeVeRif

// This is an alias to an internal value.
const CustomerTaxIDModelTypeVnTin = shared.CustomerTaxIDModelTypeVnTin

// This is an alias to an internal value.
const CustomerTaxIDModelTypeZaVat = shared.CustomerTaxIDModelTypeZaVat

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
//
// This is an alias to an internal type.
type CustomerTaxIDModelParam = shared.CustomerTaxIDModelParam

// This is an alias to an internal type.
type DimensionalPriceConfigurationModel = shared.DimensionalPriceConfigurationModel

// A dimensional price group is used to partition the result of a billable metric
// by a set of dimensions. Prices in a price group must specify the parition used
// to derive their usage.
//
// This is an alias to an internal type.
type DimensionalPriceGroupModel = shared.DimensionalPriceGroupModel

// This is an alias to an internal type.
type Discount = shared.Discount

// This is an alias to an internal type.
type DiscountUsageDiscount = shared.DiscountUsageDiscount

// This is an alias to an internal type.
type DiscountUsageDiscountDiscountType = shared.DiscountUsageDiscountDiscountType

// This is an alias to an internal value.
const DiscountUsageDiscountDiscountTypeUsage = shared.DiscountUsageDiscountDiscountTypeUsage

// This is an alias to an internal type.
type DiscountDiscountType = shared.DiscountDiscountType

// This is an alias to an internal value.
const DiscountDiscountTypePercentage = shared.DiscountDiscountTypePercentage

// This is an alias to an internal value.
const DiscountDiscountTypeTrial = shared.DiscountDiscountTypeTrial

// This is an alias to an internal value.
const DiscountDiscountTypeUsage = shared.DiscountDiscountTypeUsage

// This is an alias to an internal value.
const DiscountDiscountTypeAmount = shared.DiscountDiscountTypeAmount

// This is an alias to an internal type.
type DiscountUnionParam = shared.DiscountUnionParam

// This is an alias to an internal type.
type DiscountUsageDiscountParam = shared.DiscountUsageDiscountParam

// This is an alias to an internal type.
type DiscountOverrideModelParam = shared.DiscountOverrideModelParam

// This is an alias to an internal type.
type DiscountOverrideModelDiscountType = shared.DiscountOverrideModelDiscountType

// This is an alias to an internal value.
const DiscountOverrideModelDiscountTypePercentage = shared.DiscountOverrideModelDiscountTypePercentage

// This is an alias to an internal value.
const DiscountOverrideModelDiscountTypeUsage = shared.DiscountOverrideModelDiscountTypeUsage

// This is an alias to an internal value.
const DiscountOverrideModelDiscountTypeAmount = shared.DiscountOverrideModelDiscountTypeAmount

// This is an alias to an internal type.
type FixedFeeQuantityScheduleEntryModel = shared.FixedFeeQuantityScheduleEntryModel

// This is an alias to an internal type.
type InvoiceLevelDiscount = shared.InvoiceLevelDiscount

// This is an alias to an internal type.
type InvoiceLevelDiscountDiscountType = shared.InvoiceLevelDiscountDiscountType

// This is an alias to an internal value.
const InvoiceLevelDiscountDiscountTypePercentage = shared.InvoiceLevelDiscountDiscountTypePercentage

// This is an alias to an internal value.
const InvoiceLevelDiscountDiscountTypeAmount = shared.InvoiceLevelDiscountDiscountTypeAmount

// This is an alias to an internal value.
const InvoiceLevelDiscountDiscountTypeTrial = shared.InvoiceLevelDiscountDiscountTypeTrial

// This is an alias to an internal type.
type InvoiceLineItemModel = shared.InvoiceLineItemModel

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustment = shared.InvoiceLineItemModelAdjustment

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment = shared.InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType = shared.InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount = shared.InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment = shared.InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType = shared.InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount = shared.InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment = shared.InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType = shared.InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount = shared.InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment = shared.InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentType = shared.InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum = shared.InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment = shared.InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentType = shared.InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum = shared.InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type InvoiceLineItemModelAdjustmentsAdjustmentType = shared.InvoiceLineItemModelAdjustmentsAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsAdjustmentTypeUsageDiscount = shared.InvoiceLineItemModelAdjustmentsAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsAdjustmentTypeAmountDiscount = shared.InvoiceLineItemModelAdjustmentsAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsAdjustmentTypePercentageDiscount = shared.InvoiceLineItemModelAdjustmentsAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsAdjustmentTypeMinimum = shared.InvoiceLineItemModelAdjustmentsAdjustmentTypeMinimum

// This is an alias to an internal value.
const InvoiceLineItemModelAdjustmentsAdjustmentTypeMaximum = shared.InvoiceLineItemModelAdjustmentsAdjustmentTypeMaximum

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItem = shared.InvoiceLineItemModelSubLineItem

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsMatrixSubLineItem = shared.InvoiceLineItemModelSubLineItemsMatrixSubLineItem

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig = shared.InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsMatrixSubLineItemType = shared.InvoiceLineItemModelSubLineItemsMatrixSubLineItemType

// This is an alias to an internal value.
const InvoiceLineItemModelSubLineItemsMatrixSubLineItemTypeMatrix = shared.InvoiceLineItemModelSubLineItemsMatrixSubLineItemTypeMatrix

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsTierSubLineItem = shared.InvoiceLineItemModelSubLineItemsTierSubLineItem

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig = shared.InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsTierSubLineItemType = shared.InvoiceLineItemModelSubLineItemsTierSubLineItemType

// This is an alias to an internal value.
const InvoiceLineItemModelSubLineItemsTierSubLineItemTypeTier = shared.InvoiceLineItemModelSubLineItemsTierSubLineItemTypeTier

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsOtherSubLineItem = shared.InvoiceLineItemModelSubLineItemsOtherSubLineItem

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsOtherSubLineItemType = shared.InvoiceLineItemModelSubLineItemsOtherSubLineItemType

// This is an alias to an internal value.
const InvoiceLineItemModelSubLineItemsOtherSubLineItemTypeNull = shared.InvoiceLineItemModelSubLineItemsOtherSubLineItemTypeNull

// This is an alias to an internal type.
type InvoiceLineItemModelSubLineItemsType = shared.InvoiceLineItemModelSubLineItemsType

// This is an alias to an internal value.
const InvoiceLineItemModelSubLineItemsTypeMatrix = shared.InvoiceLineItemModelSubLineItemsTypeMatrix

// This is an alias to an internal value.
const InvoiceLineItemModelSubLineItemsTypeTier = shared.InvoiceLineItemModelSubLineItemsTypeTier

// This is an alias to an internal value.
const InvoiceLineItemModelSubLineItemsTypeNull = shared.InvoiceLineItemModelSubLineItemsTypeNull

// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
// representing the request for payment for a single subscription. This includes a
// set of line items, which correspond to prices in the subscription's plan and can
// represent fixed recurring fees or usage-based fees. They are generated at the
// end of a billing period, or as the result of an action, such as a cancellation.
//
// This is an alias to an internal type.
type InvoiceModel = shared.InvoiceModel

// This is an alias to an internal type.
type InvoiceModelInvoiceSource = shared.InvoiceModelInvoiceSource

// This is an alias to an internal value.
const InvoiceModelInvoiceSourceSubscription = shared.InvoiceModelInvoiceSourceSubscription

// This is an alias to an internal value.
const InvoiceModelInvoiceSourcePartial = shared.InvoiceModelInvoiceSourcePartial

// This is an alias to an internal value.
const InvoiceModelInvoiceSourceOneOff = shared.InvoiceModelInvoiceSourceOneOff

// This is an alias to an internal type.
type InvoiceModelStatus = shared.InvoiceModelStatus

// This is an alias to an internal value.
const InvoiceModelStatusIssued = shared.InvoiceModelStatusIssued

// This is an alias to an internal value.
const InvoiceModelStatusPaid = shared.InvoiceModelStatusPaid

// This is an alias to an internal value.
const InvoiceModelStatusSynced = shared.InvoiceModelStatusSynced

// This is an alias to an internal value.
const InvoiceModelStatusVoid = shared.InvoiceModelStatusVoid

// This is an alias to an internal value.
const InvoiceModelStatusDraft = shared.InvoiceModelStatusDraft

// This is an alias to an internal type.
type ItemExternalConnectionModel = shared.ItemExternalConnectionModel

// This is an alias to an internal type.
type ItemExternalConnectionModelExternalConnectionName = shared.ItemExternalConnectionModelExternalConnectionName

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameStripe = shared.ItemExternalConnectionModelExternalConnectionNameStripe

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameQuickbooks = shared.ItemExternalConnectionModelExternalConnectionNameQuickbooks

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameBillCom = shared.ItemExternalConnectionModelExternalConnectionNameBillCom

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameNetsuite = shared.ItemExternalConnectionModelExternalConnectionNameNetsuite

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameTaxjar = shared.ItemExternalConnectionModelExternalConnectionNameTaxjar

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameAvalara = shared.ItemExternalConnectionModelExternalConnectionNameAvalara

// This is an alias to an internal value.
const ItemExternalConnectionModelExternalConnectionNameAnrok = shared.ItemExternalConnectionModelExternalConnectionNameAnrok

// This is an alias to an internal type.
type ItemExternalConnectionModelParam = shared.ItemExternalConnectionModelParam

// The Item resource represents a sellable product or good. Items are associated
// with all line items, billable metrics, and prices and are used for defining
// external sync behavior for invoices and tax calculation purposes.
//
// This is an alias to an internal type.
type ItemModel = shared.ItemModel

// This is an alias to an internal type.
type ItemSlimModel = shared.ItemSlimModel

// This is an alias to an internal type.
type MatrixConfigModel = shared.MatrixConfigModel

// This is an alias to an internal type.
type MatrixConfigModelParam = shared.MatrixConfigModelParam

// This is an alias to an internal type.
type MatrixValueModel = shared.MatrixValueModel

// This is an alias to an internal type.
type MatrixValueModelParam = shared.MatrixValueModelParam

// This is an alias to an internal type.
type MatrixWithAllocationConfigModel = shared.MatrixWithAllocationConfigModel

// This is an alias to an internal type.
type MatrixWithAllocationConfigModelParam = shared.MatrixWithAllocationConfigModelParam

// This is an alias to an internal type.
type MaximumIntervalModel = shared.MaximumIntervalModel

// This is an alias to an internal type.
type MaximumModel = shared.MaximumModel

// This is an alias to an internal type.
type MinimumIntervalModel = shared.MinimumIntervalModel

// This is an alias to an internal type.
type MinimumModel = shared.MinimumModel

// This is an alias to an internal type.
type MutatedSubscriptionModel = shared.MutatedSubscriptionModel

// This is an alias to an internal type.
type MutatedSubscriptionModelDiscountInterval = shared.MutatedSubscriptionModelDiscountInterval

// This is an alias to an internal type.
type MutatedSubscriptionModelDiscountIntervalsDiscountType = shared.MutatedSubscriptionModelDiscountIntervalsDiscountType

// This is an alias to an internal value.
const MutatedSubscriptionModelDiscountIntervalsDiscountTypeAmount = shared.MutatedSubscriptionModelDiscountIntervalsDiscountTypeAmount

// This is an alias to an internal value.
const MutatedSubscriptionModelDiscountIntervalsDiscountTypePercentage = shared.MutatedSubscriptionModelDiscountIntervalsDiscountTypePercentage

// This is an alias to an internal value.
const MutatedSubscriptionModelDiscountIntervalsDiscountTypeUsage = shared.MutatedSubscriptionModelDiscountIntervalsDiscountTypeUsage

// This is an alias to an internal type.
type MutatedSubscriptionModelStatus = shared.MutatedSubscriptionModelStatus

// This is an alias to an internal value.
const MutatedSubscriptionModelStatusActive = shared.MutatedSubscriptionModelStatusActive

// This is an alias to an internal value.
const MutatedSubscriptionModelStatusEnded = shared.MutatedSubscriptionModelStatusEnded

// This is an alias to an internal value.
const MutatedSubscriptionModelStatusUpcoming = shared.MutatedSubscriptionModelStatusUpcoming

// This is an alias to an internal type.
type NewAccountingSyncConfigurationModelParam = shared.NewAccountingSyncConfigurationModelParam

// This is an alias to an internal type.
type NewAccountingSyncConfigurationModelAccountingProviderParam = shared.NewAccountingSyncConfigurationModelAccountingProviderParam

// This is an alias to an internal type.
type NewAdjustmentModelUnionParam = shared.NewAdjustmentModelUnionParam

// This is an alias to an internal type.
type NewAdjustmentModelNewPercentageDiscountParam = shared.NewAdjustmentModelNewPercentageDiscountParam

// This is an alias to an internal type.
type NewAdjustmentModelNewPercentageDiscountAdjustmentType = shared.NewAdjustmentModelNewPercentageDiscountAdjustmentType

// This is an alias to an internal value.
const NewAdjustmentModelNewPercentageDiscountAdjustmentTypePercentageDiscount = shared.NewAdjustmentModelNewPercentageDiscountAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type NewAdjustmentModelNewUsageDiscountParam = shared.NewAdjustmentModelNewUsageDiscountParam

// This is an alias to an internal type.
type NewAdjustmentModelNewUsageDiscountAdjustmentType = shared.NewAdjustmentModelNewUsageDiscountAdjustmentType

// This is an alias to an internal value.
const NewAdjustmentModelNewUsageDiscountAdjustmentTypeUsageDiscount = shared.NewAdjustmentModelNewUsageDiscountAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type NewAdjustmentModelNewAmountDiscountParam = shared.NewAdjustmentModelNewAmountDiscountParam

// This is an alias to an internal type.
type NewAdjustmentModelNewAmountDiscountAdjustmentType = shared.NewAdjustmentModelNewAmountDiscountAdjustmentType

// This is an alias to an internal value.
const NewAdjustmentModelNewAmountDiscountAdjustmentTypeAmountDiscount = shared.NewAdjustmentModelNewAmountDiscountAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type NewAdjustmentModelNewMinimumParam = shared.NewAdjustmentModelNewMinimumParam

// This is an alias to an internal type.
type NewAdjustmentModelNewMinimumAdjustmentType = shared.NewAdjustmentModelNewMinimumAdjustmentType

// This is an alias to an internal value.
const NewAdjustmentModelNewMinimumAdjustmentTypeMinimum = shared.NewAdjustmentModelNewMinimumAdjustmentTypeMinimum

// This is an alias to an internal type.
type NewAdjustmentModelNewMaximumParam = shared.NewAdjustmentModelNewMaximumParam

// This is an alias to an internal type.
type NewAdjustmentModelNewMaximumAdjustmentType = shared.NewAdjustmentModelNewMaximumAdjustmentType

// This is an alias to an internal value.
const NewAdjustmentModelNewMaximumAdjustmentTypeMaximum = shared.NewAdjustmentModelNewMaximumAdjustmentTypeMaximum

// This is an alias to an internal type.
type NewAdjustmentModelAdjustmentType = shared.NewAdjustmentModelAdjustmentType

// This is an alias to an internal value.
const NewAdjustmentModelAdjustmentTypePercentageDiscount = shared.NewAdjustmentModelAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const NewAdjustmentModelAdjustmentTypeUsageDiscount = shared.NewAdjustmentModelAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const NewAdjustmentModelAdjustmentTypeAmountDiscount = shared.NewAdjustmentModelAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const NewAdjustmentModelAdjustmentTypeMinimum = shared.NewAdjustmentModelAdjustmentTypeMinimum

// This is an alias to an internal value.
const NewAdjustmentModelAdjustmentTypeMaximum = shared.NewAdjustmentModelAdjustmentTypeMaximum

// This is an alias to an internal type.
type NewAllocationPriceModelParam = shared.NewAllocationPriceModelParam

// The cadence at which to allocate the amount to the customer.
//
// This is an alias to an internal type.
type NewAllocationPriceModelCadence = shared.NewAllocationPriceModelCadence

// This is an alias to an internal value.
const NewAllocationPriceModelCadenceOneTime = shared.NewAllocationPriceModelCadenceOneTime

// This is an alias to an internal value.
const NewAllocationPriceModelCadenceMonthly = shared.NewAllocationPriceModelCadenceMonthly

// This is an alias to an internal value.
const NewAllocationPriceModelCadenceQuarterly = shared.NewAllocationPriceModelCadenceQuarterly

// This is an alias to an internal value.
const NewAllocationPriceModelCadenceSemiAnnual = shared.NewAllocationPriceModelCadenceSemiAnnual

// This is an alias to an internal value.
const NewAllocationPriceModelCadenceAnnual = shared.NewAllocationPriceModelCadenceAnnual

// This is an alias to an internal value.
const NewAllocationPriceModelCadenceCustom = shared.NewAllocationPriceModelCadenceCustom

// This is an alias to an internal type.
type NewBillingCycleConfigurationModelParam = shared.NewBillingCycleConfigurationModelParam

// The unit of billing period duration.
//
// This is an alias to an internal type.
type NewBillingCycleConfigurationModelDurationUnit = shared.NewBillingCycleConfigurationModelDurationUnit

// This is an alias to an internal value.
const NewBillingCycleConfigurationModelDurationUnitDay = shared.NewBillingCycleConfigurationModelDurationUnitDay

// This is an alias to an internal value.
const NewBillingCycleConfigurationModelDurationUnitMonth = shared.NewBillingCycleConfigurationModelDurationUnitMonth

// This is an alias to an internal type.
type NewFloatingPriceModelUnionParam = shared.NewFloatingPriceModelUnionParam

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitPriceParam = shared.NewFloatingPriceModelNewFloatingUnitPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitPriceCadence = shared.NewFloatingPriceModelNewFloatingUnitPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingUnitPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingUnitPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingUnitPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingUnitPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitPriceModelType = shared.NewFloatingPriceModelNewFloatingUnitPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitPriceModelTypeUnit = shared.NewFloatingPriceModelNewFloatingUnitPriceModelTypeUnit

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingPackagePriceParam = shared.NewFloatingPriceModelNewFloatingPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingPackagePriceCadence = shared.NewFloatingPriceModelNewFloatingPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingPackagePriceModelType = shared.NewFloatingPriceModelNewFloatingPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackagePriceModelTypePackage = shared.NewFloatingPriceModelNewFloatingPackagePriceModelTypePackage

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixPriceParam = shared.NewFloatingPriceModelNewFloatingMatrixPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixPriceCadence = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingMatrixPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixPriceModelType = shared.NewFloatingPriceModelNewFloatingMatrixPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixPriceModelTypeMatrix = shared.NewFloatingPriceModelNewFloatingMatrixPriceModelTypeMatrix

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceParam = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelType = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPriceParam = shared.NewFloatingPriceModelNewFloatingTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPriceCadence = shared.NewFloatingPriceModelNewFloatingTieredPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingTieredPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPriceModelType = shared.NewFloatingPriceModelNewFloatingTieredPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPriceModelTypeTiered = shared.NewFloatingPriceModelNewFloatingTieredPriceModelTypeTiered

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredBpsPriceParam = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredBpsPriceModelType = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredBpsPriceModelTypeTieredBps = shared.NewFloatingPriceModelNewFloatingTieredBpsPriceModelTypeTieredBps

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBpsPriceParam = shared.NewFloatingPriceModelNewFloatingBpsPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBpsPriceCadence = shared.NewFloatingPriceModelNewFloatingBpsPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingBpsPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingBpsPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingBpsPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingBpsPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBpsPriceModelType = shared.NewFloatingPriceModelNewFloatingBpsPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBpsPriceModelTypeBps = shared.NewFloatingPriceModelNewFloatingBpsPriceModelTypeBps

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkBpsPriceParam = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkBpsPriceModelType = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkBpsPriceModelTypeBulkBps = shared.NewFloatingPriceModelNewFloatingBulkBpsPriceModelTypeBulkBps

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkPriceParam = shared.NewFloatingPriceModelNewFloatingBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkPriceCadence = shared.NewFloatingPriceModelNewFloatingBulkPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingBulkPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkPriceModelType = shared.NewFloatingPriceModelNewFloatingBulkPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkPriceModelTypeBulk = shared.NewFloatingPriceModelNewFloatingBulkPriceModelTypeBulk

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceParam = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelType = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPackagePriceParam = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPackagePriceModelType = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackagePriceModelTypeTieredPackage = shared.NewFloatingPriceModelNewFloatingTieredPackagePriceModelTypeTieredPackage

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedTieredPriceParam = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedTieredPriceModelType = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPriceModelTypeGroupedTiered = shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceModelTypeGroupedTiered

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceParam = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelType = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredWithMinimumPriceParam = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelType = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingPackageWithAllocationPriceParam = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelType = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation = shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceParam = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelType = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitWithPercentPriceParam = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelType = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent = shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredWithProrationPriceParam = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelType = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelTypeTieredWithProration = shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelTypeTieredWithProration

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitWithProrationPriceParam = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelType = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelTypeUnitWithProration = shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelTypeUnitWithProration

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedAllocationPriceParam = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelType = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation = shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceParam = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelType = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceParam = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelType = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceParam = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelType = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkWithProrationPriceParam = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelType = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelTypeBulkWithProration = shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelTypeBulkWithProration

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceParam = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelType = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceParam = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelType = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceParam = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelType = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceParam = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceAnnual = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceMonthly = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceQuarterly = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceOneTime = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceCustom = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelType = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPriceModelCadence = shared.NewFloatingPriceModelCadence

// This is an alias to an internal value.
const NewFloatingPriceModelCadenceAnnual = shared.NewFloatingPriceModelCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelCadenceSemiAnnual = shared.NewFloatingPriceModelCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPriceModelCadenceMonthly = shared.NewFloatingPriceModelCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPriceModelCadenceQuarterly = shared.NewFloatingPriceModelCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPriceModelCadenceOneTime = shared.NewFloatingPriceModelCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPriceModelCadenceCustom = shared.NewFloatingPriceModelCadenceCustom

// This is an alias to an internal type.
type NewFloatingPriceModelModelType = shared.NewFloatingPriceModelModelType

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeUnit = shared.NewFloatingPriceModelModelTypeUnit

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypePackage = shared.NewFloatingPriceModelModelTypePackage

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeMatrix = shared.NewFloatingPriceModelModelTypeMatrix

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeMatrixWithAllocation = shared.NewFloatingPriceModelModelTypeMatrixWithAllocation

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeTiered = shared.NewFloatingPriceModelModelTypeTiered

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeTieredBps = shared.NewFloatingPriceModelModelTypeTieredBps

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeBps = shared.NewFloatingPriceModelModelTypeBps

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeBulkBps = shared.NewFloatingPriceModelModelTypeBulkBps

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeBulk = shared.NewFloatingPriceModelModelTypeBulk

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeThresholdTotalAmount = shared.NewFloatingPriceModelModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeTieredPackage = shared.NewFloatingPriceModelModelTypeTieredPackage

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeGroupedTiered = shared.NewFloatingPriceModelModelTypeGroupedTiered

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeMaxGroupTieredPackage = shared.NewFloatingPriceModelModelTypeMaxGroupTieredPackage

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeTieredWithMinimum = shared.NewFloatingPriceModelModelTypeTieredWithMinimum

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypePackageWithAllocation = shared.NewFloatingPriceModelModelTypePackageWithAllocation

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeTieredPackageWithMinimum = shared.NewFloatingPriceModelModelTypeTieredPackageWithMinimum

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeUnitWithPercent = shared.NewFloatingPriceModelModelTypeUnitWithPercent

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeTieredWithProration = shared.NewFloatingPriceModelModelTypeTieredWithProration

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeUnitWithProration = shared.NewFloatingPriceModelModelTypeUnitWithProration

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeGroupedAllocation = shared.NewFloatingPriceModelModelTypeGroupedAllocation

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeGroupedWithProratedMinimum = shared.NewFloatingPriceModelModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeGroupedWithMeteredMinimum = shared.NewFloatingPriceModelModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeMatrixWithDisplayName = shared.NewFloatingPriceModelModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeBulkWithProration = shared.NewFloatingPriceModelModelTypeBulkWithProration

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeGroupedTieredPackage = shared.NewFloatingPriceModelModelTypeGroupedTieredPackage

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeScalableMatrixWithUnitPricing = shared.NewFloatingPriceModelModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeScalableMatrixWithTieredPricing = shared.NewFloatingPriceModelModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal value.
const NewFloatingPriceModelModelTypeCumulativeGroupedBulk = shared.NewFloatingPriceModelModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type NewReportingConfigurationModelParam = shared.NewReportingConfigurationModelParam

// This is an alias to an internal type.
type NewSubscriptionPriceModelUnionParam = shared.NewSubscriptionPriceModelUnionParam

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitPriceModelTypeUnit = shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceModelTypeUnit

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionPackagePriceParam = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionPackagePriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackagePriceModelTypePackage = shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceModelTypePackage

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMatrixPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelTypeMatrix = shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelTypeMatrix

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPriceModelTypeTiered = shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceModelTypeTiered

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelTypeTieredBps = shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelTypeTieredBps

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBpsPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBpsPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBpsPriceModelTypeBps = shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceModelTypeBps

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelTypeBulkBps = shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelTypeBulkBps

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkPriceModelTypeBulk = shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceModelTypeBulk

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceParam = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelTypeTieredPackage = shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelTypeTieredPackage

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation = shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration = shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration = shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation = shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration = shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceParam = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceParam = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceParam = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceParam = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceMonthly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceOneTime = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceCustom = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelType = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewSubscriptionPriceModelCadence = shared.NewSubscriptionPriceModelCadence

// This is an alias to an internal value.
const NewSubscriptionPriceModelCadenceAnnual = shared.NewSubscriptionPriceModelCadenceAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelCadenceSemiAnnual = shared.NewSubscriptionPriceModelCadenceSemiAnnual

// This is an alias to an internal value.
const NewSubscriptionPriceModelCadenceMonthly = shared.NewSubscriptionPriceModelCadenceMonthly

// This is an alias to an internal value.
const NewSubscriptionPriceModelCadenceQuarterly = shared.NewSubscriptionPriceModelCadenceQuarterly

// This is an alias to an internal value.
const NewSubscriptionPriceModelCadenceOneTime = shared.NewSubscriptionPriceModelCadenceOneTime

// This is an alias to an internal value.
const NewSubscriptionPriceModelCadenceCustom = shared.NewSubscriptionPriceModelCadenceCustom

// This is an alias to an internal type.
type NewSubscriptionPriceModelModelType = shared.NewSubscriptionPriceModelModelType

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeUnit = shared.NewSubscriptionPriceModelModelTypeUnit

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypePackage = shared.NewSubscriptionPriceModelModelTypePackage

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeMatrix = shared.NewSubscriptionPriceModelModelTypeMatrix

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeTiered = shared.NewSubscriptionPriceModelModelTypeTiered

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeTieredBps = shared.NewSubscriptionPriceModelModelTypeTieredBps

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeBps = shared.NewSubscriptionPriceModelModelTypeBps

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeBulkBps = shared.NewSubscriptionPriceModelModelTypeBulkBps

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeBulk = shared.NewSubscriptionPriceModelModelTypeBulk

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeThresholdTotalAmount = shared.NewSubscriptionPriceModelModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeTieredPackage = shared.NewSubscriptionPriceModelModelTypeTieredPackage

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeTieredWithMinimum = shared.NewSubscriptionPriceModelModelTypeTieredWithMinimum

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeUnitWithPercent = shared.NewSubscriptionPriceModelModelTypeUnitWithPercent

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypePackageWithAllocation = shared.NewSubscriptionPriceModelModelTypePackageWithAllocation

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeTieredWithProration = shared.NewSubscriptionPriceModelModelTypeTieredWithProration

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeUnitWithProration = shared.NewSubscriptionPriceModelModelTypeUnitWithProration

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeGroupedAllocation = shared.NewSubscriptionPriceModelModelTypeGroupedAllocation

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeGroupedWithProratedMinimum = shared.NewSubscriptionPriceModelModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeBulkWithProration = shared.NewSubscriptionPriceModelModelTypeBulkWithProration

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeScalableMatrixWithUnitPricing = shared.NewSubscriptionPriceModelModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeScalableMatrixWithTieredPricing = shared.NewSubscriptionPriceModelModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeCumulativeGroupedBulk = shared.NewSubscriptionPriceModelModelTypeCumulativeGroupedBulk

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeMaxGroupTieredPackage = shared.NewSubscriptionPriceModelModelTypeMaxGroupTieredPackage

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeGroupedWithMeteredMinimum = shared.NewSubscriptionPriceModelModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeMatrixWithDisplayName = shared.NewSubscriptionPriceModelModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const NewSubscriptionPriceModelModelTypeGroupedTieredPackage = shared.NewSubscriptionPriceModelModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type NewTaxConfigurationModelUnionParam = shared.NewTaxConfigurationModelUnionParam

// This is an alias to an internal type.
type NewTaxConfigurationModelNewAvalaraTaxConfigurationParam = shared.NewTaxConfigurationModelNewAvalaraTaxConfigurationParam

// This is an alias to an internal type.
type NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProvider = shared.NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProvider

// This is an alias to an internal value.
const NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProviderAvalara = shared.NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProviderAvalara

// This is an alias to an internal type.
type NewTaxConfigurationModelNewTaxJarConfigurationParam = shared.NewTaxConfigurationModelNewTaxJarConfigurationParam

// This is an alias to an internal type.
type NewTaxConfigurationModelNewTaxJarConfigurationTaxProvider = shared.NewTaxConfigurationModelNewTaxJarConfigurationTaxProvider

// This is an alias to an internal value.
const NewTaxConfigurationModelNewTaxJarConfigurationTaxProviderTaxjar = shared.NewTaxConfigurationModelNewTaxJarConfigurationTaxProviderTaxjar

// This is an alias to an internal type.
type NewTaxConfigurationModelTaxProvider = shared.NewTaxConfigurationModelTaxProvider

// This is an alias to an internal value.
const NewTaxConfigurationModelTaxProviderAvalara = shared.NewTaxConfigurationModelTaxProviderAvalara

// This is an alias to an internal value.
const NewTaxConfigurationModelTaxProviderTaxjar = shared.NewTaxConfigurationModelTaxProviderTaxjar

// This is an alias to an internal type.
type PackageConfigModel = shared.PackageConfigModel

// This is an alias to an internal type.
type PackageConfigModelParam = shared.PackageConfigModelParam

// This is an alias to an internal type.
type PaginationMetadata = shared.PaginationMetadata

// This is an alias to an internal type.
type PaymentAttemptModel = shared.PaymentAttemptModel

// The payment provider that attempted to collect the payment.
//
// This is an alias to an internal type.
type PaymentAttemptModelPaymentProvider = shared.PaymentAttemptModelPaymentProvider

// This is an alias to an internal value.
const PaymentAttemptModelPaymentProviderStripe = shared.PaymentAttemptModelPaymentProviderStripe

// This is an alias to an internal type.
type PercentageDiscount = shared.PercentageDiscount

// This is an alias to an internal type.
type PercentageDiscountDiscountType = shared.PercentageDiscountDiscountType

// This is an alias to an internal value.
const PercentageDiscountDiscountTypePercentage = shared.PercentageDiscountDiscountTypePercentage

// This is an alias to an internal type.
type PercentageDiscountParam = shared.PercentageDiscountParam

// This is an alias to an internal type.
type PercentageDiscountIntervalModel = shared.PercentageDiscountIntervalModel

// This is an alias to an internal type.
type PercentageDiscountIntervalModelDiscountType = shared.PercentageDiscountIntervalModelDiscountType

// This is an alias to an internal value.
const PercentageDiscountIntervalModelDiscountTypePercentage = shared.PercentageDiscountIntervalModelDiscountTypePercentage

// This is an alias to an internal type.
type PlanMinifiedModel = shared.PlanMinifiedModel

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
//
// This is an alias to an internal type.
type PlanModel = shared.PlanModel

// This is an alias to an internal type.
type PlanModelPlanPhase = shared.PlanModelPlanPhase

// This is an alias to an internal type.
type PlanModelPlanPhasesDurationUnit = shared.PlanModelPlanPhasesDurationUnit

// This is an alias to an internal value.
const PlanModelPlanPhasesDurationUnitDaily = shared.PlanModelPlanPhasesDurationUnitDaily

// This is an alias to an internal value.
const PlanModelPlanPhasesDurationUnitMonthly = shared.PlanModelPlanPhasesDurationUnitMonthly

// This is an alias to an internal value.
const PlanModelPlanPhasesDurationUnitQuarterly = shared.PlanModelPlanPhasesDurationUnitQuarterly

// This is an alias to an internal value.
const PlanModelPlanPhasesDurationUnitSemiAnnual = shared.PlanModelPlanPhasesDurationUnitSemiAnnual

// This is an alias to an internal value.
const PlanModelPlanPhasesDurationUnitAnnual = shared.PlanModelPlanPhasesDurationUnitAnnual

// This is an alias to an internal type.
type PlanModelProduct = shared.PlanModelProduct

// This is an alias to an internal type.
type PlanModelStatus = shared.PlanModelStatus

// This is an alias to an internal value.
const PlanModelStatusActive = shared.PlanModelStatusActive

// This is an alias to an internal value.
const PlanModelStatusArchived = shared.PlanModelStatusArchived

// This is an alias to an internal value.
const PlanModelStatusDraft = shared.PlanModelStatusDraft

// This is an alias to an internal type.
type PlanModelTrialConfig = shared.PlanModelTrialConfig

// This is an alias to an internal type.
type PlanModelTrialConfigTrialPeriodUnit = shared.PlanModelTrialConfigTrialPeriodUnit

// This is an alias to an internal value.
const PlanModelTrialConfigTrialPeriodUnitDays = shared.PlanModelTrialConfigTrialPeriodUnitDays

// This is an alias to an internal type.
type PriceIntervalFixedFeeQuantityTransitionModelParam = shared.PriceIntervalFixedFeeQuantityTransitionModelParam

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
//
// This is an alias to an internal type.
type PriceIntervalModel = shared.PriceIntervalModel

// This is an alias to an internal type.
type PriceIntervalModelFixedFeeQuantityTransition = shared.PriceIntervalModelFixedFeeQuantityTransition

// The Price resource represents a price that can be billed on a subscription,
// resulting in a charge on an invoice in the form of an invoice line item. Prices
// take a quantity and determine an amount to bill.
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given Price object. The model_type field
// determines the key for the configuration object that is present.
//
// For more on the types of prices, see
// [the core concepts documentation](/core-concepts#plan-and-price)
//
// This is an alias to an internal type.
type PriceModel = shared.PriceModel

// This is an alias to an internal type.
type PriceModelUnitPrice = shared.PriceModelUnitPrice

// This is an alias to an internal type.
type PriceModelUnitPriceCadence = shared.PriceModelUnitPriceCadence

// This is an alias to an internal value.
const PriceModelUnitPriceCadenceOneTime = shared.PriceModelUnitPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelUnitPriceCadenceMonthly = shared.PriceModelUnitPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelUnitPriceCadenceQuarterly = shared.PriceModelUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelUnitPriceCadenceSemiAnnual = shared.PriceModelUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelUnitPriceCadenceAnnual = shared.PriceModelUnitPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelUnitPriceCadenceCustom = shared.PriceModelUnitPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelUnitPriceModelType = shared.PriceModelUnitPriceModelType

// This is an alias to an internal value.
const PriceModelUnitPriceModelTypeUnit = shared.PriceModelUnitPriceModelTypeUnit

// This is an alias to an internal type.
type PriceModelUnitPricePriceType = shared.PriceModelUnitPricePriceType

// This is an alias to an internal value.
const PriceModelUnitPricePriceTypeUsagePrice = shared.PriceModelUnitPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelUnitPricePriceTypeFixedPrice = shared.PriceModelUnitPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelPackagePrice = shared.PriceModelPackagePrice

// This is an alias to an internal type.
type PriceModelPackagePriceCadence = shared.PriceModelPackagePriceCadence

// This is an alias to an internal value.
const PriceModelPackagePriceCadenceOneTime = shared.PriceModelPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelPackagePriceCadenceMonthly = shared.PriceModelPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelPackagePriceCadenceQuarterly = shared.PriceModelPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelPackagePriceCadenceSemiAnnual = shared.PriceModelPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelPackagePriceCadenceAnnual = shared.PriceModelPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelPackagePriceCadenceCustom = shared.PriceModelPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceModelPackagePriceModelType = shared.PriceModelPackagePriceModelType

// This is an alias to an internal value.
const PriceModelPackagePriceModelTypePackage = shared.PriceModelPackagePriceModelTypePackage

// This is an alias to an internal type.
type PriceModelPackagePricePriceType = shared.PriceModelPackagePricePriceType

// This is an alias to an internal value.
const PriceModelPackagePricePriceTypeUsagePrice = shared.PriceModelPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelPackagePricePriceTypeFixedPrice = shared.PriceModelPackagePricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelMatrixPrice = shared.PriceModelMatrixPrice

// This is an alias to an internal type.
type PriceModelMatrixPriceCadence = shared.PriceModelMatrixPriceCadence

// This is an alias to an internal value.
const PriceModelMatrixPriceCadenceOneTime = shared.PriceModelMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelMatrixPriceCadenceMonthly = shared.PriceModelMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelMatrixPriceCadenceQuarterly = shared.PriceModelMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelMatrixPriceCadenceSemiAnnual = shared.PriceModelMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelMatrixPriceCadenceAnnual = shared.PriceModelMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelMatrixPriceCadenceCustom = shared.PriceModelMatrixPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelMatrixPriceModelType = shared.PriceModelMatrixPriceModelType

// This is an alias to an internal value.
const PriceModelMatrixPriceModelTypeMatrix = shared.PriceModelMatrixPriceModelTypeMatrix

// This is an alias to an internal type.
type PriceModelMatrixPricePriceType = shared.PriceModelMatrixPricePriceType

// This is an alias to an internal value.
const PriceModelMatrixPricePriceTypeUsagePrice = shared.PriceModelMatrixPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelMatrixPricePriceTypeFixedPrice = shared.PriceModelMatrixPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelTieredPrice = shared.PriceModelTieredPrice

// This is an alias to an internal type.
type PriceModelTieredPriceCadence = shared.PriceModelTieredPriceCadence

// This is an alias to an internal value.
const PriceModelTieredPriceCadenceOneTime = shared.PriceModelTieredPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelTieredPriceCadenceMonthly = shared.PriceModelTieredPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelTieredPriceCadenceQuarterly = shared.PriceModelTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelTieredPriceCadenceSemiAnnual = shared.PriceModelTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelTieredPriceCadenceAnnual = shared.PriceModelTieredPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelTieredPriceCadenceCustom = shared.PriceModelTieredPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelTieredPriceModelType = shared.PriceModelTieredPriceModelType

// This is an alias to an internal value.
const PriceModelTieredPriceModelTypeTiered = shared.PriceModelTieredPriceModelTypeTiered

// This is an alias to an internal type.
type PriceModelTieredPricePriceType = shared.PriceModelTieredPricePriceType

// This is an alias to an internal value.
const PriceModelTieredPricePriceTypeUsagePrice = shared.PriceModelTieredPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelTieredPricePriceTypeFixedPrice = shared.PriceModelTieredPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelTieredBpsPrice = shared.PriceModelTieredBpsPrice

// This is an alias to an internal type.
type PriceModelTieredBpsPriceCadence = shared.PriceModelTieredBpsPriceCadence

// This is an alias to an internal value.
const PriceModelTieredBpsPriceCadenceOneTime = shared.PriceModelTieredBpsPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelTieredBpsPriceCadenceMonthly = shared.PriceModelTieredBpsPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelTieredBpsPriceCadenceQuarterly = shared.PriceModelTieredBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelTieredBpsPriceCadenceSemiAnnual = shared.PriceModelTieredBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelTieredBpsPriceCadenceAnnual = shared.PriceModelTieredBpsPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelTieredBpsPriceCadenceCustom = shared.PriceModelTieredBpsPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelTieredBpsPriceModelType = shared.PriceModelTieredBpsPriceModelType

// This is an alias to an internal value.
const PriceModelTieredBpsPriceModelTypeTieredBps = shared.PriceModelTieredBpsPriceModelTypeTieredBps

// This is an alias to an internal type.
type PriceModelTieredBpsPricePriceType = shared.PriceModelTieredBpsPricePriceType

// This is an alias to an internal value.
const PriceModelTieredBpsPricePriceTypeUsagePrice = shared.PriceModelTieredBpsPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelTieredBpsPricePriceTypeFixedPrice = shared.PriceModelTieredBpsPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelBpsPrice = shared.PriceModelBpsPrice

// This is an alias to an internal type.
type PriceModelBpsPriceCadence = shared.PriceModelBpsPriceCadence

// This is an alias to an internal value.
const PriceModelBpsPriceCadenceOneTime = shared.PriceModelBpsPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelBpsPriceCadenceMonthly = shared.PriceModelBpsPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelBpsPriceCadenceQuarterly = shared.PriceModelBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelBpsPriceCadenceSemiAnnual = shared.PriceModelBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelBpsPriceCadenceAnnual = shared.PriceModelBpsPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelBpsPriceCadenceCustom = shared.PriceModelBpsPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelBpsPriceModelType = shared.PriceModelBpsPriceModelType

// This is an alias to an internal value.
const PriceModelBpsPriceModelTypeBps = shared.PriceModelBpsPriceModelTypeBps

// This is an alias to an internal type.
type PriceModelBpsPricePriceType = shared.PriceModelBpsPricePriceType

// This is an alias to an internal value.
const PriceModelBpsPricePriceTypeUsagePrice = shared.PriceModelBpsPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelBpsPricePriceTypeFixedPrice = shared.PriceModelBpsPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelBulkBpsPrice = shared.PriceModelBulkBpsPrice

// This is an alias to an internal type.
type PriceModelBulkBpsPriceCadence = shared.PriceModelBulkBpsPriceCadence

// This is an alias to an internal value.
const PriceModelBulkBpsPriceCadenceOneTime = shared.PriceModelBulkBpsPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelBulkBpsPriceCadenceMonthly = shared.PriceModelBulkBpsPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelBulkBpsPriceCadenceQuarterly = shared.PriceModelBulkBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelBulkBpsPriceCadenceSemiAnnual = shared.PriceModelBulkBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelBulkBpsPriceCadenceAnnual = shared.PriceModelBulkBpsPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelBulkBpsPriceCadenceCustom = shared.PriceModelBulkBpsPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelBulkBpsPriceModelType = shared.PriceModelBulkBpsPriceModelType

// This is an alias to an internal value.
const PriceModelBulkBpsPriceModelTypeBulkBps = shared.PriceModelBulkBpsPriceModelTypeBulkBps

// This is an alias to an internal type.
type PriceModelBulkBpsPricePriceType = shared.PriceModelBulkBpsPricePriceType

// This is an alias to an internal value.
const PriceModelBulkBpsPricePriceTypeUsagePrice = shared.PriceModelBulkBpsPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelBulkBpsPricePriceTypeFixedPrice = shared.PriceModelBulkBpsPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelBulkPrice = shared.PriceModelBulkPrice

// This is an alias to an internal type.
type PriceModelBulkPriceCadence = shared.PriceModelBulkPriceCadence

// This is an alias to an internal value.
const PriceModelBulkPriceCadenceOneTime = shared.PriceModelBulkPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelBulkPriceCadenceMonthly = shared.PriceModelBulkPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelBulkPriceCadenceQuarterly = shared.PriceModelBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelBulkPriceCadenceSemiAnnual = shared.PriceModelBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelBulkPriceCadenceAnnual = shared.PriceModelBulkPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelBulkPriceCadenceCustom = shared.PriceModelBulkPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelBulkPriceModelType = shared.PriceModelBulkPriceModelType

// This is an alias to an internal value.
const PriceModelBulkPriceModelTypeBulk = shared.PriceModelBulkPriceModelTypeBulk

// This is an alias to an internal type.
type PriceModelBulkPricePriceType = shared.PriceModelBulkPricePriceType

// This is an alias to an internal value.
const PriceModelBulkPricePriceTypeUsagePrice = shared.PriceModelBulkPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelBulkPricePriceTypeFixedPrice = shared.PriceModelBulkPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelThresholdTotalAmountPrice = shared.PriceModelThresholdTotalAmountPrice

// This is an alias to an internal type.
type PriceModelThresholdTotalAmountPriceCadence = shared.PriceModelThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceCadenceOneTime = shared.PriceModelThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceCadenceMonthly = shared.PriceModelThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceCadenceQuarterly = shared.PriceModelThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceCadenceSemiAnnual = shared.PriceModelThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceCadenceAnnual = shared.PriceModelThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceCadenceCustom = shared.PriceModelThresholdTotalAmountPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelThresholdTotalAmountPriceModelType = shared.PriceModelThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.PriceModelThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// This is an alias to an internal type.
type PriceModelThresholdTotalAmountPricePriceType = shared.PriceModelThresholdTotalAmountPricePriceType

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPricePriceTypeUsagePrice = shared.PriceModelThresholdTotalAmountPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelThresholdTotalAmountPricePriceTypeFixedPrice = shared.PriceModelThresholdTotalAmountPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelTieredPackagePrice = shared.PriceModelTieredPackagePrice

// This is an alias to an internal type.
type PriceModelTieredPackagePriceCadence = shared.PriceModelTieredPackagePriceCadence

// This is an alias to an internal value.
const PriceModelTieredPackagePriceCadenceOneTime = shared.PriceModelTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelTieredPackagePriceCadenceMonthly = shared.PriceModelTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelTieredPackagePriceCadenceQuarterly = shared.PriceModelTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelTieredPackagePriceCadenceSemiAnnual = shared.PriceModelTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelTieredPackagePriceCadenceAnnual = shared.PriceModelTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelTieredPackagePriceCadenceCustom = shared.PriceModelTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceModelTieredPackagePriceModelType = shared.PriceModelTieredPackagePriceModelType

// This is an alias to an internal value.
const PriceModelTieredPackagePriceModelTypeTieredPackage = shared.PriceModelTieredPackagePriceModelTypeTieredPackage

// This is an alias to an internal type.
type PriceModelTieredPackagePricePriceType = shared.PriceModelTieredPackagePricePriceType

// This is an alias to an internal value.
const PriceModelTieredPackagePricePriceTypeUsagePrice = shared.PriceModelTieredPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelTieredPackagePricePriceTypeFixedPrice = shared.PriceModelTieredPackagePricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelGroupedTieredPrice = shared.PriceModelGroupedTieredPrice

// This is an alias to an internal type.
type PriceModelGroupedTieredPriceCadence = shared.PriceModelGroupedTieredPriceCadence

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceCadenceOneTime = shared.PriceModelGroupedTieredPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceCadenceMonthly = shared.PriceModelGroupedTieredPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceCadenceQuarterly = shared.PriceModelGroupedTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceCadenceSemiAnnual = shared.PriceModelGroupedTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceCadenceAnnual = shared.PriceModelGroupedTieredPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceCadenceCustom = shared.PriceModelGroupedTieredPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelGroupedTieredPriceModelType = shared.PriceModelGroupedTieredPriceModelType

// This is an alias to an internal value.
const PriceModelGroupedTieredPriceModelTypeGroupedTiered = shared.PriceModelGroupedTieredPriceModelTypeGroupedTiered

// This is an alias to an internal type.
type PriceModelGroupedTieredPricePriceType = shared.PriceModelGroupedTieredPricePriceType

// This is an alias to an internal value.
const PriceModelGroupedTieredPricePriceTypeUsagePrice = shared.PriceModelGroupedTieredPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelGroupedTieredPricePriceTypeFixedPrice = shared.PriceModelGroupedTieredPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelTieredWithMinimumPrice = shared.PriceModelTieredWithMinimumPrice

// This is an alias to an internal type.
type PriceModelTieredWithMinimumPriceCadence = shared.PriceModelTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceCadenceOneTime = shared.PriceModelTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceCadenceMonthly = shared.PriceModelTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceCadenceQuarterly = shared.PriceModelTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceCadenceSemiAnnual = shared.PriceModelTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceCadenceAnnual = shared.PriceModelTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceCadenceCustom = shared.PriceModelTieredWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelTieredWithMinimumPriceModelType = shared.PriceModelTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.PriceModelTieredWithMinimumPriceModelTypeTieredWithMinimum

// This is an alias to an internal type.
type PriceModelTieredWithMinimumPricePriceType = shared.PriceModelTieredWithMinimumPricePriceType

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPricePriceTypeUsagePrice = shared.PriceModelTieredWithMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelTieredWithMinimumPricePriceTypeFixedPrice = shared.PriceModelTieredWithMinimumPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelTieredPackageWithMinimumPrice = shared.PriceModelTieredPackageWithMinimumPrice

// This is an alias to an internal type.
type PriceModelTieredPackageWithMinimumPriceCadence = shared.PriceModelTieredPackageWithMinimumPriceCadence

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceCadenceOneTime = shared.PriceModelTieredPackageWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceCadenceMonthly = shared.PriceModelTieredPackageWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceCadenceQuarterly = shared.PriceModelTieredPackageWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceCadenceSemiAnnual = shared.PriceModelTieredPackageWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceCadenceAnnual = shared.PriceModelTieredPackageWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceCadenceCustom = shared.PriceModelTieredPackageWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelTieredPackageWithMinimumPriceModelType = shared.PriceModelTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.PriceModelTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// This is an alias to an internal type.
type PriceModelTieredPackageWithMinimumPricePriceType = shared.PriceModelTieredPackageWithMinimumPricePriceType

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPricePriceTypeUsagePrice = shared.PriceModelTieredPackageWithMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelTieredPackageWithMinimumPricePriceTypeFixedPrice = shared.PriceModelTieredPackageWithMinimumPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelPackageWithAllocationPrice = shared.PriceModelPackageWithAllocationPrice

// This is an alias to an internal type.
type PriceModelPackageWithAllocationPriceCadence = shared.PriceModelPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceCadenceOneTime = shared.PriceModelPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceCadenceMonthly = shared.PriceModelPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceCadenceQuarterly = shared.PriceModelPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceCadenceSemiAnnual = shared.PriceModelPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceCadenceAnnual = shared.PriceModelPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceCadenceCustom = shared.PriceModelPackageWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelPackageWithAllocationPriceModelType = shared.PriceModelPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPriceModelTypePackageWithAllocation = shared.PriceModelPackageWithAllocationPriceModelTypePackageWithAllocation

// This is an alias to an internal type.
type PriceModelPackageWithAllocationPricePriceType = shared.PriceModelPackageWithAllocationPricePriceType

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPricePriceTypeUsagePrice = shared.PriceModelPackageWithAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelPackageWithAllocationPricePriceTypeFixedPrice = shared.PriceModelPackageWithAllocationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelUnitWithPercentPrice = shared.PriceModelUnitWithPercentPrice

// This is an alias to an internal type.
type PriceModelUnitWithPercentPriceCadence = shared.PriceModelUnitWithPercentPriceCadence

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceCadenceOneTime = shared.PriceModelUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceCadenceMonthly = shared.PriceModelUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceCadenceQuarterly = shared.PriceModelUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceCadenceSemiAnnual = shared.PriceModelUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceCadenceAnnual = shared.PriceModelUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceCadenceCustom = shared.PriceModelUnitWithPercentPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelUnitWithPercentPriceModelType = shared.PriceModelUnitWithPercentPriceModelType

// This is an alias to an internal value.
const PriceModelUnitWithPercentPriceModelTypeUnitWithPercent = shared.PriceModelUnitWithPercentPriceModelTypeUnitWithPercent

// This is an alias to an internal type.
type PriceModelUnitWithPercentPricePriceType = shared.PriceModelUnitWithPercentPricePriceType

// This is an alias to an internal value.
const PriceModelUnitWithPercentPricePriceTypeUsagePrice = shared.PriceModelUnitWithPercentPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelUnitWithPercentPricePriceTypeFixedPrice = shared.PriceModelUnitWithPercentPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelMatrixWithAllocationPrice = shared.PriceModelMatrixWithAllocationPrice

// This is an alias to an internal type.
type PriceModelMatrixWithAllocationPriceCadence = shared.PriceModelMatrixWithAllocationPriceCadence

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceCadenceOneTime = shared.PriceModelMatrixWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceCadenceMonthly = shared.PriceModelMatrixWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceCadenceQuarterly = shared.PriceModelMatrixWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceCadenceSemiAnnual = shared.PriceModelMatrixWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceCadenceAnnual = shared.PriceModelMatrixWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceCadenceCustom = shared.PriceModelMatrixWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelMatrixWithAllocationPriceModelType = shared.PriceModelMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.PriceModelMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// This is an alias to an internal type.
type PriceModelMatrixWithAllocationPricePriceType = shared.PriceModelMatrixWithAllocationPricePriceType

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPricePriceTypeUsagePrice = shared.PriceModelMatrixWithAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelMatrixWithAllocationPricePriceTypeFixedPrice = shared.PriceModelMatrixWithAllocationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelTieredWithProrationPrice = shared.PriceModelTieredWithProrationPrice

// This is an alias to an internal type.
type PriceModelTieredWithProrationPriceCadence = shared.PriceModelTieredWithProrationPriceCadence

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceCadenceOneTime = shared.PriceModelTieredWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceCadenceMonthly = shared.PriceModelTieredWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceCadenceQuarterly = shared.PriceModelTieredWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceCadenceSemiAnnual = shared.PriceModelTieredWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceCadenceAnnual = shared.PriceModelTieredWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceCadenceCustom = shared.PriceModelTieredWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelTieredWithProrationPriceModelType = shared.PriceModelTieredWithProrationPriceModelType

// This is an alias to an internal value.
const PriceModelTieredWithProrationPriceModelTypeTieredWithProration = shared.PriceModelTieredWithProrationPriceModelTypeTieredWithProration

// This is an alias to an internal type.
type PriceModelTieredWithProrationPricePriceType = shared.PriceModelTieredWithProrationPricePriceType

// This is an alias to an internal value.
const PriceModelTieredWithProrationPricePriceTypeUsagePrice = shared.PriceModelTieredWithProrationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelTieredWithProrationPricePriceTypeFixedPrice = shared.PriceModelTieredWithProrationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelUnitWithProrationPrice = shared.PriceModelUnitWithProrationPrice

// This is an alias to an internal type.
type PriceModelUnitWithProrationPriceCadence = shared.PriceModelUnitWithProrationPriceCadence

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceCadenceOneTime = shared.PriceModelUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceCadenceMonthly = shared.PriceModelUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceCadenceQuarterly = shared.PriceModelUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceCadenceSemiAnnual = shared.PriceModelUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceCadenceAnnual = shared.PriceModelUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceCadenceCustom = shared.PriceModelUnitWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelUnitWithProrationPriceModelType = shared.PriceModelUnitWithProrationPriceModelType

// This is an alias to an internal value.
const PriceModelUnitWithProrationPriceModelTypeUnitWithProration = shared.PriceModelUnitWithProrationPriceModelTypeUnitWithProration

// This is an alias to an internal type.
type PriceModelUnitWithProrationPricePriceType = shared.PriceModelUnitWithProrationPricePriceType

// This is an alias to an internal value.
const PriceModelUnitWithProrationPricePriceTypeUsagePrice = shared.PriceModelUnitWithProrationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelUnitWithProrationPricePriceTypeFixedPrice = shared.PriceModelUnitWithProrationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelGroupedAllocationPrice = shared.PriceModelGroupedAllocationPrice

// This is an alias to an internal type.
type PriceModelGroupedAllocationPriceCadence = shared.PriceModelGroupedAllocationPriceCadence

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceCadenceOneTime = shared.PriceModelGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceCadenceMonthly = shared.PriceModelGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceCadenceQuarterly = shared.PriceModelGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceCadenceSemiAnnual = shared.PriceModelGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceCadenceAnnual = shared.PriceModelGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceCadenceCustom = shared.PriceModelGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelGroupedAllocationPriceModelType = shared.PriceModelGroupedAllocationPriceModelType

// This is an alias to an internal value.
const PriceModelGroupedAllocationPriceModelTypeGroupedAllocation = shared.PriceModelGroupedAllocationPriceModelTypeGroupedAllocation

// This is an alias to an internal type.
type PriceModelGroupedAllocationPricePriceType = shared.PriceModelGroupedAllocationPricePriceType

// This is an alias to an internal value.
const PriceModelGroupedAllocationPricePriceTypeUsagePrice = shared.PriceModelGroupedAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelGroupedAllocationPricePriceTypeFixedPrice = shared.PriceModelGroupedAllocationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelGroupedWithProratedMinimumPrice = shared.PriceModelGroupedWithProratedMinimumPrice

// This is an alias to an internal type.
type PriceModelGroupedWithProratedMinimumPriceCadence = shared.PriceModelGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceCadenceOneTime = shared.PriceModelGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceCadenceMonthly = shared.PriceModelGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceCadenceQuarterly = shared.PriceModelGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.PriceModelGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceCadenceAnnual = shared.PriceModelGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceCadenceCustom = shared.PriceModelGroupedWithProratedMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelGroupedWithProratedMinimumPriceModelType = shared.PriceModelGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.PriceModelGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal type.
type PriceModelGroupedWithProratedMinimumPricePriceType = shared.PriceModelGroupedWithProratedMinimumPricePriceType

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPricePriceTypeUsagePrice = shared.PriceModelGroupedWithProratedMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelGroupedWithProratedMinimumPricePriceTypeFixedPrice = shared.PriceModelGroupedWithProratedMinimumPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelGroupedWithMeteredMinimumPrice = shared.PriceModelGroupedWithMeteredMinimumPrice

// This is an alias to an internal type.
type PriceModelGroupedWithMeteredMinimumPriceCadence = shared.PriceModelGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceCadenceOneTime = shared.PriceModelGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceCadenceMonthly = shared.PriceModelGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.PriceModelGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.PriceModelGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceCadenceAnnual = shared.PriceModelGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceCadenceCustom = shared.PriceModelGroupedWithMeteredMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelGroupedWithMeteredMinimumPriceModelType = shared.PriceModelGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.PriceModelGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal type.
type PriceModelGroupedWithMeteredMinimumPricePriceType = shared.PriceModelGroupedWithMeteredMinimumPricePriceType

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPricePriceTypeUsagePrice = shared.PriceModelGroupedWithMeteredMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelGroupedWithMeteredMinimumPricePriceTypeFixedPrice = shared.PriceModelGroupedWithMeteredMinimumPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelMatrixWithDisplayNamePrice = shared.PriceModelMatrixWithDisplayNamePrice

// This is an alias to an internal type.
type PriceModelMatrixWithDisplayNamePriceCadence = shared.PriceModelMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceCadenceOneTime = shared.PriceModelMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceCadenceMonthly = shared.PriceModelMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceCadenceQuarterly = shared.PriceModelMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.PriceModelMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceCadenceAnnual = shared.PriceModelMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceCadenceCustom = shared.PriceModelMatrixWithDisplayNamePriceCadenceCustom

// This is an alias to an internal type.
type PriceModelMatrixWithDisplayNamePriceModelType = shared.PriceModelMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.PriceModelMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// This is an alias to an internal type.
type PriceModelMatrixWithDisplayNamePricePriceType = shared.PriceModelMatrixWithDisplayNamePricePriceType

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePricePriceTypeUsagePrice = shared.PriceModelMatrixWithDisplayNamePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelMatrixWithDisplayNamePricePriceTypeFixedPrice = shared.PriceModelMatrixWithDisplayNamePricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelBulkWithProrationPrice = shared.PriceModelBulkWithProrationPrice

// This is an alias to an internal type.
type PriceModelBulkWithProrationPriceCadence = shared.PriceModelBulkWithProrationPriceCadence

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceCadenceOneTime = shared.PriceModelBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceCadenceMonthly = shared.PriceModelBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceCadenceQuarterly = shared.PriceModelBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceCadenceSemiAnnual = shared.PriceModelBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceCadenceAnnual = shared.PriceModelBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceCadenceCustom = shared.PriceModelBulkWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelBulkWithProrationPriceModelType = shared.PriceModelBulkWithProrationPriceModelType

// This is an alias to an internal value.
const PriceModelBulkWithProrationPriceModelTypeBulkWithProration = shared.PriceModelBulkWithProrationPriceModelTypeBulkWithProration

// This is an alias to an internal type.
type PriceModelBulkWithProrationPricePriceType = shared.PriceModelBulkWithProrationPricePriceType

// This is an alias to an internal value.
const PriceModelBulkWithProrationPricePriceTypeUsagePrice = shared.PriceModelBulkWithProrationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelBulkWithProrationPricePriceTypeFixedPrice = shared.PriceModelBulkWithProrationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelGroupedTieredPackagePrice = shared.PriceModelGroupedTieredPackagePrice

// This is an alias to an internal type.
type PriceModelGroupedTieredPackagePriceCadence = shared.PriceModelGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceCadenceOneTime = shared.PriceModelGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceCadenceMonthly = shared.PriceModelGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceCadenceQuarterly = shared.PriceModelGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceCadenceSemiAnnual = shared.PriceModelGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceCadenceAnnual = shared.PriceModelGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceCadenceCustom = shared.PriceModelGroupedTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceModelGroupedTieredPackagePriceModelType = shared.PriceModelGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.PriceModelGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type PriceModelGroupedTieredPackagePricePriceType = shared.PriceModelGroupedTieredPackagePricePriceType

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePricePriceTypeUsagePrice = shared.PriceModelGroupedTieredPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelGroupedTieredPackagePricePriceTypeFixedPrice = shared.PriceModelGroupedTieredPackagePricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelMaxGroupTieredPackagePrice = shared.PriceModelMaxGroupTieredPackagePrice

// This is an alias to an internal type.
type PriceModelMaxGroupTieredPackagePriceCadence = shared.PriceModelMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceCadenceOneTime = shared.PriceModelMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceCadenceMonthly = shared.PriceModelMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceCadenceQuarterly = shared.PriceModelMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.PriceModelMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceCadenceAnnual = shared.PriceModelMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceCadenceCustom = shared.PriceModelMaxGroupTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceModelMaxGroupTieredPackagePriceModelType = shared.PriceModelMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.PriceModelMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal type.
type PriceModelMaxGroupTieredPackagePricePriceType = shared.PriceModelMaxGroupTieredPackagePricePriceType

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePricePriceTypeUsagePrice = shared.PriceModelMaxGroupTieredPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelMaxGroupTieredPackagePricePriceTypeFixedPrice = shared.PriceModelMaxGroupTieredPackagePricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelScalableMatrixWithUnitPricingPrice = shared.PriceModelScalableMatrixWithUnitPricingPrice

// This is an alias to an internal type.
type PriceModelScalableMatrixWithUnitPricingPriceCadence = shared.PriceModelScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.PriceModelScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.PriceModelScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.PriceModelScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.PriceModelScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.PriceModelScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceCadenceCustom = shared.PriceModelScalableMatrixWithUnitPricingPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelScalableMatrixWithUnitPricingPriceModelType = shared.PriceModelScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.PriceModelScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal type.
type PriceModelScalableMatrixWithUnitPricingPricePriceType = shared.PriceModelScalableMatrixWithUnitPricingPricePriceType

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPricePriceTypeUsagePrice = shared.PriceModelScalableMatrixWithUnitPricingPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelScalableMatrixWithUnitPricingPricePriceTypeFixedPrice = shared.PriceModelScalableMatrixWithUnitPricingPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelScalableMatrixWithTieredPricingPrice = shared.PriceModelScalableMatrixWithTieredPricingPrice

// This is an alias to an internal type.
type PriceModelScalableMatrixWithTieredPricingPriceCadence = shared.PriceModelScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.PriceModelScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.PriceModelScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.PriceModelScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.PriceModelScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.PriceModelScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceCadenceCustom = shared.PriceModelScalableMatrixWithTieredPricingPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelScalableMatrixWithTieredPricingPriceModelType = shared.PriceModelScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.PriceModelScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal type.
type PriceModelScalableMatrixWithTieredPricingPricePriceType = shared.PriceModelScalableMatrixWithTieredPricingPricePriceType

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPricePriceTypeUsagePrice = shared.PriceModelScalableMatrixWithTieredPricingPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelScalableMatrixWithTieredPricingPricePriceTypeFixedPrice = shared.PriceModelScalableMatrixWithTieredPricingPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelCumulativeGroupedBulkPrice = shared.PriceModelCumulativeGroupedBulkPrice

// This is an alias to an internal type.
type PriceModelCumulativeGroupedBulkPriceCadence = shared.PriceModelCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceCadenceOneTime = shared.PriceModelCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceCadenceMonthly = shared.PriceModelCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceCadenceQuarterly = shared.PriceModelCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.PriceModelCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceCadenceAnnual = shared.PriceModelCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceCadenceCustom = shared.PriceModelCumulativeGroupedBulkPriceCadenceCustom

// This is an alias to an internal type.
type PriceModelCumulativeGroupedBulkPriceModelType = shared.PriceModelCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.PriceModelCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type PriceModelCumulativeGroupedBulkPricePriceType = shared.PriceModelCumulativeGroupedBulkPricePriceType

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPricePriceTypeUsagePrice = shared.PriceModelCumulativeGroupedBulkPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelCumulativeGroupedBulkPricePriceTypeFixedPrice = shared.PriceModelCumulativeGroupedBulkPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceModelCadence = shared.PriceModelCadence

// This is an alias to an internal value.
const PriceModelCadenceOneTime = shared.PriceModelCadenceOneTime

// This is an alias to an internal value.
const PriceModelCadenceMonthly = shared.PriceModelCadenceMonthly

// This is an alias to an internal value.
const PriceModelCadenceQuarterly = shared.PriceModelCadenceQuarterly

// This is an alias to an internal value.
const PriceModelCadenceSemiAnnual = shared.PriceModelCadenceSemiAnnual

// This is an alias to an internal value.
const PriceModelCadenceAnnual = shared.PriceModelCadenceAnnual

// This is an alias to an internal value.
const PriceModelCadenceCustom = shared.PriceModelCadenceCustom

// This is an alias to an internal type.
type PriceModelModelType = shared.PriceModelModelType

// This is an alias to an internal value.
const PriceModelModelTypeUnit = shared.PriceModelModelTypeUnit

// This is an alias to an internal value.
const PriceModelModelTypePackage = shared.PriceModelModelTypePackage

// This is an alias to an internal value.
const PriceModelModelTypeMatrix = shared.PriceModelModelTypeMatrix

// This is an alias to an internal value.
const PriceModelModelTypeTiered = shared.PriceModelModelTypeTiered

// This is an alias to an internal value.
const PriceModelModelTypeTieredBps = shared.PriceModelModelTypeTieredBps

// This is an alias to an internal value.
const PriceModelModelTypeBps = shared.PriceModelModelTypeBps

// This is an alias to an internal value.
const PriceModelModelTypeBulkBps = shared.PriceModelModelTypeBulkBps

// This is an alias to an internal value.
const PriceModelModelTypeBulk = shared.PriceModelModelTypeBulk

// This is an alias to an internal value.
const PriceModelModelTypeThresholdTotalAmount = shared.PriceModelModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const PriceModelModelTypeTieredPackage = shared.PriceModelModelTypeTieredPackage

// This is an alias to an internal value.
const PriceModelModelTypeGroupedTiered = shared.PriceModelModelTypeGroupedTiered

// This is an alias to an internal value.
const PriceModelModelTypeTieredWithMinimum = shared.PriceModelModelTypeTieredWithMinimum

// This is an alias to an internal value.
const PriceModelModelTypeTieredPackageWithMinimum = shared.PriceModelModelTypeTieredPackageWithMinimum

// This is an alias to an internal value.
const PriceModelModelTypePackageWithAllocation = shared.PriceModelModelTypePackageWithAllocation

// This is an alias to an internal value.
const PriceModelModelTypeUnitWithPercent = shared.PriceModelModelTypeUnitWithPercent

// This is an alias to an internal value.
const PriceModelModelTypeMatrixWithAllocation = shared.PriceModelModelTypeMatrixWithAllocation

// This is an alias to an internal value.
const PriceModelModelTypeTieredWithProration = shared.PriceModelModelTypeTieredWithProration

// This is an alias to an internal value.
const PriceModelModelTypeUnitWithProration = shared.PriceModelModelTypeUnitWithProration

// This is an alias to an internal value.
const PriceModelModelTypeGroupedAllocation = shared.PriceModelModelTypeGroupedAllocation

// This is an alias to an internal value.
const PriceModelModelTypeGroupedWithProratedMinimum = shared.PriceModelModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const PriceModelModelTypeGroupedWithMeteredMinimum = shared.PriceModelModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const PriceModelModelTypeMatrixWithDisplayName = shared.PriceModelModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const PriceModelModelTypeBulkWithProration = shared.PriceModelModelTypeBulkWithProration

// This is an alias to an internal value.
const PriceModelModelTypeGroupedTieredPackage = shared.PriceModelModelTypeGroupedTieredPackage

// This is an alias to an internal value.
const PriceModelModelTypeMaxGroupTieredPackage = shared.PriceModelModelTypeMaxGroupTieredPackage

// This is an alias to an internal value.
const PriceModelModelTypeScalableMatrixWithUnitPricing = shared.PriceModelModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal value.
const PriceModelModelTypeScalableMatrixWithTieredPricing = shared.PriceModelModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal value.
const PriceModelModelTypeCumulativeGroupedBulk = shared.PriceModelModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type PriceModelPriceType = shared.PriceModelPriceType

// This is an alias to an internal value.
const PriceModelPriceTypeUsagePrice = shared.PriceModelPriceTypeUsagePrice

// This is an alias to an internal value.
const PriceModelPriceTypeFixedPrice = shared.PriceModelPriceTypeFixedPrice

// This is an alias to an internal type.
type RemoveSubscriptionAdjustmentParams = shared.RemoveSubscriptionAdjustmentParams

// This is an alias to an internal type.
type RemoveSubscriptionPriceParams = shared.RemoveSubscriptionPriceParams

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParams = shared.ReplaceSubscriptionAdjustmentParams

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParams = shared.ReplaceSubscriptionPriceParams

// This is an alias to an internal type.
type SubLineItemGroupingModel = shared.SubLineItemGroupingModel

// This is an alias to an internal type.
type SubscriptionMinifiedModel = shared.SubscriptionMinifiedModel

// A [subscription](/core-concepts#subscription) represents the purchase of a plan
// by a customer.
//
// By default, subscriptions begin on the day that they're created and renew
// automatically for each billing cycle at the cadence that's configured in the
// plan definition.
//
// Subscriptions also default to **beginning of month alignment**, which means the
// first invoice issued for the subscription will have pro-rated charges between
// the `start_date` and the first of the following month. Subsequent billing
// periods will always start and end on a month boundary (e.g. subsequent month
// starts for monthly billing).
//
// Depending on the plan configuration, any _flat_ recurring fees will be billed
// either at the beginning (in-advance) or end (in-arrears) of each billing cycle.
// Plans default to **in-advance billing**. Usage-based fees are billed in arrears
// as usage is accumulated. In the normal course of events, you can expect an
// invoice to contain usage-based charges for the previous period, and a recurring
// fee for the following period.
//
// This is an alias to an internal type.
type SubscriptionModel = shared.SubscriptionModel

// This is an alias to an internal type.
type SubscriptionModelDiscountInterval = shared.SubscriptionModelDiscountInterval

// This is an alias to an internal type.
type SubscriptionModelDiscountIntervalsDiscountType = shared.SubscriptionModelDiscountIntervalsDiscountType

// This is an alias to an internal value.
const SubscriptionModelDiscountIntervalsDiscountTypeAmount = shared.SubscriptionModelDiscountIntervalsDiscountTypeAmount

// This is an alias to an internal value.
const SubscriptionModelDiscountIntervalsDiscountTypePercentage = shared.SubscriptionModelDiscountIntervalsDiscountTypePercentage

// This is an alias to an internal value.
const SubscriptionModelDiscountIntervalsDiscountTypeUsage = shared.SubscriptionModelDiscountIntervalsDiscountTypeUsage

// This is an alias to an internal type.
type SubscriptionModelStatus = shared.SubscriptionModelStatus

// This is an alias to an internal value.
const SubscriptionModelStatusActive = shared.SubscriptionModelStatusActive

// This is an alias to an internal value.
const SubscriptionModelStatusEnded = shared.SubscriptionModelStatusEnded

// This is an alias to an internal value.
const SubscriptionModelStatusUpcoming = shared.SubscriptionModelStatusUpcoming

// This is an alias to an internal type.
type SubscriptionTrialInfoModel = shared.SubscriptionTrialInfoModel

// This is an alias to an internal type.
type SubscriptionsModel = shared.SubscriptionsModel

// This is an alias to an internal type.
type TaxAmountModel = shared.TaxAmountModel

// Thresholds are used to define the conditions under which an alert will be
// triggered.
//
// This is an alias to an internal type.
type ThresholdModel = shared.ThresholdModel

// Thresholds are used to define the conditions under which an alert will be
// triggered.
//
// This is an alias to an internal type.
type ThresholdModelParam = shared.ThresholdModelParam

// This is an alias to an internal type.
type TieredBpsConfigModel = shared.TieredBpsConfigModel

// This is an alias to an internal type.
type TieredBpsConfigModelTier = shared.TieredBpsConfigModelTier

// This is an alias to an internal type.
type TieredBpsConfigModelParam = shared.TieredBpsConfigModelParam

// This is an alias to an internal type.
type TieredBpsConfigModelTierParam = shared.TieredBpsConfigModelTierParam

// This is an alias to an internal type.
type TieredConfigModel = shared.TieredConfigModel

// This is an alias to an internal type.
type TieredConfigModelTier = shared.TieredConfigModelTier

// This is an alias to an internal type.
type TieredConfigModelParam = shared.TieredConfigModelParam

// This is an alias to an internal type.
type TieredConfigModelTierParam = shared.TieredConfigModelTierParam

// This is an alias to an internal type.
type TopUpModel = shared.TopUpModel

// Settings for invoices generated by triggered top-ups.
//
// This is an alias to an internal type.
type TopUpModelInvoiceSettings = shared.TopUpModelInvoiceSettings

// The unit of expires_after.
//
// This is an alias to an internal type.
type TopUpModelExpiresAfterUnit = shared.TopUpModelExpiresAfterUnit

// This is an alias to an internal value.
const TopUpModelExpiresAfterUnitDay = shared.TopUpModelExpiresAfterUnitDay

// This is an alias to an internal value.
const TopUpModelExpiresAfterUnitMonth = shared.TopUpModelExpiresAfterUnitMonth

// This is an alias to an internal type.
type TopUpsModel = shared.TopUpsModel

// This is an alias to an internal type.
type TrialDiscount = shared.TrialDiscount

// This is an alias to an internal type.
type TrialDiscountDiscountType = shared.TrialDiscountDiscountType

// This is an alias to an internal value.
const TrialDiscountDiscountTypeTrial = shared.TrialDiscountDiscountTypeTrial

// This is an alias to an internal type.
type TrialDiscountParam = shared.TrialDiscountParam

// This is an alias to an internal type.
type UnitConfigModel = shared.UnitConfigModel

// This is an alias to an internal type.
type UnitConfigModelParam = shared.UnitConfigModelParam

// This is an alias to an internal type.
type UsageDiscountIntervalModel = shared.UsageDiscountIntervalModel

// This is an alias to an internal type.
type UsageDiscountIntervalModelDiscountType = shared.UsageDiscountIntervalModelDiscountType

// This is an alias to an internal value.
const UsageDiscountIntervalModelDiscountTypeUsage = shared.UsageDiscountIntervalModelDiscountTypeUsage

// This is an alias to an internal type.
type UsageModel = shared.UsageModel
