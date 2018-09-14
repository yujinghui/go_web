package sessions


var providers = make(map[string]Provider)


type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestory(sid string) error
	SessionGC(maxlifetime int64)
}

