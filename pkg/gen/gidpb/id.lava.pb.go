// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
//   - protoc-gen-lava v0.0.2
//   - protoc          v5.28.2
// source: gid/id.proto

package gidpb

import rpcmeta "github.com/pubgo/lava/core/rpcmeta"

var _ = rpcmeta.Register(&rpcmeta.RpcMeta{
	Input:  new(GenerateRequest),
	Method: "example.gid.Id/Generate",
	Name:   "id.gen",
	Output: new(GenerateResponse),
	Tags: map[string]string{
		"hello":    "world",
		"internal": "true",
	},
})
