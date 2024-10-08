//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package pub

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateContainerBatch":     schema_kruise_apis_apps_pub_InPlaceUpdateContainerBatch(ref),
		"github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateContainerStatus":    schema_kruise_apis_apps_pub_InPlaceUpdateContainerStatus(ref),
		"github.com/openkruise/kruise-api/apps/pub.InPlaceUpdatePreCheckBeforeNext": schema_kruise_apis_apps_pub_InPlaceUpdatePreCheckBeforeNext(ref),
		"github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateState":              schema_kruise_apis_apps_pub_InPlaceUpdateState(ref),
		"github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateStrategy":           schema_kruise_apis_apps_pub_InPlaceUpdateStrategy(ref),
		"github.com/openkruise/kruise-api/apps/pub.Lifecycle":                       schema_kruise_apis_apps_pub_Lifecycle(ref),
		"github.com/openkruise/kruise-api/apps/pub.LifecycleHook":                   schema_kruise_apis_apps_pub_LifecycleHook(ref),
		"github.com/openkruise/kruise-api/apps/pub.RuntimeContainerHashes":          schema_kruise_apis_apps_pub_RuntimeContainerHashes(ref),
		"github.com/openkruise/kruise-api/apps/pub.RuntimeContainerMeta":            schema_kruise_apis_apps_pub_RuntimeContainerMeta(ref),
		"github.com/openkruise/kruise-api/apps/pub.RuntimeContainerMetaSet":         schema_kruise_apis_apps_pub_RuntimeContainerMetaSet(ref),
		"github.com/openkruise/kruise-api/apps/pub.UpdatePriorityOrderTerm":         schema_kruise_apis_apps_pub_UpdatePriorityOrderTerm(ref),
		"github.com/openkruise/kruise-api/apps/pub.UpdatePriorityStrategy":          schema_kruise_apis_apps_pub_UpdatePriorityStrategy(ref),
		"github.com/openkruise/kruise-api/apps/pub.UpdatePriorityWeightTerm":        schema_kruise_apis_apps_pub_UpdatePriorityWeightTerm(ref),
	}
}

func schema_kruise_apis_apps_pub_InPlaceUpdateContainerBatch(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InPlaceUpdateContainerBatch indicates the timestamp and containers for a batch update",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"timestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "Timestamp is the time for this update batch",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"containers": {
						SchemaProps: spec.SchemaProps{
							Description: "Containers is the name list of containers for this update batch",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"timestamp", "containers"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_kruise_apis_apps_pub_InPlaceUpdateContainerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InPlaceUpdateContainerStatus records the statuses of the container that are mainly used to determine whether the InPlaceUpdate is completed.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"imageID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_kruise_apis_apps_pub_InPlaceUpdatePreCheckBeforeNext(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InPlaceUpdatePreCheckBeforeNext contains the pre-check that must pass before the next containers can be in-place update.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"containersRequiredReady": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_kruise_apis_apps_pub_InPlaceUpdateState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InPlaceUpdateState records latest inplace-update state, including old statuses of containers.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"revision": {
						SchemaProps: spec.SchemaProps{
							Description: "Revision is the updated revision hash.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"updateTimestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "UpdateTimestamp is the start time when the in-place update happens.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastContainerStatuses": {
						SchemaProps: spec.SchemaProps{
							Description: "LastContainerStatuses records the before-in-place-update container statuses. It is a map from ContainerName to InPlaceUpdateContainerStatus",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateContainerStatus"),
									},
								},
							},
						},
					},
					"updateEnvFromMetadata": {
						SchemaProps: spec.SchemaProps{
							Description: "UpdateEnvFromMetadata indicates there are envs from annotations/labels that should be in-place update.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"nextContainerImages": {
						SchemaProps: spec.SchemaProps{
							Description: "NextContainerImages is the containers with lower priority that waiting for in-place update images in next batch.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"nextContainerRefMetadata": {
						SchemaProps: spec.SchemaProps{
							Description: "NextContainerRefMetadata is the containers with lower priority that waiting for in-place update labels/annotations in next batch.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
									},
								},
							},
						},
					},
					"preCheckBeforeNext": {
						SchemaProps: spec.SchemaProps{
							Description: "PreCheckBeforeNext is the pre-check that must pass before the next containers can be in-place update.",
							Ref:         ref("github.com/openkruise/kruise-api/apps/pub.InPlaceUpdatePreCheckBeforeNext"),
						},
					},
					"containerBatchesRecord": {
						SchemaProps: spec.SchemaProps{
							Description: "ContainerBatchesRecord records the update batches that have patched in this revision.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateContainerBatch"),
									},
								},
							},
						},
					},
				},
				Required: []string{"revision", "updateTimestamp", "lastContainerStatuses"},
			},
		},
		Dependencies: []string{
			"github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateContainerBatch", "github.com/openkruise/kruise-api/apps/pub.InPlaceUpdateContainerStatus", "github.com/openkruise/kruise-api/apps/pub.InPlaceUpdatePreCheckBeforeNext", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_kruise_apis_apps_pub_InPlaceUpdateStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InPlaceUpdateStrategy defines the strategies for in-place update.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"gracePeriodSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "GracePeriodSeconds is the timespan between set Pod status to not-ready and update images in Pod spec when in-place update a Pod.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
	}
}

