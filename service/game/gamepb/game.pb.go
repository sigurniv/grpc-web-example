// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/game/gamepb/game.proto

package gamepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
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

type GameSearch struct {
	Search               string   `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameSearch) Reset()         { *m = GameSearch{} }
func (m *GameSearch) String() string { return proto.CompactTextString(m) }
func (*GameSearch) ProtoMessage()    {}
func (*GameSearch) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{0}
}
func (m *GameSearch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameSearch.Unmarshal(m, b)
}
func (m *GameSearch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameSearch.Marshal(b, m, deterministic)
}
func (dst *GameSearch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameSearch.Merge(dst, src)
}
func (m *GameSearch) XXX_Size() int {
	return xxx_messageInfo_GameSearch.Size(m)
}
func (m *GameSearch) XXX_DiscardUnknown() {
	xxx_messageInfo_GameSearch.DiscardUnknown(m)
}

var xxx_messageInfo_GameSearch proto.InternalMessageInfo

func (m *GameSearch) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type SearchGameRequest struct {
	Search               *GameSearch `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SearchGameRequest) Reset()         { *m = SearchGameRequest{} }
func (m *SearchGameRequest) String() string { return proto.CompactTextString(m) }
func (*SearchGameRequest) ProtoMessage()    {}
func (*SearchGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{1}
}
func (m *SearchGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGameRequest.Unmarshal(m, b)
}
func (m *SearchGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGameRequest.Marshal(b, m, deterministic)
}
func (dst *SearchGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGameRequest.Merge(dst, src)
}
func (m *SearchGameRequest) XXX_Size() int {
	return xxx_messageInfo_SearchGameRequest.Size(m)
}
func (m *SearchGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGameRequest proto.InternalMessageInfo

func (m *SearchGameRequest) GetSearch() *GameSearch {
	if m != nil {
		return m.Search
	}
	return nil
}

type Game struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{2}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (dst *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(dst, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Game) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type SearchGameResponse struct {
	Games                []*Game  `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchGameResponse) Reset()         { *m = SearchGameResponse{} }
func (m *SearchGameResponse) String() string { return proto.CompactTextString(m) }
func (*SearchGameResponse) ProtoMessage()    {}
func (*SearchGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{3}
}
func (m *SearchGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGameResponse.Unmarshal(m, b)
}
func (m *SearchGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGameResponse.Marshal(b, m, deterministic)
}
func (dst *SearchGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGameResponse.Merge(dst, src)
}
func (m *SearchGameResponse) XXX_Size() int {
	return xxx_messageInfo_SearchGameResponse.Size(m)
}
func (m *SearchGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGameResponse proto.InternalMessageInfo

func (m *SearchGameResponse) GetGames() []*Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type GameDetailsRequest struct {
	Game                 *Game    `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameDetailsRequest) Reset()         { *m = GameDetailsRequest{} }
func (m *GameDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*GameDetailsRequest) ProtoMessage()    {}
func (*GameDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{4}
}
func (m *GameDetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameDetailsRequest.Unmarshal(m, b)
}
func (m *GameDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameDetailsRequest.Marshal(b, m, deterministic)
}
func (dst *GameDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameDetailsRequest.Merge(dst, src)
}
func (m *GameDetailsRequest) XXX_Size() int {
	return xxx_messageInfo_GameDetailsRequest.Size(m)
}
func (m *GameDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GameDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GameDetailsRequest proto.InternalMessageInfo

func (m *GameDetailsRequest) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type GameDetails struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameDetails) Reset()         { *m = GameDetails{} }
func (m *GameDetails) String() string { return proto.CompactTextString(m) }
func (*GameDetails) ProtoMessage()    {}
func (*GameDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{5}
}
func (m *GameDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameDetails.Unmarshal(m, b)
}
func (m *GameDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameDetails.Marshal(b, m, deterministic)
}
func (dst *GameDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameDetails.Merge(dst, src)
}
func (m *GameDetails) XXX_Size() int {
	return xxx_messageInfo_GameDetails.Size(m)
}
func (m *GameDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_GameDetails.DiscardUnknown(m)
}

var xxx_messageInfo_GameDetails proto.InternalMessageInfo

func (m *GameDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameDetails) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GameDetails) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *GameDetails) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type GameDetailsResponse struct {
	Details              *GameDetails `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GameDetailsResponse) Reset()         { *m = GameDetailsResponse{} }
func (m *GameDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*GameDetailsResponse) ProtoMessage()    {}
func (*GameDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_9cd1880f528bc2bc, []int{6}
}
func (m *GameDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameDetailsResponse.Unmarshal(m, b)
}
func (m *GameDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameDetailsResponse.Marshal(b, m, deterministic)
}
func (dst *GameDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameDetailsResponse.Merge(dst, src)
}
func (m *GameDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_GameDetailsResponse.Size(m)
}
func (m *GameDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GameDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GameDetailsResponse proto.InternalMessageInfo

func (m *GameDetailsResponse) GetDetails() *GameDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*GameSearch)(nil), "game.GameSearch")
	proto.RegisterType((*SearchGameRequest)(nil), "game.SearchGameRequest")
	proto.RegisterType((*Game)(nil), "game.Game")
	proto.RegisterType((*SearchGameResponse)(nil), "game.SearchGameResponse")
	proto.RegisterType((*GameDetailsRequest)(nil), "game.GameDetailsRequest")
	proto.RegisterType((*GameDetails)(nil), "game.GameDetails")
	proto.RegisterType((*GameDetailsResponse)(nil), "game.GameDetailsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	SearchGame(ctx context.Context, in *SearchGameRequest, opts ...grpc.CallOption) (*SearchGameResponse, error)
	GameDetails(ctx context.Context, in *GameDetailsRequest, opts ...grpc.CallOption) (*GameDetailsResponse, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) SearchGame(ctx context.Context, in *SearchGameRequest, opts ...grpc.CallOption) (*SearchGameResponse, error) {
	out := new(SearchGameResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/SearchGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GameDetails(ctx context.Context, in *GameDetailsRequest, opts ...grpc.CallOption) (*GameDetailsResponse, error) {
	out := new(GameDetailsResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/GameDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	SearchGame(context.Context, *SearchGameRequest) (*SearchGameResponse, error)
	GameDetails(context.Context, *GameDetailsRequest) (*GameDetailsResponse, error)
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_SearchGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).SearchGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/SearchGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).SearchGame(ctx, req.(*SearchGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GameDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GameDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/GameDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GameDetails(ctx, req.(*GameDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchGame",
			Handler:    _GameService_SearchGame_Handler,
		},
		{
			MethodName: "GameDetails",
			Handler:    _GameService_GameDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/game/gamepb/game.proto",
}

func init() {
	proto.RegisterFile("service/game/gamepb/game.proto", fileDescriptor_game_9cd1880f528bc2bc)
}

var fileDescriptor_game_9cd1880f528bc2bc = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x65, 0xd3, 0xb4, 0xb6, 0x53, 0x10, 0x3b, 0x8a, 0xae, 0x3d, 0x94, 0x10, 0x3c, 0x04, 0x94,
	0x0a, 0x55, 0xbc, 0x89, 0x50, 0x04, 0xef, 0xe9, 0xcd, 0x5b, 0xba, 0x5d, 0x74, 0xa1, 0x1f, 0x71,
	0x37, 0x15, 0xfc, 0x1b, 0xfe, 0x62, 0xd9, 0x9d, 0xcd, 0x17, 0xf1, 0x92, 0xcc, 0x9b, 0x79, 0xf3,
	0xe6, 0xcd, 0xee, 0xc2, 0xcc, 0x48, 0xfd, 0xad, 0x84, 0xbc, 0xff, 0xc8, 0x76, 0xf4, 0xc9, 0xd7,
	0xee, 0x37, 0xcf, 0xf5, 0xa1, 0x38, 0x60, 0x68, 0xe3, 0xf8, 0x06, 0xe0, 0x2d, 0xdb, 0xc9, 0x95,
	0xcc, 0xb4, 0xf8, 0xc4, 0x4b, 0x18, 0x18, 0x17, 0x71, 0x16, 0xb1, 0x64, 0x94, 0x7a, 0x14, 0x3f,
	0xc3, 0x84, 0x18, 0x96, 0x9b, 0xca, 0xaf, 0xa3, 0x34, 0x05, 0x26, 0x2d, 0xf2, 0x78, 0x71, 0x36,
	0x77, 0xea, 0xb5, 0x5c, 0xd5, 0x7e, 0x07, 0xa1, 0xcd, 0xe2, 0x29, 0x04, 0x6a, 0xe3, 0xa5, 0x03,
	0xb5, 0xc1, 0x0b, 0xe8, 0x17, 0xaa, 0xd8, 0x4a, 0x1e, 0xb8, 0x14, 0x81, 0xf8, 0x09, 0xb0, 0x39,
	0xcc, 0xe4, 0x87, 0xbd, 0x91, 0x18, 0x41, 0xdf, 0xca, 0x1b, 0xce, 0xa2, 0x5e, 0x32, 0x5e, 0x40,
	0x3d, 0x2c, 0xa5, 0x42, 0xfc, 0x08, 0x68, 0xe1, 0xab, 0x2c, 0x32, 0xb5, 0x35, 0xa5, 0xcb, 0x19,
	0xb8, 0x45, 0xbd, 0xc7, 0x66, 0x1b, 0x1d, 0x80, 0x80, 0x71, 0xa3, 0x0b, 0x11, 0xc2, 0x7d, 0x49,
	0x1f, 0xa5, 0x2e, 0xf6, 0xb6, 0x83, 0xca, 0xf6, 0x14, 0x86, 0xe2, 0xa8, 0xb5, 0xdc, 0x8b, 0x1f,
	0xde, 0x73, 0xd9, 0x0a, 0xdb, 0x95, 0x72, 0xad, 0x84, 0xe4, 0x61, 0xc4, 0x12, 0x96, 0x12, 0x88,
	0x97, 0x70, 0xde, 0xb2, 0xe6, 0x77, 0xba, 0x85, 0x93, 0x0d, 0xa5, 0xbc, 0xbd, 0x49, 0x6d, 0xaf,
	0xe4, 0x96, 0x8c, 0xc5, 0x2f, 0x23, 0xa7, 0x2b, 0xba, 0x56, 0x7c, 0x01, 0xa8, 0x8f, 0x09, 0xaf,
	0xa8, 0xb3, 0x73, 0x4b, 0x53, 0xde, 0x2d, 0xf8, 0xe9, 0xcb, 0xf6, 0xe6, 0xbc, 0x3b, 0xdb, 0x4b,
	0x5c, 0xff, 0x53, 0x21, 0x8d, 0xe5, 0xf0, 0x7d, 0x40, 0x2f, 0x6b, 0x3d, 0x70, 0xaf, 0xea, 0xe1,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x77, 0x40, 0xd9, 0x5b, 0x77, 0x02, 0x00, 0x00,
}
