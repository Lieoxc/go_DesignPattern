package main

import "fmt"

type CalculatorHandler interface {
	Execute(cal *calculator) error
	SetNext(CalculatorHandler) CalculatorHandler
	Do(cal *calculator) error
}
type calculator struct {
	Result   int
	operandA int
	operandB int
}

type Next struct {
	nextHandler CalculatorHandler
}

func (n *Next) SetNext(handler CalculatorHandler) CalculatorHandler {
	n.nextHandler = handler
	return handler
}
func (n *Next) Execute(cal *calculator) (err error) {
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(cal); err != nil {
			return
		}

		return n.nextHandler.Execute(cal)
	}

	return
}

// 加法处理器
type AddCalculator struct {
	Next
}

func (add *AddCalculator) Do(cal *calculator) (err error) {
	fmt.Println("do AddCalculator")
	val := cal.operandA + cal.operandB
	cal.Result += val
	return nil
}

// 减法处理器
type DelCalculator struct {
	Next
}

func (del *DelCalculator) Do(cal *calculator) (err error) {
	fmt.Println("do DelCalculator")
	val := cal.operandA - cal.operandB
	cal.Result += val
	return nil
}

// StartHandler 不做操作，作为第一个Handler向下转发请求
// Go 语法限制，抽象公共逻辑到通用Handler后，并不能跟继承一样让公共方法调用不通子类的实现
type StartHandler struct {
	Next
}

// Do 空Handler的Do
func (h *StartHandler) Do(cal *calculator) (err error) {
	// 空Handler 这里什么也不做 只是载体 do nothing...
	fmt.Println("do StartHandler")
	return
}

func main() {
	numberCalculatorHandler := StartHandler{}
	cal := &calculator{
		Result:   0,
		operandA: 4,
		operandB: 2,
	}
	numberCalculatorHandler.SetNext(&AddCalculator{}). //注册加法处理器
								SetNext(&DelCalculator{}) //注册减法处理器
	// 还可以继续注册其他处理器

	// 执行上面设置好的业务流程
	if err := numberCalculatorHandler.Execute(cal); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success, cal result is:", cal.Result)
}
