// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/l3/teib.proto

package vpp_l3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TeibEntry represents an tunnel endpoint information base entry.
type TeibEntry struct {
	// Interface references a tunnel interface this TEIB entry is linked to.
	Interface string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	// IP address of the peer.
	PeerAddr string `protobuf:"bytes,2,opt,name=peer_addr,json=peerAddr,proto3" json:"peer_addr,omitempty"`
	// Next hop IP address.
	NextHopAddr string `protobuf:"bytes,3,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	// VRF ID used to reach the next hop.
	VrfId                uint32   `protobuf:"varint,4,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeibEntry) Reset()         { *m = TeibEntry{} }
func (m *TeibEntry) String() string { return proto.CompactTextString(m) }
func (*TeibEntry) ProtoMessage()    {}
func (*TeibEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_43413ef30fbb968b, []int{0}
}

func (m *TeibEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeibEntry.Unmarshal(m, b)
}
func (m *TeibEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeibEntry.Marshal(b, m, deterministic)
}
func (m *TeibEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeibEntry.Merge(m, src)
}
func (m *TeibEntry) XXX_Size() int {
	return xxx_messageInfo_TeibEntry.Size(m)
}
func (m *TeibEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TeibEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TeibEntry proto.InternalMessageInfo

func (m *TeibEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *TeibEntry) GetPeerAddr() string {
	if m != nil {
		return m.PeerAddr
	}
	return ""
}

func (m *TeibEntry) GetNextHopAddr() string {
	if m != nil {
		return m.NextHopAddr
	}
	return ""
}

func (m *TeibEntry) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func init() {
	proto.RegisterType((*TeibEntry)(nil), "ligato.vpp.l3.TeibEntry")
}

func init() { proto.RegisterFile("ligato/vpp/l3/teib.proto", fileDescriptor_43413ef30fbb968b) }

var fileDescriptor_43413ef30fbb968b = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3d, 0x4b, 0x04, 0x31,
	0x14, 0x45, 0x89, 0x1f, 0x8b, 0x89, 0x4c, 0x13, 0x10, 0x02, 0x5a, 0x2c, 0x5b, 0x6d, 0x63, 0x52,
	0x44, 0x6c, 0xac, 0x14, 0x04, 0x6d, 0x17, 0x2b, 0x9b, 0x90, 0x31, 0x6f, 0xc6, 0x40, 0x48, 0x1e,
	0x8f, 0x10, 0xb4, 0xf4, 0x9f, 0xcb, 0x64, 0x04, 0xd9, 0xf6, 0x9e, 0x53, 0x9c, 0x2b, 0x54, 0x8a,
	0xb3, 0xaf, 0xc5, 0x34, 0x44, 0x93, 0xac, 0xa9, 0x10, 0x47, 0x8d, 0x54, 0x6a, 0x91, 0xc3, 0x4a,
	0x74, 0x43, 0xd4, 0xc9, 0xee, 0x7e, 0x98, 0xe0, 0x6f, 0x10, 0xc7, 0xe7, 0x5c, 0xe9, 0x5b, 0xde,
	0x08, 0x1e, 0x73, 0x05, 0x9a, 0xfc, 0x07, 0x28, 0xb6, 0x65, 0x7b, 0x7e, 0xf8, 0x1f, 0xe4, 0xb5,
	0xe0, 0x08, 0x40, 0xce, 0x87, 0x40, 0xea, 0xa4, 0xd3, 0x8b, 0x65, 0x78, 0x0c, 0x81, 0xe4, 0x4e,
	0x0c, 0x19, 0xbe, 0xaa, 0xfb, 0x2c, 0xb8, 0x0a, 0xa7, 0x5d, 0xb8, 0x5c, 0xc6, 0x97, 0x82, 0xdd,
	0xb9, 0x12, 0x9b, 0x46, 0x93, 0x8b, 0x41, 0x9d, 0x6d, 0xd9, 0x7e, 0x38, 0x9c, 0x37, 0x9a, 0x5e,
	0xc3, 0xd3, 0xfd, 0xfb, 0xdd, 0x5c, 0xf4, 0x5f, 0x57, 0xec, 0xd1, 0xb7, 0x7e, 0x86, 0x5c, 0x4d,
	0xb3, 0xa6, 0x57, 0x9b, 0xa3, 0x3b, 0x0f, 0x0d, 0xd1, 0x25, 0x3b, 0x6e, 0x3a, 0xb3, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x65, 0xf5, 0xe1, 0xfd, 0xed, 0x00, 0x00, 0x00,
}