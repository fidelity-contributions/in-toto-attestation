// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: in_toto_attestation/predicates/vulns/v0.1/vulns.proto

package v01

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Validation of all fields is left to the users of this proto.
type Vulns struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scanner       *Scanner               `protobuf:"bytes,1,opt,name=scanner,proto3" json:"scanner,omitempty"`
	ScanMetadata  *ScanMetadata          `protobuf:"bytes,2,opt,name=scan_metadata,json=scanMetadata,proto3" json:"scan_metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vulns) Reset() {
	*x = Vulns{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vulns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulns) ProtoMessage() {}

func (x *Vulns) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulns.ProtoReflect.Descriptor instead.
func (*Vulns) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{0}
}

func (x *Vulns) GetScanner() *Scanner {
	if x != nil {
		return x.Scanner
	}
	return nil
}

func (x *Vulns) GetScanMetadata() *ScanMetadata {
	if x != nil {
		return x.ScanMetadata
	}
	return nil
}

type Scanner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uri           string                 `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Version       *string                `protobuf:"bytes,2,opt,name=version,proto3,oneof" json:"version,omitempty"`
	Database      *VulnDatabase          `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	Result        *Result                `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scanner) Reset() {
	*x = Scanner{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scanner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scanner) ProtoMessage() {}

func (x *Scanner) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scanner.ProtoReflect.Descriptor instead.
func (*Scanner) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{1}
}

func (x *Scanner) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Scanner) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *Scanner) GetDatabase() *VulnDatabase {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Scanner) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type VulnDatabase struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uri           *string                `protobuf:"bytes,1,opt,name=uri,proto3,oneof" json:"uri,omitempty"`
	Version       *string                `protobuf:"bytes,2,opt,name=version,proto3,oneof" json:"version,omitempty"`
	LastUpdate    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VulnDatabase) Reset() {
	*x = VulnDatabase{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnDatabase) ProtoMessage() {}

func (x *VulnDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnDatabase.ProtoReflect.Descriptor instead.
func (*VulnDatabase) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{2}
}

func (x *VulnDatabase) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *VulnDatabase) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *VulnDatabase) GetLastUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdate
	}
	return nil
}

type Result struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Vulnerabilities []*Vulnerability       `protobuf:"bytes,3,rep,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Result) Reset() {
	*x = Result{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetVulnerabilities() []*Vulnerability {
	if x != nil {
		return x.Vulnerabilities
	}
	return nil
}

type Vulnerability struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Id            string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Severity      []*Vulnerability_Severity `protobuf:"bytes,2,rep,name=severity,proto3" json:"severity,omitempty"`
	Annotations   []*structpb.Struct        `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vulnerability) Reset() {
	*x = Vulnerability{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulnerability) ProtoMessage() {}

func (x *Vulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulnerability.ProtoReflect.Descriptor instead.
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{4}
}

func (x *Vulnerability) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vulnerability) GetSeverity() []*Vulnerability_Severity {
	if x != nil {
		return x.Severity
	}
	return nil
}

func (x *Vulnerability) GetAnnotations() []*structpb.Struct {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type ScanMetadata struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ScanStartedOn  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=scan_started_on,json=scanStartedOn,proto3" json:"scan_started_on,omitempty"`
	ScanFinishedOn *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scan_finished_on,json=scanFinishedOn,proto3" json:"scan_finished_on,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ScanMetadata) Reset() {
	*x = ScanMetadata{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScanMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanMetadata) ProtoMessage() {}

func (x *ScanMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanMetadata.ProtoReflect.Descriptor instead.
func (*ScanMetadata) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{5}
}

func (x *ScanMetadata) GetScanStartedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.ScanStartedOn
	}
	return nil
}

func (x *ScanMetadata) GetScanFinishedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.ScanFinishedOn
	}
	return nil
}

