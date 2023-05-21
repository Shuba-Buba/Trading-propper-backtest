package exchange

type OrderType string
type OrderSide string
type OrderStatus string

const (
	MarketOrder OrderType = "market"
	LimitOrder  OrderType = "limit"
)

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

const (
	PendingPlacing    OrderStatus = "pending_placing"
	Placed            OrderStatus = "placed"
	PendingCancelling OrderStatus = "pending_cancelling"
	Cancelled         OrderStatus = "cancelled"
	Filled            OrderStatus = "filled"
	// and more
)

type OrderInfo struct {
	AvgPrice     float64
	Volume       float64
	FilledVolume float64
	Fee          float64
	Symbol       string
	Type         OrderType
	Side         OrderSide
	Status       OrderStatus
}

type OrderTracker struct {
	orderInfo map[string]*OrderInfo
}

func (t *OrderTracker) getOrderInfoPtr(orderId string) *OrderInfo {
	orderInfo, ok := t.orderInfo[orderId]
	if !ok {
		panic("Order with given orderId does not exist.")
	}
	return orderInfo
}

func (t *OrderTracker) deleteOrder(orderId string) {

}

func (t *OrderTracker) GetOrderInfo(orderId string) OrderInfo {
	return *t.getOrderInfoPtr(orderId)
}

func (t *OrderTracker) SetOrderInfo(orderId string, newOrderInfo OrderInfo) {
	*t.getOrderInfoPtr(orderId) = newOrderInfo
}

func (t *OrderTracker) SetStatus(orderId string, status OrderStatus) {
	t.getOrderInfoPtr(orderId).Status = status
}

func (t *OrderTracker) GetStatus(orderId string) OrderStatus {
	return t.getOrderInfoPtr(orderId).Status
}

func (t *OrderTracker) DeleteOrder(orderId string) {
	delete(t.orderInfo, orderId)
}
