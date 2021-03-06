// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bc.proto

/*
Package bc is a generated protocol buffer package.

It is generated from these files:
	bc.proto

It has these top-level messages:
	Hash
	Program
	AssetID
	AssetAmount
	AssetDefinition
	ValueSource
	ValueDestination
	BlockHeader
	TxHeader
	TxVerifyResult
	TransactionStatus
	Mux
	Nonce
	Coinbase
	Output
	Retirement
	Issuance
	Spend
*/
package bc

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

type Hash struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hash) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *Hash) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *Hash) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *Hash) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type Program struct {
	VmVersion uint64 `protobuf:"varint,1,opt,name=vm_version,json=vmVersion" json:"vm_version,omitempty"`
	Code      []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *Program) Reset()                    { *m = Program{} }
func (m *Program) String() string            { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()               {}
func (*Program) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Program) GetVmVersion() uint64 {
	if m != nil {
		return m.VmVersion
	}
	return 0
}

func (m *Program) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// This message type duplicates Hash, above. One alternative is to
// embed a Hash inside an AssetID. But it's useful for AssetID to be
// plain old data (without pointers). Another alternative is use Hash
// in any protobuf types where an AssetID is called for, but it's
// preferable to have type safety.
type AssetID struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *AssetID) Reset()                    { *m = AssetID{} }
func (m *AssetID) String() string            { return proto.CompactTextString(m) }
func (*AssetID) ProtoMessage()               {}
func (*AssetID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AssetID) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *AssetID) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *AssetID) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *AssetID) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type AssetAmount struct {
	AssetId *AssetID `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	Amount  uint64   `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *AssetAmount) Reset()                    { *m = AssetAmount{} }