func schema_kruise_apis_apps_pub_Lifecycle(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Lifecycle contains the hooks for Pod lifecycle.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"preDelete": {
						SchemaProps: spec.SchemaProps{
							Description: "PreDelete is the hook before Pod to be deleted.",
							Ref:         ref("github.com/openkruise/kruise-api/apps/pub.LifecycleHook"),
						},
					},
					"inPlaceUpdate": {
						SchemaProps: spec.SchemaProps{
							Description: "InPlaceUpdate is the hook before Pod to update and after Pod has been updated.",
							Ref:         ref("github.com/openkruise/kruise-api/apps/pub.LifecycleHook"),
						},
					},
					"preNormal": {
						SchemaProps: spec.SchemaProps{
							Description: "PreNormal is the hook after Pod to be created and ready to be Normal.",
							Ref:         ref("github.com/openkruise/kruise-api/apps/pub.LifecycleHook"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openkruise/kruise-api/apps/pub.LifecycleHook"},
	}
}

func schema_kruise_apis_apps_pub_LifecycleHook(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"labelsHandler": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"finalizersHandler": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"markPodNotReady": {
						SchemaProps: spec.SchemaProps{
							Description: "MarkPodNotReady = true means: - Pod will be set to 'NotReady' at preparingDelete/preparingUpdate state. - Pod will be restored to 'Ready' at Updated state if it was set to 'NotReady' at preparingUpdate state. Currently, MarkPodNotReady only takes effect on InPlaceUpdate & PreDelete hook. Default to false.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_kruise_apis_apps_pub_RuntimeContainerHashes(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RuntimeContainerHashes contains the hashes of such container.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"plainHash": {
						SchemaProps: spec.SchemaProps{
							Description: "PlainHash is the hash that directly calculated from pod.spec.container[x]. Usually it is calculated by Kubelet and will be in annotation of each runtime container.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"extractedEnvFromMetadataHash": {
						SchemaProps: spec.SchemaProps{
							Description: "ExtractedEnvFromMetadataHash is the hash that calculated from pod.spec.container[x], whose envs from annotations/labels have already been extracted to the real values.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
				},
				Required: []string{"plainHash"},
			},
		},
	}
}

func schema_kruise_apis_apps_pub_RuntimeContainerMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RuntimeContainerMeta contains the meta data of a runtime container.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"containerID": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"restartCount": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"hashes": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/openkruise/kruise-api/apps/pub.RuntimeContainerHashes"),
						},
					},
				},
				Required: []string{"name", "containerID", "restartCount", "hashes"},
			},
		},
		Dependencies: []string{
			"github.com/openkruise/kruise-api/apps/pub.RuntimeContainerHashes"},
	}
}

func schema_kruise_apis_apps_pub_RuntimeContainerMetaSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RuntimeContainerMetaSet contains all the containers' meta of the Pod.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"containers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/openkruise/kruise-api/apps/pub.RuntimeContainerMeta"),
									},
								},
							},
						},
					},
				},
				Required: []string{"containers"},
			},
		},
		Dependencies: []string{
			"github.com/openkruise/kruise-api/apps/pub.RuntimeContainerMeta"},
	}
}

func schema_kruise_apis_apps_pub_UpdatePriorityOrderTerm(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UpdatePriorityOrderTerm defines order priority.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"orderedKey": {
						SchemaProps: spec.SchemaProps{
							Description: "Calculate priority by value of this key. Values of this key, will be sorted by GetInt(val). GetInt method will find the last int in value, such as getting 5 in value '5', getting 10 in value 'sts-10'.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"orderedKey"},
			},
		},
	}
}

func schema_kruise_apis_apps_pub_UpdatePriorityStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UpdatePriorityStrategy is the strategy to define priority for pods update. Only one of orderPriority and weightPriority can be set.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"orderPriority": {
						SchemaProps: spec.SchemaProps{
							Description: "Order priority terms, pods will be sorted by the value of orderedKey. For example: ``` orderPriority: - orderedKey: key1 - orderedKey: key2 ``` First, all pods which have key1 in labels will be sorted by the value of key1. Then, the left pods which have no key1 but have key2 in labels will be sorted by the value of key2 and put behind those pods have key1.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/openkruise/kruise-api/apps/pub.UpdatePriorityOrderTerm"),
									},
								},
							},
						},
					},
					"weightPriority": {
						SchemaProps: spec.SchemaProps{
							Description: "Weight priority terms, pods will be sorted by the sum of all terms weight.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/openkruise/kruise-api/apps/pub.UpdatePriorityWeightTerm"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openkruise/kruise-api/apps/pub.UpdatePriorityOrderTerm", "github.com/openkruise/kruise-api/apps/pub.UpdatePriorityWeightTerm"},
	}
}

func schema_kruise_apis_apps_pub_UpdatePriorityWeightTerm(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UpdatePriorityWeightTerm defines weight priority.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"weight": {
						SchemaProps: spec.SchemaProps{
							Description: "Weight associated with matching the corresponding matchExpressions, in the range 1-100.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"matchSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "MatchSelector is used to select by pod's labels.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
				Required: []string{"weight", "matchSelector"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}
