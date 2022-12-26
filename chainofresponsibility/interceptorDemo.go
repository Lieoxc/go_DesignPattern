package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	功能： 对一个数进行过滤，打印出能够被 2 和 4 同时整除的数
*/
const abortIndex = 31

type intercept struct {
	interceptorHandlers []interceptorFunc
	index               int
	arg                 int
}
type interceptorFunc func(c *intercept)

func (h *intercept) Use(interceptor ...interceptorFunc) {
	h.interceptorHandlers = append(h.interceptorHandlers, interceptor...)
}
func (c *intercept) Next() {
	c.index++
	for c.index < len(c.interceptorHandlers) {
		c.interceptorHandlers[c.index](c)
		c.index++
	}
}
func (c *intercept) Abort() {
	c.index = abortIndex
}
func interceptorPrint1() interceptorFunc {
	return func(c *intercept) {
		if c.arg%2 != 0 {
			fmt.Println("arg:", c.arg, "intercepted by func interceptorPrint1")
			c.Abort()
			return
		}
		c.Next()
	}
}
func interceptorPrint2() interceptorFunc {
	return func(c *intercept) {
		if c.arg%4 != 0 {
			fmt.Println("arg:", c.arg, "intercepted by func interceptorPrint2")
			c.Abort()
			return
		}
		c.Next()
	}
}
func interceptorPrint3() interceptorFunc {
	return func(c *intercept) {
		fmt.Println("arg:", c.arg, "pass all interceptor func")
		c.Next()
	}
}
func New() *intercept {
	return &intercept{
		arg: 0,
	}
}
func main() {
	interceptor := New()
	interceptor.Use(interceptorPrint1(), interceptorPrint2(), interceptorPrint3())

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second * 3) //每十分钟更新一次配置信息
		runflag := true
		for runflag {
			select {
			case <-ctx.Done():
				runflag = false
				break
			case <-ticker.C:
				interceptor.arg++
				interceptor.index = 0
				interceptor.interceptorHandlers[0](interceptor)
			}
		}
		fmt.Println("gorutine exit")
		wg.Done()
	}()
	wg.Wait()
}
