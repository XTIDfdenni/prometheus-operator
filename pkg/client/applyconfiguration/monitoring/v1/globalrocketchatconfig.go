// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// GlobalRocketChatConfigApplyConfiguration represents a declarative configuration of the GlobalRocketChatConfig type for use
// with apply.
type GlobalRocketChatConfigApplyConfiguration struct {
	APIURL  *monitoringv1.URL         `json:"apiURL,omitempty"`
	Token   *corev1.SecretKeySelector `json:"token,omitempty"`
	TokenID *corev1.SecretKeySelector `json:"tokenID,omitempty"`
}

// GlobalRocketChatConfigApplyConfiguration constructs a declarative configuration of the GlobalRocketChatConfig type for use with
// apply.
func GlobalRocketChatConfig() *GlobalRocketChatConfigApplyConfiguration {
	return &GlobalRocketChatConfigApplyConfiguration{}
}

// WithAPIURL sets the APIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIURL field is set to the value of the last call.
func (b *GlobalRocketChatConfigApplyConfiguration) WithAPIURL(value monitoringv1.URL) *GlobalRocketChatConfigApplyConfiguration {
	b.APIURL = &value
	return b
}

// WithToken sets the Token field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Token field is set to the value of the last call.
func (b *GlobalRocketChatConfigApplyConfiguration) WithToken(value corev1.SecretKeySelector) *GlobalRocketChatConfigApplyConfiguration {
	b.Token = &value
	return b
}

// WithTokenID sets the TokenID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TokenID field is set to the value of the last call.
func (b *GlobalRocketChatConfigApplyConfiguration) WithTokenID(value corev1.SecretKeySelector) *GlobalRocketChatConfigApplyConfiguration {
	b.TokenID = &value
	return b
}
