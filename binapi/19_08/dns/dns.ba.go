// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/dns.api.json

/*
Package dns is a generated VPP binary API for 'dns' module.

It consists of:
	  8 messages
	  4 services
*/
package dns

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "dns"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd464ac52
)

// DNSEnableDisable represents VPP binary API message 'dns_enable_disable'.
type DNSEnableDisable struct {
	Enable uint8
}

func (*DNSEnableDisable) GetMessageName() string {
	return "dns_enable_disable"
}
func (*DNSEnableDisable) GetCrcString() string {
	return "8050327d"
}
func (*DNSEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSEnableDisableReply represents VPP binary API message 'dns_enable_disable_reply'.
type DNSEnableDisableReply struct {
	Retval int32
}

func (*DNSEnableDisableReply) GetMessageName() string {
	return "dns_enable_disable_reply"
}
func (*DNSEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DNSEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DNSNameServerAddDel represents VPP binary API message 'dns_name_server_add_del'.
type DNSNameServerAddDel struct {
	IsIP6         uint8
	IsAdd         uint8
	ServerAddress []byte `struc:"[16]byte"`
}

func (*DNSNameServerAddDel) GetMessageName() string {
	return "dns_name_server_add_del"
}
func (*DNSNameServerAddDel) GetCrcString() string {
	return "3bb05d8c"
}
func (*DNSNameServerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSNameServerAddDelReply represents VPP binary API message 'dns_name_server_add_del_reply'.
type DNSNameServerAddDelReply struct {
	Retval int32
}

func (*DNSNameServerAddDelReply) GetMessageName() string {
	return "dns_name_server_add_del_reply"
}
func (*DNSNameServerAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DNSNameServerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DNSResolveIP represents VPP binary API message 'dns_resolve_ip'.
type DNSResolveIP struct {
	IsIP6   uint8
	Address []byte `struc:"[16]byte"`
}

func (*DNSResolveIP) GetMessageName() string {
	return "dns_resolve_ip"
}
func (*DNSResolveIP) GetCrcString() string {
	return "ae96a1a3"
}
func (*DNSResolveIP) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSResolveIPReply represents VPP binary API message 'dns_resolve_ip_reply'.
type DNSResolveIPReply struct {
	Retval int32
	Name   []byte `struc:"[256]byte"`
}

func (*DNSResolveIPReply) GetMessageName() string {
	return "dns_resolve_ip_reply"
}
func (*DNSResolveIPReply) GetCrcString() string {
	return "49ed78d6"
}
func (*DNSResolveIPReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DNSResolveName represents VPP binary API message 'dns_resolve_name'.
type DNSResolveName struct {
	Name []byte `struc:"[256]byte"`
}

func (*DNSResolveName) GetMessageName() string {
	return "dns_resolve_name"
}
func (*DNSResolveName) GetCrcString() string {
	return "c6566676"
}
func (*DNSResolveName) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSResolveNameReply represents VPP binary API message 'dns_resolve_name_reply'.
type DNSResolveNameReply struct {
	Retval     int32
	IP4Set     uint8
	IP6Set     uint8
	IP4Address []byte `struc:"[4]byte"`
	IP6Address []byte `struc:"[16]byte"`
}

func (*DNSResolveNameReply) GetMessageName() string {
	return "dns_resolve_name_reply"
}
func (*DNSResolveNameReply) GetCrcString() string {
	return "c2d758c3"
}
func (*DNSResolveNameReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*DNSEnableDisable)(nil), "dns.DNSEnableDisable")
	api.RegisterMessage((*DNSEnableDisableReply)(nil), "dns.DNSEnableDisableReply")
	api.RegisterMessage((*DNSNameServerAddDel)(nil), "dns.DNSNameServerAddDel")
	api.RegisterMessage((*DNSNameServerAddDelReply)(nil), "dns.DNSNameServerAddDelReply")
	api.RegisterMessage((*DNSResolveIP)(nil), "dns.DNSResolveIP")
	api.RegisterMessage((*DNSResolveIPReply)(nil), "dns.DNSResolveIPReply")
	api.RegisterMessage((*DNSResolveName)(nil), "dns.DNSResolveName")
	api.RegisterMessage((*DNSResolveNameReply)(nil), "dns.DNSResolveNameReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DNSEnableDisable)(nil),
		(*DNSEnableDisableReply)(nil),
		(*DNSNameServerAddDel)(nil),
		(*DNSNameServerAddDelReply)(nil),
		(*DNSResolveIP)(nil),
		(*DNSResolveIPReply)(nil),
		(*DNSResolveName)(nil),
		(*DNSResolveNameReply)(nil),
	}
}

// RPCService represents RPC service API for dns module.
type RPCService interface {
	DNSEnableDisable(ctx context.Context, in *DNSEnableDisable) (*DNSEnableDisableReply, error)
	DNSNameServerAddDel(ctx context.Context, in *DNSNameServerAddDel) (*DNSNameServerAddDelReply, error)
	DNSResolveIP(ctx context.Context, in *DNSResolveIP) (*DNSResolveIPReply, error)
	DNSResolveName(ctx context.Context, in *DNSResolveName) (*DNSResolveNameReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DNSEnableDisable(ctx context.Context, in *DNSEnableDisable) (*DNSEnableDisableReply, error) {
	out := new(DNSEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DNSNameServerAddDel(ctx context.Context, in *DNSNameServerAddDel) (*DNSNameServerAddDelReply, error) {
	out := new(DNSNameServerAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DNSResolveIP(ctx context.Context, in *DNSResolveIP) (*DNSResolveIPReply, error) {
	out := new(DNSResolveIPReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DNSResolveName(ctx context.Context, in *DNSResolveName) (*DNSResolveNameReply, error) {
	out := new(DNSResolveNameReply)
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
