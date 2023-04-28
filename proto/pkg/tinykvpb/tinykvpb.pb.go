// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tinykvpb.proto

package tinykvpb

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	kvrpcpb "github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
	raft_serverpb "github.com/pingcap-incubator/tinykv/proto/pkg/raft_serverpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("tinykvpb.proto", fileDescriptor_69e7b50ae4863887) }

var fileDescriptor_69e7b50ae4863887 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x5b, 0x09, 0x42, 0xb1, 0x34, 0x18, 0xee, 0x80, 0x2d, 0x8c, 0x20, 0xed, 0xc4, 0xa9,
	0x48, 0x80, 0xc4, 0x01, 0x81, 0xc4, 0x5a, 0x69, 0x07, 0x0f, 0xa9, 0x72, 0xc7, 0x19, 0xb9, 0xd6,
	0xb7, 0x36, 0x4a, 0x66, 0x07, 0xdb, 0x71, 0xd8, 0x9b, 0xf0, 0x48, 0x1c, 0x79, 0x04, 0x54, 0x5e,
	0x81, 0x07, 0x40, 0x49, 0xb0, 0x9b, 0xa4, 0xe9, 0x2d, 0xfe, 0xfd, 0xbf, 0xff, 0xcf, 0xaa, 0x6b,
	0xa3, 0x07, 0x26, 0x16, 0xb7, 0x89, 0xcd, 0x96, 0x93, 0x4c, 0x49, 0x23, 0xf1, 0xc8, 0xad, 0xc3,
	0x83, 0xc4, 0xaa, 0x8c, 0xbb, 0x20, 0x1c, 0x2b, 0x76, 0x6d, 0xbe, 0x6a, 0x50, 0x16, 0x94, 0x87,
	0x8f, 0xb8, 0xcc, 0x94, 0xe4, 0xa0, 0xb5, 0x54, 0xff, 0xd1, 0xd1, 0x4a, 0xae, 0x64, 0xf5, 0xf9,
	0xaa, 0xfc, 0xaa, 0xe9, 0xeb, 0xbf, 0x01, 0x0a, 0xae, 0x62, 0x71, 0x4b, 0x2c, 0x7e, 0x8b, 0xee,
	0x12, 0x7b, 0x01, 0x06, 0x8f, 0x27, 0x6e, 0x87, 0x0b, 0x30, 0x14, 0xbe, 0xe5, 0xa0, 0x4d, 0x78,
	0xd4, 0x86, 0x3a, 0x93, 0x42, 0xc3, 0xd9, 0x00, 0xbf, 0x43, 0x01, 0xb1, 0x0b, 0xce, 0x04, 0xde,
	0x4e, 0x94, 0x4b, 0xd7, 0x7b, 0xdc, 0xa1, 0xbe, 0x38, 0x45, 0x88, 0xd8, 0xb9, 0x82, 0x42, 0xc5,
	0x06, 0xf0, 0xb1, 0x1f, 0x73, 0xc8, 0x09, 0x4e, 0x7a, 0x12, 0x2f, 0xf9, 0x80, 0x46, 0xc4, 0x4e,
	0xe5, 0xcd, 0x4d, 0x6c, 0xf0, 0x13, 0x3f, 0x58, 0x03, 0x27, 0x78, 0xba, 0xc3, 0x7d, 0xfd, 0x0b,
	0x3a, 0x24, 0x76, 0xba, 0x06, 0x9e, 0x5c, 0x7d, 0x17, 0x0b, 0xc3, 0x4c, 0xae, 0x71, 0xb4, 0x1d,
	0x6f, 0x05, 0x4e, 0xf7, 0x62, 0x6f, 0xee, 0xb5, 0x14, 0x3d, 0x24, 0xf6, 0x9c, 0x19, 0xbe, 0xa6,
	0x32, 0x4d, 0x97, 0x8c, 0x27, 0xf8, 0xb9, 0x6f, 0xb5, 0xb8, 0x93, 0x46, 0xfb, 0x62, 0xef, 0xbc,
	0x44, 0x07, 0xc4, 0x52, 0xd0, 0x32, 0xb5, 0x70, 0x29, 0x79, 0x82, 0x9f, 0xf9, 0x4a, 0x83, 0x3a,
	0xdf, 0x69, 0x7f, 0xe8, 0x6d, 0xef, 0x51, 0x40, 0x59, 0x51, 0xfe, 0xd9, 0xdb, 0x53, 0xab, 0xc1,
	0xee, 0xa9, 0x39, 0xde, 0x29, 0xcf, 0xf3, 0x4e, 0x79, 0x9e, 0xf7, 0x97, 0x2b, 0xee, 0xcb, 0x33,
	0x74, 0x9f, 0xb2, 0x62, 0x06, 0x29, 0x18, 0xc0, 0x27, 0xcd, 0xb9, 0x9a, 0x39, 0x45, 0xd8, 0x17,
	0x79, 0xcb, 0x47, 0x74, 0x8f, 0xb2, 0xa2, 0xba, 0x76, 0xad, 0xbd, 0x9a, 0x37, 0xef, 0x78, 0x37,
	0x68, 0xfc, 0x84, 0x3b, 0x94, 0x5d, 0x1b, 0x1c, 0x4e, 0xda, 0xaf, 0xa7, 0x84, 0x9f, 0x41, 0x6b,
	0xb6, 0x82, 0x70, 0xdc, 0xc9, 0x66, 0x52, 0xc0, 0xd9, 0xe0, 0xe5, 0x10, 0x7f, 0x42, 0xa3, 0x85,
	0x60, 0x99, 0x5e, 0x4b, 0x83, 0x4f, 0x3b, 0x43, 0x2e, 0x98, 0xae, 0x73, 0x91, 0xec, 0x55, 0x9c,
	0x1f, 0xfe, 0xdc, 0x44, 0xc3, 0x5f, 0x9b, 0x68, 0xf8, 0x7b, 0x13, 0x0d, 0x7f, 0xfc, 0x89, 0x06,
	0xcb, 0xa0, 0x7a, 0x8f, 0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xec, 0x18, 0x44, 0xc9, 0xf8,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TinyKvClient is the client API for TinyKv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TinyKvClient interface {
	// KV commands with mvcc/txn supported.
	KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error)
	KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error)
	KvCheckTxnStatus(ctx context.Context, in *kvrpcpb.CheckTxnStatusRequest, opts ...grpc.CallOption) (*kvrpcpb.CheckTxnStatusResponse, error)
	KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error)
	KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error)
	// RawKV commands.
	RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error)
	RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error)
	RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error)
	// Raft commands (tinykv <-> tinykv).
	Raft(ctx context.Context, opts ...grpc.CallOption) (TinyKv_RaftClient, error)
	Snapshot(ctx context.Context, opts ...grpc.CallOption) (TinyKv_SnapshotClient, error)
}

