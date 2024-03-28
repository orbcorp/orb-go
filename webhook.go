// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// WebhookHeaderTimestampFormat is the format of the header X-Orb-Timestamp for webhook requests sent by Orb.
const WebhookHeaderTimestampFormat = "2006-01-02T15:04:05.999999999"

// WebhookService contains methods and other services that help with interacting
// with the Orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption

	// webhookSecret is the secret defined at the client level
	webhookSecret string
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts

	// This is a dummy response object. We need to build a request config to be able to check the webhook secret defined
	// at the client level.
	_res := struct{}{}
	cfg, err := requestconfig.NewRequestConfig(context.TODO(), http.MethodPost, "/webhooks", nil, &_res, opts...)
	if err != nil {
		panic(err)
	}

	if cfg.WebhookSecret != "" {
		r.webhookSecret = cfg.WebhookSecret
	}
	return
}

// Validates whether or not the webhook payload was sent by Orb. Pass an empty string to use the secret defined at the
// client level.
//
// An error will be raised if the webhook payload was not sent by Orb.
func (r *WebhookService) VerifySignature(payload []byte, headers http.Header, secret string, now time.Time) (err error) {
	if secret == "" {
		secret = r.webhookSecret
	}
	if secret == "" {
		return errors.New("The webhook secret must either be set using the env var, ORB_WEBHOOK_SECRET, on the client class, orb.NewClient(option.WithWebhookSecret(\"123\")}), or passed to this function")
	}

	msgSignature := headers.Values("X-Orb-Signature")
	if len(msgSignature) == 0 {
		return errors.New("could not find X-Orb-Signature header")
	}
	msgTimestamp := headers.Get("X-Orb-Timestamp")
	if len(msgTimestamp) == 0 {
		return errors.New("could not find X-Orb-Timestamp header")
	}

	timestamp, err := time.Parse(WebhookHeaderTimestampFormat, msgTimestamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp headers: %s", err)
	}

	if timestamp.Before(now.Add(-5 * time.Minute)) {
		return errors.New("value from X-Orb-Timestamp header too old")
	}
	if timestamp.After(now.Add(5 * time.Minute)) {
		return errors.New("value from X-Orb-Timestamp header too new")
	}

	secretBytes := []byte(secret)
	mac := hmac.New(sha256.New, secretBytes)
	mac.Write([]byte("v1:"))
	mac.Write([]byte(msgTimestamp))
	mac.Write([]byte(":"))
	mac.Write(payload)
	expected := mac.Sum(nil)

	for _, part := range msgSignature {
		parts := strings.Split(part, "=")
		if len(parts) != 2 {
			continue
		}
		if parts[0] != "v1" {
			continue
		}
		signature, err := hex.DecodeString(parts[1])
		if err != nil {
			continue
		}
		if hmac.Equal(signature, expected) {
			return nil
		}
	}

	return errors.New("None of the given webhook signatures match the expected signature")
}

type WebhookVerifySignatureParams struct {
}
