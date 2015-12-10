// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	CreateContainerRequest
	CreateContainerResponse
	SignalRequest
	SignalResponse
	AddProcessRequest
	User
	AddProcessResponse
	CreateCheckpointRequest
	CreateCheckpointResponse
	DeleteCheckpointRequest
	DeleteCheckpointResponse
	ListCheckpointRequest
	Checkpoint
	ListCheckpointResponse
	StateRequest
	ContainerState
	Process
	Container
	Machine
	StateResponse
	UpdateContainerRequest
	UpdateContainerResponse
	EventsRequest
	Event
*/
package types

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

type CreateContainerRequest struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	BundlePath string `protobuf:"bytes,2,opt,name=bundlePath" json:"bundlePath,omitempty"`
	Stdout     string `protobuf:"bytes,4,opt,name=stdout" json:"stdout,omitempty"`
	Stderr     string `protobuf:"bytes,5,opt,name=stderr" json:"stderr,omitempty"`
	Checkpoint string `protobuf:"bytes,7,opt,name=checkpoint" json:"checkpoint,omitempty"`
}

func (m *CreateContainerRequest) Reset()                    { *m = CreateContainerRequest{} }
func (m *CreateContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateContainerRequest) ProtoMessage()               {}
func (*CreateContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateContainerResponse struct {
}

func (m *CreateContainerResponse) Reset()                    { *m = CreateContainerResponse{} }
func (m *CreateContainerResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateContainerResponse) ProtoMessage()               {}
func (*CreateContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SignalRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Pid    uint32 `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	Signal uint32 `protobuf:"varint,3,opt,name=signal" json:"signal,omitempty"`
}

func (m *SignalRequest) Reset()                    { *m = SignalRequest{} }
func (m *SignalRequest) String() string            { return proto.CompactTextString(m) }
func (*SignalRequest) ProtoMessage()               {}
func (*SignalRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SignalResponse struct {
}

func (m *SignalResponse) Reset()                    { *m = SignalResponse{} }
func (m *SignalResponse) String() string            { return proto.CompactTextString(m) }
func (*SignalResponse) ProtoMessage()               {}
func (*SignalResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type AddProcessRequest struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Terminal bool     `protobuf:"varint,2,opt,name=terminal" json:"terminal,omitempty"`
	User     *User    `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Args     []string `protobuf:"bytes,4,rep,name=args" json:"args,omitempty"`
	Env      []string `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	Cwd      string   `protobuf:"bytes,6,opt,name=cwd" json:"cwd,omitempty"`
}

func (m *AddProcessRequest) Reset()                    { *m = AddProcessRequest{} }
func (m *AddProcessRequest) String() string            { return proto.CompactTextString(m) }
func (*AddProcessRequest) ProtoMessage()               {}
func (*AddProcessRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddProcessRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Uid            uint32   `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Gid            uint32   `protobuf:"varint,2,opt,name=gid" json:"gid,omitempty"`
	AdditionalGids []uint32 `protobuf:"varint,3,rep,name=additionalGids" json:"additionalGids,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AddProcessResponse struct {
	Pid uint32 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
}

func (m *AddProcessResponse) Reset()                    { *m = AddProcessResponse{} }
func (m *AddProcessResponse) String() string            { return proto.CompactTextString(m) }
func (*AddProcessResponse) ProtoMessage()               {}
func (*AddProcessResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type CreateCheckpointRequest struct {
	Id         string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Checkpoint *Checkpoint `protobuf:"bytes,2,opt,name=checkpoint" json:"checkpoint,omitempty"`
}

func (m *CreateCheckpointRequest) Reset()                    { *m = CreateCheckpointRequest{} }
func (m *CreateCheckpointRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCheckpointRequest) ProtoMessage()               {}
func (*CreateCheckpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateCheckpointRequest) GetCheckpoint() *Checkpoint {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

type CreateCheckpointResponse struct {
}

func (m *CreateCheckpointResponse) Reset()                    { *m = CreateCheckpointResponse{} }
func (m *CreateCheckpointResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCheckpointResponse) ProtoMessage()               {}
func (*CreateCheckpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DeleteCheckpointRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteCheckpointRequest) Reset()                    { *m = DeleteCheckpointRequest{} }
func (m *DeleteCheckpointRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCheckpointRequest) ProtoMessage()               {}
func (*DeleteCheckpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type DeleteCheckpointResponse struct {
}

func (m *DeleteCheckpointResponse) Reset()                    { *m = DeleteCheckpointResponse{} }
func (m *DeleteCheckpointResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCheckpointResponse) ProtoMessage()               {}
func (*DeleteCheckpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ListCheckpointRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ListCheckpointRequest) Reset()                    { *m = ListCheckpointRequest{} }
func (m *ListCheckpointRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCheckpointRequest) ProtoMessage()               {}
func (*ListCheckpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type Checkpoint struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Exit        bool   `protobuf:"varint,2,opt,name=exit" json:"exit,omitempty"`
	Tcp         bool   `protobuf:"varint,3,opt,name=tcp" json:"tcp,omitempty"`
	UnixSockets bool   `protobuf:"varint,4,opt,name=unixSockets" json:"unixSockets,omitempty"`
	Shell       bool   `protobuf:"varint,5,opt,name=shell" json:"shell,omitempty"`
}

func (m *Checkpoint) Reset()                    { *m = Checkpoint{} }
func (m *Checkpoint) String() string            { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()               {}
func (*Checkpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type ListCheckpointResponse struct {
	Checkpoints []*Checkpoint `protobuf:"bytes,1,rep,name=checkpoints" json:"checkpoints,omitempty"`
}

func (m *ListCheckpointResponse) Reset()                    { *m = ListCheckpointResponse{} }
func (m *ListCheckpointResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCheckpointResponse) ProtoMessage()               {}
func (*ListCheckpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListCheckpointResponse) GetCheckpoints() []*Checkpoint {
	if m != nil {
		return m.Checkpoints
	}
	return nil
}

type StateRequest struct {
}

func (m *StateRequest) Reset()                    { *m = StateRequest{} }
func (m *StateRequest) String() string            { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()               {}
func (*StateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type ContainerState struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ContainerState) Reset()                    { *m = ContainerState{} }
func (m *ContainerState) String() string            { return proto.CompactTextString(m) }
func (*ContainerState) ProtoMessage()               {}
func (*ContainerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type Process struct {
	Pid      uint32   `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Terminal bool     `protobuf:"varint,2,opt,name=terminal" json:"terminal,omitempty"`
	User     *User    `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Args     []string `protobuf:"bytes,4,rep,name=args" json:"args,omitempty"`
	Env      []string `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	Cwd      string   `protobuf:"bytes,6,opt,name=cwd" json:"cwd,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Process) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Container struct {
	Id         string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	BundlePath string     `protobuf:"bytes,3,opt,name=bundlePath" json:"bundlePath,omitempty"`
	Processes  []*Process `protobuf:"bytes,4,rep,name=processes" json:"processes,omitempty"`
	Status     string     `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *Container) GetProcesses() []*Process {
	if m != nil {
		return m.Processes
	}
	return nil
}

// Machine is information about machine on which containerd is run
type Machine struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Cpus   uint32 `protobuf:"varint,2,opt,name=cpus" json:"cpus,omitempty"`
	Memory uint64 `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
}

func (m *Machine) Reset()                    { *m = Machine{} }
func (m *Machine) String() string            { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()               {}
func (*Machine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

// StateResponse is information about containerd daemon
type StateResponse struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
	Machine    *Machine     `protobuf:"bytes,2,opt,name=machine" json:"machine,omitempty"`
}

func (m *StateResponse) Reset()                    { *m = StateResponse{} }
func (m *StateResponse) String() string            { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()               {}
func (*StateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *StateResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *StateResponse) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type UpdateContainerRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Signal uint32 `protobuf:"varint,2,opt,name=signal" json:"signal,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *UpdateContainerRequest) Reset()                    { *m = UpdateContainerRequest{} }
func (m *UpdateContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateContainerRequest) ProtoMessage()               {}
func (*UpdateContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

type UpdateContainerResponse struct {
}

func (m *UpdateContainerResponse) Reset()                    { *m = UpdateContainerResponse{} }
func (m *UpdateContainerResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateContainerResponse) ProtoMessage()               {}
func (*UpdateContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

type EventsRequest struct {
}

func (m *EventsRequest) Reset()                    { *m = EventsRequest{} }
func (m *EventsRequest) String() string            { return proto.CompactTextString(m) }
func (*EventsRequest) ProtoMessage()               {}
func (*EventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

type Event struct {
	Type       string       `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Id         string       `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Status     uint32       `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	BundlePath string       `protobuf:"bytes,4,opt,name=bundlePath" json:"bundlePath,omitempty"`
	Pid        uint32       `protobuf:"varint,5,opt,name=pid" json:"pid,omitempty"`
	Signal     uint32       `protobuf:"varint,7,opt,name=signal" json:"signal,omitempty"`
	Process    *Process     `protobuf:"bytes,8,opt,name=process" json:"process,omitempty"`
	Containers []*Container `protobuf:"bytes,9,rep,name=containers" json:"containers,omitempty"`
	Checkpoint *Checkpoint  `protobuf:"bytes,10,opt,name=checkpoint" json:"checkpoint,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *Event) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *Event) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Event) GetCheckpoint() *Checkpoint {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateContainerRequest)(nil), "types.CreateContainerRequest")
	proto.RegisterType((*CreateContainerResponse)(nil), "types.CreateContainerResponse")
	proto.RegisterType((*SignalRequest)(nil), "types.SignalRequest")
	proto.RegisterType((*SignalResponse)(nil), "types.SignalResponse")
	proto.RegisterType((*AddProcessRequest)(nil), "types.AddProcessRequest")
	proto.RegisterType((*User)(nil), "types.User")
	proto.RegisterType((*AddProcessResponse)(nil), "types.AddProcessResponse")
	proto.RegisterType((*CreateCheckpointRequest)(nil), "types.CreateCheckpointRequest")
	proto.RegisterType((*CreateCheckpointResponse)(nil), "types.CreateCheckpointResponse")
	proto.RegisterType((*DeleteCheckpointRequest)(nil), "types.DeleteCheckpointRequest")
	proto.RegisterType((*DeleteCheckpointResponse)(nil), "types.DeleteCheckpointResponse")
	proto.RegisterType((*ListCheckpointRequest)(nil), "types.ListCheckpointRequest")
	proto.RegisterType((*Checkpoint)(nil), "types.Checkpoint")
	proto.RegisterType((*ListCheckpointResponse)(nil), "types.ListCheckpointResponse")
	proto.RegisterType((*StateRequest)(nil), "types.StateRequest")
	proto.RegisterType((*ContainerState)(nil), "types.ContainerState")
	proto.RegisterType((*Process)(nil), "types.Process")
	proto.RegisterType((*Container)(nil), "types.Container")
	proto.RegisterType((*Machine)(nil), "types.Machine")
	proto.RegisterType((*StateResponse)(nil), "types.StateResponse")
	proto.RegisterType((*UpdateContainerRequest)(nil), "types.UpdateContainerRequest")
	proto.RegisterType((*UpdateContainerResponse)(nil), "types.UpdateContainerResponse")
	proto.RegisterType((*EventsRequest)(nil), "types.EventsRequest")
	proto.RegisterType((*Event)(nil), "types.Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error)
	UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (*UpdateContainerResponse, error)
	Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (*SignalResponse, error)
	AddProcess(ctx context.Context, in *AddProcessRequest, opts ...grpc.CallOption) (*AddProcessResponse, error)
	CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CreateCheckpointResponse, error)
	DeleteCheckpoint(ctx context.Context, in *DeleteCheckpointRequest, opts ...grpc.CallOption) (*DeleteCheckpointResponse, error)
	ListCheckpoint(ctx context.Context, in *ListCheckpointRequest, opts ...grpc.CallOption) (*ListCheckpointResponse, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (API_EventsClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error) {
	out := new(CreateContainerResponse)
	err := grpc.Invoke(ctx, "/types.API/CreateContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (*UpdateContainerResponse, error) {
	out := new(UpdateContainerResponse)
	err := grpc.Invoke(ctx, "/types.API/UpdateContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (*SignalResponse, error) {
	out := new(SignalResponse)
	err := grpc.Invoke(ctx, "/types.API/Signal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) AddProcess(ctx context.Context, in *AddProcessRequest, opts ...grpc.CallOption) (*AddProcessResponse, error) {
	out := new(AddProcessResponse)
	err := grpc.Invoke(ctx, "/types.API/AddProcess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CreateCheckpointResponse, error) {
	out := new(CreateCheckpointResponse)
	err := grpc.Invoke(ctx, "/types.API/CreateCheckpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteCheckpoint(ctx context.Context, in *DeleteCheckpointRequest, opts ...grpc.CallOption) (*DeleteCheckpointResponse, error) {
	out := new(DeleteCheckpointResponse)
	err := grpc.Invoke(ctx, "/types.API/DeleteCheckpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCheckpoint(ctx context.Context, in *ListCheckpointRequest, opts ...grpc.CallOption) (*ListCheckpointResponse, error) {
	out := new(ListCheckpointResponse)
	err := grpc.Invoke(ctx, "/types.API/ListCheckpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := grpc.Invoke(ctx, "/types.API/State", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (API_EventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/types.API/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_EventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type aPIEventsClient struct {
	grpc.ClientStream
}

func (x *aPIEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for API service

type APIServer interface {
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)
	UpdateContainer(context.Context, *UpdateContainerRequest) (*UpdateContainerResponse, error)
	Signal(context.Context, *SignalRequest) (*SignalResponse, error)
	AddProcess(context.Context, *AddProcessRequest) (*AddProcessResponse, error)
	CreateCheckpoint(context.Context, *CreateCheckpointRequest) (*CreateCheckpointResponse, error)
	DeleteCheckpoint(context.Context, *DeleteCheckpointRequest) (*DeleteCheckpointResponse, error)
	ListCheckpoint(context.Context, *ListCheckpointRequest) (*ListCheckpointResponse, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
	Events(*EventsRequest, API_EventsServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateContainer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_UpdateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).UpdateContainer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Signal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SignalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Signal(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_AddProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AddProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).AddProcess(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateCheckpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteCheckpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListCheckpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).State(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Events(m, &aPIEventsServer{stream})
}

type API_EventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type aPIEventsServer struct {
	grpc.ServerStream
}

func (x *aPIEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContainer",
			Handler:    _API_CreateContainer_Handler,
		},
		{
			MethodName: "UpdateContainer",
			Handler:    _API_UpdateContainer_Handler,
		},
		{
			MethodName: "Signal",
			Handler:    _API_Signal_Handler,
		},
		{
			MethodName: "AddProcess",
			Handler:    _API_AddProcess_Handler,
		},
		{
			MethodName: "CreateCheckpoint",
			Handler:    _API_CreateCheckpoint_Handler,
		},
		{
			MethodName: "DeleteCheckpoint",
			Handler:    _API_DeleteCheckpoint_Handler,
		},
		{
			MethodName: "ListCheckpoint",
			Handler:    _API_ListCheckpoint_Handler,
		},
		{
			MethodName: "State",
			Handler:    _API_State_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _API_Events_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0x5d, 0x4f, 0xdc, 0x3a,
	0x10, 0xdd, 0x25, 0x9b, 0xfd, 0x98, 0x65, 0xc3, 0x62, 0x60, 0x09, 0xd1, 0xe5, 0x02, 0xb9, 0xb7,
	0x55, 0x9f, 0x10, 0x82, 0x4a, 0xe5, 0xb1, 0x08, 0xaa, 0xaa, 0x12, 0x55, 0x11, 0x88, 0x4a, 0x7d,
	0x0c, 0x89, 0xc5, 0x46, 0x64, 0x93, 0x34, 0x76, 0x28, 0xfc, 0x8d, 0xfe, 0xb2, 0xfe, 0xa4, 0x3a,
	0x8e, 0xed, 0x7c, 0x6c, 0x52, 0x9e, 0xfa, 0xe8, 0xb1, 0x7d, 0xce, 0xcc, 0x99, 0xf1, 0x49, 0x60,
	0xe4, 0xc4, 0xfe, 0x61, 0x9c, 0x44, 0x34, 0x42, 0x3a, 0x7d, 0x8e, 0x31, 0xb1, 0x03, 0x98, 0x9d,
	0x27, 0xd8, 0xa1, 0xf8, 0x3c, 0x0a, 0xa9, 0xe3, 0x87, 0x38, 0xb9, 0xc6, 0xdf, 0x53, 0x4c, 0x28,
	0x02, 0x58, 0xf1, 0x3d, 0xb3, 0xbb, 0xdf, 0x7d, 0x33, 0x42, 0x6c, 0x71, 0x97, 0x86, 0x5e, 0x80,
	0xaf, 0x1c, 0x3a, 0x37, 0x57, 0x78, 0xcc, 0x80, 0x3e, 0xa1, 0x5e, 0x94, 0x52, 0xb3, 0x57, 0x5a,
	0xe3, 0x24, 0x31, 0x75, 0x79, 0xc7, 0x9d, 0x63, 0xf7, 0x21, 0x8e, 0xfc, 0x90, 0x9a, 0x83, 0x2c,
	0x66, 0xef, 0xc0, 0xf6, 0x12, 0x1b, 0x89, 0xa3, 0x90, 0x60, 0xfb, 0x14, 0x26, 0x37, 0xfe, 0x7d,
	0xe8, 0x04, 0x4d, 0xfc, 0x63, 0xd0, 0x62, 0xb6, 0xc8, 0x88, 0x27, 0x9c, 0x88, 0x9f, 0x34, 0xb5,
	0x6c, 0x6d, 0x4f, 0xc1, 0x90, 0x37, 0x05, 0x16, 0x85, 0xf5, 0x33, 0xcf, 0xbb, 0x4a, 0x22, 0x17,
	0x13, 0xd2, 0x84, 0x37, 0x85, 0x21, 0xc5, 0xc9, 0xc2, 0xcf, 0x40, 0x32, 0xd0, 0x21, 0xda, 0x81,
	0x5e, 0x4a, 0x70, 0xc2, 0x21, 0xc7, 0xc7, 0xe3, 0x43, 0xae, 0xce, 0xe1, 0x2d, 0x0b, 0xa1, 0x55,
	0xe8, 0x39, 0xc9, 0x3d, 0x61, 0x65, 0x6a, 0x79, 0x2a, 0x38, 0x7c, 0x64, 0x35, 0x8a, 0x85, 0xfb,
	0xc3, 0x33, 0xfb, 0xbc, 0xb8, 0x53, 0xe8, 0xf1, 0xf3, 0x2c, 0x98, 0x0a, 0xa6, 0x49, 0xb6, 0xb8,
	0x57, 0x99, 0xcf, 0xc0, 0x70, 0x3c, 0xcf, 0xa7, 0x7e, 0xc4, 0x88, 0x3f, 0xfa, 0x1e, 0x61, 0x74,
	0x1a, 0xab, 0xe0, 0x00, 0x50, 0x39, 0xdf, 0xbc, 0x0a, 0x59, 0x34, 0xc7, 0xb1, 0x2f, 0x95, 0x72,
	0x4a, 0xd3, 0xa6, 0xc2, 0x5e, 0x55, 0x44, 0x5f, 0xe1, 0xc5, 0xac, 0x8b, 0x62, 0x8a, 0x9b, 0xb6,
	0x05, 0xe6, 0x32, 0x9a, 0x10, 0xef, 0x04, 0xb6, 0x2f, 0x70, 0x80, 0x5f, 0x62, 0x62, 0xaa, 0x84,
	0xce, 0x02, 0xe7, 0xc3, 0x90, 0x01, 0x2e, 0x5f, 0x12, 0x80, 0xff, 0xc1, 0xd6, 0xa5, 0x4f, 0xe8,
	0x1f, 0xe1, 0xec, 0x6f, 0x00, 0xc5, 0x01, 0x05, 0xae, 0xa8, 0xf0, 0x93, 0x4f, 0x45, 0xa7, 0x98,
	0x2c, 0xd4, 0x8d, 0x79, 0xa3, 0x86, 0x68, 0x03, 0xc6, 0x69, 0xe8, 0x3f, 0xdd, 0x44, 0xee, 0x03,
	0xa6, 0x84, 0x4f, 0xe2, 0x10, 0x4d, 0x40, 0x27, 0x73, 0x1c, 0x04, 0x7c, 0x10, 0x87, 0xf6, 0x7b,
	0x98, 0xd5, 0xf9, 0x85, 0xc2, 0xaf, 0x61, 0x5c, 0xa8, 0x45, 0x18, 0x9b, 0xd6, 0x2c, 0x97, 0x01,
	0xab, 0x37, 0x94, 0xa9, 0x25, 0x12, 0xb7, 0xf7, 0xc1, 0x50, 0x03, 0xcc, 0x37, 0xf2, 0xe1, 0x77,
	0x68, 0x4a, 0x44, 0x39, 0x0f, 0x30, 0x10, 0xed, 0xac, 0xb4, 0xf1, 0xef, 0x0c, 0x5e, 0x00, 0x23,
	0x95, 0x4e, 0x7b, 0x8f, 0x6a, 0x8f, 0x58, 0xe3, 0xb1, 0x03, 0x18, 0xc5, 0x79, 0x9e, 0x38, 0xe7,
	0x19, 0x1f, 0x1b, 0x22, 0x05, 0x99, 0x7f, 0x51, 0x1a, 0x7f, 0xd7, 0x6c, 0x3e, 0x06, 0x9f, 0x1d,
	0x77, 0xce, 0xc8, 0xea, 0x5c, 0x6e, 0xcc, 0x0e, 0xa9, 0x37, 0xba, 0xc0, 0x8b, 0x28, 0x79, 0xe6,
	0x3c, 0x3d, 0xfb, 0x2b, 0x7b, 0xdd, 0xb9, 0x82, 0x42, 0xfa, 0xff, 0xd9, 0xa0, 0xca, 0x9c, 0xa5,
	0xf2, 0x53, 0xa9, 0xbc, 0x2a, 0x66, 0x0f, 0x06, 0x8b, 0x9c, 0x4b, 0xcc, 0xb2, 0x4c, 0x4e, 0x64,
	0x60, 0x5f, 0xc0, 0xec, 0x36, 0xf6, 0x5e, 0xb2, 0xaf, 0xc2, 0x31, 0x0a, 0x07, 0xc9, 0x4b, 0xd2,
	0xa4, 0x2d, 0x2d, 0xa1, 0x88, 0xe1, 0x5d, 0x83, 0xc9, 0x87, 0x47, 0xcc, 0xa6, 0x43, 0xf6, 0xfe,
	0x57, 0x17, 0x74, 0x1e, 0xc9, 0x2a, 0xce, 0x92, 0x11, 0x1c, 0x39, 0x5f, 0xc9, 0x1a, 0x15, 0xfe,
	0xa4, 0xa6, 0x7c, 0xaf, 0x6c, 0x69, 0x7a, 0xcd, 0xd2, 0x06, 0x7c, 0xcd, 0xea, 0x16, 0x6d, 0x31,
	0x87, 0x95, 0xba, 0x65, 0x53, 0xaa, 0xf2, 0x8d, 0x5a, 0xe4, 0xab, 0xba, 0x01, 0xb4, 0xb8, 0xc1,
	0xf1, 0x4f, 0x1d, 0xb4, 0xb3, 0xab, 0x4f, 0xe8, 0x1a, 0xd6, 0x6a, 0xee, 0x8c, 0x76, 0xe5, 0xe9,
	0xc6, 0x6f, 0x84, 0xf5, 0x6f, 0xdb, 0xb6, 0x50, 0xaf, 0x93, 0x61, 0xd6, 0xa4, 0x55, 0x98, 0xcd,
	0x8d, 0x53, 0x98, 0x6d, 0x1d, 0xe9, 0xa0, 0x77, 0xd0, 0xcf, 0x0d, 0x1f, 0x6d, 0x8a, 0xb3, 0x95,
	0x2f, 0x87, 0xb5, 0x55, 0x8b, 0xaa, 0x8b, 0xe7, 0x00, 0x85, 0xcf, 0x22, 0x53, 0x1c, 0x5b, 0xfa,
	0x54, 0x58, 0x3b, 0x0d, 0x3b, 0x0a, 0xe4, 0x16, 0xa6, 0x75, 0xef, 0x44, 0x35, 0x1d, 0xea, 0x4e,
	0x67, 0xed, 0xb5, 0xee, 0x97, 0x61, 0xeb, 0x0e, 0xaa, 0x60, 0x5b, 0xfc, 0x58, 0xc1, 0xb6, 0x5a,
	0x6f, 0x07, 0x7d, 0x01, 0xa3, 0x6a, 0x7e, 0xe8, 0x1f, 0x71, 0xa9, 0xd1, 0x93, 0xad, 0xdd, 0x96,
	0x5d, 0x05, 0xf8, 0x16, 0xf4, 0xdc, 0xf2, 0x36, 0xa4, 0xca, 0x25, 0x67, 0xb4, 0x36, 0xab, 0x41,
	0x75, 0xeb, 0x08, 0xfa, 0xf9, 0x33, 0x52, 0x2d, 0xab, 0xbc, 0x2a, 0x6b, 0xb5, 0x1c, 0xb5, 0x3b,
	0x47, 0xdd, 0xbb, 0x3e, 0xff, 0x4d, 0x39, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x82, 0xb4,
	0x20, 0xb3, 0x08, 0x00, 0x00,
}
