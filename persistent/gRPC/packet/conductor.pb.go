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

type Closer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Reason    string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Closer) Reset() {
	*x = Closer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Closer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Closer) ProtoMessage() {}

func (x *Closer) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Closer.ProtoReflect.Descriptor instead.
func (*Closer) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{0}
}

func (x *Closer) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Closer) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type AppSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exe    string   `protobuf:"bytes,3,opt,name=exe,proto3" json:"exe,omitempty"`
	Folder string   `protobuf:"bytes,4,opt,name=folder,proto3" json:"folder,omitempty"`
	Args   []string `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
	Envs   []string `protobuf:"bytes,6,rep,name=envs,proto3" json:"envs,omitempty"`
}

func (x *AppSession) Reset() {
	*x = AppSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppSession) ProtoMessage() {}

func (x *AppSession) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AppSession.ProtoReflect.Descriptor instead.
func (*AppSession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{1}
}

func (x *AppSession) GetExe() string {
	if x != nil {
		return x.Exe
	}
	return ""
}

func (x *AppSession) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

func (x *AppSession) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *AppSession) GetEnvs() []string {
	if x != nil {
		return x.Envs
	}
	return nil
}

type WorkerSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp       string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	AuthConfig      string `protobuf:"bytes,3,opt,name=authConfig,proto3" json:"authConfig,omitempty"`           // do sync
	SignalingConfig string `protobuf:"bytes,4,opt,name=signalingConfig,proto3" json:"signalingConfig,omitempty"` // do sync
	WebrtcConfig    string `protobuf:"bytes,5,opt,name=webrtcConfig,proto3" json:"webrtcConfig,omitempty"`       // do sync
	Manifest        string `protobuf:"bytes,8,opt,name=manifest,proto3" json:"manifest,omitempty"`               // not sync
}

func (x *WorkerSession) Reset() {
	*x = WorkerSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerSession) ProtoMessage() {}

func (x *WorkerSession) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use WorkerSession.ProtoReflect.Descriptor instead.
func (*WorkerSession) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{2}
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

func (x *WorkerSession) GetAuthConfig() string {
	if x != nil {
		return x.AuthConfig
	}
	return ""
}

func (x *WorkerSession) GetSignalingConfig() string {
	if x != nil {
		return x.SignalingConfig
	}
	return ""
}

func (x *WorkerSession) GetWebrtcConfig() string {
	if x != nil {
		return x.WebrtcConfig
	}
	return ""
}

func (x *WorkerSession) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

type WorkerSessions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sessions []*WorkerSession `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	Apps     []*AppSession    `protobuf:"bytes,2,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *WorkerSessions) Reset() {
	*x = WorkerSessions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerSessions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerSessions) ProtoMessage() {}

func (x *WorkerSessions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use WorkerSessions.ProtoReflect.Descriptor instead.
func (*WorkerSessions) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{3}
}

func (x *WorkerSessions) GetSessions() []*WorkerSession {
	if x != nil {
		return x.Sessions
	}
	return nil
}

func (x *WorkerSessions) GetApps() []*AppSession {
	if x != nil {
		return x.Apps
	}
	return nil
}

type Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device     string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Mountpoint string `protobuf:"bytes,2,opt,name=mountpoint,proto3" json:"mountpoint,omitempty"`
	Fstype     string `protobuf:"bytes,3,opt,name=fstype,proto3" json:"fstype,omitempty"`
	Opts       string `protobuf:"bytes,4,opt,name=opts,proto3" json:"opts,omitempty"`
}

func (x *Partition) Reset() {
	*x = Partition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{4}
}

func (x *Partition) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Partition) GetMountpoint() string {
	if x != nil {
		return x.Mountpoint
	}
	return ""
}

func (x *Partition) GetFstype() string {
	if x != nil {
		return x.Fstype
	}
	return ""
}

func (x *Partition) GetOpts() string {
	if x != nil {
		return x.Opts
	}
	return ""
}

type WorkerInfor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp  string       `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Hostname   string       `protobuf:"bytes,2,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	CPU        string       `protobuf:"bytes,3,opt,name=CPU,proto3" json:"CPU,omitempty"`
	RAM        string       `protobuf:"bytes,4,opt,name=RAM,proto3" json:"RAM,omitempty"`
	BIOS       string       `protobuf:"bytes,5,opt,name=BIOS,proto3" json:"BIOS,omitempty"`
	PublicIP   string       `protobuf:"bytes,6,opt,name=PublicIP,proto3" json:"PublicIP,omitempty"`
	PrivateIP  string       `protobuf:"bytes,7,opt,name=PrivateIP,proto3" json:"PrivateIP,omitempty"`
	GPUs       []string     `protobuf:"bytes,8,rep,name=GPUs,proto3" json:"GPUs,omitempty"`
	Disks      []string     `protobuf:"bytes,9,rep,name=disks,proto3" json:"disks,omitempty"`
	Partitions []*Partition `protobuf:"bytes,10,rep,name=partitions,proto3" json:"partitions,omitempty"`
	NICs       []string     `protobuf:"bytes,11,rep,name=NICs,proto3" json:"NICs,omitempty"`
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

func (x *WorkerInfor) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
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

func (x *WorkerInfor) GetDisks() []string {
	if x != nil {
		return x.Disks
	}
	return nil
}

func (x *WorkerInfor) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

func (x *WorkerInfor) GetNICs() []string {
	if x != nil {
		return x.NICs
	}
	return nil
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

type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineHash   string `protobuf:"bytes,1,opt,name=PipelineHash,proto3" json:"PipelineHash,omitempty"`
	PipelineString string `protobuf:"bytes,2,opt,name=PipelineString,proto3" json:"PipelineString,omitempty"`
	Plugin         string `protobuf:"bytes,3,opt,name=Plugin,proto3" json:"Plugin,omitempty"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{7}
}

