package context

import (
	"errors"
	"sync"
)

// RuleCtx 存储解析出来的规则
type RuleCtx struct {
	ruleLock sync.Mutex
	rules    map[string]interface{}
}

func NewRuleCtx() *RuleCtx {
	return &RuleCtx{
		rules: make(map[string]interface{}),
	}
}

func (r *RuleCtx) AddRule(key string, method interface{}) {
	r.ruleLock.Lock()
	r.rules[key] = method
	r.ruleLock.Unlock()
}

func (r *RuleCtx) RemoveRule(key string) error {
	r.ruleLock.Lock()
	if _, ok := r.rules[key]; ok {
		delete(r.rules, key)
		r.ruleLock.Unlock()
		return nil
	}
	r.ruleLock.Unlock()
	return errors.New("not found method")
}

func (r *RuleCtx)GetRule(key string)error{
	r.ruleLock.Lock()
	if method, ok := r.rules[key]; ok {
		f := method.(func())
		f()
		r.ruleLock.Unlock()
		return nil
	}
	r.ruleLock.Unlock()
	return errors.New("not found match rule")
}
