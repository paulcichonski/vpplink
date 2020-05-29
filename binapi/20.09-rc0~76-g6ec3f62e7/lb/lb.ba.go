// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /Users/nskrzypc/wrk/vpp/build-root/install-vpp-native/vpp/share/vpp/api/plugins/lb.api.json

/*
Package lb is a generated VPP binary API for 'lb' module.

It consists of:
	 15 enums
	  6 aliases
	  7 types
	  1 union
	 16 messages
	  8 services
*/
package lb

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
	ModuleName = "lb"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x5e4e85b1
)

// AddressFamily represents VPP binary API enum 'address_family'.
type AddressFamily uint8

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var AddressFamily_name = map[uint8]string{
	0: "ADDRESS_IP4",
	1: "ADDRESS_IP6",
}

var AddressFamily_value = map[string]uint8{
	"ADDRESS_IP4": 0,
	"ADDRESS_IP6": 1,
}

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint8(x)]
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

// IPDscp represents VPP binary API enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var IPDscp_name = map[uint8]string{
	0:  "IP_API_DSCP_CS0",
	8:  "IP_API_DSCP_CS1",
	10: "IP_API_DSCP_AF11",
	12: "IP_API_DSCP_AF12",
	14: "IP_API_DSCP_AF13",
	16: "IP_API_DSCP_CS2",
	18: "IP_API_DSCP_AF21",
	20: "IP_API_DSCP_AF22",
	22: "IP_API_DSCP_AF23",
	24: "IP_API_DSCP_CS3",
	26: "IP_API_DSCP_AF31",
	28: "IP_API_DSCP_AF32",
	30: "IP_API_DSCP_AF33",
	32: "IP_API_DSCP_CS4",
	34: "IP_API_DSCP_AF41",
	36: "IP_API_DSCP_AF42",
	38: "IP_API_DSCP_AF43",
	40: "IP_API_DSCP_CS5",
	46: "IP_API_DSCP_EF",
	48: "IP_API_DSCP_CS6",
	50: "IP_API_DSCP_CS7",
}

var IPDscp_value = map[string]uint8{
	"IP_API_DSCP_CS0":  0,
	"IP_API_DSCP_CS1":  8,
	"IP_API_DSCP_AF11": 10,
	"IP_API_DSCP_AF12": 12,
	"IP_API_DSCP_AF13": 14,
	"IP_API_DSCP_CS2":  16,
	"IP_API_DSCP_AF21": 18,
	"IP_API_DSCP_AF22": 20,
	"IP_API_DSCP_AF23": 22,
	"IP_API_DSCP_CS3":  24,
	"IP_API_DSCP_AF31": 26,
	"IP_API_DSCP_AF32": 28,
	"IP_API_DSCP_AF33": 30,
	"IP_API_DSCP_CS4":  32,
	"IP_API_DSCP_AF41": 34,
	"IP_API_DSCP_AF42": 36,
	"IP_API_DSCP_AF43": 38,
	"IP_API_DSCP_CS5":  40,
	"IP_API_DSCP_EF":   46,
	"IP_API_DSCP_CS6":  48,
	"IP_API_DSCP_CS7":  50,
}

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPEcn represents VPP binary API enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var IPEcn_name = map[uint8]string{
	0: "IP_API_ECN_NONE",
	1: "IP_API_ECN_ECT0",
	2: "IP_API_ECN_ECT1",
	3: "IP_API_ECN_CE",
}

var IPEcn_value = map[string]uint8{
	"IP_API_ECN_NONE": 0,
	"IP_API_ECN_ECT0": 1,
	"IP_API_ECN_ECT1": 2,
	"IP_API_ECN_CE":   3,
}

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPProto represents VPP binary API enum 'ip_proto'.
type IPProto uint8

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_ESP      IPProto = 50
	IP_API_PROTO_AH       IPProto = 51
	IP_API_PROTO_ICMP6    IPProto = 58
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var IPProto_name = map[uint8]string{
	0:   "IP_API_PROTO_HOPOPT",
	1:   "IP_API_PROTO_ICMP",
	2:   "IP_API_PROTO_IGMP",
	6:   "IP_API_PROTO_TCP",
	17:  "IP_API_PROTO_UDP",
	47:  "IP_API_PROTO_GRE",
	50:  "IP_API_PROTO_ESP",
	51:  "IP_API_PROTO_AH",
	58:  "IP_API_PROTO_ICMP6",
	88:  "IP_API_PROTO_EIGRP",
	89:  "IP_API_PROTO_OSPF",
	132: "IP_API_PROTO_SCTP",
	255: "IP_API_PROTO_RESERVED",
}

