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

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"tracking_plans": [
				{
					"name": "workspaces/myworkspace/tracking-plans/rs_123",
					"display_name": "Kicks App",
					"rules": {
						"identify_traits": [],
						"group_traits": [],
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

	createTime, _ := time.Parse(time.RFC3339, "2019-02-05T00:28:31Z")
	updatedTime, _ := time.Parse(time.RFC3339, "2019-02-05T00:28:31Z")
	expected := TrackingPlans{TrackingPlans: []TrackingPlan{
		{
			Name:        "workspaces/myworkspace/tracking-plans/rs_123",
			DisplayName: "Kicks App",
			Rules: Rules{
				IdentifyTraits: []Rule{},
				GroupTraits:    []Rule{},
				Events:         []Event{},
			},
			CreateTime: createTime,
			UpdateTime: updatedTime,
		},
	}}
	fmt.Printf("ACTUAL:: %+v\n", actual)
	assert.Equal(t, expected, actual)
}

func TestTrackingPlans_GetTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	testTrackingPlan := "rs_123"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testTrackingPlan)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,
			`{
				"name": "workspaces/myworkspace/tracking-plans/rs_123",
				"display_name": "Kicks App",
				"rules": {
					"identify_traits": [],
			  		"group_traits": [],
					"events": [
						{
							"name": "Order Completed",
							"description": "Who bought what",
							"version": 1,
							"rules": {
							  "$schema": "http://json-schema.org/draft-07/schema#",
							  "type": "object",
							  "properties": {
								"context": {},
								"traits": {},
								"properties": {
								  "required": [
									"product",
									"price",
									"amount"
								  ],
								  "type": "object",
								  "properties": {
									"product": {
									  "type": [
										"string"
									  ]
									},
									"amount": {
									  "type": [
										"number"
									  ]
									},
									"price": {
									  "type": [
										"number"
									  ]
									}
								  }
								}
							  }
							}
						  }
					],
					"global": {
						"$schema": "http://json-schema.org/draft-07/schema#",
						"type": "object",
						"properties": {
						  "context": {
							"required": [
							  "library"
							],
							"type": "object",
							"properties": {
							  "library": {"type": ["object"]}
							}
						  },
						  "traits": {},
						  "properties": {}
						}
					},
					"identify": {
						"$schema": "http://json-schema.org/draft-07/schema#",
						"type": "object",
						"properties": {
						  "traits": {
							"type": "object",
							"properties": {
							  "occupation": {
								"type": [
								  "string"
								]
							  },
							  "age": {
								"type": [
								  "number"
								]
							  },
							  "name": {
								"type": [
								  "string"
								]
							  }
							},
							"required": [
							  "name"
							]
						  },
						  "properties": {},
						  "context": {}
						}
					},
					"group": {
						"$schema": "http://json-schema.org/draft-07/schema#",
						"type": "object",
						"properties": {
						  "properties": {},
						  "context": {},
						  "traits": {
							"properties": {
							  "company": {
								"type": [
								  "object"
								]
							  }
							},
							"required": [
							  "company"
							],
							"type": "object"
						  }
						}
					}
				},
				"create_time": "2019-02-05T01:21:25Z",
				"update_time": "2019-02-05T01:21:25Z"
			}`)
	})

	actual, err := client.GetTrackingPlan(testTrackingPlan)
	assert.NoError(t, err)

	createTime, _ := time.Parse(time.RFC3339, "2019-02-05T01:21:25Z")
	updatedTime, _ := time.Parse(time.RFC3339, "2019-02-05T01:21:25Z")
	expected := TrackingPlan{
		Name:        "workspaces/myworkspace/tracking-plans/rs_123",
		DisplayName: "Kicks App",
		Rules: Rules{
			IdentifyTraits: []Rule{},
			GroupTraits:    []Rule{},
			Events: []Event{
				Event{
					Name:        "Order Completed",
					Description: "Who bought what",
					Version:     1,
					Rules: Rule{
						Schema: "http://json-schema.org/draft-07/schema#",
						Type:   "object",
						Properties: map[string]Rule{
							"context": Rule{},
							"traits":  Rule{},
							"properties": Rule{
								Required: []string{"product", "price", "amount"},
								Type:     "object",
								Properties: map[string]Rule{
									"product": Rule{
										Type: []interface{}{"string"},
									},
									"amount": Rule{
										Type: []interface{}{"number"},
									},
									"price": Rule{
										Type: []interface{}{"number"},
									},
								},
							},
						},
					},
				},
			},
			Global: Rule{
				Schema: "http://json-schema.org/draft-07/schema#",
				Type:   "object",
				Properties: map[string]Rule{
					"context": Rule{
						Required: []string{"library"},
						Type:     "object",
						Properties: map[string]Rule{
							"library": Rule{
								Type: []interface{}{"object"},
							},
						},
					},
					"traits":     Rule{},
					"properties": Rule{},
				},
			},
			Identify: Rule{
				Schema: "http://json-schema.org/draft-07/schema#",
				Type:   "object",
				Properties: map[string]Rule{
					"traits": Rule{
						Type: "object",
						Properties: map[string]Rule{
							"occupation": Rule{
								Type: []interface{}{"string"},
							},
							"age": Rule{
								Type: []interface{}{"number"},
							},
							"name": Rule{
								Type: []interface{}{"string"},
							},
						},
						Required: []string{"name"},
					},
					"properties": Rule{},
					"context":    Rule{},
				},
			},
			Group: Rule{
				Schema: "http://json-schema.org/draft-07/schema#",
				Type:   "object",
				Properties: map[string]Rule{
					"properties": Rule{},
					"context":    Rule{},
					"traits": Rule{
						Properties: map[string]Rule{
							"company": Rule{
								Type: []interface{}{"object"},
							},
						},
						Required: []string{"company"},
						Type:     "object",
					},
				},
			},
		},
		CreateTime: createTime,
		UpdateTime: updatedTime,
	}
	fmt.Printf("ACTUAL:: %+v\n", actual)
	assert.Equal(t, expected, actual)
}

