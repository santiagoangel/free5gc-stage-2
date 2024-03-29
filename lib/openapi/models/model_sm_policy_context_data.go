/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmPolicyContextData struct {
	AccNetChId *AccNetChId `json:"accNetChId,omitempty" bson:"accNetChId"`

	ChargEntityAddr *AccNetChargingAddress `json:"chargEntityAddr,omitempty" bson:"chargEntityAddr"`

	Gpsi string `json:"gpsi,omitempty" bson:"gpsi"`

	Supi string `json:"supi" bson:"supi"`

	InterGrpIds []string `json:"interGrpIds,omitempty" bson:"interGrpIds"`

	PduSessionId int32 `json:"pduSessionId" bson:"pduSessionId"`

	PduSessionType PduSessionType `json:"pduSessionType" bson:"pduSessionType"`

	Chargingcharacteristics string `json:"chargingcharacteristics,omitempty" bson:"chargingcharacteristics"`

	Dnn string `json:"dnn" bson:"dnn"`

	NotificationUri string `json:"notificationUri" bson:"notificationUri"`

	AccessType AccessType `json:"accessType,omitempty" bson:"accessType"`

	RatType RatType `json:"ratType,omitempty" bson:"ratType"`

	ServingNetwork *NetworkId `json:"servingNetwork,omitempty" bson:"servingNetwork"`

	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty" bson:"userLocationInfo"`

	UeTimeZone string `json:"ueTimeZone,omitempty" bson:"ueTimeZone"`

	Pei string `json:"pei,omitempty" bson:"pei"`

	Ipv4Address string `json:"ipv4Address,omitempty" bson:"ipv4Address"`

	Ipv6AddressPrefix string `json:"ipv6AddressPrefix,omitempty" bson:"ipv6AddressPrefix"`

	// Indicates the IPv4 address domain
	IpDomain string `json:"ipDomain,omitempty" bson:"ipDomain"`

	SubsSessAmbr *Ambr `json:"subsSessAmbr,omitempty" bson:"subsSessAmbr"`

	SubsDefQos *SubscribedDefaultQos `json:"subsDefQos,omitempty" bson:"subsDefQos"`

	// Contains the number of supported packet filter for signalled QoS rules.
	NumOfPackFilter int32 `json:"numOfPackFilter,omitempty" bson:"numOfPackFilter"`

	// If it is included and set to true, the online charging is applied to the PDU session.
	Online bool `json:"online,omitempty" bson:"online"`

	// If it is included and set to true, the offline charging is applied to the PDU session.
	Offline bool `json:"offline,omitempty" bson:"offline"`

	// If it is included and set to true, the 3GPP PS Data Off is activated by the UE.
	Var3gppPsDataOffStatus bool `json:"3gppPsDataOffStatus,omitempty" bson:"3gppPsDataOffStatus"`

	// If it is included and set to true, the reflective QoS is supported by the UE.
	RefQosIndication bool `json:"refQosIndication,omitempty" bson:"refQosIndication"`

	TraceReq *TraceData `json:"traceReq,omitempty" bson:"traceReq"`

	SliceInfo *Snssai `json:"sliceInfo" bson:"sliceInfo"`

	QosFlowUsage QosFlowUsage `json:"qosFlowUsage,omitempty" bson:"qosFlowUsage"`

	SuppFeat string `json:"suppFeat,omitempty" bson:"suppFeat"`
}
