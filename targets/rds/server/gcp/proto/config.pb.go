// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/rds/server/gcp/proto/config.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GCEInstances struct {
	// Zone filter, e.g. zone_filter: "us-east1-*" for all zones in us-east1
	// region.
	ZoneFilter *string `protobuf:"bytes,1,opt,name=zone_filter,json=zoneFilter" json:"zone_filter,omitempty"`
	// How often resources should be evaluated/expanded.
	ReEvalSec            *int32   `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=300" json:"re_eval_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCEInstances) Reset()         { *m = GCEInstances{} }
func (m *GCEInstances) String() string { return proto.CompactTextString(m) }
func (*GCEInstances) ProtoMessage()    {}
func (*GCEInstances) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{0}
}

func (m *GCEInstances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCEInstances.Unmarshal(m, b)
}
func (m *GCEInstances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCEInstances.Marshal(b, m, deterministic)
}
func (m *GCEInstances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCEInstances.Merge(m, src)
}
func (m *GCEInstances) XXX_Size() int {
	return xxx_messageInfo_GCEInstances.Size(m)
}
func (m *GCEInstances) XXX_DiscardUnknown() {
	xxx_messageInfo_GCEInstances.DiscardUnknown(m)
}

var xxx_messageInfo_GCEInstances proto.InternalMessageInfo

const Default_GCEInstances_ReEvalSec int32 = 300

func (m *GCEInstances) GetZoneFilter() string {
	if m != nil && m.ZoneFilter != nil {
		return *m.ZoneFilter
	}
	return ""
}

func (m *GCEInstances) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_GCEInstances_ReEvalSec
}

type RegionalForwardingRules struct {
	// Region filter, e.g. "region_filter:europe-*"
	RegionFilter         *string  `protobuf:"bytes,1,opt,name=region_filter,json=regionFilter" json:"region_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionalForwardingRules) Reset()         { *m = RegionalForwardingRules{} }
func (m *RegionalForwardingRules) String() string { return proto.CompactTextString(m) }
func (*RegionalForwardingRules) ProtoMessage()    {}
func (*RegionalForwardingRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{1}
}

func (m *RegionalForwardingRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionalForwardingRules.Unmarshal(m, b)
}
func (m *RegionalForwardingRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionalForwardingRules.Marshal(b, m, deterministic)
}
func (m *RegionalForwardingRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionalForwardingRules.Merge(m, src)
}
func (m *RegionalForwardingRules) XXX_Size() int {
	return xxx_messageInfo_RegionalForwardingRules.Size(m)
}
func (m *RegionalForwardingRules) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionalForwardingRules.DiscardUnknown(m)
}

var xxx_messageInfo_RegionalForwardingRules proto.InternalMessageInfo

func (m *RegionalForwardingRules) GetRegionFilter() string {
	if m != nil && m.RegionFilter != nil {
		return *m.RegionFilter
	}
	return ""
}

