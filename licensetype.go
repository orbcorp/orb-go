// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
)

// LicenseTypeService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLicenseTypeService] method instead.
type LicenseTypeService struct {
	Options []option.RequestOption
}

// NewLicenseTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLicenseTypeService(opts ...option.RequestOption) (r *LicenseTypeService) {
	r = &LicenseTypeService{}
	r.Options = opts
	return
}

// This endpoint is used to create a new license type.
//
// License types are used to group licenses and define billing behavior. Each
// license type has a name and a grouping key that determines how metrics are
// aggregated for billing purposes.
func (r *LicenseTypeService) New(ctx context.Context, body LicenseTypeNewParams, opts ...option.RequestOption) (res *LicenseTypeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "license_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint returns a license type identified by its license_type_id.
//
// Use this endpoint to retrieve details about a specific license type, including
// its name and grouping key.
func (r *LicenseTypeService) Get(ctx context.Context, licenseTypeID string, opts ...option.RequestOption) (res *LicenseTypeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if licenseTypeID == "" {
		err = errors.New("missing required license_type_id parameter")
		return
	}
	path := fmt.Sprintf("license_types/%s", licenseTypeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint returns a list of all license types configured for the account,
// ordered in ascending order by creation time.
//
// License types are used to group licenses and define billing behavior. Each
// license type has a name and a grouping key that determines how metrics are
// aggregated for billing purposes.
func (r *LicenseTypeService) List(ctx context.Context, query LicenseTypeListParams, opts ...option.RequestOption) (res *pagination.Page[LicenseTypeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "license_types"
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

// This endpoint returns a list of all license types configured for the account,
// ordered in ascending order by creation time.
//
// License types are used to group licenses and define billing behavior. Each
// license type has a name and a grouping key that determines how metrics are
// aggregated for billing purposes.
func (r *LicenseTypeService) ListAutoPaging(ctx context.Context, query LicenseTypeListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LicenseTypeListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
type LicenseTypeNewResponse struct {
	// The Orb-assigned unique identifier for the license type.
	ID string `json:"id,required"`
	// The key used for grouping licenses of this type. This is typically a user
	// identifier field.
	GroupingKey string `json:"grouping_key,required"`
	// The name of the license type.
	Name string                     `json:"name,required"`
	JSON licenseTypeNewResponseJSON `json:"-"`
}

// licenseTypeNewResponseJSON contains the JSON metadata for the struct
// [LicenseTypeNewResponse]
type licenseTypeNewResponseJSON struct {
	ID          apijson.Field
	GroupingKey apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseTypeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseTypeNewResponseJSON) RawJSON() string {
	return r.raw
}

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
type LicenseTypeGetResponse struct {
	// The Orb-assigned unique identifier for the license type.
	ID string `json:"id,required"`
	// The key used for grouping licenses of this type. This is typically a user
	// identifier field.
	GroupingKey string `json:"grouping_key,required"`
	// The name of the license type.
	Name string                     `json:"name,required"`
	JSON licenseTypeGetResponseJSON `json:"-"`
}

// licenseTypeGetResponseJSON contains the JSON metadata for the struct
// [LicenseTypeGetResponse]
type licenseTypeGetResponseJSON struct {
	ID          apijson.Field
	GroupingKey apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseTypeGetResponseJSON) RawJSON() string {
	return r.raw
}

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
type LicenseTypeListResponse struct {
	// The Orb-assigned unique identifier for the license type.
	ID string `json:"id,required"`
	// The key used for grouping licenses of this type. This is typically a user
	// identifier field.
	GroupingKey string `json:"grouping_key,required"`
	// The name of the license type.
	Name string                      `json:"name,required"`
	JSON licenseTypeListResponseJSON `json:"-"`
}

// licenseTypeListResponseJSON contains the JSON metadata for the struct
// [LicenseTypeListResponse]
type licenseTypeListResponseJSON struct {
	ID          apijson.Field
	GroupingKey apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseTypeListResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseTypeNewParams struct {
	// The key used for grouping licenses of this type. This is typically a user
	// identifier field.
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The name of the license type.
	Name param.Field[string] `json:"name,required"`
}

func (r LicenseTypeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseTypeListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [LicenseTypeListParams]'s query parameters as `url.Values`.
func (r LicenseTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
