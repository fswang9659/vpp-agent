// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/fib_types.api.json

/*
Package fib_types is a generated VPP binary API for 'fib_types' module.

It consists of:
	  7 enums
	  5 aliases
	  9 types
	  1 union
*/
package fib_types

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"

	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2005/ip_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "fib_types"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd6a5938
)

type AddressFamily = ip_types.AddressFamily

// FibPathFlags represents VPP binary API enum 'fib_path_flags'.
type FibPathFlags uint32

const (
	FIB_API_PATH_FLAG_NONE                 FibPathFlags = 0
	FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED FibPathFlags = 1
	FIB_API_PATH_FLAG_RESOLVE_VIA_HOST     FibPathFlags = 2
	FIB_API_PATH_FLAG_POP_PW_CW            FibPathFlags = 4
)

var FibPathFlags_name = map[uint32]string{
	0: "FIB_API_PATH_FLAG_NONE",
	1: "FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED",
	2: "FIB_API_PATH_FLAG_RESOLVE_VIA_HOST",
	4: "FIB_API_PATH_FLAG_POP_PW_CW",
}

var FibPathFlags_value = map[string]uint32{
	"FIB_API_PATH_FLAG_NONE":                 0,
	"FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED": 1,
	"FIB_API_PATH_FLAG_RESOLVE_VIA_HOST":     2,
	"FIB_API_PATH_FLAG_POP_PW_CW":            4,
}

