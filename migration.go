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
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// A migration moves a plan's subscriptions onto one of the plan's versions. Each
// subscription is migrated asynchronously, either immediately or at a scheduled
// time, and the migration reports its overall status as it progresses.
//
// MigrationService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMigrationService] method instead.
type MigrationService struct {
	Options []option.RequestOption
	// A migration moves a plan's subscriptions onto one of the plan's versions. Each
	// subscription is migrated asynchronously, either immediately or at a scheduled
	// time, and the migration reports its overall status as it progresses.
	Subscriptions *MigrationSubscriptionService
}

// NewMigrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMigrationService(opts ...option.RequestOption) (r *MigrationService) {
	r = &MigrationService{}
	r.Options = opts
	r.Subscriptions = NewMigrationSubscriptionService(opts...)
	return
}

// This endpoint creates a migration of a plan's subscriptions to one of its
// versions. With `included_subscription_ids` it migrates exactly those
// subscriptions; otherwise it migrates the plan's subscriptions that match the
// filters. Each subscription is migrated asynchronously; fetch the migration to
// follow its status.
func (r *MigrationService) New(ctx context.Context, body MigrationNewParams, opts ...option.RequestOption) (res *MigrationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "migrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint returns a migration identified by its ID.
func (r *MigrationService) Get(ctx context.Context, migrationID string, opts ...option.RequestOption) (res *MigrationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if migrationID == "" {
		err = errors.New("missing required migration_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("migrations/%s", migrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint returns a list of migrations across all plans in the account,
// optionally filtered by plan (`plan_id` or `external_plan_id`). The list of
// migrations is ordered starting from the most recently created migration. The
// response also includes pagination_metadata, which lets the caller retrieve the
// next page of results if they exist.
func (r *MigrationService) List(ctx context.Context, query MigrationListParams, opts ...option.RequestOption) (res *pagination.Page[MigrationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "migrations"
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

// This endpoint returns a list of migrations across all plans in the account,
// optionally filtered by plan (`plan_id` or `external_plan_id`). The list of
// migrations is ordered starting from the most recently created migration. The
// response also includes pagination_metadata, which lets the caller retrieve the
// next page of results if they exist.
func (r *MigrationService) ListAutoPaging(ctx context.Context, query MigrationListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MigrationListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint cancels a migration identified by its ID.
func (r *MigrationService) Cancel(ctx context.Context, migrationID string, opts ...option.RequestOption) (res *MigrationCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if migrationID == "" {
		err = errors.New("missing required migration_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("migrations/%s/cancel", migrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type MigrationSubscription struct {
	Customer shared.CustomerMinified `json:"customer" api:"required"`
	// When the subscription is migrated to the target plan version. Null unless the
	// status is 'scheduled' or 'completed'.
	EffectiveDate time.Time `json:"effective_date" api:"required,nullable" format:"date-time"`
	// Why the subscription could not be migrated. Null unless the status is 'failed'.
	FailureReason string `json:"failure_reason" api:"required,nullable"`
	// The subscription's status in the migration: 'scheduling' until the migration has
	// processed it, 'scheduled' when it will be migrated to the target version at
	// `effective_date`, 'completed' once it has been migrated, 'skipped' when the
	// migration left it unchanged, or 'failed' when it could not be migrated (see
	// `failure_reason`).
	Status MigrationSubscriptionStatus `json:"status" api:"required"`
	// The ID of the subscription.
	SubscriptionID string                    `json:"subscription_id" api:"required"`
	JSON           migrationSubscriptionJSON `json:"-"`
}

// migrationSubscriptionJSON contains the JSON metadata for the struct
// [MigrationSubscription]
type migrationSubscriptionJSON struct {
	Customer       apijson.Field
	EffectiveDate  apijson.Field
	FailureReason  apijson.Field
	Status         apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MigrationSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationSubscriptionJSON) RawJSON() string {
	return r.raw
}

// The subscription's status in the migration: 'scheduling' until the migration has
// processed it, 'scheduled' when it will be migrated to the target version at
// `effective_date`, 'completed' once it has been migrated, 'skipped' when the
// migration left it unchanged, or 'failed' when it could not be migrated (see
// `failure_reason`).
type MigrationSubscriptionStatus string

const (
	MigrationSubscriptionStatusScheduling MigrationSubscriptionStatus = "scheduling"
	MigrationSubscriptionStatusScheduled  MigrationSubscriptionStatus = "scheduled"
	MigrationSubscriptionStatusCompleted  MigrationSubscriptionStatus = "completed"
	MigrationSubscriptionStatusSkipped    MigrationSubscriptionStatus = "skipped"
	MigrationSubscriptionStatusFailed     MigrationSubscriptionStatus = "failed"
)

func (r MigrationSubscriptionStatus) IsKnown() bool {
	switch r {
	case MigrationSubscriptionStatusScheduling, MigrationSubscriptionStatusScheduled, MigrationSubscriptionStatusCompleted, MigrationSubscriptionStatusSkipped, MigrationSubscriptionStatusFailed:
		return true
	}
	return false
}

// A migration moves a plan's subscriptions onto one of the plan's versions. Each
// subscription is migrated asynchronously, either immediately or at a scheduled
// time, and the migration reports its overall status as it progresses.
type MigrationNewResponse struct {
	// Unique identifier for this plan version change.
	ID string `json:"id" api:"required"`
	// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
	// 'end_of_term' when scheduled to be at the start or end of the current billing
	// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
	// start of the current or next invoice.
	EffectiveTime MigrationNewResponseEffectiveTimeUnion `json:"effective_time" api:"required,nullable" format:"date"`
	// When the migration completed or was canceled, or null if it is still in
	// progress.
	EndedAt time.Time `json:"ended_at" api:"required,nullable" format:"date-time"`
	// The ID of the plan being migrated.
	PlanID string `json:"plan_id" api:"required"`
	// When Orb began migrating the plan's subscriptions. Null if the migration has not
	// begun, or was canceled before it began.
	StartedAt time.Time `json:"started_at" api:"required,nullable" format:"date-time"`
	// Current status of the migration: 'not_started', 'in_progress', 'completed',
	// 'action_needed', or 'canceled'.
	Status MigrationNewResponseStatus `json:"status" api:"required"`
	// The plan version subscriptions are migrated to.
	TargetPlanVersionNumber int64                    `json:"target_plan_version_number" api:"required"`
	JSON                    migrationNewResponseJSON `json:"-"`
}

// migrationNewResponseJSON contains the JSON metadata for the struct
// [MigrationNewResponse]
type migrationNewResponseJSON struct {
	ID                      apijson.Field
	EffectiveTime           apijson.Field
	EndedAt                 apijson.Field
	PlanID                  apijson.Field
	StartedAt               apijson.Field
	Status                  apijson.Field
	TargetPlanVersionNumber apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MigrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationNewResponseJSON) RawJSON() string {
	return r.raw
}

// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
// 'end_of_term' when scheduled to be at the start or end of the current billing
// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
// start of the current or next invoice.
//
// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [MigrationNewResponseEffectiveTimeString].
type MigrationNewResponseEffectiveTimeUnion interface {
	ImplementsMigrationNewResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MigrationNewResponseEffectiveTimeUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(MigrationNewResponseEffectiveTimeString("")),
		},
	)
}

type MigrationNewResponseEffectiveTimeString string

const (
	MigrationNewResponseEffectiveTimeStringStartOfTerm    MigrationNewResponseEffectiveTimeString = "start_of_term"
	MigrationNewResponseEffectiveTimeStringEndOfTerm      MigrationNewResponseEffectiveTimeString = "end_of_term"
	MigrationNewResponseEffectiveTimeStringStartOfInvoice MigrationNewResponseEffectiveTimeString = "start_of_invoice"
	MigrationNewResponseEffectiveTimeStringEndOfInvoice   MigrationNewResponseEffectiveTimeString = "end_of_invoice"
)

func (r MigrationNewResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case MigrationNewResponseEffectiveTimeStringStartOfTerm, MigrationNewResponseEffectiveTimeStringEndOfTerm, MigrationNewResponseEffectiveTimeStringStartOfInvoice, MigrationNewResponseEffectiveTimeStringEndOfInvoice:
		return true
	}
	return false
}

func (r MigrationNewResponseEffectiveTimeString) ImplementsMigrationNewResponseEffectiveTimeUnion() {}

// Current status of the migration: 'not_started', 'in_progress', 'completed',
// 'action_needed', or 'canceled'.
type MigrationNewResponseStatus string

const (
	MigrationNewResponseStatusNotStarted   MigrationNewResponseStatus = "not_started"
	MigrationNewResponseStatusInProgress   MigrationNewResponseStatus = "in_progress"
	MigrationNewResponseStatusCompleted    MigrationNewResponseStatus = "completed"
	MigrationNewResponseStatusActionNeeded MigrationNewResponseStatus = "action_needed"
	MigrationNewResponseStatusCanceled     MigrationNewResponseStatus = "canceled"
)

func (r MigrationNewResponseStatus) IsKnown() bool {
	switch r {
	case MigrationNewResponseStatusNotStarted, MigrationNewResponseStatusInProgress, MigrationNewResponseStatusCompleted, MigrationNewResponseStatusActionNeeded, MigrationNewResponseStatusCanceled:
		return true
	}
	return false
}

// A migration moves a plan's subscriptions onto one of the plan's versions. Each
// subscription is migrated asynchronously, either immediately or at a scheduled
// time, and the migration reports its overall status as it progresses.
type MigrationGetResponse struct {
	// Unique identifier for this plan version change.
	ID string `json:"id" api:"required"`
	// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
	// 'end_of_term' when scheduled to be at the start or end of the current billing
	// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
	// start of the current or next invoice.
	EffectiveTime MigrationGetResponseEffectiveTimeUnion `json:"effective_time" api:"required,nullable" format:"date"`
	// When the migration completed or was canceled, or null if it is still in
	// progress.
	EndedAt time.Time `json:"ended_at" api:"required,nullable" format:"date-time"`
	// The ID of the plan being migrated.
	PlanID string `json:"plan_id" api:"required"`
	// When Orb began migrating the plan's subscriptions. Null if the migration has not
	// begun, or was canceled before it began.
	StartedAt time.Time `json:"started_at" api:"required,nullable" format:"date-time"`
	// Current status of the migration: 'not_started', 'in_progress', 'completed',
	// 'action_needed', or 'canceled'.
	Status MigrationGetResponseStatus `json:"status" api:"required"`
	// The plan version subscriptions are migrated to.
	TargetPlanVersionNumber int64                    `json:"target_plan_version_number" api:"required"`
	JSON                    migrationGetResponseJSON `json:"-"`
}

// migrationGetResponseJSON contains the JSON metadata for the struct
// [MigrationGetResponse]
type migrationGetResponseJSON struct {
	ID                      apijson.Field
	EffectiveTime           apijson.Field
	EndedAt                 apijson.Field
	PlanID                  apijson.Field
	StartedAt               apijson.Field
	Status                  apijson.Field
	TargetPlanVersionNumber apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MigrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationGetResponseJSON) RawJSON() string {
	return r.raw
}

// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
// 'end_of_term' when scheduled to be at the start or end of the current billing
// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
// start of the current or next invoice.
//
// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [MigrationGetResponseEffectiveTimeString].
type MigrationGetResponseEffectiveTimeUnion interface {
	ImplementsMigrationGetResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MigrationGetResponseEffectiveTimeUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(MigrationGetResponseEffectiveTimeString("")),
		},
	)
}

type MigrationGetResponseEffectiveTimeString string

const (
	MigrationGetResponseEffectiveTimeStringStartOfTerm    MigrationGetResponseEffectiveTimeString = "start_of_term"
	MigrationGetResponseEffectiveTimeStringEndOfTerm      MigrationGetResponseEffectiveTimeString = "end_of_term"
	MigrationGetResponseEffectiveTimeStringStartOfInvoice MigrationGetResponseEffectiveTimeString = "start_of_invoice"
	MigrationGetResponseEffectiveTimeStringEndOfInvoice   MigrationGetResponseEffectiveTimeString = "end_of_invoice"
)

func (r MigrationGetResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case MigrationGetResponseEffectiveTimeStringStartOfTerm, MigrationGetResponseEffectiveTimeStringEndOfTerm, MigrationGetResponseEffectiveTimeStringStartOfInvoice, MigrationGetResponseEffectiveTimeStringEndOfInvoice:
		return true
	}
	return false
}

func (r MigrationGetResponseEffectiveTimeString) ImplementsMigrationGetResponseEffectiveTimeUnion() {}

// Current status of the migration: 'not_started', 'in_progress', 'completed',
// 'action_needed', or 'canceled'.
type MigrationGetResponseStatus string

const (
	MigrationGetResponseStatusNotStarted   MigrationGetResponseStatus = "not_started"
	MigrationGetResponseStatusInProgress   MigrationGetResponseStatus = "in_progress"
	MigrationGetResponseStatusCompleted    MigrationGetResponseStatus = "completed"
	MigrationGetResponseStatusActionNeeded MigrationGetResponseStatus = "action_needed"
	MigrationGetResponseStatusCanceled     MigrationGetResponseStatus = "canceled"
)

func (r MigrationGetResponseStatus) IsKnown() bool {
	switch r {
	case MigrationGetResponseStatusNotStarted, MigrationGetResponseStatusInProgress, MigrationGetResponseStatusCompleted, MigrationGetResponseStatusActionNeeded, MigrationGetResponseStatusCanceled:
		return true
	}
	return false
}

// A migration moves a plan's subscriptions onto one of the plan's versions. Each
// subscription is migrated asynchronously, either immediately or at a scheduled
// time, and the migration reports its overall status as it progresses.
type MigrationListResponse struct {
	// Unique identifier for this plan version change.
	ID string `json:"id" api:"required"`
	// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
	// 'end_of_term' when scheduled to be at the start or end of the current billing
	// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
	// start of the current or next invoice.
	EffectiveTime MigrationListResponseEffectiveTimeUnion `json:"effective_time" api:"required,nullable" format:"date"`
	// When the migration completed or was canceled, or null if it is still in
	// progress.
	EndedAt time.Time `json:"ended_at" api:"required,nullable" format:"date-time"`
	// The ID of the plan being migrated.
	PlanID string `json:"plan_id" api:"required"`
	// When Orb began migrating the plan's subscriptions. Null if the migration has not
	// begun, or was canceled before it began.
	StartedAt time.Time `json:"started_at" api:"required,nullable" format:"date-time"`
	// Current status of the migration: 'not_started', 'in_progress', 'completed',
	// 'action_needed', or 'canceled'.
	Status MigrationListResponseStatus `json:"status" api:"required"`
	// The plan version subscriptions are migrated to.
	TargetPlanVersionNumber int64                     `json:"target_plan_version_number" api:"required"`
	JSON                    migrationListResponseJSON `json:"-"`
}

// migrationListResponseJSON contains the JSON metadata for the struct
// [MigrationListResponse]
type migrationListResponseJSON struct {
	ID                      apijson.Field
	EffectiveTime           apijson.Field
	EndedAt                 apijson.Field
	PlanID                  apijson.Field
	StartedAt               apijson.Field
	Status                  apijson.Field
	TargetPlanVersionNumber apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MigrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationListResponseJSON) RawJSON() string {
	return r.raw
}

// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
// 'end_of_term' when scheduled to be at the start or end of the current billing
// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
// start of the current or next invoice.
//
// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [MigrationListResponseEffectiveTimeString].
type MigrationListResponseEffectiveTimeUnion interface {
	ImplementsMigrationListResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MigrationListResponseEffectiveTimeUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(MigrationListResponseEffectiveTimeString("")),
		},
	)
}

type MigrationListResponseEffectiveTimeString string

const (
	MigrationListResponseEffectiveTimeStringStartOfTerm    MigrationListResponseEffectiveTimeString = "start_of_term"
	MigrationListResponseEffectiveTimeStringEndOfTerm      MigrationListResponseEffectiveTimeString = "end_of_term"
	MigrationListResponseEffectiveTimeStringStartOfInvoice MigrationListResponseEffectiveTimeString = "start_of_invoice"
	MigrationListResponseEffectiveTimeStringEndOfInvoice   MigrationListResponseEffectiveTimeString = "end_of_invoice"
)

func (r MigrationListResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case MigrationListResponseEffectiveTimeStringStartOfTerm, MigrationListResponseEffectiveTimeStringEndOfTerm, MigrationListResponseEffectiveTimeStringStartOfInvoice, MigrationListResponseEffectiveTimeStringEndOfInvoice:
		return true
	}
	return false
}

