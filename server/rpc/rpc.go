package rpc

import (
	consts "sliver/client/constants"
)

// RPCResponse - Called with response data, mapped back to reqID
type RPCResponse func([]byte, error)

// RPCHandler - RPC handlers accept bytes and return bytes
type RPCHandler func([]byte, RPCResponse)

var (
	rpcHandlers = &map[string]RPCHandler{
		consts.JobsStr: rpcJobs,
		consts.MtlsStr: rpcStartMTLSListener,
		consts.DnsStr:  rpcStartDNSListener,

		consts.SessionsStr: rpcSessions,
		consts.GenerateStr: rpcGenerate,

		consts.PsStr:       rpcPs,
		consts.LsStr:       rpcLs,
		consts.RmStr:       rpcRm,
		consts.MkdirStr:    rpcMkdir,
		consts.CdStr:       rpcCd,
		consts.PwdStr:      rpcPwd,
		consts.DownloadStr: rpcDownload,
		consts.UploadStr:   rpcUpload,
	}
)

// GetRPCHandlers - Returns a map of server-side msg handlers
func GetRPCHandlers() *map[string]RPCHandler {
	return rpcHandlers
}
