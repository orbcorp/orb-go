// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
)

// LicenseUsageService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLicenseUsageService] method instead.
type LicenseUsageService struct {
	Options []option.RequestOption
}

// NewLicenseUsageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLicenseUsageService(opts ...option.RequestOption) (r *LicenseUsageService) {
	r = &LicenseUsageService{}
	r.Options = opts
	return
}

// Returns usage and remaining credits for all licenses of a given type on a
// subscription.
//
// Date range defaults to the current billing period if not specified.
func (r *LicenseUsageService) GetAllUsage(ctx context.Context, query LicenseUsageGetAllUsageParams, opts ...option.RequestOption) (res *LicenseUsageGetAllUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "licenses/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns usage and remaining credits for a specific license over a date range.
//
// Date range defaults to the current billing period if not specified.
func (r *LicenseUsageService) GetUsage(ctx context.Context, licenseID string, query LicenseUsageGetUsageParams, opts ...option.RequestOption) (res *LicenseUsageGetUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if licenseID == "" {
		err = errors.New("missing required license_id parameter")
		return
	}
	path := fmt.Sprintf("licenses/%s/usage", licenseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LicenseUsageGetAllUsageResponse struct {
	Data               []LicenseUsageGetAllUsageResponseData `json:"data,required"`
	PaginationMetadata shared.PaginationMetadata             `json:"pagination_metadata,required"`
	JSON               licenseUsageGetAllUsageResponseJSON   `json:"-"`
}

// licenseUsageGetAllUsageResponseJSON contains the JSON metadata for the struct
// [LicenseUsageGetAllUsageResponse]
type licenseUsageGetAllUsageResponseJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LicenseUsageGetAllUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseUsageGetAllUsageResponseJSON) RawJSON() string {
	return r.raw
}

// The LicenseUsage resource represents usage and remaining credits for a license
// over a date range.
//
// When grouped by 'day' only, license_id and external_license_id will be null as
// the data is aggregated across all licenses.
type LicenseUsageGetAllUsageResponseData struct {
	// The total credits allocated to this license for the period.
	AllocatedCredits float64 `json:"allocated_credits,required"`
	// The credits consumed by this license for the period.
	ConsumedCredits float64 `json:"consumed_credits,required"`
	// The end date of the usage period.
	EndDate time.Time `json:"end_date,required" format:"date"`
	// The unique identifier for the license type.
	LicenseTypeID string `json:"license_type_id,required"`
	// The pricing unit for the credits (e.g., 'credits').
	PricingUnit string `json:"pricing_unit,required"`
	// The remaining credits available for this license (allocated - consumed).
	RemainingCredits float64 `json:"remaining_credits,required"`
	// The start date of the usage period.
	StartDate time.Time `json:"start_date,required" format:"date"`
	// The unique identifier for the subscription.
	SubscriptionID string `json:"subscription_id,required"`
	// Credits consumed while the license was active (eligible for individual
	// allocation deduction).
	AllocationEligibleCredits float64 `json:"allocation_eligible_credits,nullable"`
	// The external identifier for the license. Null when grouped by day only.
	ExternalLicenseID string `json:"external_license_id,nullable"`
	// The unique identifier for the license. Null when grouped by day only.
	LicenseID string `json:"license_id,nullable"`
	// Credits consumed while the license was inactive (draws from shared pool, not
	// individual allocation).
	SharedPoolCredits float64                                 `json:"shared_pool_credits,nullable"`
	JSON              licenseUsageGetAllUsageResponseDataJSON `json:"-"`
}

// licenseUsageGetAllUsageResponseDataJSON contains the JSON metadata for the
// struct [LicenseUsageGetAllUsageResponseData]
type licenseUsageGetAllUsageResponseDataJSON struct {
	AllocatedCredits          apijson.Field
	ConsumedCredits           apijson.Field
	EndDate                   apijson.Field
	LicenseTypeID             apijson.Field
	PricingUnit               apijson.Field
	RemainingCredits          apijson.Field
	StartDate                 apijson.Field
	SubscriptionID            apijson.Field
	AllocationEligibleCredits apijson.Field
	ExternalLicenseID         apijson.Field
	LicenseID                 apijson.Field
	SharedPoolCredits         apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LicenseUsageGetAllUsageResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseUsageGetAllUsageResponseDataJSON) RawJSON() string {
	return r.raw
}

type LicenseUsageGetUsageResponse struct {
	Data               []LicenseUsageGetUsageResponseData `json:"data,required"`
	PaginationMetadata shared.PaginationMetadata          `json:"pagination_metadata,required"`
	JSON               licenseUsageGetUsageResponseJSON   `json:"-"`
}

// licenseUsageGetUsageResponseJSON contains the JSON metadata for the struct
// [LicenseUsageGetUsageResponse]
type licenseUsageGetUsageResponseJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LicenseUsageGetUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseUsageGetUsageResponseJSON) RawJSON() string {
	return r.raw
}

// The LicenseUsage resource represents usage and remaining credits for a license
// over a date range.
//
// When grouped by 'day' only, license_id and external_license_id will be null as
// the data is aggregated across all licenses.
type LicenseUsageGetUsageResponseData struct {
	// The total credits allocated to this license for the period.
	AllocatedCredits float64 `json:"allocated_credits,required"`
	// The credits consumed by this license for the period.
	ConsumedCredits float64 `json:"consumed_credits,required"`
	// The end date of the usage period.
	EndDate time.Time `json:"end_date,required" format:"date"`
	// The unique identifier for the license type.
	LicenseTypeID string `json:"license_type_id,required"`
	// The pricing unit for the credits (e.g., 'credits').
	PricingUnit string `json:"pricing_unit,required"`
	// The remaining credits available for this license (allocated - consumed).
	RemainingCredits float64 `json:"remaining_credits,required"`
	// The start date of the usage period.
	StartDate time.Time `json:"start_date,required" format:"date"`
	// The unique identifier for the subscription.
	SubscriptionID string `json:"subscription_id,required"`
	// Credits consumed while the license was active (eligible for individual
	// allocation deduction).
	AllocationEligibleCredits float64 `json:"allocation_eligible_credits,nullable"`
	// The external identifier for the license. Null when grouped by day only.
	ExternalLicenseID string `json:"external_license_id,nullable"`
	// The unique identifier for the license. Null when grouped by day only.
	LicenseID string `json:"license_id,nullable"`
	// Credits consumed while the license was inactive (draws from shared pool, not
	// individual allocation).
	SharedPoolCredits float64                              `json:"shared_pool_credits,nullable"`
	JSON              licenseUsageGetUsageResponseDataJSON `json:"-"`
}

// licenseUsageGetUsageResponseDataJSON contains the JSON metadata for the struct
// [LicenseUsageGetUsageResponseData]
type licenseUsageGetUsageResponseDataJSON struct {
	AllocatedCredits          apijson.Field
	ConsumedCredits           apijson.Field
	EndDate                   apijson.Field
	LicenseTypeID             apijson.Field
	PricingUnit               apijson.Field
	RemainingCredits          apijson.Field
	StartDate                 apijson.Field
	SubscriptionID            apijson.Field
	AllocationEligibleCredits apijson.Field
	ExternalLicenseID         apijson.Field
	LicenseID                 apijson.Field
	SharedPoolCredits         apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LicenseUsageGetUsageResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseUsageGetUsageResponseDataJSON) RawJSON() string {
	return r.raw
}

type LicenseUsageGetAllUsageParams struct {
	// The license type ID to filter licenses by.
	LicenseTypeID param.Field[string] `query:"license_type_id,required"`
	// The subscription ID to get license usage for.
	SubscriptionID param.Field[string] `query:"subscription_id,required"`
	// Pagination cursor from a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// End date for the usage period (YYYY-MM-DD). Defaults to end of current billing
	// period.
	EndDate param.Field[time.Time] `query:"end_date" format:"date"`
	// How to group the results. Valid values: 'license', 'day'. Can be combined (e.g.,
	// 'license,day').
	GroupBy param.Field[[]string] `query:"group_by"`
	// Maximum number of rows in the response data (default 20, max 100).
	Limit param.Field[int64] `query:"limit"`
	// Start date for the usage period (YYYY-MM-DD). Defaults to start of current
	// billing period.
	StartDate param.Field[time.Time] `query:"start_date" format:"date"`
}

// URLQuery serializes [LicenseUsageGetAllUsageParams]'s query parameters as
// `url.Values`.
func (r LicenseUsageGetAllUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LicenseUsageGetUsageParams struct {
	// Pagination cursor from a previous request.
	Cursor param.Field[string] `query:"cursor"`
	// End date for the usage period (YYYY-MM-DD). Defaults to end of current billing
	// period.
	EndDate param.Field[time.Time] `query:"end_date" format:"date"`
	// How to group the results. Valid values: 'license', 'day'. Can be combined (e.g.,
	// 'license,day').
	GroupBy param.Field[[]string] `query:"group_by"`
	// Maximum number of rows in the response data (default 20, max 100).
	Limit param.Field[int64] `query:"limit"`
	// Start date for the usage period (YYYY-MM-DD). Defaults to start of current
	// billing period.
	StartDate param.Field[time.Time] `query:"start_date" format:"date"`
}

// URLQuery serializes [LicenseUsageGetUsageParams]'s query parameters as
// `url.Values`.
func (r LicenseUsageGetUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
