// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /Users/aloaugus/src/vpp-api-source/vpp_build//build-root/install-vpp-native/vpp/share/vpp/api//plugins/flowprobe.api.json

/*
Package flowprobe is a generated VPP binary API for 'flowprobe' module.

It consists of:
	  8 enums
	  1 alias
	  4 messages
	  2 services
*/
package flowprobe

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "flowprobe"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xe904179b
)

// FlowprobeRecordFlags represents VPP binary API enum 'flowprobe_record_flags'.
type FlowprobeRecordFlags uint8

const (
	FLOWPROBE_RECORD_FLAG_L2 FlowprobeRecordFlags = 1
	FLOWPROBE_RECORD_FLAG_L3 FlowprobeRecordFlags = 2
	FLOWPROBE_RECORD_FLAG_L4 FlowprobeRecordFlags = 4
)

var FlowprobeRecordFlags_name = map[uint8]string{
	1: "FLOWPROBE_RECORD_FLAG_L2",
	2: "FLOWPROBE_RECORD_FLAG_L3",
	4: "FLOWPROBE_RECORD_FLAG_L4",
}

var FlowprobeRecordFlags_value = map[string]uint8{
	"FLOWPROBE_RECORD_FLAG_L2": 1,
	"FLOWPROBE_RECORD_FLAG_L3": 2,
	"FLOWPROBE_RECORD_FLAG_L4": 4,
}

func (x FlowprobeRecordFlags) String() string {
	s, ok := FlowprobeRecordFlags_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// FlowprobeWhichFlags represents VPP binary API enum 'flowprobe_which_flags'.
type FlowprobeWhichFlags uint8

const (
	FLOWPROBE_WHICH_FLAG_IP4 FlowprobeWhichFlags = 1
	FLOWPROBE_WHICH_FLAG_L2  FlowprobeWhichFlags = 2
	FLOWPROBE_WHICH_FLAG_IP6 FlowprobeWhichFlags = 4
)

var FlowprobeWhichFlags_name = map[uint8]string{
	1: "FLOWPROBE_WHICH_FLAG_IP4",
	2: "FLOWPROBE_WHICH_FLAG_L2",
	4: "FLOWPROBE_WHICH_FLAG_IP6",
}

var FlowprobeWhichFlags_value = map[string]uint8{
	"FLOWPROBE_WHICH_FLAG_IP4": 1,
	"FLOWPROBE_WHICH_FLAG_L2":  2,
	"FLOWPROBE_WHICH_FLAG_IP6": 4,
}

func (x FlowprobeWhichFlags) String() string {
	s, ok := FlowprobeWhichFlags_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfStatusFlags represents VPP binary API enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var IfStatusFlags_name = map[uint32]string{
	1: "IF_STATUS_API_FLAG_ADMIN_UP",
	2: "IF_STATUS_API_FLAG_LINK_UP",
}

var IfStatusFlags_value = map[string]uint32{
	"IF_STATUS_API_FLAG_ADMIN_UP": 1,
	"IF_STATUS_API_FLAG_LINK_UP":  2,
}

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfType represents VPP binary API enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 0
	IF_API_TYPE_SUB      IfType = 1
	IF_API_TYPE_P2P      IfType = 2
	IF_API_TYPE_PIPE     IfType = 3
)

var IfType_name = map[uint32]string{
	0: "IF_API_TYPE_HARDWARE",
	1: "IF_API_TYPE_SUB",
	2: "IF_API_TYPE_P2P",
	3: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 0,
	"IF_API_TYPE_SUB":      1,
	"IF_API_TYPE_P2P":      2,
	"IF_API_TYPE_PIPE":     3,
}

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LinkDuplex represents VPP binary API enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var LinkDuplex_name = map[uint32]string{
	0: "LINK_DUPLEX_API_UNKNOWN",
	1: "LINK_DUPLEX_API_HALF",
	2: "LINK_DUPLEX_API_FULL",
}

var LinkDuplex_value = map[string]uint32{
	"LINK_DUPLEX_API_UNKNOWN": 0,
	"LINK_DUPLEX_API_HALF":    1,
	"LINK_DUPLEX_API_FULL":    2,
}

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 0
	MTU_PROTO_API_IP4  MtuProto = 1
	MTU_PROTO_API_IP6  MtuProto = 2
	MTU_PROTO_API_MPLS MtuProto = 3
)

var MtuProto_name = map[uint32]string{
	0: "MTU_PROTO_API_L3",
	1: "MTU_PROTO_API_IP4",
	2: "MTU_PROTO_API_IP6",
	3: "MTU_PROTO_API_MPLS",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   0,
	"MTU_PROTO_API_IP4":  1,
	"MTU_PROTO_API_IP6":  2,
	"MTU_PROTO_API_MPLS": 3,
}

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RxMode represents VPP binary API enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var RxMode_name = map[uint32]string{
	0: "RX_MODE_API_UNKNOWN",
	1: "RX_MODE_API_POLLING",
	2: "RX_MODE_API_INTERRUPT",
	3: "RX_MODE_API_ADAPTIVE",
	4: "RX_MODE_API_DEFAULT",
}