func (r MigrationListResponseEffectiveTimeString) ImplementsMigrationListResponseEffectiveTimeUnion() {
}

// Current status of the migration: 'not_started', 'in_progress', 'completed',
// 'action_needed', or 'canceled'.
type MigrationListResponseStatus string

const (
	MigrationListResponseStatusNotStarted   MigrationListResponseStatus = "not_started"
	MigrationListResponseStatusInProgress   MigrationListResponseStatus = "in_progress"
	MigrationListResponseStatusCompleted    MigrationListResponseStatus = "completed"
	MigrationListResponseStatusActionNeeded MigrationListResponseStatus = "action_needed"
	MigrationListResponseStatusCanceled     MigrationListResponseStatus = "canceled"
)

func (r MigrationListResponseStatus) IsKnown() bool {
	switch r {
	case MigrationListResponseStatusNotStarted, MigrationListResponseStatusInProgress, MigrationListResponseStatusCompleted, MigrationListResponseStatusActionNeeded, MigrationListResponseStatusCanceled:
		return true
	}
	return false
}

// A migration moves a plan's subscriptions onto one of the plan's versions. Each
// subscription is migrated asynchronously, either immediately or at a scheduled
// time, and the migration reports its overall status as it progresses.
type MigrationCancelResponse struct {
	// Unique identifier for this plan version change.
	ID string `json:"id" api:"required"`
	// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
	// 'end_of_term' when scheduled to be at the start or end of the current billing
	// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
	// start of the current or next invoice.
	EffectiveTime MigrationCancelResponseEffectiveTimeUnion `json:"effective_time" api:"required,nullable" format:"date"`
	// When the migration completed or was canceled, or null if it is still in
	// progress.
	EndedAt time.Time `json:"ended_at" api:"required,nullable" format:"date-time"`
	// The ID of the plan being migrated.
	PlanID string `json:"plan_id" api:"required"`
	// When Orb began migrating the plan's subscriptions. Null if the migration has not
	// begun, or was canceled before it began.
	StartedAt time.Time `json:"started_at" api:"required,nullable" format:"date-time"`
	// Current status of the migration: 'not_started', 'in_progress', 'completed',
	// 'action_needed', or 'canceled'.
	Status MigrationCancelResponseStatus `json:"status" api:"required"`
	// The plan version subscriptions are migrated to.
	TargetPlanVersionNumber int64                       `json:"target_plan_version_number" api:"required"`
	JSON                    migrationCancelResponseJSON `json:"-"`
}

