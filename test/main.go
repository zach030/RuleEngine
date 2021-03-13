package main

import (
	"RuleEngine/constructor"
	"RuleEngine/context"
	"RuleEngine/engine"
	"fmt"
	"strconv"
	"time"
)

func method1() {
	fmt.Println("low")
}

func method2() {
	fmt.Println("mid")
}

func method3() {
	fmt.Println("high")
}

func main() {
	// 配置规则信息
	ruleCtx := context.NewRuleCtx()
	ruleCtx.AddRule("10", method1)
	ruleCtx.AddRule("20", method2)
	ruleCtx.AddRule("30", method3)
	//
	ruleConstructor := constructor.NewRuleConstructor(ruleCtx)
	// 构建规则引擎
	ruleEngine := engine.NewRuleEngine()
	// 执行规则引擎

	// test
	for i := 0; i <= 100; i++ {
		time.Sleep(1000)
		fmt.Printf("now is:%d\t",i)
		ruleEngine.ExecuteTest(strconv.Itoa(i), ruleConstructor)
	}

}