var RxMode_value = map[string]uint32{
	"RX_MODE_API_UNKNOWN":   0,
	"RX_MODE_API_POLLING":   1,
	"RX_MODE_API_INTERRUPT": 2,
	"RX_MODE_API_ADAPTIVE":  3,
	"RX_MODE_API_DEFAULT":   4,
}

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SubIfFlags represents VPP binary API enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var SubIfFlags_name = map[uint32]string{
	1:   "SUB_IF_API_FLAG_NO_TAGS",
	2:   "SUB_IF_API_FLAG_ONE_TAG",
	4:   "SUB_IF_API_FLAG_TWO_TAGS",
	8:   "SUB_IF_API_FLAG_DOT1AD",
	16:  "SUB_IF_API_FLAG_EXACT_MATCH",
	32:  "SUB_IF_API_FLAG_DEFAULT",
	64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
	128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
	254: "SUB_IF_API_FLAG_MASK_VNET",
	256: "SUB_IF_API_FLAG_DOT1AH",
}

var SubIfFlags_value = map[string]uint32{
	"SUB_IF_API_FLAG_NO_TAGS":           1,
	"SUB_IF_API_FLAG_ONE_TAG":           2,
	"SUB_IF_API_FLAG_TWO_TAGS":          4,
	"SUB_IF_API_FLAG_DOT1AD":            8,
	"SUB_IF_API_FLAG_EXACT_MATCH":       16,
	"SUB_IF_API_FLAG_DEFAULT":           32,
	"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
	"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
	"SUB_IF_API_FLAG_MASK_VNET":         254,
	"SUB_IF_API_FLAG_DOT1AH":            256,
}

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// FlowprobeParams represents VPP binary API message 'flowprobe_params'.
type FlowprobeParams struct {
	RecordFlags  FlowprobeRecordFlags
	ActiveTimer  uint32
	PassiveTimer uint32
}

func (m *FlowprobeParams) Reset()                        { *m = FlowprobeParams{} }
func (*FlowprobeParams) GetMessageName() string          { return "flowprobe_params" }
func (*FlowprobeParams) GetCrcString() string            { return "baa46c09" }
func (*FlowprobeParams) GetMessageType() api.MessageType { return api.RequestMessage }

// FlowprobeParamsReply represents VPP binary API message 'flowprobe_params_reply'.
type FlowprobeParamsReply struct {
	Retval int32
}

func (m *FlowprobeParamsReply) Reset()                        { *m = FlowprobeParamsReply{} }
func (*FlowprobeParamsReply) GetMessageName() string          { return "flowprobe_params_reply" }
func (*FlowprobeParamsReply) GetCrcString() string            { return "e8d4e804" }
func (*FlowprobeParamsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// FlowprobeTxInterfaceAddDel represents VPP binary API message 'flowprobe_tx_interface_add_del'.
type FlowprobeTxInterfaceAddDel struct {
	IsAdd     bool
	Which     FlowprobeWhichFlags
	SwIfIndex InterfaceIndex
}

func (m *FlowprobeTxInterfaceAddDel) Reset()                        { *m = FlowprobeTxInterfaceAddDel{} }
func (*FlowprobeTxInterfaceAddDel) GetMessageName() string          { return "flowprobe_tx_interface_add_del" }
func (*FlowprobeTxInterfaceAddDel) GetCrcString() string            { return "b782c976" }
func (*FlowprobeTxInterfaceAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// FlowprobeTxInterfaceAddDelReply represents VPP binary API message 'flowprobe_tx_interface_add_del_reply'.
type FlowprobeTxInterfaceAddDelReply struct {
	Retval int32
}

func (m *FlowprobeTxInterfaceAddDelReply) Reset() { *m = FlowprobeTxInterfaceAddDelReply{} }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageName() string {
	return "flowprobe_tx_interface_add_del_reply"
}
func (*FlowprobeTxInterfaceAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*FlowprobeParams)(nil), "flowprobe.FlowprobeParams")
	api.RegisterMessage((*FlowprobeParamsReply)(nil), "flowprobe.FlowprobeParamsReply")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDel)(nil), "flowprobe.FlowprobeTxInterfaceAddDel")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDelReply)(nil), "flowprobe.FlowprobeTxInterfaceAddDelReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FlowprobeParams)(nil),
		(*FlowprobeParamsReply)(nil),
		(*FlowprobeTxInterfaceAddDel)(nil),
		(*FlowprobeTxInterfaceAddDelReply)(nil),
	}
}

// RPCService represents RPC service API for flowprobe module.
type RPCService interface {
	FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error)
	FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error) {
	out := new(FlowprobeParamsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error) {
	out := new(FlowprobeTxInterfaceAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
