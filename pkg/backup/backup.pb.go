// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: backup.proto

package backup

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

	errorpb "github.com/pingcap/kvproto/pkg/errorpb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BackupEvent_Event int32

const (
	BackupEvent_Unknown       BackupEvent_Event = 0
	BackupEvent_Snapshot      BackupEvent_Event = 1
	BackupEvent_Split         BackupEvent_Event = 2
	BackupEvent_PrepareMerge  BackupEvent_Event = 3
	BackupEvent_CommitMerge   BackupEvent_Event = 4
	BackupEvent_RollbackMerge BackupEvent_Event = 5
)

var BackupEvent_Event_name = map[int32]string{
	0: "Unknown",
	1: "Snapshot",
	2: "Split",
	3: "PrepareMerge",
	4: "CommitMerge",
	5: "RollbackMerge",
}
var BackupEvent_Event_value = map[string]int32{
	"Unknown":       0,
	"Snapshot":      1,
	"Split":         2,
	"PrepareMerge":  3,
	"CommitMerge":   4,
	"RollbackMerge": 5,
}

func (x BackupEvent_Event) String() string {
	return proto.EnumName(BackupEvent_Event_name, int32(x))
}
func (BackupEvent_Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_backup_5333a12533fb1467, []int{0, 0}
}

type BackupEvent struct {
	RegionId             uint64            `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	RelatedRegionIds     []uint64          `protobuf:"varint,2,rep,packed,name=related_region_ids,json=relatedRegionIds" json:"related_region_ids,omitempty"`
	Event                BackupEvent_Event `protobuf:"varint,3,opt,name=event,proto3,enum=backup.BackupEvent_Event" json:"event,omitempty"`
	Dependency           uint64            `protobuf:"varint,4,opt,name=dependency,proto3" json:"dependency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BackupEvent) Reset()         { *m = BackupEvent{} }
func (m *BackupEvent) String() string { return proto.CompactTextString(m) }
func (*BackupEvent) ProtoMessage()    {}
func (*BackupEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_5333a12533fb1467, []int{0}
}
func (m *BackupEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupEvent.Merge(dst, src)
}
func (m *BackupEvent) XXX_Size() int {
	return m.Size()
}
func (m *BackupEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BackupEvent proto.InternalMessageInfo

func (m *BackupEvent) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *BackupEvent) GetRelatedRegionIds() []uint64 {
	if m != nil {
		return m.RelatedRegionIds
	}
	return nil
}

func (m *BackupEvent) GetEvent() BackupEvent_Event {
	if m != nil {
		return m.Event
	}
	return BackupEvent_Unknown
}

func (m *BackupEvent) GetDependency() uint64 {
	if m != nil {
		return m.Dependency
	}
	return 0
}

type BackupMeta struct {
	Events                []*BackupEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	StartBackupDependency uint64         `protobuf:"varint,2,opt,name=start_backup_dependency,json=startBackupDependency,proto3" json:"start_backup_dependency,omitempty"`
	FullBackupDependency  uint64         `protobuf:"varint,3,opt,name=full_backup_dependency,json=fullBackupDependency,proto3" json:"full_backup_dependency,omitempty"`
	IncBackupDependencies []uint64       `protobuf:"varint,4,rep,packed,name=inc_backup_dependencies,json=incBackupDependencies" json:"inc_backup_dependencies,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *BackupMeta) Reset()         { *m = BackupMeta{} }
func (m *BackupMeta) String() string { return proto.CompactTextString(m) }
func (*BackupMeta) ProtoMessage()    {}
func (*BackupMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_5333a12533fb1467, []int{1}
}
func (m *BackupMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupMeta.Merge(dst, src)
}
func (m *BackupMeta) XXX_Size() int {
	return m.Size()
}
func (m *BackupMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BackupMeta proto.InternalMessageInfo

func (m *BackupMeta) GetEvents() []*BackupEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *BackupMeta) GetStartBackupDependency() uint64 {
	if m != nil {
		return m.StartBackupDependency
	}
	return 0
}

func (m *BackupMeta) GetFullBackupDependency() uint64 {
	if m != nil {
		return m.FullBackupDependency
	}
	return 0
}

func (m *BackupMeta) GetIncBackupDependencies() []uint64 {
	if m != nil {
		return m.IncBackupDependencies
	}
	return nil
}

type BackupRequest struct {
	Context              *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Location             string           `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BackupRequest) Reset()         { *m = BackupRequest{} }
func (m *BackupRequest) String() string { return proto.CompactTextString(m) }
func (*BackupRequest) ProtoMessage()    {}
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_5333a12533fb1467, []int{2}
}
func (m *BackupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupRequest.Merge(dst, src)
}
func (m *BackupRequest) XXX_Size() int {
	return m.Size()
}
func (m *BackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BackupRequest proto.InternalMessageInfo

func (m *BackupRequest) GetContext() *kvrpcpb.Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *BackupRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type BackupResponse struct {
	RegionError          *errorpb.Error `protobuf:"bytes,1,opt,name=region_error,json=regionError" json:"region_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BackupResponse) Reset()         { *m = BackupResponse{} }
func (m *BackupResponse) String() string { return proto.CompactTextString(m) }
func (*BackupResponse) ProtoMessage()    {}
func (*BackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_5333a12533fb1467, []int{3}
}
func (m *BackupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupResponse.Merge(dst, src)
}
func (m *BackupResponse) XXX_Size() int {
	return m.Size()
}
func (m *BackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BackupResponse proto.InternalMessageInfo

func (m *BackupResponse) GetRegionError() *errorpb.Error {
	if m != nil {
		return m.RegionError
	}
	return nil
}

func init() {
	proto.RegisterType((*BackupEvent)(nil), "backup.BackupEvent")
	proto.RegisterType((*BackupMeta)(nil), "backup.BackupMeta")
	proto.RegisterType((*BackupRequest)(nil), "backup.BackupRequest")
	proto.RegisterType((*BackupResponse)(nil), "backup.BackupResponse")
	proto.RegisterEnum("backup.BackupEvent_Event", BackupEvent_Event_name, BackupEvent_Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Backup service

type BackupClient interface {
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error)
}

type backupClient struct {
	cc *grpc.ClientConn
}

func NewBackupClient(cc *grpc.ClientConn) BackupClient {
	return &backupClient{cc}
}

func (c *backupClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResponse, error) {
	out := new(BackupResponse)
	err := c.cc.Invoke(ctx, "/backup.Backup/backup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Backup service

type BackupServer interface {
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)
}

func RegisterBackupServer(s *grpc.Server, srv BackupServer) {
	s.RegisterService(&_Backup_serviceDesc, srv)
}

func _Backup_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Backup/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.Backup",
	HandlerType: (*BackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "backup",
			Handler:    _Backup_Backup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup.proto",
}

func (m *BackupEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RegionId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.RegionId))
	}
	if len(m.RelatedRegionIds) > 0 {
		dAtA2 := make([]byte, len(m.RelatedRegionIds)*10)
		var j1 int
		for _, num := range m.RelatedRegionIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintBackup(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Event != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Event))
	}
	if m.Dependency != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Dependency))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBackup(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.StartBackupDependency != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.StartBackupDependency))
	}
	if m.FullBackupDependency != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.FullBackupDependency))
	}
	if len(m.IncBackupDependencies) > 0 {
		dAtA4 := make([]byte, len(m.IncBackupDependencies)*10)
		var j3 int
		for _, num := range m.IncBackupDependencies {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintBackup(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.Context.Size()))
		n5, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.Location) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBackup(dAtA, i, uint64(len(m.Location)))
		i += copy(dAtA[i:], m.Location)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RegionError != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBackup(dAtA, i, uint64(m.RegionError.Size()))
		n6, err := m.RegionError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintBackup(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BackupEvent) Size() (n int) {
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovBackup(uint64(m.RegionId))
	}
	if len(m.RelatedRegionIds) > 0 {
		l = 0
		for _, e := range m.RelatedRegionIds {
			l += sovBackup(uint64(e))
		}
		n += 1 + sovBackup(uint64(l)) + l
	}
	if m.Event != 0 {
		n += 1 + sovBackup(uint64(m.Event))
	}
	if m.Dependency != 0 {
		n += 1 + sovBackup(uint64(m.Dependency))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupMeta) Size() (n int) {
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovBackup(uint64(l))
		}
	}
	if m.StartBackupDependency != 0 {
		n += 1 + sovBackup(uint64(m.StartBackupDependency))
	}
	if m.FullBackupDependency != 0 {
		n += 1 + sovBackup(uint64(m.FullBackupDependency))
	}
	if len(m.IncBackupDependencies) > 0 {
		l = 0
		for _, e := range m.IncBackupDependencies {
			l += sovBackup(uint64(e))
		}
		n += 1 + sovBackup(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupRequest) Size() (n int) {
	var l int
	_ = l
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovBackup(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovBackup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackupResponse) Size() (n int) {
	var l int
	_ = l
	if m.RegionError != nil {
		l = m.RegionError.Size()
		n += 1 + l + sovBackup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBackup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBackup(x uint64) (n int) {
	return sovBackup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RelatedRegionIds = append(m.RelatedRegionIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBackup
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBackup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RelatedRegionIds = append(m.RelatedRegionIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RelatedRegionIds", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			m.Event = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Event |= (BackupEvent_Event(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependency", wireType)
			}
			m.Dependency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dependency |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &BackupEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBackupDependency", wireType)
			}
			m.StartBackupDependency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBackupDependency |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullBackupDependency", wireType)
			}
			m.FullBackupDependency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FullBackupDependency |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.IncBackupDependencies = append(m.IncBackupDependencies, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBackup
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBackup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.IncBackupDependencies = append(m.IncBackupDependencies, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field IncBackupDependencies", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &kvrpcpb.Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBackup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionError == nil {
				m.RegionError = &errorpb.Error{}
			}
			if err := m.RegionError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBackup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBackup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBackup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBackup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBackup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBackup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBackup
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBackup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBackup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBackup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("backup.proto", fileDescriptor_backup_5333a12533fb1467) }

var fileDescriptor_backup_5333a12533fb1467 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0xe4, 0xd5, 0xe4, 0x3a, 0x49, 0xcd, 0xd0, 0xb4, 0x26, 0x48, 0x51, 0xe4, 0x05, 0x8a,
	0x00, 0xb9, 0x22, 0xa0, 0x4a, 0x6c, 0x1b, 0xba, 0x60, 0x51, 0x09, 0x4d, 0x85, 0x58, 0x46, 0x8e,
	0x73, 0x09, 0x23, 0xbb, 0x33, 0x66, 0x66, 0x12, 0xe0, 0x27, 0x58, 0xf3, 0x09, 0x7c, 0x0a, 0x4b,
	0x96, 0xec, 0x40, 0xe1, 0x47, 0x90, 0x67, 0x9c, 0x28, 0xb4, 0xdd, 0x78, 0xe6, 0x9e, 0x7b, 0xce,
	0x7d, 0x78, 0x0e, 0x74, 0xe6, 0x71, 0x92, 0xae, 0xf2, 0x28, 0x57, 0xd2, 0x48, 0xda, 0x74, 0xd1,
	0xa0, 0x9b, 0xae, 0x55, 0x9e, 0xe4, 0x73, 0x07, 0x0f, 0xba, 0xa8, 0x94, 0x54, 0xbb, 0xf0, 0x68,
	0x29, 0x97, 0xd2, 0x5e, 0x4f, 0x8b, 0x5b, 0x89, 0x1e, 0xaa, 0x95, 0x36, 0xf6, 0xea, 0x80, 0xf0,
	0x6b, 0x15, 0xbc, 0x73, 0x5b, 0xef, 0x62, 0x8d, 0xc2, 0xd0, 0x87, 0xd0, 0x56, 0xb8, 0xe4, 0x52,
	0xcc, 0xf8, 0x22, 0x20, 0x23, 0x32, 0xae, 0xb3, 0x96, 0x03, 0x5e, 0x2f, 0xe8, 0x53, 0xa0, 0x0a,
	0xb3, 0xd8, 0xe0, 0x62, 0xb6, 0x23, 0xe9, 0xa0, 0x3a, 0xaa, 0x8d, 0xeb, 0xcc, 0x2f, 0x33, 0xac,
	0x24, 0x6b, 0x7a, 0x0a, 0x0d, 0x2c, 0x6a, 0x06, 0xb5, 0x11, 0x19, 0xf7, 0x26, 0x0f, 0xa2, 0x72,
	0x8b, 0xbd, 0x76, 0x91, 0xfd, 0x32, 0xc7, 0xa3, 0x43, 0x80, 0x05, 0xe6, 0x28, 0x16, 0x28, 0x92,
	0x2f, 0x41, 0xdd, 0x36, 0xdf, 0x43, 0xc2, 0x04, 0x1a, 0x6e, 0x48, 0x0f, 0x0e, 0xde, 0x8a, 0x54,
	0xc8, 0x4f, 0xc2, 0xaf, 0xd0, 0x0e, 0xb4, 0xae, 0x44, 0x9c, 0xeb, 0x0f, 0xd2, 0xf8, 0x84, 0xb6,
	0xa1, 0x71, 0x95, 0x67, 0xdc, 0xf8, 0x55, 0xea, 0x43, 0xe7, 0x8d, 0xc2, 0x3c, 0x56, 0x78, 0x89,
	0x6a, 0x89, 0x7e, 0x8d, 0x1e, 0x82, 0x37, 0x95, 0xd7, 0xd7, 0xdc, 0x38, 0xa0, 0x4e, 0xef, 0x41,
	0x97, 0xc9, 0x2c, 0x2b, 0x06, 0x73, 0x50, 0x23, 0xfc, 0x4d, 0x00, 0xdc, 0x84, 0x97, 0x68, 0x62,
	0xfa, 0x04, 0x9a, 0x76, 0x38, 0x1d, 0x90, 0x51, 0x6d, 0xec, 0x4d, 0xee, 0xdf, 0xb1, 0x05, 0x2b,
	0x29, 0xf4, 0x0c, 0x4e, 0xb4, 0x89, 0x95, 0x99, 0x39, 0xce, 0x6c, 0x6f, 0x9b, 0xaa, 0xdd, 0xa6,
	0x6f, 0xd3, 0x4e, 0xfa, 0x6a, 0x97, 0xa4, 0x2f, 0xe0, 0xf8, 0xfd, 0x2a, 0xcb, 0xee, 0x90, 0xd5,
	0xac, 0xec, 0xa8, 0xc8, 0xde, 0x52, 0x9d, 0xc1, 0x09, 0x17, 0xc9, 0x2d, 0x11, 0x47, 0x1d, 0xd4,
	0xed, 0x93, 0xf4, 0xb9, 0x48, 0x6e, 0xa8, 0x38, 0xea, 0xf0, 0x1d, 0x74, 0x1d, 0xca, 0xf0, 0xe3,
	0x0a, 0xb5, 0xa1, 0x8f, 0xe1, 0x20, 0x91, 0xc2, 0xe0, 0x67, 0x63, 0x5f, 0xdc, 0x9b, 0xf8, 0xd1,
	0xd6, 0x5a, 0x53, 0x87, 0xb3, 0x2d, 0x81, 0x0e, 0xa0, 0x95, 0xc9, 0x24, 0x36, 0x5c, 0x0a, 0xbb,
	0x53, 0x9b, 0xed, 0xe2, 0x70, 0x0a, 0xbd, 0x6d, 0x61, 0x9d, 0x4b, 0xa1, 0x91, 0x3e, 0x83, 0x4e,
	0x69, 0x14, 0x6b, 0xce, 0xb2, 0x7c, 0x2f, 0xda, 0x5a, 0xf5, 0xa2, 0x38, 0x99, 0xe7, 0x38, 0x36,
	0x98, 0x4c, 0xa1, 0xe9, 0x8a, 0xd0, 0x97, 0x50, 0x3a, 0x9d, 0xf6, 0xff, 0xff, 0xe9, 0xe5, 0xdc,
	0x83, 0xe3, 0x9b, 0xb0, 0xeb, 0x1a, 0x56, 0xce, 0x1f, 0xfd, 0xfa, 0xde, 0x22, 0x3f, 0x36, 0x43,
	0xf2, 0x73, 0x33, 0x24, 0x7f, 0x36, 0x43, 0xf2, 0xed, 0xef, 0xb0, 0x02, 0xbe, 0x54, 0xcb, 0xc8,
	0xf0, 0x74, 0x1d, 0xa5, 0x6b, 0xeb, 0xfe, 0x79, 0xd3, 0x1e, 0xcf, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0xc4, 0x12, 0xbc, 0x61, 0x03, 0x00, 0x00,
}
