package model

import "github.com/shopspring/decimal"

// dto (data transfer object) 数据传输对象
// [Demain 層]

// C2S_ProductCreate 新增商品
type C2S_ProductCreate struct {
	ProductName string          `json:"product_name"`
	Currency    string          `json:"currency"`
	BaseAmount  decimal.Decimal `json:"base_amount"`
}

func (c *C2S_ProductCreate) ToDomain() (*ProductCreateParams, error) {

	// 驗證用戶參數
	if err := c.Verify(); err != nil {
		return nil, err
	}

	// 將用戶參數轉換為領域對象
	return &ProductCreateParams{
		ProductName: c.ProductName,
		Currency:    c.Currency,
		BaseAmount:  c.BaseAmount,
	}, nil
}

// 驗證商品
func (c *C2S_ProductCreate) Verify() error {
	if len(c.ProductName) == 0 || len(c.Currency) == 0 {
		return Error_VerifyFailed
	}
	// 判斷金額是否 <= 0
	if !c.BaseAmount.GreaterThan(decimal.Zero) {
		return Error_VerifyFailed
	}

	return nil
}

// S2C_ProductCreate 新增商品回應
type S2C_ProductCreate struct {
	ProductID   int64           `json:"product_id"`
	ProductName string          `json:"product_name"`
	Currency    string          `json:"currency"`
	BaseAmount  decimal.Decimal `json:"base_amount"`
}