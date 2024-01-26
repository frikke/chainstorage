// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: coinbase/crypto/rosetta/types/transaction.proto

// The stable release for rosetta types

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Used by RelatedTransaction to indicate the direction of the relation (i.e. cross-shard/cross-network
// sends may reference backward to an earlier transaction and async execution may reference forward).
// Can be used to indicate if a transaction relation is from child to parent or the reverse.
type RelatedTransaction_Direction int32

const (
	// DIRECTION_UNSPECIFIED indicating the direction is not specified.
	RelatedTransaction_DIRECTION_UNSPECIFIED RelatedTransaction_Direction = 0
	// FORWARD indicating a transaction relation is from parent to child.
	RelatedTransaction_FORWARD RelatedTransaction_Direction = 1
	// BACKWARD indicating a transaction relation is from child to parent.
	RelatedTransaction_BACKWARD RelatedTransaction_Direction = 2
)

// Enum value maps for RelatedTransaction_Direction.
var (
	RelatedTransaction_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "FORWARD",
		2: "BACKWARD",
	}
	RelatedTransaction_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"FORWARD":               1,
		"BACKWARD":              2,
	}
)

func (x RelatedTransaction_Direction) Enum() *RelatedTransaction_Direction {
	p := new(RelatedTransaction_Direction)
	*p = x
	return p
}

func (x RelatedTransaction_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelatedTransaction_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_coinbase_crypto_rosetta_types_transaction_proto_enumTypes[0].Descriptor()
}

func (RelatedTransaction_Direction) Type() protoreflect.EnumType {
	return &file_coinbase_crypto_rosetta_types_transaction_proto_enumTypes[0]
}

func (x RelatedTransaction_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelatedTransaction_Direction.Descriptor instead.
func (RelatedTransaction_Direction) EnumDescriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_transaction_proto_rawDescGZIP(), []int{2, 0}
}

// Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction_identifier uniquely identifies a transaction in a particular network and block
	// or in the mempool.
	TransactionIdentifier *TransactionIdentifier `protobuf:"bytes,1,opt,name=transaction_identifier,json=transactionIdentifier,proto3" json:"transaction_identifier,omitempty"`
	// The operations is an array of Operations that are associated with the transaction
	Operations []*Operation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// The related_transaction allows implementations to link together multiple transactions.
	// An unpopulated network identifier indicates that the related transaction is on the same network.
	RelatedTransactions []*RelatedTransaction `protobuf:"bytes,3,rep,name=related_transactions,json=relatedTransactions,proto3" json:"related_transactions,omitempty"`
	// Transactions that are related to other transactions (like a cross-shard transaction) should
	// include the tranaction_identifier of these transactions in the metadata.
	Metadata map[string]*anypb.Any `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTransactionIdentifier() *TransactionIdentifier {
	if x != nil {
		return x.TransactionIdentifier
	}
	return nil
}

func (x *Transaction) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *Transaction) GetRelatedTransactions() []*RelatedTransaction {
	if x != nil {
		return x.RelatedTransactions
	}
	return nil
}

func (x *Transaction) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// TransactionIdentifier uniquely identifies a transaction in a particular network and block or in
// the mempool.
type TransactionIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the transaction.  Any transactions that are attributable only to a block (ex: a
	// block event) should use the hash of the block as the identifier.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *TransactionIdentifier) Reset() {
	*x = TransactionIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionIdentifier) ProtoMessage() {}

func (x *TransactionIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionIdentifier.ProtoReflect.Descriptor instead.
func (*TransactionIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionIdentifier) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// RelatedTransaction allows implementations to link together multiple transactions.
type RelatedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The network_identifier specifies which network a particular object is associated with.
	// An unpopulated network identifier indicates that the related transaction is on the same network.
	NetworkIdentifier *NetworkIdentifier `protobuf:"bytes,1,opt,name=network_identifier,json=networkIdentifier,proto3" json:"network_identifier,omitempty"`
	// The transaction_identifier uniquely identifies a transaction in a particular network and block
	// or in the mempool.
	TransactionIdentifier *TransactionIdentifier `protobuf:"bytes,2,opt,name=transaction_identifier,json=transactionIdentifier,proto3" json:"transaction_identifier,omitempty"`
	// Used by RelatedTransaction to indicate the direction of the relation.
	// Can be used to indicate if a transaction relation is from child to parent or the reverse.
	Direction RelatedTransaction_Direction `protobuf:"varint,3,opt,name=direction,proto3,enum=coinbase.crypto.rosetta.types.RelatedTransaction_Direction" json:"direction,omitempty"`
}

func (x *RelatedTransaction) Reset() {
	*x = RelatedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatedTransaction) ProtoMessage() {}

func (x *RelatedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatedTransaction.ProtoReflect.Descriptor instead.
func (*RelatedTransaction) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *RelatedTransaction) GetNetworkIdentifier() *NetworkIdentifier {
	if x != nil {
		return x.NetworkIdentifier
	}
	return nil
}

func (x *RelatedTransaction) GetTransactionIdentifier() *TransactionIdentifier {
	if x != nil {
		return x.TransactionIdentifier
	}
	return nil
}

func (x *RelatedTransaction) GetDirection() RelatedTransaction_Direction {
	if x != nil {
		return x.Direction
	}
	return RelatedTransaction_DIRECTION_UNSPECIFIED
}

var File_coinbase_crypto_rosetta_types_transaction_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_transaction_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73,
	0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65,
	0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x48, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x64, 0x0a, 0x14, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74,
	0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x54, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x80, 0x03, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x12,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74,
	0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x6b, 0x0a,
	0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41,
	0x43, 0x4b, 0x57, 0x41, 0x52, 0x44, 0x10, 0x02, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_transaction_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_transaction_proto_rawDescData = file_coinbase_crypto_rosetta_types_transaction_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_transaction_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_transaction_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_transaction_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_transaction_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coinbase_crypto_rosetta_types_transaction_proto_goTypes = []interface{}{
	(RelatedTransaction_Direction)(0), // 0: coinbase.crypto.rosetta.types.RelatedTransaction.Direction
	(*Transaction)(nil),               // 1: coinbase.crypto.rosetta.types.Transaction
	(*TransactionIdentifier)(nil),     // 2: coinbase.crypto.rosetta.types.TransactionIdentifier
	(*RelatedTransaction)(nil),        // 3: coinbase.crypto.rosetta.types.RelatedTransaction
	nil,                               // 4: coinbase.crypto.rosetta.types.Transaction.MetadataEntry
	(*Operation)(nil),                 // 5: coinbase.crypto.rosetta.types.Operation
	(*NetworkIdentifier)(nil),         // 6: coinbase.crypto.rosetta.types.NetworkIdentifier
	(*anypb.Any)(nil),                 // 7: google.protobuf.Any
}
var file_coinbase_crypto_rosetta_types_transaction_proto_depIdxs = []int32{
	2, // 0: coinbase.crypto.rosetta.types.Transaction.transaction_identifier:type_name -> coinbase.crypto.rosetta.types.TransactionIdentifier
	5, // 1: coinbase.crypto.rosetta.types.Transaction.operations:type_name -> coinbase.crypto.rosetta.types.Operation
	3, // 2: coinbase.crypto.rosetta.types.Transaction.related_transactions:type_name -> coinbase.crypto.rosetta.types.RelatedTransaction
	4, // 3: coinbase.crypto.rosetta.types.Transaction.metadata:type_name -> coinbase.crypto.rosetta.types.Transaction.MetadataEntry
	6, // 4: coinbase.crypto.rosetta.types.RelatedTransaction.network_identifier:type_name -> coinbase.crypto.rosetta.types.NetworkIdentifier
	2, // 5: coinbase.crypto.rosetta.types.RelatedTransaction.transaction_identifier:type_name -> coinbase.crypto.rosetta.types.TransactionIdentifier
	0, // 6: coinbase.crypto.rosetta.types.RelatedTransaction.direction:type_name -> coinbase.crypto.rosetta.types.RelatedTransaction.Direction
	7, // 7: coinbase.crypto.rosetta.types.Transaction.MetadataEntry.value:type_name -> google.protobuf.Any
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_transaction_proto_init() }
func file_coinbase_crypto_rosetta_types_transaction_proto_init() {
	if File_coinbase_crypto_rosetta_types_transaction_proto != nil {
		return
	}
	file_coinbase_crypto_rosetta_types_operation_proto_init()
	file_coinbase_crypto_rosetta_types_network_identifier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionIdentifier); i {
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
		file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatedTransaction); i {
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
			RawDescriptor: file_coinbase_crypto_rosetta_types_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_transaction_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_transaction_proto_depIdxs,
		EnumInfos:         file_coinbase_crypto_rosetta_types_transaction_proto_enumTypes,
		MessageInfos:      file_coinbase_crypto_rosetta_types_transaction_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_transaction_proto = out.File
	file_coinbase_crypto_rosetta_types_transaction_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_transaction_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_transaction_proto_depIdxs = nil
}
