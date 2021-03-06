// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamodel.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventClass int32

const (
	EventClass_ALL    EventClass = 0
	EventClass_PHASE  EventClass = 1
	EventClass_MODULE EventClass = 2
)

var EventClass_name = map[int32]string{
	0: "ALL",
	1: "PHASE",
	2: "MODULE",
}
var EventClass_value = map[string]int32{
	"ALL":    0,
	"PHASE":  1,
	"MODULE": 2,
}

func (x EventClass) String() string {
	return proto.EnumName(EventClass_name, int32(x))
}
func (EventClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{0}
}

type Phase int32

const (
	Phase_INIT     Phase = 0
	Phase_BUILD    Phase = 1
	Phase_ANALYSIS Phase = 2
	Phase_REPORT   Phase = 3
	Phase_FAIL     Phase = 4
)

var Phase_name = map[int32]string{
	0: "INIT",
	1: "BUILD",
	2: "ANALYSIS",
	3: "REPORT",
	4: "FAIL",
}
var Phase_value = map[string]int32{
	"INIT":     0,
	"BUILD":    1,
	"ANALYSIS": 2,
	"REPORT":   3,
	"FAIL":     4,
}

func (x Phase) String() string {
	return proto.EnumName(Phase_name, int32(x))
}
func (Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{1}
}

type ExceptionType int32

const (
	ExceptionType_ERROR   ExceptionType = 0
	ExceptionType_WARNING ExceptionType = 1
)

var ExceptionType_name = map[int32]string{
	0: "ERROR",
	1: "WARNING",
}
var ExceptionType_value = map[string]int32{
	"ERROR":   0,
	"WARNING": 1,
}

func (x ExceptionType) String() string {
	return proto.EnumName(ExceptionType_name, int32(x))
}
func (ExceptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{2}
}

type FileNode_Type int32

const (
	FileNode_UNDEF        FileNode_Type = 0
	FileNode_SOURCE       FileNode_Type = 1
	FileNode_INTERMEDIATE FileNode_Type = 2
	FileNode_TARGET       FileNode_Type = 3
)

var FileNode_Type_name = map[int32]string{
	0: "UNDEF",
	1: "SOURCE",
	2: "INTERMEDIATE",
	3: "TARGET",
}
var FileNode_Type_value = map[string]int32{
	"UNDEF":        0,
	"SOURCE":       1,
	"INTERMEDIATE": 2,
	"TARGET":       3,
}

func (x FileNode_Type) String() string {
	return proto.EnumName(FileNode_Type_name, int32(x))
}
func (FileNode_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{0, 0}
}

type DiagnosticNode_Severity int32

const (
	DiagnosticNode_UNDEF   DiagnosticNode_Severity = 0
	DiagnosticNode_INFO    DiagnosticNode_Severity = 1
	DiagnosticNode_WARNING DiagnosticNode_Severity = 2
	DiagnosticNode_ERROR   DiagnosticNode_Severity = 3
)

var DiagnosticNode_Severity_name = map[int32]string{
	0: "UNDEF",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}
var DiagnosticNode_Severity_value = map[string]int32{
	"UNDEF":   0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x DiagnosticNode_Severity) String() string {
	return proto.EnumName(DiagnosticNode_Severity_name, int32(x))
}
func (DiagnosticNode_Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{2, 0}
}

