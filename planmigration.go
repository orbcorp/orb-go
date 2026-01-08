// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// PlanMigrationService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanMigrationService] method instead.
type PlanMigrationService struct {
	Options []option.RequestOption
}

// NewPlanMigrationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlanMigrationService(opts ...option.RequestOption) (r *PlanMigrationService) {
	r = &PlanMigrationService{}
	r.Options = opts
	return
}

// Fetch migration
func (r *PlanMigrationService) Get(ctx context.Context, planID string, migrationID string, opts ...option.RequestOption) (res *PlanMigrationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	if migrationID == "" {
		err = errors.New("missing required migration_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/migrations/%s", planID, migrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint returns a list of all migrations for a plan. The list of
// migrations is ordered starting from the most recently created migration. The
// response also includes pagination_metadata, which lets the caller retrieve the
// next page of results if they exist.
func (r *PlanMigrationService) List(ctx context.Context, planID string, query PlanMigrationListParams, opts ...option.RequestOption) (res *pagination.Page[PlanMigrationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/migrations", planID)
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

// This endpoint returns a list of all migrations for a plan. The list of
// migrations is ordered starting from the most recently created migration. The
// response also includes pagination_metadata, which lets the caller retrieve the
// next page of results if they exist.
func (r *PlanMigrationService) ListAutoPaging(ctx context.Context, planID string, query PlanMigrationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PlanMigrationListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, planID, query, opts...))
}

// This endpoint cancels a migration.
func (r *PlanMigrationService) Cancel(ctx context.Context, planID string, migrationID string, opts ...option.RequestOption) (res *PlanMigrationCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	if migrationID == "" {
		err = errors.New("missing required migration_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/migrations/%s/cancel", planID, migrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type PlanMigrationGetResponse struct {
	ID            string                                     `json:"id,required"`
	EffectiveTime PlanMigrationGetResponseEffectiveTimeUnion `json:"effective_time,required,nullable" format:"date"`
	PlanID        string                                     `json:"plan_id,required"`
	Status        PlanMigrationGetResponseStatus             `json:"status,required"`
	JSON          planMigrationGetResponseJSON               `json:"-"`
}

// planMigrationGetResponseJSON contains the JSON metadata for the struct
// [PlanMigrationGetResponse]
type planMigrationGetResponseJSON struct {
	ID            apijson.Field
	EffectiveTime apijson.Field
	PlanID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PlanMigrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planMigrationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [PlanMigrationGetResponseEffectiveTimeString].
type PlanMigrationGetResponseEffectiveTimeUnion interface {
	ImplementsPlanMigrationGetResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanMigrationGetResponseEffectiveTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PlanMigrationGetResponseEffectiveTimeString("")),
		},
	)
}

type PlanMigrationGetResponseEffectiveTimeString string

const (
	PlanMigrationGetResponseEffectiveTimeStringEndOfTerm PlanMigrationGetResponseEffectiveTimeString = "end_of_term"
)

func (r PlanMigrationGetResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case PlanMigrationGetResponseEffectiveTimeStringEndOfTerm:
		return true
	}
	return false
}

func (r PlanMigrationGetResponseEffectiveTimeString) ImplementsPlanMigrationGetResponseEffectiveTimeUnion() {
}

type PlanMigrationGetResponseStatus string

const (
	PlanMigrationGetResponseStatusNotStarted   PlanMigrationGetResponseStatus = "not_started"
	PlanMigrationGetResponseStatusInProgress   PlanMigrationGetResponseStatus = "in_progress"
	PlanMigrationGetResponseStatusCompleted    PlanMigrationGetResponseStatus = "completed"
	PlanMigrationGetResponseStatusActionNeeded PlanMigrationGetResponseStatus = "action_needed"
	PlanMigrationGetResponseStatusCanceled     PlanMigrationGetResponseStatus = "canceled"
)

func (r PlanMigrationGetResponseStatus) IsKnown() bool {
	switch r {
	case PlanMigrationGetResponseStatusNotStarted, PlanMigrationGetResponseStatusInProgress, PlanMigrationGetResponseStatusCompleted, PlanMigrationGetResponseStatusActionNeeded, PlanMigrationGetResponseStatusCanceled:
		return true
	}
	return false
}

type PlanMigrationListResponse struct {
	ID            string                                      `json:"id,required"`
	EffectiveTime PlanMigrationListResponseEffectiveTimeUnion `json:"effective_time,required,nullable" format:"date"`
	PlanID        string                                      `json:"plan_id,required"`
	Status        PlanMigrationListResponseStatus             `json:"status,required"`
	JSON          planMigrationListResponseJSON               `json:"-"`
}

// planMigrationListResponseJSON contains the JSON metadata for the struct
// [PlanMigrationListResponse]
type planMigrationListResponseJSON struct {
	ID            apijson.Field
	EffectiveTime apijson.Field
	PlanID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PlanMigrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planMigrationListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [PlanMigrationListResponseEffectiveTimeString].
type PlanMigrationListResponseEffectiveTimeUnion interface {
	ImplementsPlanMigrationListResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanMigrationListResponseEffectiveTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PlanMigrationListResponseEffectiveTimeString("")),
		},
	)
}

type PlanMigrationListResponseEffectiveTimeString string

const (
	PlanMigrationListResponseEffectiveTimeStringEndOfTerm PlanMigrationListResponseEffectiveTimeString = "end_of_term"
)

func (r PlanMigrationListResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case PlanMigrationListResponseEffectiveTimeStringEndOfTerm:
		return true
	}
	return false
}