type tinyKvClient struct {
	cc *grpc.ClientConn
}

func NewTinyKvClient(cc *grpc.ClientConn) TinyKvClient {
	return &tinyKvClient{cc}
}

func (c *tinyKvClient) KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error) {
	out := new(kvrpcpb.GetResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error) {
	out := new(kvrpcpb.ScanResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error) {
	out := new(kvrpcpb.PrewriteResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvPrewrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error) {
	out := new(kvrpcpb.CommitResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvCheckTxnStatus(ctx context.Context, in *kvrpcpb.CheckTxnStatusRequest, opts ...grpc.CallOption) (*kvrpcpb.CheckTxnStatusResponse, error) {
	out := new(kvrpcpb.CheckTxnStatusResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvCheckTxnStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error) {
	out := new(kvrpcpb.BatchRollbackResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvBatchRollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error) {
	out := new(kvrpcpb.ResolveLockResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvResolveLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error) {
	out := new(kvrpcpb.RawGetResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error) {
	out := new(kvrpcpb.RawPutResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error) {
	out := new(kvrpcpb.RawDeleteResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error) {
	out := new(kvrpcpb.RawScanResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) Raft(ctx context.Context, opts ...grpc.CallOption) (TinyKv_RaftClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TinyKv_serviceDesc.Streams[0], "/tinykvpb.TinyKv/Raft", opts...)
	if err != nil {
		return nil, err
	}
	x := &tinyKvRaftClient{stream}
	return x, nil
}

type TinyKv_RaftClient interface {
	Send(*raft_serverpb.RaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tinyKvRaftClient struct {
	grpc.ClientStream
}

func (x *tinyKvRaftClient) Send(m *raft_serverpb.RaftMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tinyKvRaftClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tinyKvClient) Snapshot(ctx context.Context, opts ...grpc.CallOption) (TinyKv_SnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TinyKv_serviceDesc.Streams[1], "/tinykvpb.TinyKv/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &tinyKvSnapshotClient{stream}
	return x, nil
}

type TinyKv_SnapshotClient interface {
	Send(*raft_serverpb.SnapshotChunk) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tinyKvSnapshotClient struct {
	grpc.ClientStream
}

func (x *tinyKvSnapshotClient) Send(m *raft_serverpb.SnapshotChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tinyKvSnapshotClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TinyKvServer is the server API for TinyKv service.
type TinyKvServer interface {
	// KV commands with mvcc/txn supported.
	KvGet(context.Context, *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error)
	KvScan(context.Context, *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(context.Context, *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(context.Context, *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error)
	KvCheckTxnStatus(context.Context, *kvrpcpb.CheckTxnStatusRequest) (*kvrpcpb.CheckTxnStatusResponse, error)
	KvBatchRollback(context.Context, *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error)
	KvResolveLock(context.Context, *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error)
	// RawKV commands.
	RawGet(context.Context, *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error)
	RawPut(context.Context, *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error)
	RawDelete(context.Context, *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(context.Context, *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error)
	// Raft commands (tinykv <-> tinykv).
	Raft(TinyKv_RaftServer) error
	Snapshot(TinyKv_SnapshotServer) error
}

// UnimplementedTinyKvServer can be embedded to have forward compatible implementations.
type UnimplementedTinyKvServer struct {
}

func (*UnimplementedTinyKvServer) KvGet(ctx context.Context, req *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvGet not implemented")
}
func (*UnimplementedTinyKvServer) KvScan(ctx context.Context, req *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvScan not implemented")
}
func (*UnimplementedTinyKvServer) KvPrewrite(ctx context.Context, req *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvPrewrite not implemented")
}
func (*UnimplementedTinyKvServer) KvCommit(ctx context.Context, req *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvCommit not implemented")
}
func (*UnimplementedTinyKvServer) KvCheckTxnStatus(ctx context.Context, req *kvrpcpb.CheckTxnStatusRequest) (*kvrpcpb.CheckTxnStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvCheckTxnStatus not implemented")
}
func (*UnimplementedTinyKvServer) KvBatchRollback(ctx context.Context, req *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvBatchRollback not implemented")
}
func (*UnimplementedTinyKvServer) KvResolveLock(ctx context.Context, req *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvResolveLock not implemented")
}
func (*UnimplementedTinyKvServer) RawGet(ctx context.Context, req *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawGet not implemented")
}
func (*UnimplementedTinyKvServer) RawPut(ctx context.Context, req *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawPut not implemented")
}
func (*UnimplementedTinyKvServer) RawDelete(ctx context.Context, req *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawDelete not implemented")
}
func (*UnimplementedTinyKvServer) RawScan(ctx context.Context, req *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawScan not implemented")
}
func (*UnimplementedTinyKvServer) Raft(srv TinyKv_RaftServer) error {
	return status.Errorf(codes.Unimplemented, "method Raft not implemented")
}
func (*UnimplementedTinyKvServer) Snapshot(srv TinyKv_SnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}

func RegisterTinyKvServer(s *grpc.Server, srv TinyKvServer) {
	s.RegisterService(&_TinyKv_serviceDesc, srv)
}

func _TinyKv_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvGet(ctx, req.(*kvrpcpb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvScan(ctx, req.(*kvrpcpb.ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvPrewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.PrewriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvPrewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvPrewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvPrewrite(ctx, req.(*kvrpcpb.PrewriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvCommit(ctx, req.(*kvrpcpb.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvCheckTxnStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CheckTxnStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvCheckTxnStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvCheckTxnStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvCheckTxnStatus(ctx, req.(*kvrpcpb.CheckTxnStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvBatchRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvBatchRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvBatchRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvBatchRollback(ctx, req.(*kvrpcpb.BatchRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvResolveLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ResolveLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvResolveLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvResolveLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvResolveLock(ctx, req.(*kvrpcpb.ResolveLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawGet(ctx, req.(*kvrpcpb.RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawPut(ctx, req.(*kvrpcpb.RawPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawDelete(ctx, req.(*kvrpcpb.RawDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawScan(ctx, req.(*kvrpcpb.RawScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_Raft_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TinyKvServer).Raft(&tinyKvRaftServer{stream})
}

type TinyKv_RaftServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.RaftMessage, error)
	grpc.ServerStream
}

type tinyKvRaftServer struct {
	grpc.ServerStream
}

func (x *tinyKvRaftServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tinyKvRaftServer) Recv() (*raft_serverpb.RaftMessage, error) {
	m := new(raft_serverpb.RaftMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TinyKv_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TinyKvServer).Snapshot(&tinyKvSnapshotServer{stream})
}

type TinyKv_SnapshotServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.SnapshotChunk, error)
	grpc.ServerStream
}

type tinyKvSnapshotServer struct {
	grpc.ServerStream
}

func (x *tinyKvSnapshotServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tinyKvSnapshotServer) Recv() (*raft_serverpb.SnapshotChunk, error) {
	m := new(raft_serverpb.SnapshotChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TinyKv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tinykvpb.TinyKv",
	HandlerType: (*TinyKvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvGet",
			Handler:    _TinyKv_KvGet_Handler,
		},
		{
			MethodName: "KvScan",
			Handler:    _TinyKv_KvScan_Handler,
		},
		{
			MethodName: "KvPrewrite",
			Handler:    _TinyKv_KvPrewrite_Handler,
		},
		{
			MethodName: "KvCommit",
			Handler:    _TinyKv_KvCommit_Handler,
		},
		{
			MethodName: "KvCheckTxnStatus",
			Handler:    _TinyKv_KvCheckTxnStatus_Handler,
		},
		{
			MethodName: "KvBatchRollback",
			Handler:    _TinyKv_KvBatchRollback_Handler,
		},
		{
			MethodName: "KvResolveLock",
			Handler:    _TinyKv_KvResolveLock_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _TinyKv_RawGet_Handler,
		},
		{
			MethodName: "RawPut",
			Handler:    _TinyKv_RawPut_Handler,
		},
		{
			MethodName: "RawDelete",
			Handler:    _TinyKv_RawDelete_Handler,
		},
		{
			MethodName: "RawScan",
			Handler:    _TinyKv_RawScan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Raft",
			Handler:       _TinyKv_Raft_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Snapshot",
			Handler:       _TinyKv_Snapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tinykvpb.proto",
}
