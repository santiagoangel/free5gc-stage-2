/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UpPathChgEvent struct {
	NotificationUri string `json:"notificationUri" bson:"notificationUri"`

	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.
	NotifCorreId string `json:"notifCorreId" bson:"notifCorreId"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType" bson:"dnaiChgType"`
}