// Runtime configurator variables.
type RTCVariables struct {
	RtcConfig            []*RTCVariables_RTCConfig `protobuf:"bytes,1,rep,name=rtc_config,json=rtcConfig" json:"rtc_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RTCVariables) Reset()         { *m = RTCVariables{} }
func (m *RTCVariables) String() string { return proto.CompactTextString(m) }
func (*RTCVariables) ProtoMessage()    {}
func (*RTCVariables) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{2}
}

func (m *RTCVariables) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTCVariables.Unmarshal(m, b)
}
func (m *RTCVariables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTCVariables.Marshal(b, m, deterministic)
}
func (m *RTCVariables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTCVariables.Merge(m, src)
}
func (m *RTCVariables) XXX_Size() int {
	return xxx_messageInfo_RTCVariables.Size(m)
}
func (m *RTCVariables) XXX_DiscardUnknown() {
	xxx_messageInfo_RTCVariables.DiscardUnknown(m)
}

var xxx_messageInfo_RTCVariables proto.InternalMessageInfo

func (m *RTCVariables) GetRtcConfig() []*RTCVariables_RTCConfig {
	if m != nil {
		return m.RtcConfig
	}
	return nil
}

type RTCVariables_RTCConfig struct {
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// How often RTC variables should be evaluated/expanded.
	ReEvalSec            *int32   `protobuf:"varint,2,opt,name=re_eval_sec,json=reEvalSec,def=10" json:"re_eval_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RTCVariables_RTCConfig) Reset()         { *m = RTCVariables_RTCConfig{} }
func (m *RTCVariables_RTCConfig) String() string { return proto.CompactTextString(m) }
func (*RTCVariables_RTCConfig) ProtoMessage()    {}
func (*RTCVariables_RTCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{2, 0}
}

func (m *RTCVariables_RTCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTCVariables_RTCConfig.Unmarshal(m, b)
}
func (m *RTCVariables_RTCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTCVariables_RTCConfig.Marshal(b, m, deterministic)
}
func (m *RTCVariables_RTCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTCVariables_RTCConfig.Merge(m, src)
}
func (m *RTCVariables_RTCConfig) XXX_Size() int {
	return xxx_messageInfo_RTCVariables_RTCConfig.Size(m)
}
func (m *RTCVariables_RTCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RTCVariables_RTCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RTCVariables_RTCConfig proto.InternalMessageInfo

const Default_RTCVariables_RTCConfig_ReEvalSec int32 = 10

func (m *RTCVariables_RTCConfig) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RTCVariables_RTCConfig) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_RTCVariables_RTCConfig_ReEvalSec
}

// Runtime configurator variables.
type PubSubMessages struct {
	Subscription []*PubSubMessages_Subscription `protobuf:"bytes,1,rep,name=subscription" json:"subscription,omitempty"`
	// Only for testing.
	ApiEndpoint          *string  `protobuf:"bytes,2,opt,name=api_endpoint,json=apiEndpoint" json:"api_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubSubMessages) Reset()         { *m = PubSubMessages{} }
func (m *PubSubMessages) String() string { return proto.CompactTextString(m) }
func (*PubSubMessages) ProtoMessage()    {}
func (*PubSubMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{3}
}

func (m *PubSubMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubMessages.Unmarshal(m, b)
}
func (m *PubSubMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubMessages.Marshal(b, m, deterministic)
}
func (m *PubSubMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubMessages.Merge(m, src)
}
func (m *PubSubMessages) XXX_Size() int {
	return xxx_messageInfo_PubSubMessages.Size(m)
}
func (m *PubSubMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubMessages.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubMessages proto.InternalMessageInfo

func (m *PubSubMessages) GetSubscription() []*PubSubMessages_Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *PubSubMessages) GetApiEndpoint() string {
	if m != nil && m.ApiEndpoint != nil {
		return *m.ApiEndpoint
	}
	return ""
}

type PubSubMessages_Subscription struct {
	// Subscription name. If it doesn't exist already, we try to create one.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Topic name. This is used to create the subscription if it doesn't exist
	// already.
	TopicName *string `protobuf:"bytes,2,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	// If subscription already exists, how far back to seek back on restart.
	// Note that duplicate data is fine as we filter by publish time.
	SeekBackDurationSec  *int32   `protobuf:"varint,3,opt,name=seek_back_duration_sec,json=seekBackDurationSec,def=3600" json:"seek_back_duration_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubSubMessages_Subscription) Reset()         { *m = PubSubMessages_Subscription{} }
func (m *PubSubMessages_Subscription) String() string { return proto.CompactTextString(m) }
func (*PubSubMessages_Subscription) ProtoMessage()    {}
func (*PubSubMessages_Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{3, 0}
}

func (m *PubSubMessages_Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubMessages_Subscription.Unmarshal(m, b)
}
func (m *PubSubMessages_Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubMessages_Subscription.Marshal(b, m, deterministic)
}
func (m *PubSubMessages_Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubMessages_Subscription.Merge(m, src)
}
func (m *PubSubMessages_Subscription) XXX_Size() int {
	return xxx_messageInfo_PubSubMessages_Subscription.Size(m)
}
func (m *PubSubMessages_Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubMessages_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubMessages_Subscription proto.InternalMessageInfo

const Default_PubSubMessages_Subscription_SeekBackDurationSec int32 = 3600

func (m *PubSubMessages_Subscription) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PubSubMessages_Subscription) GetTopicName() string {
	if m != nil && m.TopicName != nil {
		return *m.TopicName
	}
	return ""
}

func (m *PubSubMessages_Subscription) GetSeekBackDurationSec() int32 {
	if m != nil && m.SeekBackDurationSec != nil {
		return *m.SeekBackDurationSec
	}
	return Default_PubSubMessages_Subscription_SeekBackDurationSec
}

// GCP provider config.
type ProviderConfig struct {
	// GCP projects. If running on GCE, it defaults to the local project.
	Project []string `protobuf:"bytes,1,rep,name=project" json:"project,omitempty"`
	// GCE instances discovery options. This field should be declared for the GCE
	// instances discovery to be enabled.
	GceInstances *GCEInstances `protobuf:"bytes,2,opt,name=gce_instances,json=gceInstances" json:"gce_instances,omitempty"`
	// Regional forwarding rules discovery options. This field should be declared
	// for the regional forwarding rules discovery to be enabled.
	// NOTE: This is not supported yet.
	RegionalForwardingRules *RegionalForwardingRules `protobuf:"bytes,3,opt,name=regional_forwarding_rules,json=regionalForwardingRules" json:"regional_forwarding_rules,omitempty"`
	// RTC variables discovery options.
	RtcVariables *RTCVariables `protobuf:"bytes,4,opt,name=rtc_variables,json=rtcVariables" json:"rtc_variables,omitempty"`
	// PubSub messages discovery options.
	PubsubMessages *PubSubMessages `protobuf:"bytes,5,opt,name=pubsub_messages,json=pubsubMessages" json:"pubsub_messages,omitempty"`
	// Compute API version.
	ApiVersion           *string  `protobuf:"bytes,99,opt,name=api_version,json=apiVersion,def=v1" json:"api_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderConfig) Reset()         { *m = ProviderConfig{} }
func (m *ProviderConfig) String() string { return proto.CompactTextString(m) }
func (*ProviderConfig) ProtoMessage()    {}
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_87b43a20d7761773, []int{4}
}

