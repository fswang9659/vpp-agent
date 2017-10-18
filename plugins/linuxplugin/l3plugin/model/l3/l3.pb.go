// Code generated by protoc-gen-gogo.
// source: l3.proto
// DO NOT EDIT!

/*
Package l3 is a generated protocol buffer package.

It is generated from these files:
	l3.proto

It has these top-level messages:
	LinuxStaticRoutes
	LinuxStaticArpEntries
*/
package l3

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type LinuxStaticRoutes_Route_Namespace_NamespaceType int32

const (
	LinuxStaticRoutes_Route_Namespace_PID_REF_NS          LinuxStaticRoutes_Route_Namespace_NamespaceType = 0
	LinuxStaticRoutes_Route_Namespace_MICROSERVICE_REF_NS LinuxStaticRoutes_Route_Namespace_NamespaceType = 1
	LinuxStaticRoutes_Route_Namespace_NAMED_NS            LinuxStaticRoutes_Route_Namespace_NamespaceType = 2
	LinuxStaticRoutes_Route_Namespace_FILE_REF_NS         LinuxStaticRoutes_Route_Namespace_NamespaceType = 3
)

var LinuxStaticRoutes_Route_Namespace_NamespaceType_name = map[int32]string{
	0: "PID_REF_NS",
	1: "MICROSERVICE_REF_NS",
	2: "NAMED_NS",
	3: "FILE_REF_NS",
}
var LinuxStaticRoutes_Route_Namespace_NamespaceType_value = map[string]int32{
	"PID_REF_NS":          0,
	"MICROSERVICE_REF_NS": 1,
	"NAMED_NS":            2,
	"FILE_REF_NS":         3,
}

func (x LinuxStaticRoutes_Route_Namespace_NamespaceType) String() string {
	return proto.EnumName(LinuxStaticRoutes_Route_Namespace_NamespaceType_name, int32(x))
}

type LinuxStaticRoutes_Route_Scope_ScopeType int32

const (
	LinuxStaticRoutes_Route_Scope_GLOBAL LinuxStaticRoutes_Route_Scope_ScopeType = 0
	LinuxStaticRoutes_Route_Scope_SITE   LinuxStaticRoutes_Route_Scope_ScopeType = 1
	LinuxStaticRoutes_Route_Scope_LINK   LinuxStaticRoutes_Route_Scope_ScopeType = 2
	LinuxStaticRoutes_Route_Scope_HOST   LinuxStaticRoutes_Route_Scope_ScopeType = 3
)

var LinuxStaticRoutes_Route_Scope_ScopeType_name = map[int32]string{
	0: "GLOBAL",
	1: "SITE",
	2: "LINK",
	3: "HOST",
}
var LinuxStaticRoutes_Route_Scope_ScopeType_value = map[string]int32{
	"GLOBAL": 0,
	"SITE":   1,
	"LINK":   2,
	"HOST":   3,
}

func (x LinuxStaticRoutes_Route_Scope_ScopeType) String() string {
	return proto.EnumName(LinuxStaticRoutes_Route_Scope_ScopeType_name, int32(x))
}

type LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType int32

const (
	LinuxStaticArpEntries_ArpEntry_Namespace_PID_REF_NS          LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType = 0
	LinuxStaticArpEntries_ArpEntry_Namespace_MICROSERVICE_REF_NS LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType = 1
	LinuxStaticArpEntries_ArpEntry_Namespace_NAMED_NS            LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType = 2
	LinuxStaticArpEntries_ArpEntry_Namespace_FILE_REF_NS         LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType = 3
)

var LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType_name = map[int32]string{
	0: "PID_REF_NS",
	1: "MICROSERVICE_REF_NS",
	2: "NAMED_NS",
	3: "FILE_REF_NS",
}
var LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType_value = map[string]int32{
	"PID_REF_NS":          0,
	"MICROSERVICE_REF_NS": 1,
	"NAMED_NS":            2,
	"FILE_REF_NS":         3,
}

func (x LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType) String() string {
	return proto.EnumName(LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType_name, int32(x))
}

type LinuxStaticArpEntries_ArpEntry_NudState_NudStateType int32

