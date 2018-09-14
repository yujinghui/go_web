package session

import (
	"fmt"
	"sync"
)
var providers = make(map[string]Provider)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestory(sid string) error
	SessionGC(maxlifetime int64) 
}
type SessionManager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxlifetime int64
}

func NewManager(providerName, cookieName string, maxlifetime int64) (* SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session unkonwn providername")
	}

	return &SessionManager{provider:provider, cookieName:cookieName, maxlifetime:maxlifetime}, nil
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session : register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("sessioin: Register call twice ")
	}

	providers[name] = provider
}

