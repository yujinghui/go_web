package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

func (manager *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *SessionManager) SessionStart(writer http.ResponseWriter, request *http.Request) Session {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := request.Cookie(manager.cookieName)
	if err == nil || cookie.Value == "" {
		fmt.Printf("get cookie error %v", err)
		sid := manager.sessionId()
		session, err := manager.provider.SessionInit(sid)
		if err != nil || session == nil {
			fmt.Println("init session error")
			panic("init session error")
		}
		return session
	}
	return nil
}

func (manager *SessionManager) SessionDestory(writer http.ResponseWriter, req *http.Request) error {
		cookie, err := req.Cookie(manager.cookieName)
		if(err != nil) {
			fmt.Println("destory session error")
			return err
		}
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestory(cookie.Value)
		return nil

}

func NewManager(providerName, cookieName string, maxlifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session unkonwn providername")
	}

	return &SessionManager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
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
