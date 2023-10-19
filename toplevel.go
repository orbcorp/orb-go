// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"net/http"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// TopLevelService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTopLevelService] method instead.
type TopLevelService struct {
	Options []option.RequestOption
}

// NewTopLevelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTopLevelService(opts ...option.RequestOption) (r *TopLevelService) {
	r = &TopLevelService{}
	r.Options = opts
	return
}

// This endpoint allows you to test your connection to the Orb API and check the
// validity of your API key, passed in the Authorization header. This is
// particularly useful for checking that your environment is set up properly, and
// is a great choice for connectors and integrations.
//
// This API does not have any side-effects or return any Orb resources.
func (r *TopLevelService) Ping(ctx context.Context, opts ...option.RequestOption) (res *TopLevelPingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TopLevelPingResponse struct {
	Response string `json:"response,required"`
	JSON     topLevelPingResponseJSON
}

// topLevelPingResponseJSON contains the JSON metadata for the struct
// [TopLevelPingResponse]
type topLevelPingResponseJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TopLevelPingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
