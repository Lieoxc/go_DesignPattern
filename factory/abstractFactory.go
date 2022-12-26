package main

import "fmt"

// 定义一个工厂，需要实现  ITelevision 和  IAirConditioner 子类
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}

// 每个子类包含了什么功能
type ITelevision interface {
	Watch()
}

type IAirConditioner interface {
	SetTemperature(int)
}

// 一个 HuaWei 工厂示例
type HuaWeiFactory struct{}

// 从工厂里面则创建负责实现每个子功能 的子类
func (hf *HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTV{}
}
func (hf *HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{}
}

// HuaWeiTV 负责实现 工厂中的 Watch 方法
type HuaWeiTV struct{}

func (ht *HuaWeiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

// HuaWeiAirConditioner 负责实现 工厂中的 SetTemperature 方法
type HuaWeiAirConditioner struct{}

func (ha *HuaWeiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("HuaWei AirConditioner set temperature to %d ℃\n", temp)
}

type MiFactory struct{}

func (mf *MiFactory) CreateTelevision() ITelevision {
	return &MiTV{}
}
func (mf *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}

type MiTV struct{}

func (mt *MiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

type MiAirConditioner struct{}

func (ma *MiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("Mi AirConditioner set temperature to %d ℃\n", temp)
}

func main() {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaWeiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(25)

	factory = &MiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(26)
}
