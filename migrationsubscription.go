// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"github.com/orbcorp/orb-go/option"
)

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
