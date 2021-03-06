package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListOrgInvoicesParams creates a new ListOrgInvoicesParams object
// with the default values initialized.
func NewListOrgInvoicesParams() *ListOrgInvoicesParams {
	var ()
	return &ListOrgInvoicesParams{}
}

/*ListOrgInvoicesParams contains all the parameters to send to the API endpoint
for the list org invoices operation typically these are written to a http.Request
*/
type ListOrgInvoicesParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
}

// WithOrgname adds the orgname to the list org invoices params
func (o *ListOrgInvoicesParams) WithOrgname(Orgname string) *ListOrgInvoicesParams {
	o.Orgname = Orgname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrgInvoicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
