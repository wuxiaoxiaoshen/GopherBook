package v2

import "github.com/jinzhu/gorm"

type ShoppingCart struct {
	gorm.Model
	AccountID     uint
	ReceiptDateID uint
	OrderID       uint
	Order         Order
}
