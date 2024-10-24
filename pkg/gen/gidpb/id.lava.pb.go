// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
//   - protoc-gen-lava v0.0.1
//   - protoc          v5.28.2
// source: gid/id.proto

package gidpb

import (
	assert "github.com/pubgo/funk/assert"
	lavapbv1 "github.com/pubgo/lava/pkg/proto/lavapbv1"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// IdServiceGenerateAction Id/Generate
// Generate 生成ID
var IdServiceGenerateAction = func() *lavapbv1.RpcMeta {
	var p lavapbv1.RpcMeta
	var data = `{"name":"id.gen","tags":{"hello":"world","internal":"true"}}`
	assert.Exit(protojson.Unmarshal([]byte(data), &p))
	return &p
}()