// migrationCancelResponseJSON contains the JSON metadata for the struct
// [MigrationCancelResponse]
type migrationCancelResponseJSON struct {
	ID                      apijson.Field
	EffectiveTime           apijson.Field
	EndedAt                 apijson.Field
	PlanID                  apijson.Field
	StartedAt               apijson.Field
	Status                  apijson.Field
	TargetPlanVersionNumber apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MigrationCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationCancelResponseJSON) RawJSON() string {
	return r.raw
}

// When the migration takes effect. Can be a specific date/time, 'start_of_term' or
// 'end_of_term' when scheduled to be at the start or end of the current billing
// period, or 'start_of_invoice' or 'end_of_invoice' when scheduled to be at the
// start of the current or next invoice.
//
// Union satisfied by [shared.UnionTime], [shared.UnionTime] or
// [MigrationCancelResponseEffectiveTimeString].
type MigrationCancelResponseEffectiveTimeUnion interface {
	ImplementsMigrationCancelResponseEffectiveTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MigrationCancelResponseEffectiveTimeUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(MigrationCancelResponseEffectiveTimeString("")),
		},
	)
}

type MigrationCancelResponseEffectiveTimeString string

const (
	MigrationCancelResponseEffectiveTimeStringStartOfTerm    MigrationCancelResponseEffectiveTimeString = "start_of_term"
	MigrationCancelResponseEffectiveTimeStringEndOfTerm      MigrationCancelResponseEffectiveTimeString = "end_of_term"
	MigrationCancelResponseEffectiveTimeStringStartOfInvoice MigrationCancelResponseEffectiveTimeString = "start_of_invoice"
	MigrationCancelResponseEffectiveTimeStringEndOfInvoice   MigrationCancelResponseEffectiveTimeString = "end_of_invoice"
)

