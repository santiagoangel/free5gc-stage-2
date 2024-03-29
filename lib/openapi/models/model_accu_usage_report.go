/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AccuUsageReport struct {

	// An id referencing UsageMonitoringData objects associated with this usage report.
	RefUmIds string `json:"refUmIds" bson:"refUmIds"`

	// Unsigned integer identifying a volume in units of bytes.
	VolUsage int64 `json:"volUsage,omitempty" bson:"volUsage"`

	// Unsigned integer identifying a volume in units of bytes.
	VolUsageUplink int64 `json:"volUsageUplink,omitempty" bson:"volUsageUplink"`

	// Unsigned integer identifying a volume in units of bytes.
	VolUsageDownlink int64 `json:"volUsageDownlink,omitempty" bson:"volUsageDownlink"`

	TimeUsage int32 `json:"timeUsage,omitempty" bson:"timeUsage"`

	// Unsigned integer identifying a volume in units of bytes.
	NextVolUsage int64 `json:"nextVolUsage,omitempty" bson:"nextVolUsage"`

	// Unsigned integer identifying a volume in units of bytes.
	NextVolUsageUplink int64 `json:"nextVolUsageUplink,omitempty" bson:"nextVolUsageUplink"`

	// Unsigned integer identifying a volume in units of bytes.
	NextVolUsageDownlink int64 `json:"nextVolUsageDownlink,omitempty" bson:"nextVolUsageDownlink"`

	NextTimeUsage int32 `json:"nextTimeUsage,omitempty" bson:"nextTimeUsage"`
}
