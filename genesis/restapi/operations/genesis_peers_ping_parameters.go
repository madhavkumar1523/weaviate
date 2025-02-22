//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/semi-technologies/weaviate/genesis/models"
)

// NewGenesisPeersPingParams creates a new GenesisPeersPingParams object
// no default values defined in spec.
func NewGenesisPeersPingParams() GenesisPeersPingParams {
	return GenesisPeersPingParams{}
}

// GenesisPeersPingParams contains all the bound params for the genesis peers ping operation
// typically these are obtained from a http.Request
//
// swagger:parameters genesis.peers.ping
type GenesisPeersPingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Request Body
	  Required: true
	  In: body
	*/
	Body *models.PeerPing
	/*Name of the Weaviate peer
	  Required: true
	  In: path
	*/
	PeerID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGenesisPeersPingParams() beforehand.
func (o *GenesisPeersPingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PeerPing
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	rPeerID, rhkPeerID, _ := route.Params.GetOK("peerId")
	if err := o.bindPeerID(rPeerID, rhkPeerID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPeerID binds and validates parameter PeerID from path.
func (o *GenesisPeersPingParams) bindPeerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("peerId", "path", "strfmt.UUID", raw)
	}
	o.PeerID = *(value.(*strfmt.UUID))

	if err := o.validatePeerID(formats); err != nil {
		return err
	}

	return nil
}

// validatePeerID carries on validations for parameter PeerID
func (o *GenesisPeersPingParams) validatePeerID(formats strfmt.Registry) error {
	if err := validate.FormatOf("peerId", "path", "uuid", o.PeerID.String(), formats); err != nil {
		return err
	}
	return nil
}
