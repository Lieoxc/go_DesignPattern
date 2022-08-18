package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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
func interceptoPrint() interceptorFunc {
	return func(c *intercept) {
		fmt.Println("interceptor Print xxx 0:", c.arg)
		c.Next()
	}
}
func interceptoPrint1() interceptorFunc {
	return func(c *intercept) {
		fmt.Println("interceptor Print xxx 1:", c.arg)
		c.Next()
	}
}
func interceptoPrint2() interceptorFunc {
	return func(c *intercept) {
		fmt.Println("interceptor Print XXX 2:", c.arg)
		c.Next()
	}
}
func main() {
	interceptor := &intercept{}
	interceptor.Use(interceptoPrint(), interceptoPrint1(), interceptoPrint2())

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second * 3) //每十分钟更新一次配置信息
		runflag := true
		count := 0
		for runflag {
			select {
			case <-ctx.Done():
				runflag = false
				break
			case <-ticker.C:
				count++
				//fmt.Println("xxxxx", count)
				interceptor.index = 0
				interceptor.interceptorHandlers[0](interceptor)
			}
		}
		fmt.Println("gorutine exit")
		wg.Done()
	}()
	wg.Wait()
}
