package main

import (
	"fmt"
	"strings"

	"github.com/fenderdigital/segment-apis-go/segment"
)

func main() {
	//rules, err := ioutil.ReadFile("/Users/respinoza/Projects/data-team-operations/segment/projects/poc_terraform/create_tracking_plan/kick2_app.json")
	//if err != nil {
	//	fmt.Print(err)
	//}

	client := segment.NewClient("T1ijVmLSCqpI0i7Q6OuK7_cmF-qOLHipLIJcTBSTHNo.AVnUOWRwBDkq09TLwtzGCZ8TFgTId71dxOIOonVCBCs", "fender")
	//displayName := "TEST2"
	//s := segment.Rules{}
	//json.Unmarshal([]byte(rules), &s)
	//trackingPlan, err := client.CreateTrackingPlan(displayName, s)
	plans, _ := client.ListTrackingPlans()
	names := make(map[string]string)
	for _, element := range plans.TrackingPlans {
		nameSplit := strings.Split(element.Name, "/")
		id := nameSplit[len(nameSplit)-1]
		names[id] = element.Name
	}

	if _, ok := names["rs_1ONyQomh1iWuODcOPMvxz6gnKMm"]; !ok {
		fmt.Println("FAILED")
	} else {
		fmt.Println("SUCCESS")
	}

}

func getNameIDs(plan segment.TrackingPlan) string {
	name := plan.Name
	nameSplit := strings.Split(name, "/")
	return nameSplit[len(nameSplit)-1]
}
