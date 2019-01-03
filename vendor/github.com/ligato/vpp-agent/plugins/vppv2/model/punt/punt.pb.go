// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: punt.proto

package punt

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// L3 protocol definition
type L3Protocol int32

const (
	L3Protocol_UNDEFINED_L3 L3Protocol = 0
	L3Protocol_IPv4         L3Protocol = 4
	L3Protocol_IPv6         L3Protocol = 6
	L3Protocol_ALL          L3Protocol = 10
)

var L3Protocol_name = map[int32]string{
	0:  "UNDEFINED_L3",
	4:  "IPv4",
	6:  "IPv6",
	10: "ALL",
}
var L3Protocol_value = map[string]int32{
	"UNDEFINED_L3": 0,
	"IPv4":         4,
	"IPv6":         6,
	"ALL":          10,
}

func (x L3Protocol) String() string {
	return proto.EnumName(L3Protocol_name, int32(x))
}
func (L3Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_punt_38189aabf5c7b0cf, []int{0}
}

// L4 protocol definition
type L4Protocol int32

const (
	L4Protocol_UNDEFINED_L4 L4Protocol = 0
	L4Protocol_TCP          L4Protocol = 6
	L4Protocol_UDP          L4Protocol = 17
)

var L4Protocol_name = map[int32]string{
	0:  "UNDEFINED_L4",
	6:  "TCP",
	17: "UDP",
}
var L4Protocol_value = map[string]int32{
	"UNDEFINED_L4": 0,
	"TCP":          6,
	"UDP":          17,
}

func (x L4Protocol) String() string {
	return proto.EnumName(L4Protocol_name, int32(x))
}
func (L4Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_punt_38189aabf5c7b0cf, []int{1}
}

