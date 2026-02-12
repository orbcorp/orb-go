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
	"github.com/orbcorp/orb-go/packages/pagination"
)

// LicenseService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLicenseService] method instead.
type LicenseService struct {
	Options          []option.RequestOption
	ExternalLicenses *LicenseExternalLicenseService
	Usage            *LicenseUsageService
}

// NewLicenseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLicenseService(opts ...option.RequestOption) (r *LicenseService) {
	r = &LicenseService{}
	r.Options = opts
	r.ExternalLicenses = NewLicenseExternalLicenseService(opts...)
	r.Usage = NewLicenseUsageService(opts...)
	return
}

// This endpoint is used to create a new license for a user.
//
// If a start date is provided, the license will be activated at the **start** of
// the specified date in the customer's timezone. Otherwise, the activation time
// will default to the **start** of the current day in the customer's timezone.
func (r *LicenseService) New(ctx context.Context, body LicenseNewParams, opts ...option.RequestOption) (res *LicenseNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "licenses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch a license given an identifier.
func (r *LicenseService) Get(ctx context.Context, licenseID string, opts ...option.RequestOption) (res *LicenseGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if licenseID == "" {
		err = errors.New("missing required license_id parameter")
		return
	}
	path := fmt.Sprintf("licenses/%s", licenseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint returns a list of all licenses for a subscription.
func (r *LicenseService) List(ctx context.Context, query LicenseListParams, opts ...option.RequestOption) (res *pagination.Page[LicenseListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "licenses"
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

// This endpoint returns a list of all licenses for a subscription.
func (r *LicenseService) ListAutoPaging(ctx context.Context, query LicenseListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LicenseListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to deactivate an existing license.
//
// If an end date is provided, the license will be deactivated at the **start** of
// the specified date in the customer's timezone. Otherwise, the deactivation time
// will default to the **end** of the current day in the customer's timezone.
func (r *LicenseService) Deactivate(ctx context.Context, licenseID string, body LicenseDeactivateParams, opts ...option.RequestOption) (res *LicenseDeactivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if licenseID == "" {
		err = errors.New("missing required license_id parameter")
		return
	}
	path := fmt.Sprintf("licenses/%s/deactivate", licenseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch a license given an external license identifier.
func (r *LicenseService) GetByExternalID(ctx context.Context, externalLicenseID string, query LicenseGetByExternalIDParams, opts ...option.RequestOption) (res *LicenseGetByExternalIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalLicenseID == "" {
		err = errors.New("missing required external_license_id parameter")
		return
	}
	path := fmt.Sprintf("licenses/external_license_id/%s", externalLicenseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LicenseNewResponse struct {
	ID                string                   `json:"id,required"`
	EndDate           time.Time                `json:"end_date,required,nullable" format:"date-time"`
	ExternalLicenseID string                   `json:"external_license_id,required"`
	LicenseTypeID     string                   `json:"license_type_id,required"`
	StartDate         time.Time                `json:"start_date,required" format:"date-time"`
	Status            LicenseNewResponseStatus `json:"status,required"`
	SubscriptionID    string                   `json:"subscription_id,required"`
	JSON              licenseNewResponseJSON   `json:"-"`
}

// licenseNewResponseJSON contains the JSON metadata for the struct
// [LicenseNewResponse]
type licenseNewResponseJSON struct {
	ID                apijson.Field
	EndDate           apijson.Field
	ExternalLicenseID apijson.Field
	LicenseTypeID     apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LicenseNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseNewResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseNewResponseStatus string

const (
	LicenseNewResponseStatusActive   LicenseNewResponseStatus = "active"
	LicenseNewResponseStatusInactive LicenseNewResponseStatus = "inactive"
)

func (r LicenseNewResponseStatus) IsKnown() bool {
	switch r {
	case LicenseNewResponseStatusActive, LicenseNewResponseStatusInactive:
		return true
	}
	return false
}

type LicenseGetResponse struct {
	ID                string                   `json:"id,required"`
	EndDate           time.Time                `json:"end_date,required,nullable" format:"date-time"`
	ExternalLicenseID string                   `json:"external_license_id,required"`
	LicenseTypeID     string                   `json:"license_type_id,required"`
	StartDate         time.Time                `json:"start_date,required" format:"date-time"`
	Status            LicenseGetResponseStatus `json:"status,required"`
	SubscriptionID    string                   `json:"subscription_id,required"`
	JSON              licenseGetResponseJSON   `json:"-"`
}

// licenseGetResponseJSON contains the JSON metadata for the struct
// [LicenseGetResponse]
type licenseGetResponseJSON struct {
	ID                apijson.Field
	EndDate           apijson.Field
	ExternalLicenseID apijson.Field
	LicenseTypeID     apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LicenseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseGetResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseGetResponseStatus string

const (
	LicenseGetResponseStatusActive   LicenseGetResponseStatus = "active"
	LicenseGetResponseStatusInactive LicenseGetResponseStatus = "inactive"
)

func (r LicenseGetResponseStatus) IsKnown() bool {
	switch r {
	case LicenseGetResponseStatusActive, LicenseGetResponseStatusInactive:
		return true
	}
	return false
}

type LicenseListResponse struct {
	ID                string                    `json:"id,required"`
	EndDate           time.Time                 `json:"end_date,required,nullable" format:"date-time"`
	ExternalLicenseID string                    `json:"external_license_id,required"`
	LicenseTypeID     string                    `json:"license_type_id,required"`
	StartDate         time.Time                 `json:"start_date,required" format:"date-time"`
	Status            LicenseListResponseStatus `json:"status,required"`
	SubscriptionID    string                    `json:"subscription_id,required"`
	JSON              licenseListResponseJSON   `json:"-"`
}

// licenseListResponseJSON contains the JSON metadata for the struct
// [LicenseListResponse]
type licenseListResponseJSON struct {
	ID                apijson.Field
	EndDate           apijson.Field
	ExternalLicenseID apijson.Field
	LicenseTypeID     apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LicenseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseListResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseListResponseStatus string

const (
	LicenseListResponseStatusActive   LicenseListResponseStatus = "active"
	LicenseListResponseStatusInactive LicenseListResponseStatus = "inactive"
)

func (r LicenseListResponseStatus) IsKnown() bool {
	switch r {
	case LicenseListResponseStatusActive, LicenseListResponseStatusInactive:
		return true
	}
	return false
}

type LicenseDeactivateResponse struct {
	ID                string                          `json:"id,required"`
	EndDate           time.Time                       `json:"end_date,required,nullable" format:"date-time"`
	ExternalLicenseID string                          `json:"external_license_id,required"`
	LicenseTypeID     string                          `json:"license_type_id,required"`
	StartDate         time.Time                       `json:"start_date,required" format:"date-time"`
	Status            LicenseDeactivateResponseStatus `json:"status,required"`
	SubscriptionID    string                          `json:"subscription_id,required"`
	JSON              licenseDeactivateResponseJSON   `json:"-"`
}

// licenseDeactivateResponseJSON contains the JSON metadata for the struct
// [LicenseDeactivateResponse]
type licenseDeactivateResponseJSON struct {
	ID                apijson.Field
	EndDate           apijson.Field
	ExternalLicenseID apijson.Field
	LicenseTypeID     apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LicenseDeactivateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseDeactivateResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseDeactivateResponseStatus string

const (
	LicenseDeactivateResponseStatusActive   LicenseDeactivateResponseStatus = "active"
	LicenseDeactivateResponseStatusInactive LicenseDeactivateResponseStatus = "inactive"
)

func (r LicenseDeactivateResponseStatus) IsKnown() bool {
	switch r {
	case LicenseDeactivateResponseStatusActive, LicenseDeactivateResponseStatusInactive:
		return true
	}
	return false
}

type LicenseGetByExternalIDResponse struct {
	ID                string                               `json:"id,required"`
	EndDate           time.Time                            `json:"end_date,required,nullable" format:"date-time"`
	ExternalLicenseID string                               `json:"external_license_id,required"`
	LicenseTypeID     string                               `json:"license_type_id,required"`
	StartDate         time.Time                            `json:"start_date,required" format:"date-time"`
	Status            LicenseGetByExternalIDResponseStatus `json:"status,required"`
	SubscriptionID    string                               `json:"subscription_id,required"`
	JSON              licenseGetByExternalIDResponseJSON   `json:"-"`
}

// licenseGetByExternalIDResponseJSON contains the JSON metadata for the struct
// [LicenseGetByExternalIDResponse]
type licenseGetByExternalIDResponseJSON struct {
	ID                apijson.Field
	EndDate           apijson.Field
	ExternalLicenseID apijson.Field
	LicenseTypeID     apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LicenseGetByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseGetByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseGetByExternalIDResponseStatus string

const (
	LicenseGetByExternalIDResponseStatusActive   LicenseGetByExternalIDResponseStatus = "active"
	LicenseGetByExternalIDResponseStatusInactive LicenseGetByExternalIDResponseStatus = "inactive"
)

func (r LicenseGetByExternalIDResponseStatus) IsKnown() bool {
	switch r {
	case LicenseGetByExternalIDResponseStatusActive, LicenseGetByExternalIDResponseStatusInactive:
		return true
	}
	return false
}

type LicenseNewParams struct {
	// The external identifier for the license.
	ExternalLicenseID param.Field[string] `json:"external_license_id,required"`
	LicenseTypeID     param.Field[string] `json:"license_type_id,required"`
	SubscriptionID    param.Field[string] `json:"subscription_id,required"`
	// The end date of the license. If not provided, the license will remain active
	// until deactivated.
	EndDate param.Field[time.Time] `json:"end_date" format:"date"`
	// The start date of the license. If not provided, defaults to start of day today
	// in the customer's timezone.
	StartDate param.Field[time.Time] `json:"start_date" format:"date"`
}

func (r LicenseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseListParams struct {
	SubscriptionID param.Field[string] `query:"subscription_id,required"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor            param.Field[string] `query:"cursor"`
	ExternalLicenseID param.Field[string] `query:"external_license_id"`
	LicenseTypeID     param.Field[string] `query:"license_type_id"`
	// The number of items to fetch. Defaults to 20.
	Limit  param.Field[int64]                   `query:"limit"`
	Status param.Field[LicenseListParamsStatus] `query:"status"`
}

// URLQuery serializes [LicenseListParams]'s query parameters as `url.Values`.
func (r LicenseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LicenseListParamsStatus string

const (
	LicenseListParamsStatusActive   LicenseListParamsStatus = "active"
	LicenseListParamsStatusInactive LicenseListParamsStatus = "inactive"
)

func (r LicenseListParamsStatus) IsKnown() bool {
	switch r {
	case LicenseListParamsStatusActive, LicenseListParamsStatusInactive:
		return true
	}
	return false
}

type LicenseDeactivateParams struct {
	// The date to deactivate the license. If not provided, defaults to end of day
	// today in the customer's timezone.
	EndDate param.Field[time.Time] `json:"end_date" format:"date"`
}

func (r LicenseDeactivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseGetByExternalIDParams struct {
	// The ID of the license type to fetch the license for.
	LicenseTypeID param.Field[string] `query:"license_type_id,required"`
	// The ID of the subscription to fetch the license for.
	SubscriptionID param.Field[string] `query:"subscription_id,required"`
}

// URLQuery serializes [LicenseGetByExternalIDParams]'s query parameters as
// `url.Values`.
func (r LicenseGetByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
