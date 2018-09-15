package memory

import (
	"container/list"
	"sync"
)
type 	Provider struct {
	lock sync.Mutex
	sessions map[string]*list.Element
	list *list.list
}
var pder = &Provider{list: list.New()}

func (pder *Provider) SessionStart(sid string) (Session, error) {

	return nil
}

func (pder *Provider) SessionRead(sid string) (Session, error) {
	return nil
}

func (pder *Provider) SessionDestory(sid string) error {
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {

}

type SessionStore struct {
	sid string
	timeAccessed time.Time
	value map[interface{}]interface{}
}

func (ss *SessionStore) Set(key, value interface{}) {
	ss.value[key] = value
	pder.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(ss.sid)
	if v, ok := st.value[interface{}], ok {
		return v
	}
	return nil

}

func (ss *SessionStore) Delete(key interface{}) {
	delete(ss.value, key)
	pder.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) SessionID() {
	return ss.sid
}