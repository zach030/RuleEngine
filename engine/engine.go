package engine

import "RuleEngine/constructor"

// 提供规则的执行
type RuleEngine struct {

}

func NewRuleEngine()*RuleEngine{
	return &RuleEngine{}
}

// 执行规则
func (r *RuleEngine) Execute(constructor constructor.RuleConstructor){

}

func (r *RuleEngine) ExecuteTest(key string,constructor constructor.RuleConstructor){
	constructor.DoRule(key)
}