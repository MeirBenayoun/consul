// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package meshv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using ProxyStateTemplate within kubernetes types, where deepcopy-gen is used.
func (in *ProxyStateTemplate) DeepCopyInto(out *ProxyStateTemplate) {
	p := proto.Clone(in).(*ProxyStateTemplate)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStateTemplate. Required by controller-gen.
func (in *ProxyStateTemplate) DeepCopy() *ProxyStateTemplate {
	if in == nil {
		return nil
	}
	out := new(ProxyStateTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStateTemplate. Required by controller-gen.
func (in *ProxyStateTemplate) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ProxyState within kubernetes types, where deepcopy-gen is used.
func (in *ProxyState) DeepCopyInto(out *ProxyState) {
	p := proto.Clone(in).(*ProxyState)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyState. Required by controller-gen.
func (in *ProxyState) DeepCopy() *ProxyState {
	if in == nil {
		return nil
	}
	out := new(ProxyState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ProxyState. Required by controller-gen.
func (in *ProxyState) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}