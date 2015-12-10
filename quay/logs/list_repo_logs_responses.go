package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type ListRepoLogsReader struct {
	formats strfmt.Registry
}

func (o *ListRepoLogsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ListRepoLogsOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ListRepoLogsBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoLogsBadRequest", &result, response.Code())

	case 401:
		var result ListRepoLogsUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoLogsUnauthorized", &result, response.Code())

	case 403:
		var result ListRepoLogsForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoLogsForbidden", &result, response.Code())

	case 404:
		var result ListRepoLogsNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoLogsNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*ListRepoLogsOK

Successful invocation
*/
type ListRepoLogsOK struct {
}

func (o *ListRepoLogsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ListRepoLogsBadRequest

Bad Request
*/
type ListRepoLogsBadRequest struct {
	Payload *models.GeneralError
}

func (o *ListRepoLogsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*ListRepoLogsUnauthorized

Session required
*/
type ListRepoLogsUnauthorized struct {
}

func (o *ListRepoLogsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ListRepoLogsForbidden

Unauthorized access
*/
type ListRepoLogsForbidden struct {
}

func (o *ListRepoLogsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ListRepoLogsNotFound

Not found
*/
type ListRepoLogsNotFound struct {
}

func (o *ListRepoLogsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}