func (m *AssetAmount) String() string            { return proto.CompactTextString(m) }
func (*AssetAmount) ProtoMessage()               {}
func (*AssetAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AssetAmount) GetAssetId() *AssetID {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *AssetAmount) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AssetDefinition struct {
	IssuanceProgram *Program `protobuf:"bytes,1,opt,name=issuance_program,json=issuanceProgram" json:"issuance_program,omitempty"`
	Data            *Hash    `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *AssetDefinition) Reset()                    { *m = AssetDefinition{} }
func (m *AssetDefinition) String() string            { return proto.CompactTextString(m) }
func (*AssetDefinition) ProtoMessage()               {}
func (*AssetDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AssetDefinition) GetIssuanceProgram() *Program {
	if m != nil {
		return m.IssuanceProgram
	}
	return nil
}

func (m *AssetDefinition) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

type ValueSource struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueSource) Reset()                    { *m = ValueSource{} }
func (m *ValueSource) String() string            { return proto.CompactTextString(m) }
func (*ValueSource) ProtoMessage()               {}
func (*ValueSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ValueSource) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueSource) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueSource) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type ValueDestination struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueDestination) Reset()                    { *m = ValueDestination{} }
func (m *ValueDestination) String() string            { return proto.CompactTextString(m) }
func (*ValueDestination) ProtoMessage()               {}
func (*ValueDestination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ValueDestination) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueDestination) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueDestination) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type BlockHeader struct {
	Version               uint64             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Height                uint64             `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	PreviousBlockId       *Hash              `protobuf:"bytes,3,opt,name=previous_block_id,json=previousBlockId" json:"previous_block_id,omitempty"`
	Timestamp             uint64             `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	TransactionsRoot      *Hash              `protobuf:"bytes,5,opt,name=transactions_root,json=transactionsRoot" json:"transactions_root,omitempty"`
	TransactionStatusHash *Hash              `protobuf:"bytes,6,opt,name=transaction_status_hash,json=transactionStatusHash" json:"transaction_status_hash,omitempty"`
	Nonce                 uint64             `protobuf:"varint,7,opt,name=nonce" json:"nonce,omitempty"`
	Bits                  uint64             `protobuf:"varint,8,opt,name=bits" json:"bits,omitempty"`
	TransactionStatus     *TransactionStatus `protobuf:"bytes,9,opt,name=transaction_status,json=transactionStatus" json:"transaction_status,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BlockHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockId() *Hash {
	if m != nil {
		return m.PreviousBlockId
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetTransactionsRoot() *Hash {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *BlockHeader) GetTransactionStatusHash() *Hash {
	if m != nil {
		return m.TransactionStatusHash
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetBits() uint64 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *BlockHeader) GetTransactionStatus() *TransactionStatus {
	if m != nil {
		return m.TransactionStatus
	}
	return nil
}

type TxHeader struct {
	Version        uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SerializedSize uint64  `protobuf:"varint,2,opt,name=serialized_size,json=serializedSize" json:"serialized_size,omitempty"`
	TimeRange      uint64  `protobuf:"varint,3,opt,name=time_range,json=timeRange" json:"time_range,omitempty"`
	ResultIds      []*Hash `protobuf:"bytes,4,rep,name=result_ids,json=resultIds" json:"result_ids,omitempty"`
	ExtHash        *Hash   `protobuf:"bytes,5,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
}

func (m *TxHeader) Reset()                    { *m = TxHeader{} }
func (m *TxHeader) String() string            { return proto.CompactTextString(m) }
func (*TxHeader) ProtoMessage()               {}
func (*TxHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TxHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TxHeader) GetSerializedSize() uint64 {
	if m != nil {
		return m.SerializedSize
	}
	return 0
}

func (m *TxHeader) GetTimeRange() uint64 {
	if m != nil {
		return m.TimeRange
	}
	return 0
}

func (m *TxHeader) GetResultIds() []*Hash {
	if m != nil {
		return m.ResultIds
	}
	return nil
}

func (m *TxHeader) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

type TxVerifyResult struct {
	StatusFail bool `protobuf:"varint,1,opt,name=status_fail,json=statusFail" json:"status_fail,omitempty"`
}

func (m *TxVerifyResult) Reset()                    { *m = TxVerifyResult{} }
func (m *TxVerifyResult) String() string            { return proto.CompactTextString(m) }
func (*TxVerifyResult) ProtoMessage()               {}
func (*TxVerifyResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TxVerifyResult) GetStatusFail() bool {
	if m != nil {
		return m.StatusFail
	}
	return false
}

type TransactionStatus struct {
	Version      uint64            `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	VerifyStatus []*TxVerifyResult `protobuf:"bytes,2,rep,name=verify_status,json=verifyStatus" json:"verify_status,omitempty"`
}

func (m *TransactionStatus) Reset()                    { *m = TransactionStatus{} }
func (m *TransactionStatus) String() string            { return proto.CompactTextString(m) }
func (*TransactionStatus) ProtoMessage()               {}
func (*TransactionStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TransactionStatus) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TransactionStatus) GetVerifyStatus() []*TxVerifyResult {
	if m != nil {
		return m.VerifyStatus
	}
	return nil
}

type Mux struct {
	Sources             []*ValueSource      `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
	Program             *Program            `protobuf:"bytes,2,opt,name=program" json:"program,omitempty"`
	ExtHash             *Hash               `protobuf:"bytes,3,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessDestinations []*ValueDestination `protobuf:"bytes,4,rep,name=witness_destinations,json=witnessDestinations" json:"witness_destinations,omitempty"`
	WitnessArguments    [][]byte            `protobuf:"bytes,5,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
}

func (m *Mux) Reset()                    { *m = Mux{} }
func (m *Mux) String() string            { return proto.CompactTextString(m) }
func (*Mux) ProtoMessage()               {}
func (*Mux) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Mux) GetSources() []*ValueSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Mux) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Mux) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Mux) GetWitnessDestinations() []*ValueDestination {
	if m != nil {
		return m.WitnessDestinations
	}
	return nil
}

func (m *Mux) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

type Nonce struct {
	Program           *Program `protobuf:"bytes,1,opt,name=program" json:"program,omitempty"`
	ExtHash           *Hash    `protobuf:"bytes,2,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessArguments  [][]byte `protobuf:"bytes,3,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	WitnessAnchoredId *Hash    `protobuf:"bytes,4,opt,name=witness_anchored_id,json=witnessAnchoredId" json:"witness_anchored_id,omitempty"`
}

func (m *Nonce) Reset()                    { *m = Nonce{} }
func (m *Nonce) String() string            { return proto.CompactTextString(m) }
func (*Nonce) ProtoMessage()               {}
func (*Nonce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Nonce) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Nonce) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Nonce) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Nonce) GetWitnessAnchoredId() *Hash {
	if m != nil {
		return m.WitnessAnchoredId
	}
	return nil
}

