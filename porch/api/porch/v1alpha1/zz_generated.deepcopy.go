//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionConfig) DeepCopyInto(out *FunctionConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RequiredFields != nil {
		in, out := &in.RequiredFields, &out.RequiredFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionConfig.
func (in *FunctionConfig) DeepCopy() *FunctionConfig {
	if in == nil {
		return nil
	}
	out := new(FunctionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionEvalTaskSpec) DeepCopyInto(out *FunctionEvalTaskSpec) {
	*out = *in
	if in.FunctionRef != nil {
		in, out := &in.FunctionRef, &out.FunctionRef
		*out = new(FunctionRef)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Config.DeepCopyInto(&out.Config)
	out.Match = in.Match
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionEvalTaskSpec.
func (in *FunctionEvalTaskSpec) DeepCopy() *FunctionEvalTaskSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionEvalTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionRef) DeepCopyInto(out *FunctionRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionRef.
func (in *FunctionRef) DeepCopy() *FunctionRef {
	if in == nil {
		return nil
	}
	out := new(FunctionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	out.RepositoryRef = in.RepositoryRef
	if in.FunctionTypes != nil {
		in, out := &in.FunctionTypes, &out.FunctionTypes
		*out = make([]FunctionType, len(*in))
		copy(*out, *in)
	}
	if in.FunctionConfigs != nil {
		in, out := &in.FunctionConfigs, &out.FunctionConfigs
		*out = make([]FunctionConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitPackage) DeepCopyInto(out *GitPackage) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitPackage.
func (in *GitPackage) DeepCopy() *GitPackage {
	if in == nil {
		return nil
	}
	out := new(GitPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciPackage) DeepCopyInto(out *OciPackage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciPackage.
func (in *OciPackage) DeepCopy() *OciPackage {
	if in == nil {
		return nil
	}
	out := new(OciPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCloneTaskSpec) DeepCopyInto(out *PackageCloneTaskSpec) {
	*out = *in
	in.Upstream.DeepCopyInto(&out.Upstream)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCloneTaskSpec.
func (in *PackageCloneTaskSpec) DeepCopy() *PackageCloneTaskSpec {
	if in == nil {
		return nil
	}
	out := new(PackageCloneTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageInitTaskSpec) DeepCopyInto(out *PackageInitTaskSpec) {
	*out = *in
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageInitTaskSpec.
func (in *PackageInitTaskSpec) DeepCopy() *PackageInitTaskSpec {
	if in == nil {
		return nil
	}
	out := new(PackageInitTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackagePatchTaskSpec) DeepCopyInto(out *PackagePatchTaskSpec) {
	*out = *in
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackagePatchTaskSpec.
func (in *PackagePatchTaskSpec) DeepCopy() *PackagePatchTaskSpec {
	if in == nil {
		return nil
	}
	out := new(PackagePatchTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevision) DeepCopyInto(out *PackageRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevision.
func (in *PackageRevision) DeepCopy() *PackageRevision {
	if in == nil {
		return nil
	}
	out := new(PackageRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionList) DeepCopyInto(out *PackageRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionList.
func (in *PackageRevisionList) DeepCopy() *PackageRevisionList {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionRef) DeepCopyInto(out *PackageRevisionRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionRef.
func (in *PackageRevisionRef) DeepCopy() *PackageRevisionRef {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResources) DeepCopyInto(out *PackageRevisionResources) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResources.
func (in *PackageRevisionResources) DeepCopy() *PackageRevisionResources {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevisionResources) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResourcesList) DeepCopyInto(out *PackageRevisionResourcesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageRevisionResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResourcesList.
func (in *PackageRevisionResourcesList) DeepCopy() *PackageRevisionResourcesList {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResourcesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageRevisionResourcesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionResourcesSpec) DeepCopyInto(out *PackageRevisionResourcesSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionResourcesSpec.
func (in *PackageRevisionResourcesSpec) DeepCopy() *PackageRevisionResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionSpec) DeepCopyInto(out *PackageRevisionSpec) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionSpec.
func (in *PackageRevisionSpec) DeepCopy() *PackageRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionStatus) DeepCopyInto(out *PackageRevisionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionStatus.
func (in *PackageRevisionStatus) DeepCopy() *PackageRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryRef) DeepCopyInto(out *RepositoryRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryRef.
func (in *RepositoryRef) DeepCopy() *RepositoryRef {
	if in == nil {
		return nil
	}
	out := new(RepositoryRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		*out = new(PackageInitTaskSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = new(PackageCloneTaskSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Patch != nil {
		in, out := &in.Patch, &out.Patch
		*out = new(PackagePatchTaskSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Eval != nil {
		in, out := &in.Eval, &out.Eval
		*out = new(FunctionEvalTaskSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamPackage) DeepCopyInto(out *UpstreamPackage) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitPackage)
		**out = **in
	}
	if in.Oci != nil {
		in, out := &in.Oci, &out.Oci
		*out = new(OciPackage)
		**out = **in
	}
	if in.UpstreamRef != nil {
		in, out := &in.UpstreamRef, &out.UpstreamRef
		*out = new(PackageRevisionRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamPackage.
func (in *UpstreamPackage) DeepCopy() *UpstreamPackage {
	if in == nil {
		return nil
	}
	out := new(UpstreamPackage)
	in.DeepCopyInto(out)
	return out
}
