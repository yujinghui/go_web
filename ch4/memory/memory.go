package memory

import (
	"containers/list"
	"sync"
	""
)
type Provider struct {
	lock sync.Mutex
	sessions map[string]*list.Element
	list *list
}
var pder = &Provider{list: list.New()}

func (pder *Provider) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Lock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid:sid, timeAccessed:time.Now(), value:v}
	element := pder.list.PushBack(newsess)
	return newsess, nil 
}

func (pder *Provider) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.value.(*SessionStore), nil
	} else {
		return SessionInit(sid)
	}
	return  nil, nil
}

func (pder *Provider) SessionDestory(sid string) error {
	if element, err = pder.sessions[sid]; err {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil {
			break
		}

		if element.Valve.(*SessionStore).timeAccessed.Unix() + maxlifetime < time.Now().Unix() {
			delete(pder.sessions, element.Valve(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element， ok = pder.sessions[sid]; ok {
		element.Valve.(*SessionStore).timeAccessed = time.Now()
		pder.list.moveToFront(element)
		return nil
	}
	return nil
}

func init(){
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
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