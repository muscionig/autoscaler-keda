/*
Copyright 2024 The Knative Authors

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

package resources

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/kmeta"
	"knative.dev/serving/pkg/apis/autoscaling"
	"knative.dev/serving/pkg/autoscaler/config"
	. "knative.dev/serving/pkg/testing" //nolint:all

	hpaconfig "knative.dev/autoscaler-keda/pkg/reconciler/autoscaling/hpa/config"
	"knative.dev/autoscaler-keda/pkg/reconciler/autoscaling/hpa/helpers"
)

func TestDesiredScaledObject(t *testing.T) {
	aConfig, err := config.NewConfigFromMap(nil)
	if err != nil {
		t.Fatalf("Failed to create autoscaler config = %v", err)
	}

	autoscalerKedaConfig, err := hpaconfig.NewConfigFromMap(nil)
	if err != nil {
		t.Fatalf("Failed to create autoscaler keda config = %v", err)
	}

	scaledObjectTests := []struct {
		name             string
		wantErr          bool
		wantScaledObject *kedav1alpha1.ScaledObject
		paAnnotations    map[string]string
	}{{
		name: "cpu metric with default cm values",
		paAnnotations: map[string]string{
			autoscaling.MaxScaleAnnotationKey: "10",
			autoscaling.MetricAnnotationKey:   "cpu",
			autoscaling.TargetAnnotationKey:   "75",
		},
		wantScaledObject: ScaledObject(helpers.TestNamespace,
			helpers.TestRevision, WithAnnotations(map[string]string{
				autoscaling.MaxScaleAnnotationKey: "10",
				autoscaling.MetricAnnotationKey:   "cpu",
				autoscaling.TargetAnnotationKey:   "75",
				autoscaling.ClassAnnotationKey:    autoscaling.HPA,
			}), WithMaxScale(10), WithMinScale(1), WithScaleTargetRef(helpers.TestRevision+"-deployment"),
			WithTrigger("cpu", autoscalingv2.UtilizationMetricType, map[string]string{
				"value": "75",
			}), WithHorizontalPodAutoscalerConfig(helpers.TestRevision)),
	}, {
		name: "cpu metric with default cm values and min scale",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey: "2",
			autoscaling.MaxScaleAnnotationKey: "10",
			autoscaling.MetricAnnotationKey:   "cpu",
			autoscaling.TargetAnnotationKey:   "75",
		},
		wantScaledObject: ScaledObject(helpers.TestNamespace,
			helpers.TestRevision, WithAnnotations(map[string]string{
				autoscaling.MinScaleAnnotationKey: "2",
				autoscaling.MaxScaleAnnotationKey: "10",
				autoscaling.MetricAnnotationKey:   "cpu",
				autoscaling.TargetAnnotationKey:   "75",
				autoscaling.ClassAnnotationKey:    autoscaling.HPA,
			}), WithMaxScale(10), WithMinScale(2), WithScaleTargetRef(helpers.TestRevision+"-deployment"),
			WithTrigger("cpu", autoscalingv2.UtilizationMetricType, map[string]string{
				"value": "75",
			}), WithHorizontalPodAutoscalerConfig(helpers.TestRevision)),
	}, {
		name: "memory metric with default cm values",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey: "1",
			autoscaling.MaxScaleAnnotationKey: "10",
			autoscaling.MetricAnnotationKey:   "memory",
			autoscaling.TargetAnnotationKey:   "200",
		},
		wantScaledObject: ScaledObject(helpers.TestNamespace,
			helpers.TestRevision, WithAnnotations(map[string]string{
				autoscaling.MinScaleAnnotationKey: "1",
				autoscaling.MaxScaleAnnotationKey: "10",
				autoscaling.MetricAnnotationKey:   "memory",
				autoscaling.TargetAnnotationKey:   "200",
				autoscaling.ClassAnnotationKey:    autoscaling.HPA,
			}), WithMaxScale(10), WithMinScale(1), WithScaleTargetRef(helpers.TestRevision+"-deployment"),
			WithTrigger("memory", autoscalingv2.AverageValueMetricType, map[string]string{
				"value": "200Mi",
			}), WithHorizontalPodAutoscalerConfig(helpers.TestRevision)),
	}, {
		name: "custom metric with default cm values",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey:     "1",
			autoscaling.MaxScaleAnnotationKey:     "10",
			autoscaling.MetricAnnotationKey:       "http_requests_total",
			KedaAutoscaleAnotationPrometheusQuery: "sum(rate(http_requests_total{}[1m]))",
			autoscaling.TargetAnnotationKey:       "5",
		},
		wantScaledObject: ScaledObject(helpers.TestNamespace,
			helpers.TestRevision, WithAnnotations(map[string]string{
				autoscaling.MinScaleAnnotationKey:     "1",
				autoscaling.MaxScaleAnnotationKey:     "10",
				autoscaling.MetricAnnotationKey:       "http_requests_total",
				KedaAutoscaleAnotationPrometheusQuery: "sum(rate(http_requests_total{}[1m]))",
				autoscaling.TargetAnnotationKey:       "5",
				autoscaling.ClassAnnotationKey:        autoscaling.HPA,
			}), WithMaxScale(10), WithMinScale(1), WithPrometheusTrigger(map[string]string{
				"namespace":     helpers.TestNamespace,
				"query":         "sum(rate(http_requests_total{}[1m]))",
				"threshold":     "5",
				"serverAddress": "http://prometheus-operated.default.svc:9090",
			}), WithScaleTargetRef(helpers.TestRevision+"-deployment"), WithHorizontalPodAutoscalerConfig(helpers.TestRevision)),
	}, {
		name: "custom metric with bad prometheus address",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey:       "1",
			autoscaling.MaxScaleAnnotationKey:       "10",
			autoscaling.MetricAnnotationKey:         "http_requests_total",
			KedaAutoscaleAnotationPrometheusQuery:   "sum(rate(http_requests_total{}[1m]))",
			KedaAutoscaleAnotationPrometheusAddress: "http//9090",
			autoscaling.TargetAnnotationKey:         "5",
		},
		wantErr: true,
	}, {
		name: "custom metric with bad auth kind",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey:        "1",
			autoscaling.MaxScaleAnnotationKey:        "10",
			autoscaling.MetricAnnotationKey:          "http_requests_total",
			KedaAutoscaleAnotationPrometheusQuery:    "sum(rate(http_requests_total{}[1m]))",
			KedaAutoscaleAnotationPrometheusAuthKind: "TriggerAuth",
			KedaAutoscaleAnotationPrometheusAuthName: "keda-trigger-auth-prometheus",
			autoscaling.TargetAnnotationKey:          "5",
		},
		wantErr: true,
	}, {
		name: "custom metric with no auth name",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey:        "1",
			autoscaling.MaxScaleAnnotationKey:        "10",
			autoscaling.MetricAnnotationKey:          "http_requests_total",
			KedaAutoscaleAnotationPrometheusQuery:    "sum(rate(http_requests_total{}[1m]))",
			KedaAutoscaleAnotationPrometheusAuthKind: "TriggerAuthentication",
			autoscaling.TargetAnnotationKey:          "5",
		},
		wantErr: true,
	}, {
		name: "custom metric with default cm values with authentication",
		paAnnotations: map[string]string{
			autoscaling.MinScaleAnnotationKey:         "1",
			autoscaling.MaxScaleAnnotationKey:         "10",
			autoscaling.MetricAnnotationKey:           "http_requests_total",
			KedaAutoscaleAnotationPrometheusQuery:     "sum(rate(http_requests_total{}[1m]))",
			autoscaling.TargetAnnotationKey:           "5",
			KedaAutoscaleAnotationPrometheusAddress:   "https://thanos-querier.openshift-monitoring.svc.cluster.local:9092",
			KedaAutoscaleAnotationPrometheusAuthName:  "keda-trigger-auth-prometheus",
			KedaAutoscaleAnotationPrometheusAuthKind:  "TriggerAuthentication",
			KedaAutoscaleAnotationPrometheusAuthModes: "bearer",
		},
		wantScaledObject: ScaledObject(helpers.TestNamespace,
			helpers.TestRevision, WithAnnotations(map[string]string{
				autoscaling.MinScaleAnnotationKey:         "1",
				autoscaling.MaxScaleAnnotationKey:         "10",
				autoscaling.MetricAnnotationKey:           "http_requests_total",
				KedaAutoscaleAnotationPrometheusQuery:     "sum(rate(http_requests_total{}[1m]))",
				autoscaling.TargetAnnotationKey:           "5",
				autoscaling.ClassAnnotationKey:            autoscaling.HPA,
				KedaAutoscaleAnotationPrometheusAddress:   "https://thanos-querier.openshift-monitoring.svc.cluster.local:9092",
				KedaAutoscaleAnotationPrometheusAuthName:  "keda-trigger-auth-prometheus",
				KedaAutoscaleAnotationPrometheusAuthKind:  "TriggerAuthentication",
				KedaAutoscaleAnotationPrometheusAuthModes: "bearer",
			}), WithMaxScale(10), WithMinScale(1), WithScaleTargetRef(helpers.TestRevision+"-deployment"),
			WithAuthPrometheusTrigger(map[string]string{
				"query":         "sum(rate(http_requests_total{}[1m]))",
				"namespace":     helpers.TestNamespace,
				"threshold":     "5",
				"serverAddress": "https://thanos-querier.openshift-monitoring.svc.cluster.local:9092",
				"authModes":     "bearer",
			}, "keda-trigger-auth-prometheus", "TriggerAuthentication"),
			WithHorizontalPodAutoscalerConfig(helpers.TestRevision)),
	}}

	for _, tt := range scaledObjectTests {
		t.Run(tt.name, func(t *testing.T) {
			pa := helpers.PodAutoscaler(helpers.TestNamespace, helpers.TestRevision, WithHPAClass, helpers.WithAnnotations(tt.paAnnotations))
			scaledObject, err := DesiredScaledObject(pa, aConfig, autoscalerKedaConfig)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Failed to create desiredScaledObject, error = %v, want: %v", err, tt.wantErr)
			} else if err == nil {
				tt.wantScaledObject.OwnerReferences = []v1.OwnerReference{*kmeta.NewControllerRef(pa)}
				if diff := cmp.Diff(tt.wantScaledObject, scaledObject); diff != "" {
					t.Fatalf("ScaledObject mismatch: diff(-want,+got):\n%s", diff)
				}
			}
		})
	}
}