type FileNode struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FileNodeType         string            `protobuf:"bytes,2,opt,name=fileNodeType,proto3" json:"fileNodeType,omitempty"`
	FileType             FileNode_Type     `protobuf:"varint,3,opt,name=fileType,proto3,enum=service.FileNode_Type" json:"fileType,omitempty"`
	Path                 string            `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Name                 string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Hash                 string            `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	Broken               bool              `protobuf:"varint,7,opt,name=broken,proto3" json:"broken,omitempty"`
	DerivedFrom          []*FileNode       `protobuf:"bytes,8,rep,name=derivedFrom,proto3" json:"derivedFrom,omitempty"`
	AdditionalInfo       []*InfoNode       `protobuf:"bytes,9,rep,name=additionalInfo,proto3" json:"additionalInfo,omitempty"`
	DiagnosticInfo       []*DiagnosticNode `protobuf:"bytes,10,rep,name=diagnosticInfo,proto3" json:"diagnosticInfo,omitempty"`
	Dependencies         []*FileNode       `protobuf:"bytes,11,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FileNode) Reset()         { *m = FileNode{} }
func (m *FileNode) String() string { return proto.CompactTextString(m) }
func (*FileNode) ProtoMessage()    {}
func (*FileNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{0}
}
func (m *FileNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileNode.Unmarshal(m, b)
}
func (m *FileNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileNode.Marshal(b, m, deterministic)
}
func (dst *FileNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileNode.Merge(dst, src)
}
func (m *FileNode) XXX_Size() int {
	return xxx_messageInfo_FileNode.Size(m)
}
func (m *FileNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FileNode.DiscardUnknown(m)
}

var xxx_messageInfo_FileNode proto.InternalMessageInfo

func (m *FileNode) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FileNode) GetFileNodeType() string {
	if m != nil {
		return m.FileNodeType
	}
	return ""
}

func (m *FileNode) GetFileType() FileNode_Type {
	if m != nil {
		return m.FileType
	}
	return FileNode_UNDEF
}

func (m *FileNode) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileNode) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *FileNode) GetBroken() bool {
	if m != nil {
		return m.Broken
	}
	return false
}

func (m *FileNode) GetDerivedFrom() []*FileNode {
	if m != nil {
		return m.DerivedFrom
	}
	return nil
}

func (m *FileNode) GetAdditionalInfo() []*InfoNode {
	if m != nil {
		return m.AdditionalInfo
	}
	return nil
}

func (m *FileNode) GetDiagnosticInfo() []*DiagnosticNode {
	if m != nil {
		return m.DiagnosticInfo
	}
	return nil
}

func (m *FileNode) GetDependencies() []*FileNode {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type InfoNode struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	InfoNodeType         string               `protobuf:"bytes,2,opt,name=infoNodeType,proto3" json:"infoNodeType,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ConfidenceScore      float64              `protobuf:"fixed64,4,opt,name=confidenceScore,proto3" json:"confidenceScore,omitempty"`
	Analyzer             []*Analyzer          `protobuf:"bytes,5,rep,name=analyzer,proto3" json:"analyzer,omitempty"`
	DataNodes            []*InfoNode_DataNode `protobuf:"bytes,6,rep,name=dataNodes,proto3" json:"dataNodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InfoNode) Reset()         { *m = InfoNode{} }
func (m *InfoNode) String() string { return proto.CompactTextString(m) }
func (*InfoNode) ProtoMessage()    {}
func (*InfoNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{1}
}
func (m *InfoNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoNode.Unmarshal(m, b)
}
func (m *InfoNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoNode.Marshal(b, m, deterministic)
}
func (dst *InfoNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoNode.Merge(dst, src)
}
func (m *InfoNode) XXX_Size() int {
	return xxx_messageInfo_InfoNode.Size(m)
}
func (m *InfoNode) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoNode.DiscardUnknown(m)
}

var xxx_messageInfo_InfoNode proto.InternalMessageInfo

func (m *InfoNode) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *InfoNode) GetInfoNodeType() string {
	if m != nil {
		return m.InfoNodeType
	}
	return ""
}

func (m *InfoNode) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InfoNode) GetConfidenceScore() float64 {
	if m != nil {
		return m.ConfidenceScore
	}
	return 0
}

func (m *InfoNode) GetAnalyzer() []*Analyzer {
	if m != nil {
		return m.Analyzer
	}
	return nil
}

func (m *InfoNode) GetDataNodes() []*InfoNode_DataNode {
	if m != nil {
		return m.DataNodes
	}
	return nil
}

type InfoNode_DataNode struct {
	DataNodeType         string   `protobuf:"bytes,1,opt,name=dataNodeType,proto3" json:"dataNodeType,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoNode_DataNode) Reset()         { *m = InfoNode_DataNode{} }
