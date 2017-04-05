// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto
	content.proto
	export/export.proto
	metadata.proto
	page.proto
	route.proto
	theme.proto
	user.proto
	packages.proto
	struct.proto

It has these top-level messages:
	Error
	FieldError
	ListOptions
	UpdatePageRoutesRequest
	GetRenderedPageRequest
	GetRenderedPageResponse
	ListPageRequest
	ListPageResponse
	ListRouteResponse
	ListThemeResponse
	TLSSettingsReponse
	EnableTLSRequest
	InstallThemeRequest
	ContentMultiple
	ContentText
	ContentString
	Export
	Metadata
	Timestamp
	Page
	Content
	Route
	Theme
	ThemeTemplate
	ThemePlaceholder
	ThemeAsset
	User
	PackageRelease
	Package
	Registry
	Struct
	Value
	ListValue
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ketchup_models "github.com/octavore/ketchup/proto/ketchup/models"
import structpb "github.com/octavore/ketchup/proto/structpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_INTERNAL_SERVER_ERROR ErrorCode = 1
	ErrorCode_NOT_FOUND             ErrorCode = 2
	ErrorCode_FORBIDDEN             ErrorCode = 3
)

var ErrorCode_name = map[int32]string{
	1: "INTERNAL_SERVER_ERROR",
	2: "NOT_FOUND",
	3: "FORBIDDEN",
}
var ErrorCode_value = map[string]int32{
	"INTERNAL_SERVER_ERROR": 1,
	"NOT_FOUND":             2,
	"FORBIDDEN":             3,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRenderedPageRequest_RenderedPageFilter int32

const (
	GetRenderedPageRequest_all       GetRenderedPageRequest_RenderedPageFilter = 1
	GetRenderedPageRequest_published GetRenderedPageRequest_RenderedPageFilter = 2
	GetRenderedPageRequest_draft     GetRenderedPageRequest_RenderedPageFilter = 3
)

var GetRenderedPageRequest_RenderedPageFilter_name = map[int32]string{
	1: "all",
	2: "published",
	3: "draft",
}
var GetRenderedPageRequest_RenderedPageFilter_value = map[string]int32{
	"all":       1,
	"published": 2,
	"draft":     3,
}

func (x GetRenderedPageRequest_RenderedPageFilter) Enum() *GetRenderedPageRequest_RenderedPageFilter {
	p := new(GetRenderedPageRequest_RenderedPageFilter)
	*p = x
	return p
}
func (x GetRenderedPageRequest_RenderedPageFilter) String() string {
	return proto.EnumName(GetRenderedPageRequest_RenderedPageFilter_name, int32(x))
}
func (x *GetRenderedPageRequest_RenderedPageFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetRenderedPageRequest_RenderedPageFilter_value, data, "GetRenderedPageRequest_RenderedPageFilter")
	if err != nil {
		return err
	}
	*x = GetRenderedPageRequest_RenderedPageFilter(value)
	return nil
}
func (GetRenderedPageRequest_RenderedPageFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type ListPageRequest_ListPageFilter int32

const (
	ListPageRequest_all       ListPageRequest_ListPageFilter = 1
	ListPageRequest_published ListPageRequest_ListPageFilter = 2
	ListPageRequest_draft     ListPageRequest_ListPageFilter = 3
)

var ListPageRequest_ListPageFilter_name = map[int32]string{
	1: "all",
	2: "published",
	3: "draft",
}
var ListPageRequest_ListPageFilter_value = map[string]int32{
	"all":       1,
	"published": 2,
	"draft":     3,
}

func (x ListPageRequest_ListPageFilter) Enum() *ListPageRequest_ListPageFilter {
	p := new(ListPageRequest_ListPageFilter)
	*p = x
	return p
}
func (x ListPageRequest_ListPageFilter) String() string {
	return proto.EnumName(ListPageRequest_ListPageFilter_name, int32(x))
}
func (x *ListPageRequest_ListPageFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ListPageRequest_ListPageFilter_value, data, "ListPageRequest_ListPageFilter")
	if err != nil {
		return err
	}
	*x = ListPageRequest_ListPageFilter(value)
	return nil
}
func (ListPageRequest_ListPageFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type Error struct {
	Code             *ErrorCode    `protobuf:"varint,1,opt,name=code,enum=ketchup.api.ErrorCode" json:"code,omitempty"`
	Title            *string       `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Detail           *string       `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
	Fields           []*FieldError `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() ErrorCode {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ErrorCode_INTERNAL_SERVER_ERROR
}

func (m *Error) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Error) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *Error) GetFields() []*FieldError {
	if m != nil {
		return m.Fields
	}
	return nil
}

type FieldError struct {
	Field            *string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	Code             *string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Title            *string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Detail           *string `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FieldError) Reset()                    { *m = FieldError{} }
func (m *FieldError) String() string            { return proto.CompactTextString(m) }
func (*FieldError) ProtoMessage()               {}
func (*FieldError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FieldError) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

func (m *FieldError) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *FieldError) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *FieldError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type ListOptions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListOptions) Reset()                    { *m = ListOptions{} }
func (m *ListOptions) String() string            { return proto.CompactTextString(m) }
func (*ListOptions) ProtoMessage()               {}
func (*ListOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type UpdatePageRoutesRequest struct {
	Routes           []*ketchup_models.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdatePageRoutesRequest) Reset()                    { *m = UpdatePageRoutesRequest{} }
func (m *UpdatePageRoutesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePageRoutesRequest) ProtoMessage()               {}
func (*UpdatePageRoutesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdatePageRoutesRequest) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type GetRenderedPageRequest struct {
	Options          *GetRenderedPageRequest_RenderedPageOptions `protobuf:"bytes,1,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *GetRenderedPageRequest) Reset()                    { *m = GetRenderedPageRequest{} }
func (m *GetRenderedPageRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRenderedPageRequest) ProtoMessage()               {}
func (*GetRenderedPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRenderedPageRequest) GetOptions() *GetRenderedPageRequest_RenderedPageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetRenderedPageRequest_RenderedPageOptions struct {
	Filter           *GetRenderedPageRequest_RenderedPageFilter `protobuf:"varint,1,opt,name=filter,enum=ketchup.api.GetRenderedPageRequest_RenderedPageFilter" json:"filter,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *GetRenderedPageRequest_RenderedPageOptions) Reset() {
	*m = GetRenderedPageRequest_RenderedPageOptions{}
}
func (m *GetRenderedPageRequest_RenderedPageOptions) String() string {
	return proto.CompactTextString(m)
}
func (*GetRenderedPageRequest_RenderedPageOptions) ProtoMessage() {}
func (*GetRenderedPageRequest_RenderedPageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *GetRenderedPageRequest_RenderedPageOptions) GetFilter() GetRenderedPageRequest_RenderedPageFilter {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return GetRenderedPageRequest_all
}

type GetRenderedPageResponse struct {
	Data             *structpb.Struct `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GetRenderedPageResponse) Reset()                    { *m = GetRenderedPageResponse{} }
func (m *GetRenderedPageResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRenderedPageResponse) ProtoMessage()               {}
func (*GetRenderedPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetRenderedPageResponse) GetData() *structpb.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListPageRequest struct {
	List             *ListOptions                     `protobuf:"bytes,1,opt,name=list" json:"list,omitempty"`
	Options          *ListPageRequest_ListPageOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *ListPageRequest) Reset()                    { *m = ListPageRequest{} }
func (m *ListPageRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPageRequest) ProtoMessage()               {}
func (*ListPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListPageRequest) GetList() *ListOptions {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListPageRequest) GetOptions() *ListPageRequest_ListPageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListPageRequest_ListPageOptions struct {
	Filter           *ListPageRequest_ListPageFilter `protobuf:"varint,1,opt,name=filter,enum=ketchup.api.ListPageRequest_ListPageFilter" json:"filter,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ListPageRequest_ListPageOptions) Reset()         { *m = ListPageRequest_ListPageOptions{} }
func (m *ListPageRequest_ListPageOptions) String() string { return proto.CompactTextString(m) }
func (*ListPageRequest_ListPageOptions) ProtoMessage()    {}
func (*ListPageRequest_ListPageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

func (m *ListPageRequest_ListPageOptions) GetFilter() ListPageRequest_ListPageFilter {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return ListPageRequest_all
}

type ListPageResponse struct {
	Pages            []*ketchup_models.Page `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListPageResponse) Reset()                    { *m = ListPageResponse{} }
func (m *ListPageResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPageResponse) ProtoMessage()               {}
func (*ListPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListPageResponse) GetPages() []*ketchup_models.Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

type ListRouteResponse struct {
	Routes           []*ketchup_models.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListRouteResponse) Reset()                    { *m = ListRouteResponse{} }
func (m *ListRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRouteResponse) ProtoMessage()               {}
func (*ListRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListRouteResponse) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ListThemeResponse struct {
	Themes           []*ketchup_models.Theme `protobuf:"bytes,1,rep,name=themes" json:"themes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListThemeResponse) Reset()                    { *m = ListThemeResponse{} }
func (m *ListThemeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListThemeResponse) ProtoMessage()               {}
func (*ListThemeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListThemeResponse) GetThemes() []*ketchup_models.Theme {
	if m != nil {
		return m.Themes
	}
	return nil
}

type TLSSettingsReponse struct {
	TlsEmail         *string `protobuf:"bytes,1,opt,name=tls_email,json=tlsEmail" json:"tls_email,omitempty"`
	TlsDomain        *string `protobuf:"bytes,2,opt,name=tls_domain,json=tlsDomain" json:"tls_domain,omitempty"`
	AgreedOn         *string `protobuf:"bytes,3,opt,name=agreed_on,json=agreedOn" json:"agreed_on,omitempty"`
	TermsOfService   *string `protobuf:"bytes,4,opt,name=terms_of_service,json=termsOfService" json:"terms_of_service,omitempty"`
	HasCertificate   *bool   `protobuf:"varint,5,opt,name=has_certificate,json=hasCertificate" json:"has_certificate,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TLSSettingsReponse) Reset()                    { *m = TLSSettingsReponse{} }
func (m *TLSSettingsReponse) String() string            { return proto.CompactTextString(m) }
func (*TLSSettingsReponse) ProtoMessage()               {}
func (*TLSSettingsReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TLSSettingsReponse) GetTlsEmail() string {
	if m != nil && m.TlsEmail != nil {
		return *m.TlsEmail
	}
	return ""
}

func (m *TLSSettingsReponse) GetTlsDomain() string {
	if m != nil && m.TlsDomain != nil {
		return *m.TlsDomain
	}
	return ""
}

func (m *TLSSettingsReponse) GetAgreedOn() string {
	if m != nil && m.AgreedOn != nil {
		return *m.AgreedOn
	}
	return ""
}

func (m *TLSSettingsReponse) GetTermsOfService() string {
	if m != nil && m.TermsOfService != nil {
		return *m.TermsOfService
	}
	return ""
}

func (m *TLSSettingsReponse) GetHasCertificate() bool {
	if m != nil && m.HasCertificate != nil {
		return *m.HasCertificate
	}
	return false
}

type EnableTLSRequest struct {
	TlsEmail         *string `protobuf:"bytes,1,opt,name=tls_email,json=tlsEmail" json:"tls_email,omitempty"`
	TlsDomain        *string `protobuf:"bytes,2,opt,name=tls_domain,json=tlsDomain" json:"tls_domain,omitempty"`
	Agreed           *bool   `protobuf:"varint,3,opt,name=agreed" json:"agreed,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EnableTLSRequest) Reset()                    { *m = EnableTLSRequest{} }
func (m *EnableTLSRequest) String() string            { return proto.CompactTextString(m) }
func (*EnableTLSRequest) ProtoMessage()               {}
func (*EnableTLSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *EnableTLSRequest) GetTlsEmail() string {
	if m != nil && m.TlsEmail != nil {
		return *m.TlsEmail
	}
	return ""
}

func (m *EnableTLSRequest) GetTlsDomain() string {
	if m != nil && m.TlsDomain != nil {
		return *m.TlsDomain
	}
	return ""
}

func (m *EnableTLSRequest) GetAgreed() bool {
	if m != nil && m.Agreed != nil {
		return *m.Agreed
	}
	return false
}

type InstallThemeRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InstallThemeRequest) Reset()                    { *m = InstallThemeRequest{} }
func (m *InstallThemeRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallThemeRequest) ProtoMessage()               {}
func (*InstallThemeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *InstallThemeRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "ketchup.api.Error")
	proto.RegisterType((*FieldError)(nil), "ketchup.api.FieldError")
	proto.RegisterType((*ListOptions)(nil), "ketchup.api.ListOptions")
	proto.RegisterType((*UpdatePageRoutesRequest)(nil), "ketchup.api.UpdatePageRoutesRequest")
	proto.RegisterType((*GetRenderedPageRequest)(nil), "ketchup.api.GetRenderedPageRequest")
	proto.RegisterType((*GetRenderedPageRequest_RenderedPageOptions)(nil), "ketchup.api.GetRenderedPageRequest.RenderedPageOptions")
	proto.RegisterType((*GetRenderedPageResponse)(nil), "ketchup.api.GetRenderedPageResponse")
	proto.RegisterType((*ListPageRequest)(nil), "ketchup.api.ListPageRequest")
	proto.RegisterType((*ListPageRequest_ListPageOptions)(nil), "ketchup.api.ListPageRequest.ListPageOptions")
	proto.RegisterType((*ListPageResponse)(nil), "ketchup.api.ListPageResponse")
	proto.RegisterType((*ListRouteResponse)(nil), "ketchup.api.ListRouteResponse")
	proto.RegisterType((*ListThemeResponse)(nil), "ketchup.api.ListThemeResponse")
	proto.RegisterType((*TLSSettingsReponse)(nil), "ketchup.api.TLSSettingsReponse")
	proto.RegisterType((*EnableTLSRequest)(nil), "ketchup.api.EnableTLSRequest")
	proto.RegisterType((*InstallThemeRequest)(nil), "ketchup.api.InstallThemeRequest")
	proto.RegisterEnum("ketchup.api.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("ketchup.api.GetRenderedPageRequest_RenderedPageFilter", GetRenderedPageRequest_RenderedPageFilter_name, GetRenderedPageRequest_RenderedPageFilter_value)
	proto.RegisterEnum("ketchup.api.ListPageRequest_ListPageFilter", ListPageRequest_ListPageFilter_name, ListPageRequest_ListPageFilter_value)
}
func (m *Error) SetCode(v *ErrorCode) {
	m.Code = v
}

func (m *Error) SetTitle(v *string) {
	m.Title = v
}

func (m *Error) SetDetail(v *string) {
	m.Detail = v
}

func (m *Error) SetFields(v []*FieldError) {
	m.Fields = v
}

func (m *FieldError) SetField(v *string) {
	m.Field = v
}

func (m *FieldError) SetCode(v *string) {
	m.Code = v
}

func (m *FieldError) SetTitle(v *string) {
	m.Title = v
}

func (m *FieldError) SetDetail(v *string) {
	m.Detail = v
}

func (m *UpdatePageRoutesRequest) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func (m *GetRenderedPageRequest) SetOptions(v *GetRenderedPageRequest_RenderedPageOptions) {
	m.Options = v
}

func (m *GetRenderedPageResponse) SetData(v *structpb.Struct) {
	m.Data = v
}

func (m *ListPageRequest) SetList(v *ListOptions) {
	m.List = v
}

func (m *ListPageRequest) SetOptions(v *ListPageRequest_ListPageOptions) {
	m.Options = v
}

func (m *ListPageResponse) SetPages(v []*ketchup_models.Page) {
	m.Pages = v
}

func (m *ListRouteResponse) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func (m *ListThemeResponse) SetThemes(v []*ketchup_models.Theme) {
	m.Themes = v
}

func (m *TLSSettingsReponse) SetTlsEmail(v *string) {
	m.TlsEmail = v
}

func (m *TLSSettingsReponse) SetTlsDomain(v *string) {
	m.TlsDomain = v
}

func (m *TLSSettingsReponse) SetAgreedOn(v *string) {
	m.AgreedOn = v
}

func (m *TLSSettingsReponse) SetTermsOfService(v *string) {
	m.TermsOfService = v
}

func (m *TLSSettingsReponse) SetHasCertificate(v *bool) {
	m.HasCertificate = v
}

func (m *EnableTLSRequest) SetTlsEmail(v *string) {
	m.TlsEmail = v
}

func (m *EnableTLSRequest) SetTlsDomain(v *string) {
	m.TlsDomain = v
}

func (m *EnableTLSRequest) SetAgreed(v *bool) {
	m.Agreed = v
}

func (m *InstallThemeRequest) SetName(v *string) {
	m.Name = v
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xda, 0x4a,
	0x10, 0x95, 0xf9, 0x0a, 0x0c, 0x37, 0xc4, 0x77, 0x93, 0x00, 0x37, 0x57, 0x57, 0x42, 0x7e, 0xb9,
	0x34, 0x49, 0x5d, 0x29, 0x95, 0x9a, 0xb7, 0x4a, 0x4d, 0x30, 0x6d, 0x24, 0x04, 0xed, 0x42, 0xf2,
	0x6a, 0x6d, 0xf0, 0x00, 0x56, 0x8d, 0xd7, 0xf5, 0x2e, 0xfd, 0x13, 0x7d, 0xab, 0xfa, 0x83, 0xfa,
	0xd3, 0x2a, 0xef, 0x2e, 0xe1, 0x2b, 0x95, 0x12, 0xf5, 0x8d, 0x99, 0x39, 0x73, 0xe6, 0xec, 0xf1,
	0x30, 0x50, 0x61, 0x49, 0xe8, 0x26, 0x29, 0x97, 0x9c, 0x54, 0x3f, 0xa3, 0x1c, 0xcf, 0x16, 0x89,
	0xcb, 0x92, 0xf0, 0x04, 0x12, 0x36, 0x45, 0x5d, 0x38, 0xa9, 0xa6, 0x7c, 0x21, 0x1f, 0x02, 0x39,
	0xc3, 0xf9, 0x32, 0xf8, 0x4b, 0xc8, 0x74, 0x31, 0x96, 0x3a, 0x72, 0xbe, 0x5b, 0x50, 0xf4, 0xd2,
	0x94, 0xa7, 0xe4, 0x14, 0x0a, 0x63, 0x1e, 0x60, 0xd3, 0x6a, 0x59, 0xed, 0xda, 0x45, 0xdd, 0x5d,
	0x63, 0x76, 0x15, 0xe2, 0x9a, 0x07, 0x48, 0x15, 0x86, 0x1c, 0x41, 0x51, 0x86, 0x32, 0xc2, 0x66,
	0xae, 0x65, 0xb5, 0x2b, 0x54, 0x07, 0xa4, 0x0e, 0xa5, 0x00, 0x25, 0x0b, 0xa3, 0x66, 0x5e, 0xa5,
	0x4d, 0x44, 0x5e, 0x41, 0x69, 0x12, 0x62, 0x14, 0x88, 0x66, 0xa1, 0x95, 0x6f, 0x57, 0x2f, 0x1a,
	0x1b, 0xdc, 0xdd, 0xac, 0xa4, 0x06, 0x50, 0x03, 0x73, 0x02, 0x80, 0x55, 0x36, 0x1b, 0xa6, 0xf2,
	0x4a, 0x59, 0x85, 0xea, 0x80, 0x10, 0x23, 0x57, 0x2b, 0xd8, 0x92, 0x95, 0x7f, 0x5c, 0x56, 0x61,
	0x5d, 0x96, 0xb3, 0x0f, 0xd5, 0x5e, 0x28, 0xe4, 0x20, 0x91, 0x21, 0x8f, 0x85, 0xf3, 0x01, 0x1a,
	0xb7, 0x49, 0xc0, 0x24, 0x7e, 0x64, 0x53, 0xa4, 0x99, 0x7b, 0x82, 0xe2, 0x97, 0x05, 0x0a, 0x49,
	0x5e, 0x42, 0x49, 0xd9, 0x29, 0x9a, 0x96, 0x7a, 0xc0, 0xf1, 0xc3, 0x03, 0xe6, 0x3c, 0xc0, 0x48,
	0xb8, 0x0a, 0x4e, 0x0d, 0xc8, 0xf9, 0x96, 0x83, 0xfa, 0x7b, 0x94, 0x14, 0xe3, 0x00, 0x53, 0x0c,
	0x14, 0x9f, 0x61, 0xfa, 0x04, 0x7b, 0x5c, 0xcf, 0x53, 0xaf, 0xa9, 0x5e, 0x5c, 0x6e, 0x78, 0xf1,
	0x78, 0x97, 0xbb, 0x9e, 0x33, 0x72, 0xe9, 0x92, 0xe7, 0x04, 0xe1, 0xf0, 0x91, 0x3a, 0xe9, 0x67,
	0xa6, 0x47, 0x12, 0x53, 0xf3, 0x41, 0xdf, 0x3c, 0x77, 0x50, 0x57, 0x75, 0x53, 0xc3, 0xe2, 0x5c,
	0x02, 0xd9, 0xad, 0x92, 0x3d, 0xc8, 0xb3, 0x28, 0xb2, 0x2d, 0xb2, 0x0f, 0x95, 0x64, 0x71, 0x1f,
	0x85, 0x62, 0x86, 0x81, 0x9d, 0x23, 0x15, 0x28, 0x06, 0x29, 0x9b, 0x48, 0x3b, 0xef, 0x74, 0xa1,
	0xb1, 0x33, 0x4d, 0x24, 0x3c, 0x16, 0x48, 0xce, 0xa0, 0x10, 0x30, 0xc9, 0x8c, 0x15, 0x0d, 0x77,
	0xca, 0xf9, 0x34, 0x32, 0x7b, 0x7a, 0xbf, 0x98, 0xb8, 0x43, 0xb5, 0xa9, 0x54, 0x81, 0x9c, 0x1f,
	0x39, 0x38, 0xc8, 0xbe, 0xd7, 0xba, 0x9d, 0xe7, 0x50, 0x88, 0x42, 0x21, 0x0d, 0x41, 0x73, 0xe3,
	0x89, 0x6b, 0xdf, 0x96, 0x2a, 0x14, 0xe9, 0xae, 0xcc, 0xcf, 0xa9, 0x86, 0xf3, 0x9d, 0x86, 0x75,
	0x33, 0x96, 0xf1, 0x8e, 0xe3, 0x77, 0x2b, 0x21, 0x4b, 0xb7, 0xaf, 0xb7, 0xdc, 0x3e, 0x7b, 0x12,
	0xf3, 0x96, 0xc5, 0xaf, 0xa1, 0xb6, 0x59, 0x79, 0x8a, 0xbd, 0x6f, 0xc1, 0x5e, 0xd1, 0x1b, 0x5f,
	0x4f, 0xa1, 0x98, 0x9d, 0x82, 0xe5, 0xba, 0x1e, 0x6d, 0xaf, 0xab, 0x02, 0x6b, 0x88, 0x73, 0x05,
	0x7f, 0x67, 0xfd, 0x7a, 0x83, 0x97, 0x04, 0xcf, 0x5c, 0x78, 0xc3, 0x31, 0xca, 0xae, 0xcc, 0x3a,
	0x87, 0x3a, 0x3b, 0xbf, 0xe5, 0xd0, 0x70, 0x03, 0x72, 0x7e, 0x5a, 0x40, 0x46, 0xbd, 0xe1, 0x10,
	0xa5, 0x0c, 0xe3, 0xa9, 0xa0, 0xa8, 0x59, 0xfe, 0x85, 0x8a, 0x8c, 0x84, 0x8f, 0xf3, 0xec, 0xff,
	0xab, 0x0f, 0x40, 0x59, 0x46, 0xc2, 0xcb, 0x62, 0xf2, 0x1f, 0x40, 0x56, 0x0c, 0xf8, 0x9c, 0x85,
	0xb1, 0xb9, 0x04, 0x19, 0xbc, 0xa3, 0x12, 0x59, 0x2f, 0x9b, 0xa6, 0x88, 0x81, 0xcf, 0x63, 0x73,
	0x12, 0xca, 0x3a, 0x31, 0x88, 0x49, 0x1b, 0x6c, 0x89, 0xe9, 0x5c, 0xf8, 0x7c, 0xe2, 0x0b, 0x4c,
	0xbf, 0x86, 0x63, 0x34, 0xf7, 0xa1, 0xa6, 0xf2, 0x83, 0xc9, 0x50, 0x67, 0xc9, 0xff, 0x70, 0x30,
	0x63, 0xc2, 0x1f, 0x63, 0x2a, 0xc3, 0x49, 0x38, 0x66, 0x12, 0x9b, 0xc5, 0x96, 0xd5, 0x2e, 0xd3,
	0xda, 0x8c, 0x89, 0xeb, 0x55, 0xd6, 0x99, 0x80, 0xed, 0xc5, 0xec, 0x3e, 0xc2, 0x51, 0x6f, 0xb8,
	0xdc, 0xd0, 0x3f, 0xd1, 0x5f, 0x87, 0x92, 0x96, 0xab, 0xc4, 0x97, 0xa9, 0x89, 0x9c, 0x17, 0x70,
	0x78, 0x13, 0x0b, 0xc9, 0xa2, 0xc8, 0x38, 0xae, 0x47, 0x11, 0x28, 0xc4, 0x6c, 0x8e, 0x66, 0x8a,
	0xfa, 0x7d, 0xda, 0x81, 0xca, 0xc3, 0xed, 0x26, 0xff, 0xc0, 0xf1, 0x4d, 0x7f, 0xe4, 0xd1, 0xfe,
	0xbb, 0x9e, 0x3f, 0xf4, 0xe8, 0x9d, 0x47, 0x7d, 0x8f, 0xd2, 0x01, 0xd5, 0xfb, 0xd5, 0x1f, 0x8c,
	0xfc, 0xee, 0xe0, 0xb6, 0xdf, 0xb1, 0x73, 0x59, 0xd8, 0x1d, 0xd0, 0xab, 0x9b, 0x4e, 0xc7, 0xeb,
	0xdb, 0xf9, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0xc1, 0xf3, 0x1d, 0x71, 0x06, 0x00, 0x00,
}
