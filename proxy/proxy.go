package main

import "fmt"

// 一个交通工具的接口 （可以是 汽车，自行车，三轮车等，只要实现了 Drive方法）
type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven,call Drive")
}

// 一个司机对象
type Driver struct {
	Age int
}

type CarProxy struct {
	vehicle Vehicle
	driver  *Driver
}

// 新创建一个代理CarProxy，使用 Car 对象，来实现 Vehice 接口
func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

/*   CarProxy 代理着的 Drive 方法， 调用者不关心实现细节
实现层的Drive方法只关心 具体的Drive流程，其他的一些外部判断在代理层的 Drive 方法里面操作
*/
func (c *CarProxy) Drive() {
	//驾驶者的年龄判断，不应该在vehicle.Drive() 里面处理，所有在代理层Drive 方法里面操作
	if c.driver.Age >= 16 {
		c.vehicle.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}

func main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
	car2 := NewCarProxy(&Driver{22})
	car2.Drive()
}
