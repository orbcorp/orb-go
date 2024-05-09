// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
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

// MetricService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetricService] method instead.
type MetricService struct {
	Options []option.RequestOption
}

// NewMetricService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetricService(opts ...option.RequestOption) (r *MetricService) {
	r = &MetricService{}
	r.Options = opts
	return
}

// This endpoint is used to create a [metric](../guides/concepts##metric) using a
// SQL string. See
// [SQL support](../guides/extensibility/advanced-metrics#sql-support) for a
// description of constructing SQL queries with examples.
func (r *MetricService) New(ctx context.Context, body MetricNewParams, opts ...option.RequestOption) (res *MetricNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch [metric](../guides/concepts#metric) details given
// a metric identifier. It returns information about the metrics including its
// name, description, and item.
func (r *MetricService) List(ctx context.Context, query MetricListParams, opts ...option.RequestOption) (res *pagination.Page[MetricListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "metrics"
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

// This endpoint is used to fetch [metric](../guides/concepts#metric) details given
// a metric identifier. It returns information about the metrics including its
// name, description, and item.
func (r *MetricService) ListAutoPaging(ctx context.Context, query MetricListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MetricListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to list [metrics](../guides/concepts##metric). It returns
// information about the metrics including its name, description, and item.
func (r *MetricService) Fetch(ctx context.Context, metricID string, opts ...option.RequestOption) (res *MetricFetchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("metrics/%s", metricID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The Metric resource represents a calculation of a quantity based on events.
// Metrics are defined by the query that transforms raw usage events into
// meaningful values for your customers.
type MetricNewResponse struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// The Item resource represents a sellable product or good. Items are associated
	// with all line items, billable metrics, and prices and are used for defining
	// external sync behavior for invoices and tax calculation purposes.
	Item Item `json:"item,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string       `json:"metadata,required"`
	Name     string                  `json:"name,required"`
	Status   MetricNewResponseStatus `json:"status,required"`
	JSON     metricNewResponseJSON   `json:"-"`
}

// metricNewResponseJSON contains the JSON metadata for the struct
// [MetricNewResponse]
type metricNewResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Item        apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricNewResponseJSON) RawJSON() string {
	return r.raw
}

type MetricNewResponseStatus string

const (
	MetricNewResponseStatusActive   MetricNewResponseStatus = "active"
	MetricNewResponseStatusDraft    MetricNewResponseStatus = "draft"
	MetricNewResponseStatusArchived MetricNewResponseStatus = "archived"
)

func (r MetricNewResponseStatus) IsKnown() bool {
	switch r {
	case MetricNewResponseStatusActive, MetricNewResponseStatusDraft, MetricNewResponseStatusArchived:
		return true
	}
	return false
}

// The Metric resource represents a calculation of a quantity based on events.
// Metrics are defined by the query that transforms raw usage events into
// meaningful values for your customers.
type MetricListResponse struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// The Item resource represents a sellable product or good. Items are associated
	// with all line items, billable metrics, and prices and are used for defining
	// external sync behavior for invoices and tax calculation purposes.
	Item Item `json:"item,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string        `json:"metadata,required"`
	Name     string                   `json:"name,required"`
	Status   MetricListResponseStatus `json:"status,required"`
	JSON     metricListResponseJSON   `json:"-"`
}

// metricListResponseJSON contains the JSON metadata for the struct
// [MetricListResponse]
type metricListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Item        apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricListResponseJSON) RawJSON() string {
	return r.raw
}

type MetricListResponseStatus string

const (
	MetricListResponseStatusActive   MetricListResponseStatus = "active"
	MetricListResponseStatusDraft    MetricListResponseStatus = "draft"
	MetricListResponseStatusArchived MetricListResponseStatus = "archived"
)

func (r MetricListResponseStatus) IsKnown() bool {
	switch r {
	case MetricListResponseStatusActive, MetricListResponseStatusDraft, MetricListResponseStatusArchived:
		return true
	}
	return false
}

// The Metric resource represents a calculation of a quantity based on events.
// Metrics are defined by the query that transforms raw usage events into
// meaningful values for your customers.
type MetricFetchResponse struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// The Item resource represents a sellable product or good. Items are associated
	// with all line items, billable metrics, and prices and are used for defining
	// external sync behavior for invoices and tax calculation purposes.
	Item Item `json:"item,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string         `json:"metadata,required"`
	Name     string                    `json:"name,required"`
	Status   MetricFetchResponseStatus `json:"status,required"`
	JSON     metricFetchResponseJSON   `json:"-"`
}

// metricFetchResponseJSON contains the JSON metadata for the struct
// [MetricFetchResponse]
type metricFetchResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Item        apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricFetchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricFetchResponseJSON) RawJSON() string {
	return r.raw
}

type MetricFetchResponseStatus string

const (
	MetricFetchResponseStatusActive   MetricFetchResponseStatus = "active"
	MetricFetchResponseStatusDraft    MetricFetchResponseStatus = "draft"
	MetricFetchResponseStatusArchived MetricFetchResponseStatus = "archived"
)

func (r MetricFetchResponseStatus) IsKnown() bool {
	switch r {
	case MetricFetchResponseStatusActive, MetricFetchResponseStatusDraft, MetricFetchResponseStatusArchived:
		return true
	}
	return false
}

type MetricNewParams struct {
	// A description of the metric.
	Description param.Field[string] `json:"description,required"`
	// The id of the item
	ItemID param.Field[string] `json:"item_id,required"`
	// The name of the metric.
	Name param.Field[string] `json:"name,required"`
	// A sql string defining the metric.
	Sql param.Field[string] `json:"sql,required"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r MetricNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MetricListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [MetricListParams]'s query parameters as `url.Values`.
func (r MetricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
