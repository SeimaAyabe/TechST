package entity

// DeliveryDestination は「お届け先」テーブルのモデル
type DeliveryDestination struct {
	ID              int
	UserID          interface{}
	ZipCode         int    `form:"zipcode" binding:"required"`
	Address         string `form:"address" binding:"required"`
	TelephoneNumber string `form:"telephonenumber" binding:"required"`
	DefaultFlag     bool
}
