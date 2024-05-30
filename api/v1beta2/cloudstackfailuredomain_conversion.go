/*
Copyright 2023 The Kubernetes Authors.

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

package v1beta2

import (
	machineryconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *CloudStackFailureDomain) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1beta3.CloudStackFailureDomain)
	return Convert_v1beta2_CloudStackFailureDomain_To_v1beta3_CloudStackFailureDomain(src, dst, nil)
}

func (dst *CloudStackFailureDomain) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1beta3.CloudStackFailureDomain)
	return Convert_v1beta3_CloudStackFailureDomain_To_v1beta2_CloudStackFailureDomain(src, dst, nil)
}

func Convert_v1beta3_CloudStackFailureDomainSpec_To_v1beta2_CloudStackFailureDomainSpec(in *v1beta3.CloudStackFailureDomainSpec, out *CloudStackFailureDomainSpec, s machineryconversion.Scope) error { // nolint
	return autoConvert_v1beta3_CloudStackFailureDomainSpec_To_v1beta2_CloudStackFailureDomainSpec(in, out, s)
}
