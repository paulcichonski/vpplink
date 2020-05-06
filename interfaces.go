// Copyright (C) 2019 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpplink

import (
	"bytes"
	"fmt"
	"net"

	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/gso"
	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/interfaces"
	vppip "github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/ip"
	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/ip_neighbor"
	"github.com/calico-vpp/vpplink/binapi/20.05-rc0~540-gad1cca49e/tapv2"
	"github.com/calico-vpp/vpplink/types"
	"github.com/pkg/errors"
)

const (
	INVALID_SW_IF_INDEX = ^uint32(0)
)

type NamespaceNotFound error

func (v *VppLink) CreateTapV2(tap *types.TapV2) (swIfIndex uint32, err error) {
	v.lock.Lock()
	defer v.lock.Unlock()
	response := &tapv2.TapCreateV2Reply{}
	request := &tapv2.TapCreateV2{
		// TODO check namespace len < 64?
		// TODO set MTU?
		ID:          ^uint32(0),
		Tag:         tap.Tag,
		MacAddress:  tap.GetVppMacAddress(),
		TapFlags:    tapv2.TapFlags(tap.Flags),
		NumRxQueues: uint8(tap.RxQueues),
	}
	if tap.HostNamespace != "" {
		request.HostNamespaceSet = true
		request.HostNamespace = tap.HostNamespace
	}
	if tap.HostIfName != "" {
		request.HostIfName = tap.HostIfName
		request.HostIfNameSet = true
	}
	if tap.HostMacAddress != nil {
		request.HostMacAddr = tap.GetVppHostMacAddress()
		request.HostMacAddrSet = true
	}
	err = v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return INVALID_SW_IF_INDEX, errors.Wrap(err, "Tap creation request failed")
	} else if response.Retval == -12 {
		return INVALID_SW_IF_INDEX, nil
	} else if response.Retval != 0 {
		return INVALID_SW_IF_INDEX, fmt.Errorf("Tap creation failed (retval %d). Request: %+v", response.Retval, request)
	}
	return uint32(response.SwIfIndex), err
}

func (v *VppLink) CreateOrAttachTapV2(tap *types.TapV2) (swIfIndex uint32, err error) {
	tap.Flags |= types.TapFlagPersist | types.TapFlagAttach
	swIfIndex, err = v.CreateTapV2(tap)
	if err == nil && swIfIndex == INVALID_SW_IF_INDEX {
		tap.Flags &= ^types.TapFlagAttach
		return v.CreateTapV2(tap)
	}
	return swIfIndex, err
}

func (v *VppLink) addDelInterfaceAddress(swIfIndex uint32, addr *net.IPNet, isAdd bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	p := types.ToVppIpPrefix(addr)
	request := &interfaces.SwInterfaceAddDelAddress{
		SwIfIndex: interfaces.InterfaceIndex(swIfIndex),
		IsAdd:     isAdd,
		Prefix: interfaces.AddressWithPrefix{
			Address: interfaces.Address{
				Un: interfaces.AddressUnion(p.Address.Un),
				Af: interfaces.AddressFamily(p.Address.Af),
			},
			Len: p.Len,
		},
	}
	response := &interfaces.SwInterfaceAddDelAddressReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "Adding IP address failed: req %+v reply %+v", request, response)
	}
	return nil
}

func (v *VppLink) DelInterfaceAddress(swIfIndex uint32, addr *net.IPNet) error {
	return v.addDelInterfaceAddress(swIfIndex, addr, false)
}

func (v *VppLink) AddInterfaceAddress(swIfIndex uint32, addr *net.IPNet) error {
	return v.addDelInterfaceAddress(swIfIndex, addr, true)
}

func (v *VppLink) enableDisableInterfaceIP6(swIfIndex uint32, enable bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.SwInterfaceIP6EnableDisable{
		SwIfIndex: vppip.InterfaceIndex(swIfIndex),
		Enable:    enable,
	}
	response := &vppip.SwInterfaceIP6EnableDisableReply{}
	return v.ch.SendRequest(request).ReceiveReply(response)
}

func (v *VppLink) DisableInterfaceIP6(swIfIndex uint32) error {
	return v.enableDisableInterfaceIP6(swIfIndex, false)
}

func (v *VppLink) EnableInterfaceIP6(swIfIndex uint32) error {
	return v.enableDisableInterfaceIP6(swIfIndex, true)
}

