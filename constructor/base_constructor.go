package constructor

import (
	"RuleEngine/context"
	"fmt"
)

type BaseRuleConstructor struct {
	ruleCtx *context.RuleCtx
}

func NewRuleConstructor(ctx *context.RuleCtx) RuleConstructor {
	return &BaseRuleConstructor{
		ruleCtx: ctx,
	}
}

func (b *BaseRuleConstructor) ConstructRuleFromJson(rule string) {
	panic("implement me")
}

func (b *BaseRuleConstructor) ConstructRuleFromYaml() {
	panic("implement me")
}

func (b *BaseRuleConstructor) DoRule(key string) {
	err := b.ruleCtx.GetRule(key)
	if err != nil {
		fmt.Println("err:", err)
	}
}
