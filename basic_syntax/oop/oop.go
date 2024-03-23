package ooop

// 有一句话：为了你Go语言的工作，你需要记住：interface是一种类型
// 订单接口
type Order interface {
	// 获取订单号
	GetNo() string

	// 获取订单金额
	GetAmount() uint32

	// 获取订单备注
	GetRemarks() string
}

// "继承"， Go推荐使用组合来达到继承的目的
type VeHicleOrder interface {
	Order
	GetCourseType() uint32
}

type CultureOrder struct {
}

func (co *CultureOrder) GetNo() string {

}

GetAmount() uint32

GetRemarks() string