func (x *Pipeline) GetPipelineHash() string {
	if x != nil {
		return x.PipelineHash
	}
	return ""
}

func (x *Pipeline) GetPipelineString() string {
	if x != nil {
		return x.PipelineString
	}
	return ""
}

func (x *Pipeline) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

type Soundcard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID string    `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Api      string    `protobuf:"bytes,3,opt,name=Api,proto3" json:"Api,omitempty"`
	Pipeline *Pipeline `protobuf:"bytes,6,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *Soundcard) Reset() {
	*x = Soundcard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Soundcard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Soundcard) ProtoMessage() {}

func (x *Soundcard) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Soundcard.ProtoReflect.Descriptor instead.
func (*Soundcard) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{8}
}

func (x *Soundcard) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *Soundcard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Soundcard) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *Soundcard) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

type Microphone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID string    `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Api      string    `protobuf:"bytes,3,opt,name=Api,proto3" json:"Api,omitempty"`
	Pipeline *Pipeline `protobuf:"bytes,6,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *Microphone) Reset() {
	*x = Microphone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Microphone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Microphone) ProtoMessage() {}

func (x *Microphone) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Microphone.ProtoReflect.Descriptor instead.
func (*Microphone) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{9}
}

func (x *Microphone) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *Microphone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Microphone) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *Microphone) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

type MediaDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Soundcard  *Soundcard  `protobuf:"bytes,1,opt,name=soundcard,proto3" json:"soundcard,omitempty"`
	Microphone *Microphone `protobuf:"bytes,2,opt,name=microphone,proto3" json:"microphone,omitempty"`
}

func (x *MediaDevice) Reset() {
	*x = MediaDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conductor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaDevice) ProtoMessage() {}