func (r MigrationCancelResponseEffectiveTimeString) IsKnown() bool {
	switch r {
	case MigrationCancelResponseEffectiveTimeStringStartOfTerm, MigrationCancelResponseEffectiveTimeStringEndOfTerm, MigrationCancelResponseEffectiveTimeStringStartOfInvoice, MigrationCancelResponseEffectiveTimeStringEndOfInvoice:
		return true
	}
	return false
}

func (r MigrationCancelResponseEffectiveTimeString) ImplementsMigrationCancelResponseEffectiveTimeUnion() {
}

// Current status of the migration: 'not_started', 'in_progress', 'completed',
// 'action_needed', or 'canceled'.
type MigrationCancelResponseStatus string

const (
	MigrationCancelResponseStatusNotStarted   MigrationCancelResponseStatus = "not_started"
	MigrationCancelResponseStatusInProgress   MigrationCancelResponseStatus = "in_progress"
	MigrationCancelResponseStatusCompleted    MigrationCancelResponseStatus = "completed"
	MigrationCancelResponseStatusActionNeeded MigrationCancelResponseStatus = "action_needed"
	MigrationCancelResponseStatusCanceled     MigrationCancelResponseStatus = "canceled"
)

func (r MigrationCancelResponseStatus) IsKnown() bool {
	switch r {
	case MigrationCancelResponseStatusNotStarted, MigrationCancelResponseStatusInProgress, MigrationCancelResponseStatusCompleted, MigrationCancelResponseStatusActionNeeded, MigrationCancelResponseStatusCanceled:
		return true
	}
	return false
}