type Coinbase struct {
	WitnessDestination *ValueDestination `protobuf:"bytes,1,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	Arbitrary          []byte            `protobuf:"bytes,2,opt,name=arbitrary,proto3" json:"arbitrary,omitempty"`
}

func (m *Coinbase) Reset()                    { *m = Coinbase{} }
func (m *Coinbase) String() string            { return proto.CompactTextString(m) }
func (*Coinbase) ProtoMessage()               {}
func (*Coinbase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Coinbase) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Coinbase) GetArbitrary() []byte {
	if m != nil {
		return m.Arbitrary
	}
	return nil
}

type Output struct {
	Source         *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ControlProgram *Program     `protobuf:"bytes,2,opt,name=control_program,json=controlProgram" json:"control_program,omitempty"`
	ExtHash        *Hash        `protobuf:"bytes,3,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	Ordinal        uint64       `protobuf:"varint,4,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Output) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Output) GetControlProgram() *Program {
	if m != nil {
		return m.ControlProgram
	}
	return nil
}

func (m *Output) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Output) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Retirement struct {
	Source  *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ExtHash *Hash        `protobuf:"bytes,2,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	Ordinal uint64       `protobuf:"varint,3,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Retirement) Reset()                    { *m = Retirement{} }
func (m *Retirement) String() string            { return proto.CompactTextString(m) }
func (*Retirement) ProtoMessage()               {}
func (*Retirement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Retirement) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Retirement) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Retirement) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Issuance struct {
	AnchorId               *Hash             `protobuf:"bytes,1,opt,name=anchor_id,json=anchorId" json:"anchor_id,omitempty"`
	Value                  *AssetAmount      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	ExtHash                *Hash             `protobuf:"bytes,3,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessDestination     *ValueDestination `protobuf:"bytes,4,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessAssetDefinition *AssetDefinition  `protobuf:"bytes,5,opt,name=witness_asset_definition,json=witnessAssetDefinition" json:"witness_asset_definition,omitempty"`
	WitnessArguments       [][]byte          `protobuf:"bytes,6,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	WitnessAnchoredId      *Hash             `protobuf:"bytes,7,opt,name=witness_anchored_id,json=witnessAnchoredId" json:"witness_anchored_id,omitempty"`
	Ordinal                uint64            `protobuf:"varint,8,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Issuance) Reset()                    { *m = Issuance{} }
func (m *Issuance) String() string            { return proto.CompactTextString(m) }
func (*Issuance) ProtoMessage()               {}
func (*Issuance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Issuance) GetAnchorId() *Hash {
	if m != nil {
		return m.AnchorId
	}
	return nil
}

func (m *Issuance) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Issuance) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Issuance) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Issuance) GetWitnessAssetDefinition() *AssetDefinition {
	if m != nil {
		return m.WitnessAssetDefinition
	}
	return nil
}

func (m *Issuance) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Issuance) GetWitnessAnchoredId() *Hash {
	if m != nil {
		return m.WitnessAnchoredId
	}
	return nil
}

