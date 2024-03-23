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

// 实现基础订单
type BaseOrder struct {
}

// 获取订单号
func (o BaseOrder) GetNo() string {
	return "Base Order No."
}

// 获取订单金额
func (o BaseOrder) GetAmount() uint32 {
	return 1
}

// 获取订单备注
func (o BaseOrder) GetRemarks() string {
	return "Base Order remarks"
}

// 另一个实现基础订单
type AnotherBaseOrder struct {
}

// 获取订单号
func (o *AnotherBaseOrder) GetNo() string {
	return "Base Order No."
}

// 获取订单金额
func (o *AnotherBaseOrder) GetAmount() uint32 {
	return 1
}

// 获取订单备注
func (o *AnotherBaseOrder) GetRemarks() string {
	return "Base Order remarks"
}

// "继承"， Go推荐使用组合来达到继承的目的
// 车辆订单接口
type VehicleOrder interface {
	Order
	GetInfo() string
}

// "继承"， Go推荐使用组合来达到继承的目的
// 电车订单接口
type ElectricVehicleOrder interface {
	Order
	GetEVInfo() string
}

type NewVehicleOrder interface {
	VehicleOrder
	ElectricVehicleOrder
}

// 车辆订单的具体实现
type BaseVehicleOrder struct {
	BaseOrder   // 使用组合来“继承” BaseOrder 的方法
	VehicleInfo string
}

// 实现 VehicleOrder 接口的 GetInfo 方法
func (vo BaseVehicleOrder) GetInfo() string {
	return vo.VehicleInfo
}

// 电车订单的具体实现
type BaseElectricVehicleOrder struct {
	BaseOrder // 使用组合来“继承” BaseOrder 的方法
	EVInfo    string
}

// 实现 ElectricVehicleOrder 接口的 GetEVInfo 方法
func (evo BaseElectricVehicleOrder) GetEVInfo() string {
	return evo.EVInfo
}

// 新车辆订单的具体实现，组合了车辆订单和电车订单的功能
type BasicNewVehicleOrder struct {
	// 继承了 BaseVehicleOrder 的所有方法和属性
	BaseVehicleOrder
	// 继承了 BaseElectricVehicleOrder 的所有方法和属性
	BaseElectricVehicleOrder
}

// 重写 GetInfo 方法
func (bnvo *BasicNewVehicleOrder) GetInfo() string {
	// 自定义返回信息或调用嵌入的 VehicleOrderImpl 结构体的 GetInfo 方法
	return "New Basic Vehicle Order Info"
}

// 重写 GetEVInfo 方法
func (bnvo *BasicNewVehicleOrder) GetEVInfo() string {
	// 自定义返回信息或调用嵌入的 ElectricVehicleOrderImpl 结构体的 GetEVInfo 方法
	return "New Basic Electric Vehicle Order EVInfo"
}
