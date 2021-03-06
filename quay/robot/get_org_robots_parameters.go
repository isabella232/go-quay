package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrgRobotsParams creates a new GetOrgRobotsParams object
// with the default values initialized.
func NewGetOrgRobotsParams() *GetOrgRobotsParams {
	var ()
	return &GetOrgRobotsParams{}
}

/*GetOrgRobotsParams contains all the parameters to send to the API endpoint
for the get org robots operation typically these are written to a http.Request
*/
type GetOrgRobotsParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Permissions
	  Whether to include repostories and teams in which the robots have permission.

	*/
	Permissions *bool
}

// WithOrgname adds the orgname to the get org robots params
func (o *GetOrgRobotsParams) WithOrgname(Orgname string) *GetOrgRobotsParams {
	o.Orgname = Orgname
	return o
}

// WithPermissions adds the permissions to the get org robots params
func (o *GetOrgRobotsParams) WithPermissions(Permissions *bool) *GetOrgRobotsParams {
	o.Permissions = Permissions
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgRobotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if o.Permissions != nil {

		// query param permissions
		var qrPermissions bool
		if o.Permissions != nil {
			qrPermissions = *o.Permissions
		}
		qPermissions := swag.FormatBool(qrPermissions)
		if qPermissions != "" {
			if err := r.SetQueryParam("permissions", qPermissions); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
