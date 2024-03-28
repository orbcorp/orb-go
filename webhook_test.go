// File generated from our OpenAPI spec by Stainless.

package orb_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/orbcorp/orb-go"
	"github.com/orbcorp/orb-go/option"
)

func TestVerifySignature(t *testing.T) {
	secret := "c-UGKYdnhHh436B_sMouYAPUvXyWpzOdunZBV5dFSD8"
	payload := `{"id": "o4mmewpfNNTnjfZc", "created_at": "2024-03-27T15:42:29+00:00", "type": "resource_event.test", "properties": {"message": "A test webhook from Orb. Happy testing!"}}`
	signature := "9d25de966891ab0bc18754faf8d83d0980b44ae330fcc130b41a6cf3daf1f391"

	timestamp := "2024-03-27T15:42:29.551"
	fakeNow, err := time.Parse(orb.WebhookHeaderTimestampFormat, timestamp)
	if err != nil {
		t.Fatalf("did not expect error timestamp parsing error: %s", err.Error())
	}

	header := http.Header{}
	header.Add("X-Orb-Timestamp", timestamp)
	header.Add("X-Orb-Signature", fmt.Sprintf("v1=%s", signature))

	var testCases = map[string]struct {
		payload     string
		headers     http.Header
		secret      string
		fakeNow     time.Time
		expectedErr string
	}{
		"valid signature": {
			payload: payload,
			headers: header,
			secret:  secret,
			fakeNow: fakeNow,
		},
		"timestamp outside threshold (too old)": {
			payload:     payload,
			headers:     header,
			secret:      secret,
			fakeNow:     fakeNow.Add(6 * time.Minute),
			expectedErr: "value from X-Orb-Timestamp header too old",
		},
		"timestamp outside threshold (too new)": {
			payload:     payload,
			headers:     header,
			secret:      secret,
			fakeNow:     fakeNow.Add(-6 * time.Minute),
			expectedErr: "value from X-Orb-Timestamp header too new",
		},
		"invalid signature": {
			payload:     payload,
			headers:     header,
			secret:      "foo",
			fakeNow:     fakeNow,
			expectedErr: "None of the given webhook signatures match the expected signature",
		},
		"multiple signatures": {
			payload: payload,
			headers: http.Header{
				"X-Orb-Timestamp": []string{timestamp},
				"X-Orb-Signature": []string{"v1=my-invalid-signature", fmt.Sprintf("v1=%s", signature)},
			},
			secret:  secret,
			fakeNow: fakeNow,
		},
		"different signature version": {
			payload: payload,
			headers: http.Header{
				"X-Orb-Timestamp": []string{timestamp},
				"X-Orb-Signature": []string{fmt.Sprintf("v2=%s", signature)},
			},
			secret:      secret,
			fakeNow:     fakeNow,
			expectedErr: "None of the given webhook signatures match the expected signature",
		},
		"multiple signatures with different version": {
			payload: payload,
			headers: http.Header{
				"X-Orb-Timestamp": []string{timestamp},
				"X-Orb-Signature": []string{fmt.Sprintf("v2=%s", signature), fmt.Sprintf("v1=%s", signature)},
			},
			secret:  secret,
			fakeNow: fakeNow,
		},
		"signature version is missing": {
			payload: payload,
			headers: http.Header{
				"X-Orb-Timestamp": []string{timestamp},
				"X-Orb-Signature": []string{signature},
			},
			secret:      secret,
			fakeNow:     fakeNow,
			expectedErr: "None of the given webhook signatures match the expected signature",
		},
		"secret not set": {
			payload:     payload,
			headers:     header,
			secret:      "",
			fakeNow:     fakeNow,
			expectedErr: "The webhook secret must either be set using the env var, ORB_WEBHOOK_SECRET, on the client class, orb.NewClient(option.WithWebhookSecret(\"123\")}), or passed to this function",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			client := orb.NewClient()
			err = client.Webhooks.VerifySignature([]byte(test.payload), test.headers, test.secret, test.fakeNow)
			if test.expectedErr != "" {
				if err == nil {
					t.Fatalf("expected error: '%s', got nil", test.expectedErr)
				}
				if err.Error() != test.expectedErr {
					t.Fatalf("expected error: '%s', got: '%s'", test.expectedErr, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect error: %s", err.Error())
				}
			}
		})
	}

	t.Run("webhook secret set as client option", func(t *testing.T) {
		client := orb.NewClient(option.WithWebhookSecret(secret))
		err = client.Webhooks.VerifySignature([]byte(payload), header, "", fakeNow)
		if err != nil {
			t.Fatalf("did not expect error: %s", err.Error())
		}
	})
}
