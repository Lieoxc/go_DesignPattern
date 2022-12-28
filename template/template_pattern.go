package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BankBusinessHandler interface {
	//排队
	TakeRowNumber()
	//等待
	WaitInHead()
	//处理业务
	HandleBusiness()
	// 对服务做出评价
	Commentate()
	// 钩子方法，判断是不是VIP， VIP不用等位
	CheckVipIdentity() bool
}
type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

// 模板方法，处理银行业务
func (b *BankBusinessExecutor) ExecuteBankBusiness() {

	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBusiness()
	b.handler.Commentate()
}

// 一个默认的业务实现单元
type DefaultBusinessHandler struct {
}

//排队取号的实现
func (*DefaultBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

// 排队的实现
func (*DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

// 服务评价的实现
func (*DefaultBusinessHandler) Commentate() {
	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultBusinessHandler) CheckVipIdentity() bool {
	// 留给具体实现类实现
	return false
}

type DepositBusinessHandler struct {
	*DefaultBusinessHandler //使用默认的业务实现  取号，排队，评价等业务
	userVIP                 bool
}

// 实际的交易业务由 DepositBusinessHandler自己实现
func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}

// DepositBusinessHandler 自己实现CheckVipIdentity，这样就覆盖掉 DefaultBusinessHandler 中对应的CheckVipIdentity方法
func (dh *DepositBusinessHandler) CheckVipIdentity() bool {
	return dh.userVIP
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

func main() {
	//初始化一个交易处理器
	dh := &DepositBusinessHandler{userVIP: false}
	// 使用 dh 这个交易处理器去 生成一个交易模板
	bbe := NewBankBusinessExecutor(dh)
	// 执行交易模板任务
	bbe.ExecuteBankBusiness()
}