func (r PlanMigrationListResponseEffectiveTimeString) ImplementsPlanMigrationListResponseEffectiveTimeUnion() {
}

type PlanMigrationListResponseStatus string

const (
	PlanMigrationListResponseStatusNotStarted   PlanMigrationListResponseStatus = "not_started"
	PlanMigrationListResponseStatusInProgress   PlanMigrationListResponseStatus = "in_progress"
	PlanMigrationListResponseStatusCompleted    PlanMigrationListResponseStatus = "completed"
	PlanMigrationListResponseStatusActionNeeded PlanMigrationListResponseStatus = "action_needed"
	PlanMigrationListResponseStatusCanceled     PlanMigrationListResponseStatus = "canceled"
)

func (r PlanMigrationListResponseStatus) IsKnown() bool {
	switch r {
	case PlanMigrationListResponseStatusNotStarted, PlanMigrationListResponseStatusInProgress, PlanMigrationListResponseStatusCompleted, PlanMigrationListResponseStatusActionNeeded, PlanMigrationListResponseStatusCanceled:
		return true
	}
	return false
}

type PlanMigrationCancelResponse struct {
	ID            string                                        `json:"id,required"`
	EffectiveTime PlanMigrationCancelResponseEffectiveTimeUnion `json:"effective_time,required,nullable" format:"date"`
	PlanID        string                                        `json:"plan_id,required"`
	Status        PlanMigrationCancelResponseStatus             `json:"status,required"`
	JSON          planMigrationCancelResponseJSON               `json:"-"`
}

// planMigrationCancelResponseJSON contains the JSON metadata for the struct
// [PlanMigrationCancelResponse]
type planMigrationCancelResponseJSON struct {
	ID            apijson.Field
	EffectiveTime apijson.Field
	PlanID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PlanMigrationCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planMigrationCancelResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [PlanMigrationCancelResponseEffectiveTimeString].
type PlanMigrationCancelResponseEffectiveTimeUnion interface {
	ImplementsPlanMigrationCancelResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanMigrationCancelResponseEffectiveTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PlanMigrationCancelResponseEffectiveTimeString("")),
		},
	)
}

type PlanMigrationCancelResponseEffectiveTimeString string

const (
	PlanMigrationCancelResponseEffectiveTimeStringEndOfTerm PlanMigrationCancelResponseEffectiveTimeString = "end_of_term"
)

func (r PlanMigrationCancelResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case PlanMigrationCancelResponseEffectiveTimeStringEndOfTerm:
		return true
	}
	return false
}

func (r PlanMigrationCancelResponseEffectiveTimeString) ImplementsPlanMigrationCancelResponseEffectiveTimeUnion() {
}

type PlanMigrationCancelResponseStatus string

const (
	PlanMigrationCancelResponseStatusNotStarted   PlanMigrationCancelResponseStatus = "not_started"
	PlanMigrationCancelResponseStatusInProgress   PlanMigrationCancelResponseStatus = "in_progress"
	PlanMigrationCancelResponseStatusCompleted    PlanMigrationCancelResponseStatus = "completed"
	PlanMigrationCancelResponseStatusActionNeeded PlanMigrationCancelResponseStatus = "action_needed"
	PlanMigrationCancelResponseStatusCanceled     PlanMigrationCancelResponseStatus = "canceled"
)

func (r PlanMigrationCancelResponseStatus) IsKnown() bool {
	switch r {
	case PlanMigrationCancelResponseStatusNotStarted, PlanMigrationCancelResponseStatusInProgress, PlanMigrationCancelResponseStatusCompleted, PlanMigrationCancelResponseStatusActionNeeded, PlanMigrationCancelResponseStatusCanceled:
		return true
	}
	return false
}

type PlanMigrationListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PlanMigrationListParams]'s query parameters as
// `url.Values`.
func (r PlanMigrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
