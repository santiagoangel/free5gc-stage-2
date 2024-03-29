/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type SmPolicyUpdateContextData struct {

	// The policy control reqeust trigges which are met.
	RepPolicyCtrlReqTriggers []PolicyControlRequestTrigger `json:"repPolicyCtrlReqTriggers,omitempty" bson:"repPolicyCtrlReqTriggers"`

	// Indicates the access network charging identifier for the PCC rule(s) or whole PDU session.
	AccNetChIds []AccNetChId `json:"accNetChIds,omitempty" bson:"accNetChIds"`

	AccessType AccessType `json:"accessType,omitempty" bson:"accessType"`

	RatType RatType `json:"ratType,omitempty" bson:"ratType"`

	ServingNetwork *NetworkId `json:"servingNetwork,omitempty" bson:"servingNetwork"`

	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty" bson:"userLocationInfo"`

	UeTimeZone string `json:"ueTimeZone,omitempty" bson:"ueTimeZone"`

	RelIpv4Address string `json:"relIpv4Address,omitempty" bson:"relIpv4Address"`

	Ipv4Address string `json:"ipv4Address,omitempty" bson:"ipv4Address"`

	Ipv6AddressPrefix string `json:"ipv6AddressPrefix,omitempty" bson:"ipv6AddressPrefix"`

	RelIpv6AddressPrefix string `json:"relIpv6AddressPrefix,omitempty" bson:"relIpv6AddressPrefix"`

	RelUeMac string `json:"relUeMac,omitempty" bson:"relUeMac"`

	UeMac string `json:"ueMac,omitempty" bson:"ueMac"`

	SubsSessAmbr *Ambr `json:"subsSessAmbr,omitempty" bson:"subsSessAmbr"`

	SubsDefQos *SubscribedDefaultQos `json:"subsDefQos,omitempty" bson:"subsDefQos"`

	// Contains the number of supported packet filter for signalled QoS rules.
	NumOfPackFilter int32 `json:"numOfPackFilter,omitempty" bson:"numOfPackFilter"`

	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty" bson:"accuUsageReports"`

	// If it is included and set to true, the 3GPP PS Data Off is activated by the UE.
	Var3gppPsDataOffStatus bool `json:"3gppPsDataOffStatus,omitempty" bson:"3gppPsDataOffStatus"`

	// Report the start/stop of the application traffic and detected SDF descriptions if applicable.
	AppDetectionInfos []AppDetectionInfo `json:"appDetectionInfos,omitempty" bson:"appDetectionInfos"`

	// Used to report the PCC rule failure.
	RuleReports []RuleReport `json:"ruleReports,omitempty" bson:"ruleReports"`

	// QoS Notification Control information.
	QncReports []QosNotificationControlInfo `json:"qncReports,omitempty" bson:"qncReports"`

	UserLocationInfoTime *time.Time `json:"userLocationInfoTime,omitempty" bson:"userLocationInfoTime"`

	// Reports the changes of presence reporting area.
	RepPraInfos map[string]PresenceInfo `json:"repPraInfos,omitempty" bson:"repPraInfos"`

	UeInitResReq *UeInitiatedResourceRequest `json:"ueInitResReq,omitempty" bson:"ueInitResReq"`

	// If it is included and set to true, the reflective QoS is supported by the UE. If it is included and set to false, the reflective QoS is revoked by the UE.
	RefQosIndication bool `json:"refQosIndication,omitempty" bson:"refQosIndication"`

	QosFlowUsage QosFlowUsage `json:"qosFlowUsage,omitempty" bson:"qosFlowUsage"`

	TraceReq *TraceData `json:"traceReq,omitempty" bson:"traceReq"`
}