type Vulnerability_Severity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Score         string                 `protobuf:"bytes,2,opt,name=score,proto3" json:"score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vulnerability_Severity) Reset() {
	*x = Vulnerability_Severity{}
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vulnerability_Severity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulnerability_Severity) ProtoMessage() {}

func (x *Vulnerability_Severity) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulnerability_Severity.ProtoReflect.Descriptor instead.
func (*Vulnerability_Severity) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Vulnerability_Severity) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Vulnerability_Severity) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

var File_in_toto_attestation_predicates_vulns_v0_1_vulns_proto protoreflect.FileDescriptor

const file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDesc = "" +
	"\n" +
	"5in_toto_attestation/predicates/vulns/v0.1/vulns.proto\x12(in_toto_attestation.predicates.vulns.v01\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb1\x01\n" +
	"\x05Vulns\x12K\n" +
	"\ascanner\x18\x01 \x01(\v21.in_toto_attestation.predicates.vulns.v01.ScannerR\ascanner\x12[\n" +
	"\rscan_metadata\x18\x02 \x01(\v26.in_toto_attestation.predicates.vulns.v01.ScanMetadataR\fscanMetadata\"\xe4\x01\n" +
	"\aScanner\x12\x10\n" +
	"\x03uri\x18\x01 \x01(\tR\x03uri\x12\x1d\n" +
	"\aversion\x18\x02 \x01(\tH\x00R\aversion\x88\x01\x01\x12R\n" +
	"\bdatabase\x18\x03 \x01(\v26.in_toto_attestation.predicates.vulns.v01.VulnDatabaseR\bdatabase\x12H\n" +
	"\x06result\x18\x04 \x01(\v20.in_toto_attestation.predicates.vulns.v01.ResultR\x06resultB\n" +
	"\n" +
	"\b_version\"\x95\x01\n" +
	"\fVulnDatabase\x12\x15\n" +
	"\x03uri\x18\x01 \x01(\tH\x00R\x03uri\x88\x01\x01\x12\x1d\n" +
	"\aversion\x18\x02 \x01(\tH\x01R\aversion\x88\x01\x01\x12;\n" +
	"\vlast_update\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"lastUpdateB\x06\n" +
	"\x04_uriB\n" +
	"\n" +
	"\b_version\"k\n" +
	"\x06Result\x12a\n" +
	"\x0fvulnerabilities\x18\x03 \x03(\v27.in_toto_attestation.predicates.vulns.v01.VulnerabilityR\x0fvulnerabilities\"\xf2\x01\n" +
	"\rVulnerability\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\\\n" +
	"\bseverity\x18\x02 \x03(\v2@.in_toto_attestation.predicates.vulns.v01.Vulnerability.SeverityR\bseverity\x129\n" +
	"\vannotations\x18\x03 \x03(\v2\x17.google.protobuf.StructR\vannotations\x1a8\n" +
	"\bSeverity\x12\x16\n" +
	"\x06method\x18\x01 \x01(\tR\x06method\x12\x14\n" +
	"\x05score\x18\x02 \x01(\tR\x05score\"\x98\x01\n" +
	"\fScanMetadata\x12B\n" +
	"\x0fscan_started_on\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\rscanStartedOn\x12D\n" +
	"\x10scan_finished_on\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x0escanFinishedOnBk\n" +
	"1io.github.intoto.attestation.predicates.vulns.v01Z6github.com/in-toto/attestation/go/predicates/vulns/v01b\x06proto3"

var (
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescOnce sync.Once
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescData []byte
)

func file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescGZIP() []byte {
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescOnce.Do(func() {
		file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDesc), len(file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDesc)))
	})
	return file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDescData
}

var file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_goTypes = []any{
	(*Vulns)(nil),                  // 0: in_toto_attestation.predicates.vulns.v01.Vulns
	(*Scanner)(nil),                // 1: in_toto_attestation.predicates.vulns.v01.Scanner
	(*VulnDatabase)(nil),           // 2: in_toto_attestation.predicates.vulns.v01.VulnDatabase
	(*Result)(nil),                 // 3: in_toto_attestation.predicates.vulns.v01.Result
	(*Vulnerability)(nil),          // 4: in_toto_attestation.predicates.vulns.v01.Vulnerability
	(*ScanMetadata)(nil),           // 5: in_toto_attestation.predicates.vulns.v01.ScanMetadata
	(*Vulnerability_Severity)(nil), // 6: in_toto_attestation.predicates.vulns.v01.Vulnerability.Severity
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 8: google.protobuf.Struct
}
var file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_depIdxs = []int32{
	1,  // 0: in_toto_attestation.predicates.vulns.v01.Vulns.scanner:type_name -> in_toto_attestation.predicates.vulns.v01.Scanner
	5,  // 1: in_toto_attestation.predicates.vulns.v01.Vulns.scan_metadata:type_name -> in_toto_attestation.predicates.vulns.v01.ScanMetadata
	2,  // 2: in_toto_attestation.predicates.vulns.v01.Scanner.database:type_name -> in_toto_attestation.predicates.vulns.v01.VulnDatabase
	3,  // 3: in_toto_attestation.predicates.vulns.v01.Scanner.result:type_name -> in_toto_attestation.predicates.vulns.v01.Result
	7,  // 4: in_toto_attestation.predicates.vulns.v01.VulnDatabase.last_update:type_name -> google.protobuf.Timestamp
	4,  // 5: in_toto_attestation.predicates.vulns.v01.Result.vulnerabilities:type_name -> in_toto_attestation.predicates.vulns.v01.Vulnerability
	6,  // 6: in_toto_attestation.predicates.vulns.v01.Vulnerability.severity:type_name -> in_toto_attestation.predicates.vulns.v01.Vulnerability.Severity
	8,  // 7: in_toto_attestation.predicates.vulns.v01.Vulnerability.annotations:type_name -> google.protobuf.Struct
	7,  // 8: in_toto_attestation.predicates.vulns.v01.ScanMetadata.scan_started_on:type_name -> google.protobuf.Timestamp
	7,  // 9: in_toto_attestation.predicates.vulns.v01.ScanMetadata.scan_finished_on:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_init() }
func file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_init() {
	if File_in_toto_attestation_predicates_vulns_v0_1_vulns_proto != nil {
		return
	}
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[1].OneofWrappers = []any{}
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDesc), len(file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_goTypes,
		DependencyIndexes: file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_depIdxs,
		MessageInfos:      file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_msgTypes,
	}.Build()
	File_in_toto_attestation_predicates_vulns_v0_1_vulns_proto = out.File
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_goTypes = nil
	file_in_toto_attestation_predicates_vulns_v0_1_vulns_proto_depIdxs = nil
}