func (m *InfoNode_DataNode) String() string { return proto.CompactTextString(m) }
func (*InfoNode_DataNode) ProtoMessage()    {}
func (*InfoNode_DataNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{1, 0}
}
func (m *InfoNode_DataNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoNode_DataNode.Unmarshal(m, b)
}
func (m *InfoNode_DataNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoNode_DataNode.Marshal(b, m, deterministic)
}
func (dst *InfoNode_DataNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoNode_DataNode.Merge(dst, src)
}
func (m *InfoNode_DataNode) XXX_Size() int {
	return xxx_messageInfo_InfoNode_DataNode.Size(m)
}
func (m *InfoNode_DataNode) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoNode_DataNode.DiscardUnknown(m)
}

var xxx_messageInfo_InfoNode_DataNode proto.InternalMessageInfo

func (m *InfoNode_DataNode) GetDataNodeType() string {
	if m != nil {
		return m.DataNodeType
	}
	return ""
}

func (m *InfoNode_DataNode) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InfoNode_DataNode) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type DiagnosticNode struct {
	Uid                  string                  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DiagnosticNodeType   string                  `protobuf:"bytes,2,opt,name=diagnosticNodeType,proto3" json:"diagnosticNodeType,omitempty"`
	Severity             DiagnosticNode_Severity `protobuf:"varint,3,opt,name=severity,proto3,enum=service.DiagnosticNode_Severity" json:"severity,omitempty"`
	Message              string                  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Analyzer             []*Analyzer             `protobuf:"bytes,5,rep,name=analyzer,proto3" json:"analyzer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DiagnosticNode) Reset()         { *m = DiagnosticNode{} }
func (m *DiagnosticNode) String() string { return proto.CompactTextString(m) }
func (*DiagnosticNode) ProtoMessage()    {}
func (*DiagnosticNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{2}
}
func (m *DiagnosticNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiagnosticNode.Unmarshal(m, b)
}
func (m *DiagnosticNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiagnosticNode.Marshal(b, m, deterministic)
}
func (dst *DiagnosticNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiagnosticNode.Merge(dst, src)
}
func (m *DiagnosticNode) XXX_Size() int {
	return xxx_messageInfo_DiagnosticNode.Size(m)
}
func (m *DiagnosticNode) XXX_DiscardUnknown() {
	xxx_messageInfo_DiagnosticNode.DiscardUnknown(m)
}

var xxx_messageInfo_DiagnosticNode proto.InternalMessageInfo

func (m *DiagnosticNode) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DiagnosticNode) GetDiagnosticNodeType() string {
	if m != nil {
		return m.DiagnosticNodeType
	}
	return ""
}

func (m *DiagnosticNode) GetSeverity() DiagnosticNode_Severity {
	if m != nil {
		return m.Severity
	}
	return DiagnosticNode_UNDEF
}

func (m *DiagnosticNode) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DiagnosticNode) GetAnalyzer() []*Analyzer {
	if m != nil {
		return m.Analyzer
	}
	return nil
}

type Analyzer struct {
	Uid                  string              `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AnalyzerNodeType     string              `protobuf:"bytes,3,opt,name=analyzerNodeType,proto3" json:"analyzerNodeType,omitempty"`
	TrustLevel           int64               `protobuf:"varint,4,opt,name=trustLevel,proto3" json:"trustLevel,omitempty"`
	PathSub              []*PathSubstitution `protobuf:"bytes,5,rep,name=pathSub,proto3" json:"pathSub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Analyzer) Reset()         { *m = Analyzer{} }
func (m *Analyzer) String() string { return proto.CompactTextString(m) }
func (*Analyzer) ProtoMessage()    {}
func (*Analyzer) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{3}
}
func (m *Analyzer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Analyzer.Unmarshal(m, b)
}
func (m *Analyzer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Analyzer.Marshal(b, m, deterministic)
}
func (dst *Analyzer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Analyzer.Merge(dst, src)
}
func (m *Analyzer) XXX_Size() int {
	return xxx_messageInfo_Analyzer.Size(m)
}
func (m *Analyzer) XXX_DiscardUnknown() {
	xxx_messageInfo_Analyzer.DiscardUnknown(m)
}

var xxx_messageInfo_Analyzer proto.InternalMessageInfo

func (m *Analyzer) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Analyzer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Analyzer) GetAnalyzerNodeType() string {
	if m != nil {
		return m.AnalyzerNodeType
	}
	return ""
}

func (m *Analyzer) GetTrustLevel() int64 {
	if m != nil {
		return m.TrustLevel
	}
	return 0
}

func (m *Analyzer) GetPathSub() []*PathSubstitution {
	if m != nil {
		return m.PathSub
	}
	return nil
}

type PathSubstitution struct {
	Old                  string   `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`
	New                  string   `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathSubstitution) Reset()         { *m = PathSubstitution{} }
