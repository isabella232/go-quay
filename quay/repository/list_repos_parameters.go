package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListReposParams creates a new ListReposParams object
// with the default values initialized.
func NewListReposParams() *ListReposParams {
	var ()
	return &ListReposParams{}
}

/*ListReposParams contains all the parameters to send to the API endpoint
for the list repos operation typically these are written to a http.Request
*/
type ListReposParams struct {

	/*LastModified
	  Whether to include when the repository was last modified.

	*/
	LastModified *bool
	/*Namespace
	  Filters the repositories returned to this namespace

	*/
	Namespace *string
	/*NextPage
	  The page token for the next page

	*/
	NextPage *string
	/*Popularity
	  Whether to include the repository's popularity metric.

	*/
	Popularity *bool
	/*Public
	  Adds any repositories visible to the user by virtue of being public

	*/
	Public *bool
	/*Starred
	  Filters the repositories returned to those starred by the user

	*/
	Starred *bool
}

// WithLastModified adds the lastModified to the list repos params
func (o *ListReposParams) WithLastModified(LastModified *bool) *ListReposParams {
	o.LastModified = LastModified
	return o
}

// WithNamespace adds the namespace to the list repos params
func (o *ListReposParams) WithNamespace(Namespace *string) *ListReposParams {
	o.Namespace = Namespace
	return o
}

// WithNextPage adds the nextPage to the list repos params
func (o *ListReposParams) WithNextPage(NextPage *string) *ListReposParams {
	o.NextPage = NextPage
	return o
}

// WithPopularity adds the popularity to the list repos params
func (o *ListReposParams) WithPopularity(Popularity *bool) *ListReposParams {
	o.Popularity = Popularity
	return o
}

// WithPublic adds the public to the list repos params
func (o *ListReposParams) WithPublic(Public *bool) *ListReposParams {
	o.Public = Public
	return o
}

// WithStarred adds the starred to the list repos params
func (o *ListReposParams) WithStarred(Starred *bool) *ListReposParams {
	o.Starred = Starred
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.LastModified != nil {

		// query param last_modified
		var qrLastModified bool
		if o.LastModified != nil {
			qrLastModified = *o.LastModified
		}
		qLastModified := swag.FormatBool(qrLastModified)
		if qLastModified != "" {
			if err := r.SetQueryParam("last_modified", qLastModified); err != nil {
				return err
			}
		}

	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string
		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {
			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}

	}

	if o.NextPage != nil {

		// query param next_page
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("next_page", qNextPage); err != nil {
				return err
			}
		}

	}

	if o.Popularity != nil {

		// query param popularity
		var qrPopularity bool
		if o.Popularity != nil {
			qrPopularity = *o.Popularity
		}
		qPopularity := swag.FormatBool(qrPopularity)
		if qPopularity != "" {
			if err := r.SetQueryParam("popularity", qPopularity); err != nil {
				return err
			}
		}

	}

	if o.Public != nil {

		// query param public
		var qrPublic bool
		if o.Public != nil {
			qrPublic = *o.Public
		}
		qPublic := swag.FormatBool(qrPublic)
		if qPublic != "" {
			if err := r.SetQueryParam("public", qPublic); err != nil {
				return err
			}
		}

	}

	if o.Starred != nil {

		// query param starred
		var qrStarred bool
		if o.Starred != nil {
			qrStarred = *o.Starred
		}
		qStarred := swag.FormatBool(qrStarred)
		if qStarred != "" {
			if err := r.SetQueryParam("starred", qStarred); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