func (m *ProviderConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderConfig.Unmarshal(m, b)
}
func (m *ProviderConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderConfig.Marshal(b, m, deterministic)
}
func (m *ProviderConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderConfig.Merge(m, src)
}
func (m *ProviderConfig) XXX_Size() int {
	return xxx_messageInfo_ProviderConfig.Size(m)
}
func (m *ProviderConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderConfig proto.InternalMessageInfo

const Default_ProviderConfig_ApiVersion string = "v1"

func (m *ProviderConfig) GetProject() []string {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ProviderConfig) GetGceInstances() *GCEInstances {
	if m != nil {
		return m.GceInstances
	}
	return nil
}

func (m *ProviderConfig) GetRegionalForwardingRules() *RegionalForwardingRules {
	if m != nil {
		return m.RegionalForwardingRules
	}
	return nil
}

func (m *ProviderConfig) GetRtcVariables() *RTCVariables {
	if m != nil {
		return m.RtcVariables
	}
	return nil
}

func (m *ProviderConfig) GetPubsubMessages() *PubSubMessages {
	if m != nil {
		return m.PubsubMessages
	}
	return nil
}

func (m *ProviderConfig) GetApiVersion() string {
	if m != nil && m.ApiVersion != nil {
		return *m.ApiVersion
	}
	return Default_ProviderConfig_ApiVersion
}

func init() {
	proto.RegisterType((*GCEInstances)(nil), "cloudprober.targets.rds.gcp.GCEInstances")
	proto.RegisterType((*RegionalForwardingRules)(nil), "cloudprober.targets.rds.gcp.RegionalForwardingRules")
	proto.RegisterType((*RTCVariables)(nil), "cloudprober.targets.rds.gcp.RTCVariables")
	proto.RegisterType((*RTCVariables_RTCConfig)(nil), "cloudprober.targets.rds.gcp.RTCVariables.RTCConfig")
	proto.RegisterType((*PubSubMessages)(nil), "cloudprober.targets.rds.gcp.PubSubMessages")
	proto.RegisterType((*PubSubMessages_Subscription)(nil), "cloudprober.targets.rds.gcp.PubSubMessages.Subscription")
	proto.RegisterType((*ProviderConfig)(nil), "cloudprober.targets.rds.gcp.ProviderConfig")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/targets/rds/server/gcp/proto/config.proto", fileDescriptor_87b43a20d7761773)
}

var fileDescriptor_87b43a20d7761773 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x6a, 0xd4, 0x4e,
	0x14, 0xc6, 0xd9, 0xdd, 0x96, 0x3f, 0x39, 0x9b, 0xf6, 0x0f, 0x11, 0xec, 0x5a, 0x11, 0xd7, 0xed,
	0xcd, 0x8a, 0x90, 0x6c, 0x5b, 0x11, 0xed, 0x85, 0x17, 0xae, 0x6d, 0xf1, 0xc2, 0x52, 0xa6, 0xa5,
	0x57, 0xc2, 0x30, 0x99, 0x9c, 0xc6, 0xb1, 0x69, 0x66, 0x38, 0x33, 0x89, 0x20, 0xbe, 0x83, 0x8f,
	0xe1, 0x23, 0xf8, 0x7a, 0x92, 0x49, 0xda, 0xee, 0x42, 0x5d, 0xf4, 0x2e, 0xf3, 0x9d, 0x39, 0xdf,
	0x9c, 0xef, 0xe4, 0x07, 0xc7, 0xb9, 0x72, 0x9f, 0xab, 0x34, 0x96, 0xfa, 0x3a, 0xc9, 0xb5, 0xce,
	0x0b, 0x4c, 0x64, 0xa1, 0xab, 0xcc, 0x90, 0x4e, 0x91, 0x12, 0x27, 0x28, 0x47, 0x67, 0x13, 0xca,
	0x6c, 0x62, 0x91, 0x6a, 0xa4, 0x24, 0x97, 0x26, 0x31, 0xa4, 0x9d, 0x4e, 0xa4, 0x2e, 0x2f, 0x55,
	0x1e, 0xfb, 0x43, 0xf4, 0x78, 0xa1, 0x2d, 0xee, 0xda, 0x62, 0xca, 0x6c, 0x9c, 0x4b, 0x33, 0x39,
	0x87, 0xf0, 0x78, 0x7e, 0xf8, 0xa1, 0xb4, 0x4e, 0x94, 0x12, 0x6d, 0xf4, 0x14, 0x86, 0xdf, 0x74,
	0x89, 0xfc, 0x52, 0x15, 0x0e, 0x69, 0xd4, 0x1b, 0xf7, 0xa6, 0x01, 0x83, 0x46, 0x3a, 0xf2, 0x4a,
	0xb4, 0x03, 0x43, 0x42, 0x8e, 0xb5, 0x28, 0xb8, 0x45, 0x39, 0x4a, 0xc7, 0xbd, 0xe9, 0xfa, 0xc1,
	0x60, 0x7f, 0x36, 0x63, 0x01, 0xe1, 0x61, 0x2d, 0x8a, 0x33, 0x94, 0x93, 0xb7, 0xb0, 0xc5, 0x30,
	0x57, 0xba, 0x14, 0xc5, 0x91, 0xa6, 0xaf, 0x82, 0x32, 0x55, 0xe6, 0xac, 0x2a, 0xd0, 0x46, 0x3b,
	0xb0, 0x41, 0xbe, 0xb4, 0xfc, 0x44, 0xd8, 0x8a, 0xed, 0x23, 0x93, 0x9f, 0x3d, 0x08, 0xd9, 0xf9,
	0xfc, 0x42, 0x90, 0x12, 0x69, 0xd3, 0xc5, 0x00, 0xc8, 0x49, 0xde, 0xe6, 0x1a, 0xf5, 0xc6, 0x83,
	0xe9, 0x70, 0x6f, 0x3f, 0x5e, 0x11, 0x2c, 0x5e, 0x6c, 0x6f, 0x0e, 0x73, 0xdf, 0xca, 0x02, 0x72,
	0xb2, 0xfd, 0xdc, 0x9e, 0x43, 0x70, 0xab, 0x47, 0x11, 0xac, 0x95, 0xe2, 0x1a, 0x47, 0xbd, 0x71,
	0x7f, 0x1a, 0x30, 0xff, 0x1d, 0x4d, 0x96, 0xa3, 0xf6, 0x7d, 0xd4, 0xfe, 0xee, 0x52, 0xd2, 0x1f,
	0x7d, 0xd8, 0x3c, 0xad, 0xd2, 0xb3, 0x2a, 0xfd, 0x88, 0xd6, 0x8a, 0x1c, 0x6d, 0xf4, 0x09, 0x42,
	0x5b, 0xa5, 0x56, 0x92, 0x32, 0x4e, 0xe9, 0xb2, 0x9b, 0xf6, 0xf5, 0xca, 0x69, 0x97, 0x2d, 0xe2,
	0xb3, 0x85, 0x7e, 0xb6, 0xe4, 0x16, 0x3d, 0x83, 0x50, 0x18, 0xc5, 0xb1, 0xcc, 0x8c, 0x56, 0xa5,
	0xf3, 0x53, 0x05, 0x6c, 0x28, 0x8c, 0x3a, 0xec, 0xa4, 0xed, 0xef, 0x10, 0x2e, 0x1a, 0xdc, 0x9b,
	0xed, 0x09, 0x80, 0xd3, 0x46, 0x49, 0xee, 0x2b, 0xad, 0x49, 0xe0, 0x95, 0x93, 0xa6, 0xfc, 0x06,
	0x1e, 0x5a, 0xc4, 0x2b, 0x9e, 0x0a, 0x79, 0xc5, 0xb3, 0x8a, 0x44, 0x63, 0xe4, 0xb7, 0x30, 0xf0,
	0x5b, 0x58, 0xdb, 0x7f, 0x35, 0x9b, 0xb1, 0x07, 0xcd, 0x9d, 0x77, 0x42, 0x5e, 0xbd, 0xef, 0x6e,
	0x34, 0x1b, 0xf9, 0x35, 0x80, 0xcd, 0x53, 0xd2, 0xb5, 0xca, 0x90, 0xba, 0xe5, 0x8e, 0xe0, 0x3f,
	0x43, 0xfa, 0x0b, 0x4a, 0xe7, 0x97, 0x11, 0xb0, 0x9b, 0x63, 0x74, 0x02, 0x1b, 0xb9, 0x44, 0xae,
	0x6e, 0xf8, 0xf3, 0x93, 0x0c, 0xf7, 0x9e, 0xaf, 0x5c, 0xd6, 0x22, 0xb0, 0x2c, 0xcc, 0x25, 0xde,
	0xe1, 0x6b, 0xe0, 0x11, 0x75, 0xe0, 0xf1, 0xcb, 0x5b, 0xf2, 0x38, 0x35, 0xe8, 0xf9, 0xd1, 0x87,
	0x7b, 0x2f, 0x57, 0x63, 0x73, 0x3f, 0xb6, 0x6c, 0x8b, 0xfe, 0xc0, 0xf3, 0x09, 0x6c, 0x34, 0x64,
	0xd6, 0x37, 0xac, 0x8d, 0xd6, 0xfe, 0x22, 0xc1, 0x22, 0x9c, 0x2c, 0x24, 0x27, 0xef, 0x48, 0x3f,
	0x87, 0xff, 0x4d, 0x95, 0xda, 0x2a, 0xe5, 0xd7, 0x1d, 0x0d, 0xa3, 0x75, 0xef, 0xf8, 0xe2, 0x1f,
	0x00, 0x62, 0x9b, 0xad, 0xc7, 0x2d, 0x93, 0x3b, 0xd0, 0x10, 0xc2, 0x6b, 0x24, 0xdb, 0x20, 0x29,
	0x9b, 0xff, 0x7d, 0xd0, 0xaf, 0x77, 0x19, 0x08, 0xa3, 0x2e, 0x5a, 0xf5, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xfe, 0x3f, 0xda, 0x72, 0x04, 0x00, 0x00,
}