func (m *PathSubstitution) String() string { return proto.CompactTextString(m) }
func (*PathSubstitution) ProtoMessage()    {}
func (*PathSubstitution) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{4}
}
func (m *PathSubstitution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathSubstitution.Unmarshal(m, b)
}
func (m *PathSubstitution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathSubstitution.Marshal(b, m, deterministic)
}
func (dst *PathSubstitution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathSubstitution.Merge(dst, src)
}
func (m *PathSubstitution) XXX_Size() int {
	return xxx_messageInfo_PathSubstitution.Size(m)
}
func (m *PathSubstitution) XXX_DiscardUnknown() {
	xxx_messageInfo_PathSubstitution.DiscardUnknown(m)
}

var xxx_messageInfo_PathSubstitution proto.InternalMessageInfo

func (m *PathSubstitution) GetOld() string {
	if m != nil {
		return m.Old
	}
	return ""
}

func (m *PathSubstitution) GetNew() string {
	if m != nil {
		return m.New
	}
	return ""
}

type PackageNode struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PackageNodeType      string            `protobuf:"bytes,4,opt,name=packageNodeType,proto3" json:"packageNodeType,omitempty"`
	Targets              []*FileNode       `protobuf:"bytes,5,rep,name=targets,proto3" json:"targets,omitempty"`
	AdditionalInfo       []*InfoNode       `protobuf:"bytes,6,rep,name=additionalInfo,proto3" json:"additionalInfo,omitempty"`
	BuildConfig          string            `protobuf:"bytes,7,opt,name=buildConfig,proto3" json:"buildConfig,omitempty"`
	DiagnosticInfo       []*DiagnosticNode `protobuf:"bytes,8,rep,name=diagnosticInfo,proto3" json:"diagnosticInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PackageNode) Reset()         { *m = PackageNode{} }
func (m *PackageNode) String() string { return proto.CompactTextString(m) }
func (*PackageNode) ProtoMessage()    {}
func (*PackageNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{5}
}
func (m *PackageNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageNode.Unmarshal(m, b)
}
func (m *PackageNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageNode.Marshal(b, m, deterministic)
}
func (dst *PackageNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageNode.Merge(dst, src)
}
func (m *PackageNode) XXX_Size() int {
	return xxx_messageInfo_PackageNode.Size(m)
}
func (m *PackageNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageNode.DiscardUnknown(m)
}

var xxx_messageInfo_PackageNode proto.InternalMessageInfo

func (m *PackageNode) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PackageNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PackageNode) GetPackageNodeType() string {
	if m != nil {
		return m.PackageNodeType
	}
	return ""
}

func (m *PackageNode) GetTargets() []*FileNode {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *PackageNode) GetAdditionalInfo() []*InfoNode {
	if m != nil {
		return m.AdditionalInfo
	}
	return nil
}

func (m *PackageNode) GetBuildConfig() string {
	if m != nil {
		return m.BuildConfig
	}
	return ""
}

func (m *PackageNode) GetDiagnosticInfo() []*DiagnosticNode {
	if m != nil {
		return m.DiagnosticInfo
	}
	return nil
}

type ProjectNode struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectNodeType      string         `protobuf:"bytes,3,opt,name=projectNodeType,proto3" json:"projectNodeType,omitempty"`
	Packages             []*PackageNode `protobuf:"bytes,4,rep,name=packages,proto3" json:"packages,omitempty"`
	AdditionalInfo       []*InfoNode    `protobuf:"bytes,5,rep,name=additionalInfo,proto3" json:"additionalInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProjectNode) Reset()         { *m = ProjectNode{} }
