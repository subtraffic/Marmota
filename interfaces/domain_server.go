package interfaces

type DomainUpdater interface {
	NeedUpdate(currentIp string) bool
	GetCurrentIp() IpModel
	QueryIp()
	QueryIpOnEchoServer()
	GetHistoryRecord(count int, offset int) []IpModel
}

type IpModel interface {
	Update(currentIp string, echoSrv EchoServer)
	GetIp() string
	GetName() string
}

type EchoServer interface{}
