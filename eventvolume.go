// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// EventVolumeService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventVolumeService] method instead.
type EventVolumeService struct {
	Options []option.RequestOption
}

// NewEventVolumeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventVolumeService(opts ...option.RequestOption) (r *EventVolumeService) {
	r = &EventVolumeService{}
	r.Options = opts
	return
}

// This endpoint returns the event volume for an account in a
// [paginated list format](../reference/pagination).
//
// The event volume is aggregated by the hour and the
// [timestamp](../reference/ingest#determining-event-timestamp) field is used to
// determine which hour an event is associated with. Note, this means that
// late-arriving events increment the volume count for the hour window the
// timestamp is in, not the latest hour window.
//
// Each item in the response contains the count of events aggregated by the hour
// where the start and end time are hour-aligned and in UTC. When a specific
// timestamp is passed in for either start or end time, the response includes the
// hours the timestamp falls in.
func (r *EventVolumeService) List(ctx context.Context, query EventVolumeListParams, opts ...option.RequestOption) (res *EventVolumes, err error) {
	opts = append(r.Options[:], opts...)
	path := "events/volume"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EventVolumes struct {
	Data []EventVolumesData `json:"data,required"`
	JSON eventVolumesJSON   `json:"-"`
}

// eventVolumesJSON contains the JSON metadata for the struct [EventVolumes]
type eventVolumesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventVolumes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventVolumesJSON) RawJSON() string {
	return r.raw
}

// An EventVolume contains the event volume ingested in an hourly window. The
// timestamp used for the aggregation is the `timestamp` datetime field on events.
type EventVolumesData struct {
	// The number of events ingested with a timestamp between the timeframe
	Count          int64                `json:"count,required"`
	TimeframeEnd   time.Time            `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time            `json:"timeframe_start,required" format:"date-time"`
	JSON           eventVolumesDataJSON `json:"-"`
}

// eventVolumesDataJSON contains the JSON metadata for the struct
// [EventVolumesData]
type eventVolumesDataJSON struct {
	Count          apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventVolumesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventVolumesDataJSON) RawJSON() string {
	return r.raw
}

type EventVolumeListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// The end of the timeframe, exclusive, in which to return event volume. If not
	// specified, the current time is used. All datetime values are converted to UTC
	// time.If the specified time isn't hour-aligned, the response includes the event
	// volumecount for the hour the time falls in.
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// The start of the timeframe, inclusive, in which to return event volume. All
	// datetime values are converted to UTC time. If the specified time isn't
	// hour-aligned, the response includes the event volume count for the hour the time
	// falls in.
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
}

// URLQuery serializes [EventVolumeListParams]'s query parameters as `url.Values`.
func (r EventVolumeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
