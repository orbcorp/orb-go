// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
)

// A migration moves a plan's subscriptions onto one of the plan's versions. Each
// subscription is migrated asynchronously, either immediately or at a scheduled
// time, and the migration reports its overall status as it progresses.
//
// MigrationSubscriptionService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMigrationSubscriptionService] method instead.
type MigrationSubscriptionService struct {
	Options []option.RequestOption
}

// NewMigrationSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMigrationSubscriptionService(opts ...option.RequestOption) (r *MigrationSubscriptionService) {
	r = &MigrationSubscriptionService{}
	r.Options = opts
	return
}

// This endpoint returns the subscriptions a migration applies to, with the status
// of each one. A subscription that could not be migrated has the status `failed`
// and a `failure_reason`.
//
// The list is empty until the migration has finished selecting its subscriptions.
// Once a migration is canceled, only the subscriptions it processed before the
// cancellation are returned.
func (r *MigrationSubscriptionService) List(ctx context.Context, migrationID string, query MigrationSubscriptionListParams, opts ...option.RequestOption) (res *pagination.Page[MigrationSubscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if migrationID == "" {
		err = errors.New("missing required migration_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("migrations/%s/subscriptions", migrationID)
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

// This endpoint returns the subscriptions a migration applies to, with the status
// of each one. A subscription that could not be migrated has the status `failed`
// and a `failure_reason`.
//
// The list is empty until the migration has finished selecting its subscriptions.
// Once a migration is canceled, only the subscriptions it processed before the
// cancellation are returned.
func (r *MigrationSubscriptionService) ListAutoPaging(ctx context.Context, migrationID string, query MigrationSubscriptionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[MigrationSubscription] {
	return pagination.NewPageAutoPager(r.List(ctx, migrationID, query, opts...))
}

type MigrationSubscriptionListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// Only return subscriptions with one of these statuses.
	Status param.Field[[]MigrationSubscriptionListParamsStatus] `query:"status"`
	// Only return these subscriptions. Up to 100 IDs.
	SubscriptionID param.Field[[]string] `query:"subscription_id"`
}

// URLQuery serializes [MigrationSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r MigrationSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MigrationSubscriptionListParamsStatus string

const (
	MigrationSubscriptionListParamsStatusScheduling MigrationSubscriptionListParamsStatus = "scheduling"
	MigrationSubscriptionListParamsStatusScheduled  MigrationSubscriptionListParamsStatus = "scheduled"
	MigrationSubscriptionListParamsStatusCompleted  MigrationSubscriptionListParamsStatus = "completed"
	MigrationSubscriptionListParamsStatusSkipped    MigrationSubscriptionListParamsStatus = "skipped"
	MigrationSubscriptionListParamsStatusFailed     MigrationSubscriptionListParamsStatus = "failed"
)

func (r MigrationSubscriptionListParamsStatus) IsKnown() bool {
	switch r {
	case MigrationSubscriptionListParamsStatusScheduling, MigrationSubscriptionListParamsStatusScheduled, MigrationSubscriptionListParamsStatusCompleted, MigrationSubscriptionListParamsStatusSkipped, MigrationSubscriptionListParamsStatusFailed:
		return true
	}
	return false
}
