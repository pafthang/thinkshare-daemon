// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: conductor.proto

package packet

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

type WorkerSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp string           `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Display   *DisplaySession  `protobuf:"bytes,3,opt,name=display,proto3,oneof" json:"display,omitempty"`
	Thinkmay  *ThinkmaySession `protobuf:"bytes,4,opt,name=thinkmay,proto3,oneof" json:"thinkmay,omitempty"`
	Sunshine  *SunshineSession `protobuf:"bytes,5,opt,name=sunshine,proto3,oneof" json:"sunshine,omitempty"`
	Turn      *TurnSession     `protobuf:"bytes,6,opt,name=turn,proto3,oneof" json:"turn,omitempty"`
}

func (x *WorkerSession) Reset() {
	*x = WorkerSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerSession) ProtoMessage() {}

func (x *WorkerSession) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerSession.ProtoReflect.Descriptor instead.
func (*WorkerSession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerSession) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WorkerSession) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *WorkerSession) GetDisplay() *DisplaySession {
	if x != nil {
		return x.Display
	}
	return nil
}

func (x *WorkerSession) GetThinkmay() *ThinkmaySession {
	if x != nil {
		return x.Thinkmay
	}
	return nil
}

func (x *WorkerSession) GetSunshine() *SunshineSession {
	if x != nil {
		return x.Sunshine
	}
	return nil
}

func (x *WorkerSession) GetTurn() *TurnSession {
	if x != nil {
		return x.Turn
	}
	return nil
}

type TurnSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinPort  int32  `protobuf:"varint,1,opt,name=minPort,proto3" json:"minPort,omitempty"`  // do sync
	MaxPort  int32  `protobuf:"varint,2,opt,name=maxPort,proto3" json:"maxPort,omitempty"`  // do sync
	Port     int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`        // do sync
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"` // do sync
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"` // do sync
}

func (x *TurnSession) Reset() {
	*x = TurnSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurnSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurnSession) ProtoMessage() {}

func (x *TurnSession) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurnSession.ProtoReflect.Descriptor instead.
func (*TurnSession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{1}
}

func (x *TurnSession) GetMinPort() int32 {
	if x != nil {
		return x.MinPort
	}
	return 0
}

func (x *TurnSession) GetMaxPort() int32 {
	if x != nil {
		return x.MaxPort
	}
	return 0
}

func (x *TurnSession) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *TurnSession) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TurnSession) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DisplaySession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayIndex *int32  `protobuf:"varint,2,opt,name=DisplayIndex,proto3,oneof" json:"DisplayIndex,omitempty"` // do sync
	DisplayName  *string `protobuf:"bytes,3,opt,name=DisplayName,proto3,oneof" json:"DisplayName,omitempty"`    // do sync
	ScreenWidth  int32   `protobuf:"varint,6,opt,name=ScreenWidth,proto3" json:"ScreenWidth,omitempty"`
	ScreenHeight int32   `protobuf:"varint,7,opt,name=ScreenHeight,proto3" json:"ScreenHeight,omitempty"`
}

func (x *DisplaySession) Reset() {
	*x = DisplaySession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplaySession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplaySession) ProtoMessage() {}

func (x *DisplaySession) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplaySession.ProtoReflect.Descriptor instead.
func (*DisplaySession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{2}
}

func (x *DisplaySession) GetDisplayIndex() int32 {
	if x != nil && x.DisplayIndex != nil {
		return *x.DisplayIndex
	}
	return 0
}

func (x *DisplaySession) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *DisplaySession) GetScreenWidth() int32 {
	if x != nil {
		return x.ScreenWidth
	}
	return 0
}

func (x *DisplaySession) GetScreenHeight() int32 {
	if x != nil {
		return x.ScreenHeight
	}
	return 0
}

type ThinkmaySession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthConfig   string `protobuf:"bytes,3,opt,name=authConfig,proto3" json:"authConfig,omitempty"`     // do sync
	WebrtcConfig string `protobuf:"bytes,5,opt,name=webrtcConfig,proto3" json:"webrtcConfig,omitempty"` // do sync
}

func (x *ThinkmaySession) Reset() {
	*x = ThinkmaySession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkmaySession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkmaySession) ProtoMessage() {}

func (x *ThinkmaySession) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkmaySession.ProtoReflect.Descriptor instead.
func (*ThinkmaySession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{3}
}

func (x *ThinkmaySession) GetAuthConfig() string {
	if x != nil {
		return x.AuthConfig
	}
	return ""
}

func (x *ThinkmaySession) GetWebrtcConfig() string {
	if x != nil {
		return x.WebrtcConfig
	}
	return ""
}

type SunshineSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port     string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`         // do sync
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"` // do sync
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"` // do sync
}

func (x *SunshineSession) Reset() {
	*x = SunshineSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SunshineSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SunshineSession) ProtoMessage() {}

func (x *SunshineSession) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SunshineSession.ProtoReflect.Descriptor instead.
func (*SunshineSession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{4}
}