func (x FibPathFlags) String() string {
	s, ok := FibPathFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// FibPathNhProto represents VPP binary API enum 'fib_path_nh_proto'.
type FibPathNhProto uint32

const (
	FIB_API_PATH_NH_PROTO_IP4      FibPathNhProto = 0
	FIB_API_PATH_NH_PROTO_IP6      FibPathNhProto = 1
	FIB_API_PATH_NH_PROTO_MPLS     FibPathNhProto = 2
	FIB_API_PATH_NH_PROTO_ETHERNET FibPathNhProto = 3
	FIB_API_PATH_NH_PROTO_BIER     FibPathNhProto = 4
)

var FibPathNhProto_name = map[uint32]string{
	0: "FIB_API_PATH_NH_PROTO_IP4",
	1: "FIB_API_PATH_NH_PROTO_IP6",
	2: "FIB_API_PATH_NH_PROTO_MPLS",
	3: "FIB_API_PATH_NH_PROTO_ETHERNET",
	4: "FIB_API_PATH_NH_PROTO_BIER",
}

var FibPathNhProto_value = map[string]uint32{
	"FIB_API_PATH_NH_PROTO_IP4":      0,
	"FIB_API_PATH_NH_PROTO_IP6":      1,
	"FIB_API_PATH_NH_PROTO_MPLS":     2,
	"FIB_API_PATH_NH_PROTO_ETHERNET": 3,
	"FIB_API_PATH_NH_PROTO_BIER":     4,
}

func (x FibPathNhProto) String() string {
	s, ok := FibPathNhProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// FibPathType represents VPP binary API enum 'fib_path_type'.
type FibPathType uint32

const (
	FIB_API_PATH_TYPE_NORMAL        FibPathType = 0
	FIB_API_PATH_TYPE_LOCAL         FibPathType = 1
	FIB_API_PATH_TYPE_DROP          FibPathType = 2
	FIB_API_PATH_TYPE_UDP_ENCAP     FibPathType = 3
	FIB_API_PATH_TYPE_BIER_IMP      FibPathType = 4
	FIB_API_PATH_TYPE_ICMP_UNREACH  FibPathType = 5
	FIB_API_PATH_TYPE_ICMP_PROHIBIT FibPathType = 6
	FIB_API_PATH_TYPE_SOURCE_LOOKUP FibPathType = 7
	FIB_API_PATH_TYPE_DVR           FibPathType = 8
	FIB_API_PATH_TYPE_INTERFACE_RX  FibPathType = 9
	FIB_API_PATH_TYPE_CLASSIFY      FibPathType = 10
)

var FibPathType_name = map[uint32]string{
	0:  "FIB_API_PATH_TYPE_NORMAL",
	1:  "FIB_API_PATH_TYPE_LOCAL",
	2:  "FIB_API_PATH_TYPE_DROP",
	3:  "FIB_API_PATH_TYPE_UDP_ENCAP",
	4:  "FIB_API_PATH_TYPE_BIER_IMP",
	5:  "FIB_API_PATH_TYPE_ICMP_UNREACH",
	6:  "FIB_API_PATH_TYPE_ICMP_PROHIBIT",
	7:  "FIB_API_PATH_TYPE_SOURCE_LOOKUP",
	8:  "FIB_API_PATH_TYPE_DVR",
	9:  "FIB_API_PATH_TYPE_INTERFACE_RX",
	10: "FIB_API_PATH_TYPE_CLASSIFY",
}

var FibPathType_value = map[string]uint32{
	"FIB_API_PATH_TYPE_NORMAL":        0,
	"FIB_API_PATH_TYPE_LOCAL":         1,
	"FIB_API_PATH_TYPE_DROP":          2,
	"FIB_API_PATH_TYPE_UDP_ENCAP":     3,
	"FIB_API_PATH_TYPE_BIER_IMP":      4,
	"FIB_API_PATH_TYPE_ICMP_UNREACH":  5,
	"FIB_API_PATH_TYPE_ICMP_PROHIBIT": 6,
	"FIB_API_PATH_TYPE_SOURCE_LOOKUP": 7,
	"FIB_API_PATH_TYPE_DVR":           8,
	"FIB_API_PATH_TYPE_INTERFACE_RX":  9,
	"FIB_API_PATH_TYPE_CLASSIFY":      10,
}

func (x FibPathType) String() string {
	s, ok := FibPathType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type IPDscp = ip_types.IPDscp

type IPEcn = ip_types.IPEcn

type IPProto = ip_types.IPProto

type AddressWithPrefix = ip_types.AddressWithPrefix

type IP4Address = ip_types.IP4Address

type IP4AddressWithPrefix = ip_types.IP4AddressWithPrefix

type IP6Address = ip_types.IP6Address

type IP6AddressWithPrefix = ip_types.IP6AddressWithPrefix

type Address = ip_types.Address

// FibMplsLabel represents VPP binary API type 'fib_mpls_label'.
type FibMplsLabel struct {
	IsUniform uint8
	Label     uint32
	TTL       uint8
	Exp       uint8
}

func (*FibMplsLabel) GetTypeName() string { return "fib_mpls_label" }

// FibPath represents VPP binary API type 'fib_path'.
type FibPath struct {
	SwIfIndex  uint32
	TableID    uint32
	RpfID      uint32
	Weight     uint8
	Preference uint8
	Type       FibPathType
	Flags      FibPathFlags
	Proto      FibPathNhProto
	Nh         FibPathNh
	NLabels    uint8 `struc:"sizeof=LabelStack"` // MANUALLY FIXED, see https://jira.fd.io/browse/VPP-1261
	LabelStack []FibMplsLabel
}

func (*FibPath) GetTypeName() string { return "fib_path" }

// FibPathNh represents VPP binary API type 'fib_path_nh'.
type FibPathNh struct {
	Address            AddressUnion
	ViaLabel           uint32
	ObjID              uint32
	ClassifyTableIndex uint32
}

func (*FibPathNh) GetTypeName() string { return "fib_path_nh" }

type IP4Prefix = ip_types.IP4Prefix

type IP6Prefix = ip_types.IP6Prefix

type Mprefix = ip_types.Mprefix

type Prefix = ip_types.Prefix

type PrefixMatcher = ip_types.PrefixMatcher

type AddressUnion = ip_types.AddressUnion

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
