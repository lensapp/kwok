//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Metric{}, func(obj interface{}) { SetObjectDefaults_Metric(obj.(*Metric)) })
	scheme.AddTypeDefaultingFunc(&MetricList{}, func(obj interface{}) { SetObjectDefaults_MetricList(obj.(*MetricList)) })
	scheme.AddTypeDefaultingFunc(&Stage{}, func(obj interface{}) { SetObjectDefaults_Stage(obj.(*Stage)) })
	scheme.AddTypeDefaultingFunc(&StageList{}, func(obj interface{}) { SetObjectDefaults_StageList(obj.(*StageList)) })
	return nil
}

func SetObjectDefaults_Metric(in *Metric) {
	for i := range in.Spec.Metrics {
		a := &in.Spec.Metrics[i]
		if a.Dimension == "" {
			a.Dimension = "node"
		}
	}
}

func SetObjectDefaults_MetricList(in *MetricList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Metric(a)
	}
}

func SetObjectDefaults_Stage(in *Stage) {
	if in.Spec.ResourceRef.APIGroup == "" {
		in.Spec.ResourceRef.APIGroup = "v1"
	}
	if in.Spec.Weight == 0 {
		in.Spec.Weight = 0
	}
}

func SetObjectDefaults_StageList(in *StageList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Stage(a)
	}
}
