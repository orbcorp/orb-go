// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"github.com/orbcorp/orb-go/option"
)

// MigrationService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMigrationService] method instead.
type MigrationService struct {
	Options       []option.RequestOption
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