func (m *ProjectNode) String() string { return proto.CompactTextString(m) }
func (*ProjectNode) ProtoMessage()    {}
func (*ProjectNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{6}
}
func (m *ProjectNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectNode.Unmarshal(m, b)
}
func (m *ProjectNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectNode.Marshal(b, m, deterministic)
}
func (dst *ProjectNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectNode.Merge(dst, src)
}
func (m *ProjectNode) XXX_Size() int {
	return xxx_messageInfo_ProjectNode.Size(m)
}
func (m *ProjectNode) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectNode.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectNode proto.InternalMessageInfo

func (m *ProjectNode) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ProjectNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProjectNode) GetProjectNodeType() string {
	if m != nil {
		return m.ProjectNodeType
	}
	return ""
}

func (m *ProjectNode) GetPackages() []*PackageNode {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *ProjectNode) GetAdditionalInfo() []*InfoNode {
	if m != nil {
		return m.AdditionalInfo
	}
	return nil
}

type Event struct {
	Class                EventClass `protobuf:"varint,1,opt,name=class,proto3,enum=service.EventClass" json:"class,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{7}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetClass() EventClass {
	if m != nil {
		return m.Class
	}
	return EventClass_ALL
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type QmstrStateNode struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	QmstrStateNodeType   string   `protobuf:"bytes,2,opt,name=qmstrStateNodeType,proto3" json:"qmstrStateNodeType,omitempty"`
	Phase                Phase    `protobuf:"varint,3,opt,name=phase,proto3,enum=service.Phase" json:"phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QmstrStateNode) Reset()         { *m = QmstrStateNode{} }
func (m *QmstrStateNode) String() string { return proto.CompactTextString(m) }
func (*QmstrStateNode) ProtoMessage()    {}
func (*QmstrStateNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamodel_09fee586c8ee7519, []int{8}
}
func (m *QmstrStateNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QmstrStateNode.Unmarshal(m, b)
}
func (m *QmstrStateNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QmstrStateNode.Marshal(b, m, deterministic)
}
func (dst *QmstrStateNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QmstrStateNode.Merge(dst, src)
}
func (m *QmstrStateNode) XXX_Size() int {
	return xxx_messageInfo_QmstrStateNode.Size(m)
}
func (m *QmstrStateNode) XXX_DiscardUnknown() {
	xxx_messageInfo_QmstrStateNode.DiscardUnknown(m)
}

var xxx_messageInfo_QmstrStateNode proto.InternalMessageInfo

func (m *QmstrStateNode) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *QmstrStateNode) GetQmstrStateNodeType() string {
	if m != nil {
		return m.QmstrStateNodeType
	}
	return ""
}

func (m *QmstrStateNode) GetPhase() Phase {
	if m != nil {
		return m.Phase
	}
	return Phase_INIT
}

func init() {
	proto.RegisterType((*FileNode)(nil), "service.FileNode")
	proto.RegisterType((*InfoNode)(nil), "service.InfoNode")
	proto.RegisterType((*InfoNode_DataNode)(nil), "service.InfoNode.DataNode")
	proto.RegisterType((*DiagnosticNode)(nil), "service.DiagnosticNode")
	proto.RegisterType((*Analyzer)(nil), "service.Analyzer")
	proto.RegisterType((*PathSubstitution)(nil), "service.PathSubstitution")
	proto.RegisterType((*PackageNode)(nil), "service.PackageNode")
	proto.RegisterType((*ProjectNode)(nil), "service.ProjectNode")
	proto.RegisterType((*Event)(nil), "service.Event")
	proto.RegisterType((*QmstrStateNode)(nil), "service.QmstrStateNode")
	proto.RegisterEnum("service.EventClass", EventClass_name, EventClass_value)
	proto.RegisterEnum("service.Phase", Phase_name, Phase_value)
	proto.RegisterEnum("service.ExceptionType", ExceptionType_name, ExceptionType_value)
	proto.RegisterEnum("service.FileNode_Type", FileNode_Type_name, FileNode_Type_value)
	proto.RegisterEnum("service.DiagnosticNode_Severity", DiagnosticNode_Severity_name, DiagnosticNode_Severity_value)
}

func init() { proto.RegisterFile("datamodel.proto", fileDescriptor_datamodel_09fee586c8ee7519) }

