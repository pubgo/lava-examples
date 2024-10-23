// Code generated by protoc-gen-cloudjobs. DO NOT EDIT.
// versions:
//   - protoc-gen-cloudjobs v0.0.2
//   - protoc               v5.28.2
// source: gid/id.proto

package gidpb

import (
	"context"
	cloudjobs "github.com/pubgo/funk/component/cloudjobs"
	cloudjobpb "github.com/pubgo/funk/pkg/gen/cloudjobpb"
)

const IdServiceJobKey = "gid"

// IdServiceEventChangedKey Id/EventChanged
const IdServiceEventChangedKey = "gid.event.update"

// IdServiceProxyExecEventKey Id/ProxyExecEvent
const IdServiceProxyExecEventKey = "gid.proxy.exec"

func RegisterIdServiceEventChangedCloudJob(jobCli *cloudjobs.Client, handler func(ctx *cloudjobs.Context, req *DoProxyEventReq) error, opts ...*cloudjobpb.RegisterJobOptions) {
	var _ = cloudjobs.RegisterSubject(IdServiceEventChangedKey, new(DoProxyEventReq))
	cloudjobs.RegisterJobHandler(jobCli, IdServiceJobKey, IdServiceEventChangedKey, handler, opts...)
}

func PushIdServiceEventChangedCloudJob(jobCli *cloudjobs.Client, ctx context.Context, req *DoProxyEventReq, opts ...*cloudjobpb.PushEventOptions) error {
	return jobCli.Publish(ctx, IdServiceEventChangedKey, req, opts...)
}
func RegisterIdServiceProxyExecEventCloudJob(jobCli *cloudjobs.Client, handler func(ctx *cloudjobs.Context, req *DoProxyEventReq) error, opts ...*cloudjobpb.RegisterJobOptions) {
	var _ = cloudjobs.RegisterSubject(IdServiceProxyExecEventKey, new(DoProxyEventReq))
	cloudjobs.RegisterJobHandler(jobCli, IdServiceJobKey, IdServiceProxyExecEventKey, handler, opts...)
}

func PushIdServiceProxyExecEventCloudJob(jobCli *cloudjobs.Client, ctx context.Context, req *DoProxyEventReq, opts ...*cloudjobpb.PushEventOptions) error {
	return jobCli.Publish(ctx, IdServiceProxyExecEventKey, req, opts...)
}
