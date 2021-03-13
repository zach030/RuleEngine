package constructor

// RuleConstructor parse input string
type RuleConstructor interface {
	ConstructRuleFromJson(rule string)
	ConstructRuleFromYaml()
	DoRule(string2 string)
}



