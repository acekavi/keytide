package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// A middleware is an http.Handler decorator, so testing one means wrapping a
// stub that records whether it was reached and asserting on both the status
// and that fact. Reaching the next handler is the thing that matters: a
// middleware can return 401 and still have called through, which no assertion
// on the status code alone would catch.
func run(t *testing.T, verify Verifier, header string) (*httptest.ResponseRecorder, bool) {
	t.Helper()

	reached := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reached = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	if header != "" {
		req.Header.Set("Authorization", header)
	}

	rec := httptest.NewRecorder()
	AuthMiddleware(verify)(next).ServeHTTP(rec, req)
	return rec, reached
}

func acceptAll(token string) (string, error) { return "user-1", nil }

func TestRejectsWhenNoVerifierConfigured(t *testing.T) {
	// The previous placeholder returned true for anything starting with
	// "Bearer ", so this exact request was authorised. It must not be.
	rec, reached := run(t, RejectAll, "Bearer anything-at-all")

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
	if reached {
		t.Error("next handler was reached on a rejected token")
	}
}

func TestRejectsMissingAndMalformedHeaders(t *testing.T) {
	cases := map[string]string{
		"absent":       "",
		"empty token":  "Bearer ",
		"wrong scheme": "Basic dXNlcjpwYXNz",
		"no scheme":    "just-a-token",
	}

	for name, header := range cases {
		t.Run(name, func(t *testing.T) {
			rec, reached := run(t, acceptAll, header)
			if rec.Code != http.StatusUnauthorized {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
			}
			if reached {
				t.Error("next handler was reached")
			}
		})
	}
}

func TestPassesVerifiedRequestThrough(t *testing.T) {
	rec, reached := run(t, acceptAll, "Bearer good-token")

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if !reached {
		t.Error("next handler was not reached for a valid token")
	}
}

func TestPutsSubjectOnContext(t *testing.T) {
	var got string
	var ok bool

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, ok = SubjectFrom(r.Context())
	})

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Set("Authorization", "Bearer good-token")
	AuthMiddleware(acceptAll)(next).ServeHTTP(httptest.NewRecorder(), req)

	if !ok || got != "user-1" {
		t.Errorf("SubjectFrom = (%q, %v), want (%q, true)", got, ok, "user-1")
	}
}

func TestDoesNotLeakWhyVerificationFailed(t *testing.T) {
	// Distinguishing "expired" from "bad signature" tells an attacker which
	// part of the token to change, so the body must not carry the reason.
	secret := errors.New("token expired at 2026-01-01 for issuer https://internal")
	verify := func(string) (string, error) { return "", secret }

	rec, _ := run(t, verify, "Bearer expired")

	if body := rec.Body.String(); body != "Unauthorized\n" {
		t.Errorf("body = %q, want %q", body, "Unauthorized\n")
	}
}

func TestSendsWWWAuthenticateOn401(t *testing.T) {
	rec, _ := run(t, RejectAll, "")

	if got := rec.Header().Get("WWW-Authenticate"); got == "" {
		t.Error("401 sent without WWW-Authenticate, which RFC 9110 requires")
	}
}

func TestNilVerifierFailsClosed(t *testing.T) {
	// An AuthMiddleware constructed with no verifier must reject, not admit.
	rec, reached := run(t, nil, "Bearer anything")

	if rec.Code != http.StatusUnauthorized || reached {
		t.Errorf("nil verifier admitted a request: status=%d reached=%v", rec.Code, reached)
	}
}
