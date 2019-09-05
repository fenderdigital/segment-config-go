package segment

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTrackingPlans_ListTrackingPlans(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"tracking_plans": [
				{
					"name": "workspace/myworkspace/tracking-plans/rs_123",
					"display_name": "Kicks App",
					"rules": {
						"identify": [],
						"group": [],
						"events": []
					},
					"create_time": "2019-02-05T00:28:31Z",
					"update_time": "2019-02-05T00:28:31Z"
				}
			]
			}`)
	})

	actual, err := client.ListTrackingPlans()
	fmt.Printf("ERROR %+v\n", err)
	assert.NoError(t, err)

	currentTime, _ := time.Parse(time.RFC3339, "2019-02-05T00:28:31Z")
	updatedTime, _ := time.Parse(time.RFC3339, "2019-02-05T00:28:31Z")
	time.Parse(time.RFC3339, "2019-02-05T00:28:31Z")
	expected := TrackingPlans{TrackingPlans: []TrackingPlan{
		{
			Name:        "workspace/myworkspace/tracking-plans/rs_123",
			DisplayName: "Kicks App",
			Rules: Rules{
				Identify: []Rule{},
				Group:    []Rule{},
				Events:   []Event{},
			},
			CreateTime: currentTime,
			UpdateTime: updatedTime,
		},
	}}
	fmt.Printf("ACTUAL %+v\n", actual)
	fmt.Printf("EXPECTED %+v\n", expected)
	assert.Equal(t, expected, actual)
}
