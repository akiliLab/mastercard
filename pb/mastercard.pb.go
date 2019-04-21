// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/mastercard.proto

package mastercard

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MastercardRequest struct {
	MerchantID string `protobuf:"bytes,1,opt,name=MerchantID,proto3" json:"MerchantID,omitempty"`
	// this can be fuzzy or exact search
	Search               int64    `protobuf:"varint,2,opt,name=Search,proto3" json:"Search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MastercardRequest) Reset()         { *m = MastercardRequest{} }
func (m *MastercardRequest) String() string { return proto.CompactTextString(m) }
func (*MastercardRequest) ProtoMessage()    {}
func (*MastercardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{0}
}

func (m *MastercardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MastercardRequest.Unmarshal(m, b)
}
func (m *MastercardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MastercardRequest.Marshal(b, m, deterministic)
}
func (m *MastercardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MastercardRequest.Merge(m, src)
}
func (m *MastercardRequest) XXX_Size() int {
	return xxx_messageInfo_MastercardRequest.Size(m)
}
func (m *MastercardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MastercardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MastercardRequest proto.InternalMessageInfo

func (m *MastercardRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

func (m *MastercardRequest) GetSearch() int64 {
	if m != nil {
		return m.Search
	}
	return 0
}

type MastercardReply struct {
	MerchantIDs          *MerchantIDs `protobuf:"bytes,1,opt,name=merchantIDs,proto3" json:"merchantIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MastercardReply) Reset()         { *m = MastercardReply{} }
func (m *MastercardReply) String() string { return proto.CompactTextString(m) }
func (*MastercardReply) ProtoMessage()    {}
func (*MastercardReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{1}
}

func (m *MastercardReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MastercardReply.Unmarshal(m, b)
}
func (m *MastercardReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MastercardReply.Marshal(b, m, deterministic)
}
func (m *MastercardReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MastercardReply.Merge(m, src)
}
func (m *MastercardReply) XXX_Size() int {
	return xxx_messageInfo_MastercardReply.Size(m)
}
func (m *MastercardReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MastercardReply.DiscardUnknown(m)
}

var xxx_messageInfo_MastercardReply proto.InternalMessageInfo

func (m *MastercardReply) GetMerchantIDs() *MerchantIDs {
	if m != nil {
		return m.MerchantIDs
	}
	return nil
}

type MerchantIDs struct {
	Message              string      `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Merchant             []*Merchant `protobuf:"bytes,2,rep,name=Merchant,proto3" json:"Merchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MerchantIDs) Reset()         { *m = MerchantIDs{} }
func (m *MerchantIDs) String() string { return proto.CompactTextString(m) }
func (*MerchantIDs) ProtoMessage()    {}
func (*MerchantIDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{2}
}

func (m *MerchantIDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantIDs.Unmarshal(m, b)
}
func (m *MerchantIDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantIDs.Marshal(b, m, deterministic)
}
func (m *MerchantIDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantIDs.Merge(m, src)
}
func (m *MerchantIDs) XXX_Size() int {
	return xxx_messageInfo_MerchantIDs.Size(m)
}
func (m *MerchantIDs) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantIDs.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantIDs proto.InternalMessageInfo

func (m *MerchantIDs) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MerchantIDs) GetMerchant() []*Merchant {
	if m != nil {
		return m.Merchant
	}
	return nil
}

type Merchant struct {
	Address              *Address `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	BrandName            string   `protobuf:"bytes,3,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	MerchantCategory     string   `protobuf:"bytes,4,opt,name=MerchantCategory,proto3" json:"MerchantCategory,omitempty"`
	MerchantDbaName      string   `protobuf:"bytes,5,opt,name=MerchantDbaName,proto3" json:"MerchantDbaName,omitempty"`
	DescriptorText       string   `protobuf:"bytes,6,opt,name=DescriptorText,proto3" json:"DescriptorText,omitempty"`
	LegalCorporateName   string   `protobuf:"bytes,7,opt,name=LegalCorporateName,proto3" json:"LegalCorporateName,omitempty"`
	Comment              string   `protobuf:"bytes,8,opt,name=Comment,proto3" json:"Comment,omitempty"`
	LocationID           int64    `protobuf:"varint,9,opt,name=LocationID,proto3" json:"LocationID,omitempty"`
	SoleProprietorName   string   `protobuf:"bytes,10,opt,name=SoleProprietorName,proto3" json:"SoleProprietorName,omitempty"`
	MatchConfidenceScore int64    `protobuf:"varint,11,opt,name=MatchConfidenceScore,proto3" json:"MatchConfidenceScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Merchant) Reset()         { *m = Merchant{} }
func (m *Merchant) String() string { return proto.CompactTextString(m) }
func (*Merchant) ProtoMessage()    {}
func (*Merchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{3}
}

func (m *Merchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant.Unmarshal(m, b)
}
func (m *Merchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant.Marshal(b, m, deterministic)
}
func (m *Merchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant.Merge(m, src)
}
func (m *Merchant) XXX_Size() int {
	return xxx_messageInfo_Merchant.Size(m)
}
func (m *Merchant) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant proto.InternalMessageInfo

func (m *Merchant) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Merchant) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Merchant) GetBrandName() string {
	if m != nil {
		return m.BrandName
	}
	return ""
}

func (m *Merchant) GetMerchantCategory() string {
	if m != nil {
		return m.MerchantCategory
	}
	return ""
}

func (m *Merchant) GetMerchantDbaName() string {
	if m != nil {
		return m.MerchantDbaName
	}
	return ""
}

func (m *Merchant) GetDescriptorText() string {
	if m != nil {
		return m.DescriptorText
	}
	return ""
}

func (m *Merchant) GetLegalCorporateName() string {
	if m != nil {
		return m.LegalCorporateName
	}
	return ""
}

func (m *Merchant) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Merchant) GetLocationID() int64 {
	if m != nil {
		return m.LocationID
	}
	return 0
}

func (m *Merchant) GetSoleProprietorName() string {
	if m != nil {
		return m.SoleProprietorName
	}
	return ""
}

func (m *Merchant) GetMatchConfidenceScore() int64 {
	if m != nil {
		return m.MatchConfidenceScore
	}
	return 0
}

type CountrySubdivision struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountrySubdivision) Reset()         { *m = CountrySubdivision{} }
func (m *CountrySubdivision) String() string { return proto.CompactTextString(m) }
func (*CountrySubdivision) ProtoMessage()    {}
func (*CountrySubdivision) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{4}
}

func (m *CountrySubdivision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountrySubdivision.Unmarshal(m, b)
}
func (m *CountrySubdivision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountrySubdivision.Marshal(b, m, deterministic)
}
func (m *CountrySubdivision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountrySubdivision.Merge(m, src)
}
func (m *CountrySubdivision) XXX_Size() int {
	return xxx_messageInfo_CountrySubdivision.Size(m)
}
func (m *CountrySubdivision) XXX_DiscardUnknown() {
	xxx_messageInfo_CountrySubdivision.DiscardUnknown(m)
}

var xxx_messageInfo_CountrySubdivision proto.InternalMessageInfo

func (m *CountrySubdivision) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CountrySubdivision) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// Country contains country informations
type Country struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Country) Reset()         { *m = Country{} }
func (m *Country) String() string { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()    {}
func (*Country) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{5}
}

func (m *Country) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Country.Unmarshal(m, b)
}
func (m *Country) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Country.Marshal(b, m, deterministic)
}
func (m *Country) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Country.Merge(m, src)
}
func (m *Country) XXX_Size() int {
	return xxx_messageInfo_Country.Size(m)
}
func (m *Country) XXX_DiscardUnknown() {
	xxx_messageInfo_Country.DiscardUnknown(m)
}

var xxx_messageInfo_Country proto.InternalMessageInfo

func (m *Country) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Country) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// Address contains the full address of a merchant
type Address struct {
	Line1                string              `protobuf:"bytes,1,opt,name=Line1,proto3" json:"Line1,omitempty"`
	Line2                string              `protobuf:"bytes,2,opt,name=Line2,proto3" json:"Line2,omitempty"`
	City                 string              `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	PostalCode           string              `protobuf:"bytes,4,opt,name=PostalCode,proto3" json:"PostalCode,omitempty"`
	CountrySubdivision   *CountrySubdivision `protobuf:"bytes,5,opt,name=CountrySubdivision,proto3" json:"CountrySubdivision,omitempty"`
	Country              *Country            `protobuf:"bytes,6,opt,name=Country,proto3" json:"Country,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78319ba75afee49, []int{6}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetLine1() string {
	if m != nil {
		return m.Line1
	}
	return ""
}

func (m *Address) GetLine2() string {
	if m != nil {
		return m.Line2
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Address) GetCountrySubdivision() *CountrySubdivision {
	if m != nil {
		return m.CountrySubdivision
	}
	return nil
}

func (m *Address) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

func init() {
	proto.RegisterType((*MastercardRequest)(nil), "mastercard.MastercardRequest")
	proto.RegisterType((*MastercardReply)(nil), "mastercard.MastercardReply")
	proto.RegisterType((*MerchantIDs)(nil), "mastercard.MerchantIDs")
	proto.RegisterType((*Merchant)(nil), "mastercard.Merchant")
	proto.RegisterType((*CountrySubdivision)(nil), "mastercard.CountrySubdivision")
	proto.RegisterType((*Country)(nil), "mastercard.Country")
	proto.RegisterType((*Address)(nil), "mastercard.Address")
}

func init() { proto.RegisterFile("pb/mastercard.proto", fileDescriptor_c78319ba75afee49) }

var fileDescriptor_c78319ba75afee49 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xa6, 0xeb, 0xb6, 0xae, 0x17, 0x89, 0xc1, 0xad, 0x1a, 0x11, 0x3f, 0xa6, 0x2a, 0x0f, 0xa8,
	0x42, 0x5a, 0x61, 0xe1, 0x09, 0x89, 0x17, 0x48, 0x25, 0x34, 0xd1, 0x56, 0x55, 0x8a, 0x90, 0x78,
	0x74, 0x93, 0x5b, 0x6b, 0xd1, 0xd8, 0xc1, 0x76, 0x27, 0xfa, 0x07, 0xf2, 0xbf, 0xf0, 0x67, 0xa0,
	0xb8, 0x49, 0x6a, 0xda, 0x22, 0xf1, 0xe6, 0xfb, 0xbe, 0xef, 0xee, 0xec, 0xf3, 0x67, 0xc3, 0x45,
	0x3e, 0x7b, 0x9d, 0x31, 0x6d, 0x48, 0x25, 0x4c, 0xa5, 0xfd, 0x5c, 0x49, 0x23, 0x11, 0xb6, 0x48,
	0xf0, 0x19, 0x1e, 0x8f, 0xea, 0x28, 0xa6, 0x1f, 0x2b, 0xd2, 0x06, 0xaf, 0x00, 0x46, 0xa4, 0x92,
	0x05, 0x13, 0xe6, 0x76, 0xe0, 0x37, 0xba, 0x8d, 0x5e, 0x3b, 0x76, 0x10, 0xbc, 0x84, 0xd3, 0x29,
	0x31, 0x95, 0x2c, 0xfc, 0xa3, 0x6e, 0xa3, 0xd7, 0x8c, 0xcb, 0x28, 0x18, 0xc2, 0xb9, 0x5b, 0x2c,
	0x5f, 0xae, 0xf1, 0x1d, 0x78, 0x59, 0x9d, 0xa8, 0x6d, 0x2d, 0x2f, 0x7c, 0xd2, 0x77, 0xf6, 0xb4,
	0xad, 0xab, 0x63, 0x57, 0x1b, 0x7c, 0x03, 0xcf, 0xe1, 0xd0, 0x87, 0xd6, 0x88, 0xb4, 0x66, 0x73,
	0x2a, 0x77, 0x54, 0x85, 0xf8, 0x06, 0xce, 0x2a, 0xa1, 0x7f, 0xd4, 0x6d, 0xf6, 0xbc, 0xb0, 0x73,
	0xa8, 0x41, 0x5c, 0xab, 0x82, 0x5f, 0xcd, 0x6d, 0x0a, 0x5e, 0x43, 0xeb, 0x43, 0x9a, 0x2a, 0xd2,
	0xd5, 0xf6, 0x2e, 0xdc, 0xec, 0x92, 0x8a, 0x2b, 0x0d, 0x76, 0xc1, 0x9b, 0x2c, 0xa4, 0xa0, 0xf1,
	0x2a, 0x9b, 0x91, 0xb2, 0x13, 0x68, 0xc7, 0x2e, 0x84, 0xcf, 0xa1, 0xfd, 0x51, 0x31, 0x91, 0x8e,
	0x59, 0x46, 0x7e, 0xd3, 0xf2, 0x5b, 0x00, 0x5f, 0xc1, 0xa3, 0xaa, 0x75, 0xc4, 0x0c, 0xcd, 0xa5,
	0x5a, 0xfb, 0xc7, 0x56, 0xb4, 0x87, 0x63, 0x0f, 0xce, 0x2b, 0x6c, 0x30, 0x63, 0xb6, 0xde, 0x89,
	0x95, 0xee, 0xc2, 0xf8, 0x12, 0x1e, 0x0e, 0x48, 0x27, 0x8a, 0xe7, 0x46, 0xaa, 0x2f, 0xf4, 0xd3,
	0xf8, 0xa7, 0x56, 0xb8, 0x83, 0x62, 0x1f, 0x70, 0x48, 0x73, 0xb6, 0x8c, 0xa4, 0xca, 0xa5, 0x62,
	0x86, 0x6c, 0xd1, 0x96, 0xd5, 0x1e, 0x60, 0x8a, 0xa9, 0x47, 0x32, 0xcb, 0x48, 0x18, 0xff, 0x6c,
	0x33, 0xf5, 0x32, 0x2c, 0x4c, 0x32, 0x94, 0x09, 0x33, 0x5c, 0x8a, 0xdb, 0x81, 0xdf, 0xb6, 0x46,
	0x70, 0x90, 0xa2, 0xd3, 0x54, 0x2e, 0x69, 0xa2, 0x64, 0xae, 0x38, 0x19, 0xa9, 0x6c, 0x27, 0xd8,
	0x74, 0xda, 0x67, 0x30, 0x84, 0xce, 0x88, 0x99, 0x64, 0x11, 0x49, 0x71, 0xc7, 0x53, 0x12, 0x09,
	0x4d, 0x13, 0xa9, 0xc8, 0xf7, 0x6c, 0xe5, 0x83, 0x5c, 0xf0, 0x1e, 0x30, 0x92, 0x2b, 0x61, 0xd4,
	0x7a, 0xba, 0x9a, 0xa5, 0xfc, 0x9e, 0x6b, 0x2e, 0x05, 0x22, 0x1c, 0xdb, 0x5e, 0x1b, 0x9b, 0xd8,
	0x75, 0x81, 0x45, 0x32, 0xa5, 0xf2, 0xba, 0xec, 0x3a, 0xb8, 0x29, 0xce, 0x66, 0xb3, 0xff, 0x3b,
	0xe5, 0x77, 0xa3, 0x36, 0x0b, 0x76, 0xe0, 0x64, 0xc8, 0x05, 0xdd, 0x94, 0x49, 0x9b, 0xa0, 0x42,
	0xc3, 0x32, 0x6d, 0x13, 0xd8, 0x5a, 0xdc, 0xac, 0x4b, 0x37, 0xd8, 0x75, 0x31, 0xc0, 0x89, 0xd4,
	0xa6, 0x98, 0x78, 0x4a, 0xa5, 0x05, 0x1c, 0x04, 0xc7, 0x87, 0x0e, 0x67, 0xef, 0xdf, 0x0b, 0xaf,
	0x5c, 0x8b, 0xee, 0xab, 0xe2, 0x43, 0x63, 0xb9, 0xae, 0x8f, 0x6b, 0xbd, 0xb1, 0xe3, 0xf3, 0x92,
	0x8a, 0x2b, 0x4d, 0xf8, 0xdd, 0xfd, 0x19, 0xa6, 0xa4, 0xee, 0x79, 0x42, 0xf8, 0x15, 0x2e, 0x3f,
	0x91, 0xa9, 0x9f, 0x65, 0x4a, 0xc2, 0xf0, 0x3b, 0x4e, 0x4a, 0xe3, 0x8b, 0xbf, 0x9e, 0xdc, 0xee,
	0x97, 0xf2, 0xf4, 0xd9, 0xbf, 0xe8, 0x7c, 0xb9, 0x0e, 0x1e, 0xcc, 0x4e, 0xed, 0xcf, 0xf4, 0xf6,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x06, 0xab, 0x70, 0xb0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MastercardServiceClient is the client API for MastercardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MastercardServiceClient interface {
	GetMerchantIdentifiers(ctx context.Context, in *MastercardRequest, opts ...grpc.CallOption) (*MastercardReply, error)
}

type mastercardServiceClient struct {
	cc *grpc.ClientConn
}

func NewMastercardServiceClient(cc *grpc.ClientConn) MastercardServiceClient {
	return &mastercardServiceClient{cc}
}

func (c *mastercardServiceClient) GetMerchantIdentifiers(ctx context.Context, in *MastercardRequest, opts ...grpc.CallOption) (*MastercardReply, error) {
	out := new(MastercardReply)
	err := c.cc.Invoke(ctx, "/mastercard.MastercardService/GetMerchantIdentifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MastercardServiceServer is the server API for MastercardService service.
type MastercardServiceServer interface {
	GetMerchantIdentifiers(context.Context, *MastercardRequest) (*MastercardReply, error)
}

func RegisterMastercardServiceServer(s *grpc.Server, srv MastercardServiceServer) {
	s.RegisterService(&_MastercardService_serviceDesc, srv)
}

func _MastercardService_GetMerchantIdentifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MastercardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MastercardServiceServer).GetMerchantIdentifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mastercard.MastercardService/GetMerchantIdentifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MastercardServiceServer).GetMerchantIdentifiers(ctx, req.(*MastercardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MastercardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mastercard.MastercardService",
	HandlerType: (*MastercardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMerchantIdentifiers",
			Handler:    _MastercardService_GetMerchantIdentifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/mastercard.proto",
}