const (
	LinuxStaticArpEntries_ArpEntry_NudState_PERMANENT LinuxStaticArpEntries_ArpEntry_NudState_NudStateType = 0
	LinuxStaticArpEntries_ArpEntry_NudState_NOARP     LinuxStaticArpEntries_ArpEntry_NudState_NudStateType = 1
	LinuxStaticArpEntries_ArpEntry_NudState_REACHABLE LinuxStaticArpEntries_ArpEntry_NudState_NudStateType = 2
	LinuxStaticArpEntries_ArpEntry_NudState_STALE     LinuxStaticArpEntries_ArpEntry_NudState_NudStateType = 3
)

var LinuxStaticArpEntries_ArpEntry_NudState_NudStateType_name = map[int32]string{
	0: "PERMANENT",
	1: "NOARP",
	2: "REACHABLE",
	3: "STALE",
}
var LinuxStaticArpEntries_ArpEntry_NudState_NudStateType_value = map[string]int32{
	"PERMANENT": 0,
	"NOARP":     1,
	"REACHABLE": 2,
	"STALE":     3,
}

func (x LinuxStaticArpEntries_ArpEntry_NudState_NudStateType) String() string {
	return proto.EnumName(LinuxStaticArpEntries_ArpEntry_NudState_NudStateType_name, int32(x))
}

// static ip routes
type LinuxStaticRoutes struct {
	Route []*LinuxStaticRoutes_Route `protobuf:"bytes,1,rep,name=route" json:"route,omitempty"`
}

func (m *LinuxStaticRoutes) Reset()         { *m = LinuxStaticRoutes{} }
func (m *LinuxStaticRoutes) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticRoutes) ProtoMessage()    {}

func (m *LinuxStaticRoutes) GetRoute() []*LinuxStaticRoutes_Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type LinuxStaticRoutes_Route struct {
	Name        string                             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   *LinuxStaticRoutes_Route_Namespace `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Interface   string                             `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	Description string                             `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Scope       *LinuxStaticRoutes_Route_Scope     `protobuf:"bytes,5,opt,name=scope" json:"scope,omitempty"`
	SrcIpAddr   string                             `protobuf:"bytes,6,opt,name=src_ip_addr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr   string                             `protobuf:"bytes,7,opt,name=dst_ip_addr,proto3" json:"dst_ip_addr,omitempty"`
	GwAddr      string                             `protobuf:"bytes,8,opt,name=gw_addr,proto3" json:"gw_addr,omitempty"`
	Priority    uint32                             `protobuf:"varint,9,opt,name=priority,proto3" json:"priority,omitempty"`
	Table       uint32                             `protobuf:"varint,10,opt,name=table,proto3" json:"table,omitempty"`
}

func (m *LinuxStaticRoutes_Route) Reset()         { *m = LinuxStaticRoutes_Route{} }
func (m *LinuxStaticRoutes_Route) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticRoutes_Route) ProtoMessage()    {}

func (m *LinuxStaticRoutes_Route) GetNamespace() *LinuxStaticRoutes_Route_Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *LinuxStaticRoutes_Route) GetScope() *LinuxStaticRoutes_Route_Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

type LinuxStaticRoutes_Route_Namespace struct {
	Type         LinuxStaticRoutes_Route_Namespace_NamespaceType `protobuf:"varint,1,opt,name=type,proto3,enum=l3.LinuxStaticRoutes_Route_Namespace_NamespaceType" json:"type,omitempty"`
	Pid          uint32                                          `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Microservice string                                          `protobuf:"bytes,3,opt,name=microservice,proto3" json:"microservice,omitempty"`
	Name         string                                          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Filepath     string                                          `protobuf:"bytes,5,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (m *LinuxStaticRoutes_Route_Namespace) Reset()         { *m = LinuxStaticRoutes_Route_Namespace{} }
func (m *LinuxStaticRoutes_Route_Namespace) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticRoutes_Route_Namespace) ProtoMessage()    {}

type LinuxStaticRoutes_Route_Scope struct {
	Type LinuxStaticRoutes_Route_Scope_ScopeType `protobuf:"varint,1,opt,name=type,proto3,enum=l3.LinuxStaticRoutes_Route_Scope_ScopeType" json:"type,omitempty"`
}

