package smf_context

import (
	"free5gc/lib/pfcp/pfcpType"
)

const (
	RULE_INITIAL RuleState = 0
	RULE_CREATE  RuleState = 1
	RULE_UPDATE  RuleState = 2
	RULE_REMOVE  RuleState = 3
)

type RuleState uint8

// Packet Detection Rule. Table 7.5.2.2-1
type PDR struct {
	PDRID uint16

	Precedence         uint32
	PDI                PDI
	OuterHeaderRemoval *pfcpType.OuterHeaderRemoval

	FAR *FAR
	URR *URR
	QER *QER

	State RuleState
}

// Packet Detection. 7.5.2.2-2
type PDI struct {
	SourceInterface pfcpType.SourceInterface
	LocalFTeid      pfcpType.FTEID
	NetworkInstance pfcpType.NetworkInstance
	UEIPAddress     *pfcpType.UEIPAddress
}

// Forwarding Action Rule. 7.5.2.3-1
type FAR struct {
	FARID uint32

	ApplyAction          pfcpType.ApplyAction
	ForwardingParameters *ForwardingParameters

	BAR   *BAR
	State RuleState
}

// Forwarding Parameters. 7.5.2.3-2
type ForwardingParameters struct {
	DestinationInterface pfcpType.DestinationInterface
	NetworkInstance      pfcpType.NetworkInstance
	OuterHeaderCreation  *pfcpType.OuterHeaderCreation
}

// Buffering Action Rule
type BAR struct {
}

// QoS Enhancement Rule
type QER struct {
}

// Usage Report Rule
type URR struct {
}