type MigrationNewParams struct {
	// When the migration takes effect: 'immediately', 'requested_date' on
	// `effective_date`, 'start_of_term' or 'end_of_term' for the start or end of each
	// subscription's current billing period, or 'start_of_invoice' or 'end_of_invoice'
	// for the start of its current or next invoice.
	EffectiveTime param.Field[MigrationNewParamsEffectiveTime] `json:"effective_time" api:"required"`
	// The plan version to migrate subscriptions to. It can't be a draft version.
	TargetPlanVersionNumber param.Field[int64] `json:"target_plan_version_number" api:"required"`
	// When false, a subscription whose migration would void an issued invoice or
	// create a credit note fails instead.
	AllowInvoiceCreditOrVoid param.Field[bool] `json:"allow_invoice_credit_or_void"`
	// When true, an in-arrears price the migration ends mid-cycle bills its final line
	// item on the next scheduled invoice instead of on its end date.
	CanDeferBilling param.Field[bool] `json:"can_defer_billing"`
	// The date the migration takes effect, at midnight in each customer's timezone.
	// Required when `effective_time` is `requested_date`, and not allowed otherwise.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// Leave these subscriptions out. Can't be combined with
	// `included_subscription_ids`.
	ExcludedSubscriptionIDs param.Field[[]string] `json:"excluded_subscription_ids"`
	// The external ID of the plan whose subscriptions to migrate.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// `true` migrates only subscriptions in a trial, `false` only subscriptions not in
	// a trial.
	InTrial param.Field[bool] `json:"in_trial"`
	// Migrate only these subscriptions. When set, the filters below are ignored.
	IncludedSubscriptionIDs param.Field[[]string] `json:"included_subscription_ids"`
	// The plan whose subscriptions to migrate. Exactly one of `plan_id` and
	// `external_plan_id` is required.
	PlanID param.Field[string] `json:"plan_id"`
	// Migrate only subscriptions currently on these plan versions.
	PlanVersionNumbers param.Field[[]int64] `json:"plan_version_numbers"`
	// Why the migration was created.
	Reason param.Field[string] `json:"reason"`
	// Migrate subscriptions with these statuses. Defaults to active subscriptions.
	Statuses param.Field[[]MigrationNewParamsStatus] `json:"statuses"`
}

