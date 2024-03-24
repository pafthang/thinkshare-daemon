package persistent

import "github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"

type Persistent interface {
	Log(source string, level string, log string)
	Infor(func () *packet.WorkerInfor)
	Sessions(func ()[]packet.WorkerSession)

	RecvSession(func(*packet.WorkerSession) (*packet.WorkerSession,error))
	ClosedSession() int
	Stop()
}
