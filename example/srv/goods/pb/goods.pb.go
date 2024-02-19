// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: pb/goods.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Good struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64   `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string  `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
	Price float64 `protobuf:"fixed64,30,opt,name=Price,proto3" json:"Price,omitempty"`
	Stock int64   `protobuf:"varint,40,opt,name=Stock,proto3" json:"Stock,omitempty"`
}

func (x *Good) Reset() {
	*x = Good{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Good) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Good) ProtoMessage() {}

func (x *Good) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Good.ProtoReflect.Descriptor instead.
func (*Good) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{0}
}

func (x *Good) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Good) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Good) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Good) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GoodsStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Stock int64 `protobuf:"varint,20,opt,name=Stock,proto3" json:"Stock,omitempty"`
}

func (x *GoodsStock) Reset() {
	*x = GoodsStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsStock) ProtoMessage() {}

func (x *GoodsStock) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsStock.ProtoReflect.Descriptor instead.
func (*GoodsStock) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsStock) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GoodsStock) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Good *Good `protobuf:"bytes,10,opt,name=good,proto3" json:"good,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{2}
}

func (x *AddReq) GetGood() *Good {
	if x != nil {
		return x.Good
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Good *Good `protobuf:"bytes,10,opt,name=good,proto3" json:"good,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReq) GetGood() *Good {
	if x != nil {
		return x.Good
	}
	return nil
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GetReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Good *Good `protobuf:"bytes,10,opt,name=good,proto3" json:"good,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{5}
}

func (x *GetResp) GetGood() *Good {
	if x != nil {
		return x.Good
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int64 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *FindReq) Reset() {
	*x = FindReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindReq) ProtoMessage() {}

func (x *FindReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindReq.ProtoReflect.Descriptor instead.
func (*FindReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{7}
}

func (x *FindReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FindResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods []*Good `protobuf:"bytes,10,rep,name=goods,proto3" json:"goods,omitempty"`
}

func (x *FindResp) Reset() {
	*x = FindResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResp) ProtoMessage() {}

func (x *FindResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResp.ProtoReflect.Descriptor instead.
func (*FindResp) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{8}
}

func (x *FindResp) GetGoods() []*Good {
	if x != nil {
		return x.Goods
	}
	return nil
}

type FindByIDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs []int64 `protobuf:"varint,10,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
}

func (x *FindByIDsReq) Reset() {
	*x = FindByIDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIDsReq) ProtoMessage() {}

func (x *FindByIDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIDsReq.ProtoReflect.Descriptor instead.
func (*FindByIDsReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{9}
}

func (x *FindByIDsReq) GetIDs() []int64 {
	if x != nil {
		return x.IDs
	}
	return nil
}

type UpdateStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods []*GoodsStock `protobuf:"bytes,10,rep,name=goods,proto3" json:"goods,omitempty"`
	GID   string        `protobuf:"bytes,20,opt,name=GID,proto3" json:"GID,omitempty"`
}

func (x *UpdateStockReq) Reset() {
	*x = UpdateStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockReq) ProtoMessage() {}

func (x *UpdateStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockReq.ProtoReflect.Descriptor instead.
func (*UpdateStockReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateStockReq) GetGoods() []*GoodsStock {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *UpdateStockReq) GetGID() string {
	if x != nil {
		return x.GID
	}
	return ""
}

type SearchGoodsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	Offset  int64  `protobuf:"varint,20,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit   int64  `protobuf:"varint,30,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchGoodsReq) Reset() {
	*x = SearchGoodsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchGoodsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGoodsReq) ProtoMessage() {}

func (x *SearchGoodsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGoodsReq.ProtoReflect.Descriptor instead.
func (*SearchGoodsReq) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{11}
}

func (x *SearchGoodsReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SearchGoodsReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SearchGoodsReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchGoodsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods []*Good `protobuf:"bytes,10,rep,name=goods,proto3" json:"goods,omitempty"`
}

func (x *SearchGoodsResp) Reset() {
	*x = SearchGoodsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchGoodsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGoodsResp) ProtoMessage() {}

func (x *SearchGoodsResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGoodsResp.ProtoReflect.Descriptor instead.
func (*SearchGoodsResp) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{12}
}

func (x *SearchGoodsResp) GetGoods() []*Good {
	if x != nil {
		return x.Goods
	}
	return nil
}

type GoodsEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoodsEmpty) Reset() {
	*x = GoodsEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsEmpty) ProtoMessage() {}

func (x *GoodsEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsEmpty.ProtoReflect.Descriptor instead.
func (*GoodsEmpty) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{13}
}

var File_pb_goods_proto protoreflect.FileDescriptor

