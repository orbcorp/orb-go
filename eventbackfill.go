// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// EventBackfillService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventBackfillService] method instead.
type EventBackfillService struct {
	Options []option.RequestOption
}

// NewEventBackfillService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventBackfillService(opts ...option.RequestOption) (r *EventBackfillService) {
	r = &EventBackfillService{}
	r.Options = opts
	return
}

// Creating the backfill enables adding or replacing past events, even those that
// are older than the ingestion grace period. Performing a backfill in Orb involves
// 3 steps:
//
//  1. Create the backfill, specifying its parameters.
//  2. [Ingest](ingest) usage events, referencing the backfill (query parameter
//     `backfill_id`).
//  3. [Close](close-backfill) the backfill, propagating the update in past usage
//     throughout Orb.
//
// Changes from a backfill are not reflected until the backfill is closed, so you
// won’t need to worry about your customers seeing partially updated usage data.
// Backfills are also reversible, so you’ll be able to revert a backfill if you’ve
// made a mistake.
//
// This endpoint will return a backfill object, which contains an `id`. That `id`
// can then be used as the `backfill_id` query parameter to the event ingestion
// endpoint to associate ingested events with this backfill. The effects (e.g.
// updated usage graphs) of this backfill will not take place until the backfill is
// closed.
//
// If the `replace_existing_events` is `true`, existing events in the backfill's
// timeframe will be replaced with the newly ingested events associated with the
// backfill. If `false`, newly ingested events will be added to the existing
// events.
//
// If a `customer_id` or `external_customer_id` is specified, the backfill will
// only affect events for that customer. If neither is specified, the backfill will
// affect all customers.
func (r *EventBackfillService) New(ctx context.Context, body EventBackfillNewParams, opts ...option.RequestOption) (res *EventBackfillNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "events/backfills"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all backfills in a list format.
//
// The list of backfills is ordered starting from the most recently created
// backfill. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist. More information about pagination can be
// found in the [Pagination-metadata schema](pagination).
func (r *EventBackfillService) List(ctx context.Context, query EventBackfillListParams, opts ...option.RequestOption) (res *pagination.Page[EventBackfillListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events/backfills"
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

// This endpoint returns a list of all backfills in a list format.
//
// The list of backfills is ordered starting from the most recently created
// backfill. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist. More information about pagination can be
// found in the [Pagination-metadata schema](pagination).
func (r *EventBackfillService) ListAutoPaging(ctx context.Context, query EventBackfillListParams, opts ...option.RequestOption) *pagination.PageAutoPager[EventBackfillListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Closing a backfill makes the updated usage visible in Orb. Upon closing a
// backfill, Orb will asynchronously reflect the updated usage in invoice amounts
// and usage graphs. Once all of the updates are complete, the backfill's status
// will transition to `reflected`.
func (r *EventBackfillService) Close(ctx context.Context, backfillID string, opts ...option.RequestOption) (res *EventBackfillCloseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if backfillID == "" {
		err = errors.New("missing required backfill_id parameter")
		return
	}
	path := fmt.Sprintf("events/backfills/%s/close", backfillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint is used to fetch a backfill given an identifier.
func (r *EventBackfillService) Fetch(ctx context.Context, backfillID string, opts ...option.RequestOption) (res *EventBackfillFetchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if backfillID == "" {
		err = errors.New("missing required backfill_id parameter")
		return
	}
	path := fmt.Sprintf("events/backfills/%s", backfillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reverting a backfill undoes all the effects of closing the backfill. If the
// backfill is reflected, the status will transition to `pending_revert` while the
// effects of the backfill are undone. Once all effects are undone, the backfill
// will transition to `reverted`.
//
// If a backfill is reverted before its closed, no usage will be updated as a
// result of the backfill and it will immediately transition to `reverted`.
func (r *EventBackfillService) Revert(ctx context.Context, backfillID string, opts ...option.RequestOption) (res *EventBackfillRevertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if backfillID == "" {
		err = errors.New("missing required backfill_id parameter")
		return
	}
	path := fmt.Sprintf("events/backfills/%s/revert", backfillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type EventBackfillNewResponse struct {
	ID string `json:"id,required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time,required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer ID this backfill is scoped to. If null, this backfill is not scoped
	// to a single customer.
	CustomerID string `json:"customer_id,required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested,required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at,required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         EventBackfillNewResponseStatus `json:"status,required"`
	TimeframeEnd   time.Time                      `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                      `json:"timeframe_start,required" format:"date-time"`
	JSON           eventBackfillNewResponseJSON   `json:"-"`
}

// eventBackfillNewResponseJSON contains the JSON metadata for the struct
// [EventBackfillNewResponse]
type eventBackfillNewResponseJSON struct {
	ID             apijson.Field
	CloseTime      apijson.Field
	CreatedAt      apijson.Field
	CustomerID     apijson.Field
	EventsIngested apijson.Field
	RevertedAt     apijson.Field
	Status         apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventBackfillNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventBackfillNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type EventBackfillNewResponseStatus string

const (
	EventBackfillNewResponseStatusPending       EventBackfillNewResponseStatus = "pending"
	EventBackfillNewResponseStatusReflected     EventBackfillNewResponseStatus = "reflected"
	EventBackfillNewResponseStatusPendingRevert EventBackfillNewResponseStatus = "pending_revert"
	EventBackfillNewResponseStatusReverted      EventBackfillNewResponseStatus = "reverted"
)

func (r EventBackfillNewResponseStatus) IsKnown() bool {
	switch r {
	case EventBackfillNewResponseStatusPending, EventBackfillNewResponseStatusReflected, EventBackfillNewResponseStatusPendingRevert, EventBackfillNewResponseStatusReverted:
		return true
	}
	return false
}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type EventBackfillListResponse struct {
	ID string `json:"id,required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time,required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer ID this backfill is scoped to. If null, this backfill is not scoped
	// to a single customer.
	CustomerID string `json:"customer_id,required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested,required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at,required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         EventBackfillListResponseStatus `json:"status,required"`
	TimeframeEnd   time.Time                       `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                       `json:"timeframe_start,required" format:"date-time"`
	JSON           eventBackfillListResponseJSON   `json:"-"`
}

// eventBackfillListResponseJSON contains the JSON metadata for the struct
// [EventBackfillListResponse]
type eventBackfillListResponseJSON struct {
	ID             apijson.Field
	CloseTime      apijson.Field
	CreatedAt      apijson.Field
	CustomerID     apijson.Field
	EventsIngested apijson.Field
	RevertedAt     apijson.Field
	Status         apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventBackfillListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventBackfillListResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type EventBackfillListResponseStatus string

const (
	EventBackfillListResponseStatusPending       EventBackfillListResponseStatus = "pending"
	EventBackfillListResponseStatusReflected     EventBackfillListResponseStatus = "reflected"
	EventBackfillListResponseStatusPendingRevert EventBackfillListResponseStatus = "pending_revert"
	EventBackfillListResponseStatusReverted      EventBackfillListResponseStatus = "reverted"
)

func (r EventBackfillListResponseStatus) IsKnown() bool {
	switch r {
	case EventBackfillListResponseStatusPending, EventBackfillListResponseStatusReflected, EventBackfillListResponseStatusPendingRevert, EventBackfillListResponseStatusReverted:
		return true
	}
	return false
}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type EventBackfillCloseResponse struct {
	ID string `json:"id,required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time,required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer ID this backfill is scoped to. If null, this backfill is not scoped
	// to a single customer.
	CustomerID string `json:"customer_id,required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested,required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at,required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         EventBackfillCloseResponseStatus `json:"status,required"`
	TimeframeEnd   time.Time                        `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                        `json:"timeframe_start,required" format:"date-time"`
	JSON           eventBackfillCloseResponseJSON   `json:"-"`
}

// eventBackfillCloseResponseJSON contains the JSON metadata for the struct
// [EventBackfillCloseResponse]
type eventBackfillCloseResponseJSON struct {
	ID             apijson.Field
	CloseTime      apijson.Field
	CreatedAt      apijson.Field
	CustomerID     apijson.Field
	EventsIngested apijson.Field
	RevertedAt     apijson.Field
	Status         apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventBackfillCloseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventBackfillCloseResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type EventBackfillCloseResponseStatus string

const (
	EventBackfillCloseResponseStatusPending       EventBackfillCloseResponseStatus = "pending"
	EventBackfillCloseResponseStatusReflected     EventBackfillCloseResponseStatus = "reflected"
	EventBackfillCloseResponseStatusPendingRevert EventBackfillCloseResponseStatus = "pending_revert"
	EventBackfillCloseResponseStatusReverted      EventBackfillCloseResponseStatus = "reverted"
)

func (r EventBackfillCloseResponseStatus) IsKnown() bool {
	switch r {
	case EventBackfillCloseResponseStatusPending, EventBackfillCloseResponseStatusReflected, EventBackfillCloseResponseStatusPendingRevert, EventBackfillCloseResponseStatusReverted:
		return true
	}
	return false
}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type EventBackfillFetchResponse struct {
	ID string `json:"id,required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time,required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer ID this backfill is scoped to. If null, this backfill is not scoped
	// to a single customer.
	CustomerID string `json:"customer_id,required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested,required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at,required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         EventBackfillFetchResponseStatus `json:"status,required"`
	TimeframeEnd   time.Time                        `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                        `json:"timeframe_start,required" format:"date-time"`
	JSON           eventBackfillFetchResponseJSON   `json:"-"`
}

// eventBackfillFetchResponseJSON contains the JSON metadata for the struct
// [EventBackfillFetchResponse]
type eventBackfillFetchResponseJSON struct {
	ID             apijson.Field
	CloseTime      apijson.Field
	CreatedAt      apijson.Field
	CustomerID     apijson.Field
	EventsIngested apijson.Field
	RevertedAt     apijson.Field
	Status         apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventBackfillFetchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventBackfillFetchResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type EventBackfillFetchResponseStatus string

const (
	EventBackfillFetchResponseStatusPending       EventBackfillFetchResponseStatus = "pending"
	EventBackfillFetchResponseStatusReflected     EventBackfillFetchResponseStatus = "reflected"
	EventBackfillFetchResponseStatusPendingRevert EventBackfillFetchResponseStatus = "pending_revert"
	EventBackfillFetchResponseStatusReverted      EventBackfillFetchResponseStatus = "reverted"
)

func (r EventBackfillFetchResponseStatus) IsKnown() bool {
	switch r {
	case EventBackfillFetchResponseStatusPending, EventBackfillFetchResponseStatusReflected, EventBackfillFetchResponseStatusPendingRevert, EventBackfillFetchResponseStatusReverted:
		return true
	}
	return false
}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type EventBackfillRevertResponse struct {
	ID string `json:"id,required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time,required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The customer ID this backfill is scoped to. If null, this backfill is not scoped
	// to a single customer.
	CustomerID string `json:"customer_id,required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested,required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at,required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         EventBackfillRevertResponseStatus `json:"status,required"`
	TimeframeEnd   time.Time                         `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                         `json:"timeframe_start,required" format:"date-time"`
	JSON           eventBackfillRevertResponseJSON   `json:"-"`
}

// eventBackfillRevertResponseJSON contains the JSON metadata for the struct
// [EventBackfillRevertResponse]
type eventBackfillRevertResponseJSON struct {
	ID             apijson.Field
	CloseTime      apijson.Field
	CreatedAt      apijson.Field
	CustomerID     apijson.Field
	EventsIngested apijson.Field
	RevertedAt     apijson.Field
	Status         apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventBackfillRevertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventBackfillRevertResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type EventBackfillRevertResponseStatus string

const (
	EventBackfillRevertResponseStatusPending       EventBackfillRevertResponseStatus = "pending"
	EventBackfillRevertResponseStatusReflected     EventBackfillRevertResponseStatus = "reflected"
	EventBackfillRevertResponseStatusPendingRevert EventBackfillRevertResponseStatus = "pending_revert"
	EventBackfillRevertResponseStatusReverted      EventBackfillRevertResponseStatus = "reverted"
)

func (r EventBackfillRevertResponseStatus) IsKnown() bool {
	switch r {
	case EventBackfillRevertResponseStatusPending, EventBackfillRevertResponseStatusReflected, EventBackfillRevertResponseStatusPendingRevert, EventBackfillRevertResponseStatusReverted:
		return true
	}
	return false
}

type EventBackfillNewParams struct {
	// The (exclusive) end of the usage timeframe affected by this backfill.
	TimeframeEnd param.Field[time.Time] `json:"timeframe_end,required" format:"date-time"`
	// The (inclusive) start of the usage timeframe affected by this backfill.
	TimeframeStart param.Field[time.Time] `json:"timeframe_start,required" format:"date-time"`
	// The time at which no more events will be accepted for this backfill. The
	// backfill will automatically begin reflecting throughout Orb at the close time.
	// If not specified, it will default to 1 day after the creation of the backfill.
	CloseTime param.Field[time.Time] `json:"close_time" format:"date-time"`
	// The ID of the customer to which this backfill is scoped.
	CustomerID param.Field[string] `json:"customer_id"`
	// The external customer ID of the customer to which this backfill is scoped.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// If true, replaces all existing events in the timeframe with the newly ingested
	// events. If false, adds the newly ingested events to the existing events.
	ReplaceExistingEvents param.Field[bool] `json:"replace_existing_events"`
}

func (r EventBackfillNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventBackfillListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [EventBackfillListParams]'s query parameters as
// `url.Values`.
func (r EventBackfillListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
