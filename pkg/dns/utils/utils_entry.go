/*
 * Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 *
 */

package utils

import (
	"reflect"
	"time"

	"github.com/gardener/controller-manager-library/pkg/resources"
	"github.com/gardener/external-dns-management/pkg/dns"

	api "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
)

var _ DNSSpecification = (*DNSEntryObject)(nil)

var DNSEntryType = (*api.DNSEntry)(nil)

type DNSEntryObject struct {
	resources.Object
}

func (this *DNSEntryObject) DNSEntry() *api.DNSEntry {
	return this.Data().(*api.DNSEntry)
}

func DNSEntry(o resources.Object) *DNSEntryObject {
	if o.IsA(DNSEntryType) {
		return &DNSEntryObject{o}
	}
	return nil
}

func (this *DNSEntryObject) Spec() *api.DNSEntrySpec {
	return &this.DNSEntry().Spec
}

func (this *DNSEntryObject) StatusField() interface{} {
	return this.Status()
}

func (this *DNSEntryObject) Status() *api.DNSEntryStatus {
	return &this.DNSEntry().Status
}

func (this *DNSEntryObject) BaseStatus() *api.DNSBaseStatus {
	return &this.DNSEntry().Status.DNSBaseStatus
}

func GetDNSName(entry *api.DNSEntry) string {
	return dns.NormalizeHostname(entry.Spec.DNSName)
}

func (this *DNSEntryObject) GetDNSName() string {
	return GetDNSName(this.DNSEntry())
}

func (this *DNSEntryObject) GetSetIdentifier() string {
	if policy := this.DNSEntry().Spec.RoutingPolicy; policy != nil {
		return policy.SetIdentifier
	}
	return ""
}

func (this *DNSEntryObject) GetTargets() []string {
	return this.DNSEntry().Spec.Targets
}

func (this *DNSEntryObject) GetText() []string {
	return this.DNSEntry().Spec.Text
}

func (this *DNSEntryObject) GetOwnerId() *string {
	return this.DNSEntry().Spec.OwnerId
}

func (this *DNSEntryObject) GetTTL() *int64 {
	return this.DNSEntry().Spec.TTL
}

func (this *DNSEntryObject) GetCNameLookupInterval() *int64 {
	return this.DNSEntry().Spec.CNameLookupInterval
}

func (this *DNSEntryObject) GetReference() *api.EntryReference {
	return this.DNSEntry().Spec.Reference
}

func (this *DNSEntryObject) GetRoutingPolicy() *dns.RoutingPolicy {
	if policy := this.DNSEntry().Spec.RoutingPolicy; policy != nil {
		return &dns.RoutingPolicy{
			Type:       policy.Type,
			Parameters: policy.Parameters,
		}
	}
	return nil
}

func (this *DNSEntryObject) RefreshTime() time.Time {
	return time.Time{}
}

func (this *DNSEntryObject) ValidateSpecial() error {
	return nil
}

func (this *DNSEntryObject) AcknowledgeTargets(targets []string) bool {
	s := this.Status()
	if !reflect.DeepEqual(s.Targets, targets) {
		s.Targets = targets
		return true
	}
	return false
}

func (this *DNSEntryObject) AcknowledgeRoutingPolicy(policy *dns.RoutingPolicy) bool {
	s := this.Status()
	if s.RoutingPolicy == nil && policy == nil {
		return false
	}
	if policy == nil {
		s.RoutingPolicy = nil
		return true
	}
	statusPolicy := &api.RoutingPolicy{
		Type:          policy.Type,
		SetIdentifier: this.GetSetIdentifier(),
		Parameters:    policy.Parameters,
	}
	if !reflect.DeepEqual(s.RoutingPolicy, statusPolicy) {
		s.RoutingPolicy = statusPolicy
		return true
	}
	return false
}

func (this *DNSEntryObject) GetTargetSpec(p TargetProvider) TargetSpec {
	return BaseTargetSpec(this, p)
}

func DNSSetName(entry *api.DNSEntry) dns.DNSSetName {
	setIdentifier := ""
	if policy := entry.Spec.RoutingPolicy; policy != nil {
		setIdentifier = policy.SetIdentifier
	}
	return dns.DNSSetName{
		DNSName:       GetDNSName(entry),
		SetIdentifier: setIdentifier,
	}
}

func (this *DNSEntryObject) DNSSetName() dns.DNSSetName {
	return DNSSetName(this.DNSEntry())
}

func DNSSetNameMatcher(name dns.DNSSetName) resources.ObjectMatcher {
	return func(o resources.Object) bool {
		return DNSEntry(o).DNSSetName() == name
	}
}
