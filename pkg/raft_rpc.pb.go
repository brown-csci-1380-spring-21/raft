//
//  Brown University, CS138, Spring 2018
//
//  Purpose: Defines the Raft RPC protocol using Google's Protocol Buffers
//  syntax. See https://developers.google.com/protocol-buffers for more details.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: raft_rpc.proto

package pkg

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A log entry in Raft can be any of any of these four types.
type CommandType int32

const (
	CommandType_NOOP                  CommandType = 0
	CommandType_CLIENT_REGISTRATION   CommandType = 1
	CommandType_STATE_MACHINE_COMMAND CommandType = 2
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "NOOP",
		1: "CLIENT_REGISTRATION",
		2: "STATE_MACHINE_COMMAND",
	}
	CommandType_value = map[string]int32{
		"NOOP":                  0,
		"CLIENT_REGISTRATION":   1,
		"STATE_MACHINE_COMMAND": 2,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_raft_rpc_proto_enumTypes[0].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_raft_rpc_proto_enumTypes[0]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{0}
}

// The possible responses to a client request
type ClientStatus int32

const (
	ClientStatus_OK                   ClientStatus = 0
	ClientStatus_NOT_LEADER           ClientStatus = 1
	ClientStatus_ELECTION_IN_PROGRESS ClientStatus = 2
	ClientStatus_CLUSTER_NOT_STARTED  ClientStatus = 3
	ClientStatus_REQ_FAILED           ClientStatus = 4
)

// Enum value maps for ClientStatus.
var (
	ClientStatus_name = map[int32]string{
		0: "OK",
		1: "NOT_LEADER",
		2: "ELECTION_IN_PROGRESS",
		3: "CLUSTER_NOT_STARTED",
		4: "REQ_FAILED",
	}
	ClientStatus_value = map[string]int32{
		"OK":                   0,
		"NOT_LEADER":           1,
		"ELECTION_IN_PROGRESS": 2,
		"CLUSTER_NOT_STARTED":  3,
		"REQ_FAILED":           4,
	}
)

func (x ClientStatus) Enum() *ClientStatus {
	p := new(ClientStatus)
	*p = x
	return p
}

func (x ClientStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_raft_rpc_proto_enumTypes[1].Descriptor()
}

func (ClientStatus) Type() protoreflect.EnumType {
	return &file_raft_rpc_proto_enumTypes[1]
}

func (x ClientStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientStatus.Descriptor instead.
func (ClientStatus) EnumDescriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{1}
}

type JoinReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*RemoteNode `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *JoinReply) Reset() {
	*x = JoinReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReply) ProtoMessage() {}

func (x *JoinReply) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReply.ProtoReflect.Descriptor instead.
func (*JoinReply) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *JoinReply) GetPeers() []*RemoteNode {
	if x != nil {
		return x.Peers
	}
	return nil
}

// Represents a node in the Raft cluster
type RemoteNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoteNode) Reset() {
	*x = RemoteNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteNode) ProtoMessage() {}

func (x *RemoteNode) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteNode.ProtoReflect.Descriptor instead.
func (*RemoteNode) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteNode) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *RemoteNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of log entry (first index = 1)
	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The term that this entry was in when added
	TermId uint64 `protobuf:"varint,2,opt,name=termId,proto3" json:"termId,omitempty"`
	// Type of command associated with this entry
	Type CommandType `protobuf:"varint,3,opt,name=type,proto3,enum=raft.CommandType" json:"type,omitempty"`
	// Command associated with this log entry in the user's finite-state-machine.
	// Note that we only care about this value when type = STATE_MACHINE_COMMAND
	Command uint64 `protobuf:"varint,4,opt,name=command,proto3" json:"command,omitempty"`
	// Data associated with this log entry in the user's finite-state-machine.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// After processing this log entry, what ID to use when caching the
	// response. Use an empty string to not cache at all
	CacheId string `protobuf:"bytes,6,opt,name=cacheId,proto3" json:"cacheId,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *LogEntry) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LogEntry) GetTermId() uint64 {
	if x != nil {
		return x.TermId
	}
	return 0
}

func (x *LogEntry) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_NOOP
}