func (x *SunshineSession) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *SunshineSession) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SunshineSession) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type WorkerInfor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname  string   `protobuf:"bytes,2,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	CPU       string   `protobuf:"bytes,3,opt,name=CPU,proto3" json:"CPU,omitempty"`
	RAM       string   `protobuf:"bytes,4,opt,name=RAM,proto3" json:"RAM,omitempty"`
	BIOS      string   `protobuf:"bytes,5,opt,name=BIOS,proto3" json:"BIOS,omitempty"`
	PublicIP  string   `protobuf:"bytes,6,opt,name=PublicIP,proto3" json:"PublicIP,omitempty"`
	PrivateIP string   `protobuf:"bytes,7,opt,name=PrivateIP,proto3" json:"PrivateIP,omitempty"`
	GPUs      []string `protobuf:"bytes,8,rep,name=GPUs,proto3" json:"GPUs,omitempty"`
	TurnPort  *int32   `protobuf:"varint,9,opt,name=TurnPort,proto3,oneof" json:"TurnPort,omitempty"`
}

func (x *WorkerInfor) Reset() {
	*x = WorkerInfor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerInfor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerInfor) ProtoMessage() {}

func (x *WorkerInfor) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerInfor.ProtoReflect.Descriptor instead.
func (*WorkerInfor) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{5}
}

func (x *WorkerInfor) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *WorkerInfor) GetCPU() string {
	if x != nil {
		return x.CPU
	}
	return ""
}

func (x *WorkerInfor) GetRAM() string {
	if x != nil {
		return x.RAM
	}
	return ""
}

func (x *WorkerInfor) GetBIOS() string {
	if x != nil {
		return x.BIOS
	}
	return ""
}

func (x *WorkerInfor) GetPublicIP() string {
	if x != nil {
		return x.PublicIP
	}
	return ""
}

func (x *WorkerInfor) GetPrivateIP() string {
	if x != nil {
		return x.PrivateIP
	}
	return ""
}

func (x *WorkerInfor) GetGPUs() []string {
	if x != nil {
		return x.GPUs
	}
	return nil
}

func (x *WorkerInfor) GetTurnPort() int32 {
	if x != nil && x.TurnPort != nil {
		return *x.TurnPort
	}
	return 0
}

type WorkerLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Log       string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	Level     string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Source    string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *WorkerLog) Reset() {
	*x = WorkerLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerLog) ProtoMessage() {}

func (x *WorkerLog) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerLog.ProtoReflect.Descriptor instead.
func (*WorkerLog) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{6}
}

func (x *WorkerLog) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *WorkerLog) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *WorkerLog) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *WorkerLog) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_conductor_proto protoreflect.FileDescriptor

var file_conductor_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xcd, 0x02, 0x0a, 0x0d,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x37, 0x0a, 0x07, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x6d, 0x61, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x6d, 0x61, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x48, 0x01, 0x52, 0x08, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x6d, 0x61, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x3a, 0x0a, 0x08, 0x73, 0x75, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x75,
	0x6e, 0x73, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52,
	0x08, 0x73, 0x75, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x75, 0x72, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x48, 0x03, 0x52, 0x04, 0x74, 0x75, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x68, 0x69,
	0x6e, 0x6b, 0x6d, 0x61, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x6e, 0x73, 0x68, 0x69,
	0x6e, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x8d, 0x01, 0x0a, 0x0b,
	0x54, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0e,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x0b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x6d, 0x61,
	0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x65, 0x62, 0x72,
	0x74, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5d, 0x0a, 0x0f,
	0x53, 0x75, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x0b,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x50, 0x55, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x41, 0x4d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x52, 0x41, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x42,
	0x49, 0x4f, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x49, 0x4f, 0x53, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x50, 0x55,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x47, 0x50, 0x55, 0x73, 0x12, 0x1f, 0x0a,
	0x08, 0x54, 0x75, 0x72, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x08, 0x54, 0x75, 0x72, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x54, 0x75, 0x72, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x69, 0x0a, 0x09, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conductor_proto_rawDescOnce sync.Once
	file_conductor_proto_rawDescData = file_conductor_proto_rawDesc
)

func file_conductor_proto_rawDescGZIP() []byte {
	file_conductor_proto_rawDescOnce.Do(func() {
		file_conductor_proto_rawDescData = protoimpl.X.CompressGZIP(file_conductor_proto_rawDescData)
	})
	return file_conductor_proto_rawDescData
}

var file_conductor_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_conductor_proto_goTypes = []interface{}{
	(*WorkerSession)(nil),   // 0: protobuf.WorkerSession
	(*TurnSession)(nil),     // 1: protobuf.TurnSession
	(*DisplaySession)(nil),  // 2: protobuf.DisplaySession
	(*ThinkmaySession)(nil), // 3: protobuf.ThinkmaySession
	(*SunshineSession)(nil), // 4: protobuf.SunshineSession
	(*WorkerInfor)(nil),     // 5: protobuf.WorkerInfor
	(*WorkerLog)(nil),       // 6: protobuf.WorkerLog
}
var file_conductor_proto_depIdxs = []int32{
	2, // 0: protobuf.WorkerSession.display:type_name -> protobuf.DisplaySession
	3, // 1: protobuf.WorkerSession.thinkmay:type_name -> protobuf.ThinkmaySession
	4, // 2: protobuf.WorkerSession.sunshine:type_name -> protobuf.SunshineSession
	1, // 3: protobuf.WorkerSession.turn:type_name -> protobuf.TurnSession
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_conductor_proto_init() }
func file_conductor_proto_init() {
	if File_conductor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conductor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerSession); i {
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
		file_conductor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurnSession); i {
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
		file_conductor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplaySession); i {
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
		file_conductor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkmaySession); i {
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
		file_conductor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SunshineSession); i {
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
		file_conductor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerInfor); i {
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
		file_conductor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerLog); i {
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
	file_conductor_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_conductor_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_conductor_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conductor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conductor_proto_goTypes,
		DependencyIndexes: file_conductor_proto_depIdxs,
		MessageInfos:      file_conductor_proto_msgTypes,
	}.Build()
	File_conductor_proto = out.File
	file_conductor_proto_rawDesc = nil
	file_conductor_proto_goTypes = nil
	file_conductor_proto_depIdxs = nil
}
