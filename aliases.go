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
type AmountDiscountParam = shared.AmountDiscountParam

// This is an alias to an internal type.
type BillingCycleRelativeDate = shared.BillingCycleRelativeDate

// This is an alias to an internal value.
const BillingCycleRelativeDateStartOfTerm = shared.BillingCycleRelativeDateStartOfTerm

// This is an alias to an internal value.
const BillingCycleRelativeDateEndOfTerm = shared.BillingCycleRelativeDateEndOfTerm

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
type PercentageDiscountParam = shared.PercentageDiscountParam

// This is an alias to an internal type.
type TrialDiscount = shared.TrialDiscount

// This is an alias to an internal type.
type TrialDiscountDiscountType = shared.TrialDiscountDiscountType

// This is an alias to an internal value.
const TrialDiscountDiscountTypeTrial = shared.TrialDiscountDiscountTypeTrial

// This is an alias to an internal type.
type TrialDiscountParam = shared.TrialDiscountParam