func (x *LogEntry) GetCommand() uint64 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *LogEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LogEntry) GetCacheId() string {
	if x != nil {
		return x.CacheId
	}
	return ""
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The leader's term
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// Address of the leader sending this request
	Leader *RemoteNode `protobuf:"bytes,2,opt,name=leader,proto3" json:"leader,omitempty"`
	// The index of the log entry immediately preceding the new ones
	PrevLogIndex uint64 `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	// The term of the log entry at prevLogIndex
	PrevLogTerm uint64 `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	// The log entries the follower needs to store. Empty for heartbeat messages.
	Entries []*LogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	// The leader's commitIndex
	LeaderCommit uint64 `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
}

func (x *AppendEntriesRequest) Reset() {
	*x = AppendEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesRequest) ProtoMessage() {}

func (x *AppendEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesRequest.ProtoReflect.Descriptor instead.
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *AppendEntriesRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesRequest) GetLeader() *RemoteNode {
	if x != nil {
		return x.Leader
	}
	return nil
}

func (x *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesRequest) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntriesRequest) GetLeaderCommit() uint64 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type AppendEntriesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current term, for leader to update itself.
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// True if follower contained entry matching prevLogIndex and prevLogTerm.
	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntriesReply) Reset() {
	*x = AppendEntriesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesReply) ProtoMessage() {}

func (x *AppendEntriesReply) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesReply.ProtoReflect.Descriptor instead.
func (*AppendEntriesReply) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *AppendEntriesReply) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The candidate's current term Id
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// The candidate Id currently requesting a node to vote for it.
	Candidate *RemoteNode `protobuf:"bytes,2,opt,name=candidate,proto3" json:"candidate,omitempty"`
	// The index of the candidate's last log entry
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	// The term of the candidate's last log entry
	LastLogTerm uint64 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *RequestVoteRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetCandidate() *RemoteNode {
	if x != nil {
		return x.Candidate
	}
	return nil
}

func (x *RequestVoteRequest) GetLastLogIndex() uint64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogTerm() uint64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current term, for candidate to update itself
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// True means candidate received vote
	VoteGranted bool `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteReply) Reset() {
	*x = RequestVoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteReply) ProtoMessage() {}

