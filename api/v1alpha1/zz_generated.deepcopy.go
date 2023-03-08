//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
MIT License

Copyright (c) 2022 ngrok, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	if in.CNAMETarget != nil {
		in, out := &in.CNAMETarget, &out.CNAMETarget
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointCompression) DeepCopyInto(out *EndpointCompression) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointCompression.
func (in *EndpointCompression) DeepCopy() *EndpointCompression {
	if in == nil {
		return nil
	}
	out := new(EndpointCompression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointHeaders) DeepCopyInto(out *EndpointHeaders) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(EndpointRequestHeaders)
		(*in).DeepCopyInto(*out)
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(EndpointResponseHeaders)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointHeaders.
func (in *EndpointHeaders) DeepCopy() *EndpointHeaders {
	if in == nil {
		return nil
	}
	out := new(EndpointHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointIPPolicy) DeepCopyInto(out *EndpointIPPolicy) {
	*out = *in
	if in.IPPolicies != nil {
		in, out := &in.IPPolicies, &out.IPPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointIPPolicy.
func (in *EndpointIPPolicy) DeepCopy() *EndpointIPPolicy {
	if in == nil {
		return nil
	}
	out := new(EndpointIPPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointRequestHeaders) DeepCopyInto(out *EndpointRequestHeaders) {
	*out = *in
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointRequestHeaders.
func (in *EndpointRequestHeaders) DeepCopy() *EndpointRequestHeaders {
	if in == nil {
		return nil
	}
	out := new(EndpointRequestHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointResponseHeaders) DeepCopyInto(out *EndpointResponseHeaders) {
	*out = *in
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointResponseHeaders.
func (in *EndpointResponseHeaders) DeepCopy() *EndpointResponseHeaders {
	if in == nil {
		return nil
	}
	out := new(EndpointResponseHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointTLSTerminationAtEdge) DeepCopyInto(out *EndpointTLSTerminationAtEdge) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointTLSTerminationAtEdge.
func (in *EndpointTLSTerminationAtEdge) DeepCopy() *EndpointTLSTerminationAtEdge {
	if in == nil {
		return nil
	}
	out := new(EndpointTLSTerminationAtEdge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointWebhookVerification) DeepCopyInto(out *EndpointWebhookVerification) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SecretKeyRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointWebhookVerification.
func (in *EndpointWebhookVerification) DeepCopy() *EndpointWebhookVerification {
	if in == nil {
		return nil
	}
	out := new(EndpointWebhookVerification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSEdge) DeepCopyInto(out *HTTPSEdge) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSEdge.
func (in *HTTPSEdge) DeepCopy() *HTTPSEdge {
	if in == nil {
		return nil
	}
	out := new(HTTPSEdge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPSEdge) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSEdgeList) DeepCopyInto(out *HTTPSEdgeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPSEdge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSEdgeList.
func (in *HTTPSEdgeList) DeepCopy() *HTTPSEdgeList {
	if in == nil {
		return nil
	}
	out := new(HTTPSEdgeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPSEdgeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSEdgeRouteSpec) DeepCopyInto(out *HTTPSEdgeRouteSpec) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
	in.Backend.DeepCopyInto(&out.Backend)
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(EndpointCompression)
		**out = **in
	}
	if in.IPRestriction != nil {
		in, out := &in.IPRestriction, &out.IPRestriction
		*out = new(EndpointIPPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(EndpointHeaders)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookVerification != nil {
		in, out := &in.WebhookVerification, &out.WebhookVerification
		*out = new(EndpointWebhookVerification)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSEdgeRouteSpec.
func (in *HTTPSEdgeRouteSpec) DeepCopy() *HTTPSEdgeRouteSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPSEdgeRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSEdgeRouteStatus) DeepCopyInto(out *HTTPSEdgeRouteStatus) {
	*out = *in
	out.Backend = in.Backend
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSEdgeRouteStatus.
func (in *HTTPSEdgeRouteStatus) DeepCopy() *HTTPSEdgeRouteStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPSEdgeRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSEdgeSpec) DeepCopyInto(out *HTTPSEdgeSpec) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
	if in.Hostports != nil {
		in, out := &in.Hostports, &out.Hostports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]HTTPSEdgeRouteSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TLSTermination != nil {
		in, out := &in.TLSTermination, &out.TLSTermination
		*out = new(EndpointTLSTerminationAtEdge)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSEdgeSpec.
func (in *HTTPSEdgeSpec) DeepCopy() *HTTPSEdgeSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPSEdgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSEdgeStatus) DeepCopyInto(out *HTTPSEdgeStatus) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]HTTPSEdgeRouteStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSEdgeStatus.
func (in *HTTPSEdgeStatus) DeepCopy() *HTTPSEdgeStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPSEdgeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPolicy) DeepCopyInto(out *IPPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPolicy.
func (in *IPPolicy) DeepCopy() *IPPolicy {
	if in == nil {
		return nil
	}
	out := new(IPPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPolicyList) DeepCopyInto(out *IPPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPolicyList.
func (in *IPPolicyList) DeepCopy() *IPPolicyList {
	if in == nil {
		return nil
	}
	out := new(IPPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPolicyRule) DeepCopyInto(out *IPPolicyRule) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPolicyRule.
func (in *IPPolicyRule) DeepCopy() *IPPolicyRule {
	if in == nil {
		return nil
	}
	out := new(IPPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPolicyRuleStatus) DeepCopyInto(out *IPPolicyRuleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPolicyRuleStatus.
func (in *IPPolicyRuleStatus) DeepCopy() *IPPolicyRuleStatus {
	if in == nil {
		return nil
	}
	out := new(IPPolicyRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPolicySpec) DeepCopyInto(out *IPPolicySpec) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]IPPolicyRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPolicySpec.
func (in *IPPolicySpec) DeepCopy() *IPPolicySpec {
	if in == nil {
		return nil
	}
	out := new(IPPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPolicyStatus) DeepCopyInto(out *IPPolicyStatus) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]IPPolicyRuleStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPolicyStatus.
func (in *IPPolicyStatus) DeepCopy() *IPPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(IPPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NgrokModuleSet) DeepCopyInto(out *NgrokModuleSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Modules.DeepCopyInto(&out.Modules)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NgrokModuleSet.
func (in *NgrokModuleSet) DeepCopy() *NgrokModuleSet {
	if in == nil {
		return nil
	}
	out := new(NgrokModuleSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NgrokModuleSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NgrokModuleSetList) DeepCopyInto(out *NgrokModuleSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NgrokModuleSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NgrokModuleSetList.
func (in *NgrokModuleSetList) DeepCopy() *NgrokModuleSetList {
	if in == nil {
		return nil
	}
	out := new(NgrokModuleSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NgrokModuleSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NgrokModuleSetModules) DeepCopyInto(out *NgrokModuleSetModules) {
	*out = *in
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(EndpointCompression)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(EndpointHeaders)
		(*in).DeepCopyInto(*out)
	}
	if in.IPRestriction != nil {
		in, out := &in.IPRestriction, &out.IPRestriction
		*out = new(EndpointIPPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSTermination != nil {
		in, out := &in.TLSTermination, &out.TLSTermination
		*out = new(EndpointTLSTerminationAtEdge)
		**out = **in
	}
	if in.WebhookVerification != nil {
		in, out := &in.WebhookVerification, &out.WebhookVerification
		*out = new(EndpointWebhookVerification)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NgrokModuleSetModules.
func (in *NgrokModuleSetModules) DeepCopy() *NgrokModuleSetModules {
	if in == nil {
		return nil
	}
	out := new(NgrokModuleSetModules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPEdge) DeepCopyInto(out *TCPEdge) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPEdge.
func (in *TCPEdge) DeepCopy() *TCPEdge {
	if in == nil {
		return nil
	}
	out := new(TCPEdge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TCPEdge) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPEdgeList) DeepCopyInto(out *TCPEdgeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TCPEdge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPEdgeList.
func (in *TCPEdgeList) DeepCopy() *TCPEdgeList {
	if in == nil {
		return nil
	}
	out := new(TCPEdgeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TCPEdgeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPEdgeSpec) DeepCopyInto(out *TCPEdgeSpec) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
	in.Backend.DeepCopyInto(&out.Backend)
	if in.IPRestriction != nil {
		in, out := &in.IPRestriction, &out.IPRestriction
		*out = new(EndpointIPPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPEdgeSpec.
func (in *TCPEdgeSpec) DeepCopy() *TCPEdgeSpec {
	if in == nil {
		return nil
	}
	out := new(TCPEdgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPEdgeStatus) DeepCopyInto(out *TCPEdgeStatus) {
	*out = *in
	if in.Hostports != nil {
		in, out := &in.Hostports, &out.Hostports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Backend = in.Backend
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPEdgeStatus.
func (in *TCPEdgeStatus) DeepCopy() *TCPEdgeStatus {
	if in == nil {
		return nil
	}
	out := new(TCPEdgeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tunnel) DeepCopyInto(out *Tunnel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tunnel.
func (in *Tunnel) DeepCopy() *Tunnel {
	if in == nil {
		return nil
	}
	out := new(Tunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tunnel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelGroupBackend) DeepCopyInto(out *TunnelGroupBackend) {
	*out = *in
	out.ngrokAPICommon = in.ngrokAPICommon
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelGroupBackend.
func (in *TunnelGroupBackend) DeepCopy() *TunnelGroupBackend {
	if in == nil {
		return nil
	}
	out := new(TunnelGroupBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelGroupBackendStatus) DeepCopyInto(out *TunnelGroupBackendStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelGroupBackendStatus.
func (in *TunnelGroupBackendStatus) DeepCopy() *TunnelGroupBackendStatus {
	if in == nil {
		return nil
	}
	out := new(TunnelGroupBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelList) DeepCopyInto(out *TunnelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tunnel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelList.
func (in *TunnelList) DeepCopy() *TunnelList {
	if in == nil {
		return nil
	}
	out := new(TunnelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TunnelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelSpec) DeepCopyInto(out *TunnelSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelSpec.
func (in *TunnelSpec) DeepCopy() *TunnelSpec {
	if in == nil {
		return nil
	}
	out := new(TunnelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelStatus) DeepCopyInto(out *TunnelStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelStatus.
func (in *TunnelStatus) DeepCopy() *TunnelStatus {
	if in == nil {
		return nil
	}
	out := new(TunnelStatus)
	in.DeepCopyInto(out)
	return out
}