var IPProto_value = map[string]uint8{
	"IP_API_PROTO_HOPOPT":   0,
	"IP_API_PROTO_ICMP":     1,
	"IP_API_PROTO_IGMP":     2,
	"IP_API_PROTO_TCP":      6,
	"IP_API_PROTO_UDP":      17,
	"IP_API_PROTO_GRE":      47,
	"IP_API_PROTO_ESP":      50,
	"IP_API_PROTO_AH":       51,
	"IP_API_PROTO_ICMP6":    58,
	"IP_API_PROTO_EIGRP":    88,
	"IP_API_PROTO_OSPF":     89,
	"IP_API_PROTO_SCTP":     132,
	"IP_API_PROTO_RESERVED": 255,
}

func (x IPProto) String() string {
	s, ok := IPProto_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbEncapType represents VPP binary API enum 'lb_encap_type'.
type LbEncapType uint32

const (
	LB_API_ENCAP_TYPE_GRE4  LbEncapType = 0
	LB_API_ENCAP_TYPE_GRE6  LbEncapType = 1
	LB_API_ENCAP_TYPE_L3DSR LbEncapType = 2
	LB_API_ENCAP_TYPE_NAT4  LbEncapType = 3
	LB_API_ENCAP_TYPE_NAT6  LbEncapType = 4
	LB_API_ENCAP_N_TYPES    LbEncapType = 5
)

var LbEncapType_name = map[uint32]string{
	0: "LB_API_ENCAP_TYPE_GRE4",
	1: "LB_API_ENCAP_TYPE_GRE6",
	2: "LB_API_ENCAP_TYPE_L3DSR",
	3: "LB_API_ENCAP_TYPE_NAT4",
	4: "LB_API_ENCAP_TYPE_NAT6",
	5: "LB_API_ENCAP_N_TYPES",
}

var LbEncapType_value = map[string]uint32{
	"LB_API_ENCAP_TYPE_GRE4":  0,
	"LB_API_ENCAP_TYPE_GRE6":  1,
	"LB_API_ENCAP_TYPE_L3DSR": 2,
	"LB_API_ENCAP_TYPE_NAT4":  3,
	"LB_API_ENCAP_TYPE_NAT6":  4,
	"LB_API_ENCAP_N_TYPES":    5,
}

func (x LbEncapType) String() string {
	s, ok := LbEncapType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbLkpTypeT represents VPP binary API enum 'lb_lkp_type_t'.
type LbLkpTypeT uint32

const (
	LB_API_LKP_SAME_IP_PORT LbLkpTypeT = 0
	LB_API_LKP_DIFF_IP_PORT LbLkpTypeT = 1
	LB_API_LKP_ALL_PORT_IP  LbLkpTypeT = 2
	LB_API_LKP_N_TYPES      LbLkpTypeT = 3
)

var LbLkpTypeT_name = map[uint32]string{
	0: "LB_API_LKP_SAME_IP_PORT",
	1: "LB_API_LKP_DIFF_IP_PORT",
	2: "LB_API_LKP_ALL_PORT_IP",
	3: "LB_API_LKP_N_TYPES",
}

var LbLkpTypeT_value = map[string]uint32{
	"LB_API_LKP_SAME_IP_PORT": 0,
	"LB_API_LKP_DIFF_IP_PORT": 1,
	"LB_API_LKP_ALL_PORT_IP":  2,
	"LB_API_LKP_N_TYPES":      3,
}

func (x LbLkpTypeT) String() string {
	s, ok := LbLkpTypeT_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbNatProtocol represents VPP binary API enum 'lb_nat_protocol'.
type LbNatProtocol uint32

const (
	LB_API_NAT_PROTOCOL_UDP LbNatProtocol = 6
	LB_API_NAT_PROTOCOL_TCP LbNatProtocol = 23
	LB_API_NAT_PROTOCOL_ANY LbNatProtocol = 4.294967295e+09
)

var LbNatProtocol_name = map[uint32]string{
	6:               "LB_API_NAT_PROTOCOL_UDP",
	23:              "LB_API_NAT_PROTOCOL_TCP",
	4.294967295e+09: "LB_API_NAT_PROTOCOL_ANY",
}

var LbNatProtocol_value = map[string]uint32{
	"LB_API_NAT_PROTOCOL_UDP": 6,
	"LB_API_NAT_PROTOCOL_TCP": 23,
	"LB_API_NAT_PROTOCOL_ANY": 4.294967295e+09,
}

func (x LbNatProtocol) String() string {
	s, ok := LbNatProtocol_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbSrvType represents VPP binary API enum 'lb_srv_type'.
type LbSrvType uint32

const (
	LB_API_SRV_TYPE_CLUSTERIP LbSrvType = 0
	LB_API_SRV_TYPE_NODEPORT  LbSrvType = 1
	LB_API_SRV_N_TYPES        LbSrvType = 2
)

var LbSrvType_name = map[uint32]string{
	0: "LB_API_SRV_TYPE_CLUSTERIP",
	1: "LB_API_SRV_TYPE_NODEPORT",
	2: "LB_API_SRV_N_TYPES",
}

var LbSrvType_value = map[string]uint32{
	"LB_API_SRV_TYPE_CLUSTERIP": 0,
	"LB_API_SRV_TYPE_NODEPORT":  1,
	"LB_API_SRV_N_TYPES":        2,
}

func (x LbSrvType) String() string {
	s, ok := LbSrvType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbVipType represents VPP binary API enum 'lb_vip_type'.
type LbVipType uint32

const (
	LB_API_VIP_TYPE_IP6_GRE6  LbVipType = 0
	LB_API_VIP_TYPE_IP6_GRE4  LbVipType = 1
	LB_API_VIP_TYPE_IP4_GRE6  LbVipType = 2
	LB_API_VIP_TYPE_IP4_GRE4  LbVipType = 3
	LB_API_VIP_TYPE_IP4_L3DSR LbVipType = 4
	LB_API_VIP_TYPE_IP4_NAT4  LbVipType = 5
	LB_API_VIP_TYPE_IP6_NAT6  LbVipType = 6
	LB_API_VIP_N_TYPES        LbVipType = 7
)

var LbVipType_name = map[uint32]string{
	0: "LB_API_VIP_TYPE_IP6_GRE6",
	1: "LB_API_VIP_TYPE_IP6_GRE4",
	2: "LB_API_VIP_TYPE_IP4_GRE6",
	3: "LB_API_VIP_TYPE_IP4_GRE4",
	4: "LB_API_VIP_TYPE_IP4_L3DSR",
	5: "LB_API_VIP_TYPE_IP4_NAT4",
	6: "LB_API_VIP_TYPE_IP6_NAT6",
	7: "LB_API_VIP_N_TYPES",
}

var LbVipType_value = map[string]uint32{
	"LB_API_VIP_TYPE_IP6_GRE6":  0,
	"LB_API_VIP_TYPE_IP6_GRE4":  1,
	"LB_API_VIP_TYPE_IP4_GRE6":  2,
	"LB_API_VIP_TYPE_IP4_GRE4":  3,
	"LB_API_VIP_TYPE_IP4_L3DSR": 4,
	"LB_API_VIP_TYPE_IP4_NAT4":  5,
	"LB_API_VIP_TYPE_IP6_NAT6":  6,
	"LB_API_VIP_N_TYPES":        7,
}

func (x LbVipType) String() string {
	s, ok := LbVipType_name[uint32(x)]
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

// AddressWithPrefix represents VPP binary API alias 'address_with_prefix'.
type AddressWithPrefix Prefix

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// IP4Address represents VPP binary API alias 'ip4_address'.
type IP4Address [4]uint8

// IP4AddressWithPrefix represents VPP binary API alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address represents VPP binary API alias 'ip6_address'.
type IP6Address [16]uint8

// IP6AddressWithPrefix represents VPP binary API alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// Address represents VPP binary API type 'address'.
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string { return "address" }

// IP4Prefix represents VPP binary API type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address
	Len     uint8
}

func (*IP4Prefix) GetTypeName() string { return "ip4_prefix" }

// IP6Prefix represents VPP binary API type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address
	Len     uint8
}

func (*IP6Prefix) GetTypeName() string { return "ip6_prefix" }

// LbVip represents VPP binary API type 'lb_vip'.
type LbVip struct {
	Pfx      AddressWithPrefix
	Protocol IPProto
	Port     uint16
}

func (*LbVip) GetTypeName() string { return "lb_vip" }

// Mprefix represents VPP binary API type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string { return "mprefix" }

// Prefix represents VPP binary API type 'prefix'.
type Prefix struct {
	Address Address
	Len     uint8
}

func (*Prefix) GetTypeName() string { return "prefix" }

// PrefixMatcher represents VPP binary API type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8
	Ge uint8
}

func (*PrefixMatcher) GetTypeName() string { return "prefix_matcher" }

// AddressUnion represents VPP binary API union 'address_union'.
type AddressUnion struct {
	XXX_UnionData [16]byte
}

func (*AddressUnion) GetTypeName() string { return "address_union" }

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

// LbAddDelAs represents VPP binary API message 'lb_add_del_as'.
type LbAddDelAs struct {
	Pfx       AddressWithPrefix
	Protocol  uint8
	Port      uint16
	AsAddress Address
	IsDel     bool
	IsFlush   bool
}

func (m *LbAddDelAs) Reset()                        { *m = LbAddDelAs{} }
func (*LbAddDelAs) GetMessageName() string          { return "lb_add_del_as" }
func (*LbAddDelAs) GetCrcString() string            { return "78628987" }
func (*LbAddDelAs) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelAsReply represents VPP binary API message 'lb_add_del_as_reply'.
type LbAddDelAsReply struct {
	Retval int32
}

func (m *LbAddDelAsReply) Reset()                        { *m = LbAddDelAsReply{} }
func (*LbAddDelAsReply) GetMessageName() string          { return "lb_add_del_as_reply" }
func (*LbAddDelAsReply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelAsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAddDelIntfNat4 represents VPP binary API message 'lb_add_del_intf_nat4'.
type LbAddDelIntfNat4 struct {
	IsAdd     bool
	SwIfIndex InterfaceIndex
}

func (m *LbAddDelIntfNat4) Reset()                        { *m = LbAddDelIntfNat4{} }
func (*LbAddDelIntfNat4) GetMessageName() string          { return "lb_add_del_intf_nat4" }
func (*LbAddDelIntfNat4) GetCrcString() string            { return "47d6e753" }
func (*LbAddDelIntfNat4) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelIntfNat4Reply represents VPP binary API message 'lb_add_del_intf_nat4_reply'.
type LbAddDelIntfNat4Reply struct {
	Retval int32
}

func (m *LbAddDelIntfNat4Reply) Reset()                        { *m = LbAddDelIntfNat4Reply{} }
func (*LbAddDelIntfNat4Reply) GetMessageName() string          { return "lb_add_del_intf_nat4_reply" }
func (*LbAddDelIntfNat4Reply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelIntfNat4Reply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAddDelIntfNat6 represents VPP binary API message 'lb_add_del_intf_nat6'.
type LbAddDelIntfNat6 struct {
	IsAdd     bool
	SwIfIndex InterfaceIndex
}

func (m *LbAddDelIntfNat6) Reset()                        { *m = LbAddDelIntfNat6{} }
func (*LbAddDelIntfNat6) GetMessageName() string          { return "lb_add_del_intf_nat6" }
func (*LbAddDelIntfNat6) GetCrcString() string            { return "47d6e753" }
func (*LbAddDelIntfNat6) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelIntfNat6Reply represents VPP binary API message 'lb_add_del_intf_nat6_reply'.
type LbAddDelIntfNat6Reply struct {
	Retval int32
}

func (m *LbAddDelIntfNat6Reply) Reset()                        { *m = LbAddDelIntfNat6Reply{} }
func (*LbAddDelIntfNat6Reply) GetMessageName() string          { return "lb_add_del_intf_nat6_reply" }
func (*LbAddDelIntfNat6Reply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelIntfNat6Reply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAddDelVip represents VPP binary API message 'lb_add_del_vip'.
type LbAddDelVip struct {
	Pfx                 AddressWithPrefix
	Protocol            uint8
	Port                uint16
	Encap               LbEncapType
	Dscp                uint8
	Type                LbSrvType
	TargetPort          uint16
	NodePort            uint16
	NewFlowsTableLength uint32
	IsDel               bool
}

func (m *LbAddDelVip) Reset()                        { *m = LbAddDelVip{} }
func (*LbAddDelVip) GetMessageName() string          { return "lb_add_del_vip" }
func (*LbAddDelVip) GetCrcString() string            { return "d15b7ddc" }
func (*LbAddDelVip) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelVipReply represents VPP binary API message 'lb_add_del_vip_reply'.
type LbAddDelVipReply struct {
	Retval int32
}

func (m *LbAddDelVipReply) Reset()                        { *m = LbAddDelVipReply{} }
func (*LbAddDelVipReply) GetMessageName() string          { return "lb_add_del_vip_reply" }
func (*LbAddDelVipReply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelVipReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAsDetails represents VPP binary API message 'lb_as_details'.
type LbAsDetails struct {
	Vip        LbVip
	AppSrv     Address
	Flags      uint8
	InUseSince uint32
}

func (m *LbAsDetails) Reset()                        { *m = LbAsDetails{} }
func (*LbAsDetails) GetMessageName() string          { return "lb_as_details" }
func (*LbAsDetails) GetCrcString() string            { return "9c39f60e" }
func (*LbAsDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAsDump represents VPP binary API message 'lb_as_dump'.
type LbAsDump struct {
	Pfx      AddressWithPrefix
	Protocol uint8
	Port     uint16
}

func (m *LbAsDump) Reset()                        { *m = LbAsDump{} }
func (*LbAsDump) GetMessageName() string          { return "lb_as_dump" }
func (*LbAsDump) GetCrcString() string            { return "1063f819" }
func (*LbAsDump) GetMessageType() api.MessageType { return api.RequestMessage }

// LbConf represents VPP binary API message 'lb_conf'.
type LbConf struct {
	IP4SrcAddress        IP4Address
	IP6SrcAddress        IP6Address
	StickyBucketsPerCore uint32
	FlowTimeout          uint32
}

func (m *LbConf) Reset()                        { *m = LbConf{} }
func (*LbConf) GetMessageName() string          { return "lb_conf" }
func (*LbConf) GetCrcString() string            { return "22ddb739" }
func (*LbConf) GetMessageType() api.MessageType { return api.RequestMessage }

// LbConfReply represents VPP binary API message 'lb_conf_reply'.
type LbConfReply struct {
	Retval int32
}

func (m *LbConfReply) Reset()                        { *m = LbConfReply{} }
func (*LbConfReply) GetMessageName() string          { return "lb_conf_reply" }
func (*LbConfReply) GetCrcString() string            { return "e8d4e804" }
func (*LbConfReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbFlushVip represents VPP binary API message 'lb_flush_vip'.
type LbFlushVip struct {
	Pfx      AddressWithPrefix
	Protocol uint8
	Port     uint16
}

func (m *LbFlushVip) Reset()                        { *m = LbFlushVip{} }
func (*LbFlushVip) GetMessageName() string          { return "lb_flush_vip" }
func (*LbFlushVip) GetCrcString() string            { return "1063f819" }
func (*LbFlushVip) GetMessageType() api.MessageType { return api.RequestMessage }

// LbFlushVipReply represents VPP binary API message 'lb_flush_vip_reply'.
type LbFlushVipReply struct {
	Retval int32
}

func (m *LbFlushVipReply) Reset()                        { *m = LbFlushVipReply{} }
func (*LbFlushVipReply) GetMessageName() string          { return "lb_flush_vip_reply" }
func (*LbFlushVipReply) GetCrcString() string            { return "e8d4e804" }
func (*LbFlushVipReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbVipDetails represents VPP binary API message 'lb_vip_details'.
type LbVipDetails struct {
	Vip             LbVip
	Encap           LbEncapType
	Dscp            IPDscp
	SrvType         LbSrvType
	TargetPort      uint16
	FlowTableLength uint16
}

func (m *LbVipDetails) Reset()                        { *m = LbVipDetails{} }
func (*LbVipDetails) GetMessageName() string          { return "lb_vip_details" }
func (*LbVipDetails) GetCrcString() string            { return "08f39bed" }
func (*LbVipDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbVipDump represents VPP binary API message 'lb_vip_dump'.
type LbVipDump struct {
	Pfx        AddressWithPrefix
	PfxMatcher PrefixMatcher
	Protocol   uint8
	Port       uint16
}

func (m *LbVipDump) Reset()                        { *m = LbVipDump{} }
func (*LbVipDump) GetMessageName() string          { return "lb_vip_dump" }
func (*LbVipDump) GetCrcString() string            { return "c7bcb124" }
func (*LbVipDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*LbAddDelAs)(nil), "lb.LbAddDelAs")
	api.RegisterMessage((*LbAddDelAsReply)(nil), "lb.LbAddDelAsReply")
	api.RegisterMessage((*LbAddDelIntfNat4)(nil), "lb.LbAddDelIntfNat4")
	api.RegisterMessage((*LbAddDelIntfNat4Reply)(nil), "lb.LbAddDelIntfNat4Reply")
	api.RegisterMessage((*LbAddDelIntfNat6)(nil), "lb.LbAddDelIntfNat6")
	api.RegisterMessage((*LbAddDelIntfNat6Reply)(nil), "lb.LbAddDelIntfNat6Reply")
	api.RegisterMessage((*LbAddDelVip)(nil), "lb.LbAddDelVip")
	api.RegisterMessage((*LbAddDelVipReply)(nil), "lb.LbAddDelVipReply")
	api.RegisterMessage((*LbAsDetails)(nil), "lb.LbAsDetails")
	api.RegisterMessage((*LbAsDump)(nil), "lb.LbAsDump")
	api.RegisterMessage((*LbConf)(nil), "lb.LbConf")
	api.RegisterMessage((*LbConfReply)(nil), "lb.LbConfReply")
	api.RegisterMessage((*LbFlushVip)(nil), "lb.LbFlushVip")
	api.RegisterMessage((*LbFlushVipReply)(nil), "lb.LbFlushVipReply")
	api.RegisterMessage((*LbVipDetails)(nil), "lb.LbVipDetails")
	api.RegisterMessage((*LbVipDump)(nil), "lb.LbVipDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LbAddDelAs)(nil),
		(*LbAddDelAsReply)(nil),
		(*LbAddDelIntfNat4)(nil),
		(*LbAddDelIntfNat4Reply)(nil),
		(*LbAddDelIntfNat6)(nil),
		(*LbAddDelIntfNat6Reply)(nil),
		(*LbAddDelVip)(nil),
		(*LbAddDelVipReply)(nil),
		(*LbAsDetails)(nil),
		(*LbAsDump)(nil),
		(*LbConf)(nil),
		(*LbConfReply)(nil),
		(*LbFlushVip)(nil),
		(*LbFlushVipReply)(nil),
		(*LbVipDetails)(nil),
		(*LbVipDump)(nil),
	}
}

// RPCService represents RPC service API for lb module.
type RPCService interface {
	DumpLbAs(ctx context.Context, in *LbAsDump) (RPCService_DumpLbAsClient, error)
	DumpLbVip(ctx context.Context, in *LbVipDump) (RPCService_DumpLbVipClient, error)
	LbAddDelAs(ctx context.Context, in *LbAddDelAs) (*LbAddDelAsReply, error)
	LbAddDelIntfNat4(ctx context.Context, in *LbAddDelIntfNat4) (*LbAddDelIntfNat4Reply, error)
	LbAddDelIntfNat6(ctx context.Context, in *LbAddDelIntfNat6) (*LbAddDelIntfNat6Reply, error)
	LbAddDelVip(ctx context.Context, in *LbAddDelVip) (*LbAddDelVipReply, error)
	LbConf(ctx context.Context, in *LbConf) (*LbConfReply, error)
	LbFlushVip(ctx context.Context, in *LbFlushVip) (*LbFlushVipReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpLbAs(ctx context.Context, in *LbAsDump) (RPCService_DumpLbAsClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpLbAsClient{stream}
	return x, nil
}

type RPCService_DumpLbAsClient interface {
	Recv() (*LbAsDetails, error)
}

type serviceClient_DumpLbAsClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpLbAsClient) Recv() (*LbAsDetails, error) {
	m := new(LbAsDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpLbVip(ctx context.Context, in *LbVipDump) (RPCService_DumpLbVipClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpLbVipClient{stream}
	return x, nil
}

type RPCService_DumpLbVipClient interface {
	Recv() (*LbVipDetails, error)
}

type serviceClient_DumpLbVipClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpLbVipClient) Recv() (*LbVipDetails, error) {
	m := new(LbVipDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) LbAddDelAs(ctx context.Context, in *LbAddDelAs) (*LbAddDelAsReply, error) {
	out := new(LbAddDelAsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbAddDelIntfNat4(ctx context.Context, in *LbAddDelIntfNat4) (*LbAddDelIntfNat4Reply, error) {
	out := new(LbAddDelIntfNat4Reply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbAddDelIntfNat6(ctx context.Context, in *LbAddDelIntfNat6) (*LbAddDelIntfNat6Reply, error) {
	out := new(LbAddDelIntfNat6Reply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbAddDelVip(ctx context.Context, in *LbAddDelVip) (*LbAddDelVipReply, error) {
	out := new(LbAddDelVipReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbConf(ctx context.Context, in *LbConf) (*LbConfReply, error) {
	out := new(LbConfReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbFlushVip(ctx context.Context, in *LbFlushVip) (*LbFlushVipReply, error) {
	out := new(LbFlushVipReply)
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