func (x *RequestVoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteReply.ProtoReflect.Descriptor instead.
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *RequestVoteReply) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteReply) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique client ID associated with this client session (received
	// via a previous RegisterClient call).
	ClientId uint64 `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// A sequence number is associated to request to avoid duplicates
	SequenceNum uint64 `protobuf:"varint,2,opt,name=sequenceNum,proto3" json:"sequenceNum,omitempty"`
	// Command to be executed on the state machine; it may affect state
	StateMachineCmd uint64 `protobuf:"varint,4,opt,name=stateMachineCmd,proto3" json:"stateMachineCmd,omitempty"`
	// Data to accompany the command to the state machine; it may affect state
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *ClientRequest) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ClientRequest) GetSequenceNum() uint64 {
	if x != nil {
		return x.SequenceNum
	}
	return 0
}

func (x *ClientRequest) GetStateMachineCmd() uint64 {
	if x != nil {
		return x.StateMachineCmd
	}
	return 0
}

func (x *ClientRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClientReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OK if state machine successfully applied command
	Status ClientStatus `protobuf:"varint,1,opt,name=status,proto3,enum=raft.ClientStatus" json:"status,omitempty"`
	// The unique client ID associated with this client session
	ClientId uint64 `protobuf:"varint,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// State machine output, if successful
	Response []byte `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	// In cases where the client contacted a non-leader, the node should
	// reply with the correct current leader.
	LeaderHint *RemoteNode `protobuf:"bytes,4,opt,name=leaderHint,proto3" json:"leaderHint,omitempty"`
}

func (x *ClientReply) Reset() {
	*x = ClientReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReply) ProtoMessage() {}

func (x *ClientReply) ProtoReflect() protoreflect.Message {
	mi := &file_raft_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReply.ProtoReflect.Descriptor instead.
func (*ClientReply) Descriptor() ([]byte, []int) {
	return file_raft_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *ClientReply) GetStatus() ClientStatus {
	if x != nil {
		return x.Status
	}
	return ClientStatus_OK
}

func (x *ClientReply) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ClientReply) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ClientReply) GetLeaderHint() *RemoteNode {
	if x != nil {
		return x.LeaderHint
	}
	return nil
}

var File_raft_rpc_proto protoreflect.FileDescriptor

var file_raft_rpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x72, 0x61, 0x66, 0x74, 0x22, 0x33, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x0a, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa7, 0x01,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x74, 0x65, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x61, 0x63, 0x68, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x49, 0x64, 0x22, 0xe8, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x28, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67,
	0x54, 0x65, 0x72, 0x6d, 0x12, 0x28, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x12, 0x2e, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x48, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12,
	0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x43, 0x6d, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xa3, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x69, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x48, 0x69, 0x6e, 0x74, 0x2a, 0x4b, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x10, 0x02, 0x2a, 0x69, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f,
	0x54, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a,
	0x0a, 0x52, 0x45, 0x51, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0x95, 0x02,
	0x0a, 0x07, 0x52, 0x61, 0x66, 0x74, 0x52, 0x50, 0x43, 0x12, 0x31, 0x0a, 0x0a, 0x4a, 0x6f, 0x69,
	0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x0f, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x13,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x43, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_rpc_proto_rawDescOnce sync.Once
	file_raft_rpc_proto_rawDescData = file_raft_rpc_proto_rawDesc
)

func file_raft_rpc_proto_rawDescGZIP() []byte {
	file_raft_rpc_proto_rawDescOnce.Do(func() {
		file_raft_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_rpc_proto_rawDescData)
	})
	return file_raft_rpc_proto_rawDescData
}

var file_raft_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_raft_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_raft_rpc_proto_goTypes = []interface{}{
	(CommandType)(0),             // 0: raft.CommandType
	(ClientStatus)(0),            // 1: raft.ClientStatus
	(*JoinReply)(nil),            // 2: raft.JoinReply
	(*RemoteNode)(nil),           // 3: raft.RemoteNode
	(*LogEntry)(nil),             // 4: raft.LogEntry
	(*AppendEntriesRequest)(nil), // 5: raft.AppendEntriesRequest
	(*AppendEntriesReply)(nil),   // 6: raft.AppendEntriesReply
	(*RequestVoteRequest)(nil),   // 7: raft.RequestVoteRequest
	(*RequestVoteReply)(nil),     // 8: raft.RequestVoteReply
	(*ClientRequest)(nil),        // 9: raft.ClientRequest
	(*ClientReply)(nil),          // 10: raft.ClientReply
}
var file_raft_rpc_proto_depIdxs = []int32{
	3,  // 0: raft.JoinReply.peers:type_name -> raft.RemoteNode
	0,  // 1: raft.LogEntry.type:type_name -> raft.CommandType
	3,  // 2: raft.AppendEntriesRequest.leader:type_name -> raft.RemoteNode
	4,  // 3: raft.AppendEntriesRequest.entries:type_name -> raft.LogEntry
	3,  // 4: raft.RequestVoteRequest.candidate:type_name -> raft.RemoteNode
	1,  // 5: raft.ClientReply.status:type_name -> raft.ClientStatus
	3,  // 6: raft.ClientReply.leaderHint:type_name -> raft.RemoteNode
	3,  // 7: raft.RaftRPC.JoinCaller:input_type -> raft.RemoteNode
	5,  // 8: raft.RaftRPC.AppendEntriesCaller:input_type -> raft.AppendEntriesRequest
	7,  // 9: raft.RaftRPC.RequestVoteCaller:input_type -> raft.RequestVoteRequest
	9,  // 10: raft.RaftRPC.ClientRequestCaller:input_type -> raft.ClientRequest
	2,  // 11: raft.RaftRPC.JoinCaller:output_type -> raft.JoinReply
	6,  // 12: raft.RaftRPC.AppendEntriesCaller:output_type -> raft.AppendEntriesReply
	8,  // 13: raft.RaftRPC.RequestVoteCaller:output_type -> raft.RequestVoteReply
	10, // 14: raft.RaftRPC.ClientRequestCaller:output_type -> raft.ClientReply
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_raft_rpc_proto_init() }
func file_raft_rpc_proto_init() {
	if File_raft_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReply); i {
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
		file_raft_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteNode); i {
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
		file_raft_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_raft_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesRequest); i {
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
		file_raft_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesReply); i {
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
		file_raft_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteRequest); i {
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
		file_raft_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteReply); i {
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
		file_raft_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_raft_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReply); i {
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
			RawDescriptor: file_raft_rpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_rpc_proto_goTypes,
		DependencyIndexes: file_raft_rpc_proto_depIdxs,
		EnumInfos:         file_raft_rpc_proto_enumTypes,
		MessageInfos:      file_raft_rpc_proto_msgTypes,
	}.Build()
	File_raft_rpc_proto = out.File
	file_raft_rpc_proto_rawDesc = nil
	file_raft_rpc_proto_goTypes = nil
	file_raft_rpc_proto_depIdxs = nil
}