func (m *Issuance) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Spend struct {
	SpentOutputId      *Hash             `protobuf:"bytes,1,opt,name=spent_output_id,json=spentOutputId" json:"spent_output_id,omitempty"`
	ExtHash            *Hash             `protobuf:"bytes,2,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessDestination *ValueDestination `protobuf:"bytes,3,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessArguments   [][]byte          `protobuf:"bytes,4,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	WitnessAnchoredId  *Hash             `protobuf:"bytes,5,opt,name=witness_anchored_id,json=witnessAnchoredId" json:"witness_anchored_id,omitempty"`
	Ordinal            uint64            `protobuf:"varint,6,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Spend) Reset()                    { *m = Spend{} }
func (m *Spend) String() string            { return proto.CompactTextString(m) }
func (*Spend) ProtoMessage()               {}
func (*Spend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *Spend) GetSpentOutputId() *Hash {
	if m != nil {
		return m.SpentOutputId
	}
	return nil
}

func (m *Spend) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Spend) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Spend) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Spend) GetWitnessAnchoredId() *Hash {
	if m != nil {
		return m.WitnessAnchoredId
	}
	return nil
}

func (m *Spend) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func init() {
	proto.RegisterType((*Hash)(nil), "bc.Hash")
	proto.RegisterType((*Program)(nil), "bc.Program")
	proto.RegisterType((*AssetID)(nil), "bc.AssetID")
	proto.RegisterType((*AssetAmount)(nil), "bc.AssetAmount")
	proto.RegisterType((*AssetDefinition)(nil), "bc.AssetDefinition")
	proto.RegisterType((*ValueSource)(nil), "bc.ValueSource")
	proto.RegisterType((*ValueDestination)(nil), "bc.ValueDestination")
	proto.RegisterType((*BlockHeader)(nil), "bc.BlockHeader")
	proto.RegisterType((*TxHeader)(nil), "bc.TxHeader")
	proto.RegisterType((*TxVerifyResult)(nil), "bc.TxVerifyResult")
	proto.RegisterType((*TransactionStatus)(nil), "bc.TransactionStatus")
	proto.RegisterType((*Mux)(nil), "bc.Mux")
	proto.RegisterType((*Nonce)(nil), "bc.Nonce")
	proto.RegisterType((*Coinbase)(nil), "bc.Coinbase")
	proto.RegisterType((*Output)(nil), "bc.Output")
	proto.RegisterType((*Retirement)(nil), "bc.Retirement")
	proto.RegisterType((*Issuance)(nil), "bc.Issuance")
	proto.RegisterType((*Spend)(nil), "bc.Spend")
}

func init() { proto.RegisterFile("bc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 994 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x86, 0x44, 0x4a, 0xa2, 0x8f, 0x1c, 0xcb, 0x1e, 0x3b, 0xf9, 0x89, 0x20, 0x3f, 0x6a, 0xb0,
	0x70, 0x9d, 0xa2, 0x80, 0xe1, 0x4b, 0x7a, 0x59, 0x74, 0x51, 0xb7, 0x6e, 0x1a, 0x2d, 0xdc, 0x06,
	0x63, 0xc3, 0x5b, 0x62, 0x44, 0x8e, 0xad, 0x41, 0x25, 0x8e, 0x3a, 0x33, 0x54, 0x1d, 0x3f, 0x47,
	0x9f, 0x21, 0x0f, 0xd1, 0xc7, 0xea, 0xba, 0x8b, 0x62, 0x0e, 0x87, 0x12, 0x75, 0x71, 0x22, 0xa1,
	0xe8, 0x8e, 0xe7, 0x32, 0xe7, 0xf2, 0x9d, 0xef, 0x0c, 0x07, 0x82, 0x5e, 0x72, 0x34, 0x52, 0xd2,
	0x48, 0x52, 0xef, 0x25, 0xd1, 0x6b, 0xf0, 0xdf, 0x30, 0xdd, 0x27, 0x5b, 0x50, 0x1f, 0x1f, 0x87,
	0xb5, 0xfd, 0xda, 0xcb, 0x26, 0xad, 0x8f, 0x8f, 0x51, 0x3e, 0x09, 0xeb, 0x4e, 0x3e, 0x41, 0xf9,
	0x34, 0xf4, 0x9c, 0x7c, 0x8a, 0xf2, 0x59, 0xe8, 0x3b, 0xf9, 0x2c, 0xfa, 0x16, 0x5a, 0x6f, 0x95,
	0xbc, 0x53, 0x6c, 0x48, 0xfe, 0x0f, 0x30, 0x1e, 0xc6, 0x63, 0xae, 0xb4, 0x90, 0x19, 0x86, 0xf4,
	0xe9, 0xc6, 0x78, 0x78, 0x53, 0x28, 0x08, 0x01, 0x3f, 0x91, 0x29, 0xc7, 0xd8, 0x9b, 0x14, 0xbf,
	0xa3, 0x2e, 0xb4, 0xce, 0xb5, 0xe6, 0xa6, 0x7b, 0xf1, 0xaf, 0x0b, 0xb9, 0x84, 0x36, 0x86, 0x3a,
	0x1f, 0xca, 0x3c, 0x33, 0xe4, 0x33, 0x08, 0x98, 0x15, 0x63, 0x91, 0x62, 0xd0, 0xf6, 0x69, 0xfb,
	0xa8, 0x97, 0x1c, 0xb9, 0x6c, 0xb4, 0x85, 0xc6, 0x6e, 0x4a, 0x9e, 0x41, 0x93, 0xe1, 0x09, 0x4c,
	0xe5, 0x53, 0x27, 0x45, 0x77, 0xd0, 0x41, 0xdf, 0x0b, 0x7e, 0x2b, 0x32, 0x61, 0x6c, 0x03, 0x5f,
	0xc1, 0xb6, 0xd0, 0x3a, 0x67, 0x59, 0xc2, 0xe3, 0x51, 0xd1, 0x73, 0x35, 0xb4, 0x83, 0x81, 0x76,
	0x4a, 0xa7, 0x12, 0x97, 0x17, 0xe0, 0xa7, 0xcc, 0x30, 0x4c, 0xd0, 0x3e, 0x0d, 0xac, 0xaf, 0x85,
	0x9e, 0xa2, 0x36, 0x1a, 0x40, 0xfb, 0x86, 0x0d, 0x72, 0x7e, 0x25, 0x73, 0x95, 0x70, 0xf2, 0x1c,
	0x3c, 0xc5, 0x6f, 0x5d, 0xdc, 0xa9, 0xaf, 0x55, 0x92, 0x03, 0x68, 0x8c, 0xad, 0xab, 0x8b, 0xd4,
	0x99, 0x34, 0x54, 0xf4, 0x4c, 0x0b, 0x2b, 0x79, 0x0e, 0xc1, 0x48, 0x6a, 0xac, 0x19, 0xf1, 0xf2,
	0xe9, 0x44, 0x8e, 0x7e, 0x83, 0x6d, 0xcc, 0x76, 0xc1, 0xb5, 0x11, 0x19, 0xc3, 0xbe, 0xfe, 0xe3,
	0x94, 0x7f, 0xd7, 0xa1, 0xfd, 0xfd, 0x40, 0x26, 0xbf, 0xbe, 0xe1, 0x2c, 0xe5, 0x8a, 0x84, 0xd0,
	0x9a, 0xe5, 0x48, 0x29, 0xda, 0x59, 0xf4, 0xb9, 0xb8, 0xeb, 0x4f, 0x66, 0x51, 0x48, 0xe4, 0x15,
	0xec, 0x8c, 0x14, 0x1f, 0x0b, 0x99, 0xeb, 0xb8, 0x67, 0x23, 0xd9, 0xa1, 0x7a, 0x73, 0xe5, 0x76,
	0x4a, 0x17, 0xcc, 0xd5, 0x4d, 0xc9, 0x0b, 0xd8, 0x30, 0x62, 0xc8, 0xb5, 0x61, 0xc3, 0x11, 0xf2,
	0xc4, 0xa7, 0x53, 0x05, 0xf9, 0x12, 0x76, 0x8c, 0x62, 0x99, 0x66, 0x89, 0x2d, 0x52, 0xc7, 0x4a,
	0x4a, 0x13, 0x36, 0xe6, 0x62, 0x6e, 0x57, 0x5d, 0xa8, 0x94, 0x86, 0x7c, 0x07, 0xff, 0xab, 0xe8,
	0x62, 0x6d, 0x98, 0xc9, 0x75, 0xdc, 0x67, 0xba, 0x1f, 0x36, 0xe7, 0x0e, 0x3f, 0xad, 0x38, 0x5e,
	0xa1, 0x1f, 0x2e, 0xdc, 0x1e, 0x34, 0x32, 0x99, 0x25, 0x3c, 0x6c, 0x61, 0x49, 0x85, 0x60, 0x97,
	0xa3, 0x27, 0x8c, 0x0e, 0x03, 0x54, 0xe2, 0x37, 0xb9, 0x00, 0xb2, 0x98, 0x2b, 0xdc, 0xc0, 0x34,
	0x4f, 0x6d, 0x9a, 0xeb, 0xf9, 0x04, 0x74, 0x67, 0x21, 0x67, 0xf4, 0x67, 0x0d, 0x82, 0xeb, 0xfb,
	0x8f, 0x62, 0x7f, 0x08, 0x1d, 0xcd, 0x95, 0x60, 0x03, 0xf1, 0xc0, 0xd3, 0x58, 0x8b, 0x07, 0xee,
	0x86, 0xb0, 0x35, 0x55, 0x5f, 0x89, 0x07, 0x6e, 0xb7, 0xdc, 0xa2, 0x18, 0x2b, 0x96, 0xdd, 0x71,
	0x37, 0x6c, 0xc4, 0x95, 0x5a, 0x05, 0x39, 0x04, 0x50, 0x5c, 0xe7, 0x03, 0xbb, 0x78, 0x3a, 0xf4,
	0xf7, 0xbd, 0x19, 0x4c, 0x36, 0x0a, 0x5b, 0x37, 0xd5, 0xe4, 0x53, 0x08, 0xf8, 0xbd, 0x29, 0xa0,
	0x9b, 0xc7, 0xbd, 0xc5, 0xef, 0x8d, 0xfd, 0x88, 0x4e, 0x60, 0xeb, 0xfa, 0xfe, 0x86, 0x2b, 0x71,
	0xfb, 0x8e, 0xe2, 0x49, 0xf2, 0x09, 0xb4, 0x1d, 0xe8, 0xb7, 0x4c, 0x0c, 0xb0, 0x8b, 0x80, 0x42,
	0xa1, 0x7a, 0xcd, 0xc4, 0x20, 0xba, 0x85, 0x9d, 0x05, 0x5c, 0x3e, 0xd0, 0xf7, 0xd7, 0xf0, 0x64,
	0x8c, 0xf1, 0x4b, 0x7c, 0xeb, 0x58, 0x32, 0x41, 0x7c, 0x67, 0x52, 0xd3, 0xcd, 0xc2, 0xd1, 0xe1,
	0xfa, 0x57, 0x0d, 0xbc, 0xcb, 0xfc, 0x9e, 0x7c, 0x0e, 0x2d, 0x8d, 0xab, 0xab, 0xc3, 0x1a, 0x1e,
	0xc5, 0x1d, 0xa9, 0xac, 0x34, 0x2d, 0xed, 0xe4, 0x00, 0x5a, 0xe5, 0xbd, 0x51, 0x5f, 0xbc, 0x37,
	0x4a, 0xdb, 0x0c, 0x32, 0xde, 0x23, 0xc8, 0x90, 0x9f, 0x60, 0xef, 0x77, 0x61, 0x32, 0xae, 0x75,
	0x9c, 0x4e, 0x77, 0xb9, 0x44, 0x7c, 0x6f, 0x52, 0x43, 0x65, 0xd1, 0xe9, 0xae, 0x3b, 0x51, 0xd1,
	0x69, 0xf2, 0x05, 0xec, 0x94, 0x81, 0x98, 0xba, 0xcb, 0x87, 0x3c, 0x33, 0x3a, 0x6c, 0xec, 0x7b,
	0x2f, 0x37, 0xe9, 0xb6, 0x33, 0x9c, 0x97, 0x7a, 0x4b, 0xa6, 0xc6, 0xcf, 0x48, 0xd8, 0x4a, 0x2f,
	0xb5, 0x15, 0x7b, 0xa9, 0x3f, 0xd6, 0xcb, 0xd2, 0x12, 0xbc, 0xe5, 0x25, 0x90, 0x6f, 0x60, 0x77,
	0xe2, 0x9c, 0x25, 0x7d, 0xa9, 0x78, 0x6a, 0xaf, 0x03, 0x7f, 0x2e, 0x78, 0x19, 0xf1, 0xdc, 0xf9,
	0x74, 0xd3, 0x48, 0x42, 0xf0, 0x83, 0x14, 0x59, 0x8f, 0x69, 0x4e, 0x7e, 0x9c, 0x46, 0xa9, 0xc0,
	0xe7, 0x5a, 0x59, 0x8e, 0x1e, 0x59, 0x44, 0xcf, 0xde, 0x31, 0x4c, 0xf5, 0x84, 0x51, 0x4c, 0xbd,
	0x73, 0x3f, 0xb6, 0xa9, 0x22, 0x7a, 0x5f, 0x83, 0xe6, 0x2f, 0xb9, 0x19, 0xe5, 0x86, 0x1c, 0x42,
	0xb3, 0x60, 0x81, 0x4b, 0xb1, 0x40, 0x12, 0x67, 0x26, 0xaf, 0xa0, 0x93, 0xc8, 0xcc, 0x28, 0x39,
	0x88, 0x3f, 0xc0, 0x95, 0x2d, 0xe7, 0xf3, 0x76, 0x1d, 0xca, 0x84, 0xd0, 0x92, 0x2a, 0x15, 0x19,
	0x1b, 0xb8, 0xeb, 0xb0, 0x14, 0x23, 0x03, 0x40, 0xb9, 0x11, 0x8a, 0x5b, 0x88, 0x57, 0xaf, 0x75,
	0xa5, 0xe1, 0x56, 0xb2, 0x7a, 0xb3, 0x59, 0xff, 0xf0, 0x20, 0xe8, 0xba, 0x7f, 0x25, 0x39, 0x80,
	0x8d, 0x62, 0x9c, 0xd3, 0x1f, 0xf6, 0x34, 0x58, 0x50, 0x98, 0xba, 0xe9, 0xaa, 0xff, 0xa3, 0x95,
	0xf0, 0x78, 0x84, 0x03, 0xfe, 0x9a, 0x1c, 0xb8, 0x84, 0x70, 0x42, 0x48, 0x7c, 0x71, 0xa4, 0x93,
	0x27, 0x83, 0xbb, 0xd8, 0x76, 0x27, 0x55, 0x4e, 0x5f, 0x13, 0xf4, 0x59, 0x49, 0xd0, 0xb9, 0x57,
	0xc6, 0xd2, 0x65, 0x68, 0xae, 0xb7, 0x0c, 0xad, 0x8f, 0x2e, 0x43, 0x75, 0x2c, 0xc1, 0xec, 0x58,
	0xde, 0xd7, 0xa1, 0x71, 0x35, 0xe2, 0x59, 0x4a, 0x8e, 0xa1, 0xa3, 0x47, 0x3c, 0x33, 0xb1, 0x44,
	0x12, 0x2f, 0x9b, 0xcc, 0x13, 0x74, 0x28, 0x48, 0xde, 0x4d, 0x57, 0x63, 0xc4, 0x23, 0xb8, 0x7b,
	0x6b, 0xe2, 0xbe, 0x14, 0x28, 0x7f, 0x3d, 0xa0, 0x1a, 0x6b, 0x01, 0xd5, 0x9c, 0x01, 0xaa, 0xd7,
	0xc4, 0xd7, 0xf4, 0xd9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x00, 0x5c, 0x72, 0x59, 0x0b,
	0x00, 0x00,
}
