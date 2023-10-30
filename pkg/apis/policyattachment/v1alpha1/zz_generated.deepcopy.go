//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRPCRateLimit) DeepCopyInto(out *GRPCRateLimit) {
	*out = *in
	in.Match.DeepCopyInto(&out.Match)
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(L7RateLimit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCRateLimit.
func (in *GRPCRateLimit) DeepCopy() *GRPCRateLimit {
	if in == nil {
		return nil
	}
	out := new(GRPCRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRateLimit) DeepCopyInto(out *HTTPRateLimit) {
	*out = *in
	in.Match.DeepCopyInto(&out.Match)
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(L7RateLimit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRateLimit.
func (in *HTTPRateLimit) DeepCopy() *HTTPRateLimit {
	if in == nil {
		return nil
	}
	out := new(HTTPRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostnameRateLimit) DeepCopyInto(out *HostnameRateLimit) {
	*out = *in
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(L7RateLimit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostnameRateLimit.
func (in *HostnameRateLimit) DeepCopy() *HostnameRateLimit {
	if in == nil {
		return nil
	}
	out := new(HostnameRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7RateLimit) DeepCopyInto(out *L7RateLimit) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(RateLimitPolicyMode)
		**out = **in
	}
	if in.Backlog != nil {
		in, out := &in.Backlog, &out.Backlog
		*out = new(int)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int)
		**out = **in
	}
	if in.ResponseStatusCode != nil {
		in, out := &in.ResponseStatusCode, &out.ResponseStatusCode
		*out = new(int)
		**out = **in
	}
	if in.ResponseHeadersToAdd != nil {
		in, out := &in.ResponseHeadersToAdd, &out.ResponseHeadersToAdd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7RateLimit.
func (in *L7RateLimit) DeepCopy() *L7RateLimit {
	if in == nil {
		return nil
	}
	out := new(L7RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRateLimit) DeepCopyInto(out *PortRateLimit) {
	*out = *in
	if in.BPS != nil {
		in, out := &in.BPS, &out.BPS
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRateLimit.
func (in *PortRateLimit) DeepCopy() *PortRateLimit {
	if in == nil {
		return nil
	}
	out := new(PortRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicy) DeepCopyInto(out *RateLimitPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicy.
func (in *RateLimitPolicy) DeepCopy() *RateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyList) DeepCopyInto(out *RateLimitPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyList.
func (in *RateLimitPolicyList) DeepCopy() *RateLimitPolicyList {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicySpec) DeepCopyInto(out *RateLimitPolicySpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultBPS != nil {
		in, out := &in.DefaultBPS, &out.DefaultBPS
		*out = new(int64)
		**out = **in
	}
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]HostnameRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPRateLimits != nil {
		in, out := &in.HTTPRateLimits, &out.HTTPRateLimits
		*out = make([]HTTPRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GRPCRateLimits != nil {
		in, out := &in.GRPCRateLimits, &out.GRPCRateLimits
		*out = make([]GRPCRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultL7RateLimit != nil {
		in, out := &in.DefaultL7RateLimit, &out.DefaultL7RateLimit
		*out = new(L7RateLimit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicySpec.
func (in *RateLimitPolicySpec) DeepCopy() *RateLimitPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyStatus) DeepCopyInto(out *RateLimitPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyStatus.
func (in *RateLimitPolicyStatus) DeepCopy() *RateLimitPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteRateLimitConfig) DeepCopyInto(out *RouteRateLimitConfig) {
	*out = *in
	if in.HttpRateLimits != nil {
		in, out := &in.HttpRateLimits, &out.HttpRateLimits
		*out = make([]HTTPRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GrpcRateLimits != nil {
		in, out := &in.GrpcRateLimits, &out.GrpcRateLimits
		*out = make([]GRPCRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultRateLimit != nil {
		in, out := &in.DefaultRateLimit, &out.DefaultRateLimit
		*out = new(L7RateLimit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteRateLimitConfig.
func (in *RouteRateLimitConfig) DeepCopy() *RouteRateLimitConfig {
	if in == nil {
		return nil
	}
	out := new(RouteRateLimitConfig)
	in.DeepCopyInto(out)
	return out
}
