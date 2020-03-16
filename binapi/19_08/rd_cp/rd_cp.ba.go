// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/rd_cp.api.json

/*
Package rd_cp is a generated VPP binary API for 'rd_cp' module.

It consists of:
	  2 messages
	  1 service
*/
package rd_cp

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
	ModuleName = "rd_cp"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xc91af796
)

// IP6NdAddressAutoconfig represents VPP binary API message 'ip6_nd_address_autoconfig'.
type IP6NdAddressAutoconfig struct {
	SwIfIndex            uint32
	Enable               uint8
	InstallDefaultRoutes uint8
}

func (*IP6NdAddressAutoconfig) GetMessageName() string {
	return "ip6_nd_address_autoconfig"
}
func (*IP6NdAddressAutoconfig) GetCrcString() string {
	return "7598dda9"
}
func (*IP6NdAddressAutoconfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IP6NdAddressAutoconfigReply represents VPP binary API message 'ip6_nd_address_autoconfig_reply'.
type IP6NdAddressAutoconfigReply struct {
	Retval int32
}

func (*IP6NdAddressAutoconfigReply) GetMessageName() string {
	return "ip6_nd_address_autoconfig_reply"
}
func (*IP6NdAddressAutoconfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IP6NdAddressAutoconfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*IP6NdAddressAutoconfig)(nil), "rd_cp.IP6NdAddressAutoconfig")
	api.RegisterMessage((*IP6NdAddressAutoconfigReply)(nil), "rd_cp.IP6NdAddressAutoconfigReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IP6NdAddressAutoconfig)(nil),
		(*IP6NdAddressAutoconfigReply)(nil),
	}
}

// RPCService represents RPC service API for rd_cp module.
type RPCService interface {
	IP6NdAddressAutoconfig(ctx context.Context, in *IP6NdAddressAutoconfig) (*IP6NdAddressAutoconfigReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) IP6NdAddressAutoconfig(ctx context.Context, in *IP6NdAddressAutoconfig) (*IP6NdAddressAutoconfigReply, error) {
	out := new(IP6NdAddressAutoconfigReply)
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
