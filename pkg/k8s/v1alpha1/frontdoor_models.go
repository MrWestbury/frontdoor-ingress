package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type FrontdoorSpec struct {
	FrontdoorId string `json:"frontdoorId"`
}

type Frontdoor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec FrontdoorSpec `json:"spec"`
}

type FrontdoorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Frontdoor `json:"items"`
}

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Frontdoor) DeepCopyInto(out *Frontdoor) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = FrontdoorSpec{
		FrontdoorId: in.Spec.FrontdoorId,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Frontdoor) DeepCopyObject() runtime.Object {
	out := Frontdoor{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *FrontdoorList) DeepCopyInto(out *FrontdoorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)

	if in.Items != nil {
		out.Items = make([]Frontdoor, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

func (in *FrontdoorList) DeepCopy() *FrontdoorList {
	if in == nil {
		return nil
	}
	out := new(FrontdoorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrontdoorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