// IpRedirect allows otherwise dropped packet which destination IP address matching some of the VPP addresses
// to redirect to the defined next hop address via the TX interface
type IpRedirect struct {
	L3Protocol           L3Protocol `protobuf:"varint,1,opt,name=l3_protocol,json=l3Protocol,proto3,enum=punt.L3Protocol" json:"l3_protocol,omitempty"`
	RxInterface          string     `protobuf:"bytes,2,opt,name=rx_interface,json=rxInterface,proto3" json:"rx_interface,omitempty"`
	TxInterface          string     `protobuf:"bytes,3,opt,name=tx_interface,json=txInterface,proto3" json:"tx_interface,omitempty"`
	NextHop              string     `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IpRedirect) Reset()         { *m = IpRedirect{} }
func (m *IpRedirect) String() string { return proto.CompactTextString(m) }
func (*IpRedirect) ProtoMessage()    {}
func (*IpRedirect) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_38189aabf5c7b0cf, []int{0}
}
func (m *IpRedirect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpRedirect.Unmarshal(m, b)
}
func (m *IpRedirect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpRedirect.Marshal(b, m, deterministic)
}
func (dst *IpRedirect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpRedirect.Merge(dst, src)
}
func (m *IpRedirect) XXX_Size() int {
	return xxx_messageInfo_IpRedirect.Size(m)
}
func (m *IpRedirect) XXX_DiscardUnknown() {
	xxx_messageInfo_IpRedirect.DiscardUnknown(m)
}

var xxx_messageInfo_IpRedirect proto.InternalMessageInfo

func (m *IpRedirect) GetL3Protocol() L3Protocol {
	if m != nil {
		return m.L3Protocol
	}
	return L3Protocol_UNDEFINED_L3
}

func (m *IpRedirect) GetRxInterface() string {
	if m != nil {
		return m.RxInterface
	}
	return ""
}

func (m *IpRedirect) GetTxInterface() string {
	if m != nil {
		return m.TxInterface
	}
	return ""
}

func (m *IpRedirect) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

// allows otherwise dropped packet which destination IP address matching some of the VPP interface IP addresses to be
// punted to the host. L3 and L4 protocols can be used for filtering
type ToHost struct {
	L3Protocol           L3Protocol `protobuf:"varint,2,opt,name=l3_protocol,json=l3Protocol,proto3,enum=punt.L3Protocol" json:"l3_protocol,omitempty"`
	L4Protocol           L4Protocol `protobuf:"varint,3,opt,name=l4_protocol,json=l4Protocol,proto3,enum=punt.L4Protocol" json:"l4_protocol,omitempty"`
	Port                 uint32     `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	SocketPath           string     `protobuf:"bytes,5,opt,name=socket_path,json=socketPath,proto3" json:"socket_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ToHost) Reset()         { *m = ToHost{} }
func (m *ToHost) String() string { return proto.CompactTextString(m) }
func (*ToHost) ProtoMessage()    {}
func (*ToHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_38189aabf5c7b0cf, []int{1}
}
func (m *ToHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToHost.Unmarshal(m, b)
}
func (m *ToHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToHost.Marshal(b, m, deterministic)
}
func (dst *ToHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToHost.Merge(dst, src)
}
func (m *ToHost) XXX_Size() int {
	return xxx_messageInfo_ToHost.Size(m)
}
func (m *ToHost) XXX_DiscardUnknown() {
	xxx_messageInfo_ToHost.DiscardUnknown(m)
}

var xxx_messageInfo_ToHost proto.InternalMessageInfo

func (m *ToHost) GetL3Protocol() L3Protocol {
	if m != nil {
		return m.L3Protocol
	}
	return L3Protocol_UNDEFINED_L3
}

func (m *ToHost) GetL4Protocol() L4Protocol {
	if m != nil {
		return m.L4Protocol
	}
	return L4Protocol_UNDEFINED_L4
}

func (m *ToHost) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ToHost) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func init() {
	proto.RegisterType((*IpRedirect)(nil), "punt.IpRedirect")
	proto.RegisterType((*ToHost)(nil), "punt.ToHost")
	proto.RegisterEnum("punt.L3Protocol", L3Protocol_name, L3Protocol_value)
	proto.RegisterEnum("punt.L4Protocol", L4Protocol_name, L4Protocol_value)
}

func init() { proto.RegisterFile("punt.proto", fileDescriptor_punt_38189aabf5c7b0cf) }

var fileDescriptor_punt_38189aabf5c7b0cf = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xb1, 0x4f, 0x83, 0x40,
	0x14, 0x87, 0x4b, 0x41, 0x5a, 0x1f, 0xd5, 0x9c, 0x37, 0xe1, 0x64, 0xed, 0xd4, 0x74, 0x68, 0x54,
	0x88, 0x8b, 0x93, 0x91, 0x9a, 0x92, 0x90, 0x86, 0x90, 0x76, 0x26, 0x88, 0x67, 0x68, 0x24, 0xdc,
	0xe5, 0xfa, 0x34, 0xfc, 0x39, 0x26, 0xfe, 0xa3, 0xe6, 0x8e, 0x5a, 0x50, 0x17, 0xb7, 0x1f, 0x8f,
	0xef, 0x4b, 0xbe, 0x03, 0x10, 0x6f, 0x15, 0xce, 0x85, 0xe4, 0xc8, 0xa9, 0xa5, 0xf6, 0xe4, 0xc3,
	0x00, 0x08, 0x45, 0xc2, 0x9e, 0xb7, 0x92, 0xe5, 0x48, 0xaf, 0xc1, 0x29, 0xbd, 0x54, 0x03, 0x39,
	0x2f, 0x5d, 0x63, 0x6c, 0x4c, 0x4f, 0x6f, 0xc8, 0x5c, 0x6b, 0x91, 0x17, 0xef, 0xef, 0x09, 0x94,
	0x87, 0x4d, 0x2f, 0x61, 0x24, 0xeb, 0x74, 0x5b, 0x21, 0x93, 0x2f, 0x59, 0xce, 0xdc, 0xfe, 0xd8,
	0x98, 0x1e, 0x27, 0x8e, 0xac, 0xc3, 0xef, 0x93, 0x42, 0xb0, 0x8b, 0x98, 0x0d, 0x82, 0x1d, 0xe4,
	0x1c, 0x86, 0x15, 0xab, 0x31, 0x2d, 0xb8, 0x70, 0x2d, 0xfd, 0x7b, 0xa0, 0xbe, 0x97, 0x5c, 0x4c,
	0x3e, 0x0d, 0xb0, 0xd7, 0x7c, 0xc9, 0x77, 0x7f, 0xf2, 0xfa, 0xff, 0xc8, 0x53, 0x8a, 0xdf, 0x2a,
	0xe6, 0x0f, 0xc5, 0xef, 0x28, 0x87, 0x4d, 0x29, 0x58, 0x82, 0x4b, 0xd4, 0x1d, 0x27, 0x89, 0xde,
	0xf4, 0x02, 0x9c, 0x1d, 0xcf, 0x5f, 0x19, 0xa6, 0x22, 0xc3, 0xc2, 0x3d, 0xd2, 0x89, 0xd0, 0x9c,
	0xe2, 0x0c, 0x8b, 0xd9, 0x1d, 0x40, 0x5b, 0x40, 0x09, 0x8c, 0x36, 0xab, 0x60, 0xf1, 0x18, 0xae,
	0x16, 0x41, 0x1a, 0x79, 0xa4, 0x47, 0x87, 0x60, 0x85, 0xf1, 0xbb, 0x4f, 0xac, 0xfd, 0xba, 0x25,
	0x36, 0x1d, 0x80, 0x79, 0x1f, 0x45, 0x04, 0x66, 0x57, 0x00, 0x6d, 0xcb, 0x2f, 0xd9, 0x27, 0x3d,
	0x05, 0xae, 0x1f, 0xe2, 0xc6, 0xd8, 0x04, 0x31, 0x39, 0x7b, 0xb2, 0xf5, 0x8b, 0xbc, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x43, 0x8e, 0xc9, 0xd2, 0x01, 0x00, 0x00,
}