var fileDescriptor_datamodel_09fee586c8ee7519 = []byte{
	// 914 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0xed, 0xfc, 0x38, 0x27, 0x25, 0x35, 0x03, 0x2a, 0x66, 0x2f, 0x50, 0x64, 0x21, 0x11,
	0x0a, 0x58, 0xa8, 0x15, 0x7f, 0x02, 0x09, 0xbc, 0x8d, 0xb3, 0x58, 0xca, 0x26, 0x61, 0x92, 0x82,
	0xb8, 0x9c, 0xda, 0xd3, 0xd4, 0xac, 0x63, 0x07, 0x7b, 0x12, 0xb6, 0x3c, 0x06, 0x4f, 0xc1, 0x2d,
	0x4f, 0xc0, 0x3d, 0xef, 0xc2, 0x3b, 0xac, 0x66, 0xec, 0x71, 0x9c, 0x3f, 0xa9, 0xbd, 0x1b, 0x7f,
	0xdf, 0x77, 0x26, 0xe7, 0x3b, 0xc7, 0xe7, 0x38, 0x70, 0x16, 0x10, 0x46, 0x16, 0x49, 0x40, 0x23,
	0x7b, 0x99, 0x26, 0x2c, 0x41, 0xcd, 0x8c, 0xa6, 0xeb, 0xd0, 0xa7, 0xd6, 0xff, 0x1a, 0xe8, 0x83,
	0x30, 0xa2, 0xa3, 0x24, 0xa0, 0xc8, 0x00, 0x6d, 0x15, 0x06, 0xa6, 0xd2, 0x55, 0x7a, 0x2d, 0xcc,
	0x8f, 0xc8, 0x82, 0xd3, 0xbb, 0x82, 0x9d, 0x3d, 0x2c, 0xa9, 0xa9, 0x0a, 0x6a, 0x0b, 0x43, 0x97,
	0xa0, 0xf3, 0x67, 0xc1, 0x6b, 0x5d, 0xa5, 0xd7, 0xb9, 0x3c, 0xb7, 0x8b, 0xeb, 0x6d, 0x79, 0xb5,
	0xcd, 0x59, 0x5c, 0xea, 0x10, 0x82, 0xda, 0x92, 0xb0, 0x7b, 0xb3, 0x26, 0xee, 0x13, 0x67, 0x8e,
	0xc5, 0x64, 0x41, 0xcd, 0x7a, 0x8e, 0xf1, 0x33, 0xc7, 0xee, 0x49, 0x76, 0x6f, 0x36, 0x72, 0x8c,
	0x9f, 0xd1, 0x39, 0x34, 0x6e, 0xd3, 0xe4, 0x15, 0x8d, 0xcd, 0x66, 0x57, 0xe9, 0xe9, 0xb8, 0x78,
	0x42, 0x57, 0xd0, 0x0e, 0x68, 0x1a, 0xae, 0x69, 0x30, 0x48, 0x93, 0x85, 0xa9, 0x77, 0xb5, 0x5e,
	0xfb, 0xf2, 0xed, 0xbd, 0x54, 0x70, 0x55, 0x85, 0xbe, 0x81, 0x0e, 0x09, 0x82, 0x90, 0x85, 0x49,
	0x4c, 0x22, 0x2f, 0xbe, 0x4b, 0xcc, 0xd6, 0x4e, 0x1c, 0x07, 0x45, 0xdc, 0x8e, 0x10, 0x7d, 0x0f,
	0x9d, 0x20, 0x24, 0xf3, 0x38, 0xc9, 0x58, 0xe8, 0x8b, 0x50, 0x10, 0xa1, 0xef, 0x95, 0xa1, 0xfd,
	0x92, 0xce, 0x2f, 0xd8, 0x96, 0xa3, 0x2f, 0xe0, 0x34, 0xa0, 0x4b, 0x1a, 0x07, 0x34, 0xf6, 0x43,
	0x9a, 0x99, 0xed, 0x63, 0x19, 0x6f, 0xc9, 0xac, 0x6f, 0xa1, 0x26, 0x6a, 0xd8, 0x82, 0xfa, 0xcd,
	0xa8, 0xef, 0x0e, 0x8c, 0x13, 0x04, 0xd0, 0x98, 0x8e, 0x6f, 0xf0, 0xb5, 0x6b, 0x28, 0xc8, 0x80,
	0x53, 0x6f, 0x34, 0x73, 0xf1, 0x4b, 0xb7, 0xef, 0x39, 0x33, 0xd7, 0x50, 0x39, 0x3b, 0x73, 0xf0,
	0x0b, 0x77, 0x66, 0x68, 0xd6, 0xbf, 0x2a, 0xe8, 0xd2, 0xd1, 0xe1, 0x7e, 0x87, 0x05, 0x5b, 0xed,
	0x77, 0x15, 0xe3, 0x3d, 0x61, 0xb2, 0xd7, 0x2d, 0x2c, 0xce, 0xa8, 0x07, 0x67, 0x7e, 0x12, 0xdf,
	0x85, 0x3c, 0x47, 0x3a, 0xf5, 0x93, 0x94, 0x8a, 0xd6, 0x2a, 0x78, 0x17, 0x46, 0x9f, 0x81, 0x4e,
	0x62, 0x12, 0x3d, 0xfc, 0x49, 0x53, 0xb3, 0xbe, 0x63, 0xd8, 0x29, 0x08, 0x5c, 0x4a, 0xd0, 0xd7,
	0xd0, 0xe2, 0xef, 0x2e, 0xff, 0xf1, 0xcc, 0x6c, 0x08, 0xfd, 0xb3, 0xbd, 0xd6, 0xd8, 0xfd, 0x42,
	0x82, 0x37, 0xe2, 0x67, 0x3f, 0x83, 0x2e, 0x61, 0x6e, 0x4b, 0x12, 0xc2, 0x56, 0xee, 0x78, 0x0b,
	0x2b, 0x6d, 0xa9, 0x15, 0x5b, 0x08, 0x6a, 0x5c, 0x23, 0xad, 0xf2, 0xb3, 0xf5, 0x97, 0x0a, 0x9d,
	0xed, 0xc6, 0x1e, 0xa8, 0xa3, 0x0d, 0x28, 0xd8, 0xd2, 0x54, 0xaa, 0x79, 0x80, 0x41, 0xdf, 0x81,
	0x9e, 0xd1, 0x35, 0x4d, 0x43, 0xf6, 0x50, 0xcc, 0x50, 0xf7, 0xc8, 0x5b, 0x64, 0x4f, 0x0b, 0x1d,
	0x2e, 0x23, 0x90, 0x09, 0xcd, 0x05, 0xcd, 0x32, 0x32, 0xa7, 0xc5, 0x40, 0xc9, 0xc7, 0x27, 0x56,
	0xdb, 0xfa, 0x0a, 0x74, 0x79, 0x7d, 0xf5, 0xf5, 0xd2, 0xa1, 0xe6, 0x8d, 0x06, 0x63, 0x43, 0x41,
	0x6d, 0x68, 0xfe, 0xe2, 0xe0, 0x91, 0x37, 0x7a, 0x61, 0xa8, 0x5c, 0xe1, 0x62, 0x3c, 0xc6, 0x86,
	0x66, 0xfd, 0xa3, 0x80, 0x2e, 0xef, 0x3b, 0x50, 0x0e, 0x39, 0xda, 0x6a, 0x65, 0xb4, 0x2f, 0xc0,
	0x90, 0xbf, 0x5b, 0x16, 0x28, 0xaf, 0xf3, 0x1e, 0x8e, 0x3e, 0x00, 0x60, 0xe9, 0x2a, 0x63, 0x43,
	0xba, 0xa6, 0x91, 0xf0, 0xa8, 0xe1, 0x0a, 0x82, 0xae, 0xa0, 0xc9, 0x57, 0xc8, 0x74, 0x75, 0x5b,
	0xb8, 0x7c, 0xbf, 0x74, 0x39, 0xc9, 0xf1, 0x8c, 0x85, 0x6c, 0xc5, 0x87, 0x17, 0x4b, 0xa5, 0xf5,
	0x25, 0x18, 0xbb, 0x24, 0x4f, 0x3d, 0x89, 0xca, 0xd4, 0x93, 0x28, 0xe0, 0x48, 0x4c, 0xff, 0x28,
	0x32, 0xe7, 0x47, 0xeb, 0x6f, 0x15, 0xda, 0x13, 0xe2, 0xbf, 0x22, 0xf3, 0x63, 0x5b, 0xf3, 0x90,
	0xdd, 0x1e, 0x9c, 0x2d, 0x37, 0x41, 0xc2, 0x6d, 0xde, 0xab, 0x5d, 0x18, 0x7d, 0x02, 0x4d, 0x46,
	0xd2, 0x39, 0x65, 0xd9, 0x5e, 0xcb, 0xca, 0x8d, 0x20, 0x15, 0x07, 0xf6, 0x57, 0xe3, 0xb1, 0xfb,
	0xab, 0x0b, 0xed, 0xdb, 0x55, 0x18, 0x05, 0xd7, 0x7c, 0x42, 0xe7, 0x62, 0x99, 0xb6, 0x70, 0x15,
	0x3a, 0xb0, 0xe1, 0xf4, 0x27, 0x6d, 0x38, 0xeb, 0x3f, 0x05, 0xda, 0x93, 0x34, 0xf9, 0x8d, 0xfa,
	0xec, 0x89, 0xa5, 0xda, 0x04, 0x55, 0x5e, 0x8c, 0x5d, 0x18, 0x7d, 0x0e, 0x7a, 0x51, 0xbd, 0xcc,
	0xac, 0x89, 0xd4, 0xde, 0xad, 0x34, 0xbe, 0x2c, 0x2b, 0x2e, 0x55, 0x07, 0xea, 0x55, 0x7f, 0x64,
	0xbd, 0xac, 0x21, 0xd4, 0xdd, 0x35, 0x8d, 0x19, 0xfa, 0x18, 0xea, 0x7e, 0x44, 0xb2, 0x4c, 0xf8,
	0xe8, 0x5c, 0xbe, 0x53, 0x86, 0x0a, 0xfa, 0x9a, 0x53, 0x38, 0x57, 0x54, 0x27, 0x53, 0xdd, 0x9a,
	0x4c, 0xeb, 0x35, 0x74, 0x7e, 0x5a, 0x64, 0x2c, 0x9d, 0x32, 0xc2, 0xe8, 0xf1, 0x2d, 0xf2, 0xfb,
	0x96, 0xa6, 0xba, 0x45, 0xf6, 0x19, 0xf4, 0x21, 0xd4, 0x97, 0xf7, 0x24, 0x93, 0x9f, 0xe1, 0xce,
	0xa6, 0x16, 0x1c, 0xc5, 0x39, 0x79, 0xf1, 0x29, 0xc0, 0x26, 0x51, 0xd4, 0x04, 0xcd, 0x19, 0x0e,
	0x8d, 0x13, 0x3e, 0xcd, 0x93, 0x1f, 0x9d, 0x29, 0xff, 0x84, 0x00, 0x34, 0x5e, 0x8e, 0xfb, 0x37,
	0x43, 0xd7, 0x50, 0x2f, 0x7e, 0x80, 0xba, 0x88, 0xce, 0x97, 0x80, 0x37, 0xcb, 0x95, 0xcf, 0x6f,
	0xbc, 0x61, 0xdf, 0x50, 0xd0, 0x29, 0xe8, 0xce, 0xc8, 0x19, 0xfe, 0x3a, 0xf5, 0xa6, 0xf9, 0x87,
	0x06, 0xbb, 0x93, 0x31, 0x9e, 0x19, 0x1a, 0x97, 0x0f, 0x1c, 0x6f, 0x68, 0xd4, 0x2e, 0x3e, 0x82,
	0xb7, 0xdc, 0xd7, 0x3e, 0x5d, 0xf2, 0x52, 0xca, 0x0f, 0x57, 0xbe, 0x37, 0x4e, 0xaa, 0xfb, 0x44,
	0x79, 0x6e, 0xc2, 0x79, 0x92, 0xce, 0x6d, 0x61, 0xcc, 0x9e, 0xa7, 0x4b, 0x5f, 0xe6, 0x7f, 0xdb,
	0x10, 0xff, 0x5a, 0xae, 0xde, 0x04, 0x00, 0x00, 0xff, 0xff, 0x66, 0x5e, 0x01, 0x12, 0xc8, 0x08,
	0x00, 0x00,
}
