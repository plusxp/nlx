// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: external.proto

package external

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	api "go.nlx.io/nlx/management-api/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RequestAccessRequest struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAccessRequest) Reset()         { *m = RequestAccessRequest{} }
func (m *RequestAccessRequest) String() string { return proto.CompactTextString(m) }
func (*RequestAccessRequest) ProtoMessage()    {}
func (*RequestAccessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b7268f56e161ef5, []int{0}
}
func (m *RequestAccessRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestAccessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestAccessRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestAccessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAccessRequest.Merge(m, src)
}
func (m *RequestAccessRequest) XXX_Size() int {
	return m.Size()
}
func (m *RequestAccessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAccessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAccessRequest proto.InternalMessageInfo

type GetAccessRequestStateRequest struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccessRequestStateRequest) Reset()         { *m = GetAccessRequestStateRequest{} }
func (m *GetAccessRequestStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccessRequestStateRequest) ProtoMessage()    {}
func (*GetAccessRequestStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b7268f56e161ef5, []int{1}
}
func (m *GetAccessRequestStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAccessRequestStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAccessRequestStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAccessRequestStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccessRequestStateRequest.Merge(m, src)
}
func (m *GetAccessRequestStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAccessRequestStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccessRequestStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccessRequestStateRequest proto.InternalMessageInfo

type GetAccessRequestStateResponse struct {
	State                api.AccessRequestState `protobuf:"varint,1,opt,name=state,proto3,enum=nlx.management.AccessRequestState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetAccessRequestStateResponse) Reset()         { *m = GetAccessRequestStateResponse{} }
func (m *GetAccessRequestStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccessRequestStateResponse) ProtoMessage()    {}
func (*GetAccessRequestStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b7268f56e161ef5, []int{2}
}
func (m *GetAccessRequestStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAccessRequestStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAccessRequestStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAccessRequestStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccessRequestStateResponse.Merge(m, src)
}
func (m *GetAccessRequestStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAccessRequestStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccessRequestStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccessRequestStateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestAccessRequest)(nil), "nlx.management.external.RequestAccessRequest")
	proto.RegisterType((*GetAccessRequestStateRequest)(nil), "nlx.management.external.GetAccessRequestStateRequest")
	proto.RegisterType((*GetAccessRequestStateResponse)(nil), "nlx.management.external.GetAccessRequestStateResponse")
}

func init() { proto.RegisterFile("external.proto", fileDescriptor_2b7268f56e161ef5) }

var fileDescriptor_2b7268f56e161ef5 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x28, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcf, 0xcb, 0xa9, 0xd0,
	0xcb, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x83, 0x49, 0x4b, 0x49, 0xa7,
	0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x95, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94,
	0x54, 0x42, 0x74, 0x49, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0xe7, 0xa7, 0xe7, 0x23, 0x54, 0x81, 0x78, 0x60, 0x0e, 0x98, 0x05, 0x55, 0x2e, 0x80, 0x64,
	0x01, 0x58, 0x44, 0xc9, 0x92, 0x4b, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xc4, 0x31, 0x39,
	0x39, 0xb5, 0xb8, 0x18, 0xca, 0x11, 0x52, 0xe4, 0xe2, 0x29, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x8d, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x86, 0x8a, 0xf9,
	0x25, 0xe6, 0xa6, 0x2a, 0x39, 0x72, 0xc9, 0xb8, 0xa7, 0xa2, 0x6a, 0x0b, 0x2e, 0x49, 0x2c, 0x49,
	0x25, 0xc1, 0x88, 0x48, 0x2e, 0x59, 0x1c, 0x46, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x59,
	0x70, 0xb1, 0x16, 0x83, 0x04, 0xc0, 0x9a, 0xf9, 0x8c, 0x94, 0xf4, 0xd0, 0x42, 0x09, 0x8b, 0x56,
	0x88, 0x06, 0xa3, 0xdf, 0x8c, 0x5c, 0x22, 0xa8, 0xb2, 0x10, 0x7b, 0x85, 0x22, 0xb8, 0x78, 0x51,
	0x7c, 0x2c, 0xa4, 0xab, 0x87, 0x23, 0xe8, 0xf5, 0xb0, 0x85, 0x8c, 0x94, 0x98, 0x1e, 0x24, 0x42,
	0xf4, 0x60, 0x41, 0xad, 0xe7, 0x0a, 0x8a, 0x10, 0x25, 0x06, 0xa1, 0x0e, 0x46, 0x2e, 0x51, 0xac,
	0xde, 0x11, 0x32, 0xc5, 0x69, 0x05, 0xbe, 0x10, 0x94, 0x32, 0x23, 0x55, 0x1b, 0x24, 0xd4, 0x94,
	0x18, 0x9c, 0xa4, 0x4e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0xa3, 0x38, 0x60, 0x7a, 0x93, 0xd8, 0xc0, 0x0e, 0x37, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x70, 0xb6, 0xba, 0x4b, 0x82, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessRequestServiceClient is the client API for AccessRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessRequestServiceClient interface {
	RequestAccess(ctx context.Context, in *RequestAccessRequest, opts ...grpc.CallOption) (*types.Empty, error)
	GetAccessRequestState(ctx context.Context, in *GetAccessRequestStateRequest, opts ...grpc.CallOption) (*GetAccessRequestStateResponse, error)
}

type accessRequestServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessRequestServiceClient(cc *grpc.ClientConn) AccessRequestServiceClient {
	return &accessRequestServiceClient{cc}
}

func (c *accessRequestServiceClient) RequestAccess(ctx context.Context, in *RequestAccessRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/nlx.management.external.AccessRequestService/RequestAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRequestServiceClient) GetAccessRequestState(ctx context.Context, in *GetAccessRequestStateRequest, opts ...grpc.CallOption) (*GetAccessRequestStateResponse, error) {
	out := new(GetAccessRequestStateResponse)
	err := c.cc.Invoke(ctx, "/nlx.management.external.AccessRequestService/GetAccessRequestState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessRequestServiceServer is the server API for AccessRequestService service.
type AccessRequestServiceServer interface {
	RequestAccess(context.Context, *RequestAccessRequest) (*types.Empty, error)
	GetAccessRequestState(context.Context, *GetAccessRequestStateRequest) (*GetAccessRequestStateResponse, error)
}

// UnimplementedAccessRequestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessRequestServiceServer struct {
}

func (*UnimplementedAccessRequestServiceServer) RequestAccess(ctx context.Context, req *RequestAccessRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccess not implemented")
}
func (*UnimplementedAccessRequestServiceServer) GetAccessRequestState(ctx context.Context, req *GetAccessRequestStateRequest) (*GetAccessRequestStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessRequestState not implemented")
}

func RegisterAccessRequestServiceServer(s *grpc.Server, srv AccessRequestServiceServer) {
	s.RegisterService(&_AccessRequestService_serviceDesc, srv)
}

func _AccessRequestService_RequestAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRequestServiceServer).RequestAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.AccessRequestService/RequestAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRequestServiceServer).RequestAccess(ctx, req.(*RequestAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRequestService_GetAccessRequestState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessRequestStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRequestServiceServer).GetAccessRequestState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.AccessRequestService/GetAccessRequestState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRequestServiceServer).GetAccessRequestState(ctx, req.(*GetAccessRequestStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessRequestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nlx.management.external.AccessRequestService",
	HandlerType: (*AccessRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAccess",
			Handler:    _AccessRequestService_RequestAccess_Handler,
		},
		{
			MethodName: "GetAccessRequestState",
			Handler:    _AccessRequestService_GetAccessRequestState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external.proto",
}

func (m *RequestAccessRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestAccessRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestAccessRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintExternal(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAccessRequestStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAccessRequestStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAccessRequestStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintExternal(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAccessRequestStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAccessRequestStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAccessRequestStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.State != 0 {
		i = encodeVarintExternal(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExternal(dAtA []byte, offset int, v uint64) int {
	offset -= sovExternal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestAccessRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovExternal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAccessRequestStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovExternal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetAccessRequestStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovExternal(uint64(m.State))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExternal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExternal(x uint64) (n int) {
	return sovExternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestAccessRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestAccessRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestAccessRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternal
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
func (m *GetAccessRequestStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAccessRequestStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAccessRequestStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternal
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
func (m *GetAccessRequestStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetAccessRequestStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAccessRequestStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= api.AccessRequestState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternal
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
func skipExternal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternal
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
					return 0, ErrIntOverflowExternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExternal
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
			if length < 0 {
				return 0, ErrInvalidLengthExternal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExternal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExternal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExternal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExternal = fmt.Errorf("proto: unexpected end of group")
)