func TestTrackingPlans_CreateTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	testDisplayName := "Kicks App"
	testRules := Rules{
		Events: []Event{
			Event{
				Name:        "Order Completed",
				Description: "Who bought what",
				Version:     1,
				Rules: Rule{
					Schema: "http://json-schema.org/draft-07/schema#",
					Type:   "object",
					Properties: map[string]Rule{
						"context": Rule{},
						"traits":  Rule{},
						"properties": Rule{
							Required: []string{"product", "price", "amount"},
							Type:     "object",
							Properties: map[string]Rule{
								"product": Rule{
									Type: []interface{}{"string"},
								},
								"amount": Rule{
									Type: []interface{}{"number"},
								},
								"price": Rule{
									Type: []interface{}{"number"},
								},
							},
						},
					},
				},
			},
		},
		Global: Rule{
			Schema: "http://json-schema.org/draft-07/schema#",
			Type:   "object",
			Properties: map[string]Rule{
				"context": Rule{
					Required: []string{"library"},
					Type:     "object",
					Properties: map[string]Rule{
						"library": Rule{
							Type: []interface{}{"object"},
						},
					},
				},
				"traits":     Rule{},
				"properties": Rule{},
			},
		},
		Identify: Rule{
			Schema: "http://json-schema.org/draft-07/schema#",
			Type:   "object",
			Properties: map[string]Rule{
				"traits": Rule{
					Type: "object",
					Properties: map[string]Rule{
						"occupation": Rule{
							Type: []interface{}{"string"},
						},
						"age": Rule{
							Type: []interface{}{"number"},
						},
						"name": Rule{
							Type: []interface{}{"string"},
						},
					},
					Required: []string{"name"},
				},
				"properties": Rule{},
				"context":    Rule{},
			},
		},
		Group: Rule{
			Schema: "http://json-schema.org/draft-07/schema#",
			Type:   "object",
			Properties: map[string]Rule{
				"properties": Rule{},
				"context":    Rule{},
				"traits": Rule{
					Properties: map[string]Rule{
						"company": Rule{
							Type: []interface{}{"object"},
						},
					},
					Required: []string{"company"},
					Type:     "object",
				},
			},
		},
	}

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/",
		apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"tracking_plan": {
			"display_name": "Kicks App",
			"rules": {
				"events": [
					{
						"name": "Order Completed",
						"description": "Who bought what",
						"version": 1,
						"rules": {
						  "$schema": "http://json-schema.org/draft-07/schema#",
						  "type": "object",
						  "properties": {
							"context": {},
							"traits": {},
							"properties": {
							  "required": [
								"product",
								"price",
								"amount"
							  ],
							  "type": "object",
							  "properties": {
								"product": {
								  "type": [
									"string"
								  ]
								},
								"amount": {
								  "type": [
									"number"
								  ]
								},
								"price": {
								  "type": [
									"number"
								  ]
								}
							  }
							}
						  }
						}
					  }
				],
				"global": {
					"$schema": "http://json-schema.org/draft-07/schema#",
					"type": "object",
					"properties": {
					  "context": {
						"required": [
						  "library"
						],
						"type": "object",
						"properties": {
						  "library": {"type": ["object"]}
						}
					  },
					  "traits": {},
					  "properties": {}
					}
				},
				"identify": {
					"$schema": "http://json-schema.org/draft-07/schema#",
					"type": "object",
					"properties": {
					  "traits": {
						"type": "object",
						"properties": {
						  "occupation": {
							"type": [
							  "string"
							]
						  },
						  "age": {
							"type": [
							  "number"
							]
						  },
						  "name": {
							"type": [
							  "string"
							]
						  }
						},
						"required": [
						  "name"
						]
					  },
					  "properties": {},
					  "context": {}
					}
				},
				"group": {
					"$schema": "http://json-schema.org/draft-07/schema#",
					"type": "object",
					"properties": {
					  "properties": {},
					  "context": {},
					  "traits": {
						"properties": {
						  "company": {
							"type": [
							  "object"
							]
						  }
						},
						"required": [
						  "company"
						],
						"type": "object"
					  }
					}
				}
			}
		  }
		}`)
	})

	expected := trackingPlanCreateRequest{
		TrackingPlan: TrackingPlan{
			DisplayName: testDisplayName,
			Rules:       testRules},
	}

	actual, err := client.CreateTrackingPlan(testDisplayName, testRules)

	assert.NoError(t, err)

	assert.Equal(t, expected, actual)

}

func TestTrackingPlan_UpdateTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	testPlanName := "rs_123"
	createTime, _ := time.Parse(time.RFC3339, "2019-02-05T00:28:31Z")
	updatedTime, _ := time.Parse(time.RFC3339, "2019-02-05T00:32:15Z")
	testUpdatedPlan := TrackingPlan{
		Name:        "workspaces/myworkspace/tracking-plans/rs_123",
		DisplayName: "Kicks App - Updated",
		Rules: Rules{
			IdentifyTraits: []Rule{},
			GroupTraits:    []Rule{},
			Events: []Event{
				{
					Name:        "Product Viewed",
					Version:     1,
					Description: "Who checked out what",
					Rules: Rule{
						Schema: "http://json-schema.org/draft-04/schema#",
						Type:   "object",
						Properties: map[string]Rule{
							"traits": Rule{},
							"properties": Rule{
								Type: "object",
								Properties: map[string]Rule{
									"product": Rule{
										Type: []interface{}{"string"},
									},
								},
								Required: []string{"product"},
							},
							"context": Rule{},
						},
						Required: []string{"properties"},
					},
				},
			},
			Global: Rule{
				Schema: "http://json-schema.org/draft-04/schema#",
				Type:   "object",
				Properties: map[string]Rule{
					"context": Rule{
						Type: "object",
						Properties: map[string]Rule{
							"userAgent": {},
						},
						Required: []string{"userAgent"},
					},
					"traits":     Rule{},
					"properties": Rule{},
				},
				Required: []string{"context"},
			},
		},
		CreateTime: createTime,
		UpdateTime: updatedTime,
	}
	testPaths := []string{"tracking_plan.display_name", "tracking_plan.rules"}

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testPlanName)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"name": "workspaces/myworkspace/tracking-plans/rs_123",
			"display_name": "Kicks App - Updated",
			"rules": {
			  "identify_traits": [],
			  "group_traits": [],
			  "events": [
				{
				  "name": "Product Viewed",
				  "version": 1,
				  "description": "Who checked out what",
				  "rules": {
					"$schema": "http://json-schema.org/draft-04/schema#",
					"type": "object",
					"properties": {
					  "traits": {},
					  "properties": {
						"type": "object",
						"properties": {
						  "product": {
							"type": [
							  "string"
							]
						  }
						},
						"required": [
						  "product"
						]
					  },
					  "context": {}
					},
					"required": [
					  "properties"
					]
				  }
				}
			  ],
			  "global": {
				"required": [
				  "context"
				],
				"$schema": "http://json-schema.org/draft-04/schema#",
				"type": "object",
				"properties": {
				  "context": {
					"type": "object",
					"properties": {
					  "userAgent": {}
					},
					"required": [
					  "userAgent"
					]
				  },
				  "traits": {},
				  "properties": {}
				}
			  }
			},
			"create_time": "2019-02-05T00:28:31Z",
			"update_time": "2019-02-05T00:32:15Z"
		  }`)
	})

	expected := testUpdatedPlan

	actual, err := client.UpdateTrackingPlan(testPlanName, testPaths, testUpdatedPlan)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestTrackingPlan_CreateTrackingPlanSourceConnection(t *testing.T) {
	setup()
	defer teardown()

	testPlanName := "rs_123"
	testSrcName := "js"
	testPlanID := "rs_1Gjkh9ZKmpyHjdSZYaLTXRRgCPp"
	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace,
		TrackingPlanEndpoint, testPlanName, TrackingPlanSourceConnectionEndpoint)
	fmt.Printf("METHOD ENDPOINT: %+v\n", endpoint)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"source_name": "workspaces/test-workspace/sources/js",
			"tracking_plan_id": "rs_1Gjkh9ZKmpyHjdSZYaLTXRRgCPp"
		}`)
	})

	sourcePath := fmt.Sprintf("%s/%s/%s/%s", WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSrcName)
	expected := trackingPlanSourceConnection{
		SourceName:     sourcePath,
		TrackingPlanID: testPlanID,
	}

	actual, err := client.CreateTrackingPlanSourceConnection(testPlanName, testSrcName)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestTrackingPlan_ListTrackingPlanSourceConnections(t *testing.T) {
	setup()
	defer teardown()

	testPlanName := "rs_123"
	testSrcName := "js"
	testPlanID := "rs_1Gjkh9ZKmpyHjdSZYaLTXRRgCPp"

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testPlanName, TrackingPlanSourceConnectionEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"connections": [
			  {
				"source_name": "workspaces/test-workspace/sources/js",
				"tracking_plan_id": "rs_1Gjkh9ZKmpyHjdSZYaLTXRRgCPp"
			  }
			]
		}`)
	})

	sourcePath := fmt.Sprintf("%s/%s/%s/%s", WorkspacesEndpoint, testWorkspace, SourceEndpoint, testSrcName)
	expected := trackingPlanSourceConnections{
		Connections: []trackingPlanSourceConnection{
			{
				SourceName:     sourcePath,
				TrackingPlanID: testPlanID,
			},
		},
	}

	actual, err := client.ListTrackingPlanSourceConnections(testPlanName)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestTrackingPlan_DeleteTrackingPlan(t *testing.T) {
	setup()
	defer teardown()

	testPlanName := "rs_123"

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint, testPlanName)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {})

	err := client.DeleteTrackingPlan(testPlanName)
	assert.NoError(t, err)
}

func TestTrackingPlan_DeleteTrackingPlanSourceConnection(t *testing.T) {
	setup()
	defer teardown()

	testPlanName := "rs_123"
	srcName := "js"

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint,
		testPlanName, TrackingPlanSourceConnectionEndpoint, srcName)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {})

	err := client.DeleteTrackingPlanSourceConnection(testPlanName, srcName)
	assert.NoError(t, err)
}
