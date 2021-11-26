package ormindex

import (
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/model/ormiterator"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
)

type Index interface {
	ormkv.Codec

	Fields() []protoreflect.Name
	PrefixIterator(store kv.IndexCommitmentReadStore, prefix []protoreflect.Value, options IteratorOptions) ormiterator.Iterator
	RangeIterator(store kv.IndexCommitmentReadStore, start, end []protoreflect.Value, options IteratorOptions) ormiterator.Iterator
	ReadValueFromIndexKey(store kv.IndexCommitmentReadStore, key, value []byte, message proto.Message) error

	doNotImplement()
}

type UniqueIndex interface {
	Index
	Has(store kv.IndexCommitmentReadStore, keyValues []protoreflect.Value) (found bool, err error)
	Get(store kv.IndexCommitmentReadStore, keyValues []protoreflect.Value, message proto.Message) (found bool, err error)
}

type IteratorOptions struct {
	Reverse bool
	Cursor  []byte
}

type Indexer interface {
	OnCreate(store kv.Store, message protoreflect.Message) error
	OnUpdate(store kv.Store, new, existing protoreflect.Message) error
	OnDelete(store kv.Store, message protoreflect.Message) error

	doNotImplement()
}