/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type FailureCause string

// List of FailureCause
const (
	RULE_EVENTFailureCause     FailureCause = "PCC_RULE_EVENT"
	QOS_FLOW_EVENTFailureCause FailureCause = "PCC_QOS_FLOW_EVENT"
)
