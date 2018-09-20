package memory

import (
	"container/list"
	"goWeb/ch4/session"
	"sync"
	"time"
)

type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

var pder = &Provider{list: list.New()}

func (pder *Provider) SessionInit(sid string) (sessions.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Lock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (sessions.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *Provider) SessionDestory(sid string) error {
	if element, err := pder.sessions[sid]; err {
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

		if element.Value.(*SessionStore).timeAccessed.Unix()+maxlifetime < time.Now().Unix() {
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

func (ss *SessionStore) Set(key, value interface{}) error {
	ss.value[key] = value
	pder.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(ss.sid)
	if v, ok := ss.value[key]; ok {
		return v
	}
	return nil

}

func (ss *SessionStore) Delete(key interface{}) error {
	delete(ss.value, key)
	pder.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) SessionID() string {
	return ss.sid
}


func Init() {
	pder.sessions = make(map[string]*list.Element, 0)
	sessions.Register("memory", pder)
}