func (x *MediaDevice) ProtoReflect() protoreflect.Message {
	mi := &file_conductor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaDevice.ProtoReflect.Descriptor instead.
func (*MediaDevice) Descriptor() ([]byte, []int) {
	return file_conductor_proto_rawDescGZIP(), []int{10}
}

func (x *MediaDevice) GetSoundcard() *Soundcard {
	if x != nil {
		return x.Soundcard
	}
	return nil
}

func (x *MediaDevice) GetMicrophone() *Microphone {
	if x != nil {
		return x.Microphone
	}
	return nil
}

var File_conductor_proto protoreflect.FileDescriptor

var file_conductor_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x3e, 0x0a, 0x06, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x0a, 0x41,
	0x70, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x0d,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x62,
	0x72, 0x74, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x04,
	0x61, 0x70, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x6f, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x73, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x43, 0x50, 0x55, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x41, 0x4d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x52, 0x41, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x49, 0x4f, 0x53, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x49, 0x4f, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x49, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x49, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x50, 0x55, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x47, 0x50, 0x55, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x33, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x49, 0x43, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x49, 0x43, 0x73, 0x22, 0x69, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x6e, 0x0a, 0x08, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x22, 0x7d, 0x0a, 0x09, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x41, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x70, 0x69,
	0x12, 0x2e, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0x7e, 0x0a, 0x0a, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x41, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x70, 0x69,
	0x12, 0x2e, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0x76, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x31, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x63, 0x61, 0x72, 0x64, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x63, 0x61,
	0x72, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x0a, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0xb2, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x28, 0x01, 0x12, 0x31, 0x0a, 0x06, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x28, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_conductor_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_conductor_proto_goTypes = []interface{}{
	(*Closer)(nil),         // 0: protobuf.Closer
	(*AppSession)(nil),     // 1: protobuf.AppSession
	(*WorkerSession)(nil),  // 2: protobuf.WorkerSession
	(*WorkerSessions)(nil), // 3: protobuf.WorkerSessions
	(*Partition)(nil),      // 4: protobuf.Partition
	(*WorkerInfor)(nil),    // 5: protobuf.WorkerInfor
	(*WorkerLog)(nil),      // 6: protobuf.WorkerLog
	(*Pipeline)(nil),       // 7: protobuf.Pipeline
	(*Soundcard)(nil),      // 8: protobuf.Soundcard
	(*Microphone)(nil),     // 9: protobuf.Microphone
	(*MediaDevice)(nil),    // 10: protobuf.MediaDevice
}
var file_conductor_proto_depIdxs = []int32{
	2,  // 0: protobuf.WorkerSessions.sessions:type_name -> protobuf.WorkerSession
	1,  // 1: protobuf.WorkerSessions.apps:type_name -> protobuf.AppSession
	4,  // 2: protobuf.WorkerInfor.partitions:type_name -> protobuf.Partition
	7,  // 3: protobuf.Soundcard.pipeline:type_name -> protobuf.Pipeline
	7,  // 4: protobuf.Microphone.pipeline:type_name -> protobuf.Pipeline
	8,  // 5: protobuf.MediaDevice.soundcard:type_name -> protobuf.Soundcard
	9,  // 6: protobuf.MediaDevice.microphone:type_name -> protobuf.Microphone
	3,  // 7: protobuf.Conductor.sync:input_type -> protobuf.WorkerSessions
	5,  // 8: protobuf.Conductor.infor:input_type -> protobuf.WorkerInfor
	6,  // 9: protobuf.Conductor.logger:input_type -> protobuf.WorkerLog
	3,  // 10: protobuf.Conductor.sync:output_type -> protobuf.WorkerSessions
	0,  // 11: protobuf.Conductor.infor:output_type -> protobuf.Closer
	0,  // 12: protobuf.Conductor.logger:output_type -> protobuf.Closer
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_conductor_proto_init() }
func file_conductor_proto_init() {
	if File_conductor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conductor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Closer); i {
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
			switch v := v.(*AppSession); i {
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
		file_conductor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerSessions); i {
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
			switch v := v.(*Partition); i {
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
		file_conductor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_conductor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Soundcard); i {
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
		file_conductor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Microphone); i {
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
		file_conductor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaDevice); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conductor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
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
