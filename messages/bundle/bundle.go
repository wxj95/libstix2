// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package bundle

import (
	"github.com/freetaxii/libstix2/messages/campaign"
	"github.com/freetaxii/libstix2/messages/indicator"
	"github.com/freetaxii/libstix2/messages/malware"
	"github.com/freetaxii/libstix2/messages/relationship"
	"github.com/freetaxii/libstix2/messages/sighting"
	"github.com/freetaxii/libstix2/messages/stix"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type BundleType struct {
	MessageType   string                          `json:"type,omitempty"`
	Id            string                          `json:"id,omitempty"`
	Spec_version  string                          `json:"spec_version,omitempty"`
	Campaigns     []campaign.CampaignType         `json:"campaigns,omitempty"`
	Indicators    []indicator.IndicatorType       `json:"indicators,omitempty"`
	Malware       []malware.MalwareType           `json:"malware,omitempty"`
	Relationships []relationship.RelationshipType `json:"relationships,omitempty"`
	Sightings     []sighting.SightingType         `json:"sightings,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() BundleType {
	var obj BundleType
	obj.MessageType = "bundle"
	obj.Id = stix.NewId("bundle")
	obj.SetSpecVersion20()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------
func (this *BundleType) SetSpecVersion20() {
	this.Spec_version = "stix-2.0"
}

// ----------------------------------------------------------------------
// Public Methods - BundleType
// ----------------------------------------------------------------------

func (this *BundleType) NewCampaign() *campaign.CampaignType {
	i := campaign.New()
	slicePosition := this.addCampaign(i)
	return &this.Campaigns[slicePosition]
}

func (this *BundleType) NewIndicator() *indicator.IndicatorType {
	i := indicator.New()
	slicePosition := this.addIndicator(i)
	return &this.Indicators[slicePosition]
}

func (this *BundleType) NewMalware() *malware.MalwareType {
	i := malware.New()
	slicePosition := this.addMalware(i)
	return &this.Malware[slicePosition]
}

func (this *BundleType) NewRelationship() *relationship.RelationshipType {
	i := relationship.New()
	slicePosition := this.addRelationship(i)
	return &this.Relationships[slicePosition]
}

func (this *BundleType) NewSighting() *sighting.SightingType {
	i := sighting.New()
	slicePosition := this.addSighting(i)
	return &this.Sightings[slicePosition]
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *BundleType) addCampaign(o campaign.CampaignType) int {
	if this.Campaigns == nil {
		a := make([]campaign.CampaignType, 0)
		this.Campaigns = a
	}
	positionThatAppendWillUse := len(this.Campaigns)
	this.Campaigns = append(this.Campaigns, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addIndicator(o indicator.IndicatorType) int {
	if this.Indicators == nil {
		a := make([]indicator.IndicatorType, 0)
		this.Indicators = a
	}
	positionThatAppendWillUse := len(this.Indicators)
	this.Indicators = append(this.Indicators, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addMalware(o malware.MalwareType) int {
	if this.Malware == nil {
		a := make([]malware.MalwareType, 0)
		this.Malware = a
	}
	positionThatAppendWillUse := len(this.Malware)
	this.Malware = append(this.Malware, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addRelationship(o relationship.RelationshipType) int {
	if this.Relationships == nil {
		a := make([]relationship.RelationshipType, 0)
		this.Relationships = a
	}
	positionThatAppendWillUse := len(this.Relationships)
	this.Relationships = append(this.Relationships, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addSighting(o sighting.SightingType) int {
	if this.Sightings == nil {
		a := make([]sighting.SightingType, 0)
		this.Sightings = a
	}
	positionThatAppendWillUse := len(this.Sightings)
	this.Sightings = append(this.Sightings, o)
	return positionThatAppendWillUse
}
