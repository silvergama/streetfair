// Package Fair API.
//
// The purpose of this application is to provide an application
// to manage street fairs in the city of São Paulo.
//
//     Schemes: http
//     Host: localhost:9000
//     BasePath: /v1
//     Version: 1.0
//     Contact: Silver Gama<silver.mdg@gmail.com> https://github.com/silvergama
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"github.com/silvergama/streetfair/common/response"
	"github.com/silvergama/streetfair/fair"
)

// Not Found
// swagger:response notFound
type docNotFoundResponse struct {
	Body response.Error
}

// Bad Request
// swagger:response badRequest
type docBadRequestResponse struct {
	Body response.Error
}

// Internal Server Error
// swagger:response internalServerError
type docInternalServerErrorResponse struct {
	// in: body
	Body *response.Error
}

// No Content
// swagger:response noContent
type docNoContentResponse struct {
	// in: body
	Body *response.Success
}

// Success
// swagger:response success
type docSuccessResponse struct {
	// in: body
	Body *response.Success
}

// swagger:parameters fairsGetV1Req
type docFairsGetV1Request struct {
	// in: query
	Neighborhood string `json:"neighborhood"`
}

// swagger:parameters fairPostV1Req
type docFairPostV1Request struct {
	// in: body
	Body *fair.Fair
}

// swagger:parameters fairPutV1Req
type docFairPutV1Request struct {
	// in: path
	ID int `json:"id"`
	// in: body
	Body *fair.Fair
}

// swagger:parameters fairDeleteV1Req
type docFairDeleteV1Request struct {
	// in: path
	ID int `json:"id"`
}
