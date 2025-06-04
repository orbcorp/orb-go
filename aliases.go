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
type AmountDiscount = shared.AmountDiscount

// This is an alias to an internal type.
type AmountDiscountDiscountType = shared.AmountDiscountDiscountType

// This is an alias to an internal value.
const AmountDiscountDiscountTypeAmount = shared.AmountDiscountDiscountTypeAmount

// This is an alias to an internal type.
type AmountDiscountFilter = shared.AmountDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type AmountDiscountFiltersField = shared.AmountDiscountFiltersField

// This is an alias to an internal value.
const AmountDiscountFiltersFieldPriceID = shared.AmountDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const AmountDiscountFiltersFieldItemID = shared.AmountDiscountFiltersFieldItemID

// This is an alias to an internal value.
const AmountDiscountFiltersFieldPriceType = shared.AmountDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const AmountDiscountFiltersFieldCurrency = shared.AmountDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const AmountDiscountFiltersFieldPricingUnitID = shared.AmountDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type AmountDiscountFiltersOperator = shared.AmountDiscountFiltersOperator

// This is an alias to an internal value.
const AmountDiscountFiltersOperatorIncludes = shared.AmountDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const AmountDiscountFiltersOperatorExcludes = shared.AmountDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type AmountDiscountParam = shared.AmountDiscountParam

// This is an alias to an internal type.
type AmountDiscountFilterParam = shared.AmountDiscountFilterParam

// This is an alias to an internal type.
type BillingCycleRelativeDate = shared.BillingCycleRelativeDate

// This is an alias to an internal value.
const BillingCycleRelativeDateStartOfTerm = shared.BillingCycleRelativeDateStartOfTerm

// This is an alias to an internal value.
const BillingCycleRelativeDateEndOfTerm = shared.BillingCycleRelativeDateEndOfTerm

// This is an alias to an internal type.
type Discount = shared.Discount

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
type PaginationMetadata = shared.PaginationMetadata

// This is an alias to an internal type.
type PercentageDiscount = shared.PercentageDiscount

// This is an alias to an internal type.
type PercentageDiscountDiscountType = shared.PercentageDiscountDiscountType

// This is an alias to an internal value.
const PercentageDiscountDiscountTypePercentage = shared.PercentageDiscountDiscountTypePercentage

// This is an alias to an internal type.
type PercentageDiscountFilter = shared.PercentageDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PercentageDiscountFiltersField = shared.PercentageDiscountFiltersField

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldPriceID = shared.PercentageDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldItemID = shared.PercentageDiscountFiltersFieldItemID

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldPriceType = shared.PercentageDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldCurrency = shared.PercentageDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldPricingUnitID = shared.PercentageDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PercentageDiscountFiltersOperator = shared.PercentageDiscountFiltersOperator

// This is an alias to an internal value.
const PercentageDiscountFiltersOperatorIncludes = shared.PercentageDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const PercentageDiscountFiltersOperatorExcludes = shared.PercentageDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type PercentageDiscountParam = shared.PercentageDiscountParam

// This is an alias to an internal type.
type PercentageDiscountFilterParam = shared.PercentageDiscountFilterParam

// This is an alias to an internal type.
type TrialDiscount = shared.TrialDiscount

// This is an alias to an internal type.
type TrialDiscountDiscountType = shared.TrialDiscountDiscountType

// This is an alias to an internal value.
const TrialDiscountDiscountTypeTrial = shared.TrialDiscountDiscountTypeTrial

// This is an alias to an internal type.
type TrialDiscountFilter = shared.TrialDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type TrialDiscountFiltersField = shared.TrialDiscountFiltersField

// This is an alias to an internal value.
const TrialDiscountFiltersFieldPriceID = shared.TrialDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const TrialDiscountFiltersFieldItemID = shared.TrialDiscountFiltersFieldItemID

// This is an alias to an internal value.
const TrialDiscountFiltersFieldPriceType = shared.TrialDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const TrialDiscountFiltersFieldCurrency = shared.TrialDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const TrialDiscountFiltersFieldPricingUnitID = shared.TrialDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type TrialDiscountFiltersOperator = shared.TrialDiscountFiltersOperator

// This is an alias to an internal value.
const TrialDiscountFiltersOperatorIncludes = shared.TrialDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const TrialDiscountFiltersOperatorExcludes = shared.TrialDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type TrialDiscountParam = shared.TrialDiscountParam

// This is an alias to an internal type.
type TrialDiscountFilterParam = shared.TrialDiscountFilterParam

// This is an alias to an internal type.
type UsageDiscount = shared.UsageDiscount

// This is an alias to an internal type.
type UsageDiscountDiscountType = shared.UsageDiscountDiscountType

// This is an alias to an internal value.
const UsageDiscountDiscountTypeUsage = shared.UsageDiscountDiscountTypeUsage

// This is an alias to an internal type.
type UsageDiscountFilter = shared.UsageDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type UsageDiscountFiltersField = shared.UsageDiscountFiltersField

// This is an alias to an internal value.
const UsageDiscountFiltersFieldPriceID = shared.UsageDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const UsageDiscountFiltersFieldItemID = shared.UsageDiscountFiltersFieldItemID

// This is an alias to an internal value.
const UsageDiscountFiltersFieldPriceType = shared.UsageDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const UsageDiscountFiltersFieldCurrency = shared.UsageDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const UsageDiscountFiltersFieldPricingUnitID = shared.UsageDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type UsageDiscountFiltersOperator = shared.UsageDiscountFiltersOperator

// This is an alias to an internal value.
const UsageDiscountFiltersOperatorIncludes = shared.UsageDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const UsageDiscountFiltersOperatorExcludes = shared.UsageDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type UsageDiscountParam = shared.UsageDiscountParam

// This is an alias to an internal type.
type UsageDiscountFilterParam = shared.UsageDiscountFilterParam
