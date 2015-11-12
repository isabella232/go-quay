package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetOrganizationApplicationReader struct {
	formats strfmt.Registry
}

func (o *GetOrganizationApplicationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetOrganizationApplicationOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetOrganizationApplicationBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationApplicationBadRequest", &result, response.Code())

	case 401:
		var result GetOrganizationApplicationUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationApplicationUnauthorized", &result, response.Code())

	case 403:
		var result GetOrganizationApplicationForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationApplicationForbidden", &result, response.Code())

	case 404:
		var result GetOrganizationApplicationNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationApplicationNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetOrganizationApplicationOK struct {
}

func (o *GetOrganizationApplicationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetOrganizationApplicationBadRequest struct {
}

func (o *GetOrganizationApplicationBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetOrganizationApplicationUnauthorized struct {
}

func (o *GetOrganizationApplicationUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetOrganizationApplicationForbidden struct {
}

func (o *GetOrganizationApplicationForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetOrganizationApplicationNotFound struct {
}

func (o *GetOrganizationApplicationNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}