func (v *VppLink) SearchInterfaceWithTag(tag string) (err error, swIfIndex uint32) {
	v.lock.Lock()
	defer v.lock.Unlock()

	swIfIndex = INVALID_SW_IF_INDEX
	request := &interfaces.SwInterfaceDump{}
	stream := v.ch.SendMultiRequest(request)
	for {
		response := &interfaces.SwInterfaceDetails{}
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP interfaces: %v", err)
			return err, INVALID_SW_IF_INDEX
		}
		if stop {
			break
		}
		intfTag := string(bytes.Trim([]byte(response.Tag), "\x00"))
		v.log.Debugf("found interface %d, tag: %s (len %d)", response.SwIfIndex, intfTag, len(intfTag))
		if intfTag == tag {
			swIfIndex = uint32(response.SwIfIndex)
		}
	}
	if swIfIndex == INVALID_SW_IF_INDEX {
		v.log.Errorf("Interface with tag %s not found", tag)
		return errors.New("Interface not found"), INVALID_SW_IF_INDEX
	}
	return nil, swIfIndex
}

func (v *VppLink) SearchInterfaceWithName(name string) (err error, swIfIndex uint32) {
	v.lock.Lock()
	defer v.lock.Unlock()

	swIfIndex = INVALID_SW_IF_INDEX
	request := &interfaces.SwInterfaceDump{
		SwIfIndex: interfaces.InterfaceIndex(INVALID_SW_IF_INDEX),
		// TODO: filter by name with NameFilter
	}
	reqCtx := v.ch.SendMultiRequest(request)
	for {
		response := &interfaces.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("SwInterfaceDump failed: %v", err)
			return err, INVALID_SW_IF_INDEX
		}
		if stop {
			break
		}
		interfaceName := string(bytes.Trim([]byte(response.InterfaceName), "\x00"))
		v.log.Debugf("Found interface: %s", interfaceName)
		if interfaceName == name {
			swIfIndex = uint32(response.SwIfIndex)
		}

	}
	if swIfIndex == INVALID_SW_IF_INDEX {
		v.log.Errorf("Interface %s not found", name)
		return errors.New("Interface not found"), INVALID_SW_IF_INDEX
	}
	return nil, swIfIndex
}

func (v *VppLink) GetInterfaceDetails(swIfIndex uint32) (i *types.VppInterfaceDetails, err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &interfaces.SwInterfaceDump{
		SwIfIndex: interfaces.InterfaceIndex(swIfIndex),
	}
	stream := v.ch.SendMultiRequest(request)
	for {
		response := &interfaces.SwInterfaceDetails{}
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP interfaces: %v", err)
			return nil, err
		}
		if stop {
			break
		}
		if uint32(response.SwIfIndex) != swIfIndex {
			v.log.Debugf("Got interface that doesn't match filter, fix vpp")
			continue
		}
		v.log.Debugf("found interface %d", response.SwIfIndex)
		i = &types.VppInterfaceDetails{
			SwIfIndex: uint32(response.SwIfIndex),
			IsUp:      response.Flags&interfaces.IF_STATUS_API_FLAG_ADMIN_UP > 0,
			Name:      response.InterfaceName,
			Tag:       response.Tag,
			Type:      response.InterfaceDevType,
		}
	}
	if i == nil {
		return nil, errors.New("Interface not found")
	}
	return i, nil
}

func (v *VppLink) interfaceAdminUpDown(swIfIndex uint32, up bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	var f interfaces.IfStatusFlags = 0
	if up {
		f |= interfaces.IF_STATUS_API_FLAG_ADMIN_UP
	}
	// Set interface down
	request := &interfaces.SwInterfaceSetFlags{
		SwIfIndex: interfaces.InterfaceIndex(swIfIndex),
		Flags:     f,
	}
	response := &interfaces.SwInterfaceSetFlagsReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "setting interface down failed")
	}
	return nil
}

func (v *VppLink) InterfaceAdminDown(swIfIndex uint32) error {
	return v.interfaceAdminUpDown(swIfIndex, false)
}

func (v *VppLink) InterfaceAdminUp(swIfIndex uint32) error {
	return v.interfaceAdminUpDown(swIfIndex, true)
}

