// Package getoembed defines a GetOEmbedHandler outside of the
// configure_*.go file so it will be easier to regenerate that file if
// the swagger.yaml changes
package getoembed

import (
	_ "github.com/GannettDigital/{{ RepoName }}/config" // bootstrap config
	"github.com/GannettDigital/{{ RepoName }}/swag/models"
	"github.com/GannettDigital/{{ RepoName }}/swag/restapi/operations"
	middleware "github.com/go-openapi/runtime/middleware"
)

// New returns a operations.GetOEmbedHandler
func New() operations.GetOEmbedHandler {
	return getOEmbed{}
}

type getOEmbed struct{}

// GetOEmbedHandler a simple handler for the service
func (getOEmbed) Handle(params operations.GetOEmbedParams) middleware.Responder {
	return operations.NewGetOEmbedOK().WithPayload(
		&models.OEmbed{
			HTML:    `<div><em>Hello, World!</em></div>`,
			Type:    "html",
			Version: "1.0",
		},
	)
}
