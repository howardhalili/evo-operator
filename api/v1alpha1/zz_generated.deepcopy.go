// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"smartmobilelabs.com/evo/evo-operator/controllers/k8sdynamic"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppReportedData) DeepCopyInto(out *AppReportedData) {
	*out = *in
	if in.PrivateNetworkIpAddress != nil {
		in, out := &in.PrivateNetworkIpAddress, &out.PrivateNetworkIpAddress
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppReportedData.
func (in *AppReportedData) DeepCopy() *AppReportedData {
	if in == nil {
		return nil
	}
	out := new(AppReportedData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvoDomainInfo) DeepCopyInto(out *EvoDomainInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvoDomainInfo.
func (in *EvoDomainInfo) DeepCopy() *EvoDomainInfo {
	if in == nil {
		return nil
	}
	out := new(EvoDomainInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvoPorts) DeepCopyInto(out *EvoPorts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvoPorts.
func (in *EvoPorts) DeepCopy() *EvoPorts {
	if in == nil {
		return nil
	}
	out := new(EvoPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.AdditionalRoutes != nil {
		in, out := &in.AdditionalRoutes, &out.AdditionalRoutes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkAccess) DeepCopyInto(out *PrivateNetworkAccess) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkAccess.
func (in *PrivateNetworkAccess) DeepCopy() *PrivateNetworkAccess {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmlEvo) DeepCopyInto(out *SmlEvo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmlEvo.
func (in *SmlEvo) DeepCopy() *SmlEvo {
	if in == nil {
		return nil
	}
	out := new(SmlEvo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmlEvo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmlEvoList) DeepCopyInto(out *SmlEvoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SmlEvo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmlEvoList.
func (in *SmlEvoList) DeepCopy() *SmlEvoList {
	if in == nil {
		return nil
	}
	out := new(SmlEvoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmlEvoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmlEvoSpec) DeepCopyInto(out *SmlEvoSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(EvoPorts)
		**out = **in
	}
	if in.PrivateNetworkAccess != nil {
		in, out := &in.PrivateNetworkAccess, &out.PrivateNetworkAccess
		*out = new(PrivateNetworkAccess)
		(*in).DeepCopyInto(*out)
	}
	out.Domain = in.Domain
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmlEvoSpec.
func (in *SmlEvoSpec) DeepCopy() *SmlEvoSpec {
	if in == nil {
		return nil
	}
	out := new(SmlEvoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmlEvoStatus) DeepCopyInto(out *SmlEvoStatus) {
	*out = *in
	if in.PrevSpec != nil {
		in, out := &in.PrevSpec, &out.PrevSpec
		*out = new(SmlEvoSpec)
		(*in).DeepCopyInto(*out)
	}
	in.AppReportedData.DeepCopyInto(&out.AppReportedData)
	if in.AppliedResources != nil {
		in, out := &in.AppliedResources, &out.AppliedResources
		*out = make([]k8sdynamic.ResourceDescriptor, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmlEvoStatus.
func (in *SmlEvoStatus) DeepCopy() *SmlEvoStatus {
	if in == nil {
		return nil
	}
	out := new(SmlEvoStatus)
	in.DeepCopyInto(out)
	return out
}
