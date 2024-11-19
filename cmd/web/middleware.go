package main

import (
	"net/http"

	"github.com/gorilla/csrf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	CSRF := csrf.Protect(
		[]byte("a-32-byte-long-key-goes-here"),
		csrf.Path("/"),
		// instruct the browser to never send cookies during cross site requests
		csrf.SameSite(csrf.SameSiteStrictMode),
	)
	return CSRF(next)
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
