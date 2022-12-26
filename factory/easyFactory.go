package main

import "fmt"

//********************简单工厂模式******************//
type Printer interface {
	Print(name string) string
}

type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return fmt.Sprintf("你好，%s", name)
}

type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return fmt.Sprintf("hello,%s", name)
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}
func runEasyFactoryCode() {
	printer := NewPrinter("cn")
	fmt.Println(printer.Print("小明"))

	printer = NewPrinter("en")
	fmt.Println(printer.Print("kobe"))
}

func main() {
	runEasyFactoryCode()
}
