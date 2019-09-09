package segment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

//ListTrackingPlans returns all tracking plans for a workspace
func (c *Client) ListTrackingPlans() (TrackingPlans, error) {
	var p TrackingPlans
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
		nil)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return p, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return p, nil
}

// GetTrackingPlan returns information about a tracking plan
func (c *Client) GetTrackingPlan(planName string) (TrackingPlan, error) {
	var p TrackingPlan
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s", WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, planName),
		nil)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return p, errors.Wrap(err, "failed to unmarshal tracking plan response")
	}
	return p, nil
}

// CreateTrackingPlan creates a new tracking plan
func (c *Client) CreateTrackingPlan(displayName string, rules Rules) (trackingPlanCreateRequest, error) {
	var p trackingPlanCreateRequest

	plan := trackingPlanCreateRequest{
		TrackingPlan: TrackingPlan{
			DisplayName: displayName,
			Rules:       rules,
		}}
	data, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
		plan)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return p, errors.Wrap(err, "failed to unmarshall tracking plan response")
	}

	return p, nil

}

// UpdateTrackingPlan upates an existing tracking plan
func (c *Client) UpdateTrackingPlan(planName string, paths []string, updatedPlan TrackingPlan) (TrackingPlan, error) {
	var p TrackingPlan
	req := trackingPlanUpdateRequest{TrackingPlan: updatedPlan, UpdateMask: UpdateMask{Paths: paths}}
	data, err := c.doRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s/%s/", WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, planName), req)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return p, errors.Wrap(err, "failed to unmarshal tracking plan response")
	}

	return p, nil
}
