/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AuthorizedDefaultQos struct {
	Var5qi int32 `json:"5qi,omitempty" bson:"5qi"`

	Arp *Arp `json:"arp,omitempty" bson:"arp"`

	PriorityLevel int32 `json:"priorityLevel,omitempty" bson:"priorityLevel"`

	AverWindow int32 `json:"averWindow,omitempty" bson:"averWindow"`

	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" bson:"maxDataBurstVol"`
}
