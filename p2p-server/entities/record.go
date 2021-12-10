package entities
// Yue Yu
import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

type (
	NodeRecord  map[string]string
	TokenRecord map[string]string

	NodeTokenRecord struct {
		nodeRecord  NodeRecord
		tokenRecord TokenRecord
		m           sync.Mutex
	}
)

func GenToken(name string) string {
	token := sha256.Sum256([]byte(name + time.Now().String()))

	return fmt.Sprintf("%x", token)
}

func CreateNodeTokenRecord() *NodeTokenRecord {
	r := &NodeTokenRecord{}
	r.init()

	return r
}

func (r *NodeTokenRecord) init() {
	r.nodeRecord = make(NodeRecord)
	r.tokenRecord = make(TokenRecord)
}

func (r *NodeTokenRecord) GetNodes() []string {
	r.m.Lock()
	defer r.m.Unlock()

	var list []string
	for k := range r.nodeRecord {
		list = append(list, k)
	}

	return list
}

func (r *NodeTokenRecord) GetTokens() []string {
	r.m.Lock()
	defer r.m.Unlock()

	var list []string
	for k := range r.tokenRecord {
		list = append(list, k)
	}

	return list
}

func (r *NodeTokenRecord) Add(token string, name string) {
	r.m.Lock()
	defer r.m.Unlock()

	r.tokenRecord[token] = name
	r.nodeRecord[name] = token
}

func (r *NodeTokenRecord) RemoveByToken(token string) {
	r.m.Lock()
	defer r.m.Unlock()

	if name, find := r.tokenRecord[token]; find {
		delete(r.tokenRecord, token)
		delete(r.nodeRecord, name)
	}
}

func (r *NodeTokenRecord) RemoveByName(name string) {
	r.m.Lock()
	defer r.m.Unlock()

	if token, find := r.nodeRecord[name]; find {
		delete(r.tokenRecord, token)
		delete(r.nodeRecord, name)
	}
}

func (r *NodeTokenRecord) FindByToken(token string) (string, bool) {
	r.m.Lock()
	defer r.m.Unlock()

	val, find := r.tokenRecord[token]

	return val, find
}

func (r *NodeTokenRecord) FindByName(name string) (string, bool) {
	r.m.Lock()
	defer r.m.Unlock()

	val, find := r.nodeRecord[name]

	return val, find
}
