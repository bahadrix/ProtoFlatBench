// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package profile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile_Gender int32

const (
	Profile_Male   Profile_Gender = 0
	Profile_Female Profile_Gender = 1
)

var Profile_Gender_name = map[int32]string{
	0: "Male",
	1: "Female",
}
var Profile_Gender_value = map[string]int32{
	"Male":   0,
	"Female": 1,
}

func (x Profile_Gender) Enum() *Profile_Gender {
	p := new(Profile_Gender)
	*p = x
	return p
}
func (x Profile_Gender) String() string {
	return proto.EnumName(Profile_Gender_name, int32(x))
}
func (x *Profile_Gender) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Profile_Gender_value, data, "Profile_Gender")
	if err != nil {
		return err
	}
	*x = Profile_Gender(value)
	return nil
}
func (Profile_Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Profile struct {
	Id               *int64          `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string         `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Gender           *Profile_Gender `protobuf:"varint,3,req,name=gender,enum=Profile_Gender" json:"gender,omitempty"`
	ProfileImage     []byte          `protobuf:"bytes,4,req,name=profileImage" json:"profileImage,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Profile) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Profile) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Profile) GetGender() Profile_Gender {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return Profile_Male
}

func (m *Profile) GetProfileImage() []byte {
	if m != nil {
		return m.ProfileImage
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterEnum("Profile_Gender", Profile_Gender_name, Profile_Gender_value)
}

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x03, 0xd2, 0x25, 0xf9, 0x4a, 0x55, 0x5c, 0xec, 0x01, 0x10, 0x01, 0x21,
	0x2e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x66, 0x21, 0x1e, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x26, 0x20, 0x8f, 0x53, 0x48, 0x9e, 0x8b, 0x2d, 0x3d, 0x35, 0x2f, 0x25,
	0xb5, 0x48, 0x82, 0x19, 0xc8, 0xe7, 0x33, 0xe2, 0xd7, 0x83, 0xea, 0xd1, 0x73, 0x07, 0x0b, 0x0b,
	0x89, 0x70, 0xf1, 0x40, 0x8d, 0xf5, 0xcc, 0x4d, 0x4c, 0x4f, 0x95, 0x60, 0x01, 0x2a, 0xe3, 0x51,
	0x92, 0xe3, 0x62, 0x83, 0xca, 0x73, 0x70, 0xb1, 0xf8, 0x26, 0xe6, 0xa4, 0x0a, 0x30, 0x00, 0x2d,
	0x61, 0x73, 0x4b, 0xcd, 0x05, 0xb1, 0x19, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x60, 0x35,
	0xb6, 0x8b, 0x00, 0x00, 0x00,
}