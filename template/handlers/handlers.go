// Package handlers defines functions that set up the the middleware
// for the service.
package handlers

import (
	_ "github.com/GannettDigital/{{ RepoName }}/config" // init config
	"github.com/rs/cors"
	"net/http"
)

// SetupMiddlewares The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// SetupGlobalMiddleware The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = cors.Default().Handler(handler)
	return handler
}
