// Code generated by protoc-gen-go from "riak.proto"
// DO NOT EDIT!

package riak

import proto "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf

type RpbIndexReq_IndexQueryType int32

const (
	RpbIndexReq_eq    RpbIndexReq_IndexQueryType = 0
	RpbIndexReq_range RpbIndexReq_IndexQueryType = 1
)

var RpbIndexReq_IndexQueryType_name = map[int32]string{
	0: "eq",
	1: "range",
}
var RpbIndexReq_IndexQueryType_value = map[string]int32{
	"eq":    0,
	"range": 1,
}

// NewRpbIndexReq_IndexQueryType is deprecated. Use x.Enum() instead.
func NewRpbIndexReq_IndexQueryType(x RpbIndexReq_IndexQueryType) *RpbIndexReq_IndexQueryType {
	e := RpbIndexReq_IndexQueryType(x)
	return &e
}
func (x RpbIndexReq_IndexQueryType) Enum() *RpbIndexReq_IndexQueryType {
	p := new(RpbIndexReq_IndexQueryType)
	*p = x
	return p
}
func (x RpbIndexReq_IndexQueryType) String() string {
	return proto.EnumName(RpbIndexReq_IndexQueryType_name, int32(x))
}

type RpbErrorResp struct {
	Errmsg           []byte  `protobuf:"bytes,1,req,name=errmsg" json:"errmsg,omitempty"`
	Errcode          *uint32 `protobuf:"varint,2,req,name=errcode" json:"errcode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *RpbErrorResp) Reset()         { *this = RpbErrorResp{} }
func (this *RpbErrorResp) String() string { return proto.CompactTextString(this) }

type RpbGetServerInfoResp struct {
	Node             []byte `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	ServerVersion    []byte `protobuf:"bytes,2,opt,name=server_version" json:"server_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbGetServerInfoResp) Reset()         { *this = RpbGetServerInfoResp{} }
func (this *RpbGetServerInfoResp) String() string { return proto.CompactTextString(this) }

type RpbPair struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbPair) Reset()         { *this = RpbPair{} }
func (this *RpbPair) String() string { return proto.CompactTextString(this) }

type RpbGetClientIdResp struct {
	ClientId         []byte `protobuf:"bytes,1,req,name=client_id" json:"client_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbGetClientIdResp) Reset()         { *this = RpbGetClientIdResp{} }
func (this *RpbGetClientIdResp) String() string { return proto.CompactTextString(this) }

type RpbSetClientIdReq struct {
	ClientId         []byte `protobuf:"bytes,1,req,name=client_id" json:"client_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbSetClientIdReq) Reset()         { *this = RpbSetClientIdReq{} }
func (this *RpbSetClientIdReq) String() string { return proto.CompactTextString(this) }

type RpbGetReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	R                *uint32 `protobuf:"varint,3,opt,name=r" json:"r,omitempty"`
	Pr               *uint32 `protobuf:"varint,4,opt,name=pr" json:"pr,omitempty"`
	BasicQuorum      *bool   `protobuf:"varint,5,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk       *bool   `protobuf:"varint,6,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	IfModified       []byte  `protobuf:"bytes,7,opt,name=if_modified" json:"if_modified,omitempty"`
	Head             *bool   `protobuf:"varint,8,opt,name=head" json:"head,omitempty"`
	Deletedvclock    *bool   `protobuf:"varint,9,opt,name=deletedvclock" json:"deletedvclock,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *RpbGetReq) Reset()         { *this = RpbGetReq{} }
func (this *RpbGetReq) String() string { return proto.CompactTextString(this) }

type RpbGetResp struct {
	Content          []*RpbContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Vclock           []byte        `protobuf:"bytes,2,opt,name=vclock" json:"vclock,omitempty"`
	Unchanged        *bool         `protobuf:"varint,3,opt,name=unchanged" json:"unchanged,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (this *RpbGetResp) Reset()         { *this = RpbGetResp{} }
func (this *RpbGetResp) String() string { return proto.CompactTextString(this) }

type RpbPutReq struct {
	Bucket           []byte      `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte      `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Vclock           []byte      `protobuf:"bytes,3,opt,name=vclock" json:"vclock,omitempty"`
	Content          *RpbContent `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	W                *uint32     `protobuf:"varint,5,opt,name=w" json:"w,omitempty"`
	Dw               *uint32     `protobuf:"varint,6,opt,name=dw" json:"dw,omitempty"`
	ReturnBody       *bool       `protobuf:"varint,7,opt,name=return_body" json:"return_body,omitempty"`
	Pw               *uint32     `protobuf:"varint,8,opt,name=pw" json:"pw,omitempty"`
	IfNotModified    *bool       `protobuf:"varint,9,opt,name=if_not_modified" json:"if_not_modified,omitempty"`
	IfNoneMatch      *bool       `protobuf:"varint,10,opt,name=if_none_match" json:"if_none_match,omitempty"`
	ReturnHead       *bool       `protobuf:"varint,11,opt,name=return_head" json:"return_head,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (this *RpbPutReq) Reset()         { *this = RpbPutReq{} }
func (this *RpbPutReq) String() string { return proto.CompactTextString(this) }

type RpbPutResp struct {
	Content          []*RpbContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Vclock           []byte        `protobuf:"bytes,2,opt,name=vclock" json:"vclock,omitempty"`
	Key              []byte        `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (this *RpbPutResp) Reset()         { *this = RpbPutResp{} }
func (this *RpbPutResp) String() string { return proto.CompactTextString(this) }

type RpbDelReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Rw               *uint32 `protobuf:"varint,3,opt,name=rw" json:"rw,omitempty"`
	Vclock           []byte  `protobuf:"bytes,4,opt,name=vclock" json:"vclock,omitempty"`
	R                *uint32 `protobuf:"varint,5,opt,name=r" json:"r,omitempty"`
	W                *uint32 `protobuf:"varint,6,opt,name=w" json:"w,omitempty"`
	Pr               *uint32 `protobuf:"varint,7,opt,name=pr" json:"pr,omitempty"`
	Pw               *uint32 `protobuf:"varint,8,opt,name=pw" json:"pw,omitempty"`
	Dw               *uint32 `protobuf:"varint,9,opt,name=dw" json:"dw,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *RpbDelReq) Reset()         { *this = RpbDelReq{} }
func (this *RpbDelReq) String() string { return proto.CompactTextString(this) }

type RpbListBucketsResp struct {
	Buckets          [][]byte `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *RpbListBucketsResp) Reset()         { *this = RpbListBucketsResp{} }
func (this *RpbListBucketsResp) String() string { return proto.CompactTextString(this) }

type RpbListKeysReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbListKeysReq) Reset()         { *this = RpbListKeysReq{} }
func (this *RpbListKeysReq) String() string { return proto.CompactTextString(this) }

type RpbListKeysResp struct {
	Keys             [][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Done             *bool    `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *RpbListKeysResp) Reset()         { *this = RpbListKeysResp{} }
func (this *RpbListKeysResp) String() string { return proto.CompactTextString(this) }

type RpbGetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbGetBucketReq) Reset()         { *this = RpbGetBucketReq{} }
func (this *RpbGetBucketReq) String() string { return proto.CompactTextString(this) }

type RpbGetBucketResp struct {
	Props            *RpbBucketProps `protobuf:"bytes,1,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *RpbGetBucketResp) Reset()         { *this = RpbGetBucketResp{} }
func (this *RpbGetBucketResp) String() string { return proto.CompactTextString(this) }

type RpbSetBucketReq struct {
	Bucket           []byte          `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *RpbSetBucketReq) Reset()         { *this = RpbSetBucketReq{} }
func (this *RpbSetBucketReq) String() string { return proto.CompactTextString(this) }

type RpbMapRedReq struct {
	Request          []byte `protobuf:"bytes,1,req,name=request" json:"request,omitempty"`
	ContentType      []byte `protobuf:"bytes,2,req,name=content_type" json:"content_type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbMapRedReq) Reset()         { *this = RpbMapRedReq{} }
func (this *RpbMapRedReq) String() string { return proto.CompactTextString(this) }

type RpbMapRedResp struct {
	Phase            *uint32 `protobuf:"varint,1,opt,name=phase" json:"phase,omitempty"`
	Response         []byte  `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Done             *bool   `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *RpbMapRedResp) Reset()         { *this = RpbMapRedResp{} }
func (this *RpbMapRedResp) String() string { return proto.CompactTextString(this) }

type RpbIndexReq struct {
	Bucket           []byte                      `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Index            []byte                      `protobuf:"bytes,2,req,name=index" json:"index,omitempty"`
	Qtype            *RpbIndexReq_IndexQueryType `protobuf:"varint,3,req,name=qtype,enum=RpbIndexReq_IndexQueryType" json:"qtype,omitempty"`
	Key              []byte                      `protobuf:"bytes,4,opt,name=key" json:"key,omitempty"`
	RangeMin         []byte                      `protobuf:"bytes,5,opt,name=range_min" json:"range_min,omitempty"`
	RangeMax         []byte                      `protobuf:"bytes,6,opt,name=range_max" json:"range_max,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (this *RpbIndexReq) Reset()         { *this = RpbIndexReq{} }
func (this *RpbIndexReq) String() string { return proto.CompactTextString(this) }

type RpbIndexResp struct {
	Keys             [][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *RpbIndexResp) Reset()         { *this = RpbIndexResp{} }
func (this *RpbIndexResp) String() string { return proto.CompactTextString(this) }

type RpbContent struct {
	Value            []byte     `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	ContentType      []byte     `protobuf:"bytes,2,opt,name=content_type" json:"content_type,omitempty"`
	Charset          []byte     `protobuf:"bytes,3,opt,name=charset" json:"charset,omitempty"`
	ContentEncoding  []byte     `protobuf:"bytes,4,opt,name=content_encoding" json:"content_encoding,omitempty"`
	Vtag             []byte     `protobuf:"bytes,5,opt,name=vtag" json:"vtag,omitempty"`
	Links            []*RpbLink `protobuf:"bytes,6,rep,name=links" json:"links,omitempty"`
	LastMod          *uint32    `protobuf:"varint,7,opt,name=last_mod" json:"last_mod,omitempty"`
	LastModUsecs     *uint32    `protobuf:"varint,8,opt,name=last_mod_usecs" json:"last_mod_usecs,omitempty"`
	Usermeta         []*RpbPair `protobuf:"bytes,9,rep,name=usermeta" json:"usermeta,omitempty"`
	Indexes          []*RpbPair `protobuf:"bytes,10,rep,name=indexes" json:"indexes,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *RpbContent) Reset()         { *this = RpbContent{} }
func (this *RpbContent) String() string { return proto.CompactTextString(this) }

type RpbLink struct {
	Bucket           []byte `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key              []byte `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Tag              []byte `protobuf:"bytes,3,opt,name=tag" json:"tag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *RpbLink) Reset()         { *this = RpbLink{} }
func (this *RpbLink) String() string { return proto.CompactTextString(this) }

type RpbBucketProps struct {
	NVal             *uint32 `protobuf:"varint,1,opt,name=n_val" json:"n_val,omitempty"`
	AllowMult        *bool   `protobuf:"varint,2,opt,name=allow_mult" json:"allow_mult,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *RpbBucketProps) Reset()         { *this = RpbBucketProps{} }
func (this *RpbBucketProps) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("RpbIndexReq_IndexQueryType", RpbIndexReq_IndexQueryType_name, RpbIndexReq_IndexQueryType_value)
}