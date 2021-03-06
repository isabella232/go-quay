package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewConductSearchParams creates a new ConductSearchParams object
// with the default values initialized.
func NewConductSearchParams() *ConductSearchParams {
	var ()
	return &ConductSearchParams{}
}

/*ConductSearchParams contains all the parameters to send to the API endpoint
for the conduct search operation typically these are written to a http.Request
*/
type ConductSearchParams struct {

	/*Query
	  The search query.

	*/
	Query *string
}

// WithQuery adds the query to the conduct search params
func (o *ConductSearchParams) WithQuery(Query *string) *ConductSearchParams {
	o.Query = Query
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ConductSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