func (v *VppLink) GetInterfaceNeighbors(swIfIndex uint32, isIPv6 bool) (err error, neighbors []types.Neighbor) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &ip_neighbor.IPNeighborDump{
		SwIfIndex: ip_neighbor.InterfaceIndex(swIfIndex),
		Af:        ip_neighbor.ADDRESS_IP4,
	}
	if isIPv6 {
		request.Af = ip_neighbor.ADDRESS_IP6
	}
	response := &ip_neighbor.IPNeighborDetails{}
	stream := v.ch.SendMultiRequest(request)
	for {
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			v.log.Errorf("error listing VPP neighbors: %v", err)
			return err, nil
		}
		if stop {
			return nil, neighbors
		}
		vppNeighbor := response.Neighbor
		neighbors = append(neighbors, types.Neighbor{
			SwIfIndex:    uint32(vppNeighbor.SwIfIndex),
			Flags:        types.FromVppNeighborFlags(vppNeighbor.Flags),
			IP:           types.FromVppNeighborAddress(vppNeighbor.IPAddress),
			HardwareAddr: types.FromVppNeighborMacAddress(vppNeighbor.MacAddress),
		})
	}
}

func (v *VppLink) DelTap(swIfIndex uint32) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &tapv2.TapDeleteV2{
		SwIfIndex: tapv2.InterfaceIndex(swIfIndex),
	}
	response := &tapv2.TapDeleteV2Reply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "failed to delete tap from VPP")
	}
	return nil
}

func (v *VppLink) interfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32, isAdd bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &interfaces.SwInterfaceSetUnnumbered{
		SwIfIndex:           interfaces.InterfaceIndex(swIfIndex),
		UnnumberedSwIfIndex: interfaces.InterfaceIndex(unnumberedSwIfIndex),
		IsAdd:               isAdd,
	}
	response := &interfaces.SwInterfaceSetUnnumberedReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "setting interface unnumbered failed %d -> %d", unnumberedSwIfIndex, swIfIndex)
	}
	return nil
}

func (v *VppLink) AddrList(swIfIndex uint32, isv6 bool) (addresses []types.IfAddress, err error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	request := &vppip.IPAddressDump{
		SwIfIndex: vppip.InterfaceIndex(swIfIndex),
		IsIPv6:    isv6,
	}
	stream := v.ch.SendMultiRequest(request)
	for {
		response := &vppip.IPAddressDetails{}
		stop, err := stream.ReceiveReply(response)
		if err != nil {
			return addresses, errors.Wrapf(err, "error listing VPP interfaces addresses")
		}
		if stop {
			break
		}
		address := types.IfAddress{
			SwIfIndex: uint32(response.SwIfIndex),
			IPNet:     *types.FromVppIpPrefix(vppip.Prefix(response.Prefix)),
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (v *VppLink) InterfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error {
	return v.interfaceSetUnnumbered(unnumberedSwIfIndex, swIfIndex, true)
}

func (v *VppLink) InterfaceUnsetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error {
	return v.interfaceSetUnnumbered(unnumberedSwIfIndex, swIfIndex, false)
}

func (v *VppLink) PuntRedirect(sourceSwIfIndex, destSwIfIndex uint32, nh net.IP) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	request := &vppip.IPPuntRedirect{
		Punt: vppip.PuntRedirect{
			RxSwIfIndex: vppip.InterfaceIndex(sourceSwIfIndex),
			TxSwIfIndex: vppip.InterfaceIndex(destSwIfIndex),
			Nh:          types.ToVppIpAddress(nh),
		},
		IsAdd: true,
	}
	response := &vppip.IPPuntRedirectReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil || response.Retval != 0 {
		return fmt.Errorf("cannot set punt in VPP: %v %d", err, response.Retval)
	}
	return nil
}

func (v *VppLink) enableDisableGso(swIfIndex uint32, enable bool) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	request := &gso.FeatureGsoEnableDisable{
		SwIfIndex:     gso.InterfaceIndex(swIfIndex),
		EnableDisable: enable,
	}
	response := &gso.FeatureGsoEnableDisableReply{}
	err := v.ch.SendRequest(request).ReceiveReply(response)
	if err != nil || response.Retval != 0 {
		return fmt.Errorf("cannot configure gso: %v %d", err, response.Retval)
	}
	return nil
}

func (v *VppLink) EnableGSOFeature(swIfIndex uint32) error {
	return v.enableDisableGso(swIfIndex, true)
}

func (v *VppLink) DisableGSOFeature(swIfIndex uint32) error {
	return v.enableDisableGso(swIfIndex, false)
}
