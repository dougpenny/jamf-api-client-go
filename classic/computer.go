package classic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Computers returns all enrolled computer devices
func (j *Client) Computers() (*ComputerList, error) {
	ep := fmt.Sprintf("%s/%s", j.Endpoint, computersContext)
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building JAMF computer query request")
	}

	res := &ComputerList{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computers from %s", ep)
	}
	return res, nil
}

// ComputerDetails returns the details for a specific computer given its ID
func (j *Client) ComputerDetails(identifier interface{}) (*Computer, error) {
	ep, err := EndpointBuilder(j.Endpoint, computersContext, identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF query request endpoint for computer: %v", identifier)
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", ep, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error building JAMF computer request for computer: %v (%s)", identifier, ep)
	}

	res := &Computer{}
	if err := j.makeAPIrequest(req, &res); err != nil {
		return nil, errors.Wrapf(err, "unable to query enrolled computer for computer: %v (%s)", identifier, ep)
	}
	return res, nil
}