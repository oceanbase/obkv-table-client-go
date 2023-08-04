/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObAddrVersion uint8

const (
	ObAddrIPV4 = 4
	ObAddrIPV6 = 6
)

const Ipv6Len = 16

type ObAddr struct {
	ObUniVersionHeader
	version ObAddrVersion
	ip      []uint8
	port    int32
}

func (a *ObAddr) IpToString() string {
	if a.IsIPv6() {
		return a.GetIPv6().String()
	}

	return a.GetIPv4().String()
}

func (a *ObAddr) IsIPv6() bool {
	return a.version == ObAddrIPV6
}

func (a *ObAddr) GetIPv4() net.IP {
	if a.IsIPv6() {
		return nil
	}
	return net.IPv4(a.ip[0], a.ip[1], a.ip[2], a.ip[3])
}

func (a *ObAddr) GetIPv6() net.IP {
	if !a.IsIPv6() {
		return nil
	}
	return net.IP(a.ip[:])
}

func (a *ObAddr) Ip() []uint8 {
	return a.ip
}

func (a *ObAddr) SetIp(ip []uint8) {
	a.ip = ip
}

func (a *ObAddr) Port() int32 {
	return a.port
}

func (a *ObAddr) SetPort(port int32) {
	a.port = port
}

func NewObAddr() *ObAddr {
	return &ObAddr{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		version: ObAddrIPV4,
		ip:      make([]uint8, Ipv6Len),
		port:    0,
	}
}

func (a *ObAddr) PayloadLen() int {
	return a.PayloadContentLen() + a.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (a *ObAddr) PayloadContentLen() int {
	totalLen := 0
	totalLen += 1 // version

	ip1 := binary.BigEndian.Uint32(a.ip[0:4])
	ip2 := binary.BigEndian.Uint32(a.ip[4:8])
	ip3 := binary.BigEndian.Uint32(a.ip[8:12])
	ip4 := binary.BigEndian.Uint32(a.ip[12:16])

	totalLen += util.EncodedLengthByVi32(int32(ip1))
	totalLen += util.EncodedLengthByVi32(int32(ip2))
	totalLen += util.EncodedLengthByVi32(int32(ip3))
	totalLen += util.EncodedLengthByVi32(int32(ip4))

	totalLen += util.EncodedLengthByVi32(a.port)

	a.ObUniVersionHeader.SetContentLength(totalLen)
	return a.ObUniVersionHeader.ContentLength()
}

func (a *ObAddr) Encode(buffer *bytes.Buffer) {
	a.ObUniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, uint8(a.version))

	ip1 := binary.BigEndian.Uint32(a.ip[0:4])
	ip2 := binary.BigEndian.Uint32(a.ip[4:8])
	ip3 := binary.BigEndian.Uint32(a.ip[8:12])
	ip4 := binary.BigEndian.Uint32(a.ip[12:16])

	util.EncodeVi32(buffer, int32(ip1))
	util.EncodeVi32(buffer, int32(ip2))
	util.EncodeVi32(buffer, int32(ip3))
	util.EncodeVi32(buffer, int32(ip4))

	util.EncodeVi32(buffer, a.port)
}

func (a *ObAddr) Decode(buffer *bytes.Buffer) {
	a.ObUniVersionHeader.Decode(buffer)

	a.version = ObAddrVersion(util.Uint8(buffer))

	var ip1, ip2, ip3, ip4 uint32
	ip1 = uint32(util.DecodeVi32(buffer))
	ip2 = uint32(util.DecodeVi32(buffer))
	ip3 = uint32(util.DecodeVi32(buffer))
	ip4 = uint32(util.DecodeVi32(buffer))
	binary.BigEndian.PutUint32(a.ip[0:], ip1)
	binary.BigEndian.PutUint32(a.ip[4:], ip2)
	binary.BigEndian.PutUint32(a.ip[8:], ip3)
	binary.BigEndian.PutUint32(a.ip[12:], ip4)

	a.port = util.DecodeVi32(buffer)
}

func (a *ObAddr) String() string {
	var ObUniVersionHeaderStr = "nil"
	if a.ObUniVersionHeader != (ObUniVersionHeader{}) {
		ObUniVersionHeaderStr = a.ObUniVersionHeader.String()
	}

	ipAddr := net.IP(a.ip[:])
	return "ObAddr{" +
		"ObUniVersionHeader:" + ObUniVersionHeaderStr + ", " +
		"version:" + strconv.Itoa(int(a.version)) + ", " +
		"ip:" + ipAddr.String() + ", " +
		"port:" + strconv.Itoa(int(a.port)) +
		"}"
}
