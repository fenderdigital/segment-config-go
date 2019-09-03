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
