//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttainableNode) DeepCopyInto(out *AttainableNode) {
	*out = *in
	in.state.DeepCopyInto(&out.state)
	if in.unknownFields != nil {
		in, out := &in.unknownFields, &out.unknownFields
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttainableNode.
func (in *AttainableNode) DeepCopy() *AttainableNode {
	if in == nil {
		return nil
	}
	out := new(AttainableNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WakeNodeUpRequest) DeepCopyInto(out *WakeNodeUpRequest) {
	*out = *in
	in.state.DeepCopyInto(&out.state)
	if in.unknownFields != nil {
		in, out := &in.unknownFields, &out.unknownFields
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WakeNodeUpRequest.
func (in *WakeNodeUpRequest) DeepCopy() *WakeNodeUpRequest {
	if in == nil {
		return nil
	}
	out := new(WakeNodeUpRequest)
	in.DeepCopyInto(out)
	return out
}