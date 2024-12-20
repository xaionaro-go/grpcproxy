// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: proxy.proto

package proxy_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NetworkProtocol int32

const (
	NetworkProtocol_networkProtocolUndefined NetworkProtocol = 0
	NetworkProtocol_TCP                      NetworkProtocol = 1
	NetworkProtocol_UDP                      NetworkProtocol = 2
)

// Enum value maps for NetworkProtocol.
var (
	NetworkProtocol_name = map[int32]string{
		0: "networkProtocolUndefined",
		1: "TCP",
		2: "UDP",
	}
	NetworkProtocol_value = map[string]int32{
		"networkProtocolUndefined": 0,
		"TCP":                      1,
		"UDP":                      2,
	}
)

func (x NetworkProtocol) Enum() *NetworkProtocol {
	p := new(NetworkProtocol)
	*p = x
	return p
}

func (x NetworkProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_proto_enumTypes[0].Descriptor()
}

func (NetworkProtocol) Type() protoreflect.EnumType {
	return &file_proxy_proto_enumTypes[0]
}

func (x NetworkProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkProtocol.Descriptor instead.
func (NetworkProtocol) EnumDescriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

type ConnectionStatus int32

const (
	ConnectionStatus_connectionStatusUndefined ConnectionStatus = 0
	ConnectionStatus_Connecting                ConnectionStatus = 1
	ConnectionStatus_Connected                 ConnectionStatus = 2
	ConnectionStatus_Rejected                  ConnectionStatus = 3
	ConnectionStatus_Finishing                 ConnectionStatus = 4
	ConnectionStatus_Finished                  ConnectionStatus = 5
	ConnectionStatus_Reset                     ConnectionStatus = 6
	ConnectionStatus_Unreachable               ConnectionStatus = 7
	ConnectionStatus_TimedOut                  ConnectionStatus = 8
	ConnectionStatus_UnableToResolve           ConnectionStatus = 9
	ConnectionStatus_BrokenProxyConnection     ConnectionStatus = 10
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0:  "connectionStatusUndefined",
		1:  "Connecting",
		2:  "Connected",
		3:  "Rejected",
		4:  "Finishing",
		5:  "Finished",
		6:  "Reset",
		7:  "Unreachable",
		8:  "TimedOut",
		9:  "UnableToResolve",
		10: "BrokenProxyConnection",
	}
	ConnectionStatus_value = map[string]int32{
		"connectionStatusUndefined": 0,
		"Connecting":                1,
		"Connected":                 2,
		"Rejected":                  3,
		"Finishing":                 4,
		"Finished":                  5,
		"Reset":                     6,
		"Unreachable":               7,
		"TimedOut":                  8,
		"UnableToResolve":           9,
		"BrokenProxyConnection":     10,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_proto_enumTypes[1].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_proxy_proto_enumTypes[1]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol NetworkProtocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=proxy.NetworkProtocol" json:"protocol,omitempty"`
	Hostname string          `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port     uint32          `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetProtocol() NetworkProtocol {
	if x != nil {
		return x.Protocol
	}
	return NetworkProtocol_networkProtocolUndefined
}

func (x *ConnectRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ConnectRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ConnectionStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionID     uint64           `protobuf:"varint,1,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
	ConnectionStatus ConnectionStatus `protobuf:"varint,2,opt,name=connectionStatus,proto3,enum=proxy.ConnectionStatus" json:"connectionStatus,omitempty"`
	ErrorMessage     *string          `protobuf:"bytes,3,opt,name=errorMessage,proto3,oneof" json:"errorMessage,omitempty"`
	LocalIPAddress   []byte           `protobuf:"bytes,4,opt,name=localIPAddress,proto3,oneof" json:"localIPAddress,omitempty"`
	LocalPort        *uint32          `protobuf:"varint,5,opt,name=localPort,proto3,oneof" json:"localPort,omitempty"`
	RemoteIPAddress  []byte           `protobuf:"bytes,6,opt,name=remoteIPAddress,proto3,oneof" json:"remoteIPAddress,omitempty"`
	RemotePort       *uint32          `protobuf:"varint,7,opt,name=remotePort,proto3,oneof" json:"remotePort,omitempty"`
}

func (x *ConnectionStatusUpdate) Reset() {
	*x = ConnectionStatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionStatusUpdate) ProtoMessage() {}

func (x *ConnectionStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionStatusUpdate.ProtoReflect.Descriptor instead.
func (*ConnectionStatusUpdate) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionStatusUpdate) GetConnectionID() uint64 {
	if x != nil {
		return x.ConnectionID
	}
	return 0
}

func (x *ConnectionStatusUpdate) GetConnectionStatus() ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return ConnectionStatus_connectionStatusUndefined
}

func (x *ConnectionStatusUpdate) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *ConnectionStatusUpdate) GetLocalIPAddress() []byte {
	if x != nil {
		return x.LocalIPAddress
	}
	return nil
}

func (x *ConnectionStatusUpdate) GetLocalPort() uint32 {
	if x != nil && x.LocalPort != nil {
		return *x.LocalPort
	}
	return 0
}

func (x *ConnectionStatusUpdate) GetRemoteIPAddress() []byte {
	if x != nil {
		return x.RemoteIPAddress
	}
	return nil
}

func (x *ConnectionStatusUpdate) GetRemotePort() uint32 {
	if x != nil && x.RemotePort != nil {
		return *x.RemotePort
	}
	return 0
}

type MessageForward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageForwardOneOf:
	//
	//	*MessageForward_Connect
	//	*MessageForward_SendPayload
	MessageForwardOneOf isMessageForward_MessageForwardOneOf `protobuf_oneof:"MessageForwardOneOf"`
}

func (x *MessageForward) Reset() {
	*x = MessageForward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageForward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageForward) ProtoMessage() {}

func (x *MessageForward) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageForward.ProtoReflect.Descriptor instead.
func (*MessageForward) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{2}
}

func (m *MessageForward) GetMessageForwardOneOf() isMessageForward_MessageForwardOneOf {
	if m != nil {
		return m.MessageForwardOneOf
	}
	return nil
}

func (x *MessageForward) GetConnect() *ConnectRequest {
	if x, ok := x.GetMessageForwardOneOf().(*MessageForward_Connect); ok {
		return x.Connect
	}
	return nil
}

func (x *MessageForward) GetSendPayload() []byte {
	if x, ok := x.GetMessageForwardOneOf().(*MessageForward_SendPayload); ok {
		return x.SendPayload
	}
	return nil
}

type isMessageForward_MessageForwardOneOf interface {
	isMessageForward_MessageForwardOneOf()
}

type MessageForward_Connect struct {
	Connect *ConnectRequest `protobuf:"bytes,1,opt,name=connect,proto3,oneof"`
}

type MessageForward_SendPayload struct {
	SendPayload []byte `protobuf:"bytes,2,opt,name=sendPayload,proto3,oneof"`
}

func (*MessageForward_Connect) isMessageForward_MessageForwardOneOf() {}

func (*MessageForward_SendPayload) isMessageForward_MessageForwardOneOf() {}

type MessageBack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageBackOneOf:
	//
	//	*MessageBack_ConnectionUpdate
	//	*MessageBack_ReceivedPayload
	MessageBackOneOf isMessageBack_MessageBackOneOf `protobuf_oneof:"MessageBackOneOf"`
}

func (x *MessageBack) Reset() {
	*x = MessageBack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBack) ProtoMessage() {}

func (x *MessageBack) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBack.ProtoReflect.Descriptor instead.
func (*MessageBack) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{3}
}

func (m *MessageBack) GetMessageBackOneOf() isMessageBack_MessageBackOneOf {
	if m != nil {
		return m.MessageBackOneOf
	}
	return nil
}

func (x *MessageBack) GetConnectionUpdate() *ConnectionStatusUpdate {
	if x, ok := x.GetMessageBackOneOf().(*MessageBack_ConnectionUpdate); ok {
		return x.ConnectionUpdate
	}
	return nil
}

func (x *MessageBack) GetReceivedPayload() []byte {
	if x, ok := x.GetMessageBackOneOf().(*MessageBack_ReceivedPayload); ok {
		return x.ReceivedPayload
	}
	return nil
}

type isMessageBack_MessageBackOneOf interface {
	isMessageBack_MessageBackOneOf()
}

type MessageBack_ConnectionUpdate struct {
	ConnectionUpdate *ConnectionStatusUpdate `protobuf:"bytes,1,opt,name=connectionUpdate,proto3,oneof"`
}

type MessageBack_ReceivedPayload struct {
	ReceivedPayload []byte `protobuf:"bytes,2,opt,name=receivedPayload,proto3,oneof"`
}

func (*MessageBack_ConnectionUpdate) isMessageBack_MessageBackOneOf() {}

func (*MessageBack_ReceivedPayload) isMessageBack_MessageBackOneOf() {}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x22, 0x74, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xa3, 0x03, 0x0a, 0x16, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x01, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x50, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x0a, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x7e, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65,
	0x6e, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4f, 0x6e, 0x65, 0x4f, 0x66,
	0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x12, 0x4b, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a,
	0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x2a, 0x41, 0x0a,
	0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x1c, 0x0a, 0x18, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x43, 0x50, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x02,
	0x2a, 0xd5, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6e, 0x67, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10, 0x05, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x72,
	0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x69,
	0x6d, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x10, 0x09, 0x12, 0x19, 0x0a,
	0x15, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0a, 0x32, 0x48, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x61, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proxy_proto_goTypes = []interface{}{
	(NetworkProtocol)(0),           // 0: proxy.NetworkProtocol
	(ConnectionStatus)(0),          // 1: proxy.ConnectionStatus
	(*ConnectRequest)(nil),         // 2: proxy.ConnectRequest
	(*ConnectionStatusUpdate)(nil), // 3: proxy.ConnectionStatusUpdate
	(*MessageForward)(nil),         // 4: proxy.MessageForward
	(*MessageBack)(nil),            // 5: proxy.MessageBack
}
var file_proxy_proto_depIdxs = []int32{
	0, // 0: proxy.ConnectRequest.protocol:type_name -> proxy.NetworkProtocol
	1, // 1: proxy.ConnectionStatusUpdate.connectionStatus:type_name -> proxy.ConnectionStatus
	2, // 2: proxy.MessageForward.connect:type_name -> proxy.ConnectRequest
	3, // 3: proxy.MessageBack.connectionUpdate:type_name -> proxy.ConnectionStatusUpdate
	4, // 4: proxy.NetworkProxy.Proxy:input_type -> proxy.MessageForward
	5, // 5: proxy.NetworkProxy.Proxy:output_type -> proxy.MessageBack
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionStatusUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageForward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_proxy_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MessageForward_Connect)(nil),
		(*MessageForward_SendPayload)(nil),
	}
	file_proxy_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MessageBack_ConnectionUpdate)(nil),
		(*MessageBack_ReceivedPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		EnumInfos:         file_proxy_proto_enumTypes,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}
