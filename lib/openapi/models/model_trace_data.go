/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type TraceData struct {
	TraceRef                 string     `json:"traceRef" yaml:"traceRef" bson:"traceRef" mapstructure:"TraceRef"`
	TraceDepth               TraceDepth `json:"traceDepth" yaml:"traceDepth" bson:"traceDepth" mapstructure:"TraceDepth"`
	NeTypeList               string     `json:"neTypeList" yaml:"neTypeList" bson:"neTypeList" mapstructure:"NeTypeList"`
	EventList                string     `json:"eventList" yaml:"eventList" bson:"eventList" mapstructure:"EventList"`
	CollectionEntityIpv4Addr string     `json:"collectionEntityIpv4Addr,omitempty" yaml:"collectionEntityIpv4Addr" bson:"collectionEntityIpv4Addr" mapstructure:"CollectionEntityIpv4Addr"`
	CollectionEntityIpv6Addr string     `json:"collectionEntityIpv6Addr,omitempty" yaml:"collectionEntityIpv6Addr" bson:"collectionEntityIpv6Addr" mapstructure:"CollectionEntityIpv6Addr"`
	InterfaceList            string     `json:"interfaceList,omitempty" yaml:"interfaceList" bson:"interfaceList" mapstructure:"InterfaceList"`
}