func (m *LinuxStaticRoutes_Route_Scope) Reset()         { *m = LinuxStaticRoutes_Route_Scope{} }
func (m *LinuxStaticRoutes_Route_Scope) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticRoutes_Route_Scope) ProtoMessage()    {}

// static arp entires
type LinuxStaticArpEntries struct {
	ArpEntry []*LinuxStaticArpEntries_ArpEntry `protobuf:"bytes,1,rep,name=arp_entry" json:"arp_entry,omitempty"`
}

func (m *LinuxStaticArpEntries) Reset()         { *m = LinuxStaticArpEntries{} }
func (m *LinuxStaticArpEntries) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticArpEntries) ProtoMessage()    {}

func (m *LinuxStaticArpEntries) GetArpEntry() []*LinuxStaticArpEntries_ArpEntry {
	if m != nil {
		return m.ArpEntry
	}
	return nil
}

type LinuxStaticArpEntries_ArpEntry struct {
	Name      string                                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace *LinuxStaticArpEntries_ArpEntry_Namespace `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Interface string                                    `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	State     *LinuxStaticArpEntries_ArpEntry_NudState  `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	IpAddr    string                                    `protobuf:"bytes,5,opt,name=ip_addr,proto3" json:"ip_addr,omitempty"`
	HwAddress string                                    `protobuf:"bytes,6,opt,name=hw_address,proto3" json:"hw_address,omitempty"`
}

func (m *LinuxStaticArpEntries_ArpEntry) Reset()         { *m = LinuxStaticArpEntries_ArpEntry{} }
func (m *LinuxStaticArpEntries_ArpEntry) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticArpEntries_ArpEntry) ProtoMessage()    {}

func (m *LinuxStaticArpEntries_ArpEntry) GetNamespace() *LinuxStaticArpEntries_ArpEntry_Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *LinuxStaticArpEntries_ArpEntry) GetState() *LinuxStaticArpEntries_ArpEntry_NudState {
	if m != nil {
		return m.State
	}
	return nil
}

type LinuxStaticArpEntries_ArpEntry_Namespace struct {
	Type         LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType `protobuf:"varint,1,opt,name=type,proto3,enum=l3.LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType" json:"type,omitempty"`
	Pid          uint32                                                 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Microservice string                                                 `protobuf:"bytes,3,opt,name=microservice,proto3" json:"microservice,omitempty"`
	Name         string                                                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Filepath     string                                                 `protobuf:"bytes,5,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (m *LinuxStaticArpEntries_ArpEntry_Namespace) Reset() {
	*m = LinuxStaticArpEntries_ArpEntry_Namespace{}
}
func (m *LinuxStaticArpEntries_ArpEntry_Namespace) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticArpEntries_ArpEntry_Namespace) ProtoMessage()    {}

type LinuxStaticArpEntries_ArpEntry_NudState struct {
	Type LinuxStaticArpEntries_ArpEntry_NudState_NudStateType `protobuf:"varint,1,opt,name=type,proto3,enum=l3.LinuxStaticArpEntries_ArpEntry_NudState_NudStateType" json:"type,omitempty"`
}

func (m *LinuxStaticArpEntries_ArpEntry_NudState) Reset() {
	*m = LinuxStaticArpEntries_ArpEntry_NudState{}
}
func (m *LinuxStaticArpEntries_ArpEntry_NudState) String() string { return proto.CompactTextString(m) }
func (*LinuxStaticArpEntries_ArpEntry_NudState) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("l3.LinuxStaticRoutes_Route_Namespace_NamespaceType", LinuxStaticRoutes_Route_Namespace_NamespaceType_name, LinuxStaticRoutes_Route_Namespace_NamespaceType_value)
	proto.RegisterEnum("l3.LinuxStaticRoutes_Route_Scope_ScopeType", LinuxStaticRoutes_Route_Scope_ScopeType_name, LinuxStaticRoutes_Route_Scope_ScopeType_value)
	proto.RegisterEnum("l3.LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType", LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType_name, LinuxStaticArpEntries_ArpEntry_Namespace_NamespaceType_value)
	proto.RegisterEnum("l3.LinuxStaticArpEntries_ArpEntry_NudState_NudStateType", LinuxStaticArpEntries_ArpEntry_NudState_NudStateType_name, LinuxStaticArpEntries_ArpEntry_NudState_NudStateType_value)
}
