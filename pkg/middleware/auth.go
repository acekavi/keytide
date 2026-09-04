package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// contextKey keys values this package puts on a request context.
//
// An unexported struct type rather than a plain string: a string key can
// collide with one set by any other package in the process, silently.
type contextKey struct{ name string }

var subjectKey = &contextKey{"subject"}

// SubjectFrom returns the authenticated subject, if the request passed through
// AuthMiddleware.
func SubjectFrom(ctx context.Context) (string, bool) {
	subject, ok := ctx.Value(subjectKey).(string)
	return subject, ok
}

// ErrNoVerifier is returned by the placeholder verifier.
var ErrNoVerifier = errors.New("middleware: no token verifier configured")

// Verifier turns a bearer token into a subject, or returns an error.
//
// This is the seam. Real verification means checking the signature, pinning
// the algorithm, and validating the exp / aud / iss claims — which needs a JWT
// library and a key source, neither of which belongs in a middleware package.
// Expressing it as a function type means swapping the placeholder for the real
// thing changes one value at startup and touches nothing else.
type Verifier func(token string) (subject string, err error)

// RejectAll is the default verifier: it authenticates nobody.
//
// This replaces a function that returned true for any string beginning with
// "Bearer ". That was a more dangerous kind of placeholder than it looked:
// it resembled authentication and passed every request, so a reader skimming
// the file could reasonably believe the endpoint was protected. This one fails
// closed — an unconfigured server rejects traffic rather than serving it to
// anyone who sends a header.
func RejectAll(string) (string, error) {
	return "", ErrNoVerifier
}

// AuthMiddleware rejects any request whose bearer token the verifier will not
// accept, and puts the resulting subject on the request context.
//
// The signature is func(http.Handler) http.Handler, the standard shape, so
// this composes with any router and any other middleware.
func AuthMiddleware(verify Verifier) func(http.Handler) http.Handler {
	if verify == nil {
		verify = RejectAll
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CutPrefix rather than HasPrefix: the token is what follows the
			// scheme, and a header of exactly "Bearer " carries no token.
			raw, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
			if !ok || raw == "" {
				unauthorized(w)
				return
			}

			subject, err := verify(raw)
			if err != nil {
				// The reason is deliberately not echoed to the caller.
				// Distinguishing "expired" from "bad signature" from "unknown
				// issuer" tells an attacker which part of the token to change.
				unauthorized(w)
				return
			}

			ctx := context.WithValue(r.Context(), subjectKey, subject)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	// RFC 9110 requires WWW-Authenticate on a 401; it tells a client which
	// scheme to retry with.
	w.Header().Set("WWW-Authenticate", `Bearer realm="keytide"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