func (r MigrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When the migration takes effect: 'immediately', 'requested_date' on
// `effective_date`, 'start_of_term' or 'end_of_term' for the start or end of each
// subscription's current billing period, or 'start_of_invoice' or 'end_of_invoice'
// for the start of its current or next invoice.
type MigrationNewParamsEffectiveTime string

const (
	MigrationNewParamsEffectiveTimeImmediately    MigrationNewParamsEffectiveTime = "immediately"
	MigrationNewParamsEffectiveTimeRequestedDate  MigrationNewParamsEffectiveTime = "requested_date"
	MigrationNewParamsEffectiveTimeStartOfTerm    MigrationNewParamsEffectiveTime = "start_of_term"
	MigrationNewParamsEffectiveTimeEndOfTerm      MigrationNewParamsEffectiveTime = "end_of_term"
	MigrationNewParamsEffectiveTimeStartOfInvoice MigrationNewParamsEffectiveTime = "start_of_invoice"
	MigrationNewParamsEffectiveTimeEndOfInvoice   MigrationNewParamsEffectiveTime = "end_of_invoice"
)

func (r MigrationNewParamsEffectiveTime) IsKnown() bool {
	switch r {
	case MigrationNewParamsEffectiveTimeImmediately, MigrationNewParamsEffectiveTimeRequestedDate, MigrationNewParamsEffectiveTimeStartOfTerm, MigrationNewParamsEffectiveTimeEndOfTerm, MigrationNewParamsEffectiveTimeStartOfInvoice, MigrationNewParamsEffectiveTimeEndOfInvoice:
		return true
	}
	return false
}

type MigrationNewParamsStatus string

const (
	MigrationNewParamsStatusActive   MigrationNewParamsStatus = "active"
	MigrationNewParamsStatusUpcoming MigrationNewParamsStatus = "upcoming"
)

func (r MigrationNewParamsStatus) IsKnown() bool {
	switch r {
	case MigrationNewParamsStatusActive, MigrationNewParamsStatusUpcoming:
		return true
	}
	return false
}

type MigrationListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor         param.Field[string] `query:"cursor"`
	ExternalPlanID param.Field[string] `query:"external_plan_id"`
	// The number of items to fetch. Defaults to 20.
	Limit  param.Field[int64]  `query:"limit"`
	PlanID param.Field[string] `query:"plan_id"`
}

// URLQuery serializes [MigrationListParams]'s query parameters as `url.Values`.
func (r MigrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