var file_pb_goods_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x56, 0x0a, 0x04, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x0a,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x22, 0x26, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x04, 0x67, 0x6f,
	0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x22, 0x29, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x67,
	0x6f, 0x6f, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2a, 0x0a, 0x08,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x49, 0x44, 0x73, 0x22, 0x48, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x05,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x47, 0x49, 0x44, 0x22, 0x58, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x31,
	0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1e, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x95, 0x03, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x36, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_goods_proto_rawDescOnce sync.Once
	file_pb_goods_proto_rawDescData = file_pb_goods_proto_rawDesc
)

func file_pb_goods_proto_rawDescGZIP() []byte {
	file_pb_goods_proto_rawDescOnce.Do(func() {
		file_pb_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_goods_proto_rawDescData)
	})
	return file_pb_goods_proto_rawDescData
}

var file_pb_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pb_goods_proto_goTypes = []interface{}{
	(*Good)(nil),            // 0: pb.Good
	(*GoodsStock)(nil),      // 1: pb.GoodsStock
	(*AddReq)(nil),          // 2: pb.AddReq
	(*UpdateReq)(nil),       // 3: pb.UpdateReq
	(*GetReq)(nil),          // 4: pb.GetReq
	(*GetResp)(nil),         // 5: pb.GetResp
	(*DeleteReq)(nil),       // 6: pb.DeleteReq
	(*FindReq)(nil),         // 7: pb.FindReq
	(*FindResp)(nil),        // 8: pb.FindResp
	(*FindByIDsReq)(nil),    // 9: pb.FindByIDsReq
	(*UpdateStockReq)(nil),  // 10: pb.UpdateStockReq
	(*SearchGoodsReq)(nil),  // 11: pb.SearchGoodsReq
	(*SearchGoodsResp)(nil), // 12: pb.SearchGoodsResp
	(*GoodsEmpty)(nil),      // 13: pb.GoodsEmpty
}
var file_pb_goods_proto_depIdxs = []int32{
	0,  // 0: pb.AddReq.good:type_name -> pb.Good
	0,  // 1: pb.UpdateReq.good:type_name -> pb.Good
	0,  // 2: pb.GetResp.good:type_name -> pb.Good
	0,  // 3: pb.FindResp.goods:type_name -> pb.Good
	1,  // 4: pb.UpdateStockReq.goods:type_name -> pb.GoodsStock
	0,  // 5: pb.SearchGoodsResp.goods:type_name -> pb.Good
	2,  // 6: pb.GoodsService.Add:input_type -> pb.AddReq
	3,  // 7: pb.GoodsService.Update:input_type -> pb.UpdateReq
	4,  // 8: pb.GoodsService.Get:input_type -> pb.GetReq
	6,  // 9: pb.GoodsService.Delete:input_type -> pb.DeleteReq
	7,  // 10: pb.GoodsService.Find:input_type -> pb.FindReq
	9,  // 11: pb.GoodsService.FindByIds:input_type -> pb.FindByIDsReq
	10, // 12: pb.GoodsService.UpdateStock:input_type -> pb.UpdateStockReq
	10, // 13: pb.GoodsService.UpdateStockBack:input_type -> pb.UpdateStockReq
	11, // 14: pb.GoodsService.SearchGoods:input_type -> pb.SearchGoodsReq
	13, // 15: pb.GoodsService.Add:output_type -> pb.GoodsEmpty
	13, // 16: pb.GoodsService.Update:output_type -> pb.GoodsEmpty
	5,  // 17: pb.GoodsService.Get:output_type -> pb.GetResp
	13, // 18: pb.GoodsService.Delete:output_type -> pb.GoodsEmpty
	8,  // 19: pb.GoodsService.Find:output_type -> pb.FindResp
	8,  // 20: pb.GoodsService.FindByIds:output_type -> pb.FindResp
	13, // 21: pb.GoodsService.UpdateStock:output_type -> pb.GoodsEmpty
	13, // 22: pb.GoodsService.UpdateStockBack:output_type -> pb.GoodsEmpty
	12, // 23: pb.GoodsService.SearchGoods:output_type -> pb.SearchGoodsResp
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pb_goods_proto_init() }
func file_pb_goods_proto_init() {
	if File_pb_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Good); i {
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
		file_pb_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsStock); i {
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
		file_pb_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_pb_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_pb_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_pb_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
		file_pb_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_pb_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindReq); i {
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
		file_pb_goods_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindResp); i {
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
		file_pb_goods_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIDsReq); i {
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
		file_pb_goods_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockReq); i {
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
		file_pb_goods_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchGoodsReq); i {
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
		file_pb_goods_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchGoodsResp); i {
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
		file_pb_goods_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsEmpty); i {
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
			RawDescriptor: file_pb_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_goods_proto_goTypes,
		DependencyIndexes: file_pb_goods_proto_depIdxs,
		MessageInfos:      file_pb_goods_proto_msgTypes,
	}.Build()
	File_pb_goods_proto = out.File
	file_pb_goods_proto_rawDesc = nil
	file_pb_goods_proto_goTypes = nil
	file_pb_goods_proto_depIdxs = nil
}