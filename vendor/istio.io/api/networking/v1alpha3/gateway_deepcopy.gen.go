// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha3

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using Gateway within kubernetes types, where deepcopy-gen is used.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	p := proto.Clone(in).(*Gateway)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway. Required by controller-gen.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Gateway. Required by controller-gen.
func (in *Gateway) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Server within kubernetes types, where deepcopy-gen is used.
func (in *Server) DeepCopyInto(out *Server) {
	p := proto.Clone(in).(*Server)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server. Required by controller-gen.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Server. Required by controller-gen.
func (in *Server) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Port within kubernetes types, where deepcopy-gen is used.
func (in *Port) DeepCopyInto(out *Port) {
	p := proto.Clone(in).(*Port)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port. Required by controller-gen.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Port. Required by controller-gen.
func (in *Port) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServerTLSSettings within kubernetes types, where deepcopy-gen is used.
func (in *ServerTLSSettings) DeepCopyInto(out *ServerTLSSettings) {
	p := proto.Clone(in).(*ServerTLSSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSSettings. Required by controller-gen.
func (in *ServerTLSSettings) DeepCopy() *ServerTLSSettings {
	if in == nil {
		return nil
	}
	out := new(ServerTLSSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLSSettings. Required by controller-gen.
func (in *ServerTLSSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
