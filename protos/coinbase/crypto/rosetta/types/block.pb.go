// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: coinbase/crypto/rosetta/types/block.proto

// The stable release for rosetta types

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Blocks contain an array of Transactions that occurred at a particular BlockIdentifier.
// A hard requirement for blocks returned by Rosetta implementations is that they MUST be
// immutable: once a client has requested and received a block identified by a specific
// BlockIndentifier, all future calls for that same BlockIdentifier must return the same
// block contents.
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BlockIdentifier uniquely identifies a block in a particular network.
	BlockIdentifier *BlockIdentifier `protobuf:"bytes,2,opt,name=block_identifier,json=blockIdentifier,proto3" json:"block_identifier,omitempty"`
	// BlockIdentifier uniquely identifies a block in a particular network.
	ParentBlockIdentifier *BlockIdentifier `protobuf:"bytes,3,opt,name=parent_block_identifier,json=parentBlockIdentifier,proto3" json:"parent_block_identifier,omitempty"`
	// The timestamp of the block since the Unix Epoch. The rosetta spec only
	// support timestamp in milliseconds granularity.
	// (-- api-linter: core::0142::time-field-names=disabled
	//
	//	go/aip/not-precedent: This field was called timestamp in rosetta spec. --)
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Transactions within the block.
	Transactions []*Transaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// Metadata for the block.
	Metadata map[string]*anypb.Any `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetBlockIdentifier() *BlockIdentifier {
	if x != nil {
		return x.BlockIdentifier
	}
	return nil
}

func (x *Block) GetParentBlockIdentifier() *BlockIdentifier {
	if x != nil {
		return x.ParentBlockIdentifier
	}
	return nil
}

func (x *Block) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// The block identifier uniquely identifies a block in a particular network.
type BlockIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The block height.  Note that this this should be >= 0
	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The block hash.
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BlockIdentifier) Reset() {
	*x = BlockIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockIdentifier) ProtoMessage() {}

func (x *BlockIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockIdentifier.ProtoReflect.Descriptor instead.
func (*BlockIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_block_proto_rawDescGZIP(), []int{1}
}

func (x *BlockIdentifier) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BlockIdentifier) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_coinbase_crypto_rosetta_types_block_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_block_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73,
	0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x04, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x59, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f,
	0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x17,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x15, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4e,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x51,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0f, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_block_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_block_proto_rawDescData = file_coinbase_crypto_rosetta_types_block_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_block_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_block_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_block_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_block_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_crypto_rosetta_types_block_proto_goTypes = []interface{}{
	(*Block)(nil),                 // 0: coinbase.crypto.rosetta.types.Block
	(*BlockIdentifier)(nil),       // 1: coinbase.crypto.rosetta.types.BlockIdentifier
	nil,                           // 2: coinbase.crypto.rosetta.types.Block.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Transaction)(nil),           // 4: coinbase.crypto.rosetta.types.Transaction
	(*anypb.Any)(nil),             // 5: google.protobuf.Any
}
var file_coinbase_crypto_rosetta_types_block_proto_depIdxs = []int32{
	1, // 0: coinbase.crypto.rosetta.types.Block.block_identifier:type_name -> coinbase.crypto.rosetta.types.BlockIdentifier
	1, // 1: coinbase.crypto.rosetta.types.Block.parent_block_identifier:type_name -> coinbase.crypto.rosetta.types.BlockIdentifier
	3, // 2: coinbase.crypto.rosetta.types.Block.timestamp:type_name -> google.protobuf.Timestamp
	4, // 3: coinbase.crypto.rosetta.types.Block.transactions:type_name -> coinbase.crypto.rosetta.types.Transaction
	2, // 4: coinbase.crypto.rosetta.types.Block.metadata:type_name -> coinbase.crypto.rosetta.types.Block.MetadataEntry
	5, // 5: coinbase.crypto.rosetta.types.Block.MetadataEntry.value:type_name -> google.protobuf.Any
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_block_proto_init() }
func file_coinbase_crypto_rosetta_types_block_proto_init() {
	if File_coinbase_crypto_rosetta_types_block_proto != nil {
		return
	}
	file_coinbase_crypto_rosetta_types_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coinbase_crypto_rosetta_types_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockIdentifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_crypto_rosetta_types_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_block_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_block_proto_depIdxs,
		MessageInfos:      file_coinbase_crypto_rosetta_types_block_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_block_proto = out.File
	file_coinbase_crypto_rosetta_types_block_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_block_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_block_proto_depIdxs = nil
}
