// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// PlanExternalPlanIDService contains methods and other services that help with
// interacting with the orb API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPlanExternalPlanIDService] method
// instead.
type PlanExternalPlanIDService struct {
	Options []option.RequestOption
}

// NewPlanExternalPlanIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPlanExternalPlanIDService(opts ...option.RequestOption) (r *PlanExternalPlanIDService) {
	r = &PlanExternalPlanIDService{}
	r.Options = opts
	return
}

// This endpoint is used to fetch [plan](../guides/concepts##plan-and-price)
// details given an external_plan_id identifier. It returns information about the
// prices included in the plan and their configuration, as well as the product that
// the plan is attached to.
//
// ## Serialized prices
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given [Price](../guides/concepts#plan-and-price)
// object. The `model_type` field determines the key for the configuration object
// that is present. A detailed explanation of price types can be found in the
// [Price schema](../guides/concepts#plan-and-price).
func (r *PlanExternalPlanIDService) Update(ctx context.Context, otherExternalPlanID string, body PlanExternalPlanIDUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("plans/external_plan_id/%s", otherExternalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch [plan](../guides/concepts##plan-and-price)
// details given an external_plan_id identifier. It returns information about the
// prices included in the plan and their configuration, as well as the product that
// the plan is attached to.
//
// ## Serialized prices
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given [Price](../guides/concepts#plan-and-price)
// object. The `model_type` field determines the key for the configuration object
// that is present. A detailed explanation of price types can be found in the
// [Price schema](../guides/concepts#plan-and-price). "
func (r *PlanExternalPlanIDService) Fetch(ctx context.Context, externalPlanID string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("plans/external_plan_id/%s", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PlanExternalPlanIDUpdateParams struct {
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
